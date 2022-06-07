// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: proto/channel/v1alpha1/broadcast.proto

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

type BroadcastData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	AdminId   string `protobuf:"bytes,3,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	ArticleId string `protobuf:"bytes,4,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
	Cid       string `protobuf:"bytes,5,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *BroadcastData) Reset() {
	*x = BroadcastData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastData) ProtoMessage() {}

func (x *BroadcastData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastData.ProtoReflect.Descriptor instead.
func (*BroadcastData) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_broadcast_proto_rawDescGZIP(), []int{0}
}

func (x *BroadcastData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BroadcastData) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *BroadcastData) GetAdminId() string {
	if x != nil {
		return x.AdminId
	}
	return ""
}

func (x *BroadcastData) GetArticleId() string {
	if x != nil {
		return x.ArticleId
	}
	return ""
}

func (x *BroadcastData) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type CreateBroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	AdminId   string `protobuf:"bytes,2,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	ArticleId string `protobuf:"bytes,3,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
}

func (x *CreateBroadcastRequest) Reset() {
	*x = CreateBroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBroadcastRequest) ProtoMessage() {}

func (x *CreateBroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBroadcastRequest.ProtoReflect.Descriptor instead.
func (*CreateBroadcastRequest) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_broadcast_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBroadcastRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *CreateBroadcastRequest) GetAdminId() string {
	if x != nil {
		return x.AdminId
	}
	return ""
}

func (x *CreateBroadcastRequest) GetArticleId() string {
	if x != nil {
		return x.ArticleId
	}
	return ""
}

type CreateBroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *CreateBroadcastResponse) Reset() {
	*x = CreateBroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBroadcastResponse) ProtoMessage() {}

func (x *CreateBroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBroadcastResponse.ProtoReflect.Descriptor instead.
func (*CreateBroadcastResponse) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_broadcast_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBroadcastResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateBroadcastResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type GetAllBroadcastsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *GetAllBroadcastsRequest) Reset() {
	*x = GetAllBroadcastsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBroadcastsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBroadcastsRequest) ProtoMessage() {}

func (x *GetAllBroadcastsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBroadcastsRequest.ProtoReflect.Descriptor instead.
func (*GetAllBroadcastsRequest) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_broadcast_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllBroadcastsRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type GetAllBroadcastsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       string           `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Broadcasts []*BroadcastData `protobuf:"bytes,2,rep,name=broadcasts,proto3" json:"broadcasts,omitempty"`
}

func (x *GetAllBroadcastsResponse) Reset() {
	*x = GetAllBroadcastsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBroadcastsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBroadcastsResponse) ProtoMessage() {}

func (x *GetAllBroadcastsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBroadcastsResponse.ProtoReflect.Descriptor instead.
func (*GetAllBroadcastsResponse) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_broadcast_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllBroadcastsResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetAllBroadcastsResponse) GetBroadcasts() []*BroadcastData {
	if x != nil {
		return x.Broadcasts
	}
	return nil
}

type DeleteBroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId   string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	AdminId     string `protobuf:"bytes,2,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	BroadcastId string `protobuf:"bytes,3,opt,name=broadcast_id,json=broadcastId,proto3" json:"broadcast_id,omitempty"`
}

func (x *DeleteBroadcastRequest) Reset() {
	*x = DeleteBroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBroadcastRequest) ProtoMessage() {}

func (x *DeleteBroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBroadcastRequest.ProtoReflect.Descriptor instead.
func (*DeleteBroadcastRequest) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_broadcast_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBroadcastRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *DeleteBroadcastRequest) GetAdminId() string {
	if x != nil {
		return x.AdminId
	}
	return ""
}

func (x *DeleteBroadcastRequest) GetBroadcastId() string {
	if x != nil {
		return x.BroadcastId
	}
	return ""
}

type DeleteBroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *DeleteBroadcastResponse) Reset() {
	*x = DeleteBroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBroadcastResponse) ProtoMessage() {}

func (x *DeleteBroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_channel_v1alpha1_broadcast_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBroadcastResponse.ProtoReflect.Descriptor instead.
func (*DeleteBroadcastResponse) Descriptor() ([]byte, []int) {
	return file_proto_channel_v1alpha1_broadcast_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBroadcastResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeleteBroadcastResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_proto_channel_v1alpha1_broadcast_proto protoreflect.FileDescriptor

var file_proto_channel_v1alpha1_broadcast_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69,
	0x64, 0x22, 0x71, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x38, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x73, 0x22, 0x75,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x88, 0x03, 0x0a, 0x09, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x7c, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x32, 0x2e, 0x68, 0x76,
	0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x73, 0x12, 0x33, 0x2e, 0x68, 0x76, 0x78,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x34, 0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x32, 0x2e, 0x68, 0x76, 0x78,
	0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x68, 0x76, 0x78, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_channel_v1alpha1_broadcast_proto_rawDescOnce sync.Once
	file_proto_channel_v1alpha1_broadcast_proto_rawDescData = file_proto_channel_v1alpha1_broadcast_proto_rawDesc
)

func file_proto_channel_v1alpha1_broadcast_proto_rawDescGZIP() []byte {
	file_proto_channel_v1alpha1_broadcast_proto_rawDescOnce.Do(func() {
		file_proto_channel_v1alpha1_broadcast_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_channel_v1alpha1_broadcast_proto_rawDescData)
	})
	return file_proto_channel_v1alpha1_broadcast_proto_rawDescData
}

var file_proto_channel_v1alpha1_broadcast_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_channel_v1alpha1_broadcast_proto_goTypes = []interface{}{
	(*BroadcastData)(nil),            // 0: hvx.channel.v1alpha1.proto.BroadcastData
	(*CreateBroadcastRequest)(nil),   // 1: hvx.channel.v1alpha1.proto.CreateBroadcastRequest
	(*CreateBroadcastResponse)(nil),  // 2: hvx.channel.v1alpha1.proto.CreateBroadcastResponse
	(*GetAllBroadcastsRequest)(nil),  // 3: hvx.channel.v1alpha1.proto.GetAllBroadcastsRequest
	(*GetAllBroadcastsResponse)(nil), // 4: hvx.channel.v1alpha1.proto.GetAllBroadcastsResponse
	(*DeleteBroadcastRequest)(nil),   // 5: hvx.channel.v1alpha1.proto.DeleteBroadcastRequest
	(*DeleteBroadcastResponse)(nil),  // 6: hvx.channel.v1alpha1.proto.DeleteBroadcastResponse
}
var file_proto_channel_v1alpha1_broadcast_proto_depIdxs = []int32{
	0, // 0: hvx.channel.v1alpha1.proto.GetAllBroadcastsResponse.broadcasts:type_name -> hvx.channel.v1alpha1.proto.BroadcastData
	1, // 1: hvx.channel.v1alpha1.proto.Broadcast.CreateBroadcast:input_type -> hvx.channel.v1alpha1.proto.CreateBroadcastRequest
	3, // 2: hvx.channel.v1alpha1.proto.Broadcast.GetAllBroadcasts:input_type -> hvx.channel.v1alpha1.proto.GetAllBroadcastsRequest
	5, // 3: hvx.channel.v1alpha1.proto.Broadcast.DeleteBroadcast:input_type -> hvx.channel.v1alpha1.proto.DeleteBroadcastRequest
	2, // 4: hvx.channel.v1alpha1.proto.Broadcast.CreateBroadcast:output_type -> hvx.channel.v1alpha1.proto.CreateBroadcastResponse
	4, // 5: hvx.channel.v1alpha1.proto.Broadcast.GetAllBroadcasts:output_type -> hvx.channel.v1alpha1.proto.GetAllBroadcastsResponse
	6, // 6: hvx.channel.v1alpha1.proto.Broadcast.DeleteBroadcast:output_type -> hvx.channel.v1alpha1.proto.DeleteBroadcastResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_channel_v1alpha1_broadcast_proto_init() }
func file_proto_channel_v1alpha1_broadcast_proto_init() {
	if File_proto_channel_v1alpha1_broadcast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_channel_v1alpha1_broadcast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastData); i {
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
		file_proto_channel_v1alpha1_broadcast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBroadcastRequest); i {
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
		file_proto_channel_v1alpha1_broadcast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBroadcastResponse); i {
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
		file_proto_channel_v1alpha1_broadcast_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBroadcastsRequest); i {
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
		file_proto_channel_v1alpha1_broadcast_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBroadcastsResponse); i {
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
		file_proto_channel_v1alpha1_broadcast_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBroadcastRequest); i {
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
		file_proto_channel_v1alpha1_broadcast_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBroadcastResponse); i {
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
			RawDescriptor: file_proto_channel_v1alpha1_broadcast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_channel_v1alpha1_broadcast_proto_goTypes,
		DependencyIndexes: file_proto_channel_v1alpha1_broadcast_proto_depIdxs,
		MessageInfos:      file_proto_channel_v1alpha1_broadcast_proto_msgTypes,
	}.Build()
	File_proto_channel_v1alpha1_broadcast_proto = out.File
	file_proto_channel_v1alpha1_broadcast_proto_rawDesc = nil
	file_proto_channel_v1alpha1_broadcast_proto_goTypes = nil
	file_proto_channel_v1alpha1_broadcast_proto_depIdxs = nil
}