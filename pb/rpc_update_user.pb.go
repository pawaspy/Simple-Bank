// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: rpc_update_user.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName      *string                `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3,oneof" json:"full_name,omitempty"`
	Email         *string                `protobuf:"bytes,3,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Password      *string                `protobuf:"bytes,4,opt,name=password,proto3,oneof" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_rpc_update_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_update_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_rpc_update_user_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateUserRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	mi := &file_rpc_update_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_update_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_rpc_update_user_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_rpc_update_user_proto protoreflect.FileDescriptor

const file_rpc_update_user_proto_rawDesc = "" +
	"\n" +
	"\x15rpc_update_user.proto\x12\x02pb\x1a\n" +
	"user.proto\"\xb2\x01\n" +
	"\x11UpdateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12 \n" +
	"\tfull_name\x18\x02 \x01(\tH\x00R\bfullName\x88\x01\x01\x12\x19\n" +
	"\x05email\x18\x03 \x01(\tH\x01R\x05email\x88\x01\x01\x12\x1f\n" +
	"\bpassword\x18\x04 \x01(\tH\x02R\bpassword\x88\x01\x01B\f\n" +
	"\n" +
	"_full_nameB\b\n" +
	"\x06_emailB\v\n" +
	"\t_password\"2\n" +
	"\x12UpdateUserResponse\x12\x1c\n" +
	"\x04user\x18\x01 \x01(\v2\b.pb.UserR\x04userB#Z!github.com/pawaspy/simple_bank/pbb\x06proto3"

var (
	file_rpc_update_user_proto_rawDescOnce sync.Once
	file_rpc_update_user_proto_rawDescData []byte
)

func file_rpc_update_user_proto_rawDescGZIP() []byte {
	file_rpc_update_user_proto_rawDescOnce.Do(func() {
		file_rpc_update_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_update_user_proto_rawDesc), len(file_rpc_update_user_proto_rawDesc)))
	})
	return file_rpc_update_user_proto_rawDescData
}

var file_rpc_update_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_update_user_proto_goTypes = []any{
	(*UpdateUserRequest)(nil),  // 0: pb.UpdateUserRequest
	(*UpdateUserResponse)(nil), // 1: pb.UpdateUserResponse
	(*User)(nil),               // 2: pb.User
}
var file_rpc_update_user_proto_depIdxs = []int32{
	2, // 0: pb.UpdateUserResponse.user:type_name -> pb.User
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_update_user_proto_init() }
func file_rpc_update_user_proto_init() {
	if File_rpc_update_user_proto != nil {
		return
	}
	file_user_proto_init()
	file_rpc_update_user_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_update_user_proto_rawDesc), len(file_rpc_update_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_update_user_proto_goTypes,
		DependencyIndexes: file_rpc_update_user_proto_depIdxs,
		MessageInfos:      file_rpc_update_user_proto_msgTypes,
	}.Build()
	File_rpc_update_user_proto = out.File
	file_rpc_update_user_proto_goTypes = nil
	file_rpc_update_user_proto_depIdxs = nil
}
