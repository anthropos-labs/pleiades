// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: pkg/server/kv.proto

package server

import (
	database "a13s.io/pleiades/api/v1/database"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pkg_server_kv_proto protoreflect.FileDescriptor

var file_pkg_server_kv_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6b, 0x76, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1b, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb8, 0x01, 0x0a, 0x0e, 0x4b,
	0x56, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x61, 0x31, 0x33, 0x73, 0x2e, 0x69, 0x6f,
	0x2f, 0x70, 0x6c, 0x65, 0x69, 0x61, 0x64, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_server_kv_proto_goTypes = []interface{}{
	(*database.GetRequest)(nil),     // 0: database.GetRequest
	(*database.PutRequest)(nil),     // 1: database.PutRequest
	(*database.DeleteRequest)(nil),  // 2: database.DeleteRequest
	(*database.GetResponse)(nil),    // 3: database.GetResponse
	(*database.PutReply)(nil),       // 4: database.PutReply
	(*database.DeleteResponse)(nil), // 5: database.DeleteResponse
}
var file_pkg_server_kv_proto_depIdxs = []int32{
	0, // 0: server.KVStoreService.Get:input_type -> database.GetRequest
	1, // 1: server.KVStoreService.Put:input_type -> database.PutRequest
	2, // 2: server.KVStoreService.Delete:input_type -> database.DeleteRequest
	3, // 3: server.KVStoreService.Get:output_type -> database.GetResponse
	4, // 4: server.KVStoreService.Put:output_type -> database.PutReply
	5, // 5: server.KVStoreService.Delete:output_type -> database.DeleteResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_server_kv_proto_init() }
func file_pkg_server_kv_proto_init() {
	if File_pkg_server_kv_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_server_kv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_server_kv_proto_goTypes,
		DependencyIndexes: file_pkg_server_kv_proto_depIdxs,
	}.Build()
	File_pkg_server_kv_proto = out.File
	file_pkg_server_kv_proto_rawDesc = nil
	file_pkg_server_kv_proto_goTypes = nil
	file_pkg_server_kv_proto_depIdxs = nil
}