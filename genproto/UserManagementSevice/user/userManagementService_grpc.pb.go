// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: UserManagementSevice/userManagementService.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	UserManagementService_GetUserById_FullMethodName       = "/User.UserManagementService/GetUserById"
	UserManagementService_UpdateUserById_FullMethodName    = "/User.UserManagementService/UpdateUserById"
	UserManagementService_DeleteUserById_FullMethodName    = "/User.UserManagementService/DeleteUserById"
	UserManagementService_GetUserProfile_FullMethodName    = "/User.UserManagementService/GetUserProfile"
	UserManagementService_UpdateUserProfile_FullMethodName = "/User.UserManagementService/UpdateUserProfile"
)

// UserManagementServiceClient is the client API for UserManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagementServiceClient interface {
	GetUserById(ctx context.Context, in *IdUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	UpdateUserById(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	DeleteUserById(ctx context.Context, in *IdUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	GetUserProfile(ctx context.Context, in *IdUserRequest, opts ...grpc.CallOption) (*UserProfileResponse, error)
	UpdateUserProfile(ctx context.Context, in *UserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error)
}

type userManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementServiceClient(cc grpc.ClientConnInterface) UserManagementServiceClient {
	return &userManagementServiceClient{cc}
}

func (c *userManagementServiceClient) GetUserById(ctx context.Context, in *IdUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserManagementService_GetUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) UpdateUserById(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserManagementService_UpdateUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) DeleteUserById(ctx context.Context, in *IdUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserManagementService_DeleteUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) GetUserProfile(ctx context.Context, in *IdUserRequest, opts ...grpc.CallOption) (*UserProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserProfileResponse)
	err := c.cc.Invoke(ctx, UserManagementService_GetUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) UpdateUserProfile(ctx context.Context, in *UserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserProfileResponse)
	err := c.cc.Invoke(ctx, UserManagementService_UpdateUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServiceServer is the server API for UserManagementService service.
// All implementations must embed UnimplementedUserManagementServiceServer
// for forward compatibility
type UserManagementServiceServer interface {
	GetUserById(context.Context, *IdUserRequest) (*UserResponse, error)
	UpdateUserById(context.Context, *UpdateUserRequest) (*UserResponse, error)
	DeleteUserById(context.Context, *IdUserRequest) (*DeleteUserResponse, error)
	GetUserProfile(context.Context, *IdUserRequest) (*UserProfileResponse, error)
	UpdateUserProfile(context.Context, *UserProfileRequest) (*UserProfileResponse, error)
	mustEmbedUnimplementedUserManagementServiceServer()
}

// UnimplementedUserManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagementServiceServer struct {
}

func (UnimplementedUserManagementServiceServer) GetUserById(context.Context, *IdUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserManagementServiceServer) UpdateUserById(context.Context, *UpdateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserById not implemented")
}
func (UnimplementedUserManagementServiceServer) DeleteUserById(context.Context, *IdUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserById not implemented")
}
func (UnimplementedUserManagementServiceServer) GetUserProfile(context.Context, *IdUserRequest) (*UserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserManagementServiceServer) UpdateUserProfile(context.Context, *UserProfileRequest) (*UserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (UnimplementedUserManagementServiceServer) mustEmbedUnimplementedUserManagementServiceServer() {}

// UnsafeUserManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServiceServer will
// result in compilation errors.
type UnsafeUserManagementServiceServer interface {
	mustEmbedUnimplementedUserManagementServiceServer()
}

func RegisterUserManagementServiceServer(s grpc.ServiceRegistrar, srv UserManagementServiceServer) {
	s.RegisterService(&UserManagementService_ServiceDesc, srv)
}

func _UserManagementService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).GetUserById(ctx, req.(*IdUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_UpdateUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).UpdateUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_UpdateUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).UpdateUserById(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_DeleteUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).DeleteUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_DeleteUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).DeleteUserById(ctx, req.(*IdUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).GetUserProfile(ctx, req.(*IdUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_UpdateUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).UpdateUserProfile(ctx, req.(*UserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagementService_ServiceDesc is the grpc.ServiceDesc for UserManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "User.UserManagementService",
	HandlerType: (*UserManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserManagementService_GetUserById_Handler,
		},
		{
			MethodName: "UpdateUserById",
			Handler:    _UserManagementService_UpdateUserById_Handler,
		},
		{
			MethodName: "DeleteUserById",
			Handler:    _UserManagementService_DeleteUserById_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserManagementService_GetUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserManagementService_UpdateUserProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "UserManagementSevice/userManagementService.proto",
}
