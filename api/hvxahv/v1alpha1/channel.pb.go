// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: api/hvxahv/v1alpha1/channel.proto

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

type ChannelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Bio       string `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	Avatar    string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IsPrivate bool   `protobuf:"varint,5,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
}

func (x *ChannelData) Reset() {
	*x = ChannelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1alpha1_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelData) ProtoMessage() {}

func (x *ChannelData) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1alpha1_channel_proto_msgTypes[0]
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
	return file_api_hvxahv_v1alpha1_channel_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChannelData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelData) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *ChannelData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ChannelData) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hvxahv_v1alpha1_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_api_hvxahv_v1alpha1_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_api_hvxahv_v1alpha1_channel_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_hvxahv_v1alpha1_channel_proto protoreflect.FileDescriptor

var file_api_hvxahv_v1alpha1_channel_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0b, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x35, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x5b, 0x0a, 0x07,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x50, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x22, 0x2e, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1c, 0x2e, 0x68, 0x76, 0x78, 0x61,
	0x68, 0x76, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x73, 0x69, 0x73, 0x6d, 0x2f, 0x68,
	0x76, 0x78, 0x61, 0x68, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x76, 0x78, 0x61, 0x68, 0x76,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_hvxahv_v1alpha1_channel_proto_rawDescOnce sync.Once
	file_api_hvxahv_v1alpha1_channel_proto_rawDescData = file_api_hvxahv_v1alpha1_channel_proto_rawDesc
)

func file_api_hvxahv_v1alpha1_channel_proto_rawDescGZIP() []byte {
	file_api_hvxahv_v1alpha1_channel_proto_rawDescOnce.Do(func() {
		file_api_hvxahv_v1alpha1_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_hvxahv_v1alpha1_channel_proto_rawDescData)
	})
	return file_api_hvxahv_v1alpha1_channel_proto_rawDescData
}

var file_api_hvxahv_v1alpha1_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_hvxahv_v1alpha1_channel_proto_goTypes = []interface{}{
	(*ChannelData)(nil), // 0: hvxahv.v1alpha1.proto.ChannelData
	(*Reply)(nil),       // 1: hvxahv.v1alpha1.proto.Reply
}
var file_api_hvxahv_v1alpha1_channel_proto_depIdxs = []int32{
	0, // 0: hvxahv.v1alpha1.proto.Channel.NewChannel:input_type -> hvxahv.v1alpha1.proto.ChannelData
	1, // 1: hvxahv.v1alpha1.proto.Channel.NewChannel:output_type -> hvxahv.v1alpha1.proto.Reply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_hvxahv_v1alpha1_channel_proto_init() }
func file_api_hvxahv_v1alpha1_channel_proto_init() {
	if File_api_hvxahv_v1alpha1_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_hvxahv_v1alpha1_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_hvxahv_v1alpha1_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
			RawDescriptor: file_api_hvxahv_v1alpha1_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_hvxahv_v1alpha1_channel_proto_goTypes,
		DependencyIndexes: file_api_hvxahv_v1alpha1_channel_proto_depIdxs,
		MessageInfos:      file_api_hvxahv_v1alpha1_channel_proto_msgTypes,
	}.Build()
	File_api_hvxahv_v1alpha1_channel_proto = out.File
	file_api_hvxahv_v1alpha1_channel_proto_rawDesc = nil
	file_api_hvxahv_v1alpha1_channel_proto_goTypes = nil
	file_api_hvxahv_v1alpha1_channel_proto_depIdxs = nil
}
