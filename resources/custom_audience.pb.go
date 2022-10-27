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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: google/ads/googleads/v12/resources/custom_audience.proto

package resources

import (
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

// A custom audience. This is a list of users by interest.
type CustomAudience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the custom audience.
	// Custom audience resource names have the form:
	//
	// `customers/{customer_id}/customAudiences/{custom_audience_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. ID of the custom audience.
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. Status of this custom audience. Indicates whether the custom audience is
	// enabled or removed.
	Status enums.CustomAudienceStatusEnum_CustomAudienceStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v12.enums.CustomAudienceStatusEnum_CustomAudienceStatus" json:"status,omitempty"`
	// Name of the custom audience. It should be unique for all custom audiences
	// created by a customer.
	// This field is required for creating operations.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the custom audience.
	// ("INTEREST" OR "PURCHASE_INTENT" is not allowed for newly created custom
	// audience but kept for existing audiences)
	Type enums.CustomAudienceTypeEnum_CustomAudienceType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v12.enums.CustomAudienceTypeEnum_CustomAudienceType" json:"type,omitempty"`
	// Description of this custom audience.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// List of custom audience members that this custom audience is composed of.
	// Members can be added during CustomAudience creation. If members are
	// presented in UPDATE operation, existing members will be overridden.
	Members []*CustomAudienceMember `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *CustomAudience) Reset() {
	*x = CustomAudience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_resources_custom_audience_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomAudience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomAudience) ProtoMessage() {}

func (x *CustomAudience) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_resources_custom_audience_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomAudience.ProtoReflect.Descriptor instead.
func (*CustomAudience) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_resources_custom_audience_proto_rawDescGZIP(), []int{0}
}

func (x *CustomAudience) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CustomAudience) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomAudience) GetStatus() enums.CustomAudienceStatusEnum_CustomAudienceStatus {
	if x != nil {
		return x.Status
	}
	return enums.CustomAudienceStatusEnum_CustomAudienceStatus(0)
}

func (x *CustomAudience) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomAudience) GetType() enums.CustomAudienceTypeEnum_CustomAudienceType {
	if x != nil {
		return x.Type
	}
	return enums.CustomAudienceTypeEnum_CustomAudienceType(0)
}

func (x *CustomAudience) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CustomAudience) GetMembers() []*CustomAudienceMember {
	if x != nil {
		return x.Members
	}
	return nil
}

// A member of custom audience. A member can be a KEYWORD, URL,
// PLACE_CATEGORY or APP. It can only be created or removed but not changed.
type CustomAudienceMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of custom audience member, KEYWORD, URL, PLACE_CATEGORY or APP.
	MemberType enums.CustomAudienceMemberTypeEnum_CustomAudienceMemberType `protobuf:"varint,1,opt,name=member_type,json=memberType,proto3,enum=google.ads.googleads.v12.enums.CustomAudienceMemberTypeEnum_CustomAudienceMemberType" json:"member_type,omitempty"`
	// The CustomAudienceMember value. One field is populated depending on the
	// member type.
	//
	// Types that are assignable to Value:
	//	*CustomAudienceMember_Keyword
	//	*CustomAudienceMember_Url
	//	*CustomAudienceMember_PlaceCategory
	//	*CustomAudienceMember_App
	Value isCustomAudienceMember_Value `protobuf_oneof:"value"`
}

func (x *CustomAudienceMember) Reset() {
	*x = CustomAudienceMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_resources_custom_audience_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomAudienceMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomAudienceMember) ProtoMessage() {}

func (x *CustomAudienceMember) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_resources_custom_audience_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomAudienceMember.ProtoReflect.Descriptor instead.
func (*CustomAudienceMember) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_resources_custom_audience_proto_rawDescGZIP(), []int{1}
}

func (x *CustomAudienceMember) GetMemberType() enums.CustomAudienceMemberTypeEnum_CustomAudienceMemberType {
	if x != nil {
		return x.MemberType
	}
	return enums.CustomAudienceMemberTypeEnum_CustomAudienceMemberType(0)
}

func (m *CustomAudienceMember) GetValue() isCustomAudienceMember_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *CustomAudienceMember) GetKeyword() string {
	if x, ok := x.GetValue().(*CustomAudienceMember_Keyword); ok {
		return x.Keyword
	}
	return ""
}

func (x *CustomAudienceMember) GetUrl() string {
	if x, ok := x.GetValue().(*CustomAudienceMember_Url); ok {
		return x.Url
	}
	return ""
}

func (x *CustomAudienceMember) GetPlaceCategory() int64 {
	if x, ok := x.GetValue().(*CustomAudienceMember_PlaceCategory); ok {
		return x.PlaceCategory
	}
	return 0
}

func (x *CustomAudienceMember) GetApp() string {
	if x, ok := x.GetValue().(*CustomAudienceMember_App); ok {
		return x.App
	}
	return ""
}

type isCustomAudienceMember_Value interface {
	isCustomAudienceMember_Value()
}

type CustomAudienceMember_Keyword struct {
	// A keyword or keyword phrase — at most 10 words and 80 characters.
	// Languages with double-width characters such as Chinese, Japanese,
	// or Korean, are allowed 40 characters, which describes the user's
	// interests or actions.
	Keyword string `protobuf:"bytes,2,opt,name=keyword,proto3,oneof"`
}

type CustomAudienceMember_Url struct {
	// An HTTP URL, protocol-included — at most 2048 characters, which includes
	// contents users have interests in.
	Url string `protobuf:"bytes,3,opt,name=url,proto3,oneof"`
}

type CustomAudienceMember_PlaceCategory struct {
	// A place type described by a place category users visit.
	PlaceCategory int64 `protobuf:"varint,4,opt,name=place_category,json=placeCategory,proto3,oneof"`
}

type CustomAudienceMember_App struct {
	// A package name of Android apps which users installed such as
	// com.google.example.
	App string `protobuf:"bytes,5,opt,name=app,proto3,oneof"`
}

func (*CustomAudienceMember_Keyword) isCustomAudienceMember_Value() {}

func (*CustomAudienceMember_Url) isCustomAudienceMember_Value() {}

func (*CustomAudienceMember_PlaceCategory) isCustomAudienceMember_Value() {}

func (*CustomAudienceMember_App) isCustomAudienceMember_Value() {}

var File_google_ads_googleads_v12_resources_custom_audience_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_resources_custom_audience_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x40,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x04, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f,
	0xe0, 0x41, 0x05, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x6a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x3a, 0x6a, 0xea, 0x41, 0x67, 0x0a, 0x27, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x3c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x7d, 0x22, 0x84, 0x02, 0x0a, 0x14, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x76, 0x0a, 0x0b,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x55, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0d,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x03, 0x61, 0x70, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x61, 0x70,
	0x70, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x85, 0x02, 0x0a, 0x26, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64,
	0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_resources_custom_audience_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_resources_custom_audience_proto_rawDescData = file_google_ads_googleads_v12_resources_custom_audience_proto_rawDesc
)

func file_google_ads_googleads_v12_resources_custom_audience_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_resources_custom_audience_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_resources_custom_audience_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_resources_custom_audience_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_resources_custom_audience_proto_rawDescData
}

var file_google_ads_googleads_v12_resources_custom_audience_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v12_resources_custom_audience_proto_goTypes = []interface{}{
	(*CustomAudience)(nil),                                           // 0: google.ads.googleads.v12.resources.CustomAudience
	(*CustomAudienceMember)(nil),                                     // 1: google.ads.googleads.v12.resources.CustomAudienceMember
	(enums.CustomAudienceStatusEnum_CustomAudienceStatus)(0),         // 2: google.ads.googleads.v12.enums.CustomAudienceStatusEnum.CustomAudienceStatus
	(enums.CustomAudienceTypeEnum_CustomAudienceType)(0),             // 3: google.ads.googleads.v12.enums.CustomAudienceTypeEnum.CustomAudienceType
	(enums.CustomAudienceMemberTypeEnum_CustomAudienceMemberType)(0), // 4: google.ads.googleads.v12.enums.CustomAudienceMemberTypeEnum.CustomAudienceMemberType
}
var file_google_ads_googleads_v12_resources_custom_audience_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v12.resources.CustomAudience.status:type_name -> google.ads.googleads.v12.enums.CustomAudienceStatusEnum.CustomAudienceStatus
	3, // 1: google.ads.googleads.v12.resources.CustomAudience.type:type_name -> google.ads.googleads.v12.enums.CustomAudienceTypeEnum.CustomAudienceType
	1, // 2: google.ads.googleads.v12.resources.CustomAudience.members:type_name -> google.ads.googleads.v12.resources.CustomAudienceMember
	4, // 3: google.ads.googleads.v12.resources.CustomAudienceMember.member_type:type_name -> google.ads.googleads.v12.enums.CustomAudienceMemberTypeEnum.CustomAudienceMemberType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_resources_custom_audience_proto_init() }
func file_google_ads_googleads_v12_resources_custom_audience_proto_init() {
	if File_google_ads_googleads_v12_resources_custom_audience_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_resources_custom_audience_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomAudience); i {
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
		file_google_ads_googleads_v12_resources_custom_audience_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomAudienceMember); i {
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
	file_google_ads_googleads_v12_resources_custom_audience_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CustomAudienceMember_Keyword)(nil),
		(*CustomAudienceMember_Url)(nil),
		(*CustomAudienceMember_PlaceCategory)(nil),
		(*CustomAudienceMember_App)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v12_resources_custom_audience_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_resources_custom_audience_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_resources_custom_audience_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v12_resources_custom_audience_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_resources_custom_audience_proto = out.File
	file_google_ads_googleads_v12_resources_custom_audience_proto_rawDesc = nil
	file_google_ads_googleads_v12_resources_custom_audience_proto_goTypes = nil
	file_google_ads_googleads_v12_resources_custom_audience_proto_depIdxs = nil
}
