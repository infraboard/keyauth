// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package verifycode

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// VerifyCodeServiceClient is the client API for VerifyCodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerifyCodeServiceClient interface {
	IssueCode(ctx context.Context, in *IssueCodeRequest, opts ...grpc.CallOption) (*IssueCodeResponse, error)
	CheckCode(ctx context.Context, in *CheckCodeRequest, opts ...grpc.CallOption) (*Code, error)
}

type verifyCodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVerifyCodeServiceClient(cc grpc.ClientConnInterface) VerifyCodeServiceClient {
	return &verifyCodeServiceClient{cc}
}

func (c *verifyCodeServiceClient) IssueCode(ctx context.Context, in *IssueCodeRequest, opts ...grpc.CallOption) (*IssueCodeResponse, error) {
	out := new(IssueCodeResponse)
	err := c.cc.Invoke(ctx, "/keyauth.verifycode.VerifyCodeService/IssueCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifyCodeServiceClient) CheckCode(ctx context.Context, in *CheckCodeRequest, opts ...grpc.CallOption) (*Code, error) {
	out := new(Code)
	err := c.cc.Invoke(ctx, "/keyauth.verifycode.VerifyCodeService/CheckCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifyCodeServiceServer is the server API for VerifyCodeService service.
// All implementations must embed UnimplementedVerifyCodeServiceServer
// for forward compatibility
type VerifyCodeServiceServer interface {
	IssueCode(context.Context, *IssueCodeRequest) (*IssueCodeResponse, error)
	CheckCode(context.Context, *CheckCodeRequest) (*Code, error)
	mustEmbedUnimplementedVerifyCodeServiceServer()
}

// UnimplementedVerifyCodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVerifyCodeServiceServer struct {
}

func (UnimplementedVerifyCodeServiceServer) IssueCode(context.Context, *IssueCodeRequest) (*IssueCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueCode not implemented")
}
func (UnimplementedVerifyCodeServiceServer) CheckCode(context.Context, *CheckCodeRequest) (*Code, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCode not implemented")
}
func (UnimplementedVerifyCodeServiceServer) mustEmbedUnimplementedVerifyCodeServiceServer() {}

// UnsafeVerifyCodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerifyCodeServiceServer will
// result in compilation errors.
type UnsafeVerifyCodeServiceServer interface {
	mustEmbedUnimplementedVerifyCodeServiceServer()
}

func RegisterVerifyCodeServiceServer(s *grpc.Server, srv VerifyCodeServiceServer) {
	s.RegisterService(&_VerifyCodeService_serviceDesc, srv)
}

func _VerifyCodeService_IssueCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyCodeServiceServer).IssueCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.verifycode.VerifyCodeService/IssueCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyCodeServiceServer).IssueCode(ctx, req.(*IssueCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifyCodeService_CheckCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifyCodeServiceServer).CheckCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.verifycode.VerifyCodeService/CheckCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifyCodeServiceServer).CheckCode(ctx, req.(*CheckCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerifyCodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keyauth.verifycode.VerifyCodeService",
	HandlerType: (*VerifyCodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueCode",
			Handler:    _VerifyCodeService_IssueCode_Handler,
		},
		{
			MethodName: "CheckCode",
			Handler:    _VerifyCodeService_CheckCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/verifycode/pb/service.proto",
}
