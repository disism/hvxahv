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

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsClient interface {
	// Create the Actor first, and then use the returned ActorID to create a unique account of the current instance account system.
	// The username in the account system is unique, and the Actor may have the same username in different instances.
	Create(ctx context.Context, in *NewCreate, opts ...grpc.CallOption) (*Reply, error)
	//  Log in to the account system interface.
	Verify(ctx context.Context, in *NewVerify, opts ...grpc.CallOption) (*VerifyReply, error)
	GetActorByAccountUsername(ctx context.Context, in *NewAccountUsername, opts ...grpc.CallOption) (*ActorData, error)
	GetAccountByUsername(ctx context.Context, in *NewAccountUsername, opts ...grpc.CallOption) (*AccountData, error)
	DeleteAccount(ctx context.Context, in *NewAccountDelete, opts ...grpc.CallOption) (*Reply, error)
	// EditAccountUsername The account username will be updated with the account ID, along with the Actor preferred username.
	EditAccountUsername(ctx context.Context, in *NewEditAccountUsername, opts ...grpc.CallOption) (*Reply, error)
	EditAccountPassword(ctx context.Context, in *NewEditAccountPassword, opts ...grpc.CallOption) (*Reply, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) Create(ctx context.Context, in *NewCreate, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Verify(ctx context.Context, in *NewVerify, opts ...grpc.CallOption) (*VerifyReply, error) {
	out := new(VerifyReply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetActorByAccountUsername(ctx context.Context, in *NewAccountUsername, opts ...grpc.CallOption) (*ActorData, error) {
	out := new(ActorData)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/GetActorByAccountUsername", in, out, opts...)
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

func (c *accountsClient) DeleteAccount(ctx context.Context, in *NewAccountDelete, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditAccountUsername(ctx context.Context, in *NewEditAccountUsername, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/EditAccountUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditAccountPassword(ctx context.Context, in *NewEditAccountPassword, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/EditAccountPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
// All implementations must embed UnimplementedAccountsServer
// for forward compatibility
type AccountsServer interface {
	// Create the Actor first, and then use the returned ActorID to create a unique account of the current instance account system.
	// The username in the account system is unique, and the Actor may have the same username in different instances.
	Create(context.Context, *NewCreate) (*Reply, error)
	//  Log in to the account system interface.
	Verify(context.Context, *NewVerify) (*VerifyReply, error)
	GetActorByAccountUsername(context.Context, *NewAccountUsername) (*ActorData, error)
	GetAccountByUsername(context.Context, *NewAccountUsername) (*AccountData, error)
	DeleteAccount(context.Context, *NewAccountDelete) (*Reply, error)
	// EditAccountUsername The account username will be updated with the account ID, along with the Actor preferred username.
	EditAccountUsername(context.Context, *NewEditAccountUsername) (*Reply, error)
	EditAccountPassword(context.Context, *NewEditAccountPassword) (*Reply, error)
	mustEmbedUnimplementedAccountsServer()
}

// UnimplementedAccountsServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (UnimplementedAccountsServer) Create(context.Context, *NewCreate) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccountsServer) Verify(context.Context, *NewVerify) (*VerifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAccountsServer) GetActorByAccountUsername(context.Context, *NewAccountUsername) (*ActorData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActorByAccountUsername not implemented")
}
func (UnimplementedAccountsServer) GetAccountByUsername(context.Context, *NewAccountUsername) (*AccountData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByUsername not implemented")
}
func (UnimplementedAccountsServer) DeleteAccount(context.Context, *NewAccountDelete) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountsServer) EditAccountUsername(context.Context, *NewEditAccountUsername) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAccountUsername not implemented")
}
func (UnimplementedAccountsServer) EditAccountPassword(context.Context, *NewEditAccountPassword) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAccountPassword not implemented")
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

func _Accounts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCreate)
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
		return srv.(AccountsServer).Create(ctx, req.(*NewCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewVerify)
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
		return srv.(AccountsServer).Verify(ctx, req.(*NewVerify))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetActorByAccountUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetActorByAccountUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/GetActorByAccountUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetActorByAccountUsername(ctx, req.(*NewAccountUsername))
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

func _Accounts_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).DeleteAccount(ctx, req.(*NewAccountDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_EditAccountUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEditAccountUsername)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).EditAccountUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/EditAccountUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).EditAccountUsername(ctx, req.(*NewEditAccountUsername))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_EditAccountPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEditAccountPassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).EditAccountPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/EditAccountPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).EditAccountPassword(ctx, req.(*NewEditAccountPassword))
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
			MethodName: "Create",
			Handler:    _Accounts_Create_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Accounts_Verify_Handler,
		},
		{
			MethodName: "GetActorByAccountUsername",
			Handler:    _Accounts_GetActorByAccountUsername_Handler,
		},
		{
			MethodName: "GetAccountByUsername",
			Handler:    _Accounts_GetAccountByUsername_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Accounts_DeleteAccount_Handler,
		},
		{
			MethodName: "EditAccountUsername",
			Handler:    _Accounts_EditAccountUsername_Handler,
		},
		{
			MethodName: "EditAccountPassword",
			Handler:    _Accounts_EditAccountPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/accounts/v1alpha1/accounts.proto",
}
