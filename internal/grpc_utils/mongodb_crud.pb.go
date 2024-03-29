// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: mongodb_crud.proto

package grpc_utils

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

// Query anime configurations
type QueryAnimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActiveType int32    `protobuf:"varint,1,opt,name=activeType,proto3" json:"activeType,omitempty"`
	Names      []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"` // if ['*'], return the whole list
}

func (x *QueryAnimeRequest) Reset() {
	*x = QueryAnimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongodb_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAnimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAnimeRequest) ProtoMessage() {}

func (x *QueryAnimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongodb_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAnimeRequest.ProtoReflect.Descriptor instead.
func (*QueryAnimeRequest) Descriptor() ([]byte, []int) {
	return file_mongodb_crud_proto_rawDescGZIP(), []int{0}
}

func (x *QueryAnimeRequest) GetActiveType() int32 {
	if x != nil {
		return x.ActiveType
	}
	return 0
}

func (x *QueryAnimeRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type QueryAnimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names  []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	RssUrl []string `protobuf:"bytes,3,rep,name=rssUrl,proto3" json:"rssUrl,omitempty"`
}

func (x *QueryAnimeResponse) Reset() {
	*x = QueryAnimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongodb_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAnimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAnimeResponse) ProtoMessage() {}

func (x *QueryAnimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mongodb_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAnimeResponse.ProtoReflect.Descriptor instead.
func (*QueryAnimeResponse) Descriptor() ([]byte, []int) {
	return file_mongodb_crud_proto_rawDescGZIP(), []int{1}
}

func (x *QueryAnimeResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *QueryAnimeResponse) GetRssUrl() []string {
	if x != nil {
		return x.RssUrl
	}
	return nil
}

// Add/Modify anime configurations
type UpdateAnimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names        []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	RssUrls      []string `protobuf:"bytes,3,rep,name=rssUrls,proto3" json:"rssUrls,omitempty"`
	RuleVersions []string `protobuf:"bytes,4,rep,name=ruleVersions,proto3" json:"ruleVersions,omitempty"`
	RuleRegexs   []string `protobuf:"bytes,5,rep,name=ruleRegexs,proto3" json:"ruleRegexs,omitempty"`
	IsActives    []bool   `protobuf:"varint,6,rep,packed,name=isActives,proto3" json:"isActives,omitempty"`
}

func (x *UpdateAnimeRequest) Reset() {
	*x = UpdateAnimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongodb_crud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnimeRequest) ProtoMessage() {}

func (x *UpdateAnimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongodb_crud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnimeRequest.ProtoReflect.Descriptor instead.
func (*UpdateAnimeRequest) Descriptor() ([]byte, []int) {
	return file_mongodb_crud_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAnimeRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *UpdateAnimeRequest) GetRssUrls() []string {
	if x != nil {
		return x.RssUrls
	}
	return nil
}

func (x *UpdateAnimeRequest) GetRuleVersions() []string {
	if x != nil {
		return x.RuleVersions
	}
	return nil
}

func (x *UpdateAnimeRequest) GetRuleRegexs() []string {
	if x != nil {
		return x.RuleRegexs
	}
	return nil
}

func (x *UpdateAnimeRequest) GetIsActives() []bool {
	if x != nil {
		return x.IsActives
	}
	return nil
}

type UpdateAnimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessCount int64    `protobuf:"varint,1,opt,name=successCount,proto3" json:"successCount,omitempty"`
	FailedList   []string `protobuf:"bytes,2,rep,name=failedList,proto3" json:"failedList,omitempty"`
}

func (x *UpdateAnimeResponse) Reset() {
	*x = UpdateAnimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongodb_crud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnimeResponse) ProtoMessage() {}

func (x *UpdateAnimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mongodb_crud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnimeResponse.ProtoReflect.Descriptor instead.
func (*UpdateAnimeResponse) Descriptor() ([]byte, []int) {
	return file_mongodb_crud_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAnimeResponse) GetSuccessCount() int64 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

func (x *UpdateAnimeResponse) GetFailedList() []string {
	if x != nil {
		return x.FailedList
	}
	return nil
}

// Delete anime configurations
type DelAnimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelAll bool     `protobuf:"varint,1,opt,name=delAll,proto3" json:"delAll,omitempty"`
	Names  []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *DelAnimeRequest) Reset() {
	*x = DelAnimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongodb_crud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelAnimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelAnimeRequest) ProtoMessage() {}

func (x *DelAnimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mongodb_crud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelAnimeRequest.ProtoReflect.Descriptor instead.
func (*DelAnimeRequest) Descriptor() ([]byte, []int) {
	return file_mongodb_crud_proto_rawDescGZIP(), []int{4}
}

func (x *DelAnimeRequest) GetDelAll() bool {
	if x != nil {
		return x.DelAll
	}
	return false
}

func (x *DelAnimeRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type DelAnimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessCount int64 `protobuf:"varint,1,opt,name=successCount,proto3" json:"successCount,omitempty"`
}

func (x *DelAnimeResponse) Reset() {
	*x = DelAnimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mongodb_crud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelAnimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelAnimeResponse) ProtoMessage() {}

func (x *DelAnimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mongodb_crud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelAnimeResponse.ProtoReflect.Descriptor instead.
func (*DelAnimeResponse) Descriptor() ([]byte, []int) {
	return file_mongodb_crud_proto_rawDescGZIP(), []int{5}
}

func (x *DelAnimeResponse) GetSuccessCount() int64 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

var File_mongodb_crud_proto protoreflect.FileDescriptor

var file_mongodb_crud_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x5f, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x22, 0x49, 0x0a, 0x11, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x22, 0xa6, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x75, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x75, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x65, 0x78,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x67,
	0x65, 0x78, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x22, 0x59, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x41, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x64, 0x65, 0x6c, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x36, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb3, 0x02, 0x0a, 0x11, 0x4d, 0x69, 0x6b, 0x61, 0x6e, 0x61,
	0x6e, 0x69, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x43, 0x72, 0x75, 0x64, 0x12, 0x5f, 0x0a, 0x0a, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x2e, 0x6d, 0x69, 0x6b, 0x61,
	0x6e, 0x61, 0x6e, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x2e, 0x6d, 0x69,
	0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x59, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x2e, 0x6d,
	0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x74, 0x69,
	0x6c, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x75,
	0x74, 0x69, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mongodb_crud_proto_rawDescOnce sync.Once
	file_mongodb_crud_proto_rawDescData = file_mongodb_crud_proto_rawDesc
)

func file_mongodb_crud_proto_rawDescGZIP() []byte {
	file_mongodb_crud_proto_rawDescOnce.Do(func() {
		file_mongodb_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_mongodb_crud_proto_rawDescData)
	})
	return file_mongodb_crud_proto_rawDescData
}

var file_mongodb_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mongodb_crud_proto_goTypes = []interface{}{
	(*QueryAnimeRequest)(nil),   // 0: mikanani_grpc_utils.QueryAnimeRequest
	(*QueryAnimeResponse)(nil),  // 1: mikanani_grpc_utils.QueryAnimeResponse
	(*UpdateAnimeRequest)(nil),  // 2: mikanani_grpc_utils.UpdateAnimeRequest
	(*UpdateAnimeResponse)(nil), // 3: mikanani_grpc_utils.UpdateAnimeResponse
	(*DelAnimeRequest)(nil),     // 4: mikanani_grpc_utils.DelAnimeRequest
	(*DelAnimeResponse)(nil),    // 5: mikanani_grpc_utils.DelAnimeResponse
}
var file_mongodb_crud_proto_depIdxs = []int32{
	0, // 0: mikanani_grpc_utils.MikananiMongoCrud.QueryAnime:input_type -> mikanani_grpc_utils.QueryAnimeRequest
	2, // 1: mikanani_grpc_utils.MikananiMongoCrud.UpdateAnime:input_type -> mikanani_grpc_utils.UpdateAnimeRequest
	4, // 2: mikanani_grpc_utils.MikananiMongoCrud.DelAnime:input_type -> mikanani_grpc_utils.DelAnimeRequest
	1, // 3: mikanani_grpc_utils.MikananiMongoCrud.QueryAnime:output_type -> mikanani_grpc_utils.QueryAnimeResponse
	3, // 4: mikanani_grpc_utils.MikananiMongoCrud.UpdateAnime:output_type -> mikanani_grpc_utils.UpdateAnimeResponse
	5, // 5: mikanani_grpc_utils.MikananiMongoCrud.DelAnime:output_type -> mikanani_grpc_utils.DelAnimeResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mongodb_crud_proto_init() }
func file_mongodb_crud_proto_init() {
	if File_mongodb_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mongodb_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAnimeRequest); i {
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
		file_mongodb_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAnimeResponse); i {
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
		file_mongodb_crud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnimeRequest); i {
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
		file_mongodb_crud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnimeResponse); i {
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
		file_mongodb_crud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelAnimeRequest); i {
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
		file_mongodb_crud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelAnimeResponse); i {
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
			RawDescriptor: file_mongodb_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mongodb_crud_proto_goTypes,
		DependencyIndexes: file_mongodb_crud_proto_depIdxs,
		MessageInfos:      file_mongodb_crud_proto_msgTypes,
	}.Build()
	File_mongodb_crud_proto = out.File
	file_mongodb_crud_proto_rawDesc = nil
	file_mongodb_crud_proto_goTypes = nil
	file_mongodb_crud_proto_depIdxs = nil
}
