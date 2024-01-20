// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: fix.proto

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

type Fix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName     string                 `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName      string                 `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Birth         string                 `protobuf:"bytes,3,opt,name=Birth,proto3" json:"Birth,omitempty"`
	Country       string                 `protobuf:"bytes,4,opt,name=Country,proto3" json:"Country,omitempty"`
	Gender        string                 `protobuf:"bytes,5,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Blood         string                 `protobuf:"bytes,6,opt,name=Blood,proto3" json:"Blood,omitempty"`
	Age           int32                  `protobuf:"varint,7,opt,name=age,proto3" json:"age,omitempty"`
	Constellation string                 `protobuf:"bytes,8,opt,name=constellation,proto3" json:"constellation,omitempty"`
	Certification bool                   `protobuf:"varint,9,opt,name=certification,proto3" json:"certification,omitempty"`
	CreateAt      *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
}

func (x *Fix) Reset() {
	*x = Fix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fix_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fix) ProtoMessage() {}

func (x *Fix) ProtoReflect() protoreflect.Message {
	mi := &file_fix_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fix.ProtoReflect.Descriptor instead.
func (*Fix) Descriptor() ([]byte, []int) {
	return file_fix_proto_rawDescGZIP(), []int{0}
}

func (x *Fix) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Fix) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Fix) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

func (x *Fix) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Fix) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Fix) GetBlood() string {
	if x != nil {
		return x.Blood
	}
	return ""
}

func (x *Fix) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Fix) GetConstellation() string {
	if x != nil {
		return x.Constellation
	}
	return ""
}

func (x *Fix) GetCertification() bool {
	if x != nil {
		return x.Certification
	}
	return false
}

func (x *Fix) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

var File_fix_proto protoreflect.FileDescriptor

var file_fix_proto_rawDesc = []byte{
	0x0a, 0x09, 0x66, 0x69, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb3, 0x02, 0x0a, 0x03, 0x46, 0x69, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x6c,
	0x6f, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x42, 0x6c, 0x6f, 0x6f, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fix_proto_rawDescOnce sync.Once
	file_fix_proto_rawDescData = file_fix_proto_rawDesc
)

func file_fix_proto_rawDescGZIP() []byte {
	file_fix_proto_rawDescOnce.Do(func() {
		file_fix_proto_rawDescData = protoimpl.X.CompressGZIP(file_fix_proto_rawDescData)
	})
	return file_fix_proto_rawDescData
}

var file_fix_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fix_proto_goTypes = []interface{}{
	(*Fix)(nil),                   // 0: pb.Fix
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_fix_proto_depIdxs = []int32{
	1, // 0: pb.Fix.CreateAt:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fix_proto_init() }
func file_fix_proto_init() {
	if File_fix_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fix_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fix); i {
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
			RawDescriptor: file_fix_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fix_proto_goTypes,
		DependencyIndexes: file_fix_proto_depIdxs,
		MessageInfos:      file_fix_proto_msgTypes,
	}.Build()
	File_fix_proto = out.File
	file_fix_proto_rawDesc = nil
	file_fix_proto_goTypes = nil
	file_fix_proto_depIdxs = nil
}
