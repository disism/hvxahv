// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: api/activity/v1alpha1/activity.proto

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

type Inboxes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId      string `protobuf:"bytes,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	ActivityId   string `protobuf:"bytes,2,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	ActivityType string `protobuf:"bytes,3,opt,name=activity_type,json=activityType,proto3" json:"activity_type,omitempty"`
	ActivityBody []byte `protobuf:"bytes,4,opt,name=activity_body,json=activityBody,proto3" json:"activity_body,omitempty"`
}

func (x *Inboxes) Reset() {
	*x = Inboxes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inboxes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inboxes) ProtoMessage() {}

func (x *Inboxes) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inboxes.ProtoReflect.Descriptor instead.
func (*Inboxes) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{0}
}

func (x *Inboxes) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

func (x *Inboxes) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *Inboxes) GetActivityType() string {
	if x != nil {
		return x.ActivityType
	}
	return ""
}

func (x *Inboxes) GetActivityBody() []byte {
	if x != nil {
		return x.ActivityBody
	}
	return nil
}

type CreateInboxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateInboxRequest) Reset() {
	*x = CreateInboxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInboxRequest) ProtoMessage() {}

func (x *CreateInboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInboxRequest.ProtoReflect.Descriptor instead.
func (*CreateInboxRequest) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInboxRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateInboxRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateInboxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateInboxResponse) Reset() {
	*x = CreateInboxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInboxResponse) ProtoMessage() {}

func (x *CreateInboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInboxResponse.ProtoReflect.Descriptor instead.
func (*CreateInboxResponse) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{2}
}

func (x *CreateInboxResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateInboxResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type GetInboxByActivityIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId string `protobuf:"bytes,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
}

func (x *GetInboxByActivityIDRequest) Reset() {
	*x = GetInboxByActivityIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInboxByActivityIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInboxByActivityIDRequest) ProtoMessage() {}

func (x *GetInboxByActivityIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInboxByActivityIDRequest.ProtoReflect.Descriptor instead.
func (*GetInboxByActivityIDRequest) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{3}
}

func (x *GetInboxByActivityIDRequest) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

type GetInboxByActivityIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Inbox *Inboxes `protobuf:"bytes,2,opt,name=inbox,proto3" json:"inbox,omitempty"`
}

func (x *GetInboxByActivityIDResponse) Reset() {
	*x = GetInboxByActivityIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInboxByActivityIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInboxByActivityIDResponse) ProtoMessage() {}

func (x *GetInboxByActivityIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInboxByActivityIDResponse.ProtoReflect.Descriptor instead.
func (*GetInboxByActivityIDResponse) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{4}
}

func (x *GetInboxByActivityIDResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetInboxByActivityIDResponse) GetInbox() *Inboxes {
	if x != nil {
		return x.Inbox
	}
	return nil
}

type GetInboxesByActivityIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetInboxesByActivityIDRequest) Reset() {
	*x = GetInboxesByActivityIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInboxesByActivityIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInboxesByActivityIDRequest) ProtoMessage() {}

func (x *GetInboxesByActivityIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInboxesByActivityIDRequest.ProtoReflect.Descriptor instead.
func (*GetInboxesByActivityIDRequest) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{5}
}

func (x *GetInboxesByActivityIDRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type GetInboxesByActivityIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string     `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Inboxes []*Inboxes `protobuf:"bytes,2,rep,name=inboxes,proto3" json:"inboxes,omitempty"`
}

func (x *GetInboxesByActivityIDResponse) Reset() {
	*x = GetInboxesByActivityIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInboxesByActivityIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInboxesByActivityIDResponse) ProtoMessage() {}

func (x *GetInboxesByActivityIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInboxesByActivityIDResponse.ProtoReflect.Descriptor instead.
func (*GetInboxesByActivityIDResponse) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{6}
}

func (x *GetInboxesByActivityIDResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetInboxesByActivityIDResponse) GetInboxes() []*Inboxes {
	if x != nil {
		return x.Inboxes
	}
	return nil
}

type CreateOutboxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOutboxRequest) Reset() {
	*x = CreateOutboxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOutboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOutboxRequest) ProtoMessage() {}

func (x *CreateOutboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOutboxRequest.ProtoReflect.Descriptor instead.
func (*CreateOutboxRequest) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{7}
}

type CreateOutboxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOutboxResponse) Reset() {
	*x = CreateOutboxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOutboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOutboxResponse) ProtoMessage() {}

func (x *CreateOutboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOutboxResponse.ProtoReflect.Descriptor instead.
func (*CreateOutboxResponse) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{8}
}

type GetOutboxByActivityIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOutboxByActivityIDRequest) Reset() {
	*x = GetOutboxByActivityIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOutboxByActivityIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOutboxByActivityIDRequest) ProtoMessage() {}

func (x *GetOutboxByActivityIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOutboxByActivityIDRequest.ProtoReflect.Descriptor instead.
func (*GetOutboxByActivityIDRequest) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{9}
}

type GetOutboxByActivityIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOutboxByActivityIDResponse) Reset() {
	*x = GetOutboxByActivityIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOutboxByActivityIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOutboxByActivityIDResponse) ProtoMessage() {}

func (x *GetOutboxByActivityIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_activity_v1alpha1_activity_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOutboxByActivityIDResponse.ProtoReflect.Descriptor instead.
func (*GetOutboxByActivityIDResponse) Descriptor() ([]byte, []int) {
	return file_api_activity_v1alpha1_activity_proto_rawDescGZIP(), []int{10}
}

var File_api_activity_v1alpha1_activity_proto protoreflect.FileDescriptor

var file_api_activity_v1alpha1_activity_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01,
	0x0a, 0x07, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42, 0x6f, 0x64, 0x79, 0x22,
	0x3c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78,
	0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78,
	0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x69, 0x6e, 0x62, 0x6f,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x05, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x22, 0x3e,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x42, 0x79, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x6e,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x42, 0x79, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x62, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x07, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x22, 0x15,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x75, 0x74, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a,
	0x1d, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf2,
	0x04, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x66, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x29, 0x2e, 0x68, 0x76, 0x78,
	0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x81, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78,
	0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x32, 0x2e, 0x68,
	0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x42, 0x79, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f,
	0x78, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x44, 0x12, 0x34, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x62, 0x6f, 0x78, 0x65, 0x73, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68,
	0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x42, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x69, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f,
	0x78, 0x12, 0x2a, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x62,
	0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x33, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x68, 0x76,
	0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x42, 0x79, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_activity_v1alpha1_activity_proto_rawDescOnce sync.Once
	file_api_activity_v1alpha1_activity_proto_rawDescData = file_api_activity_v1alpha1_activity_proto_rawDesc
)

func file_api_activity_v1alpha1_activity_proto_rawDescGZIP() []byte {
	file_api_activity_v1alpha1_activity_proto_rawDescOnce.Do(func() {
		file_api_activity_v1alpha1_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_activity_v1alpha1_activity_proto_rawDescData)
	})
	return file_api_activity_v1alpha1_activity_proto_rawDescData
}

var file_api_activity_v1alpha1_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_activity_v1alpha1_activity_proto_goTypes = []interface{}{
	(*Inboxes)(nil),                        // 0: hvxahv.v1alpha1.proto.Inboxes
	(*CreateInboxRequest)(nil),             // 1: hvxahv.v1alpha1.proto.CreateInboxRequest
	(*CreateInboxResponse)(nil),            // 2: hvxahv.v1alpha1.proto.CreateInboxResponse
	(*GetInboxByActivityIDRequest)(nil),    // 3: hvxahv.v1alpha1.proto.GetInboxByActivityIDRequest
	(*GetInboxByActivityIDResponse)(nil),   // 4: hvxahv.v1alpha1.proto.GetInboxByActivityIDResponse
	(*GetInboxesByActivityIDRequest)(nil),  // 5: hvxahv.v1alpha1.proto.GetInboxesByActivityIDRequest
	(*GetInboxesByActivityIDResponse)(nil), // 6: hvxahv.v1alpha1.proto.GetInboxesByActivityIDResponse
	(*CreateOutboxRequest)(nil),            // 7: hvxahv.v1alpha1.proto.CreateOutboxRequest
	(*CreateOutboxResponse)(nil),           // 8: hvxahv.v1alpha1.proto.CreateOutboxResponse
	(*GetOutboxByActivityIDRequest)(nil),   // 9: hvxahv.v1alpha1.proto.GetOutboxByActivityIDRequest
	(*GetOutboxByActivityIDResponse)(nil),  // 10: hvxahv.v1alpha1.proto.GetOutboxByActivityIDResponse
}
var file_api_activity_v1alpha1_activity_proto_depIdxs = []int32{
	0,  // 0: hvxahv.v1alpha1.proto.GetInboxByActivityIDResponse.inbox:type_name -> hvxahv.v1alpha1.proto.Inboxes
	0,  // 1: hvxahv.v1alpha1.proto.GetInboxesByActivityIDResponse.inboxes:type_name -> hvxahv.v1alpha1.proto.Inboxes
	1,  // 2: hvxahv.v1alpha1.proto.Activity.CreateInbox:input_type -> hvxahv.v1alpha1.proto.CreateInboxRequest
	3,  // 3: hvxahv.v1alpha1.proto.Activity.GetInboxByActivityID:input_type -> hvxahv.v1alpha1.proto.GetInboxByActivityIDRequest
	5,  // 4: hvxahv.v1alpha1.proto.Activity.GetInboxesByActivityID:input_type -> hvxahv.v1alpha1.proto.GetInboxesByActivityIDRequest
	7,  // 5: hvxahv.v1alpha1.proto.Activity.CreateOutbox:input_type -> hvxahv.v1alpha1.proto.CreateOutboxRequest
	9,  // 6: hvxahv.v1alpha1.proto.Activity.GetOutboxByActivityID:input_type -> hvxahv.v1alpha1.proto.GetOutboxByActivityIDRequest
	2,  // 7: hvxahv.v1alpha1.proto.Activity.CreateInbox:output_type -> hvxahv.v1alpha1.proto.CreateInboxResponse
	4,  // 8: hvxahv.v1alpha1.proto.Activity.GetInboxByActivityID:output_type -> hvxahv.v1alpha1.proto.GetInboxByActivityIDResponse
	6,  // 9: hvxahv.v1alpha1.proto.Activity.GetInboxesByActivityID:output_type -> hvxahv.v1alpha1.proto.GetInboxesByActivityIDResponse
	8,  // 10: hvxahv.v1alpha1.proto.Activity.CreateOutbox:output_type -> hvxahv.v1alpha1.proto.CreateOutboxResponse
	10, // 11: hvxahv.v1alpha1.proto.Activity.GetOutboxByActivityID:output_type -> hvxahv.v1alpha1.proto.GetOutboxByActivityIDResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_activity_v1alpha1_activity_proto_init() }
func file_api_activity_v1alpha1_activity_proto_init() {
	if File_api_activity_v1alpha1_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_activity_v1alpha1_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inboxes); i {
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
		file_api_activity_v1alpha1_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInboxRequest); i {
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
		file_api_activity_v1alpha1_activity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInboxResponse); i {
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
		file_api_activity_v1alpha1_activity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInboxByActivityIDRequest); i {
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
		file_api_activity_v1alpha1_activity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInboxByActivityIDResponse); i {
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
		file_api_activity_v1alpha1_activity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInboxesByActivityIDRequest); i {
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
		file_api_activity_v1alpha1_activity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInboxesByActivityIDResponse); i {
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
		file_api_activity_v1alpha1_activity_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOutboxRequest); i {
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
		file_api_activity_v1alpha1_activity_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOutboxResponse); i {
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
		file_api_activity_v1alpha1_activity_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOutboxByActivityIDRequest); i {
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
		file_api_activity_v1alpha1_activity_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOutboxByActivityIDResponse); i {
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
			RawDescriptor: file_api_activity_v1alpha1_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_activity_v1alpha1_activity_proto_goTypes,
		DependencyIndexes: file_api_activity_v1alpha1_activity_proto_depIdxs,
		MessageInfos:      file_api_activity_v1alpha1_activity_proto_msgTypes,
	}.Build()
	File_api_activity_v1alpha1_activity_proto = out.File
	file_api_activity_v1alpha1_activity_proto_rawDesc = nil
	file_api_activity_v1alpha1_activity_proto_goTypes = nil
	file_api_activity_v1alpha1_activity_proto_depIdxs = nil
}
