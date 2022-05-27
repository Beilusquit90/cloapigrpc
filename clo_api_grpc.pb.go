// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: clo_api.proto

package cloapi_grpc

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

// CloapiRPCClient is the client API for CloapiRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloapiRPCClient interface {
	// For delete Whitepages From All cloaks
	CreateWhitePage(ctx context.Context, in *FtpPage, opts ...grpc.CallOption) (*GRPCResponce, error)
	DeleteWhitePage(ctx context.Context, in *FtpPage, opts ...grpc.CallOption) (*GRPCResponce, error)
	DeleteAllWhitePages(ctx context.Context, in *FtpPage, opts ...grpc.CallOption) (*GRPCResponce, error)
	GetWhitePage(ctx context.Context, in *FtpPage, opts ...grpc.CallOption) (*FtpPage, error)
}

type cloapiRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewCloapiRPCClient(cc grpc.ClientConnInterface) CloapiRPCClient {
	return &cloapiRPCClient{cc}
}

func (c *cloapiRPCClient) CreateWhitePage(ctx context.Context, in *FtpPage, opts ...grpc.CallOption) (*GRPCResponce, error) {
	out := new(GRPCResponce)
	err := c.cc.Invoke(ctx, "/cloapi_grpc.CloapiRPC/CreateWhitePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloapiRPCClient) DeleteWhitePage(ctx context.Context, in *FtpPage, opts ...grpc.CallOption) (*GRPCResponce, error) {
	out := new(GRPCResponce)
	err := c.cc.Invoke(ctx, "/cloapi_grpc.CloapiRPC/DeleteWhitePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloapiRPCClient) DeleteAllWhitePages(ctx context.Context, in *FtpPage, opts ...grpc.CallOption) (*GRPCResponce, error) {
	out := new(GRPCResponce)
	err := c.cc.Invoke(ctx, "/cloapi_grpc.CloapiRPC/DeleteAllWhitePages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloapiRPCClient) GetWhitePage(ctx context.Context, in *FtpPage, opts ...grpc.CallOption) (*FtpPage, error) {
	out := new(FtpPage)
	err := c.cc.Invoke(ctx, "/cloapi_grpc.CloapiRPC/GetWhitePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloapiRPCServer is the server API for CloapiRPC service.
// All implementations must embed UnimplementedCloapiRPCServer
// for forward compatibility
type CloapiRPCServer interface {
	// For delete Whitepages From All cloaks
	CreateWhitePage(context.Context, *FtpPage) (*GRPCResponce, error)
	DeleteWhitePage(context.Context, *FtpPage) (*GRPCResponce, error)
	DeleteAllWhitePages(context.Context, *FtpPage) (*GRPCResponce, error)
	GetWhitePage(context.Context, *FtpPage) (*FtpPage, error)
	mustEmbedUnimplementedCloapiRPCServer()
}

// UnimplementedCloapiRPCServer must be embedded to have forward compatible implementations.
type UnimplementedCloapiRPCServer struct {
}

func (UnimplementedCloapiRPCServer) CreateWhitePage(context.Context, *FtpPage) (*GRPCResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWhitePage not implemented")
}
func (UnimplementedCloapiRPCServer) DeleteWhitePage(context.Context, *FtpPage) (*GRPCResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWhitePage not implemented")
}
func (UnimplementedCloapiRPCServer) DeleteAllWhitePages(context.Context, *FtpPage) (*GRPCResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllWhitePages not implemented")
}
func (UnimplementedCloapiRPCServer) GetWhitePage(context.Context, *FtpPage) (*FtpPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWhitePage not implemented")
}
func (UnimplementedCloapiRPCServer) mustEmbedUnimplementedCloapiRPCServer() {}

// UnsafeCloapiRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloapiRPCServer will
// result in compilation errors.
type UnsafeCloapiRPCServer interface {
	mustEmbedUnimplementedCloapiRPCServer()
}

func RegisterCloapiRPCServer(s grpc.ServiceRegistrar, srv CloapiRPCServer) {
	s.RegisterService(&CloapiRPC_ServiceDesc, srv)
}

func _CloapiRPC_CreateWhitePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FtpPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloapiRPCServer).CreateWhitePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloapi_grpc.CloapiRPC/CreateWhitePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloapiRPCServer).CreateWhitePage(ctx, req.(*FtpPage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloapiRPC_DeleteWhitePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FtpPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloapiRPCServer).DeleteWhitePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloapi_grpc.CloapiRPC/DeleteWhitePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloapiRPCServer).DeleteWhitePage(ctx, req.(*FtpPage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloapiRPC_DeleteAllWhitePages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FtpPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloapiRPCServer).DeleteAllWhitePages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloapi_grpc.CloapiRPC/DeleteAllWhitePages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloapiRPCServer).DeleteAllWhitePages(ctx, req.(*FtpPage))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloapiRPC_GetWhitePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FtpPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloapiRPCServer).GetWhitePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloapi_grpc.CloapiRPC/GetWhitePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloapiRPCServer).GetWhitePage(ctx, req.(*FtpPage))
	}
	return interceptor(ctx, in, info, handler)
}

// CloapiRPC_ServiceDesc is the grpc.ServiceDesc for CloapiRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloapiRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloapi_grpc.CloapiRPC",
	HandlerType: (*CloapiRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWhitePage",
			Handler:    _CloapiRPC_CreateWhitePage_Handler,
		},
		{
			MethodName: "DeleteWhitePage",
			Handler:    _CloapiRPC_DeleteWhitePage_Handler,
		},
		{
			MethodName: "DeleteAllWhitePages",
			Handler:    _CloapiRPC_DeleteAllWhitePages_Handler,
		},
		{
			MethodName: "GetWhitePage",
			Handler:    _CloapiRPC_GetWhitePage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clo_api.proto",
}
