/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package blaze

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/lni/dragonboat/v3"
	"github.com/multiformats/go-multiaddr"
	"github.com/rs/zerolog"
)

type resultCode int

const (
	timeoutResultCode resultCode = iota
	completedResultCode
	terminatedResultCode
	rejectedResultCode
	droppedResultCode
	abortedResultCode
	committedResultCode
)

type OperationResult struct {
	Status resultCode
	Index  uint64
	Data   []byte
}

// MembershipEntry is the struct used to describe Raft cluster membership query
// results.
type MembershipEntry struct {
	// ConfigChangeId is the Raft entry index of the last applied membership
	// change entry.
	ConfigChangeId uint64
	// Nodes is a map of NodeID values to NodeHost Raft addresses for all regular
	// Raft nodes.
	Replicas map[uint64]string
	// Observers is a map of NodeID values to NodeHost Raft addresses for all
	// observers in the Raft cluster.
	Observers map[uint64]string
	// Witnesses is a map of NodeID values to NodeHost Raft addresses for all
	// witnesses in the Raft cluster.
	Witnesses map[uint64]string
	// Removed is a set of NodeID values that have been removed from the Raft
	// cluster. They are not allowed to be added back to the cluster.
	Removed map[uint64]struct{}
}

func requestResultAdapter(req dragonboat.RequestResult) resultCode {
	switch {
	case req.Timeout():
		return timeoutResultCode
	case req.Committed():
		return committedResultCode
	case req.Completed():
		return completedResultCode
	case req.Terminated():
		return terminatedResultCode
	case req.Aborted():
		return abortedResultCode
	case req.Rejected():
		return rejectedResultCode
	case req.Dropped():
		return droppedResultCode
	default:
		return droppedResultCode
	}
}

var (
	_ IShardManager = (*clusterManager)(nil)

	defaultTimeout = 3000 * time.Millisecond

	errNoConfigChangeId = errors.New("no config change id")
)

func newClusterManager(nodeHost *dragonboat.NodeHost, logger zerolog.Logger) *clusterManager {
	l := logger.With().Str("component", "cluster-manager").Logger()
	return &clusterManager{l, nodeHost}
}

type clusterManager struct {
	logger zerolog.Logger
	nh     *dragonboat.NodeHost
}

func (c *clusterManager) NewShard(cfg IClusterConfig) error {
	l := c.logger.With().Uint64("shard", cfg.ShardId()).Uint64("replica", cfg.ReplicaId()).Logger()
	l.Info().Msg("creating new shard")

	clusterConfig := cfg.Adapt()

	members := make(map[uint64]string)
	members[clusterConfig.ClusterID] = c.nh.RaftAddress()
	l.Debug().Str("raft-address", c.nh.RaftAddress()).Msg("adding self to members")

	switch cfg.StateMachineType() {
	case testStateMachineType:
		l.Info().Msg("creating test state machine")
		return c.nh.StartCluster(members, false, newTestStateMachine, clusterConfig)
	}

	l.Error().Msg("unknown state machine type")
	return ErrUnsupportedStateMachine
}

func (c *clusterManager) GetLeaderId(shardId uint64) (leader uint64, ok bool, err error) {
	c.logger.Info().Uint64("shard", shardId).Msg("getting leader id")
	return c.nh.GetLeaderID(shardId)
}

func (c *clusterManager) StopReplica(shardId uint64) (*OperationResult, error) {
	c.logger.Info().Uint64("shard", shardId).Msg("stopping replica")
	err := c.nh.StopCluster(shardId)
	if err != nil {
		c.logger.Error().Err(err).Uint64("shard", shardId).Msg("failed to stop replica")
		return nil, err
	}
	return &OperationResult{}, nil
}

func (c *clusterManager) DeleteReplica(cfg IClusterConfig, timeout time.Duration) error {
	l := c.logger.With().Uint64("shard", cfg.ShardId()).Uint64("replica", cfg.ReplicaId()).Logger()
	l.Info().Msg("deleting replica")

	if timeout == 0 {
		l.Debug().Msg("using default timeout")
		timeout = defaultTimeout
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	members, err := c.GetShardMembers(cfg.ShardId())
	if err != nil {
		l.Error().Err(err).Msg("failed to get shard members")
		return err
	}

	if members.ConfigChangeId == 0 {
		l.Error().Err(errNoConfigChangeId).Msg("failed to get config change id from shard members")
		return errNoConfigChangeId
	}

	err = c.nh.SyncRequestDeleteNode(ctx, cfg.ShardId(), cfg.ReplicaId(), members.ConfigChangeId)
	if err != nil {
		l.Error().Err(err).Msg("failed to delete replica")
		return err
	}
	return nil
}

func (c *clusterManager) GetShardMembers(shardId uint64) (*MembershipEntry, error) {
	l := c.logger.With().Uint64("shard", shardId).Logger()
	l.Debug().Msg("getting shard members")

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(defaultTimeout))
	defer cancel()

	membership, err := c.nh.SyncGetClusterMembership(ctx, shardId)
	if err != nil {
		l.Error().Err(err).Msg("failed to get shard membership")
		return nil, err
	}

	return &MembershipEntry{
		ConfigChangeId: membership.ConfigChangeID,
		Replicas:       membership.Nodes,
		Observers:      membership.Observers,
		Witnesses:      membership.Witnesses,
	}, nil
}

func (c *clusterManager) RemoveData(shardId, replicaId uint64) error {
	c.logger.Info().Uint64("shard", shardId).Uint64("replica", replicaId).Msg("removing data")

	err := c.nh.RemoveData(shardId, replicaId)
	if err != nil {
		c.logger.Error().Err(err).Uint64("shard", shardId).Uint64("replica", replicaId).Msg("failed to remove data")
	}
	return err
}

func (c *clusterManager) AddReplica(cfg IClusterConfig, newHost multiaddr.Multiaddr, timeout time.Duration) error {
	l := c.logger.With().Uint64("shard", cfg.ShardId()).Uint64("replica", cfg.ReplicaId()).Logger()
	l.Info().Msg("adding replica")

	if timeout == 0 {
		l.Debug().Msg("using default timeout")
		timeout = defaultTimeout
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	members, err := c.GetShardMembers(cfg.ShardId())
	if err != nil {
		l.Error().Err(err).Msg("failed to get shard members")
		return err
	}

	err = c.nh.SyncRequestAddNode(ctx, cfg.ShardId(), cfg.ReplicaId(), newHost.String(), members.ConfigChangeId)
	if err != nil {
		l.Error().Err(err).Msg("failed to add replica")
	}
	return err
}

func (c *clusterManager) AddShardObserver(cfg IClusterConfig, newHost multiaddr.Multiaddr, timeout time.Duration) error {
	l := c.logger.With().Uint64("shard", cfg.ShardId()).Uint64("replica", cfg.ReplicaId()).Logger()
	l.Info().Msg("adding shard observer")

	if timeout == 0 {
		l.Debug().Msg("using default timeout")
		timeout = defaultTimeout
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	members, err := c.GetShardMembers(cfg.ShardId())
	if err != nil {
		l.Error().Err(err).Msg("failed to get shard members")
		return err
	}

	err = c.nh.SyncRequestAddObserver(ctx, cfg.ShardId(), cfg.ReplicaId(), newHost.String(), members.ConfigChangeId)
	if err != nil {
		l.Error().Err(err).Msg("failed to add shard observer")
	}
	return err
}

func (c *clusterManager) AddShardWitness(cfg IClusterConfig, newHost multiaddr.Multiaddr, timeout time.Duration) error {
	l := c.logger.With().Uint64("shard", cfg.ShardId()).Uint64("replica", cfg.ReplicaId()).Logger()
	l.Info().Msg("adding shard observer")

	if timeout == 0 {
		l.Debug().Msg("using default timeout")
		timeout = defaultTimeout
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	members, err := c.GetShardMembers(cfg.ShardId())
	if err != nil {
		l.Error().Err(err).Msg("failed to get shard members")
		return err
	}

	err = c.nh.SyncRequestAddWitness(ctx, cfg.ShardId(), cfg.ReplicaId(), newHost.String(), members.ConfigChangeId)
	if err != nil {
		l.Error().Err(err).Msg("failed to add shard observer")
	}
	return err
}