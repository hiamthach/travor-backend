// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: package_service.proto

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

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DesId        uint64  `protobuf:"varint,2,opt,name=des_id,json=desId,proto3" json:"des_id,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Details      string  `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
	Price        int32   `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	ImgUrl       string  `protobuf:"bytes,6,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
	Duration     string  `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
	NumberPeople int32   `protobuf:"varint,8,opt,name=number_people,json=numberPeople,proto3" json:"number_people,omitempty"`
	Types        []*Type `protobuf:"bytes,9,rep,name=types,proto3" json:"types,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Package) GetDesId() uint64 {
	if x != nil {
		return x.DesId
	}
	return 0
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *Package) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Package) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *Package) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Package) GetNumberPeople() int32 {
	if x != nil {
		return x.NumberPeople
	}
	return 0
}

func (x *Package) GetTypes() []*Type {
	if x != nil {
		return x.Types
	}
	return nil
}

type PackageStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PackageStats) Reset() {
	*x = PackageStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageStats) ProtoMessage() {}

func (x *PackageStats) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageStats.ProtoReflect.Descriptor instead.
func (*PackageStats) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{1}
}

func (x *PackageStats) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PackageStats) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPkgStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPkgStatsRequest) Reset() {
	*x = GetPkgStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPkgStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPkgStatsRequest) ProtoMessage() {}

func (x *GetPkgStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPkgStatsRequest.ProtoReflect.Descriptor instead.
func (*GetPkgStatsRequest) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{2}
}

type GetPkgStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packages []*PackageStats `protobuf:"bytes,1,rep,name=packages,proto3" json:"packages,omitempty"`
}

func (x *GetPkgStatsResponse) Reset() {
	*x = GetPkgStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPkgStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPkgStatsResponse) ProtoMessage() {}

func (x *GetPkgStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPkgStatsResponse.ProtoReflect.Descriptor instead.
func (*GetPkgStatsResponse) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPkgStatsResponse) GetPackages() []*PackageStats {
	if x != nil {
		return x.Packages
	}
	return nil
}

type PackageID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PackageID) Reset() {
	*x = PackageID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageID) ProtoMessage() {}

func (x *PackageID) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageID.ProtoReflect.Descriptor instead.
func (*PackageID) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{4}
}

func (x *PackageID) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IsValidPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *IsValidPackageResponse) Reset() {
	*x = IsValidPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidPackageResponse) ProtoMessage() {}

func (x *IsValidPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidPackageResponse.ProtoReflect.Descriptor instead.
func (*IsValidPackageResponse) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{5}
}

func (x *IsValidPackageResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type GetPackagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keyword  string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *GetPackagesRequest) Reset() {
	*x = GetPackagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackagesRequest) ProtoMessage() {}

func (x *GetPackagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackagesRequest.ProtoReflect.Descriptor instead.
func (*GetPackagesRequest) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetPackagesRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPackagesRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetPackagesRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type GetPackagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packages   []*Package     `protobuf:"bytes,1,rep,name=packages,proto3" json:"packages,omitempty"`
	Pagination *PaginationRes `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetPackagesResponse) Reset() {
	*x = GetPackagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPackagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPackagesResponse) ProtoMessage() {}

func (x *GetPackagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPackagesResponse.ProtoReflect.Descriptor instead.
func (*GetPackagesResponse) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetPackagesResponse) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

func (x *GetPackagesResponse) GetPagination() *PaginationRes {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type CreatePackageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DesId        uint64   `protobuf:"varint,1,opt,name=des_id,json=desId,proto3" json:"des_id,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Details      string   `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	Price        int32    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	ImgUrl       string   `protobuf:"bytes,5,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
	Duration     string   `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	NumberPeople int32    `protobuf:"varint,7,opt,name=number_people,json=numberPeople,proto3" json:"number_people,omitempty"`
	Types        []uint64 `protobuf:"varint,8,rep,packed,name=types,proto3" json:"types,omitempty"`
}

func (x *CreatePackageReq) Reset() {
	*x = CreatePackageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePackageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePackageReq) ProtoMessage() {}

func (x *CreatePackageReq) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePackageReq.ProtoReflect.Descriptor instead.
func (*CreatePackageReq) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreatePackageReq) GetDesId() uint64 {
	if x != nil {
		return x.DesId
	}
	return 0
}

func (x *CreatePackageReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePackageReq) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *CreatePackageReq) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreatePackageReq) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *CreatePackageReq) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *CreatePackageReq) GetNumberPeople() int32 {
	if x != nil {
		return x.NumberPeople
	}
	return 0
}

func (x *CreatePackageReq) GetTypes() []uint64 {
	if x != nil {
		return x.Types
	}
	return nil
}

type UpdatePackageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DesId        uint64   `protobuf:"varint,2,opt,name=des_id,json=desId,proto3" json:"des_id,omitempty"`
	Name         string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Details      string   `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
	Price        int32    `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	ImgUrl       string   `protobuf:"bytes,6,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
	Duration     string   `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
	NumberPeople int32    `protobuf:"varint,8,opt,name=number_people,json=numberPeople,proto3" json:"number_people,omitempty"`
	Types        []uint64 `protobuf:"varint,9,rep,packed,name=types,proto3" json:"types,omitempty"`
}

func (x *UpdatePackageReq) Reset() {
	*x = UpdatePackageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePackageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePackageReq) ProtoMessage() {}

func (x *UpdatePackageReq) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePackageReq.ProtoReflect.Descriptor instead.
func (*UpdatePackageReq) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{9}
}

func (x *UpdatePackageReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePackageReq) GetDesId() uint64 {
	if x != nil {
		return x.DesId
	}
	return 0
}

func (x *UpdatePackageReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePackageReq) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *UpdatePackageReq) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdatePackageReq) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *UpdatePackageReq) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *UpdatePackageReq) GetNumberPeople() int32 {
	if x != nil {
		return x.NumberPeople
	}
	return 0
}

func (x *UpdatePackageReq) GetTypes() []uint64 {
	if x != nil {
		return x.Types
	}
	return nil
}

type DeletePackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeletePackageResponse) Reset() {
	*x = DeletePackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePackageResponse) ProtoMessage() {}

func (x *DeletePackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_package_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePackageResponse.ProtoReflect.Descriptor instead.
func (*DeletePackageResponse) Descriptor() ([]byte, []int) {
	return file_package_service_proto_rawDescGZIP(), []int{10}
}

func (x *DeletePackageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_package_service_proto protoreflect.FileDescriptor

var file_package_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xee, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64,
	0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x64, 0x65, 0x73,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x12, 0x1e, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x32, 0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6b, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x50, 0x6b, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x1b, 0x0a, 0x09, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x16,
	0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x71, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x31, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xdd, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x64, 0x65, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x22, 0xed, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x64, 0x65, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0xa4, 0x04, 0x0a, 0x0e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x44, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6b, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6b, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0b, 0x12, 0x09, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e,
	0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x48,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x1a, 0x0e, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x51, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x44, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x76, 0x6f, 0x72, 0x2d,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_package_service_proto_rawDescOnce sync.Once
	file_package_service_proto_rawDescData = file_package_service_proto_rawDesc
)

func file_package_service_proto_rawDescGZIP() []byte {
	file_package_service_proto_rawDescOnce.Do(func() {
		file_package_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_service_proto_rawDescData)
	})
	return file_package_service_proto_rawDescData
}

var file_package_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_package_service_proto_goTypes = []interface{}{
	(*Package)(nil),                // 0: pb.Package
	(*PackageStats)(nil),           // 1: pb.PackageStats
	(*GetPkgStatsRequest)(nil),     // 2: pb.GetPkgStatsRequest
	(*GetPkgStatsResponse)(nil),    // 3: pb.GetPkgStatsResponse
	(*PackageID)(nil),              // 4: pb.PackageID
	(*IsValidPackageResponse)(nil), // 5: pb.IsValidPackageResponse
	(*GetPackagesRequest)(nil),     // 6: pb.GetPackagesRequest
	(*GetPackagesResponse)(nil),    // 7: pb.GetPackagesResponse
	(*CreatePackageReq)(nil),       // 8: pb.CreatePackageReq
	(*UpdatePackageReq)(nil),       // 9: pb.UpdatePackageReq
	(*DeletePackageResponse)(nil),  // 10: pb.DeletePackageResponse
	(*Type)(nil),                   // 11: pb.Type
	(*PaginationRes)(nil),          // 12: pb.PaginationRes
}
var file_package_service_proto_depIdxs = []int32{
	11, // 0: pb.Package.types:type_name -> pb.Type
	1,  // 1: pb.GetPkgStatsResponse.packages:type_name -> pb.PackageStats
	0,  // 2: pb.GetPackagesResponse.packages:type_name -> pb.Package
	12, // 3: pb.GetPackagesResponse.pagination:type_name -> pb.PaginationRes
	4,  // 4: pb.PackageService.IsValidPackage:input_type -> pb.PackageID
	2,  // 5: pb.PackageService.GetAll:input_type -> pb.GetPkgStatsRequest
	6,  // 6: pb.PackageService.GetPackages:input_type -> pb.GetPackagesRequest
	4,  // 7: pb.PackageService.GetPackage:input_type -> pb.PackageID
	8,  // 8: pb.PackageService.CreatePackage:input_type -> pb.CreatePackageReq
	9,  // 9: pb.PackageService.UpdatePackage:input_type -> pb.UpdatePackageReq
	4,  // 10: pb.PackageService.DeletePackage:input_type -> pb.PackageID
	5,  // 11: pb.PackageService.IsValidPackage:output_type -> pb.IsValidPackageResponse
	3,  // 12: pb.PackageService.GetAll:output_type -> pb.GetPkgStatsResponse
	7,  // 13: pb.PackageService.GetPackages:output_type -> pb.GetPackagesResponse
	0,  // 14: pb.PackageService.GetPackage:output_type -> pb.Package
	0,  // 15: pb.PackageService.CreatePackage:output_type -> pb.Package
	0,  // 16: pb.PackageService.UpdatePackage:output_type -> pb.Package
	10, // 17: pb.PackageService.DeletePackage:output_type -> pb.DeletePackageResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_package_service_proto_init() }
func file_package_service_proto_init() {
	if File_package_service_proto != nil {
		return
	}
	file_pagination_proto_init()
	file_type_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_package_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_package_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageStats); i {
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
		file_package_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPkgStatsRequest); i {
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
		file_package_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPkgStatsResponse); i {
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
		file_package_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageID); i {
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
		file_package_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidPackageResponse); i {
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
		file_package_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackagesRequest); i {
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
		file_package_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPackagesResponse); i {
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
		file_package_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePackageReq); i {
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
		file_package_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePackageReq); i {
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
		file_package_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePackageResponse); i {
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
			RawDescriptor: file_package_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_package_service_proto_goTypes,
		DependencyIndexes: file_package_service_proto_depIdxs,
		MessageInfos:      file_package_service_proto_msgTypes,
	}.Build()
	File_package_service_proto = out.File
	file_package_service_proto_rawDesc = nil
	file_package_service_proto_goTypes = nil
	file_package_service_proto_depIdxs = nil
}
