// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: google/cloud/baremetalsolution/v2/lun.proto

package baremetalsolutionpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The possible states for the LUN.
type Lun_State int32

const (
	// The LUN is in an unknown state.
	Lun_STATE_UNSPECIFIED Lun_State = 0
	// The LUN is being created.
	Lun_CREATING Lun_State = 1
	// The LUN is being updated.
	Lun_UPDATING Lun_State = 2
	// The LUN is ready for use.
	Lun_READY Lun_State = 3
	// The LUN has been requested to be deleted.
	Lun_DELETING Lun_State = 4
)

// Enum value maps for Lun_State.
var (
	Lun_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "UPDATING",
		3: "READY",
		4: "DELETING",
	}
	Lun_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"UPDATING":          2,
		"READY":             3,
		"DELETING":          4,
	}
)

func (x Lun_State) Enum() *Lun_State {
	p := new(Lun_State)
	*p = x
	return p
}

func (x Lun_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Lun_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_baremetalsolution_v2_lun_proto_enumTypes[0].Descriptor()
}

func (Lun_State) Type() protoreflect.EnumType {
	return &file_google_cloud_baremetalsolution_v2_lun_proto_enumTypes[0]
}

func (x Lun_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Lun_State.Descriptor instead.
func (Lun_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_lun_proto_rawDescGZIP(), []int{0, 0}
}

// Display the operating systems present for the LUN multiprotocol type.
type Lun_MultiprotocolType int32

const (
	// Server has no OS specified.
	Lun_MULTIPROTOCOL_TYPE_UNSPECIFIED Lun_MultiprotocolType = 0
	// Server with Linux OS.
	Lun_LINUX Lun_MultiprotocolType = 1
)

// Enum value maps for Lun_MultiprotocolType.
var (
	Lun_MultiprotocolType_name = map[int32]string{
		0: "MULTIPROTOCOL_TYPE_UNSPECIFIED",
		1: "LINUX",
	}
	Lun_MultiprotocolType_value = map[string]int32{
		"MULTIPROTOCOL_TYPE_UNSPECIFIED": 0,
		"LINUX":                          1,
	}
)

func (x Lun_MultiprotocolType) Enum() *Lun_MultiprotocolType {
	p := new(Lun_MultiprotocolType)
	*p = x
	return p
}

func (x Lun_MultiprotocolType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Lun_MultiprotocolType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_baremetalsolution_v2_lun_proto_enumTypes[1].Descriptor()
}

func (Lun_MultiprotocolType) Type() protoreflect.EnumType {
	return &file_google_cloud_baremetalsolution_v2_lun_proto_enumTypes[1]
}

func (x Lun_MultiprotocolType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Lun_MultiprotocolType.Descriptor instead.
func (Lun_MultiprotocolType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_lun_proto_rawDescGZIP(), []int{0, 1}
}

// The storage types for a LUN.
type Lun_StorageType int32

const (
	// The storage type for this LUN is unknown.
	Lun_STORAGE_TYPE_UNSPECIFIED Lun_StorageType = 0
	// This storage type for this LUN is SSD.
	Lun_SSD Lun_StorageType = 1
	// This storage type for this LUN is HDD.
	Lun_HDD Lun_StorageType = 2
)

// Enum value maps for Lun_StorageType.
var (
	Lun_StorageType_name = map[int32]string{
		0: "STORAGE_TYPE_UNSPECIFIED",
		1: "SSD",
		2: "HDD",
	}
	Lun_StorageType_value = map[string]int32{
		"STORAGE_TYPE_UNSPECIFIED": 0,
		"SSD":                      1,
		"HDD":                      2,
	}
)

func (x Lun_StorageType) Enum() *Lun_StorageType {
	p := new(Lun_StorageType)
	*p = x
	return p
}

func (x Lun_StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Lun_StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_baremetalsolution_v2_lun_proto_enumTypes[2].Descriptor()
}

func (Lun_StorageType) Type() protoreflect.EnumType {
	return &file_google_cloud_baremetalsolution_v2_lun_proto_enumTypes[2]
}

func (x Lun_StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Lun_StorageType.Descriptor instead.
func (Lun_StorageType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_lun_proto_rawDescGZIP(), []int{0, 2}
}

// A storage volume logical unit number (LUN).
type Lun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The name of the LUN.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An identifier for the LUN, generated by the backend.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	// The state of this storage volume.
	State Lun_State `protobuf:"varint,2,opt,name=state,proto3,enum=google.cloud.baremetalsolution.v2.Lun_State" json:"state,omitempty"`
	// The size of this LUN, in gigabytes.
	SizeGb int64 `protobuf:"varint,3,opt,name=size_gb,json=sizeGb,proto3" json:"size_gb,omitempty"`
	// The LUN multiprotocol type ensures the characteristics of the LUN are
	// optimized for each operating system.
	MultiprotocolType Lun_MultiprotocolType `protobuf:"varint,4,opt,name=multiprotocol_type,json=multiprotocolType,proto3,enum=google.cloud.baremetalsolution.v2.Lun_MultiprotocolType" json:"multiprotocol_type,omitempty"`
	// Display the storage volume for this LUN.
	StorageVolume string `protobuf:"bytes,5,opt,name=storage_volume,json=storageVolume,proto3" json:"storage_volume,omitempty"`
	// Display if this LUN can be shared between multiple physical servers.
	Shareable bool `protobuf:"varint,6,opt,name=shareable,proto3" json:"shareable,omitempty"`
	// Display if this LUN is a boot LUN.
	BootLun bool `protobuf:"varint,7,opt,name=boot_lun,json=bootLun,proto3" json:"boot_lun,omitempty"`
	// The storage type for this LUN.
	StorageType Lun_StorageType `protobuf:"varint,8,opt,name=storage_type,json=storageType,proto3,enum=google.cloud.baremetalsolution.v2.Lun_StorageType" json:"storage_type,omitempty"`
	// The WWID for this LUN.
	Wwid string `protobuf:"bytes,9,opt,name=wwid,proto3" json:"wwid,omitempty"`
}

func (x *Lun) Reset() {
	*x = Lun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lun) ProtoMessage() {}

func (x *Lun) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lun.ProtoReflect.Descriptor instead.
func (*Lun) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_lun_proto_rawDescGZIP(), []int{0}
}

func (x *Lun) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lun) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Lun) GetState() Lun_State {
	if x != nil {
		return x.State
	}
	return Lun_STATE_UNSPECIFIED
}

func (x *Lun) GetSizeGb() int64 {
	if x != nil {
		return x.SizeGb
	}
	return 0
}

func (x *Lun) GetMultiprotocolType() Lun_MultiprotocolType {
	if x != nil {
		return x.MultiprotocolType
	}
	return Lun_MULTIPROTOCOL_TYPE_UNSPECIFIED
}

func (x *Lun) GetStorageVolume() string {
	if x != nil {
		return x.StorageVolume
	}
	return ""
}

func (x *Lun) GetShareable() bool {
	if x != nil {
		return x.Shareable
	}
	return false
}

func (x *Lun) GetBootLun() bool {
	if x != nil {
		return x.BootLun
	}
	return false
}

func (x *Lun) GetStorageType() Lun_StorageType {
	if x != nil {
		return x.StorageType
	}
	return Lun_STORAGE_TYPE_UNSPECIFIED
}

func (x *Lun) GetWwid() string {
	if x != nil {
		return x.Wwid
	}
	return ""
}

// Message for requesting storage lun information.
type GetLunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetLunRequest) Reset() {
	*x = GetLunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLunRequest) ProtoMessage() {}

func (x *GetLunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLunRequest.ProtoReflect.Descriptor instead.
func (*GetLunRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_lun_proto_rawDescGZIP(), []int{1}
}

func (x *GetLunRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Message for requesting a list of storage volume luns.
type ListLunsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Parent value for ListLunsRequest.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Requested page size. The server might return fewer items than requested.
	// If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A token identifying a page of results from the server.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListLunsRequest) Reset() {
	*x = ListLunsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLunsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLunsRequest) ProtoMessage() {}

func (x *ListLunsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLunsRequest.ProtoReflect.Descriptor instead.
func (*ListLunsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_lun_proto_rawDescGZIP(), []int{2}
}

func (x *ListLunsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListLunsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListLunsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message containing the list of storage volume luns.
type ListLunsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of luns.
	Luns []*Lun `protobuf:"bytes,1,rep,name=luns,proto3" json:"luns,omitempty"`
	// A token identifying a page of results from the server.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Locations that could not be reached.
	Unreachable []string `protobuf:"bytes,3,rep,name=unreachable,proto3" json:"unreachable,omitempty"`
}

func (x *ListLunsResponse) Reset() {
	*x = ListLunsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLunsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLunsResponse) ProtoMessage() {}

func (x *ListLunsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLunsResponse.ProtoReflect.Descriptor instead.
func (*ListLunsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_baremetalsolution_v2_lun_proto_rawDescGZIP(), []int{3}
}

func (x *ListLunsResponse) GetLuns() []*Lun {
	if x != nil {
		return x.Luns
	}
	return nil
}

func (x *ListLunsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListLunsResponse) GetUnreachable() []string {
	if x != nil {
		return x.Unreachable
	}
	return nil
}

var File_google_cloud_baremetalsolution_v2_lun_proto protoreflect.FileDescriptor

var file_google_cloud_baremetalsolution_v2_lun_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x06, 0x0a,
	0x03, 0x4c, 0x75, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x4c, 0x75, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x67, 0x62, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x47, 0x62, 0x12, 0x67, 0x0a, 0x12, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x75, 0x6e, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xfa, 0x41, 0x29,
	0x0a, 0x27, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x74, 0x5f, 0x6c,
	0x75, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x74, 0x4c, 0x75,
	0x6e, 0x12, 0x55, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x75, 0x6e, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x77, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x77, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x04, 0x22, 0x42, 0x0a, 0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49,
	0x4e, 0x55, 0x58, 0x10, 0x01, 0x22, 0x3d, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x53, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x48,
	0x44, 0x44, 0x10, 0x02, 0x3a, 0x6e, 0xea, 0x41, 0x6b, 0x0a, 0x24, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x6e, 0x12,
	0x43, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73,
	0x2f, 0x7b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x7d, 0x2f, 0x6c, 0x75, 0x6e, 0x73, 0x2f, 0x7b,
	0x6c, 0x75, 0x6e, 0x7d, 0x22, 0x51, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x75, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x62, 0x61, 0x72,
	0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75,
	0x6e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x4c, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x29, 0x0a, 0x27, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x98, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x6c, 0x75, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x75, 0x6e, 0x52, 0x04, 0x6c, 0x75, 0x6e,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72,
	0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x75, 0x6e, 0x72, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x42, 0xf6, 0x01, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x42, 0x08, 0x4c, 0x75, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x52, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x3b, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x42, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x42, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x24,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x42,
	0x61, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_baremetalsolution_v2_lun_proto_rawDescOnce sync.Once
	file_google_cloud_baremetalsolution_v2_lun_proto_rawDescData = file_google_cloud_baremetalsolution_v2_lun_proto_rawDesc
)

func file_google_cloud_baremetalsolution_v2_lun_proto_rawDescGZIP() []byte {
	file_google_cloud_baremetalsolution_v2_lun_proto_rawDescOnce.Do(func() {
		file_google_cloud_baremetalsolution_v2_lun_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_baremetalsolution_v2_lun_proto_rawDescData)
	})
	return file_google_cloud_baremetalsolution_v2_lun_proto_rawDescData
}

var file_google_cloud_baremetalsolution_v2_lun_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_baremetalsolution_v2_lun_proto_goTypes = []interface{}{
	(Lun_State)(0),             // 0: google.cloud.baremetalsolution.v2.Lun.State
	(Lun_MultiprotocolType)(0), // 1: google.cloud.baremetalsolution.v2.Lun.MultiprotocolType
	(Lun_StorageType)(0),       // 2: google.cloud.baremetalsolution.v2.Lun.StorageType
	(*Lun)(nil),                // 3: google.cloud.baremetalsolution.v2.Lun
	(*GetLunRequest)(nil),      // 4: google.cloud.baremetalsolution.v2.GetLunRequest
	(*ListLunsRequest)(nil),    // 5: google.cloud.baremetalsolution.v2.ListLunsRequest
	(*ListLunsResponse)(nil),   // 6: google.cloud.baremetalsolution.v2.ListLunsResponse
}
var file_google_cloud_baremetalsolution_v2_lun_proto_depIdxs = []int32{
	0, // 0: google.cloud.baremetalsolution.v2.Lun.state:type_name -> google.cloud.baremetalsolution.v2.Lun.State
	1, // 1: google.cloud.baremetalsolution.v2.Lun.multiprotocol_type:type_name -> google.cloud.baremetalsolution.v2.Lun.MultiprotocolType
	2, // 2: google.cloud.baremetalsolution.v2.Lun.storage_type:type_name -> google.cloud.baremetalsolution.v2.Lun.StorageType
	3, // 3: google.cloud.baremetalsolution.v2.ListLunsResponse.luns:type_name -> google.cloud.baremetalsolution.v2.Lun
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_baremetalsolution_v2_lun_proto_init() }
func file_google_cloud_baremetalsolution_v2_lun_proto_init() {
	if File_google_cloud_baremetalsolution_v2_lun_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lun); i {
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
		file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLunRequest); i {
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
		file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLunsRequest); i {
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
		file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLunsResponse); i {
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
			RawDescriptor: file_google_cloud_baremetalsolution_v2_lun_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_baremetalsolution_v2_lun_proto_goTypes,
		DependencyIndexes: file_google_cloud_baremetalsolution_v2_lun_proto_depIdxs,
		EnumInfos:         file_google_cloud_baremetalsolution_v2_lun_proto_enumTypes,
		MessageInfos:      file_google_cloud_baremetalsolution_v2_lun_proto_msgTypes,
	}.Build()
	File_google_cloud_baremetalsolution_v2_lun_proto = out.File
	file_google_cloud_baremetalsolution_v2_lun_proto_rawDesc = nil
	file_google_cloud_baremetalsolution_v2_lun_proto_goTypes = nil
	file_google_cloud_baremetalsolution_v2_lun_proto_depIdxs = nil
}
