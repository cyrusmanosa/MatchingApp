// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: rpc_ResetPasswordEmail.proto

package pb

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

type ResetPwEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ResetPwEmailRequest) Reset() {
	*x = ResetPwEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_ResetPasswordEmail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPwEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPwEmailRequest) ProtoMessage() {}

func (x *ResetPwEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ResetPasswordEmail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPwEmailRequest.ProtoReflect.Descriptor instead.
func (*ResetPwEmailRequest) Descriptor() ([]byte, []int) {
	return file_rpc_ResetPasswordEmail_proto_rawDescGZIP(), []int{0}
}

func (x *ResetPwEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ResetPwEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link   string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResetPwEmailResponse) Reset() {
	*x = ResetPwEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_ResetPasswordEmail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPwEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPwEmailResponse) ProtoMessage() {}

func (x *ResetPwEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ResetPasswordEmail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPwEmailResponse.ProtoReflect.Descriptor instead.
func (*ResetPwEmailResponse) Descriptor() ([]byte, []int) {
	return file_rpc_ResetPasswordEmail_proto_rawDescGZIP(), []int{1}
}

func (x *ResetPwEmailResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *ResetPwEmailResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ResetPwEmailConfirmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ResetPwEmailConfirmRequest) Reset() {
	*x = ResetPwEmailConfirmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_ResetPasswordEmail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPwEmailConfirmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPwEmailConfirmRequest) ProtoMessage() {}

func (x *ResetPwEmailConfirmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_ResetPasswordEmail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPwEmailConfirmRequest.ProtoReflect.Descriptor instead.
func (*ResetPwEmailConfirmRequest) Descriptor() ([]byte, []int) {
	return file_rpc_ResetPasswordEmail_proto_rawDescGZIP(), []int{2}
}

func (x *ResetPwEmailConfirmRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResetPwEmailConfirmRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_rpc_ResetPasswordEmail_proto protoreflect.FileDescriptor

var file_rpc_ResetPasswordEmail_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x5f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x2b, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x77, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x42, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x77, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x4e, 0x0a, 0x1a, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x77, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x42, 0x0c, 0x5a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_ResetPasswordEmail_proto_rawDescOnce sync.Once
	file_rpc_ResetPasswordEmail_proto_rawDescData = file_rpc_ResetPasswordEmail_proto_rawDesc
)

func file_rpc_ResetPasswordEmail_proto_rawDescGZIP() []byte {
	file_rpc_ResetPasswordEmail_proto_rawDescOnce.Do(func() {
		file_rpc_ResetPasswordEmail_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_ResetPasswordEmail_proto_rawDescData)
	})
	return file_rpc_ResetPasswordEmail_proto_rawDescData
}

var file_rpc_ResetPasswordEmail_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_ResetPasswordEmail_proto_goTypes = []interface{}{
	(*ResetPwEmailRequest)(nil),        // 0: pb.ResetPwEmailRequest
	(*ResetPwEmailResponse)(nil),       // 1: pb.ResetPwEmailResponse
	(*ResetPwEmailConfirmRequest)(nil), // 2: pb.ResetPwEmailConfirmRequest
}
var file_rpc_ResetPasswordEmail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_ResetPasswordEmail_proto_init() }
func file_rpc_ResetPasswordEmail_proto_init() {
	if File_rpc_ResetPasswordEmail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_ResetPasswordEmail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPwEmailRequest); i {
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
		file_rpc_ResetPasswordEmail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPwEmailResponse); i {
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
		file_rpc_ResetPasswordEmail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPwEmailConfirmRequest); i {
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
			RawDescriptor: file_rpc_ResetPasswordEmail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_ResetPasswordEmail_proto_goTypes,
		DependencyIndexes: file_rpc_ResetPasswordEmail_proto_depIdxs,
		MessageInfos:      file_rpc_ResetPasswordEmail_proto_msgTypes,
	}.Build()
	File_rpc_ResetPasswordEmail_proto = out.File
	file_rpc_ResetPasswordEmail_proto_rawDesc = nil
	file_rpc_ResetPasswordEmail_proto_goTypes = nil
	file_rpc_ResetPasswordEmail_proto_depIdxs = nil
}