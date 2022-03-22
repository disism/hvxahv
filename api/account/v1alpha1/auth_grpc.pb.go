// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/account/v1alpha1/auth.proto

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// Verify Authentication Sign in with a valid user.
	// When the authentication is successful.
	// The client ID must be registered to the device table
	// So you must submit a UA identifier in addition to the username and password when you sign in.
	// A valid Token is returned,
	// This returned TOKEN must be carried in subsequent API access operations.
	// Returns a device id which is generated by the server as a hash using the uuid.
	// Returns a public key.
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	// GetPublicKeyByAccountUsername obtaining the public key by account name.
	// Verify the account.
	// Use this method when verifying the signature (ActivityPub) or when exchanging keys.
	GetPublicKeyByAccountUsername(ctx context.Context, in *GetPublicKeyByAccountUsernameRequest, opts ...grpc.CallOption) (*GetPublicKeyByAccountUsernameResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Auth/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetPublicKeyByAccountUsername(ctx context.Context, in *GetPublicKeyByAccountUsernameRequest, opts ...grpc.CallOption) (*GetPublicKeyByAccountUsernameResponse, error) {
	out := new(GetPublicKeyByAccountUsernameResponse)
	err := c.cc.Invoke(ctx, "/hvxahv.v1alpha1.proto.Auth/GetPublicKeyByAccountUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// Verify Authentication Sign in with a valid user.
	// When the authentication is successful.
	// The client ID must be registered to the device table
	// So you must submit a UA identifier in addition to the username and password when you sign in.
	// A valid Token is returned,
	// This returned TOKEN must be carried in subsequent API access operations.
	// Returns a device id which is generated by the server as a hash using the uuid.
	// Returns a public key.
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	// GetPublicKeyByAccountUsername obtaining the public key by account name.
	// Verify the account.
	// Use this method when verifying the signature (ActivityPub) or when exchanging keys.
	GetPublicKeyByAccountUsername(context.Context, *GetPublicKeyByAccountUsernameRequest) (*GetPublicKeyByAccountUsernameResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAuthServer) GetPublicKeyByAccountUsername(context.Context, *GetPublicKeyByAccountUsernameRequest) (*GetPublicKeyByAccountUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKeyByAccountUsername not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Auth/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetPublicKeyByAccountUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyByAccountUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetPublicKeyByAccountUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvxahv.v1alpha1.proto.Auth/GetPublicKeyByAccountUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetPublicKeyByAccountUsername(ctx, req.(*GetPublicKeyByAccountUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvxahv.v1alpha1.proto.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _Auth_Verify_Handler,
		},
		{
			MethodName: "GetPublicKeyByAccountUsername",
			Handler:    _Auth_GetPublicKeyByAccountUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/account/v1alpha1/auth.proto",
}
