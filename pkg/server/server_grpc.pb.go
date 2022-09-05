// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: pkg/server/server.proto

package server

import (
	raft "a13s.io/pleiades/api/v1/raft"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShardManagerClient is the client API for ShardManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShardManagerClient interface {
	AddReplica(ctx context.Context, in *raft.AddReplicaRequest, opts ...grpc.CallOption) (*raft.AddReplicaReply, error)
	AddShardObserver(ctx context.Context, in *raft.AddShardObserverRequest, opts ...grpc.CallOption) (*raft.AddShardObserverRequest, error)
	AddShardWitness(ctx context.Context, in *raft.AddShardWitnessRequest, opts ...grpc.CallOption) (*raft.AddShardWitnessRequest, error)
	DeleteReplica(ctx context.Context, in *raft.DeleteReplicaRequest, opts ...grpc.CallOption) (*raft.DeleteReplicaReply, error)
	GetLeaderId(ctx context.Context, in *raft.GetLeaderIdRequest, opts ...grpc.CallOption) (*raft.GetLeaderIdReply, error)
	GetShardMembers(ctx context.Context, in *raft.GetShardMembersRequest, opts ...grpc.CallOption) (*raft.GetShardMembersReply, error)
	NewShard(ctx context.Context, in *raft.NewShardRequest, opts ...grpc.CallOption) (*raft.NewShardReply, error)
	RemoveData(ctx context.Context, in *raft.RemoveDataRequest, opts ...grpc.CallOption) (*raft.RemoveDataReply, error)
	StopReplica(ctx context.Context, in *raft.StopReplicaRequest, opts ...grpc.CallOption) (*raft.StopReplicaReply, error)
}

type shardManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewShardManagerClient(cc grpc.ClientConnInterface) ShardManagerClient {
	return &shardManagerClient{cc}
}

func (c *shardManagerClient) AddReplica(ctx context.Context, in *raft.AddReplicaRequest, opts ...grpc.CallOption) (*raft.AddReplicaReply, error) {
	out := new(raft.AddReplicaReply)
	err := c.cc.Invoke(ctx, "/server.ShardManager/AddReplica", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardManagerClient) AddShardObserver(ctx context.Context, in *raft.AddShardObserverRequest, opts ...grpc.CallOption) (*raft.AddShardObserverRequest, error) {
	out := new(raft.AddShardObserverRequest)
	err := c.cc.Invoke(ctx, "/server.ShardManager/AddShardObserver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardManagerClient) AddShardWitness(ctx context.Context, in *raft.AddShardWitnessRequest, opts ...grpc.CallOption) (*raft.AddShardWitnessRequest, error) {
	out := new(raft.AddShardWitnessRequest)
	err := c.cc.Invoke(ctx, "/server.ShardManager/AddShardWitness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardManagerClient) DeleteReplica(ctx context.Context, in *raft.DeleteReplicaRequest, opts ...grpc.CallOption) (*raft.DeleteReplicaReply, error) {
	out := new(raft.DeleteReplicaReply)
	err := c.cc.Invoke(ctx, "/server.ShardManager/DeleteReplica", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardManagerClient) GetLeaderId(ctx context.Context, in *raft.GetLeaderIdRequest, opts ...grpc.CallOption) (*raft.GetLeaderIdReply, error) {
	out := new(raft.GetLeaderIdReply)
	err := c.cc.Invoke(ctx, "/server.ShardManager/GetLeaderId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardManagerClient) GetShardMembers(ctx context.Context, in *raft.GetShardMembersRequest, opts ...grpc.CallOption) (*raft.GetShardMembersReply, error) {
	out := new(raft.GetShardMembersReply)
	err := c.cc.Invoke(ctx, "/server.ShardManager/GetShardMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardManagerClient) NewShard(ctx context.Context, in *raft.NewShardRequest, opts ...grpc.CallOption) (*raft.NewShardReply, error) {
	out := new(raft.NewShardReply)
	err := c.cc.Invoke(ctx, "/server.ShardManager/NewShard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardManagerClient) RemoveData(ctx context.Context, in *raft.RemoveDataRequest, opts ...grpc.CallOption) (*raft.RemoveDataReply, error) {
	out := new(raft.RemoveDataReply)
	err := c.cc.Invoke(ctx, "/server.ShardManager/RemoveData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shardManagerClient) StopReplica(ctx context.Context, in *raft.StopReplicaRequest, opts ...grpc.CallOption) (*raft.StopReplicaReply, error) {
	out := new(raft.StopReplicaReply)
	err := c.cc.Invoke(ctx, "/server.ShardManager/StopReplica", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShardManagerServer is the server API for ShardManager service.
// All implementations must embed UnimplementedShardManagerServer
// for forward compatibility
type ShardManagerServer interface {
	AddReplica(context.Context, *raft.AddReplicaRequest) (*raft.AddReplicaReply, error)
	AddShardObserver(context.Context, *raft.AddShardObserverRequest) (*raft.AddShardObserverRequest, error)
	AddShardWitness(context.Context, *raft.AddShardWitnessRequest) (*raft.AddShardWitnessRequest, error)
	DeleteReplica(context.Context, *raft.DeleteReplicaRequest) (*raft.DeleteReplicaReply, error)
	GetLeaderId(context.Context, *raft.GetLeaderIdRequest) (*raft.GetLeaderIdReply, error)
	GetShardMembers(context.Context, *raft.GetShardMembersRequest) (*raft.GetShardMembersReply, error)
	NewShard(context.Context, *raft.NewShardRequest) (*raft.NewShardReply, error)
	RemoveData(context.Context, *raft.RemoveDataRequest) (*raft.RemoveDataReply, error)
	StopReplica(context.Context, *raft.StopReplicaRequest) (*raft.StopReplicaReply, error)
	mustEmbedUnimplementedShardManagerServer()
}

// UnimplementedShardManagerServer must be embedded to have forward compatible implementations.
type UnimplementedShardManagerServer struct {
}

func (UnimplementedShardManagerServer) AddReplica(context.Context, *raft.AddReplicaRequest) (*raft.AddReplicaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReplica not implemented")
}
func (UnimplementedShardManagerServer) AddShardObserver(context.Context, *raft.AddShardObserverRequest) (*raft.AddShardObserverRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShardObserver not implemented")
}
func (UnimplementedShardManagerServer) AddShardWitness(context.Context, *raft.AddShardWitnessRequest) (*raft.AddShardWitnessRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShardWitness not implemented")
}
func (UnimplementedShardManagerServer) DeleteReplica(context.Context, *raft.DeleteReplicaRequest) (*raft.DeleteReplicaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReplica not implemented")
}
func (UnimplementedShardManagerServer) GetLeaderId(context.Context, *raft.GetLeaderIdRequest) (*raft.GetLeaderIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderId not implemented")
}
func (UnimplementedShardManagerServer) GetShardMembers(context.Context, *raft.GetShardMembersRequest) (*raft.GetShardMembersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShardMembers not implemented")
}
func (UnimplementedShardManagerServer) NewShard(context.Context, *raft.NewShardRequest) (*raft.NewShardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewShard not implemented")
}
func (UnimplementedShardManagerServer) RemoveData(context.Context, *raft.RemoveDataRequest) (*raft.RemoveDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveData not implemented")
}
func (UnimplementedShardManagerServer) StopReplica(context.Context, *raft.StopReplicaRequest) (*raft.StopReplicaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopReplica not implemented")
}
func (UnimplementedShardManagerServer) mustEmbedUnimplementedShardManagerServer() {}

// UnsafeShardManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShardManagerServer will
// result in compilation errors.
type UnsafeShardManagerServer interface {
	mustEmbedUnimplementedShardManagerServer()
}

func RegisterShardManagerServer(s grpc.ServiceRegistrar, srv ShardManagerServer) {
	s.RegisterService(&ShardManager_ServiceDesc, srv)
}

func _ShardManager_AddReplica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.AddReplicaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardManagerServer).AddReplica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ShardManager/AddReplica",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardManagerServer).AddReplica(ctx, req.(*raft.AddReplicaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardManager_AddShardObserver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.AddShardObserverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardManagerServer).AddShardObserver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ShardManager/AddShardObserver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardManagerServer).AddShardObserver(ctx, req.(*raft.AddShardObserverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardManager_AddShardWitness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.AddShardWitnessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardManagerServer).AddShardWitness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ShardManager/AddShardWitness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardManagerServer).AddShardWitness(ctx, req.(*raft.AddShardWitnessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardManager_DeleteReplica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.DeleteReplicaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardManagerServer).DeleteReplica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ShardManager/DeleteReplica",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardManagerServer).DeleteReplica(ctx, req.(*raft.DeleteReplicaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardManager_GetLeaderId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.GetLeaderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardManagerServer).GetLeaderId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ShardManager/GetLeaderId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardManagerServer).GetLeaderId(ctx, req.(*raft.GetLeaderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardManager_GetShardMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.GetShardMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardManagerServer).GetShardMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ShardManager/GetShardMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardManagerServer).GetShardMembers(ctx, req.(*raft.GetShardMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardManager_NewShard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.NewShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardManagerServer).NewShard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ShardManager/NewShard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardManagerServer).NewShard(ctx, req.(*raft.NewShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardManager_RemoveData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.RemoveDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardManagerServer).RemoveData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ShardManager/RemoveData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardManagerServer).RemoveData(ctx, req.(*raft.RemoveDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShardManager_StopReplica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.StopReplicaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardManagerServer).StopReplica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.ShardManager/StopReplica",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardManagerServer).StopReplica(ctx, req.(*raft.StopReplicaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShardManager_ServiceDesc is the grpc.ServiceDesc for ShardManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShardManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.ShardManager",
	HandlerType: (*ShardManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddReplica",
			Handler:    _ShardManager_AddReplica_Handler,
		},
		{
			MethodName: "AddShardObserver",
			Handler:    _ShardManager_AddShardObserver_Handler,
		},
		{
			MethodName: "AddShardWitness",
			Handler:    _ShardManager_AddShardWitness_Handler,
		},
		{
			MethodName: "DeleteReplica",
			Handler:    _ShardManager_DeleteReplica_Handler,
		},
		{
			MethodName: "GetLeaderId",
			Handler:    _ShardManager_GetLeaderId_Handler,
		},
		{
			MethodName: "GetShardMembers",
			Handler:    _ShardManager_GetShardMembers_Handler,
		},
		{
			MethodName: "NewShard",
			Handler:    _ShardManager_NewShard_Handler,
		},
		{
			MethodName: "RemoveData",
			Handler:    _ShardManager_RemoveData_Handler,
		},
		{
			MethodName: "StopReplica",
			Handler:    _ShardManager_StopReplica_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/server/server.proto",
}

// RaftHostClient is the client API for RaftHost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftHostClient interface {
	Compact(ctx context.Context, in *raft.CompactRequest, opts ...grpc.CallOption) (*raft.CompactReply, error)
	GetHostConfig(ctx context.Context, in *raft.GetHostConfigRequest, opts ...grpc.CallOption) (*raft.GetHostConfigReply, error)
	LeaderTransfer(ctx context.Context, in *raft.LeaderTransferRequest, opts ...grpc.CallOption) (*raft.LeaderTransferReply, error)
	Snapshot(ctx context.Context, in *raft.SnapshotRequest, opts ...grpc.CallOption) (*raft.SnapshotReply, error)
	Stop(ctx context.Context, in *raft.StopRequest, opts ...grpc.CallOption) (*raft.StopReply, error)
}

type raftHostClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftHostClient(cc grpc.ClientConnInterface) RaftHostClient {
	return &raftHostClient{cc}
}

func (c *raftHostClient) Compact(ctx context.Context, in *raft.CompactRequest, opts ...grpc.CallOption) (*raft.CompactReply, error) {
	out := new(raft.CompactReply)
	err := c.cc.Invoke(ctx, "/server.RaftHost/Compact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftHostClient) GetHostConfig(ctx context.Context, in *raft.GetHostConfigRequest, opts ...grpc.CallOption) (*raft.GetHostConfigReply, error) {
	out := new(raft.GetHostConfigReply)
	err := c.cc.Invoke(ctx, "/server.RaftHost/GetHostConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftHostClient) LeaderTransfer(ctx context.Context, in *raft.LeaderTransferRequest, opts ...grpc.CallOption) (*raft.LeaderTransferReply, error) {
	out := new(raft.LeaderTransferReply)
	err := c.cc.Invoke(ctx, "/server.RaftHost/LeaderTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftHostClient) Snapshot(ctx context.Context, in *raft.SnapshotRequest, opts ...grpc.CallOption) (*raft.SnapshotReply, error) {
	out := new(raft.SnapshotReply)
	err := c.cc.Invoke(ctx, "/server.RaftHost/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftHostClient) Stop(ctx context.Context, in *raft.StopRequest, opts ...grpc.CallOption) (*raft.StopReply, error) {
	out := new(raft.StopReply)
	err := c.cc.Invoke(ctx, "/server.RaftHost/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftHostServer is the server API for RaftHost service.
// All implementations must embed UnimplementedRaftHostServer
// for forward compatibility
type RaftHostServer interface {
	Compact(context.Context, *raft.CompactRequest) (*raft.CompactReply, error)
	GetHostConfig(context.Context, *raft.GetHostConfigRequest) (*raft.GetHostConfigReply, error)
	LeaderTransfer(context.Context, *raft.LeaderTransferRequest) (*raft.LeaderTransferReply, error)
	Snapshot(context.Context, *raft.SnapshotRequest) (*raft.SnapshotReply, error)
	Stop(context.Context, *raft.StopRequest) (*raft.StopReply, error)
	mustEmbedUnimplementedRaftHostServer()
}

// UnimplementedRaftHostServer must be embedded to have forward compatible implementations.
type UnimplementedRaftHostServer struct {
}

func (UnimplementedRaftHostServer) Compact(context.Context, *raft.CompactRequest) (*raft.CompactReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compact not implemented")
}
func (UnimplementedRaftHostServer) GetHostConfig(context.Context, *raft.GetHostConfigRequest) (*raft.GetHostConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostConfig not implemented")
}
func (UnimplementedRaftHostServer) LeaderTransfer(context.Context, *raft.LeaderTransferRequest) (*raft.LeaderTransferReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaderTransfer not implemented")
}
func (UnimplementedRaftHostServer) Snapshot(context.Context, *raft.SnapshotRequest) (*raft.SnapshotReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}
func (UnimplementedRaftHostServer) Stop(context.Context, *raft.StopRequest) (*raft.StopReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedRaftHostServer) mustEmbedUnimplementedRaftHostServer() {}

// UnsafeRaftHostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftHostServer will
// result in compilation errors.
type UnsafeRaftHostServer interface {
	mustEmbedUnimplementedRaftHostServer()
}

func RegisterRaftHostServer(s grpc.ServiceRegistrar, srv RaftHostServer) {
	s.RegisterService(&RaftHost_ServiceDesc, srv)
}

func _RaftHost_Compact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.CompactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftHostServer).Compact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.RaftHost/Compact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftHostServer).Compact(ctx, req.(*raft.CompactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftHost_GetHostConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.GetHostConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftHostServer).GetHostConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.RaftHost/GetHostConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftHostServer).GetHostConfig(ctx, req.(*raft.GetHostConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftHost_LeaderTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.LeaderTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftHostServer).LeaderTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.RaftHost/LeaderTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftHostServer).LeaderTransfer(ctx, req.(*raft.LeaderTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftHost_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftHostServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.RaftHost/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftHostServer).Snapshot(ctx, req.(*raft.SnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftHost_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raft.StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftHostServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.RaftHost/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftHostServer).Stop(ctx, req.(*raft.StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RaftHost_ServiceDesc is the grpc.ServiceDesc for RaftHost service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftHost_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.RaftHost",
	HandlerType: (*RaftHostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compact",
			Handler:    _RaftHost_Compact_Handler,
		},
		{
			MethodName: "GetHostConfig",
			Handler:    _RaftHost_GetHostConfig_Handler,
		},
		{
			MethodName: "LeaderTransfer",
			Handler:    _RaftHost_LeaderTransfer_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _RaftHost_Snapshot_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _RaftHost_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/server/server.proto",
}
