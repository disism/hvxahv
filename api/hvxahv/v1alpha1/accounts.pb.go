// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: api/hvxahv/v1alpha1/accounts.proto

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

// AccountsData List of all data of the account service
type AccountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Mail       string `protobuf:"bytes,4,opt,name=Mail,proto3" json:"Mail,omitempty"`
	Avatar     string `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Bio        string `protobuf:"bytes,6,opt,name=Bio,proto3" json:"Bio,omitempty"`
	Name       string `protobuf:"bytes,7,opt,name=Name,proto3" json:"Name,omitempty"`
	Phone      string `protobuf:"bytes,8,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Private    int32  `protobuf:"varint,9,opt,name=Private,proto3" json:"Private,omitempty"`
	PrivateKey string `protobuf:"bytes,10,opt,name=PrivateKey,proto3" json:"PrivateKey,omitempty"`
	PublicKey  string `protobuf:"bytes,11,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
}

func (x *AccountData) Reset() {
	*x = AccountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountData) ProtoMessage() {}

func (x *AccountData) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountData.ProtoReflect.Descriptor instead.
func (*AccountData) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1alpha1_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *AccountData) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AccountData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AccountData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AccountData) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *AccountData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *AccountData) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *AccountData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountData) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AccountData) GetPrivate() int32 {
	if x != nil {
		return x.Private
	}
	return 0
}

func (x *AccountData) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *AccountData) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// AccountsReply The return value, code and message of the Accounts operation.
type AccountsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *AccountsReply) Reset() {
	*x = AccountsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsReply) ProtoMessage() {}

func (x *AccountsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsReply.ProtoReflect.Descriptor instead.
func (*AccountsReply) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1alpha1_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *AccountsReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AccountsReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// AccountByName Functions operated by username.
type AccountByName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *AccountByName) Reset() {
	*x = AccountByName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountByName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountByName) ProtoMessage() {}

func (x *AccountByName) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountByName.ProtoReflect.Descriptor instead.
func (*AccountByName) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1alpha1_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *AccountByName) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// NewAccountData Data needed to create a new account.
type NewAccountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Mail     string `protobuf:"bytes,3,opt,name=Mail,proto3" json:"Mail,omitempty"`
}

func (x *NewAccountData) Reset() {
	*x = NewAccountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAccountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAccountData) ProtoMessage() {}

func (x *NewAccountData) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAccountData.ProtoReflect.Descriptor instead.
func (*NewAccountData) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1alpha1_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *NewAccountData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NewAccountData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *NewAccountData) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

// LoginData Account login, use email and password.
type LoginData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail     string `protobuf:"bytes,1,opt,name=Mail,proto3" json:"Mail,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginData) Reset() {
	*x = LoginData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginData) ProtoMessage() {}

func (x *LoginData) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginData.ProtoReflect.Descriptor instead.
func (*LoginData) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1alpha1_accounts_proto_rawDescGZIP(), []int{4}
}

func (x *LoginData) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *LoginData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// LoginReply The data returned by the login account needs to return a token.
type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1alpha1_accounts_proto_rawDescGZIP(), []int{5}
}

func (x *LoginReply) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_api_hvxahv_v1alpha1_accounts_proto protoreflect.FileDescriptor

var file_api_hvxahv_v1alpha1_accounts_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x0b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x55,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x42, 0x69, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x3d, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x4d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d, 0x61, 0x69,
	0x6c, 0x22, 0x3b, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3e,
	0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xd5,
	0x03, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x5b, 0x0a, 0x0a, 0x4e,
	0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x68, 0x76, 0x78, 0x61,
	0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x24, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x68, 0x76, 0x78, 0x61,
	0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x24, 0x2e,
	0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x24, 0x2e, 0x68,
	0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x22, 0x2e, 0x68, 0x76, 0x78, 0x61,
	0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x20, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x21, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x73, 0x69, 0x73, 0x6d, 0x2f, 0x68, 0x76, 0x78, 0x61,
	0x68, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_hvxahv_v1alpha1_accounts_proto_rawDescOnce sync.Once
	file_api_hvxahv_v1alpha1_accounts_proto_rawDescData = file_api_hvxahv_v1alpha1_accounts_proto_rawDesc
)

func file_api_hvxahv_v1alpha1_accounts_proto_rawDescGZIP() []byte {
	file_api_hvxahv_v1alpha1_accounts_proto_rawDescOnce.Do(func() {
		file_api_hvxahv_v1alpha1_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_hvxahv_v1alpha1_accounts_proto_rawDescData)
	})
	return file_api_hvxahv_v1alpha1_accounts_proto_rawDescData
}

var file_api_hvxahv_v1alpha1_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_hvxahv_v1alpha1_accounts_proto_goTypes = []interface{}{
	(*AccountData)(nil),    // 0: hvxahv.v1alpha1.proto.AccountData
	(*AccountsReply)(nil),  // 1: hvxahv.v1alpha1.proto.AccountsReply
	(*AccountByName)(nil),  // 2: hvxahv.v1alpha1.proto.AccountByName
	(*NewAccountData)(nil), // 3: hvxahv.v1alpha1.proto.NewAccountData
	(*LoginData)(nil),      // 4: hvxahv.v1alpha1.proto.LoginData
	(*LoginReply)(nil),     // 5: hvxahv.v1alpha1.proto.LoginReply
}
var file_api_hvxahv_v1alpha1_accounts_proto_depIdxs = []int32{
	3, // 0: hvxahv.v1alpha1.proto.Accounts.NewAccount:input_type -> hvxahv.v1alpha1.proto.NewAccountData
	0, // 1: hvxahv.v1alpha1.proto.Accounts.UpdateAccount:input_type -> hvxahv.v1alpha1.proto.AccountData
	2, // 2: hvxahv.v1alpha1.proto.Accounts.DeleteAccount:input_type -> hvxahv.v1alpha1.proto.AccountByName
	2, // 3: hvxahv.v1alpha1.proto.Accounts.FindAccount:input_type -> hvxahv.v1alpha1.proto.AccountByName
	4, // 4: hvxahv.v1alpha1.proto.Accounts.LoginAccount:input_type -> hvxahv.v1alpha1.proto.LoginData
	1, // 5: hvxahv.v1alpha1.proto.Accounts.NewAccount:output_type -> hvxahv.v1alpha1.proto.AccountsReply
	1, // 6: hvxahv.v1alpha1.proto.Accounts.UpdateAccount:output_type -> hvxahv.v1alpha1.proto.AccountsReply
	1, // 7: hvxahv.v1alpha1.proto.Accounts.DeleteAccount:output_type -> hvxahv.v1alpha1.proto.AccountsReply
	0, // 8: hvxahv.v1alpha1.proto.Accounts.FindAccount:output_type -> hvxahv.v1alpha1.proto.AccountData
	5, // 9: hvxahv.v1alpha1.proto.Accounts.LoginAccount:output_type -> hvxahv.v1alpha1.proto.LoginReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_hvxahv_v1alpha1_accounts_proto_init() }
func file_api_hvxahv_v1alpha1_accounts_proto_init() {
	if File_api_hvxahv_v1alpha1_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountData); i {
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
		file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountsReply); i {
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
		file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountByName); i {
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
		file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAccountData); i {
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
		file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginData); i {
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
		file_api_hvxahv_v1alpha1_accounts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
			RawDescriptor: file_api_hvxahv_v1alpha1_accounts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_hvxahv_v1alpha1_accounts_proto_goTypes,
		DependencyIndexes: file_api_hvxahv_v1alpha1_accounts_proto_depIdxs,
		MessageInfos:      file_api_hvxahv_v1alpha1_accounts_proto_msgTypes,
	}.Build()
	File_api_hvxahv_v1alpha1_accounts_proto = out.File
	file_api_hvxahv_v1alpha1_accounts_proto_rawDesc = nil
	file_api_hvxahv_v1alpha1_accounts_proto_goTypes = nil
	file_api_hvxahv_v1alpha1_accounts_proto_depIdxs = nil
}
