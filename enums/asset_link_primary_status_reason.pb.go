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
// source: google/ads/googleads/v13/enums/asset_link_primary_status_reason.proto

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

// Enum Provides insight into why an asset is not serving or not serving
// at full capacity for a particular link level. These reasons are
// aggregated to determine a final PrimaryStatus.
// For example, a sitelink might be paused by the user,
// but also limited in serving due to violation of an alcohol policy. In
// this case, the PrimaryStatus will be returned as PAUSED, since the asset's
// effective status is determined by its paused state.
type AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason int32

const (
	// Not specified.
	AssetLinkPrimaryStatusReasonEnum_UNSPECIFIED AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason = 0
	// Used for return value only. Represents value unknown in this version.
	AssetLinkPrimaryStatusReasonEnum_UNKNOWN AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason = 1
	// The asset is paused for its linked rollup level. Contributes to a
	// PrimaryStatus of PAUSED.
	AssetLinkPrimaryStatusReasonEnum_ASSET_LINK_PAUSED AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason = 2
	// The asset is removed for its linked rollup level. Contributes to a
	// PrimaryStatus of REMOVED.
	AssetLinkPrimaryStatusReasonEnum_ASSET_LINK_REMOVED AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason = 3
	// The asset has been marked as disapproved. Contributes to a PrimaryStatus
	// of NOT_ELIGIBLE
	AssetLinkPrimaryStatusReasonEnum_ASSET_DISAPPROVED AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason = 4
	// The asset has not completed policy review. Contributes to a PrimaryStatus
	// of PENDING.
	AssetLinkPrimaryStatusReasonEnum_ASSET_UNDER_REVIEW AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason = 5
	// The asset is approved with policies applied. Contributes to a
	// PrimaryStatus of LIMITED.
	AssetLinkPrimaryStatusReasonEnum_ASSET_APPROVED_LABELED AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason = 6
)

// Enum value maps for AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason.
var (
	AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ASSET_LINK_PAUSED",
		3: "ASSET_LINK_REMOVED",
		4: "ASSET_DISAPPROVED",
		5: "ASSET_UNDER_REVIEW",
		6: "ASSET_APPROVED_LABELED",
	}
	AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason_value = map[string]int32{
		"UNSPECIFIED":            0,
		"UNKNOWN":                1,
		"ASSET_LINK_PAUSED":      2,
		"ASSET_LINK_REMOVED":     3,
		"ASSET_DISAPPROVED":      4,
		"ASSET_UNDER_REVIEW":     5,
		"ASSET_APPROVED_LABELED": 6,
	}
)

func (x AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason) Enum() *AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason {
	p := new(AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason)
	*p = x
	return p
}

func (x AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_enumTypes[0].Descriptor()
}

func (AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_enumTypes[0]
}

func (x AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason.Descriptor instead.
func (AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDescGZIP(), []int{0, 0}
}

// Provides the reason of a primary status.
// For example: a sitelink may be paused for a particular campaign.
type AssetLinkPrimaryStatusReasonEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetLinkPrimaryStatusReasonEnum) Reset() {
	*x = AssetLinkPrimaryStatusReasonEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetLinkPrimaryStatusReasonEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetLinkPrimaryStatusReasonEnum) ProtoMessage() {}

func (x *AssetLinkPrimaryStatusReasonEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetLinkPrimaryStatusReasonEnum.ProtoReflect.Descriptor instead.
func (*AssetLinkPrimaryStatusReasonEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x20, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xb6, 0x01, 0x0a,
	0x1c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x41,
	0x53, 0x53, 0x45, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b,
	0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x53,
	0x53, 0x45, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10,
	0x04, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52,
	0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x53, 0x53,
	0x45, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x5f, 0x4c, 0x41, 0x42, 0x45,
	0x4c, 0x45, 0x44, 0x10, 0x06, 0x42, 0xfb, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x21, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x33, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDescData = file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDesc
)

func file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDescData
}

var file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_goTypes = []interface{}{
	(AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason)(0), // 0: google.ads.googleads.v13.enums.AssetLinkPrimaryStatusReasonEnum.AssetLinkPrimaryStatusReason
	(*AssetLinkPrimaryStatusReasonEnum)(nil),                           // 1: google.ads.googleads.v13.enums.AssetLinkPrimaryStatusReasonEnum
}
var file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_init() }
func file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_init() {
	if File_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetLinkPrimaryStatusReasonEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto = out.File
	file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_rawDesc = nil
	file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_goTypes = nil
	file_google_ads_googleads_v13_enums_asset_link_primary_status_reason_proto_depIdxs = nil
}
