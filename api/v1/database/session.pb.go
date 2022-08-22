// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: api/v1/database/session.proto

package database

import (
	v1 "a13s.io/pleiades/api/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SessionPayload_MethodName int32

const (
	SessionPayload_NEW_SESSION   SessionPayload_MethodName = 0
	SessionPayload_CLOSE_SESSION SessionPayload_MethodName = 1
)

// Enum value maps for SessionPayload_MethodName.
var (
	SessionPayload_MethodName_name = map[int32]string{
		0: "NEW_SESSION",
		1: "CLOSE_SESSION",
	}
	SessionPayload_MethodName_value = map[string]int32{
		"NEW_SESSION":   0,
		"CLOSE_SESSION": 1,
	}
)

func (x SessionPayload_MethodName) Enum() *SessionPayload_MethodName {
	p := new(SessionPayload_MethodName)
	*p = x
	return p
}

func (x SessionPayload_MethodName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SessionPayload_MethodName) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_database_session_proto_enumTypes[0].Descriptor()
}

func (SessionPayload_MethodName) Type() protoreflect.EnumType {
	return &file_api_v1_database_session_proto_enumTypes[0]
}

func (x SessionPayload_MethodName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SessionPayload_MethodName.Descriptor instead.
func (SessionPayload_MethodName) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_database_session_proto_rawDescGZIP(), []int{0, 0}
}

type SessionPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*SessionPayload_NewSessionRequest
	//	*SessionPayload_NewSessionResponse
	//	*SessionPayload_CloseSessionRequest
	//	*SessionPayload_CloseSessionResponse
	//	*SessionPayload_Error
	Type   isSessionPayload_Type     `protobuf_oneof:"Type"`
	Method SessionPayload_MethodName `protobuf:"varint,6,opt,name=method,proto3,enum=database.SessionPayload_MethodName" json:"method,omitempty"`
}

func (x *SessionPayload) Reset() {
	*x = SessionPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionPayload) ProtoMessage() {}

func (x *SessionPayload) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionPayload.ProtoReflect.Descriptor instead.
func (*SessionPayload) Descriptor() ([]byte, []int) {
	return file_api_v1_database_session_proto_rawDescGZIP(), []int{0}
}

func (m *SessionPayload) GetType() isSessionPayload_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *SessionPayload) GetNewSessionRequest() *NewSessionRequest {
	if x, ok := x.GetType().(*SessionPayload_NewSessionRequest); ok {
		return x.NewSessionRequest
	}
	return nil
}

func (x *SessionPayload) GetNewSessionResponse() *NewSessionResponse {
	if x, ok := x.GetType().(*SessionPayload_NewSessionResponse); ok {
		return x.NewSessionResponse
	}
	return nil
}

func (x *SessionPayload) GetCloseSessionRequest() *CloseSessionRequest {
	if x, ok := x.GetType().(*SessionPayload_CloseSessionRequest); ok {
		return x.CloseSessionRequest
	}
	return nil
}

func (x *SessionPayload) GetCloseSessionResponse() *CloseSessionResponse {
	if x, ok := x.GetType().(*SessionPayload_CloseSessionResponse); ok {
		return x.CloseSessionResponse
	}
	return nil
}

func (x *SessionPayload) GetError() *v1.DBError {
	if x, ok := x.GetType().(*SessionPayload_Error); ok {
		return x.Error
	}
	return nil
}

func (x *SessionPayload) GetMethod() SessionPayload_MethodName {
	if x != nil {
		return x.Method
	}
	return SessionPayload_NEW_SESSION
}

type isSessionPayload_Type interface {
	isSessionPayload_Type()
}

type SessionPayload_NewSessionRequest struct {
	NewSessionRequest *NewSessionRequest `protobuf:"bytes,1,opt,name=NewSessionRequest,proto3,oneof"`
}

type SessionPayload_NewSessionResponse struct {
	NewSessionResponse *NewSessionResponse `protobuf:"bytes,2,opt,name=NewSessionResponse,proto3,oneof"`
}

type SessionPayload_CloseSessionRequest struct {
	CloseSessionRequest *CloseSessionRequest `protobuf:"bytes,3,opt,name=CloseSessionRequest,proto3,oneof"`
}

type SessionPayload_CloseSessionResponse struct {
	CloseSessionResponse *CloseSessionResponse `protobuf:"bytes,4,opt,name=CloseSessionResponse,proto3,oneof"`
}

type SessionPayload_Error struct {
	Error *v1.DBError `protobuf:"bytes,5,opt,name=Error,proto3,oneof"`
}

func (*SessionPayload_NewSessionRequest) isSessionPayload_Type() {}

func (*SessionPayload_NewSessionResponse) isSessionPayload_Type() {}

func (*SessionPayload_CloseSessionRequest) isSessionPayload_Type() {}

func (*SessionPayload_CloseSessionResponse) isSessionPayload_Type() {}

func (*SessionPayload_Error) isSessionPayload_Type() {}

type CloseSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Timeout int64    `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *CloseSessionRequest) Reset() {
	*x = CloseSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseSessionRequest) ProtoMessage() {}

func (x *CloseSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseSessionRequest.ProtoReflect.Descriptor instead.
func (*CloseSessionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_database_session_proto_rawDescGZIP(), []int{1}
}

func (x *CloseSessionRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *CloseSessionRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type CloseSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Timeout int64    `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *CloseSessionResponse) Reset() {
	*x = CloseSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseSessionResponse) ProtoMessage() {}

func (x *CloseSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseSessionResponse.ProtoReflect.Descriptor instead.
func (*CloseSessionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_database_session_proto_rawDescGZIP(), []int{2}
}

func (x *CloseSessionResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *CloseSessionResponse) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type ProposeSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Timeout int64    `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *ProposeSessionRequest) Reset() {
	*x = ProposeSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeSessionRequest) ProtoMessage() {}

func (x *ProposeSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeSessionRequest.ProtoReflect.Descriptor instead.
func (*ProposeSessionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_database_session_proto_rawDescGZIP(), []int{3}
}

func (x *ProposeSessionRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *ProposeSessionRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId   uint64 `protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ClientId    uint64 `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	SessionId   uint64 `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	RespondedTo uint64 `protobuf:"varint,4,opt,name=responded_to,json=respondedTo,proto3" json:"responded_to,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_api_v1_database_session_proto_rawDescGZIP(), []int{4}
}

func (x *Session) GetClusterId() uint64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *Session) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *Session) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *Session) GetRespondedTo() uint64 {
	if x != nil {
		return x.RespondedTo
	}
	return 0
}

type NewSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId uint64 `protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ClientId  uint64 `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *NewSessionRequest) Reset() {
	*x = NewSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_session_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSessionRequest) ProtoMessage() {}

func (x *NewSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_session_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSessionRequest.ProtoReflect.Descriptor instead.
func (*NewSessionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_database_session_proto_rawDescGZIP(), []int{5}
}

func (x *NewSessionRequest) GetClusterId() uint64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *NewSessionRequest) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type NewSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId uint64 `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *NewSessionResponse) Reset() {
	*x = NewSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_session_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewSessionResponse) ProtoMessage() {}

func (x *NewSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_session_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewSessionResponse.ProtoReflect.Descriptor instead.
func (*NewSessionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_database_session_proto_rawDescGZIP(), []int{6}
}

func (x *NewSessionResponse) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

var File_api_v1_database_session_proto protoreflect.FileDescriptor

var file_api_v1_database_session_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8,
	0x03, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x4b, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x11, 0x4e, 0x65, 0x77,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e,
	0x0a, 0x12, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x12, 0x4e, 0x65, 0x77, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x13, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x13, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x54, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x00, 0x52, 0x14, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x44, 0x42, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22,
	0x30, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x4e, 0x45, 0x57, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10,
	0x01, 0x42, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5c, 0x0a, 0x13, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x5d, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x5e, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2b, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x6f,
	0x22, 0x4f, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x33, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0x8f, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x4e, 0x65, 0x77,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x4e, 0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x22, 0x5a, 0x20, 0x61, 0x31, 0x33, 0x73,
	0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6c, 0x65, 0x69, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_database_session_proto_rawDescOnce sync.Once
	file_api_v1_database_session_proto_rawDescData = file_api_v1_database_session_proto_rawDesc
)

func file_api_v1_database_session_proto_rawDescGZIP() []byte {
	file_api_v1_database_session_proto_rawDescOnce.Do(func() {
		file_api_v1_database_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_database_session_proto_rawDescData)
	})
	return file_api_v1_database_session_proto_rawDescData
}

var file_api_v1_database_session_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_database_session_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_v1_database_session_proto_goTypes = []interface{}{
	(SessionPayload_MethodName)(0), // 0: database.SessionPayload.MethodName
	(*SessionPayload)(nil),         // 1: database.SessionPayload
	(*CloseSessionRequest)(nil),    // 2: database.CloseSessionRequest
	(*CloseSessionResponse)(nil),   // 3: database.CloseSessionResponse
	(*ProposeSessionRequest)(nil),  // 4: database.ProposeSessionRequest
	(*Session)(nil),                // 5: database.Session
	(*NewSessionRequest)(nil),      // 6: database.NewSessionRequest
	(*NewSessionResponse)(nil),     // 7: database.NewSessionResponse
	(*v1.DBError)(nil),             // 8: database.DBError
}
var file_api_v1_database_session_proto_depIdxs = []int32{
	6,  // 0: database.SessionPayload.NewSessionRequest:type_name -> database.NewSessionRequest
	7,  // 1: database.SessionPayload.NewSessionResponse:type_name -> database.NewSessionResponse
	2,  // 2: database.SessionPayload.CloseSessionRequest:type_name -> database.CloseSessionRequest
	3,  // 3: database.SessionPayload.CloseSessionResponse:type_name -> database.CloseSessionResponse
	8,  // 4: database.SessionPayload.Error:type_name -> database.DBError
	0,  // 5: database.SessionPayload.method:type_name -> database.SessionPayload.MethodName
	5,  // 6: database.CloseSessionRequest.session:type_name -> database.Session
	5,  // 7: database.CloseSessionResponse.session:type_name -> database.Session
	5,  // 8: database.ProposeSessionRequest.session:type_name -> database.Session
	6,  // 9: database.SessionService.NewSession:input_type -> database.NewSessionRequest
	5,  // 10: database.SessionService.CloseSession:input_type -> database.Session
	7,  // 11: database.SessionService.NewSession:output_type -> database.NewSessionResponse
	5,  // 12: database.SessionService.CloseSession:output_type -> database.Session
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_v1_database_session_proto_init() }
func file_api_v1_database_session_proto_init() {
	if File_api_v1_database_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_database_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_database_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseSessionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_database_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseSessionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_database_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeSessionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_database_session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_database_session_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSessionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_database_session_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewSessionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_v1_database_session_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SessionPayload_NewSessionRequest)(nil),
		(*SessionPayload_NewSessionResponse)(nil),
		(*SessionPayload_CloseSessionRequest)(nil),
		(*SessionPayload_CloseSessionResponse)(nil),
		(*SessionPayload_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_database_session_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_database_session_proto_goTypes,
		DependencyIndexes: file_api_v1_database_session_proto_depIdxs,
		EnumInfos:         file_api_v1_database_session_proto_enumTypes,
		MessageInfos:      file_api_v1_database_session_proto_msgTypes,
	}.Build()
	File_api_v1_database_session_proto = out.File
	file_api_v1_database_session_proto_rawDesc = nil
	file_api_v1_database_session_proto_goTypes = nil
	file_api_v1_database_session_proto_depIdxs = nil
}
