// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/v1alpha1/device/device.proto

package device

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Device    string `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Device) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Device) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Device) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type IsExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IsExistRequest) Reset() {
	*x = IsExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsExistRequest) ProtoMessage() {}

func (x *IsExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsExistRequest.ProtoReflect.Descriptor instead.
func (*IsExistRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{1}
}

func (x *IsExistRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IsExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExist bool `protobuf:"varint,1,opt,name=is_exist,json=isExist,proto3" json:"is_exist,omitempty"`
}

func (x *IsExistResponse) Reset() {
	*x = IsExistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsExistResponse) ProtoMessage() {}

func (x *IsExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsExistResponse.ProtoReflect.Descriptor instead.
func (*IsExistResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{2}
}

func (x *IsExistResponse) GetIsExist() bool {
	if x != nil {
		return x.IsExist
	}
	return false
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Ua        string `protobuf:"bytes,2,opt,name=ua,proto3" json:"ua,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CreateRequest) GetUa() string {
	if x != nil {
		return x.Ua
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{4}
}

func (x *CreateResponse) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{5}
}

func (x *GetRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type GetDevicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string    `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Devices []*Device `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *GetDevicesResponse) Reset() {
	*x = GetDevicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDevicesResponse) ProtoMessage() {}

func (x *GetDevicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDevicesResponse.ProtoReflect.Descriptor instead.
func (*GetDevicesResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{6}
}

func (x *GetDevicesResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetDevicesResponse) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type DeleteDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *DeleteDevicesRequest) Reset() {
	*x = DeleteDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDevicesRequest) ProtoMessage() {}

func (x *DeleteDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDevicesRequest.ProtoReflect.Descriptor instead.
func (*DeleteDevicesRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDevicesRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type DeleteDevicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *DeleteDevicesResponse) Reset() {
	*x = DeleteDevicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDevicesResponse) ProtoMessage() {}

func (x *DeleteDevicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDevicesResponse.ProtoReflect.Descriptor instead.
func (*DeleteDevicesResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteDevicesResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeleteDevicesResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_device_device_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_device_device_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_device_device_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeleteResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_proto_v1alpha1_device_device_proto protoreflect.FileDescriptor

var file_proto_v1alpha1_device_device_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x68, 0x76, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6e, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x20, 0x0a, 0x0e, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x22,
	0x3e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x75, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x75, 0x61, 0x22,
	0x2d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x29,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x15, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2c,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xb3, 0x05, 0x0a, 0x07, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x6a, 0x0a, 0x07, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12,
	0x2d, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x68, 0x76, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x67, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x68, 0x76, 0x78,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x29, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x76,
	0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x31, 0x2e, 0x68, 0x76, 0x78,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x80, 0x01,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x0e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x7c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x33, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2f, 0x68, 0x76, 0x78, 0x2f, 0x41, 0x50, 0x49, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1alpha1_device_device_proto_rawDescOnce sync.Once
	file_proto_v1alpha1_device_device_proto_rawDescData = file_proto_v1alpha1_device_device_proto_rawDesc
)

func file_proto_v1alpha1_device_device_proto_rawDescGZIP() []byte {
	file_proto_v1alpha1_device_device_proto_rawDescOnce.Do(func() {
		file_proto_v1alpha1_device_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1alpha1_device_device_proto_rawDescData)
	})
	return file_proto_v1alpha1_device_device_proto_rawDescData
}

var file_proto_v1alpha1_device_device_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_v1alpha1_device_device_proto_goTypes = []interface{}{
	(*Device)(nil),                // 0: hvx.api.v1alpha1.device.proto.Device
	(*IsExistRequest)(nil),        // 1: hvx.api.v1alpha1.device.proto.IsExistRequest
	(*IsExistResponse)(nil),       // 2: hvx.api.v1alpha1.device.proto.IsExistResponse
	(*CreateRequest)(nil),         // 3: hvx.api.v1alpha1.device.proto.CreateRequest
	(*CreateResponse)(nil),        // 4: hvx.api.v1alpha1.device.proto.CreateResponse
	(*GetRequest)(nil),            // 5: hvx.api.v1alpha1.device.proto.GetRequest
	(*GetDevicesResponse)(nil),    // 6: hvx.api.v1alpha1.device.proto.GetDevicesResponse
	(*DeleteDevicesRequest)(nil),  // 7: hvx.api.v1alpha1.device.proto.DeleteDevicesRequest
	(*DeleteDevicesResponse)(nil), // 8: hvx.api.v1alpha1.device.proto.DeleteDevicesResponse
	(*DeleteRequest)(nil),         // 9: hvx.api.v1alpha1.device.proto.DeleteRequest
	(*DeleteResponse)(nil),        // 10: hvx.api.v1alpha1.device.proto.DeleteResponse
	(*emptypb.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_proto_v1alpha1_device_device_proto_depIdxs = []int32{
	0,  // 0: hvx.api.v1alpha1.device.proto.GetDevicesResponse.devices:type_name -> hvx.api.v1alpha1.device.proto.Device
	1,  // 1: hvx.api.v1alpha1.device.proto.Devices.IsExist:input_type -> hvx.api.v1alpha1.device.proto.IsExistRequest
	3,  // 2: hvx.api.v1alpha1.device.proto.Devices.Create:input_type -> hvx.api.v1alpha1.device.proto.CreateRequest
	5,  // 3: hvx.api.v1alpha1.device.proto.Devices.Get:input_type -> hvx.api.v1alpha1.device.proto.GetRequest
	11, // 4: hvx.api.v1alpha1.device.proto.Devices.GetDevices:input_type -> google.protobuf.Empty
	9,  // 5: hvx.api.v1alpha1.device.proto.Devices.Delete:input_type -> hvx.api.v1alpha1.device.proto.DeleteRequest
	7,  // 6: hvx.api.v1alpha1.device.proto.Devices.DeleteDevices:input_type -> hvx.api.v1alpha1.device.proto.DeleteDevicesRequest
	2,  // 7: hvx.api.v1alpha1.device.proto.Devices.IsExist:output_type -> hvx.api.v1alpha1.device.proto.IsExistResponse
	4,  // 8: hvx.api.v1alpha1.device.proto.Devices.Create:output_type -> hvx.api.v1alpha1.device.proto.CreateResponse
	0,  // 9: hvx.api.v1alpha1.device.proto.Devices.Get:output_type -> hvx.api.v1alpha1.device.proto.Device
	6,  // 10: hvx.api.v1alpha1.device.proto.Devices.GetDevices:output_type -> hvx.api.v1alpha1.device.proto.GetDevicesResponse
	10, // 11: hvx.api.v1alpha1.device.proto.Devices.Delete:output_type -> hvx.api.v1alpha1.device.proto.DeleteResponse
	8,  // 12: hvx.api.v1alpha1.device.proto.Devices.DeleteDevices:output_type -> hvx.api.v1alpha1.device.proto.DeleteDevicesResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_v1alpha1_device_device_proto_init() }
func file_proto_v1alpha1_device_device_proto_init() {
	if File_proto_v1alpha1_device_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1alpha1_device_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_device_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsExistRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_device_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsExistResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_device_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_device_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_device_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_device_device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDevicesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_device_device_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDevicesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_device_device_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDevicesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_device_device_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_device_device_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1alpha1_device_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1alpha1_device_device_proto_goTypes,
		DependencyIndexes: file_proto_v1alpha1_device_device_proto_depIdxs,
		MessageInfos:      file_proto_v1alpha1_device_device_proto_msgTypes,
	}.Build()
	File_proto_v1alpha1_device_device_proto = out.File
	file_proto_v1alpha1_device_device_proto_rawDesc = nil
	file_proto_v1alpha1_device_device_proto_goTypes = nil
	file_proto_v1alpha1_device_device_proto_depIdxs = nil
}
