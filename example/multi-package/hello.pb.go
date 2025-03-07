// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: hello.proto

package multi_package

import (
	context "context"
	bar "github.com/anushaunni/gripmock/example/multi-package/bar"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

var File_hello_proto protoreflect.FileDescriptor

var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x62, 0x61,
	0x72, 0x2f, 0x62, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x66, 0x6f, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2c, 0x0a, 0x08, 0x47, 0x72, 0x69, 0x70, 0x6d, 0x6f,
	0x63, 0x6b, 0x12, 0x20, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x08, 0x2e, 0x62, 0x61,
	0x72, 0x2e, 0x42, 0x61, 0x72, 0x1a, 0x0d, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x67, 0x72, 0x69,
	0x70, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x2d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_hello_proto_goTypes = []interface{}{
	(*bar.Bar)(nil),  // 0: bar.Bar
	(*Response)(nil), // 1: foo.Response
}
var file_hello_proto_depIdxs = []int32{
	0, // 0: multi_package.Gripmock.Greet:input_type -> bar.Bar
	1, // 1: multi_package.Gripmock.Greet:output_type -> foo.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	file_foo_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
	}.Build()
	File_hello_proto = out.File
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GripmockClient is the client API for Gripmock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GripmockClient interface {
	Greet(ctx context.Context, in *bar.Bar, opts ...grpc.CallOption) (*Response, error)
}

type gripmockClient struct {
	cc grpc.ClientConnInterface
}

func NewGripmockClient(cc grpc.ClientConnInterface) GripmockClient {
	return &gripmockClient{cc}
}

func (c *gripmockClient) Greet(ctx context.Context, in *bar.Bar, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/multi_package.Gripmock/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GripmockServer is the server API for Gripmock service.
type GripmockServer interface {
	Greet(context.Context, *bar.Bar) (*Response, error)
}

// UnimplementedGripmockServer can be embedded to have forward compatible implementations.
type UnimplementedGripmockServer struct {
}

func (*UnimplementedGripmockServer) Greet(context.Context, *bar.Bar) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGripmockServer(s *grpc.Server, srv GripmockServer) {
	s.RegisterService(&_Gripmock_serviceDesc, srv)
}

func _Gripmock_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bar.Bar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripmockServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multi_package.Gripmock/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripmockServer).Greet(ctx, req.(*bar.Bar))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gripmock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "multi_package.Gripmock",
	HandlerType: (*GripmockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Gripmock_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
