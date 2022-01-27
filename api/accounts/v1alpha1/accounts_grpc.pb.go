// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/accounts/v1alpha1/accounts.proto

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

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsClient interface {
	IsExist(ctx context.Context, in *NewAccountUsername, opts ...grpc.CallOption) (*IsExistReply, error)
	// Create the Actor first, and then use the returned ActorID to create a unique account of the current instance account system.
	// The username in the account system is unique, and the Actor may have the same username in different instances.
	Create(ctx context.Context, in *NewAccountCreate, opts ...grpc.CallOption) (*Reply, error)
	//  Log in to the account system interface.
	Verify(ctx context.Context, in *NewAccountVerify, opts ...grpc.CallOption) (*VerifyAccountReply, error)
	GetAccountByUsername(ctx context.Context, in *NewAccountUsername, opts ...grpc.CallOption) (*AccountData, error)
	Delete(ctx context.Context, in *NewAccountDelete, opts ...grpc.CallOption) (*Reply, error)
	// EditUsername The account username will be updated with the account ID, along with the Actor preferred username.
	EditUsername(ctx context.Context, in *NewEditAccountUsername, opts ...grpc.CallOption) (*Reply, error)
	EditPassword(ctx context.Context, in *NewEditAccountPassword, opts ...grpc.CallOption) (*Reply, error)
	EditMail(ctx context.Context, in *NewEditAccountMail, opts ...grpc.CallOption) (*Reply, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) IsExist(ctx context.Context, in *NewAccountUsername, opts ...grpc.CallOption) (*IsExistReply, error) {
	out := new(IsExistReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/IsExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Create(ctx context.Context, in *NewAccountCreate, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Verify(ctx context.Context, in *NewAccountVerify, opts ...grpc.CallOption) (*VerifyAccountReply, error) {
	out := new(VerifyAccountReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetAccountByUsername(ctx context.Context, in *NewAccountUsername, opts ...grpc.CallOption) (*AccountData, error) {
	out := new(AccountData)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/GetAccountByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Delete(ctx context.Context, in *NewAccountDelete, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditUsername(ctx context.Context, in *NewEditAccountUsername, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/EditUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditPassword(ctx context.Context, in *NewEditAccountPassword, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/EditPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditMail(ctx context.Context, in *NewEditAccountMail, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/EditMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
// All implementations must embed UnimplementedAccountsServer
// for forward compatibility
type AccountsServer interface {
	IsExist(context.Context, *NewAccountUsername) (*IsExistReply, error)
	// Create the Actor first, and then use the returned ActorID to create a unique account of the current instance account system.
	// The username in the account system is unique, and the Actor may have the same username in different instances.
	Create(context.Context, *NewAccountCreate) (*Reply, error)
	//  Log in to the account system interface.
	Verify(context.Context, *NewAccountVerify) (*VerifyAccountReply, error)
	GetAccountByUsername(context.Context, *NewAccountUsername) (*AccountData, error)
	Delete(context.Context, *NewAccountDelete) (*Reply, error)
	// EditUsername The account username will be updated with the account ID, along with the Actor preferred username.
	EditUsername(context.Context, *NewEditAccountUsername) (*Reply, error)
	EditPassword(context.Context, *NewEditAccountPassword) (*Reply, error)
	EditMail(context.Context, *NewEditAccountMail) (*Reply, error)
	mustEmbedUnimplementedAccountsServer()
}

// UnimplementedAccountsServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (UnimplementedAccountsServer) IsExist(context.Context, *NewAccountUsername) (*IsExistReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsExist not implemented")
}
func (UnimplementedAccountsServer) Create(context.Context, *NewAccountCreate) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccountsServer) Verify(context.Context, *NewAccountVerify) (*VerifyAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAccountsServer) GetAccountByUsername(context.Context, *NewAccountUsername) (*AccountData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByUsername not implemented")
}
func (UnimplementedAccountsServer) Delete(context.Context, *NewAccountDelete) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAccountsServer) EditUsername(context.Context, *NewEditAccountUsername) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUsername not implemented")
}
func (UnimplementedAccountsServer) EditPassword(context.Context, *NewEditAccountPassword) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPassword not implemented")
}
func (UnimplementedAccountsServer) EditMail(context.Context, *NewEditAccountMail) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMail not implemented")
}
func (UnimplementedAccountsServer) mustEmbedUnimplementedAccountsServer() {}

// UnsafeAccountsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsServer will
// result in compilation errors.
type UnsafeAccountsServer interface {
	mustEmbedUnimplementedAccountsServer()
}

func RegisterAccountsServer(s grpc.ServiceRegistrar, srv AccountsServer) {
	s.RegisterService(&Accounts_ServiceDesc, srv)
}

func _Accounts_IsExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).IsExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/IsExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).IsExist(ctx, req.(*NewAccountUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Create(ctx, req.(*NewAccountCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountVerify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Verify(ctx, req.(*NewAccountVerify))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetAccountByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetAccountByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/GetAccountByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetAccountByUsername(ctx, req.(*NewAccountUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Delete(ctx, req.(*NewAccountDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_EditUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEditAccountUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).EditUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/EditUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).EditUsername(ctx, req.(*NewEditAccountUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_EditPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEditAccountPassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).EditPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/EditPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).EditPassword(ctx, req.(*NewEditAccountPassword))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_EditMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEditAccountMail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).EditMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/EditMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).EditMail(ctx, req.(*NewEditAccountMail))
	}
	return interceptor(ctx, in, info, handler)
}

// Accounts_ServiceDesc is the grpc.ServiceDesc for Accounts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Accounts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsExist",
			Handler:    _Accounts_IsExist_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Accounts_Create_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Accounts_Verify_Handler,
		},
		{
			MethodName: "GetAccountByUsername",
			Handler:    _Accounts_GetAccountByUsername_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Accounts_Delete_Handler,
		},
		{
			MethodName: "EditUsername",
			Handler:    _Accounts_EditUsername_Handler,
		},
		{
			MethodName: "EditPassword",
			Handler:    _Accounts_EditPassword_Handler,
		},
		{
			MethodName: "EditMail",
			Handler:    _Accounts_EditMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/accounts/v1alpha1/accounts.proto",
}

// ActorsClient is the client API for Actors service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActorsClient interface {
	GetActorByAccountUsername(ctx context.Context, in *NewAccountUsername, opts ...grpc.CallOption) (*ActorData, error)
	GetActorsByPreferredUsername(ctx context.Context, in *NewActorPreferredUsername, opts ...grpc.CallOption) (*ActorsData, error)
}

type actorsClient struct {
	cc grpc.ClientConnInterface
}

func NewActorsClient(cc grpc.ClientConnInterface) ActorsClient {
	return &actorsClient{cc}
}

func (c *actorsClient) GetActorByAccountUsername(ctx context.Context, in *NewAccountUsername, opts ...grpc.CallOption) (*ActorData, error) {
	out := new(ActorData)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Actors/GetActorByAccountUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actorsClient) GetActorsByPreferredUsername(ctx context.Context, in *NewActorPreferredUsername, opts ...grpc.CallOption) (*ActorsData, error) {
	out := new(ActorsData)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Actors/GetActorsByPreferredUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActorsServer is the server API for Actors service.
// All implementations must embed UnimplementedActorsServer
// for forward compatibility
type ActorsServer interface {
	GetActorByAccountUsername(context.Context, *NewAccountUsername) (*ActorData, error)
	GetActorsByPreferredUsername(context.Context, *NewActorPreferredUsername) (*ActorsData, error)
	mustEmbedUnimplementedActorsServer()
}

// UnimplementedActorsServer must be embedded to have forward compatible implementations.
type UnimplementedActorsServer struct {
}

func (UnimplementedActorsServer) GetActorByAccountUsername(context.Context, *NewAccountUsername) (*ActorData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActorByAccountUsername not implemented")
}
func (UnimplementedActorsServer) GetActorsByPreferredUsername(context.Context, *NewActorPreferredUsername) (*ActorsData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActorsByPreferredUsername not implemented")
}
func (UnimplementedActorsServer) mustEmbedUnimplementedActorsServer() {}

// UnsafeActorsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActorsServer will
// result in compilation errors.
type UnsafeActorsServer interface {
	mustEmbedUnimplementedActorsServer()
}

func RegisterActorsServer(s grpc.ServiceRegistrar, srv ActorsServer) {
	s.RegisterService(&Actors_ServiceDesc, srv)
}

func _Actors_GetActorByAccountUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorsServer).GetActorByAccountUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Actors/GetActorByAccountUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorsServer).GetActorByAccountUsername(ctx, req.(*NewAccountUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actors_GetActorsByPreferredUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewActorPreferredUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActorsServer).GetActorsByPreferredUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Actors/GetActorsByPreferredUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActorsServer).GetActorsByPreferredUsername(ctx, req.(*NewActorPreferredUsername))
	}
	return interceptor(ctx, in, info, handler)
}

// Actors_ServiceDesc is the grpc.ServiceDesc for Actors service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Actors_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Actors",
	HandlerType: (*ActorsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActorByAccountUsername",
			Handler:    _Actors_GetActorByAccountUsername_Handler,
		},
		{
			MethodName: "GetActorsByPreferredUsername",
			Handler:    _Actors_GetActorsByPreferredUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/accounts/v1alpha1/accounts.proto",
}
