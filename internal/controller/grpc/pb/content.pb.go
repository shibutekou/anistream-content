// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: content.proto

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

type TitleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	TitleOrig   string   `protobuf:"bytes,2,opt,name=titleOrig,proto3" json:"titleOrig,omitempty"`
	Year        int32    `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	KinopoiskID string   `protobuf:"bytes,4,opt,name=kinopoiskID,proto3" json:"kinopoiskID,omitempty"`
	ShikimoriID string   `protobuf:"bytes,5,opt,name=shikimoriID,proto3" json:"shikimoriID,omitempty"`
	ImdbID      string   `protobuf:"bytes,6,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Screenshots []string `protobuf:"bytes,7,rep,name=screenshots,proto3" json:"screenshots,omitempty"`
}

func (x *TitleInfo) Reset() {
	*x = TitleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TitleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TitleInfo) ProtoMessage() {}

func (x *TitleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TitleInfo.ProtoReflect.Descriptor instead.
func (*TitleInfo) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{0}
}

func (x *TitleInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TitleInfo) GetTitleOrig() string {
	if x != nil {
		return x.TitleOrig
	}
	return ""
}

func (x *TitleInfo) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *TitleInfo) GetKinopoiskID() string {
	if x != nil {
		return x.KinopoiskID
	}
	return ""
}

func (x *TitleInfo) GetShikimoriID() string {
	if x != nil {
		return x.ShikimoriID
	}
	return ""
}

func (x *TitleInfo) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *TitleInfo) GetScreenshots() []string {
	if x != nil {
		return x.Screenshots
	}
	return nil
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{1}
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type TitleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Svc string `protobuf:"bytes,1,opt,name=svc,proto3" json:"svc,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *TitleFilter) Reset() {
	*x = TitleFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TitleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TitleFilter) ProtoMessage() {}

func (x *TitleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TitleFilter.ProtoReflect.Descriptor instead.
func (*TitleFilter) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{2}
}

func (x *TitleFilter) GetSvc() string {
	if x != nil {
		return x.Svc
	}
	return ""
}

func (x *TitleFilter) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type GetTitleInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *TitleFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetTitleInfoRequest) Reset() {
	*x = GetTitleInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTitleInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTitleInfoRequest) ProtoMessage() {}

func (x *GetTitleInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTitleInfoRequest.ProtoReflect.Descriptor instead.
func (*GetTitleInfoRequest) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{3}
}

func (x *GetTitleInfoRequest) GetFilter() *TitleFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetTitleInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TitleInfo []*TitleInfo `protobuf:"bytes,1,rep,name=titleInfo,proto3" json:"titleInfo,omitempty"`
}

func (x *GetTitleInfoReply) Reset() {
	*x = GetTitleInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTitleInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTitleInfoReply) ProtoMessage() {}

func (x *GetTitleInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTitleInfoReply.ProtoReflect.Descriptor instead.
func (*GetTitleInfoReply) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{4}
}

func (x *GetTitleInfoReply) GetTitleInfo() []*TitleInfo {
	if x != nil {
		return x.TitleInfo
	}
	return nil
}

type GetLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *TitleFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetLinkRequest) Reset() {
	*x = GetLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkRequest) ProtoMessage() {}

func (x *GetLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkRequest.ProtoReflect.Descriptor instead.
func (*GetLinkRequest) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{5}
}

func (x *GetLinkRequest) GetFilter() *TitleFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetLinkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link *Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *GetLinkReply) Reset() {
	*x = GetLinkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkReply) ProtoMessage() {}

func (x *GetLinkReply) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkReply.ProtoReflect.Descriptor instead.
func (*GetLinkReply) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{6}
}

func (x *GetLinkReply) GetLink() *Link {
	if x != nil {
		return x.Link
	}
	return nil
}

var File_content_proto protoreflect.FileDescriptor

var file_content_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd1, 0x01, 0x0a, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4f, 0x72, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4f, 0x72, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x6b, 0x69, 0x6e, 0x6f, 0x70, 0x6f, 0x69,
	0x73, 0x6b, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b, 0x69, 0x6e, 0x6f,
	0x70, 0x6f, 0x69, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x68, 0x69, 0x6b, 0x69,
	0x6d, 0x6f, 0x72, 0x69, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x68,
	0x69, 0x6b, 0x69, 0x6d, 0x6f, 0x72, 0x69, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64,
	0x62, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68,
	0x6f, 0x74, 0x73, 0x22, 0x18, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x31, 0x0a,
	0x0b, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x76, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x76, 0x63, 0x12, 0x10,
	0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x22, 0x3b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x3d, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x36, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x32,
	0x75, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_content_proto_rawDescOnce sync.Once
	file_content_proto_rawDescData = file_content_proto_rawDesc
)

func file_content_proto_rawDescGZIP() []byte {
	file_content_proto_rawDescOnce.Do(func() {
		file_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_proto_rawDescData)
	})
	return file_content_proto_rawDescData
}

var file_content_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_content_proto_goTypes = []interface{}{
	(*TitleInfo)(nil),           // 0: TitleInfo
	(*Link)(nil),                // 1: Link
	(*TitleFilter)(nil),         // 2: TitleFilter
	(*GetTitleInfoRequest)(nil), // 3: GetTitleInfoRequest
	(*GetTitleInfoReply)(nil),   // 4: GetTitleInfoReply
	(*GetLinkRequest)(nil),      // 5: GetLinkRequest
	(*GetLinkReply)(nil),        // 6: GetLinkReply
}
var file_content_proto_depIdxs = []int32{
	2, // 0: GetTitleInfoRequest.filter:type_name -> TitleFilter
	0, // 1: GetTitleInfoReply.titleInfo:type_name -> TitleInfo
	2, // 2: GetLinkRequest.filter:type_name -> TitleFilter
	1, // 3: GetLinkReply.link:type_name -> Link
	3, // 4: ContentService.GetTitleInfo:input_type -> GetTitleInfoRequest
	5, // 5: ContentService.GetLink:input_type -> GetLinkRequest
	4, // 6: ContentService.GetTitleInfo:output_type -> GetTitleInfoReply
	6, // 7: ContentService.GetLink:output_type -> GetLinkReply
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_content_proto_init() }
func file_content_proto_init() {
	if File_content_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TitleInfo); i {
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
		file_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
		file_content_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TitleFilter); i {
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
		file_content_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTitleInfoRequest); i {
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
		file_content_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTitleInfoReply); i {
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
		file_content_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkRequest); i {
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
		file_content_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkReply); i {
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
			RawDescriptor: file_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_content_proto_goTypes,
		DependencyIndexes: file_content_proto_depIdxs,
		MessageInfos:      file_content_proto_msgTypes,
	}.Build()
	File_content_proto = out.File
	file_content_proto_rawDesc = nil
	file_content_proto_goTypes = nil
	file_content_proto_depIdxs = nil
}
