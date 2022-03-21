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

// ActorClient is the client API for Actor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActorClient interface {
	CreateActor(ctx context.Context, in *CreateActorRequest, opts ...grpc.CallOption) (*CreateActorResponse, error)
	// GetActorByAccountUsername Get by username returns the actor
	// identified by the specified username.
	GetActorByAccountUsername(ctx context.Context, in *GetActorByAccountUsernameRequest, opts ...grpc.CallOption) (*AccountDataResponse, error)
	// GetActorsByPreferredUsername Returns the set of actors by the
	// specified GetActorsByPreferredUsername by PreferredUsername.
	GetActorsByPreferredUsername(ctx context.Context, in *GetActorsByPreferredUsernameRequest, opts ...grpc.CallOption) (*GetActorsByPreferredUsernameResponse, error)
	// GetActorByAddress Returns the actor identified by the specified address.
	GetActorByAddress(ctx context.Context, in *GetActorByAddressRequest, opts ...grpc.CallOption) (*AccountDataResponse, error)
	// EditActor Edit the actor's profile and view the structure to get
	// the allowed and changeable parameters.
	// By username.
	EditActor(ctx context.Context, in *EditActorRequest, opts ...grpc.CallOption) (*EditActorResponse, error)
	DeleteActorByChannelID(ctx context.Context, in *DeleteActorByChannelIDRequest, opts ...grpc.CallOption) (*DeleteActorByChannelIDResponse, error)
}

type actorClient struct {
	cc grpc.ClientConnInterface
}

func NewActorClient(cc grpc.ClientConnInterface) ActorClient {
	return &actorClient{cc}
}

func (c *actorClient) CreateActor(ctx context.Context, in *CreateActorRequest, opts ...grpc.CallOption) (*CreateActorResponse, error) {
	out := new(CreateActorResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Actor/CreateActor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actorClient) GetActorByAccountUsername(ctx context.Context, in *GetActorByAccountUsernameRequest, opts ...grpc.CallOption) (*AccountDataResponse, error) {
	out := new(AccountDataResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Actor/GetActorByAccountUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actorClient) GetActorsByPreferredUsername(ctx context.Context, in *GetActorsByPreferredUsernameRequest, opts ...grpc.CallOption) (*GetActorsByPreferredUsernameResponse, error) {
	out := new(GetActorsByPreferredUsernameResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Actor/GetActorsByPreferredUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actorClient) GetActorByAddress(ctx context.Context, in *GetActorByAddressRequest, opts ...grpc.CallOption) (*AccountDataResponse, error) {
	out := new(AccountDataResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Actor/GetActorByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actorClient) EditActor(ctx context.Context, in *EditActorRequest, opts ...grpc.CallOption) (*EditActorResponse, error) {
	out := new(EditActorResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Actor/EditActor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actorClient) DeleteActorByChannelID(ctx context.Context, in *DeleteActorByChannelIDRequest, opts ...grpc.CallOption) (*DeleteActorByChannelIDResponse, error) {
	out := new(DeleteActorByChannelIDResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Actor/DeleteActorByChannelID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActorServer is the server API for Actor service.
// All implementations must embed UnimplementedActorServer
// for forward compatibility
type ActorServer interface {
	CreateActor(context.Context, *CreateActorRequest) (*CreateActorResponse, error)
	// GetActorByAccountUsername Get by username returns the actor
	// identified by the specified username.
	GetActorByAccountUsername(context.Context, *GetActorByAccountUsernameRequest) (*AccountDataResponse, error)
	// GetActorsByPreferredUsername Returns the set of actors by the
	// specified GetActorsByPreferredUsername by PreferredUsername.
	GetActorsByPreferredUsername(context.Context, *GetActorsByPreferredUsernameRequest) (*GetActorsByPreferredUsernameResponse, error)
	// GetActorByAddress Returns the actor identified by the specified address.
	GetActorByAddress(context.Context, *GetActorByAddressRequest) (*AccountDataResponse, error)
	// EditActor Edit the actor's profile and view the structure to get
	// the allowed and changeable parameters.
	// By username.
	EditActor(context.Context, *EditActorRequest) (*EditActorResponse, error)
	DeleteActorByChannelID(context.Context, *DeleteActorByChannelIDRequest) (*DeleteActorByChannelIDResponse, error)
	mustEmbedUnimplementedActorServer()
}

// UnimplementedActorServer must be embedded to have forward compatible implementations.
type UnimplementedActorServer struct {
}

func (UnimplementedActorServer) CreateActor(context.Context, *CreateActorRequest) (*CreateActorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActor not implemented")
}
func (UnimplementedActorServer) GetActorByAccountUsername(context.Context, *GetActorByAccountUsernameRequest) (*AccountDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActorByAccountUsername not implemented")
}
func (UnimplementedActorServer) GetActorsByPreferredUsername(context.Context, *GetActorsByPreferredUsernameRequest) (*GetActorsByPreferredUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActorsByPreferredUsername not implemented")
}
func (UnimplementedActorServer) GetActorByAddress(context.Context, *GetActorByAddressRequest) (*AccountDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActorByAddress not implemented")
}
func (UnimplementedActorServer) EditActor(context.Context, *EditActorRequest) (*EditActorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditActor not implemented")
}
func (UnimplementedActorServer) DeleteActorByChannelID(context.Context, *DeleteActorByChannelIDRequest) (*DeleteActorByChannelIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteActorByChannelID not implemented")
}
func (UnimplementedActorServer) mustEmbedUnimplementedActorServer() {}

// UnsafeActorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActorServer will
// result in compilation errors.
type UnsafeActorServer interface {
	mustEmbedUnimplementedActorServer()
}

func RegisterActorServer(s grpc.ServiceRegistrar, srv ActorServer) {
	s.RegisterService(&Actor_ServiceDesc, srv)
}

func _Actor_CreateActor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorServer).CreateActor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Actor/CreateActor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorServer).CreateActor(ctx, req.(*CreateActorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actor_GetActorByAccountUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActorByAccountUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorServer).GetActorByAccountUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Actor/GetActorByAccountUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorServer).GetActorByAccountUsername(ctx, req.(*GetActorByAccountUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actor_GetActorsByPreferredUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActorsByPreferredUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorServer).GetActorsByPreferredUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Actor/GetActorsByPreferredUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorServer).GetActorsByPreferredUsername(ctx, req.(*GetActorsByPreferredUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actor_GetActorByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActorByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorServer).GetActorByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Actor/GetActorByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorServer).GetActorByAddress(ctx, req.(*GetActorByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actor_EditActor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditActorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorServer).EditActor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Actor/EditActor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorServer).EditActor(ctx, req.(*EditActorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actor_DeleteActorByChannelID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActorByChannelIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorServer).DeleteActorByChannelID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Actor/DeleteActorByChannelID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorServer).DeleteActorByChannelID(ctx, req.(*DeleteActorByChannelIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Actor_ServiceDesc is the grpc.ServiceDesc for Actor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Actor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Actor",
	HandlerType: (*ActorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateActor",
			Handler:    _Actor_CreateActor_Handler,
		},
		{
			MethodName: "GetActorByAccountUsername",
			Handler:    _Actor_GetActorByAccountUsername_Handler,
		},
		{
			MethodName: "GetActorsByPreferredUsername",
			Handler:    _Actor_GetActorsByPreferredUsername_Handler,
		},
		{
			MethodName: "GetActorByAddress",
			Handler:    _Actor_GetActorByAddress_Handler,
		},
		{
			MethodName: "EditActor",
			Handler:    _Actor_EditActor_Handler,
		},
		{
			MethodName: "DeleteActorByChannelID",
			Handler:    _Actor_DeleteActorByChannelID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/account/v1alpha1/actor.proto",
}