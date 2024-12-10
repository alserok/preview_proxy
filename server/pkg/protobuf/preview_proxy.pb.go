// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: server/pkg/protobuf/preview_proxy.proto

package proto

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

type DownloadThumbnailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoUrls []string `protobuf:"bytes,1,rep,name=video_urls,json=videoUrls,proto3" json:"video_urls,omitempty"`
	Async     bool     `protobuf:"varint,2,opt,name=async,proto3" json:"async,omitempty"`
}

func (x *DownloadThumbnailReq) Reset() {
	*x = DownloadThumbnailReq{}
	mi := &file_server_pkg_protobuf_preview_proxy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadThumbnailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadThumbnailReq) ProtoMessage() {}

func (x *DownloadThumbnailReq) ProtoReflect() protoreflect.Message {
	mi := &file_server_pkg_protobuf_preview_proxy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadThumbnailReq.ProtoReflect.Descriptor instead.
func (*DownloadThumbnailReq) Descriptor() ([]byte, []int) {
	return file_server_pkg_protobuf_preview_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadThumbnailReq) GetVideoUrls() []string {
	if x != nil {
		return x.VideoUrls
	}
	return nil
}

func (x *DownloadThumbnailReq) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

type DownloadThumbnailRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failed uint32            `protobuf:"varint,1,opt,name=failed,proto3" json:"failed,omitempty"`
	Total  uint32            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Videos []*VideoThumbnail `protobuf:"bytes,3,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *DownloadThumbnailRes) Reset() {
	*x = DownloadThumbnailRes{}
	mi := &file_server_pkg_protobuf_preview_proxy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadThumbnailRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadThumbnailRes) ProtoMessage() {}

func (x *DownloadThumbnailRes) ProtoReflect() protoreflect.Message {
	mi := &file_server_pkg_protobuf_preview_proxy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadThumbnailRes.ProtoReflect.Descriptor instead.
func (*DownloadThumbnailRes) Descriptor() ([]byte, []int) {
	return file_server_pkg_protobuf_preview_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadThumbnailRes) GetFailed() uint32 {
	if x != nil {
		return x.Failed
	}
	return 0
}

func (x *DownloadThumbnailRes) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DownloadThumbnailRes) GetVideos() []*VideoThumbnail {
	if x != nil {
		return x.Videos
	}
	return nil
}

type VideoThumbnail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoUrl     string `protobuf:"bytes,1,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	ThumbnailUrl string `protobuf:"bytes,2,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
}

func (x *VideoThumbnail) Reset() {
	*x = VideoThumbnail{}
	mi := &file_server_pkg_protobuf_preview_proxy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoThumbnail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoThumbnail) ProtoMessage() {}

func (x *VideoThumbnail) ProtoReflect() protoreflect.Message {
	mi := &file_server_pkg_protobuf_preview_proxy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoThumbnail.ProtoReflect.Descriptor instead.
func (*VideoThumbnail) Descriptor() ([]byte, []int) {
	return file_server_pkg_protobuf_preview_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *VideoThumbnail) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *VideoThumbnail) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

var File_server_pkg_protobuf_preview_proxy_proto protoreflect.FileDescriptor

var file_server_pkg_protobuf_preview_proxy_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x14, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x22, 0x6d, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x06,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x73, 0x22, 0x52, 0x0a, 0x0e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x32, 0x52, 0x0a, 0x0c, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x42, 0x0a, 0x12, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x15, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x42, 0x3c, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x73, 0x65,
	0x72, 0x6f, 0x6b, 0x2f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_server_pkg_protobuf_preview_proxy_proto_rawDescOnce sync.Once
	file_server_pkg_protobuf_preview_proxy_proto_rawDescData = file_server_pkg_protobuf_preview_proxy_proto_rawDesc
)

func file_server_pkg_protobuf_preview_proxy_proto_rawDescGZIP() []byte {
	file_server_pkg_protobuf_preview_proxy_proto_rawDescOnce.Do(func() {
		file_server_pkg_protobuf_preview_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_pkg_protobuf_preview_proxy_proto_rawDescData)
	})
	return file_server_pkg_protobuf_preview_proxy_proto_rawDescData
}

var file_server_pkg_protobuf_preview_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_server_pkg_protobuf_preview_proxy_proto_goTypes = []any{
	(*DownloadThumbnailReq)(nil), // 0: DownloadThumbnailReq
	(*DownloadThumbnailRes)(nil), // 1: DownloadThumbnailRes
	(*VideoThumbnail)(nil),       // 2: VideoThumbnail
}
var file_server_pkg_protobuf_preview_proxy_proto_depIdxs = []int32{
	2, // 0: DownloadThumbnailRes.videos:type_name -> VideoThumbnail
	0, // 1: PreviewProxy.DownloadThumbnails:input_type -> DownloadThumbnailReq
	1, // 2: PreviewProxy.DownloadThumbnails:output_type -> DownloadThumbnailRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_server_pkg_protobuf_preview_proxy_proto_init() }
func file_server_pkg_protobuf_preview_proxy_proto_init() {
	if File_server_pkg_protobuf_preview_proxy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_pkg_protobuf_preview_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_pkg_protobuf_preview_proxy_proto_goTypes,
		DependencyIndexes: file_server_pkg_protobuf_preview_proxy_proto_depIdxs,
		MessageInfos:      file_server_pkg_protobuf_preview_proxy_proto_msgTypes,
	}.Build()
	File_server_pkg_protobuf_preview_proxy_proto = out.File
	file_server_pkg_protobuf_preview_proxy_proto_rawDesc = nil
	file_server_pkg_protobuf_preview_proxy_proto_goTypes = nil
	file_server_pkg_protobuf_preview_proxy_proto_depIdxs = nil
}
