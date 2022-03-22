// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/channel/v1alpha1/subscriber.proto

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

// SubscriberServiceClient is the client API for SubscriberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriberServiceClient interface {
	AddSubscriber(ctx context.Context, in *AddSubscriberRequest, opts ...grpc.CallOption) (*AddSubscriberResponse, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
	RemoveSubscriber(ctx context.Context, in *RemoveSubscriberRequest, opts ...grpc.CallOption) (*RemoveSubscriberResponse, error)
	GetSubscribers(ctx context.Context, in *GetSubscribersRequest, opts ...grpc.CallOption) (*GetSubscribersResponse, error)
}

type subscriberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriberServiceClient(cc grpc.ClientConnInterface) SubscriberServiceClient {
	return &subscriberServiceClient{cc}
}

func (c *subscriberServiceClient) AddSubscriber(ctx context.Context, in *AddSubscriberRequest, opts ...grpc.CallOption) (*AddSubscriberResponse, error) {
	out := new(AddSubscriberResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.SubscriberService/AddSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriberServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.SubscriberService/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriberServiceClient) RemoveSubscriber(ctx context.Context, in *RemoveSubscriberRequest, opts ...grpc.CallOption) (*RemoveSubscriberResponse, error) {
	out := new(RemoveSubscriberResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.SubscriberService/RemoveSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriberServiceClient) GetSubscribers(ctx context.Context, in *GetSubscribersRequest, opts ...grpc.CallOption) (*GetSubscribersResponse, error) {
	out := new(GetSubscribersResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.SubscriberService/GetSubscribers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriberServiceServer is the server API for SubscriberService service.
// All implementations must embed UnimplementedSubscriberServiceServer
// for forward compatibility
type SubscriberServiceServer interface {
	AddSubscriber(context.Context, *AddSubscriberRequest) (*AddSubscriberResponse, error)
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
	RemoveSubscriber(context.Context, *RemoveSubscriberRequest) (*RemoveSubscriberResponse, error)
	GetSubscribers(context.Context, *GetSubscribersRequest) (*GetSubscribersResponse, error)
	mustEmbedUnimplementedSubscriberServiceServer()
}

// UnimplementedSubscriberServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriberServiceServer struct {
}

func (UnimplementedSubscriberServiceServer) AddSubscriber(context.Context, *AddSubscriberRequest) (*AddSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscriber not implemented")
}
func (UnimplementedSubscriberServiceServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedSubscriberServiceServer) RemoveSubscriber(context.Context, *RemoveSubscriberRequest) (*RemoveSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSubscriber not implemented")
}
func (UnimplementedSubscriberServiceServer) GetSubscribers(context.Context, *GetSubscribersRequest) (*GetSubscribersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscribers not implemented")
}
func (UnimplementedSubscriberServiceServer) mustEmbedUnimplementedSubscriberServiceServer() {}

// UnsafeSubscriberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriberServiceServer will
// result in compilation errors.
type UnsafeSubscriberServiceServer interface {
	mustEmbedUnimplementedSubscriberServiceServer()
}

func RegisterSubscriberServiceServer(s grpc.ServiceRegistrar, srv SubscriberServiceServer) {
	s.RegisterService(&SubscriberService_ServiceDesc, srv)
}

func _SubscriberService_AddSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberServiceServer).AddSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.SubscriberService/AddSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberServiceServer).AddSubscriber(ctx, req.(*AddSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriberService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.SubscriberService/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberServiceServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriberService_RemoveSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberServiceServer).RemoveSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.SubscriberService/RemoveSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberServiceServer).RemoveSubscriber(ctx, req.(*RemoveSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriberService_GetSubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscribersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberServiceServer).GetSubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.SubscriberService/GetSubscribers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberServiceServer).GetSubscribers(ctx, req.(*GetSubscribersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriberService_ServiceDesc is the grpc.ServiceDesc for SubscriberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.SubscriberService",
	HandlerType: (*SubscriberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSubscriber",
			Handler:    _SubscriberService_AddSubscriber_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _SubscriberService_Unsubscribe_Handler,
		},
		{
			MethodName: "RemoveSubscriber",
			Handler:    _SubscriberService_RemoveSubscriber_Handler,
		},
		{
			MethodName: "GetSubscribers",
			Handler:    _SubscriberService_GetSubscribers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/channel/v1alpha1/subscriber.proto",
}
