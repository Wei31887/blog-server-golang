// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: server_blog.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_server_blog_proto protoreflect.FileDescriptor

var file_server_blog_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x65,
	0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x57, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x42, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x0d, 0x5a, 0x0b, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_server_blog_proto_goTypes = []interface{}{
	(*LoginBloggerRequest)(nil),  // 0: pb.LoginBloggerRequest
	(*LoginBloggerResponse)(nil), // 1: pb.LoginBloggerResponse
}
var file_server_blog_proto_depIdxs = []int32{
	0, // 0: pb.BlogServer.LoginBlogger:input_type -> pb.LoginBloggerRequest
	1, // 1: pb.BlogServer.LoginBlogger:output_type -> pb.LoginBloggerResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_blog_proto_init() }
func file_server_blog_proto_init() {
	if File_server_blog_proto != nil {
		return
	}
	file_rpc_login_blogger_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_blog_proto_goTypes,
		DependencyIndexes: file_server_blog_proto_depIdxs,
	}.Build()
	File_server_blog_proto = out.File
	file_server_blog_proto_rawDesc = nil
	file_server_blog_proto_goTypes = nil
	file_server_blog_proto_depIdxs = nil
}
