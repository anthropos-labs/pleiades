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

	"a13s.io/pleiades/pkg/conf"
	"github.com/lni/dragonboat/v3"
	dconfig "github.com/lni/dragonboat/v3/config"
	dlog "github.com/lni/dragonboat/v3/logger"
	"github.com/rs/zerolog"
)

var (
	_ INodeHost = (*Node)(nil)
)

func init() {
	dlog.SetLoggerFactory(conf.DragonboatLoggerFactory)
}

type Node struct {
	logger zerolog.Logger
	nh     *dragonboat.NodeHost

	notifyOnCommit bool
	clusterManager ICluster
	sessionManager ISession
	storeManager   IStore
}

// NewRaftControlNode creates a new Node instance.
func NewRaftControlNode(nodeHostConfig dconfig.NodeHostConfig, logger zerolog.Logger) (*Node, error) {
	l := logger.With().Str("component", "node").Logger()

	nh, err := dragonboat.NewNodeHost(nodeHostConfig)
	if err != nil {
		l.Error().Err(err).Msg("failed to create node host")
		return nil, err
	}

	node := &Node{logger: l, nh: nh}

	if nodeHostConfig.NotifyCommit {
		node.notifyOnCommit = true
	}

	return node, nil
}

// NewOrGetClusterManager creates a new ICluster instance or gets the existing one.
func (n *Node) NewOrGetClusterManager() (ICluster, error) {
	if n.clusterManager == nil {
		n.clusterManager = newClusterManager(n.logger, n.nh)
	}
	return n.clusterManager, nil
}

// NewOrGetSessionManager creates a new ISession instance or gets the existing one.
func (n *Node) NewOrGetSessionManager() (ISession, error) {
	if n.sessionManager == nil {
		n.sessionManager = newSessionManager(n.logger, n.nh)
	}
	return n.sessionManager, nil
}

// NewOrGetStoreManager creates a new IStore instance or gets the existing one.
func (n *Node) NewOrGetStoreManager() (IStore, error) {
	if n.storeManager == nil {
		n.storeManager = newStoreManager(n.logger, n.nh)
	}
	return n.storeManager, nil
}

func (n *Node) NotifyOnCommit() bool {
	return n.notifyOnCommit
}

func (n *Node) GetLeaderID(clusterID uint64) (uint64, bool, error) {
	return n.nh.GetLeaderID(clusterID)
}

func (n *Node) GetNodeUser(clusterID uint64) (dragonboat.INodeUser, error) {
	return n.nh.GetNodeUser(clusterID)
}

func (n *Node) ID() string {
	return n.nh.ID()
}

func (n *Node) RaftAddress() string {
	return n.nh.RaftAddress()
}

func (n *Node) RemoveData(clusterID uint64, nodeID uint64) error {
	return n.nh.RemoveData(clusterID, nodeID)
}

func (n *Node) RequestAddNode(clusterID uint64, nodeID uint64, target dragonboat.Target, configChangeIndex uint64, timeout time.Duration) (*dragonboat.RequestState, error) {
	return n.nh.RequestAddNode(clusterID, nodeID, target, configChangeIndex, timeout)
}

func (n *Node) RequestAddObserver(clusterID uint64, nodeID uint64, target dragonboat.Target, configChangeIndex uint64, timeout time.Duration) (*dragonboat.RequestState, error) {
	return n.nh.RequestAddObserver(clusterID, nodeID, target, configChangeIndex, timeout)
}

func (n *Node) RequestAddWitness(clusterID uint64, nodeID uint64, target dragonboat.Target, configChangeIndex uint64, timeout time.Duration) (*dragonboat.RequestState, error) {
	return n.nh.RequestAddWitness(clusterID, nodeID, target, configChangeIndex, timeout)
}

func (n *Node) RequestCompaction(clusterID uint64, nodeID uint64) (*dragonboat.SysOpState, error) {
	return n.nh.RequestCompaction(clusterID, nodeID)
}

func (n *Node) RequestDeleteNode(clusterID uint64, nodeID uint64, configChangeIndex uint64, timeout time.Duration) (*dragonboat.RequestState, error) {
	return n.nh.RequestDeleteNode(clusterID, nodeID, configChangeIndex, timeout)
}

func (n *Node) RequestLeaderTransfer(clusterID uint64, targetNodeID uint64) error {
	return n.nh.RequestLeaderTransfer(clusterID, targetNodeID)
}

func (n *Node) RequestSnapshot(clusterID uint64, opt dragonboat.SnapshotOption, timeout time.Duration) (*dragonboat.RequestState, error) {
	return n.nh.RequestSnapshot(clusterID, opt, timeout)
}

func (n *Node) Stop() {
	n.nh.Stop()
}

func (n *Node) StopNode(clusterID uint64, nodeID uint64) error {
	return n.nh.StopNode(clusterID, nodeID)
}

func (n *Node) SyncRemoveData(ctx context.Context, clusterID uint64, nodeID uint64) error {
	return n.nh.SyncRemoveData(ctx, clusterID, nodeID)
}

func (n *Node) SyncRequestAddNode(ctx context.Context, clusterID uint64, nodeID uint64, target string, configChangeIndex uint64) error {
	return n.nh.SyncRequestAddNode(ctx, clusterID, nodeID, target, configChangeIndex)
}

func (n *Node) SyncRequestAddObserver(ctx context.Context, clusterID uint64, nodeID uint64, target string, configChangeIndex uint64) error {
	return n.nh.SyncRequestAddObserver(ctx, clusterID, nodeID, target, configChangeIndex)
}

func (n *Node) SyncRequestAddWitness(ctx context.Context, clusterID uint64, nodeID uint64, target string, configChangeIndex uint64) error {
	return n.nh.SyncRequestAddWitness(ctx, clusterID, nodeID, target, configChangeIndex)
}

func (n *Node) SyncRequestDeleteNode(ctx context.Context, clusterID uint64, nodeID uint64, configChangeIndex uint64) error {
	return n.nh.SyncRequestDeleteNode(ctx, clusterID, nodeID, configChangeIndex)
}

func (n *Node) SyncRequestSnapshot(ctx context.Context, clusterID uint64, opt dragonboat.SnapshotOption) (uint64, error) {
	return n.nh.SyncRequestSnapshot(ctx, clusterID, opt)
}
