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
// 	protoc        v4.24.4
// source: google/ads/googleads/v15/errors/asset_group_error.proto

package errors

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

// Enum describing possible asset group errors.
type AssetGroupErrorEnum_AssetGroupError int32

const (
	// Enum unspecified.
	AssetGroupErrorEnum_UNSPECIFIED AssetGroupErrorEnum_AssetGroupError = 0
	// The received error code is not known in this version.
	AssetGroupErrorEnum_UNKNOWN AssetGroupErrorEnum_AssetGroupError = 1
	// Each asset group in a single campaign must have a unique name.
	AssetGroupErrorEnum_DUPLICATE_NAME AssetGroupErrorEnum_AssetGroupError = 2
	// Cannot add asset group for the campaign type.
	AssetGroupErrorEnum_CANNOT_ADD_ASSET_GROUP_FOR_CAMPAIGN_TYPE AssetGroupErrorEnum_AssetGroupError = 3
	// Not enough headline asset for a valid asset group.
	AssetGroupErrorEnum_NOT_ENOUGH_HEADLINE_ASSET AssetGroupErrorEnum_AssetGroupError = 4
	// Not enough long headline asset for a valid asset group.
	AssetGroupErrorEnum_NOT_ENOUGH_LONG_HEADLINE_ASSET AssetGroupErrorEnum_AssetGroupError = 5
	// Not enough description headline asset for a valid asset group.
	AssetGroupErrorEnum_NOT_ENOUGH_DESCRIPTION_ASSET AssetGroupErrorEnum_AssetGroupError = 6
	// Not enough business name asset for a valid asset group.
	AssetGroupErrorEnum_NOT_ENOUGH_BUSINESS_NAME_ASSET AssetGroupErrorEnum_AssetGroupError = 7
	// Not enough marketing image asset for a valid asset group.
	AssetGroupErrorEnum_NOT_ENOUGH_MARKETING_IMAGE_ASSET AssetGroupErrorEnum_AssetGroupError = 8
	// Not enough square marketing image asset for a valid asset group.
	AssetGroupErrorEnum_NOT_ENOUGH_SQUARE_MARKETING_IMAGE_ASSET AssetGroupErrorEnum_AssetGroupError = 9
	// Not enough logo asset for a valid asset group.
	AssetGroupErrorEnum_NOT_ENOUGH_LOGO_ASSET AssetGroupErrorEnum_AssetGroupError = 10
	// Final url and shopping merchant url does not have the same domain.
	AssetGroupErrorEnum_FINAL_URL_SHOPPING_MERCHANT_HOME_PAGE_URL_DOMAINS_DIFFER AssetGroupErrorEnum_AssetGroupError = 11
	// Path1 required when path2 is set.
	AssetGroupErrorEnum_PATH1_REQUIRED_WHEN_PATH2_IS_SET AssetGroupErrorEnum_AssetGroupError = 12
	// At least one short description asset is required for a valid asset group.
	AssetGroupErrorEnum_SHORT_DESCRIPTION_REQUIRED AssetGroupErrorEnum_AssetGroupError = 13
	// Final url field is required for asset group.
	AssetGroupErrorEnum_FINAL_URL_REQUIRED AssetGroupErrorEnum_AssetGroupError = 14
	// Final url contains invalid domain name.
	AssetGroupErrorEnum_FINAL_URL_CONTAINS_INVALID_DOMAIN_NAME AssetGroupErrorEnum_AssetGroupError = 15
	// Ad customizers are not supported in asset group's text field.
	AssetGroupErrorEnum_AD_CUSTOMIZER_NOT_SUPPORTED AssetGroupErrorEnum_AssetGroupError = 16
	// Cannot mutate asset group for campaign with removed status.
	AssetGroupErrorEnum_CANNOT_MUTATE_ASSET_GROUP_FOR_REMOVED_CAMPAIGN AssetGroupErrorEnum_AssetGroupError = 17
)

// Enum value maps for AssetGroupErrorEnum_AssetGroupError.
var (
	AssetGroupErrorEnum_AssetGroupError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "DUPLICATE_NAME",
		3:  "CANNOT_ADD_ASSET_GROUP_FOR_CAMPAIGN_TYPE",
		4:  "NOT_ENOUGH_HEADLINE_ASSET",
		5:  "NOT_ENOUGH_LONG_HEADLINE_ASSET",
		6:  "NOT_ENOUGH_DESCRIPTION_ASSET",
		7:  "NOT_ENOUGH_BUSINESS_NAME_ASSET",
		8:  "NOT_ENOUGH_MARKETING_IMAGE_ASSET",
		9:  "NOT_ENOUGH_SQUARE_MARKETING_IMAGE_ASSET",
		10: "NOT_ENOUGH_LOGO_ASSET",
		11: "FINAL_URL_SHOPPING_MERCHANT_HOME_PAGE_URL_DOMAINS_DIFFER",
		12: "PATH1_REQUIRED_WHEN_PATH2_IS_SET",
		13: "SHORT_DESCRIPTION_REQUIRED",
		14: "FINAL_URL_REQUIRED",
		15: "FINAL_URL_CONTAINS_INVALID_DOMAIN_NAME",
		16: "AD_CUSTOMIZER_NOT_SUPPORTED",
		17: "CANNOT_MUTATE_ASSET_GROUP_FOR_REMOVED_CAMPAIGN",
	}
	AssetGroupErrorEnum_AssetGroupError_value = map[string]int32{
		"UNSPECIFIED":    0,
		"UNKNOWN":        1,
		"DUPLICATE_NAME": 2,
		"CANNOT_ADD_ASSET_GROUP_FOR_CAMPAIGN_TYPE":                 3,
		"NOT_ENOUGH_HEADLINE_ASSET":                                4,
		"NOT_ENOUGH_LONG_HEADLINE_ASSET":                           5,
		"NOT_ENOUGH_DESCRIPTION_ASSET":                             6,
		"NOT_ENOUGH_BUSINESS_NAME_ASSET":                           7,
		"NOT_ENOUGH_MARKETING_IMAGE_ASSET":                         8,
		"NOT_ENOUGH_SQUARE_MARKETING_IMAGE_ASSET":                  9,
		"NOT_ENOUGH_LOGO_ASSET":                                    10,
		"FINAL_URL_SHOPPING_MERCHANT_HOME_PAGE_URL_DOMAINS_DIFFER": 11,
		"PATH1_REQUIRED_WHEN_PATH2_IS_SET":                         12,
		"SHORT_DESCRIPTION_REQUIRED":                               13,
		"FINAL_URL_REQUIRED":                                       14,
		"FINAL_URL_CONTAINS_INVALID_DOMAIN_NAME":                   15,
		"AD_CUSTOMIZER_NOT_SUPPORTED":                              16,
		"CANNOT_MUTATE_ASSET_GROUP_FOR_REMOVED_CAMPAIGN":           17,
	}
)

func (x AssetGroupErrorEnum_AssetGroupError) Enum() *AssetGroupErrorEnum_AssetGroupError {
	p := new(AssetGroupErrorEnum_AssetGroupError)
	*p = x
	return p
}

func (x AssetGroupErrorEnum_AssetGroupError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetGroupErrorEnum_AssetGroupError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v15_errors_asset_group_error_proto_enumTypes[0].Descriptor()
}

func (AssetGroupErrorEnum_AssetGroupError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v15_errors_asset_group_error_proto_enumTypes[0]
}

func (x AssetGroupErrorEnum_AssetGroupError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetGroupErrorEnum_AssetGroupError.Descriptor instead.
func (AssetGroupErrorEnum_AssetGroupError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible asset group errors.
type AssetGroupErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetGroupErrorEnum) Reset() {
	*x = AssetGroupErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_errors_asset_group_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetGroupErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupErrorEnum) ProtoMessage() {}

func (x *AssetGroupErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_errors_asset_group_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupErrorEnum.ProtoReflect.Descriptor instead.
func (*AssetGroupErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v15_errors_asset_group_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x9d, 0x05, 0x0a, 0x13, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0x85, 0x05, 0x0a, 0x0f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x2c, 0x0a, 0x28, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e,
	0x4f, 0x55, 0x47, 0x48, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x41, 0x53,
	0x53, 0x45, 0x54, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f,
	0x55, 0x47, 0x48, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e,
	0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x4e, 0x4f, 0x54,
	0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x06, 0x12, 0x22, 0x0a, 0x1e, 0x4e,
	0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x42, 0x55, 0x53, 0x49, 0x4e, 0x45,
	0x53, 0x53, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x07, 0x12,
	0x24, 0x0a, 0x20, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x4d, 0x41,
	0x52, 0x4b, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x53,
	0x53, 0x45, 0x54, 0x10, 0x08, 0x12, 0x2b, 0x0a, 0x27, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f,
	0x55, 0x47, 0x48, 0x5f, 0x53, 0x51, 0x55, 0x41, 0x52, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45,
	0x54, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54,
	0x10, 0x09, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48,
	0x5f, 0x4c, 0x4f, 0x47, 0x4f, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x0a, 0x12, 0x3c, 0x0a,
	0x38, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x53, 0x48, 0x4f, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x52, 0x43, 0x48, 0x41, 0x4e, 0x54, 0x5f, 0x48, 0x4f, 0x4d,
	0x45, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49,
	0x4e, 0x53, 0x5f, 0x44, 0x49, 0x46, 0x46, 0x45, 0x52, 0x10, 0x0b, 0x12, 0x24, 0x0a, 0x20, 0x50,
	0x41, 0x54, 0x48, 0x31, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x57, 0x48,
	0x45, 0x4e, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x32, 0x5f, 0x49, 0x53, 0x5f, 0x53, 0x45, 0x54, 0x10,
	0x0c, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52,
	0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x0d, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0e, 0x12, 0x2a, 0x0a, 0x26, 0x46, 0x49, 0x4e,
	0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x10, 0x0f, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x49, 0x5a, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f,
	0x52, 0x54, 0x45, 0x44, 0x10, 0x10, 0x12, 0x32, 0x0a, 0x2e, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x4d, 0x55, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x47, 0x52,
	0x4f, 0x55, 0x50, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x5f,
	0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x11, 0x42, 0xf4, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x42, 0x14, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x35, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x35, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x35, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDescData = file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDesc
)

func file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDescData
}

var file_google_ads_googleads_v15_errors_asset_group_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v15_errors_asset_group_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v15_errors_asset_group_error_proto_goTypes = []interface{}{
	(AssetGroupErrorEnum_AssetGroupError)(0), // 0: google.ads.googleads.v15.errors.AssetGroupErrorEnum.AssetGroupError
	(*AssetGroupErrorEnum)(nil),              // 1: google.ads.googleads.v15.errors.AssetGroupErrorEnum
}
var file_google_ads_googleads_v15_errors_asset_group_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v15_errors_asset_group_error_proto_init() }
func file_google_ads_googleads_v15_errors_asset_group_error_proto_init() {
	if File_google_ads_googleads_v15_errors_asset_group_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_errors_asset_group_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetGroupErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v15_errors_asset_group_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_errors_asset_group_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v15_errors_asset_group_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v15_errors_asset_group_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_errors_asset_group_error_proto = out.File
	file_google_ads_googleads_v15_errors_asset_group_error_proto_rawDesc = nil
	file_google_ads_googleads_v15_errors_asset_group_error_proto_goTypes = nil
	file_google_ads_googleads_v15_errors_asset_group_error_proto_depIdxs = nil
}
