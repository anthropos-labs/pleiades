// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: api/v1/database/kv.proto

package database

import (
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

type EventPayload_MethodName int32

const (
	EventPayload_UNKNOWN EventPayload_MethodName = 0
)

// Enum value maps for EventPayload_MethodName.
var (
	EventPayload_MethodName_name = map[int32]string{
		0: "UNKNOWN",
	}
	EventPayload_MethodName_value = map[string]int32{
		"UNKNOWN": 0,
	}
)

func (x EventPayload_MethodName) Enum() *EventPayload_MethodName {
	p := new(EventPayload_MethodName)
	*p = x
	return p
}

func (x EventPayload_MethodName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventPayload_MethodName) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_database_kv_proto_enumTypes[0].Descriptor()
}

func (EventPayload_MethodName) Type() protoreflect.EnumType {
	return &file_api_v1_database_kv_proto_enumTypes[0]
}

func (x EventPayload_MethodName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventPayload_MethodName.Descriptor instead.
func (EventPayload_MethodName) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_database_kv_proto_rawDescGZIP(), []int{0, 0}
}

type Event_EventType int32

const (
	Event_GET    Event_EventType = 0
	Event_PUT    Event_EventType = 1
	Event_DELETE Event_EventType = 2
)

// Enum value maps for Event_EventType.
var (
	Event_EventType_name = map[int32]string{
		0: "GET",
		1: "PUT",
		2: "DELETE",
	}
	Event_EventType_value = map[string]int32{
		"GET":    0,
		"PUT":    1,
		"DELETE": 2,
	}
)

func (x Event_EventType) Enum() *Event_EventType {
	p := new(Event_EventType)
	*p = x
	return p
}

func (x Event_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_database_kv_proto_enumTypes[1].Descriptor()
}

func (Event_EventType) Type() protoreflect.EnumType {
	return &file_api_v1_database_kv_proto_enumTypes[1]
}

func (x Event_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_EventType.Descriptor instead.
func (Event_EventType) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_database_kv_proto_rawDescGZIP(), []int{3, 0}
}

type EventPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*EventPayload_KeyValueOperation
	Type   isEventPayload_Type     `protobuf_oneof:"Type"`
	Method EventPayload_MethodName `protobuf:"varint,2,opt,name=Method,proto3,enum=database.EventPayload_MethodName" json:"Method,omitempty"`
}

func (x *EventPayload) Reset() {
	*x = EventPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPayload) ProtoMessage() {}

func (x *EventPayload) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPayload.ProtoReflect.Descriptor instead.
func (*EventPayload) Descriptor() ([]byte, []int) {
	return file_api_v1_database_kv_proto_rawDescGZIP(), []int{0}
}

func (m *EventPayload) GetType() isEventPayload_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *EventPayload) GetKeyValueOperation() *KeyValueOperation {
	if x, ok := x.GetType().(*EventPayload_KeyValueOperation); ok {
		return x.KeyValueOperation
	}
	return nil
}

func (x *EventPayload) GetMethod() EventPayload_MethodName {
	if x != nil {
		return x.Method
	}
	return EventPayload_UNKNOWN
}

type isEventPayload_Type interface {
	isEventPayload_Type()
}

type EventPayload_KeyValueOperation struct {
	KeyValueOperation *KeyValueOperation `protobuf:"bytes,1,opt,name=KeyValueOperation,proto3,oneof"`
}

func (*EventPayload_KeyValueOperation) isEventPayload_Type() {}

type KeyValueOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=Session,proto3" json:"Session,omitempty"`
	Event   *Event   `protobuf:"bytes,2,opt,name=Event,proto3" json:"Event,omitempty"`
}

func (x *KeyValueOperation) Reset() {
	*x = KeyValueOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValueOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueOperation) ProtoMessage() {}

func (x *KeyValueOperation) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueOperation.ProtoReflect.Descriptor instead.
func (*KeyValueOperation) Descriptor() ([]byte, []int) {
	return file_api_v1_database_kv_proto_rawDescGZIP(), []int{1}
}

func (x *KeyValueOperation) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *KeyValueOperation) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key is the key in bytes. An empty key is not allowed.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// create_revision is the revision of last creation on this key.
	CreateRevision int64 `protobuf:"varint,2,opt,name=create_revision,json=createRevision,proto3" json:"create_revision,omitempty"`
	// mod_revision is the revision of last modification on this key.
	ModRevision int64 `protobuf:"varint,3,opt,name=mod_revision,json=modRevision,proto3" json:"mod_revision,omitempty"`
	// version is the version of the key. A deletion resets
	// the version to zero and any modification of the key
	// increases its version.
	Version int64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	// value is the value held by the key, in bytes.
	Value []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	// lease is the ID of the lease that attached to key.
	// When the attached lease expires, the key will be deleted.
	// If lease is 0, then no lease is attached to the key.
	Lease int64 `protobuf:"varint,6,opt,name=lease,proto3" json:"lease,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_kv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_kv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_api_v1_database_kv_proto_rawDescGZIP(), []int{2}
}

func (x *KeyValue) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *KeyValue) GetCreateRevision() int64 {
	if x != nil {
		return x.CreateRevision
	}
	return 0
}

func (x *KeyValue) GetModRevision() int64 {
	if x != nil {
		return x.ModRevision
	}
	return 0
}

func (x *KeyValue) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *KeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KeyValue) GetLease() int64 {
	if x != nil {
		return x.Lease
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type is the kind of event. If type is a PUT, it indicates
	// new data has been stored to the key. If type is a DELETE,
	// it indicates the key was deleted.
	Type Event_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=database.Event_EventType" json:"type,omitempty"`
	// kv holds the KeyValue for the event.
	// A PUT event contains current kv pair.
	// A PUT event with kv.Version=1 indicates the creation of a key.
	// A DELETE/EXPIRE event contains the deleted key with
	// its modification revision set to the revision of deletion.
	Kv *KeyValue `protobuf:"bytes,2,opt,name=kv,proto3" json:"kv,omitempty"`
	// prev_kv holds the key-value pair before the event happens.
	PrevKv *KeyValue `protobuf:"bytes,3,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_database_kv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_database_kv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_api_v1_database_kv_proto_rawDescGZIP(), []int{3}
}

func (x *Event) GetType() Event_EventType {
	if x != nil {
		return x.Type
	}
	return Event_GET
}

func (x *Event) GetKv() *KeyValue {
	if x != nil {
		return x.Kv
	}
	return nil
}

func (x *Event) GetPrevKv() *KeyValue {
	if x != nil {
		return x.PrevKv
	}
	return nil
}

var File_api_v1_database_kv_proto protoreflect.FileDescriptor

var file_api_v1_database_kv_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x4b, 0x0a, 0x11, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x11,
	0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x39, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x19, 0x0a, 0x0a,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x42, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x67, 0x0a, 0x11, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x08, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x02, 0x6b, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x02, 0x6b, 0x76, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x6b,
	0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x76, 0x4b, 0x76, 0x22, 0x29, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x42, 0x22,
	0x5a, 0x20, 0x61, 0x31, 0x33, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6c, 0x65, 0x69, 0x61, 0x64,
	0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_database_kv_proto_rawDescOnce sync.Once
	file_api_v1_database_kv_proto_rawDescData = file_api_v1_database_kv_proto_rawDesc
)

func file_api_v1_database_kv_proto_rawDescGZIP() []byte {
	file_api_v1_database_kv_proto_rawDescOnce.Do(func() {
		file_api_v1_database_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_database_kv_proto_rawDescData)
	})
	return file_api_v1_database_kv_proto_rawDescData
}

var file_api_v1_database_kv_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_v1_database_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_database_kv_proto_goTypes = []interface{}{
	(EventPayload_MethodName)(0), // 0: database.EventPayload.MethodName
	(Event_EventType)(0),         // 1: database.Event.EventType
	(*EventPayload)(nil),         // 2: database.EventPayload
	(*KeyValueOperation)(nil),    // 3: database.KeyValueOperation
	(*KeyValue)(nil),             // 4: database.KeyValue
	(*Event)(nil),                // 5: database.Event
	(*Session)(nil),              // 6: database.Session
}
var file_api_v1_database_kv_proto_depIdxs = []int32{
	3, // 0: database.EventPayload.KeyValueOperation:type_name -> database.KeyValueOperation
	0, // 1: database.EventPayload.Method:type_name -> database.EventPayload.MethodName
	6, // 2: database.KeyValueOperation.Session:type_name -> database.Session
	5, // 3: database.KeyValueOperation.Event:type_name -> database.Event
	1, // 4: database.Event.type:type_name -> database.Event.EventType
	4, // 5: database.Event.kv:type_name -> database.KeyValue
	4, // 6: database.Event.prev_kv:type_name -> database.KeyValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_api_v1_database_kv_proto_init() }
func file_api_v1_database_kv_proto_init() {
	if File_api_v1_database_kv_proto != nil {
		return
	}
	file_api_v1_database_session_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_database_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPayload); i {
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
		file_api_v1_database_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValueOperation); i {
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
		file_api_v1_database_kv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
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
		file_api_v1_database_kv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
	file_api_v1_database_kv_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EventPayload_KeyValueOperation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_database_kv_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_database_kv_proto_goTypes,
		DependencyIndexes: file_api_v1_database_kv_proto_depIdxs,
		EnumInfos:         file_api_v1_database_kv_proto_enumTypes,
		MessageInfos:      file_api_v1_database_kv_proto_msgTypes,
	}.Build()
	File_api_v1_database_kv_proto = out.File
	file_api_v1_database_kv_proto_rawDesc = nil
	file_api_v1_database_kv_proto_goTypes = nil
	file_api_v1_database_kv_proto_depIdxs = nil
}
