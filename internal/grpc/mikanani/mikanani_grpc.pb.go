// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.3
// source: mikanani_grpc.proto

package mikanani

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnimeMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid            int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name           string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DownloadBitmap int64    `protobuf:"varint,3,opt,name=downloadBitmap,proto3" json:"downloadBitmap,omitempty"`
	IsActive       int32    `protobuf:"varint,4,opt,name=isActive,proto3" json:"isActive,omitempty"`
	Tags           []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *AnimeMeta) Reset() {
	*x = AnimeMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMeta) ProtoMessage() {}

func (x *AnimeMeta) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMeta.ProtoReflect.Descriptor instead.
func (*AnimeMeta) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *AnimeMeta) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AnimeMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnimeMeta) GetDownloadBitmap() int64 {
	if x != nil {
		return x.DownloadBitmap
	}
	return 0
}

func (x *AnimeMeta) GetIsActive() int32 {
	if x != nil {
		return x.IsActive
	}
	return 0
}

func (x *AnimeMeta) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type AnimeDoc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	RssUrl string `protobuf:"bytes,2,opt,name=rssUrl,proto3" json:"rssUrl,omitempty"`
	Rule   string `protobuf:"bytes,3,opt,name=rule,proto3" json:"rule,omitempty"`
	Regex  string `protobuf:"bytes,4,opt,name=regex,proto3" json:"regex,omitempty"`
}

func (x *AnimeDoc) Reset() {
	*x = AnimeDoc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeDoc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeDoc) ProtoMessage() {}

func (x *AnimeDoc) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeDoc.ProtoReflect.Descriptor instead.
func (*AnimeDoc) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *AnimeDoc) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AnimeDoc) GetRssUrl() string {
	if x != nil {
		return x.RssUrl
	}
	return ""
}

func (x *AnimeDoc) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *AnimeDoc) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

// ----- ListAnimeMeta
type ListAnimeMetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartIndex   int64 `protobuf:"varint,1,opt,name=startIndex,proto3" json:"startIndex,omitempty"` // if startIndex == endIndex == -1, return all
	EndIndex     int64 `protobuf:"varint,2,opt,name=endIndex,proto3" json:"endIndex,omitempty"`
	StatusFilter int32 `protobuf:"varint,3,opt,name=statusFilter,proto3" json:"statusFilter,omitempty"` // 0: all, 1: only-active, 2: only-inactive
}

func (x *ListAnimeMetaRequest) Reset() {
	*x = ListAnimeMetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAnimeMetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAnimeMetaRequest) ProtoMessage() {}

func (x *ListAnimeMetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAnimeMetaRequest.ProtoReflect.Descriptor instead.
func (*ListAnimeMetaRequest) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *ListAnimeMetaRequest) GetStartIndex() int64 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

func (x *ListAnimeMetaRequest) GetEndIndex() int64 {
	if x != nil {
		return x.EndIndex
	}
	return 0
}

func (x *ListAnimeMetaRequest) GetStatusFilter() int32 {
	if x != nil {
		return x.StatusFilter
	}
	return 0
}

type ListAnimeMetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemCount  int64        `protobuf:"varint,1,opt,name=itemCount,proto3" json:"itemCount,omitempty"`
	AnimeMetas []*AnimeMeta `protobuf:"bytes,2,rep,name=animeMetas,proto3" json:"animeMetas,omitempty"`
}

func (x *ListAnimeMetaResponse) Reset() {
	*x = ListAnimeMetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAnimeMetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAnimeMetaResponse) ProtoMessage() {}

func (x *ListAnimeMetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAnimeMetaResponse.ProtoReflect.Descriptor instead.
func (*ListAnimeMetaResponse) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *ListAnimeMetaResponse) GetItemCount() int64 {
	if x != nil {
		return x.ItemCount
	}
	return 0
}

func (x *ListAnimeMetaResponse) GetAnimeMetas() []*AnimeMeta {
	if x != nil {
		return x.AnimeMetas
	}
	return nil
}

// ---- GetAnimeDoc
type GetAnimeDocRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetAnimeDocRequest) Reset() {
	*x = GetAnimeDocRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnimeDocRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnimeDocRequest) ProtoMessage() {}

func (x *GetAnimeDocRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnimeDocRequest.ProtoReflect.Descriptor instead.
func (*GetAnimeDocRequest) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *GetAnimeDocRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetAnimeDocResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeDoc *AnimeDoc `protobuf:"bytes,1,opt,name=animeDoc,proto3" json:"animeDoc,omitempty"`
}

func (x *GetAnimeDocResponse) Reset() {
	*x = GetAnimeDocResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnimeDocResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnimeDocResponse) ProtoMessage() {}

func (x *GetAnimeDocResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnimeDocResponse.ProtoReflect.Descriptor instead.
func (*GetAnimeDocResponse) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *GetAnimeDocResponse) GetAnimeDoc() *AnimeDoc {
	if x != nil {
		return x.AnimeDoc
	}
	return nil
}

// ---- UpdateAnimeDoc
type UpdateAnimeDocRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateAnimeDoc *AnimeDoc `protobuf:"bytes,1,opt,name=updateAnimeDoc,proto3" json:"updateAnimeDoc,omitempty"`
}

func (x *UpdateAnimeDocRequest) Reset() {
	*x = UpdateAnimeDocRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnimeDocRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnimeDocRequest) ProtoMessage() {}

func (x *UpdateAnimeDocRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnimeDocRequest.ProtoReflect.Descriptor instead.
func (*UpdateAnimeDocRequest) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAnimeDocRequest) GetUpdateAnimeDoc() *AnimeDoc {
	if x != nil {
		return x.UpdateAnimeDoc
	}
	return nil
}

// ---- UpdateAnimeMeta
type UpdateAnimeMetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateAnimeMeta *AnimeMeta `protobuf:"bytes,1,opt,name=updateAnimeMeta,proto3" json:"updateAnimeMeta,omitempty"`
}

func (x *UpdateAnimeMetaRequest) Reset() {
	*x = UpdateAnimeMetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnimeMetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnimeMetaRequest) ProtoMessage() {}

func (x *UpdateAnimeMetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnimeMetaRequest.ProtoReflect.Descriptor instead.
func (*UpdateAnimeMetaRequest) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAnimeMetaRequest) GetUpdateAnimeMeta() *AnimeMeta {
	if x != nil {
		return x.UpdateAnimeMeta
	}
	return nil
}

// ---- InsertAnimeItem
type InsertAnimeItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InsertAnimeMeta *AnimeMeta `protobuf:"bytes,1,opt,name=insertAnimeMeta,proto3" json:"insertAnimeMeta,omitempty"`
	InsertAnimeDoc  *AnimeDoc  `protobuf:"bytes,2,opt,name=insertAnimeDoc,proto3" json:"insertAnimeDoc,omitempty"`
}

func (x *InsertAnimeItemRequest) Reset() {
	*x = InsertAnimeItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertAnimeItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertAnimeItemRequest) ProtoMessage() {}

func (x *InsertAnimeItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertAnimeItemRequest.ProtoReflect.Descriptor instead.
func (*InsertAnimeItemRequest) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *InsertAnimeItemRequest) GetInsertAnimeMeta() *AnimeMeta {
	if x != nil {
		return x.InsertAnimeMeta
	}
	return nil
}

func (x *InsertAnimeItemRequest) GetInsertAnimeDoc() *AnimeDoc {
	if x != nil {
		return x.InsertAnimeDoc
	}
	return nil
}

type InsertAnimeItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *InsertAnimeItemResponse) Reset() {
	*x = InsertAnimeItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertAnimeItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertAnimeItemResponse) ProtoMessage() {}

func (x *InsertAnimeItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertAnimeItemResponse.ProtoReflect.Descriptor instead.
func (*InsertAnimeItemResponse) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{9}
}

func (x *InsertAnimeItemResponse) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

// ---- DeleteAnimeItem
type DeleteAnimeItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DeleteAnimeItemRequest) Reset() {
	*x = DeleteAnimeItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAnimeItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAnimeItemRequest) ProtoMessage() {}

func (x *DeleteAnimeItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAnimeItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteAnimeItemRequest) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteAnimeItemRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

// ---- GetAnimeCount
type GetAnimeCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAnimeCountResponse) Reset() {
	*x = GetAnimeCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mikanani_grpc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnimeCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnimeCountResponse) ProtoMessage() {}

func (x *GetAnimeCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mikanani_grpc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnimeCountResponse.ProtoReflect.Descriptor instead.
func (*GetAnimeCountResponse) Descriptor() ([]byte, []int) {
	return file_mikanani_grpc_proto_rawDescGZIP(), []int{11}
}

func (x *GetAnimeCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_mikanani_grpc_proto protoreflect.FileDescriptor

var file_mikanani_grpc_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x09, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x69, 0x74, 0x6d, 0x61, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x5e, 0x0a, 0x08, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44,
	0x6f, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x22, 0x76, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e,
	0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x6a,
	0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x69, 0x6b, 0x61,
	0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x22, 0x26, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x61, 0x6e, 0x69,
	0x6d, 0x65, 0x44, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x69,
	0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x63, 0x52,
	0x08, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x63, 0x22, 0x53, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x44, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x69, 0x6b,
	0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x63, 0x52, 0x0e,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x63, 0x22, 0x57,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e,
	0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x93, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x69,
	0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x0f, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x3a, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x44, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x69, 0x6b, 0x61,
	0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x63, 0x52, 0x0e, 0x69,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x63, 0x22, 0x2b, 0x0a,
	0x17, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb2, 0x06, 0x0a, 0x0f, 0x4d, 0x69, 0x6b, 0x61, 0x6e, 0x61,
	0x6e, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1e, 0x2e, 0x6d, 0x69, 0x6b,
	0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x69, 0x6b,
	0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x5e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x63, 0x12, 0x1c,
	0x2e, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d,
	0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x44, 0x6f, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x64, 0x6f, 0x63, 0x2f, 0x7b, 0x75, 0x69, 0x64, 0x7d, 0x12,
	0x61, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f,
	0x63, 0x12, 0x1f, 0x2e, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x3a, 0x01, 0x2a, 0x1a, 0x0b, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x64,
	0x6f, 0x63, 0x12, 0x64, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x20, 0x2e, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x1a, 0x0c, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2d, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x6a, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x2e, 0x6d, 0x69,
	0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x1a, 0x07, 0x2f, 0x69, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x12, 0x60, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6e,
	0x69, 0x6d, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x2e, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61,
	0x6e, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x2f, 0x75, 0x69, 0x64, 0x12, 0x62, 0x0a, 0x14, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x12, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x2d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x5e, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61,
	0x6e, 0x69, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x6b, 0x61, 0x6e, 0x61, 0x6e,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mikanani_grpc_proto_rawDescOnce sync.Once
	file_mikanani_grpc_proto_rawDescData = file_mikanani_grpc_proto_rawDesc
)

func file_mikanani_grpc_proto_rawDescGZIP() []byte {
	file_mikanani_grpc_proto_rawDescOnce.Do(func() {
		file_mikanani_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mikanani_grpc_proto_rawDescData)
	})
	return file_mikanani_grpc_proto_rawDescData
}

var file_mikanani_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_mikanani_grpc_proto_goTypes = []interface{}{
	(*AnimeMeta)(nil),               // 0: mikanani.AnimeMeta
	(*AnimeDoc)(nil),                // 1: mikanani.AnimeDoc
	(*ListAnimeMetaRequest)(nil),    // 2: mikanani.ListAnimeMetaRequest
	(*ListAnimeMetaResponse)(nil),   // 3: mikanani.ListAnimeMetaResponse
	(*GetAnimeDocRequest)(nil),      // 4: mikanani.GetAnimeDocRequest
	(*GetAnimeDocResponse)(nil),     // 5: mikanani.GetAnimeDocResponse
	(*UpdateAnimeDocRequest)(nil),   // 6: mikanani.UpdateAnimeDocRequest
	(*UpdateAnimeMetaRequest)(nil),  // 7: mikanani.UpdateAnimeMetaRequest
	(*InsertAnimeItemRequest)(nil),  // 8: mikanani.InsertAnimeItemRequest
	(*InsertAnimeItemResponse)(nil), // 9: mikanani.InsertAnimeItemResponse
	(*DeleteAnimeItemRequest)(nil),  // 10: mikanani.DeleteAnimeItemRequest
	(*GetAnimeCountResponse)(nil),   // 11: mikanani.GetAnimeCountResponse
	(*emptypb.Empty)(nil),           // 12: google.protobuf.Empty
}
var file_mikanani_grpc_proto_depIdxs = []int32{
	0,  // 0: mikanani.ListAnimeMetaResponse.animeMetas:type_name -> mikanani.AnimeMeta
	1,  // 1: mikanani.GetAnimeDocResponse.animeDoc:type_name -> mikanani.AnimeDoc
	1,  // 2: mikanani.UpdateAnimeDocRequest.updateAnimeDoc:type_name -> mikanani.AnimeDoc
	0,  // 3: mikanani.UpdateAnimeMetaRequest.updateAnimeMeta:type_name -> mikanani.AnimeMeta
	0,  // 4: mikanani.InsertAnimeItemRequest.insertAnimeMeta:type_name -> mikanani.AnimeMeta
	1,  // 5: mikanani.InsertAnimeItemRequest.insertAnimeDoc:type_name -> mikanani.AnimeDoc
	2,  // 6: mikanani.MikananiService.ListAnimeMeta:input_type -> mikanani.ListAnimeMetaRequest
	4,  // 7: mikanani.MikananiService.GetAnimeDoc:input_type -> mikanani.GetAnimeDocRequest
	6,  // 8: mikanani.MikananiService.UpdateAnimeDoc:input_type -> mikanani.UpdateAnimeDocRequest
	7,  // 9: mikanani.MikananiService.UpdateAnimeMeta:input_type -> mikanani.UpdateAnimeMetaRequest
	8,  // 10: mikanani.MikananiService.InsertAnimeItem:input_type -> mikanani.InsertAnimeItemRequest
	10, // 11: mikanani.MikananiService.DeleteAnimeItem:input_type -> mikanani.DeleteAnimeItemRequest
	12, // 12: mikanani.MikananiService.DispatchDownloadTask:input_type -> google.protobuf.Empty
	12, // 13: mikanani.MikananiService.GetAnimeCount:input_type -> google.protobuf.Empty
	3,  // 14: mikanani.MikananiService.ListAnimeMeta:output_type -> mikanani.ListAnimeMetaResponse
	5,  // 15: mikanani.MikananiService.GetAnimeDoc:output_type -> mikanani.GetAnimeDocResponse
	12, // 16: mikanani.MikananiService.UpdateAnimeDoc:output_type -> google.protobuf.Empty
	12, // 17: mikanani.MikananiService.UpdateAnimeMeta:output_type -> google.protobuf.Empty
	9,  // 18: mikanani.MikananiService.InsertAnimeItem:output_type -> mikanani.InsertAnimeItemResponse
	12, // 19: mikanani.MikananiService.DeleteAnimeItem:output_type -> google.protobuf.Empty
	12, // 20: mikanani.MikananiService.DispatchDownloadTask:output_type -> google.protobuf.Empty
	11, // 21: mikanani.MikananiService.GetAnimeCount:output_type -> mikanani.GetAnimeCountResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_mikanani_grpc_proto_init() }
func file_mikanani_grpc_proto_init() {
	if File_mikanani_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mikanani_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMeta); i {
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
		file_mikanani_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeDoc); i {
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
		file_mikanani_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAnimeMetaRequest); i {
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
		file_mikanani_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAnimeMetaResponse); i {
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
		file_mikanani_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnimeDocRequest); i {
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
		file_mikanani_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnimeDocResponse); i {
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
		file_mikanani_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnimeDocRequest); i {
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
		file_mikanani_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnimeMetaRequest); i {
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
		file_mikanani_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertAnimeItemRequest); i {
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
		file_mikanani_grpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertAnimeItemResponse); i {
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
		file_mikanani_grpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAnimeItemRequest); i {
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
		file_mikanani_grpc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnimeCountResponse); i {
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
			RawDescriptor: file_mikanani_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mikanani_grpc_proto_goTypes,
		DependencyIndexes: file_mikanani_grpc_proto_depIdxs,
		MessageInfos:      file_mikanani_grpc_proto_msgTypes,
	}.Build()
	File_mikanani_grpc_proto = out.File
	file_mikanani_grpc_proto_rawDesc = nil
	file_mikanani_grpc_proto_goTypes = nil
	file_mikanani_grpc_proto_depIdxs = nil
}