// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: search.proto

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

type SH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultID []int32 `protobuf:"varint,1,rep,packed,name=resultID,proto3" json:"resultID,omitempty"`
	Rank     string  `protobuf:"bytes,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Len      int32   `protobuf:"varint,3,opt,name=len,proto3" json:"len,omitempty"`
}

func (x *SH) Reset() {
	*x = SH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SH) ProtoMessage() {}

func (x *SH) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SH.ProtoReflect.Descriptor instead.
func (*SH) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *SH) GetResultID() []int32 {
	if x != nil {
		return x.ResultID
	}
	return nil
}

func (x *SH) GetRank() string {
	if x != nil {
		return x.Rank
	}
	return ""
}

func (x *SH) GetLen() int32 {
	if x != nil {
		return x.Len
	}
	return 0
}

type SL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultID []int32 `protobuf:"varint,1,rep,packed,name=resultID,proto3" json:"resultID,omitempty"`
	Rank     string  `protobuf:"bytes,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Len      int32   `protobuf:"varint,3,opt,name=len,proto3" json:"len,omitempty"`
}

func (x *SL) Reset() {
	*x = SL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SL) ProtoMessage() {}

func (x *SL) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SL.ProtoReflect.Descriptor instead.
func (*SL) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *SL) GetResultID() []int32 {
	if x != nil {
		return x.ResultID
	}
	return nil
}

func (x *SL) GetRank() string {
	if x != nil {
		return x.Rank
	}
	return ""
}

func (x *SL) GetLen() int32 {
	if x != nil {
		return x.Len
	}
	return 0
}

type SA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultID []int32 `protobuf:"varint,1,rep,packed,name=resultID,proto3" json:"resultID,omitempty"`
	Rank     string  `protobuf:"bytes,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Len      int32   `protobuf:"varint,3,opt,name=len,proto3" json:"len,omitempty"`
}

func (x *SA) Reset() {
	*x = SA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SA) ProtoMessage() {}

func (x *SA) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SA.ProtoReflect.Descriptor instead.
func (*SA) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2}
}

func (x *SA) GetResultID() []int32 {
	if x != nil {
		return x.ResultID
	}
	return nil
}

func (x *SA) GetRank() string {
	if x != nil {
		return x.Rank
	}
	return ""
}

func (x *SA) GetLen() int32 {
	if x != nil {
		return x.Len
	}
	return 0
}

var File_search_proto protoreflect.FileDescriptor

var file_search_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x46, 0x0a, 0x02, 0x73, 0x48, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6c, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x02, 0x73, 0x4c,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6c,
	0x65, 0x6e, 0x22, 0x46, 0x0a, 0x02, 0x73, 0x41, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6c, 0x65, 0x6e, 0x42, 0x0c, 0x5a, 0x0a, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData = file_search_proto_rawDesc
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_proto_rawDescData)
	})
	return file_search_proto_rawDescData
}

var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_search_proto_goTypes = []interface{}{
	(*SH)(nil), // 0: pb.sH
	(*SL)(nil), // 1: pb.sL
	(*SA)(nil), // 2: pb.sA
}
var file_search_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SH); i {
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
		file_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SL); i {
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
		file_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SA); i {
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
			RawDescriptor: file_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		MessageInfos:      file_search_proto_msgTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_rawDesc = nil
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}