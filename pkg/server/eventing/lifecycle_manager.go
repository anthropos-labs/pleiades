/*
 * Copyright (c) 2022 Anthropos Labs, Inc.
 *
 * Licensed under the PolyForm Strict License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://gitlab.com/anthropos-labs/pleiades/-/blob/mainline/LICENSE
 */

package eventing

import (
	"context"

	raftv1 "a13s.io/api/raft/v1"
	"a13s.io/pleiades/pkg/fsm"
	"a13s.io/pleiades/pkg/messaging"
	"a13s.io/pleiades/pkg/server/runtime"
	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var (
	evSingleton *messaging.RaftEventHandler
)

type LifecycleManagerBuilderParams struct {
	fx.In

	StreamClient *messaging.EmbeddedMessagingStreamClient
	PubSubClient *messaging.EmbeddedMessagingPubSubClient
	ShardManager runtime.IShardManager
	RaftHost     runtime.IHost
	Logger       zerolog.Logger
}

type LifecycleManagerBuilderResults struct {
	fx.Out

	Runner *LifecycleManager
}

func NewLifecycleManager(lc fx.Lifecycle, params LifecycleManagerBuilderParams) LifecycleManagerBuilderResults {
	l := params.Logger.With().Str("component", "lifecycle-manager").Logger()

	store, err := fsm.NewSystemStore(l)
	if err != nil {
		l.Fatal().Err(err).Msg("can't create shard storage")
	}

	evSingleton = messaging.NewRaftEventHandler(params.PubSubClient, params.StreamClient, l)

	runner := &LifecycleManager{
		logger: l,
		store: store,
		shardManager: params.ShardManager,
		pubSubClient: params.PubSubClient,
		eventHandler: evSingleton,
		raftHost: params.RaftHost,
	}
	runner.registerCallbacks()

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			// we start the shards before we start the listener to prevent random startup issues
			err = runner.StartShards()
			if err != nil {
				l.Error().Err(err).Msg("can't ")
				return err
			}

			go evSingleton.Run()
			return nil
		},
		OnStop: func(_ context.Context) error {
			evSingleton.Stop()
			return runner.StopShards()
		},
	})

	return LifecycleManagerBuilderResults{Runner: runner}
}

type LifecycleManager struct {
	logger       zerolog.Logger
	store        *fsm.SystemStore
	eventHandler *messaging.RaftEventHandler
	shardManager runtime.IShardManager
	raftHost     runtime.IHost
	pubSubClient *messaging.EmbeddedMessagingPubSubClient
}

// StartShards will attempt to boot any replicas that were running on the node
func (l *LifecycleManager) StartShards() error {
	all, err := l.store.GetAll()
	if err != nil {
		if !errors.Is(err, fsm.ErrNoShards) {
			l.logger.Error().Err(err).Msg("failed to get all shards")
			return err
		}
	} else {
		l.logger.Info().Msg("no existing shards found")
		if len(all) == 0 {
			return nil
		}
	}

	raftAddr := l.raftHost.RaftAddress()

	// nb (sienna): I hate this, rewrite it later...
	for _, shard := range all {
		for replicaId, addr := range shard.GetReplicas() {
			if addr != raftAddr {
				continue
			}
			err = l.shardManager.StartReplica(&raftv1.StartReplicaRequest{
				ShardId:   shard.GetShardId(),
				ReplicaId: replicaId,
				Type:      shard.GetType(),
				Restart:   true,
			})
			if err != nil {
				l.logger.Error().
					Err(err).
					Uint64("shard-id", shard.GetShardId()).
					Uint64("replica-id", replicaId).
					Msg("can't start replica")
			}
		}
		for replicaId, addr := range shard.GetObservers() {
			if addr != raftAddr {
				continue
			}
			err = l.shardManager.StartReplicaObserver(&raftv1.StartReplicaObserverRequest{
				ShardId:   shard.GetShardId(),
				ReplicaId: replicaId,
				Type:      shard.GetType(),
				Restart:   true,
			})
			if err != nil {
				l.logger.Error().
					Err(err).
					Uint64("shard-id", shard.GetShardId()).
					Uint64("replica-id", replicaId).
					Msg("can't start replica observer")
			}
		}
	}

	return nil
}

// StopShards will attempt to stop any replicas that were running on the node
func (l *LifecycleManager) StopShards() error {
	all, err := l.store.GetAll()
	if !errors.Is(err, fsm.ErrNoShards) {
		l.logger.Error().Err(err).Msg("failed to get all shards")
		return err
	}

	raftAddr := l.raftHost.RaftAddress()

	for _, shard := range all {
		for replicaId, addr := range shard.GetReplicas() {
			if addr != raftAddr {
				continue
			}
			_, err = l.shardManager.StopReplica(shard.GetShardId(), replicaId)
			if err != nil {
				l.logger.Error().
					Err(err).
					Uint64("shard-id", shard.GetShardId()).
					Uint64("replica-id", replicaId).
					Msg("can't start replica")
			}
		}
	}

	err = l.store.Close()
	if err != nil {
		l.logger.Error().Err(err).Msg("can't safely close system config store")
	}

	return nil
}

func (l *LifecycleManager) registerCallbacks() {
	l.eventHandler.RegisterCallback("leader-update", raftv1.Event_EVENT_LEADER_UPDATED, l.handleLeaderUpdate)
}

func (l *LifecycleManager) handleLeaderUpdate(event *raftv1.RaftEvent) {
	l.logger.Debug().Interface("payload", event).Msg("leader update recieved")

	// safety check
	if event.Typ != raftv1.EventType_EVENT_TYPE_RAFT {
		l.logger.Error().Msg("event type mismatched")
		return
	}

	lu := event.GetLeaderUpdate()
	if lu == nil {
		l.logger.Error().Msg("leader event is empty")
		return
	}

	members, err := l.shardManager.GetShardMembers(lu.GetShardId())
	if err != nil {
		l.logger.Error().Err(err).Uint64("shard-id", lu.GetShardId()).Msg("can't get shard members")
		return
	}

	update := &raftv1.ShardState{
		LastUpdated:    event.GetTimestamp(),
		ShardId:        lu.GetShardId(),
		ConfigChangeId: members.ConfigChangeId,
		Replicas:       members.Replicas,
		Observers:      members.Observers,
		Witnesses:      members.Witnesses,
		Removed: func() map[uint64]string {
			m := make(map[uint64]string)
			for k := range members.Removed {
				m[k] = ""
			}
			return m
		}(),
		Type: raftv1.StateMachineType_STATE_MACHINE_TYPE_KV,
	}

	err = l.store.Put(update)
	if err != nil {
		l.logger.Error().Err(err).Msg("can't put the update")
	}

	return
}
