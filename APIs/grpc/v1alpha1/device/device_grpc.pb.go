// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/v1alpha1/device/device.proto

package device

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DevicesClient is the client API for Devices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevicesClient interface {
	// IsExist Gets whether the device is authenticated by the device id.
	// This method returns a boolean value. Returns true if it does not exist
	// Otherwise returns false.
	IsExist(ctx context.Context, in *IsExistRequest, opts ...grpc.CallOption) (*IsExistResponse, error)
	// Create To create the logged-in device data, pass the account id
	// and user agent and generate a hash as a unique identifier
	// for the device.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get Get detailed data about the device by its unique device ID.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Device, error)
	// Get a list of all logged-in devices by account ID.
	GetDevices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDevicesResponse, error)
	// Delete a device by its device id.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// DeleteDeviceByWithAccountID Delete all login devices by account id.
	DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*DeleteDevicesResponse, error)
}

type devicesClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicesClient(cc grpc.ClientConnInterface) DevicesClient {
	return &devicesClient{cc}
}

func (c *devicesClient) IsExist(ctx context.Context, in *IsExistRequest, opts ...grpc.CallOption) (*IsExistResponse, error) {
	out := new(IsExistResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.device.proto.Devices/IsExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.device.proto.Devices/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.device.proto.Devices/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) GetDevices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDevicesResponse, error) {
	out := new(GetDevicesResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.device.proto.Devices/GetDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.device.proto.Devices/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*DeleteDevicesResponse, error) {
	out := new(DeleteDevicesResponse)
	err := c.cc.Invoke(ctx, "/hvx.api.v1alpha1.device.proto.Devices/DeleteDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServer is the server API for Devices service.
// All implementations should embed UnimplementedDevicesServer
// for forward compatibility
type DevicesServer interface {
	// IsExist Gets whether the device is authenticated by the device id.
	// This method returns a boolean value. Returns true if it does not exist
	// Otherwise returns false.
	IsExist(context.Context, *IsExistRequest) (*IsExistResponse, error)
	// Create To create the logged-in device data, pass the account id
	// and user agent and generate a hash as a unique identifier
	// for the device.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get Get detailed data about the device by its unique device ID.
	Get(context.Context, *GetRequest) (*Device, error)
	// Get a list of all logged-in devices by account ID.
	GetDevices(context.Context, *emptypb.Empty) (*GetDevicesResponse, error)
	// Delete a device by its device id.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// DeleteDeviceByWithAccountID Delete all login devices by account id.
	DeleteDevices(context.Context, *DeleteDevicesRequest) (*DeleteDevicesResponse, error)
}

// UnimplementedDevicesServer should be embedded to have forward compatible implementations.
type UnimplementedDevicesServer struct {
}

func (UnimplementedDevicesServer) IsExist(context.Context, *IsExistRequest) (*IsExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsExist not implemented")
}
func (UnimplementedDevicesServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDevicesServer) Get(context.Context, *GetRequest) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDevicesServer) GetDevices(context.Context, *emptypb.Empty) (*GetDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}
func (UnimplementedDevicesServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDevicesServer) DeleteDevices(context.Context, *DeleteDevicesRequest) (*DeleteDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevices not implemented")
}

// UnsafeDevicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevicesServer will
// result in compilation errors.
type UnsafeDevicesServer interface {
	mustEmbedUnimplementedDevicesServer()
}

func RegisterDevicesServer(s grpc.ServiceRegistrar, srv DevicesServer) {
	s.RegisterService(&Devices_ServiceDesc, srv)
}

func _Devices_IsExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).IsExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.device.proto.Devices/IsExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).IsExist(ctx, req.(*IsExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.device.proto.Devices/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.device.proto.Devices/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_GetDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).GetDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.device.proto.Devices/GetDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).GetDevices(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.device.proto.Devices/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_DeleteDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).DeleteDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hvx.api.v1alpha1.device.proto.Devices/DeleteDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).DeleteDevices(ctx, req.(*DeleteDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Devices_ServiceDesc is the grpc.ServiceDesc for Devices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Devices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hvx.api.v1alpha1.device.proto.Devices",
	HandlerType: (*DevicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsExist",
			Handler:    _Devices_IsExist_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Devices_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Devices_Get_Handler,
		},
		{
			MethodName: "GetDevices",
			Handler:    _Devices_GetDevices_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Devices_Delete_Handler,
		},
		{
			MethodName: "DeleteDevices",
			Handler:    _Devices_DeleteDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/device/device.proto",
}
