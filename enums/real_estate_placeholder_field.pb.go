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
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: google/ads/googleads/v13/enums/real_estate_placeholder_field.proto

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

// Possible values for Real Estate placeholder fields.
type RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField int32

const (
	// Not specified.
	RealEstatePlaceholderFieldEnum_UNSPECIFIED RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	RealEstatePlaceholderFieldEnum_UNKNOWN RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 1
	// Data Type: STRING. Unique ID.
	RealEstatePlaceholderFieldEnum_LISTING_ID RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 2
	// Data Type: STRING. Main headline with listing name to be shown in dynamic
	// ad.
	RealEstatePlaceholderFieldEnum_LISTING_NAME RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 3
	// Data Type: STRING. City name to be shown in dynamic ad.
	RealEstatePlaceholderFieldEnum_CITY_NAME RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 4
	// Data Type: STRING. Description of listing to be shown in dynamic ad.
	RealEstatePlaceholderFieldEnum_DESCRIPTION RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 5
	// Data Type: STRING. Complete listing address, including postal code.
	RealEstatePlaceholderFieldEnum_ADDRESS RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 6
	// Data Type: STRING. Price to be shown in the ad.
	// Example: "100.00 USD"
	RealEstatePlaceholderFieldEnum_PRICE RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 7
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	RealEstatePlaceholderFieldEnum_FORMATTED_PRICE RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 8
	// Data Type: URL. Image to be displayed in the ad.
	RealEstatePlaceholderFieldEnum_IMAGE_URL RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 9
	// Data Type: STRING. Type of property (house, condo, apartment, etc.) used
	// to group like items together for recommendation engine.
	RealEstatePlaceholderFieldEnum_PROPERTY_TYPE RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 10
	// Data Type: STRING. Type of listing (resale, rental, foreclosure, etc.)
	// used to group like items together for recommendation engine.
	RealEstatePlaceholderFieldEnum_LISTING_TYPE RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 11
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	RealEstatePlaceholderFieldEnum_CONTEXTUAL_KEYWORDS RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 12
	// Data Type: URL_LIST. Final URLs to be used in ad when using Upgraded
	// URLs; the more specific the better (for example, the individual URL of a
	// specific listing and its location).
	RealEstatePlaceholderFieldEnum_FINAL_URLS RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 13
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	RealEstatePlaceholderFieldEnum_FINAL_MOBILE_URLS RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 14
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	RealEstatePlaceholderFieldEnum_TRACKING_URL RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 15
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//
	//	scheme.
	//
	// host_path: identifies the specific content within your application.
	RealEstatePlaceholderFieldEnum_ANDROID_APP_LINK RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 16
	// Data Type: STRING_LIST. List of recommended listing IDs to show together
	// with this item.
	RealEstatePlaceholderFieldEnum_SIMILAR_LISTING_IDS RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 17
	// Data Type: STRING. iOS app link.
	RealEstatePlaceholderFieldEnum_IOS_APP_LINK RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 18
	// Data Type: INT64. iOS app store ID.
	RealEstatePlaceholderFieldEnum_IOS_APP_STORE_ID RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 19
)

// Enum value maps for RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField.
var (
	RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "LISTING_ID",
		3:  "LISTING_NAME",
		4:  "CITY_NAME",
		5:  "DESCRIPTION",
		6:  "ADDRESS",
		7:  "PRICE",
		8:  "FORMATTED_PRICE",
		9:  "IMAGE_URL",
		10: "PROPERTY_TYPE",
		11: "LISTING_TYPE",
		12: "CONTEXTUAL_KEYWORDS",
		13: "FINAL_URLS",
		14: "FINAL_MOBILE_URLS",
		15: "TRACKING_URL",
		16: "ANDROID_APP_LINK",
		17: "SIMILAR_LISTING_IDS",
		18: "IOS_APP_LINK",
		19: "IOS_APP_STORE_ID",
	}
	RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField_value = map[string]int32{
		"UNSPECIFIED":         0,
		"UNKNOWN":             1,
		"LISTING_ID":          2,
		"LISTING_NAME":        3,
		"CITY_NAME":           4,
		"DESCRIPTION":         5,
		"ADDRESS":             6,
		"PRICE":               7,
		"FORMATTED_PRICE":     8,
		"IMAGE_URL":           9,
		"PROPERTY_TYPE":       10,
		"LISTING_TYPE":        11,
		"CONTEXTUAL_KEYWORDS": 12,
		"FINAL_URLS":          13,
		"FINAL_MOBILE_URLS":   14,
		"TRACKING_URL":        15,
		"ANDROID_APP_LINK":    16,
		"SIMILAR_LISTING_IDS": 17,
		"IOS_APP_LINK":        18,
		"IOS_APP_STORE_ID":    19,
	}
)

func (x RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField) Enum() *RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField {
	p := new(RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField)
	*p = x
	return p
}

func (x RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_enumTypes[0]
}

func (x RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField.Descriptor instead.
func (RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

// Values for Real Estate placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type RealEstatePlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RealEstatePlaceholderFieldEnum) Reset() {
	*x = RealEstatePlaceholderFieldEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealEstatePlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealEstatePlaceholderFieldEnum) ProtoMessage() {}

func (x *RealEstatePlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealEstatePlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*RealEstatePlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x22, 0xa9, 0x03, 0x0a, 0x1e, 0x52, 0x65, 0x61, 0x6c, 0x45, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x86, 0x03, 0x0a, 0x1a, 0x52, 0x65, 0x61, 0x6c,
	0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x49, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f,
	0x4e, 0x41, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x49, 0x54, 0x59, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53,
	0x53, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x07, 0x12, 0x13,
	0x0a, 0x0f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x49, 0x43,
	0x45, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x52, 0x4c,
	0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0b, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x54, 0x45,
	0x58, 0x54, 0x55, 0x41, 0x4c, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x0c,
	0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0d,
	0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45,
	0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b,
	0x49, 0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4e, 0x44,
	0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x10, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x49, 0x4d, 0x49, 0x4c, 0x41, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x11, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4f, 0x53, 0x5f,
	0x41, 0x50, 0x50, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x12, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4f,
	0x53, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x13,
	0x42, 0xf9, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1f, 0x52, 0x65, 0x61, 0x6c, 0x45, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2,
	0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31,
	0x33, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDescData = file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_goTypes = []interface{}{
	(RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField)(0), // 0: google.ads.googleads.v13.enums.RealEstatePlaceholderFieldEnum.RealEstatePlaceholderField
	(*RealEstatePlaceholderFieldEnum)(nil),                         // 1: google.ads.googleads.v13.enums.RealEstatePlaceholderFieldEnum
}
var file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_init() }
func file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_init() {
	if File_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RealEstatePlaceholderFieldEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto = out.File
	file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v13_enums_real_estate_placeholder_field_proto_depIdxs = nil
}
