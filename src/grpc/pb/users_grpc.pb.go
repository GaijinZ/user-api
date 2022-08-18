// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: src/grpc/proto/users.proto

package pb

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

// UserApiClient is the client API for UserApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserApiClient interface {
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error)
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *AllUsersRequest, opts ...grpc.CallOption) (UserApi_GetUsersClient, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserRes, error)
}

type userApiClient struct {
	cc grpc.ClientConnInterface
}

func NewUserApiClient(cc grpc.ClientConnInterface) UserApiClient {
	return &userApiClient{cc}
}

func (c *userApiClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error) {
	out := new(CreateUserRes)
	err := c.cc.Invoke(ctx, "/pb.UserApi/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.UserApi/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) GetUsers(ctx context.Context, in *AllUsersRequest, opts ...grpc.CallOption) (UserApi_GetUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserApi_ServiceDesc.Streams[0], "/pb.UserApi/GetUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userApiGetUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserApi_GetUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userApiGetUsersClient struct {
	grpc.ClientStream
}

func (x *userApiGetUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userApiClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.UserApi/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApiClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserRes, error) {
	out := new(DeleteUserRes)
	err := c.cc.Invoke(ctx, "/pb.UserApi/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserApiServer is the server API for UserApi service.
// All implementations must embed UnimplementedUserApiServer
// for forward compatibility
type UserApiServer interface {
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error)
	GetUser(context.Context, *GetUserReq) (*User, error)
	GetUsers(*AllUsersRequest, UserApi_GetUsersServer) error
	UpdateUser(context.Context, *UpdateUserReq) (*User, error)
	DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserRes, error)
	mustEmbedUnimplementedUserApiServer()
}

// UnimplementedUserApiServer must be embedded to have forward compatible implementations.
type UnimplementedUserApiServer struct {
}

func (UnimplementedUserApiServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserApiServer) GetUser(context.Context, *GetUserReq) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserApiServer) GetUsers(*AllUsersRequest, UserApi_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserApiServer) UpdateUser(context.Context, *UpdateUserReq) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserApiServer) DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserApiServer) mustEmbedUnimplementedUserApiServer() {}

// UnsafeUserApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserApiServer will
// result in compilation errors.
type UnsafeUserApiServer interface {
	mustEmbedUnimplementedUserApiServer()
}

func RegisterUserApiServer(s grpc.ServiceRegistrar, srv UserApiServer) {
	s.RegisterService(&UserApi_ServiceDesc, srv)
}

func _UserApi_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserApi/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApi_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserApi/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApi_GetUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AllUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserApiServer).GetUsers(m, &userApiGetUsersServer{stream})
}

type UserApi_GetUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userApiGetUsersServer struct {
	grpc.ServerStream
}

func (x *userApiGetUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UserApi_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserApi/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApi_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApiServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserApi/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApiServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserApi_ServiceDesc is the grpc.ServiceDesc for UserApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserApi",
	HandlerType: (*UserApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserApi_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserApi_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserApi_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserApi_DeleteUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsers",
			Handler:       _UserApi_GetUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "src/grpc/proto/users.proto",
}
