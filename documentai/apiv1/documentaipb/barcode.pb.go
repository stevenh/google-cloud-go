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
// source: google/cloud/documentai/v1/barcode.proto

package documentaipb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Encodes the detailed information of a barcode.
type Barcode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format of a barcode.
	// The supported formats are:
	//
	//	CODE_128: Code 128 type.
	//	CODE_39: Code 39 type.
	//	CODE_93: Code 93 type.
	//	CODABAR: Codabar type.
	//	DATA_MATRIX: 2D Data Matrix type.
	//	ITF: ITF type.
	//	EAN_13: EAN-13 type.
	//	EAN_8: EAN-8 type.
	//	QR_CODE: 2D QR code type.
	//	UPC_A: UPC-A type.
	//	UPC_E: UPC-E type.
	//	PDF417: PDF417 type.
	//	AZTEC: 2D Aztec code type.
	//	DATABAR: GS1 DataBar code type.
	Format string `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty"`
	// Value format describes the format of the value that a barcode
	// encodes.
	// The supported formats are:
	//
	//	CONTACT_INFO: Contact information.
	//	EMAIL: Email address.
	//	ISBN: ISBN identifier.
	//	PHONE: Phone number.
	//	PRODUCT: Product.
	//	SMS: SMS message.
	//	TEXT: Text string.
	//	URL: URL address.
	//	WIFI: Wifi information.
	//	GEO: Geo-localization.
	//	CALENDAR_EVENT: Calendar event.
	//	DRIVER_LICENSE: Driver's license.
	ValueFormat string `protobuf:"bytes,2,opt,name=value_format,json=valueFormat,proto3" json:"value_format,omitempty"`
	// Raw value encoded in the barcode.
	// For example, 'MEBKM:TITLE:Google;URL:https://www.google.com;;'.
	RawValue string `protobuf:"bytes,3,opt,name=raw_value,json=rawValue,proto3" json:"raw_value,omitempty"`
}

func (x *Barcode) Reset() {
	*x = Barcode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_documentai_v1_barcode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Barcode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Barcode) ProtoMessage() {}

func (x *Barcode) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_documentai_v1_barcode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Barcode.ProtoReflect.Descriptor instead.
func (*Barcode) Descriptor() ([]byte, []int) {
	return file_google_cloud_documentai_v1_barcode_proto_rawDescGZIP(), []int{0}
}

func (x *Barcode) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Barcode) GetValueFormat() string {
	if x != nil {
		return x.ValueFormat
	}
	return ""
}

func (x *Barcode) GetRawValue() string {
	if x != nil {
		return x.RawValue
	}
	return ""
}

var File_google_cloud_documentai_v1_barcode_proto protoreflect.FileDescriptor

var file_google_cloud_documentai_v1_barcode_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x61, 0x0a, 0x07, 0x42, 0x61, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x61, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x61, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xd0, 0x01, 0x0a, 0x1e, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x42, 0x61,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x69, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x49, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_documentai_v1_barcode_proto_rawDescOnce sync.Once
	file_google_cloud_documentai_v1_barcode_proto_rawDescData = file_google_cloud_documentai_v1_barcode_proto_rawDesc
)

func file_google_cloud_documentai_v1_barcode_proto_rawDescGZIP() []byte {
	file_google_cloud_documentai_v1_barcode_proto_rawDescOnce.Do(func() {
		file_google_cloud_documentai_v1_barcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_documentai_v1_barcode_proto_rawDescData)
	})
	return file_google_cloud_documentai_v1_barcode_proto_rawDescData
}

var file_google_cloud_documentai_v1_barcode_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_documentai_v1_barcode_proto_goTypes = []interface{}{
	(*Barcode)(nil), // 0: google.cloud.documentai.v1.Barcode
}
var file_google_cloud_documentai_v1_barcode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_documentai_v1_barcode_proto_init() }
func file_google_cloud_documentai_v1_barcode_proto_init() {
	if File_google_cloud_documentai_v1_barcode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_documentai_v1_barcode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Barcode); i {
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
			RawDescriptor: file_google_cloud_documentai_v1_barcode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_documentai_v1_barcode_proto_goTypes,
		DependencyIndexes: file_google_cloud_documentai_v1_barcode_proto_depIdxs,
		MessageInfos:      file_google_cloud_documentai_v1_barcode_proto_msgTypes,
	}.Build()
	File_google_cloud_documentai_v1_barcode_proto = out.File
	file_google_cloud_documentai_v1_barcode_proto_rawDesc = nil
	file_google_cloud_documentai_v1_barcode_proto_goTypes = nil
	file_google_cloud_documentai_v1_barcode_proto_depIdxs = nil
}
