// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/meeting/space.proto

package proto

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

type Space struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lng       float64 `protobuf:"fixed64,3,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat       float64 `protobuf:"fixed64,4,opt,name=lat,proto3" json:"lat,omitempty"`
	Status    string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Space) Reset() {
	*x = Space{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meeting_space_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Space) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Space) ProtoMessage() {}

func (x *Space) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meeting_space_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Space.ProtoReflect.Descriptor instead.
func (*Space) Descriptor() ([]byte, []int) {
	return file_proto_meeting_space_proto_rawDescGZIP(), []int{0}
}

func (x *Space) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Space) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Space) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *Space) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Space) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Space) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ReqGetAllSpaces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortBy string `protobuf:"bytes,1,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	Order  string `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *ReqGetAllSpaces) Reset() {
	*x = ReqGetAllSpaces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meeting_space_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetAllSpaces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetAllSpaces) ProtoMessage() {}

func (x *ReqGetAllSpaces) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meeting_space_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetAllSpaces.ProtoReflect.Descriptor instead.
func (*ReqGetAllSpaces) Descriptor() ([]byte, []int) {
	return file_proto_meeting_space_proto_rawDescGZIP(), []int{1}
}

func (x *ReqGetAllSpaces) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ReqGetAllSpaces) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type ReqGetSpaces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//???????????????????????????????????????
	//@inject_tag: bson:"page"
	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	//?????????????????????????????????10
	//@inject_tag: bson:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// ????????????
	//@inject_tag: bson:"sortBy"
	SortBy string `protobuf:"bytes,3,opt,name=sortBy,proto3" json:"sortBy,omitempty"`
	// ?????? desc?????? asc??????
	//@inject_tag: bson:"order"
	Order string `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
	// Types that are assignable to OneStatus:
	//	*ReqGetSpaces_Status
	OneStatus isReqGetSpaces_OneStatus `protobuf_oneof:"one_status"`
}

func (x *ReqGetSpaces) Reset() {
	*x = ReqGetSpaces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meeting_space_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetSpaces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetSpaces) ProtoMessage() {}

func (x *ReqGetSpaces) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meeting_space_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetSpaces.ProtoReflect.Descriptor instead.
func (*ReqGetSpaces) Descriptor() ([]byte, []int) {
	return file_proto_meeting_space_proto_rawDescGZIP(), []int{2}
}

func (x *ReqGetSpaces) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ReqGetSpaces) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ReqGetSpaces) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ReqGetSpaces) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (m *ReqGetSpaces) GetOneStatus() isReqGetSpaces_OneStatus {
	if m != nil {
		return m.OneStatus
	}
	return nil
}

func (x *ReqGetSpaces) GetStatus() string {
	if x, ok := x.GetOneStatus().(*ReqGetSpaces_Status); ok {
		return x.Status
	}
	return ""
}

type isReqGetSpaces_OneStatus interface {
	isReqGetSpaces_OneStatus()
}

type ReqGetSpaces_Status struct {
	Status string `protobuf:"bytes,5,opt,name=status,proto3,oneof"`
}

func (*ReqGetSpaces_Status) isReqGetSpaces_OneStatus() {}

type ReqGetSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReqGetSpace) Reset() {
	*x = ReqGetSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meeting_space_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetSpace) ProtoMessage() {}

func (x *ReqGetSpace) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meeting_space_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetSpace.ProtoReflect.Descriptor instead.
func (*ReqGetSpace) Descriptor() ([]byte, []int) {
	return file_proto_meeting_space_proto_rawDescGZIP(), []int{3}
}

func (x *ReqGetSpace) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReqCreateSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lng  float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat  float64 `protobuf:"fixed64,3,opt,name=lat,proto3" json:"lat,omitempty"`
	// Types that are assignable to OneStatus:
	//	*ReqCreateSpace_Status
	OneStatus isReqCreateSpace_OneStatus `protobuf_oneof:"one_status"`
}

func (x *ReqCreateSpace) Reset() {
	*x = ReqCreateSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meeting_space_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqCreateSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqCreateSpace) ProtoMessage() {}

func (x *ReqCreateSpace) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meeting_space_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqCreateSpace.ProtoReflect.Descriptor instead.
func (*ReqCreateSpace) Descriptor() ([]byte, []int) {
	return file_proto_meeting_space_proto_rawDescGZIP(), []int{4}
}

func (x *ReqCreateSpace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqCreateSpace) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *ReqCreateSpace) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (m *ReqCreateSpace) GetOneStatus() isReqCreateSpace_OneStatus {
	if m != nil {
		return m.OneStatus
	}
	return nil
}

func (x *ReqCreateSpace) GetStatus() string {
	if x, ok := x.GetOneStatus().(*ReqCreateSpace_Status); ok {
		return x.Status
	}
	return ""
}

type isReqCreateSpace_OneStatus interface {
	isReqCreateSpace_OneStatus()
}

type ReqCreateSpace_Status struct {
	Status string `protobuf:"bytes,4,opt,name=status,proto3,oneof"`
}

func (*ReqCreateSpace_Status) isReqCreateSpace_OneStatus() {}

type ReqUpdateSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lng  float64 `protobuf:"fixed64,3,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat  float64 `protobuf:"fixed64,4,opt,name=lat,proto3" json:"lat,omitempty"`
	// Types that are assignable to OneStatus:
	//	*ReqUpdateSpace_Status
	OneStatus isReqUpdateSpace_OneStatus `protobuf_oneof:"one_status"`
}

func (x *ReqUpdateSpace) Reset() {
	*x = ReqUpdateSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meeting_space_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUpdateSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUpdateSpace) ProtoMessage() {}

func (x *ReqUpdateSpace) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meeting_space_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUpdateSpace.ProtoReflect.Descriptor instead.
func (*ReqUpdateSpace) Descriptor() ([]byte, []int) {
	return file_proto_meeting_space_proto_rawDescGZIP(), []int{5}
}

func (x *ReqUpdateSpace) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReqUpdateSpace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqUpdateSpace) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *ReqUpdateSpace) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (m *ReqUpdateSpace) GetOneStatus() isReqUpdateSpace_OneStatus {
	if m != nil {
		return m.OneStatus
	}
	return nil
}

func (x *ReqUpdateSpace) GetStatus() string {
	if x, ok := x.GetOneStatus().(*ReqUpdateSpace_Status); ok {
		return x.Status
	}
	return ""
}

type isReqUpdateSpace_OneStatus interface {
	isReqUpdateSpace_OneStatus()
}

type ReqUpdateSpace_Status struct {
	Status string `protobuf:"bytes,5,opt,name=status,proto3,oneof"`
}

func (*ReqUpdateSpace_Status) isReqUpdateSpace_OneStatus() {}

type ReqDelSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReqDelSpace) Reset() {
	*x = ReqDelSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meeting_space_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqDelSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqDelSpace) ProtoMessage() {}

func (x *ReqDelSpace) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meeting_space_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqDelSpace.ProtoReflect.Descriptor instead.
func (*ReqDelSpace) Descriptor() ([]byte, []int) {
	return file_proto_meeting_space_proto_rawDescGZIP(), []int{6}
}

func (x *ReqDelSpace) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReqUpdateSpaceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to OneStatus:
	//	*ReqUpdateSpaceStatus_Status
	OneStatus isReqUpdateSpaceStatus_OneStatus `protobuf_oneof:"one_status"`
}

func (x *ReqUpdateSpaceStatus) Reset() {
	*x = ReqUpdateSpaceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_meeting_space_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUpdateSpaceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUpdateSpaceStatus) ProtoMessage() {}

func (x *ReqUpdateSpaceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_meeting_space_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUpdateSpaceStatus.ProtoReflect.Descriptor instead.
func (*ReqUpdateSpaceStatus) Descriptor() ([]byte, []int) {
	return file_proto_meeting_space_proto_rawDescGZIP(), []int{7}
}

func (x *ReqUpdateSpaceStatus) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (m *ReqUpdateSpaceStatus) GetOneStatus() isReqUpdateSpaceStatus_OneStatus {
	if m != nil {
		return m.OneStatus
	}
	return nil
}

func (x *ReqUpdateSpaceStatus) GetStatus() string {
	if x, ok := x.GetOneStatus().(*ReqUpdateSpaceStatus_Status); ok {
		return x.Status
	}
	return ""
}

type isReqUpdateSpaceStatus_OneStatus interface {
	isReqUpdateSpaceStatus_OneStatus()
}

type ReqUpdateSpaceStatus_Status struct {
	Status string `protobuf:"bytes,2,opt,name=status,proto3,oneof"`
}

func (*ReqUpdateSpaceStatus_Status) isReqUpdateSpaceStatus_OneStatus() {}

var File_proto_meeting_space_proto protoreflect.FileDescriptor

var file_proto_meeting_space_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x05, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3f, 0x0a, 0x0f,
	0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x94, 0x01,
	0x0a, 0x0c, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x1d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x70, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x18, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x6f, 0x6e,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x44,
	0x65, 0x6c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x6f, 0x6e, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xbe, 0x03, 0x0a, 0x0c, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x1a, 0x12, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14,
	0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71,
	0x44, 0x65, 0x6c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65,
	0x71, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_meeting_space_proto_rawDescOnce sync.Once
	file_proto_meeting_space_proto_rawDescData = file_proto_meeting_space_proto_rawDesc
)

func file_proto_meeting_space_proto_rawDescGZIP() []byte {
	file_proto_meeting_space_proto_rawDescOnce.Do(func() {
		file_proto_meeting_space_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_meeting_space_proto_rawDescData)
	})
	return file_proto_meeting_space_proto_rawDescData
}

var file_proto_meeting_space_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_meeting_space_proto_goTypes = []interface{}{
	(*Space)(nil),                // 0: meeting.Space
	(*ReqGetAllSpaces)(nil),      // 1: meeting.ReqGetAllSpaces
	(*ReqGetSpaces)(nil),         // 2: meeting.ReqGetSpaces
	(*ReqGetSpace)(nil),          // 3: meeting.ReqGetSpace
	(*ReqCreateSpace)(nil),       // 4: meeting.ReqCreateSpace
	(*ReqUpdateSpace)(nil),       // 5: meeting.ReqUpdateSpace
	(*ReqDelSpace)(nil),          // 6: meeting.ReqDelSpace
	(*ReqUpdateSpaceStatus)(nil), // 7: meeting.ReqUpdateSpaceStatus
	(*Response)(nil),             // 8: response.Response
}
var file_proto_meeting_space_proto_depIdxs = []int32{
	1, // 0: meeting.SpaceService.GetAllSpaces:input_type -> meeting.ReqGetAllSpaces
	2, // 1: meeting.SpaceService.GetSpaces:input_type -> meeting.ReqGetSpaces
	3, // 2: meeting.SpaceService.GetSpace:input_type -> meeting.ReqGetSpace
	4, // 3: meeting.SpaceService.CreateSpace:input_type -> meeting.ReqCreateSpace
	5, // 4: meeting.SpaceService.UpdateSpace:input_type -> meeting.ReqUpdateSpace
	6, // 5: meeting.SpaceService.DelSpace:input_type -> meeting.ReqDelSpace
	7, // 6: meeting.SpaceService.UpdateSpaceStatus:input_type -> meeting.ReqUpdateSpaceStatus
	8, // 7: meeting.SpaceService.GetAllSpaces:output_type -> response.Response
	8, // 8: meeting.SpaceService.GetSpaces:output_type -> response.Response
	8, // 9: meeting.SpaceService.GetSpace:output_type -> response.Response
	8, // 10: meeting.SpaceService.CreateSpace:output_type -> response.Response
	8, // 11: meeting.SpaceService.UpdateSpace:output_type -> response.Response
	8, // 12: meeting.SpaceService.DelSpace:output_type -> response.Response
	8, // 13: meeting.SpaceService.UpdateSpaceStatus:output_type -> response.Response
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_meeting_space_proto_init() }
func file_proto_meeting_space_proto_init() {
	if File_proto_meeting_space_proto != nil {
		return
	}
	file_proto_response_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_meeting_space_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Space); i {
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
		file_proto_meeting_space_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetAllSpaces); i {
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
		file_proto_meeting_space_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetSpaces); i {
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
		file_proto_meeting_space_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetSpace); i {
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
		file_proto_meeting_space_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqCreateSpace); i {
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
		file_proto_meeting_space_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUpdateSpace); i {
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
		file_proto_meeting_space_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqDelSpace); i {
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
		file_proto_meeting_space_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUpdateSpaceStatus); i {
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
	file_proto_meeting_space_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ReqGetSpaces_Status)(nil),
	}
	file_proto_meeting_space_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ReqCreateSpace_Status)(nil),
	}
	file_proto_meeting_space_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ReqUpdateSpace_Status)(nil),
	}
	file_proto_meeting_space_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*ReqUpdateSpaceStatus_Status)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_meeting_space_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_meeting_space_proto_goTypes,
		DependencyIndexes: file_proto_meeting_space_proto_depIdxs,
		MessageInfos:      file_proto_meeting_space_proto_msgTypes,
	}.Build()
	File_proto_meeting_space_proto = out.File
	file_proto_meeting_space_proto_rawDesc = nil
	file_proto_meeting_space_proto_goTypes = nil
	file_proto_meeting_space_proto_depIdxs = nil
}
