// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: google/ads/googleads/v14/enums/media_type.proto

package enums

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

// The type of media.
type MediaTypeEnum_MediaType int32

const (
	// The media type has not been specified.
	MediaTypeEnum_UNSPECIFIED MediaTypeEnum_MediaType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	MediaTypeEnum_UNKNOWN MediaTypeEnum_MediaType = 1
	// Static image, used for image ad.
	MediaTypeEnum_IMAGE MediaTypeEnum_MediaType = 2
	// Small image, used for map ad.
	MediaTypeEnum_ICON MediaTypeEnum_MediaType = 3
	// ZIP file, used in fields of template ads.
	MediaTypeEnum_MEDIA_BUNDLE MediaTypeEnum_MediaType = 4
	// Audio file.
	MediaTypeEnum_AUDIO MediaTypeEnum_MediaType = 5
	// Video file.
	MediaTypeEnum_VIDEO MediaTypeEnum_MediaType = 6
	// Animated image, such as animated GIF.
	MediaTypeEnum_DYNAMIC_IMAGE MediaTypeEnum_MediaType = 7
)

// Enum value maps for MediaTypeEnum_MediaType.
var (
	MediaTypeEnum_MediaType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "IMAGE",
		3: "ICON",
		4: "MEDIA_BUNDLE",
		5: "AUDIO",
		6: "VIDEO",
		7: "DYNAMIC_IMAGE",
	}
	MediaTypeEnum_MediaType_value = map[string]int32{
		"UNSPECIFIED":   0,
		"UNKNOWN":       1,
		"IMAGE":         2,
		"ICON":          3,
		"MEDIA_BUNDLE":  4,
		"AUDIO":         5,
		"VIDEO":         6,
		"DYNAMIC_IMAGE": 7,
	}
)

func (x MediaTypeEnum_MediaType) Enum() *MediaTypeEnum_MediaType {
	p := new(MediaTypeEnum_MediaType)
	*p = x
	return p
}

func (x MediaTypeEnum_MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaTypeEnum_MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v14_enums_media_type_proto_enumTypes[0].Descriptor()
}

func (MediaTypeEnum_MediaType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v14_enums_media_type_proto_enumTypes[0]
}

func (x MediaTypeEnum_MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaTypeEnum_MediaType.Descriptor instead.
func (MediaTypeEnum_MediaType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_enums_media_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the types of media.
type MediaTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MediaTypeEnum) Reset() {
	*x = MediaTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_enums_media_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaTypeEnum) ProtoMessage() {}

func (x *MediaTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_enums_media_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaTypeEnum.ProtoReflect.Descriptor instead.
func (*MediaTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_enums_media_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v14_enums_media_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v14_enums_media_type_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x22, 0x8a, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x45,
	0x6e, 0x75, 0x6d, 0x22, 0x79, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x43, 0x4f,
	0x4e, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x42, 0x55, 0x4e,
	0x44, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x05,
	0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x44,
	0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x07, 0x42, 0xe8,
	0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x0e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x34, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x34, 0x5c, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x34, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v14_enums_media_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v14_enums_media_type_proto_rawDescData = file_google_ads_googleads_v14_enums_media_type_proto_rawDesc
)

func file_google_ads_googleads_v14_enums_media_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v14_enums_media_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v14_enums_media_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v14_enums_media_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v14_enums_media_type_proto_rawDescData
}

var file_google_ads_googleads_v14_enums_media_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v14_enums_media_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v14_enums_media_type_proto_goTypes = []interface{}{
	(MediaTypeEnum_MediaType)(0), // 0: google.ads.googleads.v14.enums.MediaTypeEnum.MediaType
	(*MediaTypeEnum)(nil),        // 1: google.ads.googleads.v14.enums.MediaTypeEnum
}
var file_google_ads_googleads_v14_enums_media_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v14_enums_media_type_proto_init() }
func file_google_ads_googleads_v14_enums_media_type_proto_init() {
	if File_google_ads_googleads_v14_enums_media_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v14_enums_media_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v14_enums_media_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v14_enums_media_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v14_enums_media_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v14_enums_media_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v14_enums_media_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v14_enums_media_type_proto = out.File
	file_google_ads_googleads_v14_enums_media_type_proto_rawDesc = nil
	file_google_ads_googleads_v14_enums_media_type_proto_goTypes = nil
	file_google_ads_googleads_v14_enums_media_type_proto_depIdxs = nil
}
