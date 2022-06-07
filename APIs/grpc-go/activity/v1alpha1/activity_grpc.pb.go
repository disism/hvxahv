// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/activity/v1alpha1/activity.proto

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

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityClient interface {
	Inbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxResponse, error)
	GetInboxByActivityID(ctx context.Context, in *GetInboxByActivityIDRequest, opts ...grpc.CallOption) (*GetInboxByActivityIDResponse, error)
	GetInboxesByAccountID(ctx context.Context, in *GetInboxesByAccountIDRequest, opts ...grpc.CallOption) (*GetInboxesByAccountIDResponse, error)
	DeleteInboxByInboxesID(ctx context.Context, in *DeleteInboxByInboxesIDRequest, opts ...grpc.CallOption) (*DeleteInboxByInboxesIDResponse, error)
	CreateOutbox(ctx context.Context, in *CreateOutboxRequest, opts ...grpc.CallOption) (*CreateOutboxResponse, error)
	GetOutboxByActivityID(ctx context.Context, in *GetOutboxByActivityIDRequest, opts ...grpc.CallOption) (*GetOutboxByActivityIDResponse, error)
}

type activityClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityClient(cc grpc.ClientConnInterface) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) Inbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxResponse, error) {
	out := new(InboxResponse)
	err := c.cc.Invoke(ctx, "/hvx.activity.v1alpha1.proto.Activity/Inbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) GetInboxByActivityID(ctx context.Context, in *GetInboxByActivityIDRequest, opts ...grpc.CallOption) (*GetInboxByActivityIDResponse, error) {
	out := new(GetInboxByActivityIDResponse)
	err := c.cc.Invoke(ctx, "/hvx.activity.v1alpha1.proto.Activity/GetInboxByActivityID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) GetInboxesByAccountID(ctx context.Context, in *GetInboxesByAccountIDRequest, opts ...grpc.CallOption) (*GetInboxesByAccountIDResponse, error) {
	out := new(GetInboxesByAccountIDResponse)
	err := c.cc.Invoke(ctx, "/hvx.activity.v1alpha1.proto.Activity/GetInboxesByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) DeleteInboxByInboxesID(ctx context.Context, in *DeleteInboxByInboxesIDRequest, opts ...grpc.CallOption) (*DeleteInboxByInboxesIDResponse, error) {
	out := new(DeleteInboxByInboxesIDResponse)
	err := c.cc.Invoke(ctx, "/hvx.activity.v1alpha1.proto.Activity/DeleteInboxByInboxesID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) CreateOutbox(ctx context.Context, in *CreateOutboxRequest, opts ...grpc.CallOption) (*CreateOutboxResponse, error) {
	out := new(CreateOutboxResponse)
	err := c.cc.Invoke(ctx, "/hvx.activity.v1alpha1.proto.Activity/CreateOutbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) GetOutboxByActivityID(ctx context.Context, in *GetOutboxByActivityIDRequest, opts ...grpc.CallOption) (*GetOutboxByActivityIDResponse, error) {
	out := new(GetOutboxByActivityIDResponse)
	err := c.cc.Invoke(ctx, "/hvx.activity.v1alpha1.proto.Activity/GetOutboxByActivityID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
// All implementations should embed UnimplementedActivityServer
// for forward compatibility
type ActivityServer interface {
	Inbox(context.Context, *InboxRequest) (*InboxResponse, error)
	GetInboxByActivityID(context.Context, *GetInboxByActivityIDRequest) (*GetInboxByActivityIDResponse, error)
	GetInboxesByAccountID(context.Context, *GetInboxesByAccountIDRequest) (*GetInboxesByAccountIDResponse, error)
	DeleteInboxByInboxesID(context.Context, *DeleteInboxByInboxesIDRequest) (*DeleteInboxByInboxesIDResponse, error)
	CreateOutbox(context.Context, *CreateOutboxRequest) (*CreateOutboxResponse, error)
	GetOutboxByActivityID(context.Context, *GetOutboxByActivityIDRequest) (*GetOutboxByActivityIDResponse, error)
}

// UnimplementedActivityServer should be embedded to have forward compatible implementations.
type UnimplementedActivityServer struct {
}

func (UnimplementedActivityServer) Inbox(context.Context, *InboxRequest) (*InboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inbox not implemented")
}
func (UnimplementedActivityServer) GetInboxByActivityID(context.Context, *GetInboxByActivityIDRequest) (*GetInboxByActivityIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInboxByActivityID not implemented")
}
func (UnimplementedActivityServer) GetInboxesByAccountID(context.Context, *GetInboxesByAccountIDRequest) (*GetInboxesByAccountIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInboxesByAccountID not implemented")
}
func (UnimplementedActivityServer) DeleteInboxByInboxesID(context.Context, *DeleteInboxByInboxesIDRequest) (*DeleteInboxByInboxesIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInboxByInboxesID not implemented")
}
func (UnimplementedActivityServer) CreateOutbox(context.Context, *CreateOutboxRequest) (*CreateOutboxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOutbox not implemented")
}
func (UnimplementedActivityServer) GetOutboxByActivityID(context.Context, *GetOutboxByActivityIDRequest) (*GetOutboxByActivityIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOutboxByActivityID not implemented")
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

func _Activity_Inbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).Inbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.activity.v1alpha1.proto.Activity/Inbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).Inbox(ctx, req.(*InboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_GetInboxByActivityID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInboxByActivityIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetInboxByActivityID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.activity.v1alpha1.proto.Activity/GetInboxByActivityID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetInboxByActivityID(ctx, req.(*GetInboxByActivityIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_GetInboxesByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInboxesByAccountIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetInboxesByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.activity.v1alpha1.proto.Activity/GetInboxesByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetInboxesByAccountID(ctx, req.(*GetInboxesByAccountIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_DeleteInboxByInboxesID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInboxByInboxesIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).DeleteInboxByInboxesID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.activity.v1alpha1.proto.Activity/DeleteInboxByInboxesID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).DeleteInboxByInboxesID(ctx, req.(*DeleteInboxByInboxesIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_CreateOutbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOutboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).CreateOutbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.activity.v1alpha1.proto.Activity/CreateOutbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).CreateOutbox(ctx, req.(*CreateOutboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_GetOutboxByActivityID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOutboxByActivityIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetOutboxByActivityID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.activity.v1alpha1.proto.Activity/GetOutboxByActivityID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetOutboxByActivityID(ctx, req.(*GetOutboxByActivityIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_ServiceDesc is the grpc.ServiceDesc for Activity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvx.activity.v1alpha1.proto.Activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Inbox",
			Handler:    _Activity_Inbox_Handler,
		},
		{
			MethodName: "GetInboxByActivityID",
			Handler:    _Activity_GetInboxByActivityID_Handler,
		},
		{
			MethodName: "GetInboxesByAccountID",
			Handler:    _Activity_GetInboxesByAccountID_Handler,
		},
		{
			MethodName: "DeleteInboxByInboxesID",
			Handler:    _Activity_DeleteInboxByInboxesID_Handler,
		},
		{
			MethodName: "CreateOutbox",
			Handler:    _Activity_CreateOutbox_Handler,
		},
		{
			MethodName: "GetOutboxByActivityID",
			Handler:    _Activity_GetOutboxByActivityID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/activity/v1alpha1/activity.proto",
}
