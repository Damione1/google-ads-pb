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
// source: google/ads/googleads/v13/resources/ad_group_criterion_customizer.proto

package resources

import (
	common "github.com/shenzhencenter/google-ads-pb/common"
	enums "github.com/shenzhencenter/google-ads-pb/enums"
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

// A customizer value for the associated CustomizerAttribute at the
// AdGroupCriterion level.
type AdGroupCriterionCustomizer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the ad group criterion customizer.
	// Ad group criterion customizer resource names have the form:
	//
	// `customers/{customer_id}/adGroupCriterionCustomizers/{ad_group_id}~{criterion_id}~{customizer_attribute_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The ad group criterion to which the customizer attribute is
	// linked. It must be a keyword criterion.
	AdGroupCriterion *string `protobuf:"bytes,2,opt,name=ad_group_criterion,json=adGroupCriterion,proto3,oneof" json:"ad_group_criterion,omitempty"`
	// Required. Immutable. The customizer attribute which is linked to the ad
	// group criterion.
	CustomizerAttribute string `protobuf:"bytes,3,opt,name=customizer_attribute,json=customizerAttribute,proto3" json:"customizer_attribute,omitempty"`
	// Output only. The status of the ad group criterion customizer.
	Status enums.CustomizerValueStatusEnum_CustomizerValueStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v13.enums.CustomizerValueStatusEnum_CustomizerValueStatus" json:"status,omitempty"`
	// Required. The value to associate with the customizer attribute at this
	// level. The value must be of the type specified for the CustomizerAttribute.
	Value *common.CustomizerValue `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AdGroupCriterionCustomizer) Reset() {
	*x = AdGroupCriterionCustomizer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupCriterionCustomizer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupCriterionCustomizer) ProtoMessage() {}

func (x *AdGroupCriterionCustomizer) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupCriterionCustomizer.ProtoReflect.Descriptor instead.
func (*AdGroupCriterionCustomizer) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDescGZIP(), []int{0}
}

func (x *AdGroupCriterionCustomizer) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AdGroupCriterionCustomizer) GetAdGroupCriterion() string {
	if x != nil && x.AdGroupCriterion != nil {
		return *x.AdGroupCriterion
	}
	return ""
}

func (x *AdGroupCriterionCustomizer) GetCustomizerAttribute() string {
	if x != nil {
		return x.CustomizerAttribute
	}
	return ""
}

func (x *AdGroupCriterionCustomizer) GetStatus() enums.CustomizerValueStatusEnum_CustomizerValueStatus {
	if x != nil {
		return x.Status
	}
	return enums.CustomizerValueStatusEnum_CustomizerValueStatus(0)
}

func (x *AdGroupCriterionCustomizer) GetValue() *common.CustomizerValue {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x36, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x33, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca,
	0x05, 0x0a, 0x1a, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x60, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x3b, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x35, 0x0a, 0x33, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65,
	0x72, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x64, 0x0a, 0x12, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xe0, 0x41, 0x05,
	0xfa, 0x41, 0x2b, 0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x10, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x6a, 0x0a, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x65, 0x72, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x37, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2e, 0x0a, 0x2c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69,
	0x7a, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x13, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x12, 0x6c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x4f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x4b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0xa5, 0x01, 0xea,
	0x41, 0xa1, 0x01, 0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x6a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x61,
	0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x63, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x42, 0x91, 0x02, 0x0a, 0x26,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x33, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1f, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x69, 0x7a,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x33, 0x5c, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x33, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDescData = file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDesc
)

func file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDescData)
	})
	return file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDescData
}

var file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_goTypes = []interface{}{
	(*AdGroupCriterionCustomizer)(nil),                         // 0: google.ads.googleads.v13.resources.AdGroupCriterionCustomizer
	(enums.CustomizerValueStatusEnum_CustomizerValueStatus)(0), // 1: google.ads.googleads.v13.enums.CustomizerValueStatusEnum.CustomizerValueStatus
	(*common.CustomizerValue)(nil),                             // 2: google.ads.googleads.v13.common.CustomizerValue
}
var file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v13.resources.AdGroupCriterionCustomizer.status:type_name -> google.ads.googleads.v13.enums.CustomizerValueStatusEnum.CustomizerValueStatus
	2, // 1: google.ads.googleads.v13.resources.AdGroupCriterionCustomizer.value:type_name -> google.ads.googleads.v13.common.CustomizerValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_init() }
func file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_init() {
	if File_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupCriterionCustomizer); i {
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
	file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto = out.File
	file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_rawDesc = nil
	file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_goTypes = nil
	file_google_ads_googleads_v13_resources_ad_group_criterion_customizer_proto_depIdxs = nil
}
