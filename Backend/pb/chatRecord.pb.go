// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: chatRecord.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChatRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetID  int32                  `protobuf:"varint,1,opt,name=TargetID,proto3" json:"TargetID,omitempty"`
	RoleType  string                 `protobuf:"bytes,2,opt,name=RoleType,proto3" json:"RoleType,omitempty"`
	MediaType string                 `protobuf:"bytes,3,opt,name=MediaType,proto3" json:"MediaType,omitempty"`
	Media     string                 `protobuf:"bytes,4,opt,name=Media,proto3" json:"Media,omitempty"`
	IsRead    bool                   `protobuf:"varint,5,opt,name=IsRead,proto3" json:"IsRead,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *ChatRecord) Reset() {
	*x = ChatRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRecord) ProtoMessage() {}

func (x *ChatRecord) ProtoReflect() protoreflect.Message {
	mi := &file_chatRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRecord.ProtoReflect.Descriptor instead.
func (*ChatRecord) Descriptor() ([]byte, []int) {
	return file_chatRecord_proto_rawDescGZIP(), []int{0}
}

func (x *ChatRecord) GetTargetID() int32 {
	if x != nil {
		return x.TargetID
	}
	return 0
}

func (x *ChatRecord) GetRoleType() string {
	if x != nil {
		return x.RoleType
	}
	return ""
}

func (x *ChatRecord) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *ChatRecord) GetMedia() string {
	if x != nil {
		return x.Media
	}
	return ""
}

func (x *ChatRecord) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *ChatRecord) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_chatRecord_proto protoreflect.FileDescriptor

var file_chatRecord_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatRecord_proto_rawDescOnce sync.Once
	file_chatRecord_proto_rawDescData = file_chatRecord_proto_rawDesc
)

func file_chatRecord_proto_rawDescGZIP() []byte {
	file_chatRecord_proto_rawDescOnce.Do(func() {
		file_chatRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatRecord_proto_rawDescData)
	})
	return file_chatRecord_proto_rawDescData
}

var file_chatRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chatRecord_proto_goTypes = []interface{}{
	(*ChatRecord)(nil),            // 0: pb.ChatRecord
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_chatRecord_proto_depIdxs = []int32{
	1, // 0: pb.ChatRecord.CreatedAt:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chatRecord_proto_init() }
func file_chatRecord_proto_init() {
	if File_chatRecord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRecord); i {
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
			RawDescriptor: file_chatRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chatRecord_proto_goTypes,
		DependencyIndexes: file_chatRecord_proto_depIdxs,
		MessageInfos:      file_chatRecord_proto_msgTypes,
	}.Build()
	File_chatRecord_proto = out.File
	file_chatRecord_proto_rawDesc = nil
	file_chatRecord_proto_goTypes = nil
	file_chatRecord_proto_depIdxs = nil
}
