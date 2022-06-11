// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.30
// source: config_server.proto

package pb

import (
	context "context"
	errors "errors"
	protojson "google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_config_server_proto struct{}

func (drpcEncoding_File_config_server_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_config_server_proto) MarshalAppend(buf []byte, msg drpc.Message) ([]byte, error) {
	return proto.MarshalOptions{}.MarshalAppend(buf, msg.(proto.Message))
}

func (drpcEncoding_File_config_server_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_config_server_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	return protojson.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_config_server_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return protojson.Unmarshal(buf, msg.(proto.Message))
}

type DRPCConfigServiceClient interface {
	DRPCConn() drpc.Conn

	GetConfig(ctx context.Context, in *ConfigRequest) (*ConfigResponse, error)
}

type drpcConfigServiceClient struct {
	cc drpc.Conn
}

func NewDRPCConfigServiceClient(cc drpc.Conn) DRPCConfigServiceClient {
	return &drpcConfigServiceClient{cc}
}

func (c *drpcConfigServiceClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcConfigServiceClient) GetConfig(ctx context.Context, in *ConfigRequest) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/ConfigService/GetConfig", drpcEncoding_File_config_server_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCConfigServiceServer interface {
	GetConfig(context.Context, *ConfigRequest) (*ConfigResponse, error)
}

type DRPCConfigServiceUnimplementedServer struct{}

func (s *DRPCConfigServiceUnimplementedServer) GetConfig(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCConfigServiceDescription struct{}

func (DRPCConfigServiceDescription) NumMethods() int { return 1 }

func (DRPCConfigServiceDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/ConfigService/GetConfig", drpcEncoding_File_config_server_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCConfigServiceServer).
					GetConfig(
						ctx,
						in1.(*ConfigRequest),
					)
			}, DRPCConfigServiceServer.GetConfig, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterConfigService(mux drpc.Mux, impl DRPCConfigServiceServer) error {
	return mux.Register(impl, DRPCConfigServiceDescription{})
}

type DRPCConfigService_GetConfigStream interface {
	drpc.Stream
	SendAndClose(*ConfigResponse) error
}

type drpcConfigService_GetConfigStream struct {
	drpc.Stream
}

func (x *drpcConfigService_GetConfigStream) SendAndClose(m *ConfigResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_config_server_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}