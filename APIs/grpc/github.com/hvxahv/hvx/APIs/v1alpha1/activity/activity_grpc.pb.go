// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/v1alpha1/activity/activity.proto

package activity

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

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityClient interface {
	// New Design...
	// Activity create an activity that is sent to the activity pub instance server of the specified actor (server to server interactions),
	// When sending messages to the Mastodon server, we found that Mastodon must verify the user's identity by signing,
	// and if the actor's private key is stored locally,
	// we need to sign locally and then submit the signature to the server before sending,
	// but we found serious performance problems when doing rsa private key signing on the client side,
	// which is why we put the key supporting activitypub feature on the server.
	// So in the tradeoff of privacy, we decided to design two key systems,
	// one for asymmetric encryption of accounts and one for activitypub key pairs.
	Activity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	ArticleActivity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
}

type activityClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityClient(cc grpc.ClientConnInterface) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) Activity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.activity.proto.Activity/Activity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) ArticleActivity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.activity.proto.Activity/ArticleActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
// All implementations should embed UnimplementedActivityServer
// for forward compatibility
type ActivityServer interface {
	// New Design...
	// Activity create an activity that is sent to the activity pub instance server of the specified actor (server to server interactions),
	// When sending messages to the Mastodon server, we found that Mastodon must verify the user's identity by signing,
	// and if the actor's private key is stored locally,
	// we need to sign locally and then submit the signature to the server before sending,
	// but we found serious performance problems when doing rsa private key signing on the client side,
	// which is why we put the key supporting activitypub feature on the server.
	// So in the tradeoff of privacy, we decided to design two key systems,
	// one for asymmetric encryption of accounts and one for activitypub key pairs.
	Activity(context.Context, *ActivityRequest) (*ActivityResponse, error)
	ArticleActivity(context.Context, *ActivityRequest) (*ActivityResponse, error)
}

// UnimplementedActivityServer should be embedded to have forward compatible implementations.
type UnimplementedActivityServer struct {
}

func (UnimplementedActivityServer) Activity(context.Context, *ActivityRequest) (*ActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activity not implemented")
}
func (UnimplementedActivityServer) ArticleActivity(context.Context, *ActivityRequest) (*ActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleActivity not implemented")
}

// UnsafeActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServer will
// result in compilation errors.
type UnsafeActivityServer interface {
	mustEmbedUnimplementedActivityServer()
}

func RegisterActivityServer(s grpc.ServiceRegistrar, srv ActivityServer) {
	s.RegisterService(&Activity_ServiceDesc, srv)
}

func _Activity_Activity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).Activity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.activity.proto.Activity/Activity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).Activity(ctx, req.(*ActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_ArticleActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).ArticleActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.activity.proto.Activity/ArticleActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).ArticleActivity(ctx, req.(*ActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_ServiceDesc is the grpc.ServiceDesc for Activity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvx.api.v1alpha1.activity.proto.Activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activity",
			Handler:    _Activity_Activity_Handler,
		},
		{
			MethodName: "ArticleActivity",
			Handler:    _Activity_ArticleActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/activity/activity.proto",
}
