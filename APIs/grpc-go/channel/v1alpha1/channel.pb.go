// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: proto/channel/v1alpha1/channel.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreferredUsername string `protobuf:"bytes,1,opt,name=preferredUsername,proto3" json:"preferredUsername,omitempty"`
	AccountId         string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *CreateChannelRequest) Reset() {
	*x = CreateChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelRequest) ProtoMessage() {}

func (x *CreateChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelRequest.ProtoReflect.Descriptor instead.
func (*CreateChannelRequest) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_channel_proto_rawDescGZIP(), []int{0}
}

func (x *CreateChannelRequest) GetPreferredUsername() string {
	if x != nil {
		return x.PreferredUsername
	}
	return ""
}

func (x *CreateChannelRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type CreateChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *CreateChannelResponse) Reset() {
	*x = CreateChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChannelResponse) ProtoMessage() {}

func (x *CreateChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChannelResponse.ProtoReflect.Descriptor instead.
func (*CreateChannelResponse) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_channel_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChannelResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateChannelResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type GetChannelsByAccountIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetChannelsByAccountIDRequest) Reset() {
	*x = GetChannelsByAccountIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChannelsByAccountIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelsByAccountIDRequest) ProtoMessage() {}

func (x *GetChannelsByAccountIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelsByAccountIDRequest.ProtoReflect.Descriptor instead.
func (*GetChannelsByAccountIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_channel_proto_rawDescGZIP(), []int{2}
}

func (x *GetChannelsByAccountIDRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type ChannelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *ChannelData) Reset() {
	*x = ChannelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelData) ProtoMessage() {}

func (x *ChannelData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelData.ProtoReflect.Descriptor instead.
func (*ChannelData) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_channel_proto_rawDescGZIP(), []int{3}
}

func (x *ChannelData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChannelData) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type GetChannelsByAccountIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string         `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Channels []*ChannelData `protobuf:"bytes,2,rep,name=channels,proto3" json:"channels,omitempty"`
}

func (x *GetChannelsByAccountIDResponse) Reset() {
	*x = GetChannelsByAccountIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChannelsByAccountIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelsByAccountIDResponse) ProtoMessage() {}

func (x *GetChannelsByAccountIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelsByAccountIDResponse.ProtoReflect.Descriptor instead.
func (*GetChannelsByAccountIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_channel_proto_rawDescGZIP(), []int{4}
}

func (x *GetChannelsByAccountIDResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetChannelsByAccountIDResponse) GetChannels() []*ChannelData {
	if x != nil {
		return x.Channels
	}
	return nil
}

type DeleteChannelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *DeleteChannelRequest) Reset() {
	*x = DeleteChannelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChannelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChannelRequest) ProtoMessage() {}

func (x *DeleteChannelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChannelRequest.ProtoReflect.Descriptor instead.
func (*DeleteChannelRequest) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_channel_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteChannelRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *DeleteChannelRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type DeleteChannelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *DeleteChannelResponse) Reset() {
	*x = DeleteChannelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChannelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChannelResponse) ProtoMessage() {}

func (x *DeleteChannelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChannelResponse.ProtoReflect.Descriptor instead.
func (*DeleteChannelResponse) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_channel_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteChannelResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeleteChannelResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type DeleteAllChannelsByAccountIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *DeleteAllChannelsByAccountIDRequest) Reset() {
	*x = DeleteAllChannelsByAccountIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllChannelsByAccountIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllChannelsByAccountIDRequest) ProtoMessage() {}

func (x *DeleteAllChannelsByAccountIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllChannelsByAccountIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllChannelsByAccountIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_channel_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAllChannelsByAccountIDRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type DeleteAllChannelsByAccountIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *DeleteAllChannelsByAccountIDResponse) Reset() {
	*x = DeleteAllChannelsByAccountIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllChannelsByAccountIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllChannelsByAccountIDResponse) ProtoMessage() {}

func (x *DeleteAllChannelsByAccountIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_channel_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllChannelsByAccountIDResponse.ProtoReflect.Descriptor instead.
func (*DeleteAllChannelsByAccountIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_channel_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAllChannelsByAccountIDResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeleteAllChannelsByAccountIDResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_proto_channel_v1alpha1_channel_proto protoreflect.FileDescriptor

var file_proto_channel_v1alpha1_channel_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x63, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3e, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0b, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x43,
	0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x73, 0x22, 0x54, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x44, 0x0a, 0x23,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x50, 0x0a, 0x24, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x32, 0xb3, 0x04, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x76, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x12, 0x30, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x91, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x39, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x42, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x30, 0x2e,
	0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0xa3, 0x01, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x6c, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x3f, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_channel_v1alpha1_channel_proto_rawDescOnce sync.Once
	file_proto_channel_v1alpha1_channel_proto_rawDescData = file_proto_channel_v1alpha1_channel_proto_rawDesc
)

func file_proto_channel_v1alpha1_channel_proto_rawDescGZIP() []byte {
	file_proto_channel_v1alpha1_channel_proto_rawDescOnce.Do(func() {
		file_proto_channel_v1alpha1_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_channel_v1alpha1_channel_proto_rawDescData)
	})
	return file_proto_channel_v1alpha1_channel_proto_rawDescData
}

var file_proto_channel_v1alpha1_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_channel_v1alpha1_channel_proto_goTypes = []interface{}{
	(*CreateChannelRequest)(nil),                 // 0: hvx.channel.v1alpha1.proto.CreateChannelRequest
	(*CreateChannelResponse)(nil),                // 1: hvx.channel.v1alpha1.proto.CreateChannelResponse
	(*GetChannelsByAccountIDRequest)(nil),        // 2: hvx.channel.v1alpha1.proto.GetChannelsByAccountIDRequest
	(*ChannelData)(nil),                          // 3: hvx.channel.v1alpha1.proto.ChannelData
	(*GetChannelsByAccountIDResponse)(nil),       // 4: hvx.channel.v1alpha1.proto.GetChannelsByAccountIDResponse
	(*DeleteChannelRequest)(nil),                 // 5: hvx.channel.v1alpha1.proto.DeleteChannelRequest
	(*DeleteChannelResponse)(nil),                // 6: hvx.channel.v1alpha1.proto.DeleteChannelResponse
	(*DeleteAllChannelsByAccountIDRequest)(nil),  // 7: hvx.channel.v1alpha1.proto.DeleteAllChannelsByAccountIDRequest
	(*DeleteAllChannelsByAccountIDResponse)(nil), // 8: hvx.channel.v1alpha1.proto.DeleteAllChannelsByAccountIDResponse
}
var file_proto_channel_v1alpha1_channel_proto_depIdxs = []int32{
	3, // 0: hvx.channel.v1alpha1.proto.GetChannelsByAccountIDResponse.channels:type_name -> hvx.channel.v1alpha1.proto.ChannelData
	0, // 1: hvx.channel.v1alpha1.proto.Channel.CreateChannel:input_type -> hvx.channel.v1alpha1.proto.CreateChannelRequest
	2, // 2: hvx.channel.v1alpha1.proto.Channel.GetChannelsByAccountID:input_type -> hvx.channel.v1alpha1.proto.GetChannelsByAccountIDRequest
	5, // 3: hvx.channel.v1alpha1.proto.Channel.DeleteChannel:input_type -> hvx.channel.v1alpha1.proto.DeleteChannelRequest
	7, // 4: hvx.channel.v1alpha1.proto.Channel.DeleteAllChannelsByAccountID:input_type -> hvx.channel.v1alpha1.proto.DeleteAllChannelsByAccountIDRequest
	1, // 5: hvx.channel.v1alpha1.proto.Channel.CreateChannel:output_type -> hvx.channel.v1alpha1.proto.CreateChannelResponse
	4, // 6: hvx.channel.v1alpha1.proto.Channel.GetChannelsByAccountID:output_type -> hvx.channel.v1alpha1.proto.GetChannelsByAccountIDResponse
	6, // 7: hvx.channel.v1alpha1.proto.Channel.DeleteChannel:output_type -> hvx.channel.v1alpha1.proto.DeleteChannelResponse
	8, // 8: hvx.channel.v1alpha1.proto.Channel.DeleteAllChannelsByAccountID:output_type -> hvx.channel.v1alpha1.proto.DeleteAllChannelsByAccountIDResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_channel_v1alpha1_channel_proto_init() }
func file_proto_channel_v1alpha1_channel_proto_init() {
	if File_proto_channel_v1alpha1_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_channel_v1alpha1_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChannelRequest); i {
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
		file_proto_channel_v1alpha1_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChannelResponse); i {
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
		file_proto_channel_v1alpha1_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChannelsByAccountIDRequest); i {
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
		file_proto_channel_v1alpha1_channel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelData); i {
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
		file_proto_channel_v1alpha1_channel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChannelsByAccountIDResponse); i {
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
		file_proto_channel_v1alpha1_channel_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChannelRequest); i {
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
		file_proto_channel_v1alpha1_channel_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChannelResponse); i {
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
		file_proto_channel_v1alpha1_channel_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllChannelsByAccountIDRequest); i {
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
		file_proto_channel_v1alpha1_channel_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllChannelsByAccountIDResponse); i {
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
			RawDescriptor: file_proto_channel_v1alpha1_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_channel_v1alpha1_channel_proto_goTypes,
		DependencyIndexes: file_proto_channel_v1alpha1_channel_proto_depIdxs,
		MessageInfos:      file_proto_channel_v1alpha1_channel_proto_msgTypes,
	}.Build()
	File_proto_channel_v1alpha1_channel_proto = out.File
	file_proto_channel_v1alpha1_channel_proto_rawDesc = nil
	file_proto_channel_v1alpha1_channel_proto_goTypes = nil
	file_proto_channel_v1alpha1_channel_proto_depIdxs = nil
}