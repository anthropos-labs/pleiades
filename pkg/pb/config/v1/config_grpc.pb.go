// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc config.2.0
// - protoc             (unknown)
// source: config/config/config.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go config.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigServiceClient interface {
	// Retrieve one or more configurations
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
	// Store a raft configuration
	PutRaftConfiguration(ctx context.Context, in *PutRaftConfigurationRequest, opts ...grpc.CallOption) (*PutRaftConfigurationResponse, error)
	GetRaftConfiguration(ctx context.Context, in *GetRaftConfigurationRequest, opts ...grpc.CallOption) (*GetRaftConfigurationResponse, error)
	ListRaftConfiguration(ctx context.Context, in *ListRaftConfigurationRequest, opts ...grpc.CallOption) (*ListRaftConfigurationResponse, error)
}

type configServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServiceClient(cc grpc.ClientConnInterface) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) PutRaftConfiguration(ctx context.Context, in *PutRaftConfigurationRequest, opts ...grpc.CallOption) (*PutRaftConfigurationResponse, error) {
	out := new(PutRaftConfigurationResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigService/PutRaftConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetRaftConfiguration(ctx context.Context, in *GetRaftConfigurationRequest, opts ...grpc.CallOption) (*GetRaftConfigurationResponse, error) {
	out := new(GetRaftConfigurationResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigService/GetRaftConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) ListRaftConfiguration(ctx context.Context, in *ListRaftConfigurationRequest, opts ...grpc.CallOption) (*ListRaftConfigurationResponse, error) {
	out := new(ListRaftConfigurationResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigService/ListRaftConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
// All implementations must embed UnimplementedConfigServiceServer
// for forward compatibility
type ConfigServiceServer interface {
	// Retrieve one or more configurations
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	// Store a raft configuration
	PutRaftConfiguration(context.Context, *PutRaftConfigurationRequest) (*PutRaftConfigurationResponse, error)
	GetRaftConfiguration(context.Context, *GetRaftConfigurationRequest) (*GetRaftConfigurationResponse, error)
	ListRaftConfiguration(context.Context, *ListRaftConfigurationRequest) (*ListRaftConfigurationResponse, error)
	mustEmbedUnimplementedConfigServiceServer()
}

// UnimplementedConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (UnimplementedConfigServiceServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedConfigServiceServer) PutRaftConfiguration(context.Context, *PutRaftConfigurationRequest) (*PutRaftConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutRaftConfiguration not implemented")
}
func (UnimplementedConfigServiceServer) GetRaftConfiguration(context.Context, *GetRaftConfigurationRequest) (*GetRaftConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRaftConfiguration not implemented")
}
func (UnimplementedConfigServiceServer) ListRaftConfiguration(context.Context, *ListRaftConfigurationRequest) (*ListRaftConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRaftConfiguration not implemented")
}
func (UnimplementedConfigServiceServer) mustEmbedUnimplementedConfigServiceServer() {}

// UnsafeConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServiceServer will
// result in compilation errors.
type UnsafeConfigServiceServer interface {
	mustEmbedUnimplementedConfigServiceServer()
}

func RegisterConfigServiceServer(s grpc.ServiceRegistrar, srv ConfigServiceServer) {
	s.RegisterService(&ConfigService_ServiceDesc, srv)
}

func _ConfigService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_PutRaftConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRaftConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).PutRaftConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/PutRaftConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).PutRaftConfiguration(ctx, req.(*PutRaftConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetRaftConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRaftConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetRaftConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/GetRaftConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetRaftConfiguration(ctx, req.(*GetRaftConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_ListRaftConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRaftConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).ListRaftConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/ListRaftConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).ListRaftConfiguration(ctx, req.(*ListRaftConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigService_ServiceDesc is the grpc.ServiceDesc for ConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _ConfigService_GetConfig_Handler,
		},
		{
			MethodName: "PutRaftConfiguration",
			Handler:    _ConfigService_PutRaftConfiguration_Handler,
		},
		{
			MethodName: "GetRaftConfiguration",
			Handler:    _ConfigService_GetRaftConfiguration_Handler,
		},
		{
			MethodName: "ListRaftConfiguration",
			Handler:    _ConfigService_ListRaftConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/config/config.proto",
}
