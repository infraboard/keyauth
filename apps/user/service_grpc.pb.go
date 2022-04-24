// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: apps/user/pb/service.proto

package user

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
	// 查询用户
	QueryAccount(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*Set, error)
	// 获取账号Profile
	DescribeAccount(ctx context.Context, in *DescribeAccountRequest, opts ...grpc.CallOption) (*User, error)
	// 创建用户
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*User, error)
	// 警用账号
	BlockAccount(ctx context.Context, in *BlockAccountRequest, opts ...grpc.CallOption) (*User, error)
	// 警用账号
	UnBlockAccount(ctx context.Context, in *UnBlockAccountRequest, opts ...grpc.CallOption) (*User, error)
	// DeleteAccount 删除用户
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*User, error)
	// 更新用户
	UpdateAccountProfile(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*User, error)
	// 修改用户密码
	UpdateAccountPassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*Password, error)
	// GeneratePassword 生成符合检测强度的随机密码
	GeneratePassword(ctx context.Context, in *GeneratePasswordRequest, opts ...grpc.CallOption) (*GeneratePasswordResponse, error)
	// 开启或关闭OTP
	UpdateOTPStatus(ctx context.Context, in *UpdateOTPStatusRequest, opts ...grpc.CallOption) (*User, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) QueryAccount(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.user.Service/QueryAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeAccount(ctx context.Context, in *DescribeAccountRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.user.Service/DescribeAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.user.Service/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BlockAccount(ctx context.Context, in *BlockAccountRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.user.Service/BlockAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UnBlockAccount(ctx context.Context, in *UnBlockAccountRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.user.Service/UnBlockAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.user.Service/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateAccountProfile(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.user.Service/UpdateAccountProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateAccountPassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*Password, error) {
	out := new(Password)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.user.Service/UpdateAccountPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GeneratePassword(ctx context.Context, in *GeneratePasswordRequest, opts ...grpc.CallOption) (*GeneratePasswordResponse, error) {
	out := new(GeneratePasswordResponse)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.user.Service/GeneratePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateOTPStatus(ctx context.Context, in *UpdateOTPStatusRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/infraboard.keyauth.user.Service/UpdateOTPStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// 查询用户
	QueryAccount(context.Context, *QueryAccountRequest) (*Set, error)
	// 获取账号Profile
	DescribeAccount(context.Context, *DescribeAccountRequest) (*User, error)
	// 创建用户
	CreateAccount(context.Context, *CreateAccountRequest) (*User, error)
	// 警用账号
	BlockAccount(context.Context, *BlockAccountRequest) (*User, error)
	// 警用账号
	UnBlockAccount(context.Context, *UnBlockAccountRequest) (*User, error)
	// DeleteAccount 删除用户
	DeleteAccount(context.Context, *DeleteAccountRequest) (*User, error)
	// 更新用户
	UpdateAccountProfile(context.Context, *UpdateAccountRequest) (*User, error)
	// 修改用户密码
	UpdateAccountPassword(context.Context, *UpdatePasswordRequest) (*Password, error)
	// GeneratePassword 生成符合检测强度的随机密码
	GeneratePassword(context.Context, *GeneratePasswordRequest) (*GeneratePasswordResponse, error)
	// 开启或关闭OTP
	UpdateOTPStatus(context.Context, *UpdateOTPStatusRequest) (*User, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) QueryAccount(context.Context, *QueryAccountRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAccount not implemented")
}
func (UnimplementedServiceServer) DescribeAccount(context.Context, *DescribeAccountRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAccount not implemented")
}
func (UnimplementedServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedServiceServer) BlockAccount(context.Context, *BlockAccountRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockAccount not implemented")
}
func (UnimplementedServiceServer) UnBlockAccount(context.Context, *UnBlockAccountRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockAccount not implemented")
}
func (UnimplementedServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedServiceServer) UpdateAccountProfile(context.Context, *UpdateAccountRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountProfile not implemented")
}
func (UnimplementedServiceServer) UpdateAccountPassword(context.Context, *UpdatePasswordRequest) (*Password, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountPassword not implemented")
}
func (UnimplementedServiceServer) GeneratePassword(context.Context, *GeneratePasswordRequest) (*GeneratePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePassword not implemented")
}
func (UnimplementedServiceServer) UpdateOTPStatus(context.Context, *UpdateOTPStatusRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOTPStatus not implemented")
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

func _Service_QueryAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.user.Service/QueryAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryAccount(ctx, req.(*QueryAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.user.Service/DescribeAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeAccount(ctx, req.(*DescribeAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.user.Service/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_BlockAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).BlockAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.user.Service/BlockAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).BlockAccount(ctx, req.(*BlockAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UnBlockAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnBlockAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UnBlockAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.user.Service/UnBlockAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UnBlockAccount(ctx, req.(*UnBlockAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.user.Service/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateAccountProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateAccountProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.user.Service/UpdateAccountProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateAccountProfile(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateAccountPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateAccountPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.user.Service/UpdateAccountPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateAccountPassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GeneratePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GeneratePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.user.Service/GeneratePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GeneratePassword(ctx, req.(*GeneratePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateOTPStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOTPStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateOTPStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.keyauth.user.Service/UpdateOTPStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateOTPStatus(ctx, req.(*UpdateOTPStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.keyauth.user.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryAccount",
			Handler:    _Service_QueryAccount_Handler,
		},
		{
			MethodName: "DescribeAccount",
			Handler:    _Service_DescribeAccount_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Service_CreateAccount_Handler,
		},
		{
			MethodName: "BlockAccount",
			Handler:    _Service_BlockAccount_Handler,
		},
		{
			MethodName: "UnBlockAccount",
			Handler:    _Service_UnBlockAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Service_DeleteAccount_Handler,
		},
		{
			MethodName: "UpdateAccountProfile",
			Handler:    _Service_UpdateAccountProfile_Handler,
		},
		{
			MethodName: "UpdateAccountPassword",
			Handler:    _Service_UpdateAccountPassword_Handler,
		},
		{
			MethodName: "GeneratePassword",
			Handler:    _Service_GeneratePassword_Handler,
		},
		{
			MethodName: "UpdateOTPStatus",
			Handler:    _Service_UpdateOTPStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/user/pb/service.proto",
}
