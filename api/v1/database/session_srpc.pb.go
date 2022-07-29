// Code generated by protoc-gen-srpc. DO NOT EDIT.
// protoc-gen-srpc version: v0.8.6
// source: api/v1/database/session.proto

package database

import (
	context "context"
	srpc "github.com/aperturerobotics/starpc/srpc"
)

type SRPCSessionServiceClient interface {
	SRPCClient() srpc.Client

	NewSession(ctx context.Context, in *NewSessionRequest) (*NewSessionResponse, error)
	CloseSession(ctx context.Context, in *Session) (*Session, error)
	ProposeSession(ctx context.Context, in *ProposeSessionRequest) (SRPCSessionService_ProposeSessionClient, error)
}

type srpcSessionServiceClient struct {
	cc srpc.Client
}

func NewSRPCSessionServiceClient(cc srpc.Client) SRPCSessionServiceClient {
	return &srpcSessionServiceClient{cc}
}

func (c *srpcSessionServiceClient) SRPCClient() srpc.Client { return c.cc }

func (c *srpcSessionServiceClient) NewSession(ctx context.Context, in *NewSessionRequest) (*NewSessionResponse, error) {
	out := new(NewSessionResponse)
	err := c.cc.Invoke(ctx, "database.SessionService", "NewSession", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srpcSessionServiceClient) CloseSession(ctx context.Context, in *Session) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "database.SessionService", "CloseSession", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srpcSessionServiceClient) ProposeSession(ctx context.Context, in *ProposeSessionRequest) (SRPCSessionService_ProposeSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, "database.SessionService", "ProposeSession", in)
	if err != nil {
		return nil, err
	}
	strm := &srpcSessionService_ProposeSessionClient{stream}
	if err := strm.CloseSend(); err != nil {
		return nil, err
	}
	return strm, nil
}

type SRPCSessionService_ProposeSessionClient interface {
	srpc.Stream
	Recv() (*IndexState, error)
	RecvTo(*IndexState) error
}

type srpcSessionService_ProposeSessionClient struct {
	srpc.Stream
}

func (x *srpcSessionService_ProposeSessionClient) Recv() (*IndexState, error) {
	m := new(IndexState)
	if err := x.MsgRecv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *srpcSessionService_ProposeSessionClient) RecvTo(m *IndexState) error {
	return x.MsgRecv(m)
}

type SRPCSessionServiceServer interface {
	NewSession(context.Context, *NewSessionRequest) (*NewSessionResponse, error)
	CloseSession(context.Context, *Session) (*Session, error)
	ProposeSession(*ProposeSessionRequest, SRPCSessionService_ProposeSessionStream) error
}

type SRPCSessionServiceUnimplementedServer struct{}

func (s *SRPCSessionServiceUnimplementedServer) NewSession(context.Context, *NewSessionRequest) (*NewSessionResponse, error) {
	return nil, srpc.ErrUnimplemented
}

func (s *SRPCSessionServiceUnimplementedServer) CloseSession(context.Context, *Session) (*Session, error) {
	return nil, srpc.ErrUnimplemented
}

func (s *SRPCSessionServiceUnimplementedServer) ProposeSession(*ProposeSessionRequest, SRPCSessionService_ProposeSessionStream) error {
	return srpc.ErrUnimplemented
}

const SRPCSessionServiceServiceID = "database.SessionService"

type SRPCSessionServiceHandler struct {
	impl SRPCSessionServiceServer
}

func (SRPCSessionServiceHandler) GetServiceID() string { return SRPCSessionServiceServiceID }

func (SRPCSessionServiceHandler) GetMethodIDs() []string {
	return []string{
		"NewSession",
		"CloseSession",
		"ProposeSession",
	}
}

func (d *SRPCSessionServiceHandler) InvokeMethod(
	serviceID, methodID string,
	strm srpc.Stream,
) (bool, error) {
	if serviceID != "" && serviceID != d.GetServiceID() {
		return false, nil
	}

	switch methodID {
	case "NewSession":
		return true, d.InvokeMethod_NewSession(d.impl, strm)
	case "CloseSession":
		return true, d.InvokeMethod_CloseSession(d.impl, strm)
	case "ProposeSession":
		return true, d.InvokeMethod_ProposeSession(d.impl, strm)
	default:
		return false, nil
	}
}

func (SRPCSessionServiceHandler) InvokeMethod_NewSession(impl SRPCSessionServiceServer, strm srpc.Stream) error {
	req := new(NewSessionRequest)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	out, err := impl.NewSession(strm.Context(), req)
	if err != nil {
		return err
	}
	return strm.MsgSend(out)
}

func (SRPCSessionServiceHandler) InvokeMethod_CloseSession(impl SRPCSessionServiceServer, strm srpc.Stream) error {
	req := new(Session)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	out, err := impl.CloseSession(strm.Context(), req)
	if err != nil {
		return err
	}
	return strm.MsgSend(out)
}

func (SRPCSessionServiceHandler) InvokeMethod_ProposeSession(impl SRPCSessionServiceServer, strm srpc.Stream) error {
	req := new(ProposeSessionRequest)
	if err := strm.MsgRecv(req); err != nil {
		return err
	}
	serverStrm := &srpcSessionService_ProposeSessionStream{strm}
	return impl.ProposeSession(req, serverStrm)
}

func SRPCRegisterSessionService(mux srpc.Mux, impl SRPCSessionServiceServer) error {
	return mux.Register(&SRPCSessionServiceHandler{impl: impl})
}

type SRPCSessionService_NewSessionStream interface {
	srpc.Stream
	SendAndClose(*NewSessionResponse) error
}

type srpcSessionService_NewSessionStream struct {
	srpc.Stream
}

func (x *srpcSessionService_NewSessionStream) SendAndClose(m *NewSessionResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type SRPCSessionService_CloseSessionStream interface {
	srpc.Stream
	SendAndClose(*Session) error
}

type srpcSessionService_CloseSessionStream struct {
	srpc.Stream
}

func (x *srpcSessionService_CloseSessionStream) SendAndClose(m *Session) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type SRPCSessionService_ProposeSessionStream interface {
	srpc.Stream
	Send(*IndexState) error
}

type srpcSessionService_ProposeSessionStream struct {
	srpc.Stream
}

func (x *srpcSessionService_ProposeSessionStream) Send(m *IndexState) error {
	return x.MsgSend(m)
}