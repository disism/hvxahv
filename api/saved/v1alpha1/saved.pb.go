// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: api/saved/v1alpha1/saved.proto

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

type Save struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Cid         string `protobuf:"bytes,5,opt,name=cid,proto3" json:"cid,omitempty"`
	Types       string `protobuf:"bytes,6,opt,name=types,proto3" json:"types,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Save) Reset() {
	*x = Save{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Save) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Save) ProtoMessage() {}

func (x *Save) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Save.ProtoReflect.Descriptor instead.
func (*Save) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{0}
}

func (x *Save) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Save) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Save) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Save) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *Save) GetTypes() string {
	if x != nil {
		return x.Types
	}
	return ""
}

func (x *Save) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateSavedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId   string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Cid         string `protobuf:"bytes,4,opt,name=cid,proto3" json:"cid,omitempty"`
	Types       string `protobuf:"bytes,5,opt,name=types,proto3" json:"types,omitempty"`
	IsPrivate   bool   `protobuf:"varint,6,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
}

func (x *CreateSavedRequest) Reset() {
	*x = CreateSavedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSavedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSavedRequest) ProtoMessage() {}

func (x *CreateSavedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSavedRequest.ProtoReflect.Descriptor instead.
func (*CreateSavedRequest) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSavedRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CreateSavedRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSavedRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateSavedRequest) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *CreateSavedRequest) GetTypes() string {
	if x != nil {
		return x.Types
	}
	return ""
}

func (x *CreateSavedRequest) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

type CreateSavedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *CreateSavedResponse) Reset() {
	*x = CreateSavedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSavedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSavedResponse) ProtoMessage() {}

func (x *CreateSavedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSavedResponse.ProtoReflect.Descriptor instead.
func (*CreateSavedResponse) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSavedResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateSavedResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type GetSavesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetSavesRequest) Reset() {
	*x = GetSavesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSavesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSavesRequest) ProtoMessage() {}

func (x *GetSavesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSavesRequest.ProtoReflect.Descriptor instead.
func (*GetSavesRequest) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{3}
}

func (x *GetSavesRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type GetSavesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string  `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Saves []*Save `protobuf:"bytes,3,rep,name=saves,proto3" json:"saves,omitempty"`
}

func (x *GetSavesResponse) Reset() {
	*x = GetSavesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSavesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSavesResponse) ProtoMessage() {}

func (x *GetSavesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSavesResponse.ProtoReflect.Descriptor instead.
func (*GetSavesResponse) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{4}
}

func (x *GetSavesResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetSavesResponse) GetSaves() []*Save {
	if x != nil {
		return x.Saves
	}
	return nil
}

type GetSavedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSavedRequest) Reset() {
	*x = GetSavedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSavedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSavedRequest) ProtoMessage() {}

func (x *GetSavedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSavedRequest.ProtoReflect.Descriptor instead.
func (*GetSavedRequest) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{5}
}

func (x *GetSavedRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *GetSavedRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EditSavedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EditSavedRequest) Reset() {
	*x = EditSavedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditSavedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditSavedRequest) ProtoMessage() {}

func (x *EditSavedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditSavedRequest.ProtoReflect.Descriptor instead.
func (*EditSavedRequest) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{6}
}

func (x *EditSavedRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EditSavedRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *EditSavedRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EditSavedRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type EditSavedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *EditSavedResponse) Reset() {
	*x = EditSavedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditSavedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditSavedResponse) ProtoMessage() {}

func (x *EditSavedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditSavedResponse.ProtoReflect.Descriptor instead.
func (*EditSavedResponse) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{7}
}

func (x *EditSavedResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *EditSavedResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type DeleteSavedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *DeleteSavedRequest) Reset() {
	*x = DeleteSavedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSavedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSavedRequest) ProtoMessage() {}

func (x *DeleteSavedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSavedRequest.ProtoReflect.Descriptor instead.
func (*DeleteSavedRequest) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteSavedRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteSavedRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type DeleteSavedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *DeleteSavedResponse) Reset() {
	*x = DeleteSavedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSavedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSavedResponse) ProtoMessage() {}

func (x *DeleteSavedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSavedResponse.ProtoReflect.Descriptor instead.
func (*DeleteSavedResponse) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteSavedResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeleteSavedResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type DeleteAllSavesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *DeleteAllSavesRequest) Reset() {
	*x = DeleteAllSavesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllSavesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllSavesRequest) ProtoMessage() {}

func (x *DeleteAllSavesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllSavesRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllSavesRequest) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteAllSavesRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type DeleteAllSavesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Reply string `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *DeleteAllSavesResponse) Reset() {
	*x = DeleteAllSavesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllSavesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllSavesResponse) ProtoMessage() {}

func (x *DeleteAllSavesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_saved_v1alpha1_saved_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllSavesResponse.ProtoReflect.Descriptor instead.
func (*DeleteAllSavesResponse) Descriptor() ([]byte, []int) {
	return file_api_saved_v1alpha1_saved_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteAllSavesResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DeleteAllSavesResponse) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_api_saved_v1alpha1_saved_proto protoreflect.FileDescriptor

var file_api_saved_v1alpha1_saved_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xaf, 0x01,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22,
	0x3f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x59, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x61,
	0x76, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x68, 0x76, 0x78, 0x61,
	0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x05, 0x73, 0x61, 0x76, 0x65, 0x73, 0x22, 0x40, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x77, 0x0a, 0x10, 0x45, 0x64, 0x69, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x11, 0x45, 0x64, 0x69, 0x74,
	0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x43, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x36, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x53, 0x61, 0x76, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x6c, 0x6c, 0x53, 0x61, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xdc, 0x04, 0x0a, 0x05, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x12, 0x66, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x76,
	0x65, 0x64, 0x12, 0x29, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x26, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a,
	0x09, 0x45, 0x64, 0x69, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x27, 0x2e, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x66, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x64, 0x12, 0x29,
	0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x76,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x68, 0x76, 0x78, 0x61,
	0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x76, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x53, 0x61, 0x76, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x68, 0x76, 0x78, 0x61,
	0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x53, 0x61, 0x76, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x53, 0x61, 0x76, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x68, 0x76,
	0x78, 0x61, 0x68, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_saved_v1alpha1_saved_proto_rawDescOnce sync.Once
	file_api_saved_v1alpha1_saved_proto_rawDescData = file_api_saved_v1alpha1_saved_proto_rawDesc
)

func file_api_saved_v1alpha1_saved_proto_rawDescGZIP() []byte {
	file_api_saved_v1alpha1_saved_proto_rawDescOnce.Do(func() {
		file_api_saved_v1alpha1_saved_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_saved_v1alpha1_saved_proto_rawDescData)
	})
	return file_api_saved_v1alpha1_saved_proto_rawDescData
}

var file_api_saved_v1alpha1_saved_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_saved_v1alpha1_saved_proto_goTypes = []interface{}{
	(*Save)(nil),                   // 0: hvxahv.v1alpha1.proto.Save
	(*CreateSavedRequest)(nil),     // 1: hvxahv.v1alpha1.proto.CreateSavedRequest
	(*CreateSavedResponse)(nil),    // 2: hvxahv.v1alpha1.proto.CreateSavedResponse
	(*GetSavesRequest)(nil),        // 3: hvxahv.v1alpha1.proto.GetSavesRequest
	(*GetSavesResponse)(nil),       // 4: hvxahv.v1alpha1.proto.GetSavesResponse
	(*GetSavedRequest)(nil),        // 5: hvxahv.v1alpha1.proto.GetSavedRequest
	(*EditSavedRequest)(nil),       // 6: hvxahv.v1alpha1.proto.EditSavedRequest
	(*EditSavedResponse)(nil),      // 7: hvxahv.v1alpha1.proto.EditSavedResponse
	(*DeleteSavedRequest)(nil),     // 8: hvxahv.v1alpha1.proto.DeleteSavedRequest
	(*DeleteSavedResponse)(nil),    // 9: hvxahv.v1alpha1.proto.DeleteSavedResponse
	(*DeleteAllSavesRequest)(nil),  // 10: hvxahv.v1alpha1.proto.DeleteAllSavesRequest
	(*DeleteAllSavesResponse)(nil), // 11: hvxahv.v1alpha1.proto.DeleteAllSavesResponse
}
var file_api_saved_v1alpha1_saved_proto_depIdxs = []int32{
	0,  // 0: hvxahv.v1alpha1.proto.GetSavesResponse.saves:type_name -> hvxahv.v1alpha1.proto.Save
	1,  // 1: hvxahv.v1alpha1.proto.Saved.CreateSaved:input_type -> hvxahv.v1alpha1.proto.CreateSavedRequest
	3,  // 2: hvxahv.v1alpha1.proto.Saved.GetSaves:input_type -> hvxahv.v1alpha1.proto.GetSavesRequest
	5,  // 3: hvxahv.v1alpha1.proto.Saved.GetSaved:input_type -> hvxahv.v1alpha1.proto.GetSavedRequest
	6,  // 4: hvxahv.v1alpha1.proto.Saved.EditSaved:input_type -> hvxahv.v1alpha1.proto.EditSavedRequest
	8,  // 5: hvxahv.v1alpha1.proto.Saved.DeleteSaved:input_type -> hvxahv.v1alpha1.proto.DeleteSavedRequest
	10, // 6: hvxahv.v1alpha1.proto.Saved.DeleteAllSaves:input_type -> hvxahv.v1alpha1.proto.DeleteAllSavesRequest
	2,  // 7: hvxahv.v1alpha1.proto.Saved.CreateSaved:output_type -> hvxahv.v1alpha1.proto.CreateSavedResponse
	4,  // 8: hvxahv.v1alpha1.proto.Saved.GetSaves:output_type -> hvxahv.v1alpha1.proto.GetSavesResponse
	0,  // 9: hvxahv.v1alpha1.proto.Saved.GetSaved:output_type -> hvxahv.v1alpha1.proto.Save
	7,  // 10: hvxahv.v1alpha1.proto.Saved.EditSaved:output_type -> hvxahv.v1alpha1.proto.EditSavedResponse
	9,  // 11: hvxahv.v1alpha1.proto.Saved.DeleteSaved:output_type -> hvxahv.v1alpha1.proto.DeleteSavedResponse
	11, // 12: hvxahv.v1alpha1.proto.Saved.DeleteAllSaves:output_type -> hvxahv.v1alpha1.proto.DeleteAllSavesResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_saved_v1alpha1_saved_proto_init() }
func file_api_saved_v1alpha1_saved_proto_init() {
	if File_api_saved_v1alpha1_saved_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_saved_v1alpha1_saved_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Save); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSavedRequest); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSavedResponse); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSavesRequest); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSavesResponse); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSavedRequest); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditSavedRequest); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditSavedResponse); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSavedRequest); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSavedResponse); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllSavesRequest); i {
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
		file_api_saved_v1alpha1_saved_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllSavesResponse); i {
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
			RawDescriptor: file_api_saved_v1alpha1_saved_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_saved_v1alpha1_saved_proto_goTypes,
		DependencyIndexes: file_api_saved_v1alpha1_saved_proto_depIdxs,
		MessageInfos:      file_api_saved_v1alpha1_saved_proto_msgTypes,
	}.Build()
	File_api_saved_v1alpha1_saved_proto = out.File
	file_api_saved_v1alpha1_saved_proto_rawDesc = nil
	file_api_saved_v1alpha1_saved_proto_goTypes = nil
	file_api_saved_v1alpha1_saved_proto_depIdxs = nil
}
