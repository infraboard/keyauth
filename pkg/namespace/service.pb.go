//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: pkg/namespace/pb/service.proto

package namespace

import (
	_ "github.com/infraboard/mcube/pb/http"
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

var File_pkg_namespace_pb_service_proto protoreflect.FileDescriptor

var file_pkg_namespace_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x1a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x85, 0x05, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x97, 0x01, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x29, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x3b, 0xa2, 0xa3, 0x8c, 0x4d, 0x36, 0x2a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x30, 0x01, 0x42, 0x13, 0x0a, 0x05,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x42, 0x10, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x48, 0x01, 0x12, 0x8d, 0x01, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x22, 0x39, 0xa2, 0xa3, 0x8c, 0x4d,
	0x34, 0x2a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x30, 0x01, 0x42, 0x13,
	0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x42, 0x0e, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x48, 0x01, 0x12, 0x98, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x38, 0xa2, 0xa3, 0x8c, 0x4d, 0x33, 0x2a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x30, 0x01, 0x42, 0x13, 0x0a, 0x05, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x12, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42,
	0x0d, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x03, 0x67, 0x65, 0x74, 0x48, 0x01,
	0x12, 0xab, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x4f, 0xa2,
	0xa3, 0x8c, 0x4d, 0x4a, 0x1a, 0x0f, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x2f, 0x3a, 0x69, 0x64, 0x22, 0x03, 0x47, 0x45, 0x54, 0x2a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x30, 0x01, 0x42, 0x13, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x12, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x10, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_namespace_pb_service_proto_goTypes = []interface{}{
	(*CreateNamespaceRequest)(nil),   // 0: keyauth.namespace.CreateNamespaceRequest
	(*QueryNamespaceRequest)(nil),    // 1: keyauth.namespace.QueryNamespaceRequest
	(*DescriptNamespaceRequest)(nil), // 2: keyauth.namespace.DescriptNamespaceRequest
	(*DeleteNamespaceRequest)(nil),   // 3: keyauth.namespace.DeleteNamespaceRequest
	(*Namespace)(nil),                // 4: keyauth.namespace.Namespace
	(*Set)(nil),                      // 5: keyauth.namespace.Set
}
var file_pkg_namespace_pb_service_proto_depIdxs = []int32{
	0, // 0: keyauth.namespace.NamespaceService.CreateNamespace:input_type -> keyauth.namespace.CreateNamespaceRequest
	1, // 1: keyauth.namespace.NamespaceService.QueryNamespace:input_type -> keyauth.namespace.QueryNamespaceRequest
	2, // 2: keyauth.namespace.NamespaceService.DescribeNamespace:input_type -> keyauth.namespace.DescriptNamespaceRequest
	3, // 3: keyauth.namespace.NamespaceService.DeleteNamespace:input_type -> keyauth.namespace.DeleteNamespaceRequest
	4, // 4: keyauth.namespace.NamespaceService.CreateNamespace:output_type -> keyauth.namespace.Namespace
	5, // 5: keyauth.namespace.NamespaceService.QueryNamespace:output_type -> keyauth.namespace.Set
	4, // 6: keyauth.namespace.NamespaceService.DescribeNamespace:output_type -> keyauth.namespace.Namespace
	4, // 7: keyauth.namespace.NamespaceService.DeleteNamespace:output_type -> keyauth.namespace.Namespace
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_namespace_pb_service_proto_init() }
func file_pkg_namespace_pb_service_proto_init() {
	if File_pkg_namespace_pb_service_proto != nil {
		return
	}
	file_pkg_namespace_pb_namespace_proto_init()
	file_pkg_namespace_pb_request_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_namespace_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_namespace_pb_service_proto_goTypes,
		DependencyIndexes: file_pkg_namespace_pb_service_proto_depIdxs,
	}.Build()
	File_pkg_namespace_pb_service_proto = out.File
	file_pkg_namespace_pb_service_proto_rawDesc = nil
	file_pkg_namespace_pb_service_proto_goTypes = nil
	file_pkg_namespace_pb_service_proto_depIdxs = nil
}
