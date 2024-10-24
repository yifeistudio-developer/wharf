// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: alps.proto

package alps

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

type CredentialType int32

const (
	CredentialType_ACCESS_TOKEN CredentialType = 0
	CredentialType_TICKET       CredentialType = 1
)

// Enum value maps for CredentialType.
var (
	CredentialType_name = map[int32]string{
		0: "ACCESS_TOKEN",
		1: "TICKET",
	}
	CredentialType_value = map[string]int32{
		"ACCESS_TOKEN": 0,
		"TICKET":       1,
	}
)

func (x CredentialType) Enum() *CredentialType {
	p := new(CredentialType)
	*p = x
	return p
}

func (x CredentialType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CredentialType) Descriptor() protoreflect.EnumDescriptor {
	return file_alps_proto_enumTypes[0].Descriptor()
}

func (CredentialType) Type() protoreflect.EnumType {
	return &file_alps_proto_enumTypes[0]
}

func (x CredentialType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CredentialType.Descriptor instead.
func (CredentialType) EnumDescriptor() ([]byte, []int) {
	return file_alps_proto_rawDescGZIP(), []int{0}
}

// 凭证
type CredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credential string         `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
	Type       CredentialType `protobuf:"varint,2,opt,name=type,proto3,enum=grpc.CredentialType" json:"type,omitempty"`
}

func (x *CredentialRequest) Reset() {
	*x = CredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialRequest) ProtoMessage() {}

func (x *CredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialRequest.ProtoReflect.Descriptor instead.
func (*CredentialRequest) Descriptor() ([]byte, []int) {
	return file_alps_proto_rawDescGZIP(), []int{0}
}

func (x *CredentialRequest) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

func (x *CredentialRequest) GetType() CredentialType {
	if x != nil {
		return x.Type
	}
	return CredentialType_ACCESS_TOKEN
}

// 账户概要信息
type PrincipalsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32                       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg        string                      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Principals *PrincipalsReply_Principals `protobuf:"bytes,3,opt,name=principals,proto3" json:"principals,omitempty"`
}

func (x *PrincipalsReply) Reset() {
	*x = PrincipalsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrincipalsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrincipalsReply) ProtoMessage() {}

func (x *PrincipalsReply) ProtoReflect() protoreflect.Message {
	mi := &file_alps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrincipalsReply.ProtoReflect.Descriptor instead.
func (*PrincipalsReply) Descriptor() ([]byte, []int) {
	return file_alps_proto_rawDescGZIP(), []int{1}
}

func (x *PrincipalsReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PrincipalsReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *PrincipalsReply) GetPrincipals() *PrincipalsReply_Principals {
	if x != nil {
		return x.Principals
	}
	return nil
}

type PrincipalsReply_Principals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 账号
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 昵称
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// 头像
	Avatar string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *PrincipalsReply_Principals) Reset() {
	*x = PrincipalsReply_Principals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alps_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrincipalsReply_Principals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrincipalsReply_Principals) ProtoMessage() {}

func (x *PrincipalsReply_Principals) ProtoReflect() protoreflect.Message {
	mi := &file_alps_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrincipalsReply_Principals.ProtoReflect.Descriptor instead.
func (*PrincipalsReply_Principals) Descriptor() ([]byte, []int) {
	return file_alps_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PrincipalsReply_Principals) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PrincipalsReply_Principals) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PrincipalsReply_Principals) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

var File_alps_proto protoreflect.FileDescriptor

var file_alps_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x6c, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x22, 0x5d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0xd7, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x40, 0x0a, 0x0a, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x73, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x1a, 0x5c, 0x0a,
	0x0a, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2a, 0x2e, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x0c, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x01, 0x32, 0x61, 0x0a, 0x15, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x34,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x69, 0x66,
	0x65, 0x69, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x72, 0x2f, 0x77, 0x68, 0x61, 0x72, 0x66, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x61, 0x6c, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alps_proto_rawDescOnce sync.Once
	file_alps_proto_rawDescData = file_alps_proto_rawDesc
)

func file_alps_proto_rawDescGZIP() []byte {
	file_alps_proto_rawDescOnce.Do(func() {
		file_alps_proto_rawDescData = protoimpl.X.CompressGZIP(file_alps_proto_rawDescData)
	})
	return file_alps_proto_rawDescData
}

var file_alps_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_alps_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_alps_proto_goTypes = []any{
	(CredentialType)(0),                // 0: grpc.CredentialType
	(*CredentialRequest)(nil),          // 1: grpc.CredentialRequest
	(*PrincipalsReply)(nil),            // 2: grpc.PrincipalsReply
	(*PrincipalsReply_Principals)(nil), // 3: grpc.PrincipalsReply.Principals
}
var file_alps_proto_depIdxs = []int32{
	0, // 0: grpc.CredentialRequest.type:type_name -> grpc.CredentialType
	3, // 1: grpc.PrincipalsReply.principals:type_name -> grpc.PrincipalsReply.Principals
	1, // 2: grpc.AuthenticationService.GetAccountPrincipals:input_type -> grpc.CredentialRequest
	2, // 3: grpc.AuthenticationService.GetAccountPrincipals:output_type -> grpc.PrincipalsReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_alps_proto_init() }
func file_alps_proto_init() {
	if File_alps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alps_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CredentialRequest); i {
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
		file_alps_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PrincipalsReply); i {
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
		file_alps_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PrincipalsReply_Principals); i {
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
			RawDescriptor: file_alps_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alps_proto_goTypes,
		DependencyIndexes: file_alps_proto_depIdxs,
		EnumInfos:         file_alps_proto_enumTypes,
		MessageInfos:      file_alps_proto_msgTypes,
	}.Build()
	File_alps_proto = out.File
	file_alps_proto_rawDesc = nil
	file_alps_proto_goTypes = nil
	file_alps_proto_depIdxs = nil
}
