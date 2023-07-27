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
// source: google/ads/googleads/v14/common/final_app_url.proto

package common

import (
	enums "github.com/shenzhencenter/google-ads-pb/enums"
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

// A URL for deep linking into an app for the given operating system.
type FinalAppUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The operating system targeted by this URL. Required.
	OsType enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType `protobuf:"varint,1,opt,name=os_type,json=osType,proto3,enum=google.ads.googleads.v14.enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType" json:"os_type,omitempty"`
	// The app deep link URL. Deep links specify a location in an app that
	// corresponds to the content you'd like to show, and should be of the form
	// {scheme}://{host_path}
	// The scheme identifies which app to open. For your app, you can use a custom
	// scheme that starts with the app's name. The host and path specify the
	// unique location in the app where your content exists.
	// Example: "exampleapp://productid_1234". Required.
	Url *string `protobuf:"bytes,3,opt,name=url,proto3,oneof" json:"url,omitempty"`
}

func (x *FinalAppUrl) Reset() {
	*x = FinalAppUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v14_common_final_app_url_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalAppUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalAppUrl) ProtoMessage() {}

func (x *FinalAppUrl) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v14_common_final_app_url_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalAppUrl.ProtoReflect.Descriptor instead.
func (*FinalAppUrl) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v14_common_final_app_url_proto_rawDescGZIP(), []int{0}
}

func (x *FinalAppUrl) GetOsType() enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType {
	if x != nil {
		return x.OsType
	}
	return enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType(0)
}

func (x *FinalAppUrl) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

var File_google_ads_googleads_v14_common_final_app_url_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v14_common_final_app_url_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x46,
	0x69, 0x6e, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x55, 0x72, 0x6c, 0x12, 0x70, 0x0a, 0x07, 0x6f, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x57, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x55, 0x72, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x72,
	0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0xf0, 0x01, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x34, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x42, 0x10, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x55, 0x72, 0x6c,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x34,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x34, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x34, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x34, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v14_common_final_app_url_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v14_common_final_app_url_proto_rawDescData = file_google_ads_googleads_v14_common_final_app_url_proto_rawDesc
)

func file_google_ads_googleads_v14_common_final_app_url_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v14_common_final_app_url_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v14_common_final_app_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v14_common_final_app_url_proto_rawDescData)
	})
	return file_google_ads_googleads_v14_common_final_app_url_proto_rawDescData
}

var file_google_ads_googleads_v14_common_final_app_url_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v14_common_final_app_url_proto_goTypes = []interface{}{
	(*FinalAppUrl)(nil), // 0: google.ads.googleads.v14.common.FinalAppUrl
	(enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType)(0), // 1: google.ads.googleads.v14.enums.AppUrlOperatingSystemTypeEnum.AppUrlOperatingSystemType
}
var file_google_ads_googleads_v14_common_final_app_url_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v14.common.FinalAppUrl.os_type:type_name -> google.ads.googleads.v14.enums.AppUrlOperatingSystemTypeEnum.AppUrlOperatingSystemType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v14_common_final_app_url_proto_init() }
func file_google_ads_googleads_v14_common_final_app_url_proto_init() {
	if File_google_ads_googleads_v14_common_final_app_url_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v14_common_final_app_url_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalAppUrl); i {
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
	file_google_ads_googleads_v14_common_final_app_url_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v14_common_final_app_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v14_common_final_app_url_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v14_common_final_app_url_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v14_common_final_app_url_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v14_common_final_app_url_proto = out.File
	file_google_ads_googleads_v14_common_final_app_url_proto_rawDesc = nil
	file_google_ads_googleads_v14_common_final_app_url_proto_goTypes = nil
	file_google_ads_googleads_v14_common_final_app_url_proto_depIdxs = nil
}
