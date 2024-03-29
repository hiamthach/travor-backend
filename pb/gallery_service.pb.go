// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: gallery_service.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url   string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	DesId uint64 `protobuf:"varint,3,opt,name=desId,proto3" json:"desId,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gallery_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_gallery_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_gallery_service_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Image) GetDesId() uint64 {
	if x != nil {
		return x.DesId
	}
	return 0
}

type AddImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	DesId uint64 `protobuf:"varint,2,opt,name=desId,proto3" json:"desId,omitempty"`
}

func (x *AddImageRequest) Reset() {
	*x = AddImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gallery_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddImageRequest) ProtoMessage() {}

func (x *AddImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gallery_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddImageRequest.ProtoReflect.Descriptor instead.
func (*AddImageRequest) Descriptor() ([]byte, []int) {
	return file_gallery_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddImageRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddImageRequest) GetDesId() uint64 {
	if x != nil {
		return x.DesId
	}
	return 0
}

type AddImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *AddImageResponse) Reset() {
	*x = AddImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gallery_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddImageResponse) ProtoMessage() {}

func (x *AddImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gallery_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddImageResponse.ProtoReflect.Descriptor instead.
func (*AddImageResponse) Descriptor() ([]byte, []int) {
	return file_gallery_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddImageResponse) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type AddListImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls  []string `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	DesId uint64   `protobuf:"varint,2,opt,name=desId,proto3" json:"desId,omitempty"`
}

func (x *AddListImagesRequest) Reset() {
	*x = AddListImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gallery_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddListImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddListImagesRequest) ProtoMessage() {}

func (x *AddListImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gallery_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddListImagesRequest.ProtoReflect.Descriptor instead.
func (*AddListImagesRequest) Descriptor() ([]byte, []int) {
	return file_gallery_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddListImagesRequest) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *AddListImagesRequest) GetDesId() uint64 {
	if x != nil {
		return x.DesId
	}
	return 0
}

type AddListImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *AddListImagesResponse) Reset() {
	*x = AddListImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gallery_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddListImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddListImagesResponse) ProtoMessage() {}

func (x *AddListImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gallery_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddListImagesResponse.ProtoReflect.Descriptor instead.
func (*AddListImagesResponse) Descriptor() ([]byte, []int) {
	return file_gallery_service_proto_rawDescGZIP(), []int{4}
}

func (x *AddListImagesResponse) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type GetImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DesId uint64 `protobuf:"varint,1,opt,name=desId,proto3" json:"desId,omitempty"`
}

func (x *GetImagesRequest) Reset() {
	*x = GetImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gallery_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesRequest) ProtoMessage() {}

func (x *GetImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gallery_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesRequest.ProtoReflect.Descriptor instead.
func (*GetImagesRequest) Descriptor() ([]byte, []int) {
	return file_gallery_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetImagesRequest) GetDesId() uint64 {
	if x != nil {
		return x.DesId
	}
	return 0
}

type GetImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *GetImagesResponse) Reset() {
	*x = GetImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gallery_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesResponse) ProtoMessage() {}

func (x *GetImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gallery_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesResponse.ProtoReflect.Descriptor instead.
func (*GetImagesResponse) Descriptor() ([]byte, []int) {
	return file_gallery_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetImagesResponse) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type DeleteImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteImageRequest) Reset() {
	*x = DeleteImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gallery_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImageRequest) ProtoMessage() {}

func (x *DeleteImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gallery_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImageRequest.ProtoReflect.Descriptor instead.
func (*DeleteImageRequest) Descriptor() ([]byte, []int) {
	return file_gallery_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteImageRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteImageResponse) Reset() {
	*x = DeleteImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gallery_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImageResponse) ProtoMessage() {}

func (x *DeleteImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gallery_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImageResponse.ProtoReflect.Descriptor instead.
func (*DeleteImageResponse) Descriptor() ([]byte, []int) {
	return file_gallery_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteImageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_gallery_service_proto protoreflect.FileDescriptor

var file_gallery_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x05, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x64, 0x65, 0x73, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x64, 0x65, 0x73, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x14, 0x41, 0x64,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x73, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x64, 0x65, 0x73, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x15,
	0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x64, 0x65, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x64, 0x65, 0x73,
	0x49, 0x64, 0x22, 0x36, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0xef, 0x02, 0x0a, 0x0e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x60, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x73, 0x49, 0x64, 0x7d, 0x12, 0x57, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x2a, 0x0f, 0x2f, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x72, 0x61, 0x76, 0x6f, 0x72, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gallery_service_proto_rawDescOnce sync.Once
	file_gallery_service_proto_rawDescData = file_gallery_service_proto_rawDesc
)

func file_gallery_service_proto_rawDescGZIP() []byte {
	file_gallery_service_proto_rawDescOnce.Do(func() {
		file_gallery_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_gallery_service_proto_rawDescData)
	})
	return file_gallery_service_proto_rawDescData
}

var file_gallery_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gallery_service_proto_goTypes = []interface{}{
	(*Image)(nil),                 // 0: pb.Image
	(*AddImageRequest)(nil),       // 1: pb.AddImageRequest
	(*AddImageResponse)(nil),      // 2: pb.AddImageResponse
	(*AddListImagesRequest)(nil),  // 3: pb.AddListImagesRequest
	(*AddListImagesResponse)(nil), // 4: pb.AddListImagesResponse
	(*GetImagesRequest)(nil),      // 5: pb.GetImagesRequest
	(*GetImagesResponse)(nil),     // 6: pb.GetImagesResponse
	(*DeleteImageRequest)(nil),    // 7: pb.DeleteImageRequest
	(*DeleteImageResponse)(nil),   // 8: pb.DeleteImageResponse
}
var file_gallery_service_proto_depIdxs = []int32{
	0, // 0: pb.AddImageResponse.image:type_name -> pb.Image
	0, // 1: pb.AddListImagesResponse.images:type_name -> pb.Image
	0, // 2: pb.GetImagesResponse.images:type_name -> pb.Image
	1, // 3: pb.GalleryService.AddImage:input_type -> pb.AddImageRequest
	3, // 4: pb.GalleryService.AddListImages:input_type -> pb.AddListImagesRequest
	5, // 5: pb.GalleryService.GetImages:input_type -> pb.GetImagesRequest
	7, // 6: pb.GalleryService.DeleteImage:input_type -> pb.DeleteImageRequest
	2, // 7: pb.GalleryService.AddImage:output_type -> pb.AddImageResponse
	4, // 8: pb.GalleryService.AddListImages:output_type -> pb.AddListImagesResponse
	6, // 9: pb.GalleryService.GetImages:output_type -> pb.GetImagesResponse
	8, // 10: pb.GalleryService.DeleteImage:output_type -> pb.DeleteImageResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gallery_service_proto_init() }
func file_gallery_service_proto_init() {
	if File_gallery_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gallery_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_gallery_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddImageRequest); i {
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
		file_gallery_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddImageResponse); i {
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
		file_gallery_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddListImagesRequest); i {
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
		file_gallery_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddListImagesResponse); i {
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
		file_gallery_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImagesRequest); i {
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
		file_gallery_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImagesResponse); i {
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
		file_gallery_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImageRequest); i {
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
		file_gallery_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImageResponse); i {
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
			RawDescriptor: file_gallery_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gallery_service_proto_goTypes,
		DependencyIndexes: file_gallery_service_proto_depIdxs,
		MessageInfos:      file_gallery_service_proto_msgTypes,
	}.Build()
	File_gallery_service_proto = out.File
	file_gallery_service_proto_rawDesc = nil
	file_gallery_service_proto_goTypes = nil
	file_gallery_service_proto_depIdxs = nil
}
