// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/v1alpha1/account/account.proto

package account

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
	// IsExist checks if the account exists. The account name(username) is specified in the request.
	// If the account exists, it returns true. Otherwise, it returns false.
	IsExist(ctx context.Context, in *IsExistRequest, opts ...grpc.CallOption) (*IsExistResponse, error)
	// Create a new account.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get the account by username.
	GetByUsername(ctx context.Context, in *GetByUsernameRequest, opts ...grpc.CallOption) (*GetByUsernameResponse, error)
	// Delete the account.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// EditUsername Update the account username.
	// will change the account username and the preferred_username in Actor.
	EditUsername(ctx context.Context, in *EditUsernameRequest, opts ...grpc.CallOption) (*EditUsernameResponse, error)
	// EditPassword Update the account password.
	EditPassword(ctx context.Context, in *EditPasswordRequest, opts ...grpc.CallOption) (*EditPasswordResponse, error)
	// EditEmail Edit the unique email for the account.
	EditEmail(ctx context.Context, in *EditEmailRequest, opts ...grpc.CallOption) (*EditEmailResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequest, opts ...grpc.CallOption) (*GetPrivateKeyResponse, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) IsExist(ctx context.Context, in *IsExistRequest, opts ...grpc.CallOption) (*IsExistResponse, error) {
	out := new(IsExistResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.account.proto.Accounts/IsExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.account.proto.Accounts/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetByUsername(ctx context.Context, in *GetByUsernameRequest, opts ...grpc.CallOption) (*GetByUsernameResponse, error) {
	out := new(GetByUsernameResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.account.proto.Accounts/GetByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.account.proto.Accounts/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditUsername(ctx context.Context, in *EditUsernameRequest, opts ...grpc.CallOption) (*EditUsernameResponse, error) {
	out := new(EditUsernameResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.account.proto.Accounts/EditUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditPassword(ctx context.Context, in *EditPasswordRequest, opts ...grpc.CallOption) (*EditPasswordResponse, error) {
	out := new(EditPasswordResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.account.proto.Accounts/EditPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) EditEmail(ctx context.Context, in *EditEmailRequest, opts ...grpc.CallOption) (*EditEmailResponse, error) {
	out := new(EditEmailResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.account.proto.Accounts/EditEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.account.proto.Accounts/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequest, opts ...grpc.CallOption) (*GetPrivateKeyResponse, error) {
	out := new(GetPrivateKeyResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.account.proto.Accounts/GetPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
// All implementations should embed UnimplementedAccountsServer
// for forward compatibility
type AccountsServer interface {
	// IsExist checks if the account exists. The account name(username) is specified in the request.
	// If the account exists, it returns true. Otherwise, it returns false.
	IsExist(context.Context, *IsExistRequest) (*IsExistResponse, error)
	// Create a new account.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get the account by username.
	GetByUsername(context.Context, *GetByUsernameRequest) (*GetByUsernameResponse, error)
	// Delete the account.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// EditUsername Update the account username.
	// will change the account username and the preferred_username in Actor.
	EditUsername(context.Context, *EditUsernameRequest) (*EditUsernameResponse, error)
	// EditPassword Update the account password.
	EditPassword(context.Context, *EditPasswordRequest) (*EditPasswordResponse, error)
	// EditEmail Edit the unique email for the account.
	EditEmail(context.Context, *EditEmailRequest) (*EditEmailResponse, error)
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	GetPrivateKey(context.Context, *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error)
}

// UnimplementedAccountsServer should be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (UnimplementedAccountsServer) IsExist(context.Context, *IsExistRequest) (*IsExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsExist not implemented")
}
func (UnimplementedAccountsServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccountsServer) GetByUsername(context.Context, *GetByUsernameRequest) (*GetByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUsername not implemented")
}
func (UnimplementedAccountsServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
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
func (UnimplementedAccountsServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAccountsServer) GetPrivateKey(context.Context, *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateKey not implemented")
}

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
		FullMethod: "/hvx.api.v1alpha1.account.proto.Accounts/IsExist",
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
		FullMethod: "/hvx.api.v1alpha1.account.proto.Accounts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.account.proto.Accounts/GetByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetByUsername(ctx, req.(*GetByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.account.proto.Accounts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Delete(ctx, req.(*DeleteRequest))
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
		FullMethod: "/hvx.api.v1alpha1.account.proto.Accounts/EditUsername",
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
		FullMethod: "/hvx.api.v1alpha1.account.proto.Accounts/EditPassword",
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
		FullMethod: "/hvx.api.v1alpha1.account.proto.Accounts/EditEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).EditEmail(ctx, req.(*EditEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.account.proto.Accounts/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.account.proto.Accounts/GetPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetPrivateKey(ctx, req.(*GetPrivateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Accounts_ServiceDesc is the grpc.ServiceDesc for Accounts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Accounts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvx.api.v1alpha1.account.proto.Accounts",
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
			MethodName: "GetByUsername",
			Handler:    _Accounts_GetByUsername_Handler,
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
			MethodName: "EditEmail",
			Handler:    _Accounts_EditEmail_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Accounts_Verify_Handler,
		},
		{
			MethodName: "GetPrivateKey",
			Handler:    _Accounts_GetPrivateKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/account/account.proto",
}
