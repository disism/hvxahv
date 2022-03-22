// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/account/v1alpha1/account.proto

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
	// IsExist Check the account creation status by username,
	// return true if not created, otherwise return false, used
	// when checking if the user exists.
	IsExist(ctx context.Context, in *IsExistRequest, opts ...grpc.CallOption) (*IsExistResponse, error)
	// Create a new account.
	// Actor is created first and then the returned ActorID is used to
	//  create a unique account for the current instance account system.
	// The Actor's PreferredUsername is used to identify the actor in the
	// current instance account system. The username in the account system
	// is unique, but the Actor's PreferredUsername is not unique,
	// as it may have the same username in different instances.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get the account by username.
	// Internal methods in the program should not be exposed to API
	// structures for users to call. There is no other solution for
	// the time being, so we will use the code for now and modify it later.
	GetAccountByUsername(ctx context.Context, in *GetAccountByUsernameRequest, opts ...grpc.CallOption) (*GetAccountByUsernameResponse, error)
	// First you need the password as a parameter to verify that the account
	// is correct, and then delete the account system and actor table when
	// the password is verified to be correct.
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
	// EditUsername Update the account username.
	// will change both the username and preferred_username in the
	// accounts and actors tables.
	// As with password changes, you need to take down all clients
	// and log back in to issue the token.
	EditUsername(ctx context.Context, in *EditUsernameRequest, opts ...grpc.CallOption) (*EditUsernameResponse, error)
	// EditPassword Update the account password.
	// When changing the password for an account, all client authorizations
	// need to be removed and all devices that have logged in must log in
	// again to perform all account operations; this verification is
	// verified in the REST API TOKEN.
	EditPassword(ctx context.Context, in *EditPasswordRequest, opts ...grpc.CallOption) (*EditPasswordResponse, error)
	// EditEmail Edit the unique email for the account.
	// All devices should be taken offline.
	EditEmail(ctx context.Context, in *EditEmailRequest, opts ...grpc.CallOption) (*EditEmailResponse, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) IsExist(ctx context.Context, in *IsExistRequest, opts ...grpc.CallOption) (*IsExistResponse, error) {
	out := new(IsExistResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/IsExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetAccountByUsername(ctx context.Context, in *GetAccountByUsernameRequest, opts ...grpc.CallOption) (*GetAccountByUsernameResponse, error) {
	out := new(GetAccountByUsernameResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/GetAccountByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditUsername(ctx context.Context, in *EditUsernameRequest, opts ...grpc.CallOption) (*EditUsernameResponse, error) {
	out := new(EditUsernameResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/EditUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditPassword(ctx context.Context, in *EditPasswordRequest, opts ...grpc.CallOption) (*EditPasswordResponse, error) {
	out := new(EditPasswordResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/EditPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditEmail(ctx context.Context, in *EditEmailRequest, opts ...grpc.CallOption) (*EditEmailResponse, error) {
	out := new(EditEmailResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Accounts/EditEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
// All implementations must embed UnimplementedAccountsServer
// for forward compatibility
type AccountsServer interface {
	// IsExist Check the account creation status by username,
	// return true if not created, otherwise return false, used
	// when checking if the user exists.
	IsExist(context.Context, *IsExistRequest) (*IsExistResponse, error)
	// Create a new account.
	// Actor is created first and then the returned ActorID is used to
	//  create a unique account for the current instance account system.
	// The Actor's PreferredUsername is used to identify the actor in the
	// current instance account system. The username in the account system
	// is unique, but the Actor's PreferredUsername is not unique,
	// as it may have the same username in different instances.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get the account by username.
	// Internal methods in the program should not be exposed to API
	// structures for users to call. There is no other solution for
	// the time being, so we will use the code for now and modify it later.
	GetAccountByUsername(context.Context, *GetAccountByUsernameRequest) (*GetAccountByUsernameResponse, error)
	// First you need the password as a parameter to verify that the account
	// is correct, and then delete the account system and actor table when
	// the password is verified to be correct.
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
	// EditUsername Update the account username.
	// will change both the username and preferred_username in the
	// accounts and actors tables.
	// As with password changes, you need to take down all clients
	// and log back in to issue the token.
	EditUsername(context.Context, *EditUsernameRequest) (*EditUsernameResponse, error)
	// EditPassword Update the account password.
	// When changing the password for an account, all client authorizations
	// need to be removed and all devices that have logged in must log in
	// again to perform all account operations; this verification is
	// verified in the REST API TOKEN.
	EditPassword(context.Context, *EditPasswordRequest) (*EditPasswordResponse, error)
	// EditEmail Edit the unique email for the account.
	// All devices should be taken offline.
	EditEmail(context.Context, *EditEmailRequest) (*EditEmailResponse, error)
	mustEmbedUnimplementedAccountsServer()
}

// UnimplementedAccountsServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (UnimplementedAccountsServer) IsExist(context.Context, *IsExistRequest) (*IsExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsExist not implemented")
}
func (UnimplementedAccountsServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccountsServer) GetAccountByUsername(context.Context, *GetAccountByUsernameRequest) (*GetAccountByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByUsername not implemented")
}
func (UnimplementedAccountsServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountsServer) EditUsername(context.Context, *EditUsernameRequest) (*EditUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUsername not implemented")
}
func (UnimplementedAccountsServer) EditPassword(context.Context, *EditPasswordRequest) (*EditPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPassword not implemented")
}
func (UnimplementedAccountsServer) EditEmail(context.Context, *EditEmailRequest) (*EditEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditEmail not implemented")
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
	in := new(IsExistRequest)
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
		return srv.(AccountsServer).IsExist(ctx, req.(*IsExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
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
		return srv.(AccountsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetAccountByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByUsernameRequest)
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
		return srv.(AccountsServer).GetAccountByUsername(ctx, req.(*GetAccountByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
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
		return srv.(AccountsServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_EditUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUsernameRequest)
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
		return srv.(AccountsServer).EditUsername(ctx, req.(*EditUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_EditPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditPasswordRequest)
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
		return srv.(AccountsServer).EditPassword(ctx, req.(*EditPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_EditEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).EditEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Accounts/EditEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).EditEmail(ctx, req.(*EditEmailRequest))
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
			MethodName: "GetAccountByUsername",
			Handler:    _Accounts_GetAccountByUsername_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Accounts_DeleteAccount_Handler,
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
			MethodName: "EditEmail",
			Handler:    _Accounts_EditEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/account/v1alpha1/account.proto",
}
