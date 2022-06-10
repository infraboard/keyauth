// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: apps/namespace/pb/service.proto

package namespace

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
	QueryNamespace(ctx context.Context, in *QueryNamespaceRequest, opts ...grpc.CallOption) (*Set, error)
	DescribeNamespace(ctx context.Context, in *DescriptNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.namespace.Service/CreateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryNamespace(ctx context.Context, in *QueryNamespaceRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.namespace.Service/QueryNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeNamespace(ctx context.Context, in *DescriptNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.namespace.Service/DescribeNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.namespace.Service/DeleteNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*Namespace, error)
	QueryNamespace(context.Context, *QueryNamespaceRequest) (*Set, error)
	DescribeNamespace(context.Context, *DescriptNamespaceRequest) (*Namespace, error)
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*Namespace, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateNamespace(context.Context, *CreateNamespaceRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (UnimplementedServiceServer) QueryNamespace(context.Context, *QueryNamespaceRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNamespace not implemented")
}
func (UnimplementedServiceServer) DescribeNamespace(context.Context, *DescriptNamespaceRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeNamespace not implemented")
}
func (UnimplementedServiceServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.namespace.Service/CreateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.namespace.Service/QueryNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryNamespace(ctx, req.(*QueryNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescriptNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.namespace.Service/DescribeNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeNamespace(ctx, req.(*DescriptNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.namespace.Service/DeleteNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.keyauth.namespace.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNamespace",
			Handler:    _Service_CreateNamespace_Handler,
		},
		{
			MethodName: "QueryNamespace",
			Handler:    _Service_QueryNamespace_Handler,
		},
		{
			MethodName: "DescribeNamespace",
			Handler:    _Service_DescribeNamespace_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _Service_DeleteNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/namespace/pb/service.proto",
}