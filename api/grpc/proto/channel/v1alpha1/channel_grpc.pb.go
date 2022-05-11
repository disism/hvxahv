// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/channel/v1alpha1/channel.proto

package v1alpha

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

// ChannelClient is the client API for Channel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelClient interface {
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error)
	GetChannelsByAccountID(ctx context.Context, in *GetChannelsByAccountIDRequest, opts ...grpc.CallOption) (*GetChannelsByAccountIDResponse, error)
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error)
	DeleteAllChannelsByAccountID(ctx context.Context, in *DeleteAllChannelsByAccountIDRequest, opts ...grpc.CallOption) (*DeleteAllChannelsByAccountIDResponse, error)
}

type channelClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelClient(cc grpc.ClientConnInterface) ChannelClient {
	return &channelClient{cc}
}

func (c *channelClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error) {
	out := new(CreateChannelResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Channel/CreateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) GetChannelsByAccountID(ctx context.Context, in *GetChannelsByAccountIDRequest, opts ...grpc.CallOption) (*GetChannelsByAccountIDResponse, error) {
	out := new(GetChannelsByAccountIDResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Channel/GetChannelsByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*DeleteChannelResponse, error) {
	out := new(DeleteChannelResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Channel/DeleteChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelClient) DeleteAllChannelsByAccountID(ctx context.Context, in *DeleteAllChannelsByAccountIDRequest, opts ...grpc.CallOption) (*DeleteAllChannelsByAccountIDResponse, error) {
	out := new(DeleteAllChannelsByAccountIDResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Channel/DeleteAllChannelsByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServer is the server API for Channel service.
// All implementations should embed UnimplementedChannelServer
// for forward compatibility
type ChannelServer interface {
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)
	GetChannelsByAccountID(context.Context, *GetChannelsByAccountIDRequest) (*GetChannelsByAccountIDResponse, error)
	DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error)
	DeleteAllChannelsByAccountID(context.Context, *DeleteAllChannelsByAccountIDRequest) (*DeleteAllChannelsByAccountIDResponse, error)
}

// UnimplementedChannelServer should be embedded to have forward compatible implementations.
type UnimplementedChannelServer struct {
}

func (UnimplementedChannelServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedChannelServer) GetChannelsByAccountID(context.Context, *GetChannelsByAccountIDRequest) (*GetChannelsByAccountIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannelsByAccountID not implemented")
}
func (UnimplementedChannelServer) DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChannel not implemented")
}
func (UnimplementedChannelServer) DeleteAllChannelsByAccountID(context.Context, *DeleteAllChannelsByAccountIDRequest) (*DeleteAllChannelsByAccountIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllChannelsByAccountID not implemented")
}

// UnsafeChannelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelServer will
// result in compilation errors.
type UnsafeChannelServer interface {
	mustEmbedUnimplementedChannelServer()
}

func RegisterChannelServer(s grpc.ServiceRegistrar, srv ChannelServer) {
	s.RegisterService(&Channel_ServiceDesc, srv)
}

func _Channel_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Channel/CreateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_GetChannelsByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelsByAccountIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).GetChannelsByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Channel/GetChannelsByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).GetChannelsByAccountID(ctx, req.(*GetChannelsByAccountIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Channel/DeleteChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Channel_DeleteAllChannelsByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllChannelsByAccountIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServer).DeleteAllChannelsByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Channel/DeleteAllChannelsByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServer).DeleteAllChannelsByAccountID(ctx, req.(*DeleteAllChannelsByAccountIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Channel_ServiceDesc is the grpc.ServiceDesc for Channel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Channel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Channel",
	HandlerType: (*ChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChannel",
			Handler:    _Channel_CreateChannel_Handler,
		},
		{
			MethodName: "GetChannelsByAccountID",
			Handler:    _Channel_GetChannelsByAccountID_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _Channel_DeleteChannel_Handler,
		},
		{
			MethodName: "DeleteAllChannelsByAccountID",
			Handler:    _Channel_DeleteAllChannelsByAccountID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/channel/v1alpha1/channel.proto",
}
