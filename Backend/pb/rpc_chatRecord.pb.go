// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: rpc_chatRecord.proto

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

// Create Chat Table
type CreateChatTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *CreateChatTableRequest) Reset() {
	*x = CreateChatTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatTableRequest) ProtoMessage() {}

func (x *CreateChatTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatTableRequest.ProtoReflect.Descriptor instead.
func (*CreateChatTableRequest) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{0}
}

func (x *CreateChatTableRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

// Create
type CreateChatRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    int32  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TargetID  int32  `protobuf:"varint,2,opt,name=TargetID,proto3" json:"TargetID,omitempty"`
	RoleType  string `protobuf:"bytes,3,opt,name=RoleType,proto3" json:"RoleType,omitempty"`
	MediaType string `protobuf:"bytes,4,opt,name=MediaType,proto3" json:"MediaType,omitempty"`
	Media     string `protobuf:"bytes,6,opt,name=Media,proto3" json:"Media,omitempty"`
}

func (x *CreateChatRecordRequest) Reset() {
	*x = CreateChatRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRecordRequest) ProtoMessage() {}

func (x *CreateChatRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRecordRequest.ProtoReflect.Descriptor instead.
func (*CreateChatRecordRequest) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChatRecordRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateChatRecordRequest) GetTargetID() int32 {
	if x != nil {
		return x.TargetID
	}
	return 0
}

func (x *CreateChatRecordRequest) GetRoleType() string {
	if x != nil {
		return x.RoleType
	}
	return ""
}

func (x *CreateChatRecordRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *CreateChatRecordRequest) GetMedia() string {
	if x != nil {
		return x.Media
	}
	return ""
}

type CreateChatRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatRecord *ChatRecord `protobuf:"bytes,1,opt,name=ChatRecord,proto3" json:"ChatRecord,omitempty"`
}

func (x *CreateChatRecordResponse) Reset() {
	*x = CreateChatRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRecordResponse) ProtoMessage() {}

func (x *CreateChatRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRecordResponse.ProtoReflect.Descriptor instead.
func (*CreateChatRecordResponse) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{2}
}

func (x *CreateChatRecordResponse) GetChatRecord() *ChatRecord {
	if x != nil {
		return x.ChatRecord
	}
	return nil
}

// GetRecord
type GetChatRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int32 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TargetID int32 `protobuf:"varint,2,opt,name=TargetID,proto3" json:"TargetID,omitempty"`
}

func (x *GetChatRecordRequest) Reset() {
	*x = GetChatRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatRecordRequest) ProtoMessage() {}

func (x *GetChatRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatRecordRequest.ProtoReflect.Descriptor instead.
func (*GetChatRecordRequest) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{3}
}

func (x *GetChatRecordRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetChatRecordRequest) GetTargetID() int32 {
	if x != nil {
		return x.TargetID
	}
	return 0
}

type GetChatRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatRecordNoID []*ChatRecordNoID `protobuf:"bytes,1,rep,name=ChatRecordNoID,proto3" json:"ChatRecordNoID,omitempty"`
}

func (x *GetChatRecordResponse) Reset() {
	*x = GetChatRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatRecordResponse) ProtoMessage() {}

func (x *GetChatRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatRecordResponse.ProtoReflect.Descriptor instead.
func (*GetChatRecordResponse) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{4}
}

func (x *GetChatRecordResponse) GetChatRecordNoID() []*ChatRecordNoID {
	if x != nil {
		return x.ChatRecordNoID
	}
	return nil
}

// GetTargetID
type GetTargetIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetTargetIDRequest) Reset() {
	*x = GetTargetIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTargetIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTargetIDRequest) ProtoMessage() {}

func (x *GetTargetIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTargetIDRequest.ProtoReflect.Descriptor instead.
func (*GetTargetIDRequest) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{5}
}

func (x *GetTargetIDRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GetTargetIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetID []int32 `protobuf:"varint,1,rep,packed,name=TargetID,proto3" json:"TargetID,omitempty"`
}

func (x *GetTargetIDResponse) Reset() {
	*x = GetTargetIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTargetIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTargetIDResponse) ProtoMessage() {}

func (x *GetTargetIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTargetIDResponse.ProtoReflect.Descriptor instead.
func (*GetTargetIDResponse) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{6}
}

func (x *GetTargetIDResponse) GetTargetID() []int32 {
	if x != nil {
		return x.TargetID
	}
	return nil
}

// GetLastMsg
type GetLastMsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int32 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TargetID int32 `protobuf:"varint,2,opt,name=TargetID,proto3" json:"TargetID,omitempty"`
}

func (x *GetLastMsgRequest) Reset() {
	*x = GetLastMsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastMsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastMsgRequest) ProtoMessage() {}

func (x *GetLastMsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastMsgRequest.ProtoReflect.Descriptor instead.
func (*GetLastMsgRequest) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{7}
}

func (x *GetLastMsgRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetLastMsgRequest) GetTargetID() int32 {
	if x != nil {
		return x.TargetID
	}
	return 0
}

type GetLastMsgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaType string `protobuf:"bytes,1,opt,name=MediaType,proto3" json:"MediaType,omitempty"`
	Media     string `protobuf:"bytes,2,opt,name=Media,proto3" json:"Media,omitempty"`
	Isread    bool   `protobuf:"varint,3,opt,name=Isread,proto3" json:"Isread,omitempty"`
}

func (x *GetLastMsgResponse) Reset() {
	*x = GetLastMsgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastMsgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastMsgResponse) ProtoMessage() {}

func (x *GetLastMsgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastMsgResponse.ProtoReflect.Descriptor instead.
func (*GetLastMsgResponse) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{8}
}

func (x *GetLastMsgResponse) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *GetLastMsgResponse) GetMedia() string {
	if x != nil {
		return x.Media
	}
	return ""
}

func (x *GetLastMsgResponse) GetIsread() bool {
	if x != nil {
		return x.Isread
	}
	return false
}

// Update
type UpdateChatRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    int32                  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TargetID  int32                  `protobuf:"varint,2,opt,name=TargetID,proto3" json:"TargetID,omitempty"`
	RoleType  string                 `protobuf:"bytes,3,opt,name=RoleType,proto3" json:"RoleType,omitempty"`
	MediaType string                 `protobuf:"bytes,4,opt,name=MediaType,proto3" json:"MediaType,omitempty"`
	Media     string                 `protobuf:"bytes,5,opt,name=Media,proto3" json:"Media,omitempty"`
	CreateAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
}

func (x *UpdateChatRecordRequest) Reset() {
	*x = UpdateChatRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChatRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChatRecordRequest) ProtoMessage() {}

func (x *UpdateChatRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChatRecordRequest.ProtoReflect.Descriptor instead.
func (*UpdateChatRecordRequest) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateChatRecordRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UpdateChatRecordRequest) GetTargetID() int32 {
	if x != nil {
		return x.TargetID
	}
	return 0
}

func (x *UpdateChatRecordRequest) GetRoleType() string {
	if x != nil {
		return x.RoleType
	}
	return ""
}

func (x *UpdateChatRecordRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *UpdateChatRecordRequest) GetMedia() string {
	if x != nil {
		return x.Media
	}
	return ""
}

func (x *UpdateChatRecordRequest) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

type UpdateChatRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatRecord *ChatRecord `protobuf:"bytes,1,opt,name=ChatRecord,proto3" json:"ChatRecord,omitempty"`
}

func (x *UpdateChatRecordResponse) Reset() {
	*x = UpdateChatRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChatRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChatRecordResponse) ProtoMessage() {}

func (x *UpdateChatRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChatRecordResponse.ProtoReflect.Descriptor instead.
func (*UpdateChatRecordResponse) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateChatRecordResponse) GetChatRecord() *ChatRecord {
	if x != nil {
		return x.ChatRecord
	}
	return nil
}

type UpdateReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int32 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TargetID int32 `protobuf:"varint,2,opt,name=TargetID,proto3" json:"TargetID,omitempty"`
}

func (x *UpdateReadRequest) Reset() {
	*x = UpdateReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReadRequest) ProtoMessage() {}

func (x *UpdateReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReadRequest.ProtoReflect.Descriptor instead.
func (*UpdateReadRequest) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateReadRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UpdateReadRequest) GetTargetID() int32 {
	if x != nil {
		return x.TargetID
	}
	return 0
}

// Delete
type DeleteChatRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int32                  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TargetID int32                  `protobuf:"varint,2,opt,name=TargetID,proto3" json:"TargetID,omitempty"`
	CreateAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
}

func (x *DeleteChatRecordRequest) Reset() {
	*x = DeleteChatRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_chatRecord_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatRecordRequest) ProtoMessage() {}

func (x *DeleteChatRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_chatRecord_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatRecordRequest.ProtoReflect.Descriptor instead.
func (*DeleteChatRecordRequest) Descriptor() ([]byte, []int) {
	return file_rpc_chatRecord_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteChatRecordRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *DeleteChatRecordRequest) GetTargetID() int32 {
	if x != nil {
		return x.TargetID
	}
	return 0
}

func (x *DeleteChatRecordRequest) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

var File_rpc_chatRecord_proto protoreflect.FileDescriptor

var file_rpc_chatRecord_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x10, 0x63, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x49, 0x44, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x9d, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x4a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x22, 0x4a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x22, 0x53, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f,
	0x49, 0x44, 0x52, 0x0e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f,
	0x49, 0x44, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x44, 0x22, 0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x22, 0x60, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x72, 0x65, 0x61, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x72, 0x65, 0x61, 0x64, 0x22, 0xd5,
	0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x36,
	0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x4a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0a, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x22, 0x47, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x22, 0x85, 0x01, 0x0a, 0x17,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x08, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_chatRecord_proto_rawDescOnce sync.Once
	file_rpc_chatRecord_proto_rawDescData = file_rpc_chatRecord_proto_rawDesc
)

func file_rpc_chatRecord_proto_rawDescGZIP() []byte {
	file_rpc_chatRecord_proto_rawDescOnce.Do(func() {
		file_rpc_chatRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_chatRecord_proto_rawDescData)
	})
	return file_rpc_chatRecord_proto_rawDescData
}

var file_rpc_chatRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_rpc_chatRecord_proto_goTypes = []interface{}{
	(*CreateChatTableRequest)(nil),   // 0: pb.CreateChatTableRequest
	(*CreateChatRecordRequest)(nil),  // 1: pb.CreateChatRecordRequest
	(*CreateChatRecordResponse)(nil), // 2: pb.CreateChatRecordResponse
	(*GetChatRecordRequest)(nil),     // 3: pb.GetChatRecordRequest
	(*GetChatRecordResponse)(nil),    // 4: pb.GetChatRecordResponse
	(*GetTargetIDRequest)(nil),       // 5: pb.GetTargetIDRequest
	(*GetTargetIDResponse)(nil),      // 6: pb.GetTargetIDResponse
	(*GetLastMsgRequest)(nil),        // 7: pb.GetLastMsgRequest
	(*GetLastMsgResponse)(nil),       // 8: pb.GetLastMsgResponse
	(*UpdateChatRecordRequest)(nil),  // 9: pb.UpdateChatRecordRequest
	(*UpdateChatRecordResponse)(nil), // 10: pb.UpdateChatRecordResponse
	(*UpdateReadRequest)(nil),        // 11: pb.UpdateReadRequest
	(*DeleteChatRecordRequest)(nil),  // 12: pb.DeleteChatRecordRequest
	(*ChatRecord)(nil),               // 13: pb.ChatRecord
	(*ChatRecordNoID)(nil),           // 14: pb.ChatRecordNoID
	(*timestamppb.Timestamp)(nil),    // 15: google.protobuf.Timestamp
}
var file_rpc_chatRecord_proto_depIdxs = []int32{
	13, // 0: pb.CreateChatRecordResponse.ChatRecord:type_name -> pb.ChatRecord
	14, // 1: pb.GetChatRecordResponse.ChatRecordNoID:type_name -> pb.ChatRecordNoID
	15, // 2: pb.UpdateChatRecordRequest.CreateAt:type_name -> google.protobuf.Timestamp
	13, // 3: pb.UpdateChatRecordResponse.ChatRecord:type_name -> pb.ChatRecord
	15, // 4: pb.DeleteChatRecordRequest.CreateAt:type_name -> google.protobuf.Timestamp
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_rpc_chatRecord_proto_init() }
func file_rpc_chatRecord_proto_init() {
	if File_rpc_chatRecord_proto != nil {
		return
	}
	file_chatRecord_proto_init()
	file_chatRecordNoID_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_chatRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatTableRequest); i {
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
		file_rpc_chatRecord_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatRecordRequest); i {
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
		file_rpc_chatRecord_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatRecordResponse); i {
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
		file_rpc_chatRecord_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatRecordRequest); i {
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
		file_rpc_chatRecord_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatRecordResponse); i {
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
		file_rpc_chatRecord_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTargetIDRequest); i {
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
		file_rpc_chatRecord_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTargetIDResponse); i {
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
		file_rpc_chatRecord_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastMsgRequest); i {
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
		file_rpc_chatRecord_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastMsgResponse); i {
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
		file_rpc_chatRecord_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChatRecordRequest); i {
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
		file_rpc_chatRecord_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChatRecordResponse); i {
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
		file_rpc_chatRecord_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReadRequest); i {
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
		file_rpc_chatRecord_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatRecordRequest); i {
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
			RawDescriptor: file_rpc_chatRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_chatRecord_proto_goTypes,
		DependencyIndexes: file_rpc_chatRecord_proto_depIdxs,
		MessageInfos:      file_rpc_chatRecord_proto_msgTypes,
	}.Build()
	File_rpc_chatRecord_proto = out.File
	file_rpc_chatRecord_proto_rawDesc = nil
	file_rpc_chatRecord_proto_goTypes = nil
	file_rpc_chatRecord_proto_depIdxs = nil
}
