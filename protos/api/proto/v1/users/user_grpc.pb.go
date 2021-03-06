// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: user.proto

package users

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

// UserDetailClient is the client API for UserDetail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDetailClient interface {
	GetUser(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailResponse, error)
	GetUsersList(ctx context.Context, in *ListOfUserDetailsRequest, opts ...grpc.CallOption) (*ListOfUserDetailsResponse, error)
}

type userDetailClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDetailClient(cc grpc.ClientConnInterface) UserDetailClient {
	return &userDetailClient{cc}
}

func (c *userDetailClient) GetUser(ctx context.Context, in *UserDetailRequest, opts ...grpc.CallOption) (*UserDetailResponse, error) {
	out := new(UserDetailResponse)
	err := c.cc.Invoke(ctx, "/UserDetail/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDetailClient) GetUsersList(ctx context.Context, in *ListOfUserDetailsRequest, opts ...grpc.CallOption) (*ListOfUserDetailsResponse, error) {
	out := new(ListOfUserDetailsResponse)
	err := c.cc.Invoke(ctx, "/UserDetail/GetUsersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDetailServer is the server API for UserDetail service.
// All implementations must embed UnimplementedUserDetailServer
// for forward compatibility
type UserDetailServer interface {
	GetUser(context.Context, *UserDetailRequest) (*UserDetailResponse, error)
	GetUsersList(context.Context, *ListOfUserDetailsRequest) (*ListOfUserDetailsResponse, error)
	mustEmbedUnimplementedUserDetailServer()
}

// UnimplementedUserDetailServer must be embedded to have forward compatible implementations.
type UnimplementedUserDetailServer struct {
}

func (UnimplementedUserDetailServer) GetUser(context.Context, *UserDetailRequest) (*UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserDetailServer) GetUsersList(context.Context, *ListOfUserDetailsRequest) (*ListOfUserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersList not implemented")
}
func (UnimplementedUserDetailServer) mustEmbedUnimplementedUserDetailServer() {}

// UnsafeUserDetailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDetailServer will
// result in compilation errors.
type UnsafeUserDetailServer interface {
	mustEmbedUnimplementedUserDetailServer()
}

func RegisterUserDetailServer(s grpc.ServiceRegistrar, srv UserDetailServer) {
	s.RegisterService(&UserDetail_ServiceDesc, srv)
}

func _UserDetail_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDetailServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserDetail/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDetailServer).GetUser(ctx, req.(*UserDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDetail_GetUsersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOfUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDetailServer).GetUsersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserDetail/GetUsersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDetailServer).GetUsersList(ctx, req.(*ListOfUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserDetail_ServiceDesc is the grpc.ServiceDesc for UserDetail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDetail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserDetail",
	HandlerType: (*UserDetailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserDetail_GetUser_Handler,
		},
		{
			MethodName: "GetUsersList",
			Handler:    _UserDetail_GetUsersList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
