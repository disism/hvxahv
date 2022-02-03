// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

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

// NotifyClient is the client API for Notify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifyClient interface {
	Subscription(ctx context.Context, in *NewNotifySubscription, opts ...grpc.CallOption) (*NotifySubscriptionReply, error)
	Push(ctx context.Context, in *NewNotifyPush, opts ...grpc.CallOption) (*NotifyPushReply, error)
}

type notifyClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifyClient(cc grpc.ClientConnInterface) NotifyClient {
	return &notifyClient{cc}
}

func (c *notifyClient) Subscription(ctx context.Context, in *NewNotifySubscription, opts ...grpc.CallOption) (*NotifySubscriptionReply, error) {
	out := new(NotifySubscriptionReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Notify/Subscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyClient) Push(ctx context.Context, in *NewNotifyPush, opts ...grpc.CallOption) (*NotifyPushReply, error) {
	out := new(NotifyPushReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Notify/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyServer is the server API for Notify service.
// All implementations must embed UnimplementedNotifyServer
// for forward compatibility
type NotifyServer interface {
	Subscription(context.Context, *NewNotifySubscription) (*NotifySubscriptionReply, error)
	Push(context.Context, *NewNotifyPush) (*NotifyPushReply, error)
	mustEmbedUnimplementedNotifyServer()
}

// UnimplementedNotifyServer must be embedded to have forward compatible implementations.
type UnimplementedNotifyServer struct {
}

func (UnimplementedNotifyServer) Subscription(context.Context, *NewNotifySubscription) (*NotifySubscriptionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscription not implemented")
}
func (UnimplementedNotifyServer) Push(context.Context, *NewNotifyPush) (*NotifyPushReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedNotifyServer) mustEmbedUnimplementedNotifyServer() {}

// UnsafeNotifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifyServer will
// result in compilation errors.
type UnsafeNotifyServer interface {
	mustEmbedUnimplementedNotifyServer()
}

func RegisterNotifyServer(s grpc.ServiceRegistrar, srv NotifyServer) {
	s.RegisterService(&Notify_ServiceDesc, srv)
}

func _Notify_Subscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewNotifySubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).Subscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Notify/Subscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).Subscription(ctx, req.(*NewNotifySubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notify_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewNotifyPush)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Notify/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyServer).Push(ctx, req.(*NewNotifyPush))
	}
	return interceptor(ctx, in, info, handler)
}

// Notify_ServiceDesc is the grpc.ServiceDesc for Notify service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notify_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Notify",
	HandlerType: (*NotifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscription",
			Handler:    _Notify_Subscription_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _Notify_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/notify/v1alpha1/notify.proto",
}
