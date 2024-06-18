// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: user_info.proto

package sign

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserInfoRespond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Avatar   string `protobuf:"bytes,5,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Gender   string `protobuf:"bytes,6,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Nickname string `protobuf:"bytes,7,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
}

func (x *UserInfoRespond) Reset() {
	*x = UserInfoRespond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRespond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRespond) ProtoMessage() {}

func (x *UserInfoRespond) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRespond.ProtoReflect.Descriptor instead.
func (*UserInfoRespond) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfoRespond) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfoRespond) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfoRespond) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserInfoRespond) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfoRespond) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfoRespond) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserInfoRespond) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

var File_user_info_proto protoreflect.FileDescriptor

var file_user_info_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0x29, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x73, 0x69, 0x67, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_info_proto_rawDescOnce sync.Once
	file_user_info_proto_rawDescData = file_user_info_proto_rawDesc
)

func file_user_info_proto_rawDescGZIP() []byte {
	file_user_info_proto_rawDescOnce.Do(func() {
		file_user_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_info_proto_rawDescData)
	})
	return file_user_info_proto_rawDescData
}

var file_user_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_info_proto_goTypes = []interface{}{
	(*UserInfoRequest)(nil), // 0: sign.UserInfoRequest
	(*UserInfoRespond)(nil), // 1: sign.UserInfoRespond
}
var file_user_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_info_proto_init() }
func file_user_info_proto_init() {
	if File_user_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRequest); i {
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
		file_user_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRespond); i {
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
			RawDescriptor: file_user_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_info_proto_goTypes,
		DependencyIndexes: file_user_info_proto_depIdxs,
		MessageInfos:      file_user_info_proto_msgTypes,
	}.Build()
	File_user_info_proto = out.File
	file_user_info_proto_rawDesc = nil
	file_user_info_proto_goTypes = nil
	file_user_info_proto_depIdxs = nil
}
