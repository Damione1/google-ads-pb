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
// source: google/ads/googleads/v15/services/extension_feed_item_service.proto

package services

import (
	enums "github.com/Damione1/google-ads-pb/enums"
	resources "github.com/Damione1/google-ads-pb/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for
// [ExtensionFeedItemService.MutateExtensionFeedItems][google.ads.googleads.v15.services.ExtensionFeedItemService.MutateExtensionFeedItems].
type MutateExtensionFeedItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose extension feed items are being
	// modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual extension feed
	// items.
	Operations []*ExtensionFeedItemOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,5,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v15.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
}

func (x *MutateExtensionFeedItemsRequest) Reset() {
	*x = MutateExtensionFeedItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateExtensionFeedItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateExtensionFeedItemsRequest) ProtoMessage() {}

func (x *MutateExtensionFeedItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateExtensionFeedItemsRequest.ProtoReflect.Descriptor instead.
func (*MutateExtensionFeedItemsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateExtensionFeedItemsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateExtensionFeedItemsRequest) GetOperations() []*ExtensionFeedItemOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateExtensionFeedItemsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateExtensionFeedItemsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateExtensionFeedItemsRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, update, remove) on an extension feed item.
type ExtensionFeedItemOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//
	//	*ExtensionFeedItemOperation_Create
	//	*ExtensionFeedItemOperation_Update
	//	*ExtensionFeedItemOperation_Remove
	Operation isExtensionFeedItemOperation_Operation `protobuf_oneof:"operation"`
}

func (x *ExtensionFeedItemOperation) Reset() {
	*x = ExtensionFeedItemOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionFeedItemOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionFeedItemOperation) ProtoMessage() {}

func (x *ExtensionFeedItemOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionFeedItemOperation.ProtoReflect.Descriptor instead.
func (*ExtensionFeedItemOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescGZIP(), []int{1}
}

func (x *ExtensionFeedItemOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *ExtensionFeedItemOperation) GetOperation() isExtensionFeedItemOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *ExtensionFeedItemOperation) GetCreate() *resources.ExtensionFeedItem {
	if x, ok := x.GetOperation().(*ExtensionFeedItemOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *ExtensionFeedItemOperation) GetUpdate() *resources.ExtensionFeedItem {
	if x, ok := x.GetOperation().(*ExtensionFeedItemOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *ExtensionFeedItemOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*ExtensionFeedItemOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isExtensionFeedItemOperation_Operation interface {
	isExtensionFeedItemOperation_Operation()
}

type ExtensionFeedItemOperation_Create struct {
	// Create operation: No resource name is expected for the new extension
	// feed item.
	Create *resources.ExtensionFeedItem `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type ExtensionFeedItemOperation_Update struct {
	// Update operation: The extension feed item is expected to have a
	// valid resource name.
	Update *resources.ExtensionFeedItem `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type ExtensionFeedItemOperation_Remove struct {
	// Remove operation: A resource name for the removed extension feed item
	// is expected, in this format:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*ExtensionFeedItemOperation_Create) isExtensionFeedItemOperation_Operation() {}

func (*ExtensionFeedItemOperation_Update) isExtensionFeedItemOperation_Operation() {}

func (*ExtensionFeedItemOperation_Remove) isExtensionFeedItemOperation_Operation() {}

// Response message for an extension feed item mutate.
type MutateExtensionFeedItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateExtensionFeedItemResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateExtensionFeedItemsResponse) Reset() {
	*x = MutateExtensionFeedItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateExtensionFeedItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateExtensionFeedItemsResponse) ProtoMessage() {}

func (x *MutateExtensionFeedItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateExtensionFeedItemsResponse.ProtoReflect.Descriptor instead.
func (*MutateExtensionFeedItemsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateExtensionFeedItemsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateExtensionFeedItemsResponse) GetResults() []*MutateExtensionFeedItemResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the extension feed item mutate.
type MutateExtensionFeedItemResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated extension feed item with only mutable fields after mutate. The
	// field will only be returned when response_content_type is set to
	// "MUTABLE_RESOURCE".
	ExtensionFeedItem *resources.ExtensionFeedItem `protobuf:"bytes,2,opt,name=extension_feed_item,json=extensionFeedItem,proto3" json:"extension_feed_item,omitempty"`
}

func (x *MutateExtensionFeedItemResult) Reset() {
	*x = MutateExtensionFeedItemResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateExtensionFeedItemResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateExtensionFeedItemResult) ProtoMessage() {}

func (x *MutateExtensionFeedItemResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateExtensionFeedItemResult.ProtoReflect.Descriptor instead.
func (*MutateExtensionFeedItemResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateExtensionFeedItemResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateExtensionFeedItemResult) GetExtensionFeedItem() *resources.ExtensionFeedItem {
	if x != nil {
		return x.ExtensionFeedItem
	}
	return nil
}

var File_google_ads_googleads_v15_services_extension_feed_item_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65,
	0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfa, 0x02, 0x0a, 0x1f, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x62, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x7f, 0x0a, 0x15,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd3, 0x02,
	0x0a, 0x1a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x4f, 0x0a, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xfa, 0x41, 0x2c,
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x00, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xc6, 0x01, 0x0a, 0x20, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x5a, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xdc, 0x01, 0x0a,
	0x1d, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x54,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65,
	0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x65, 0x0a, 0x13, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x32, 0xe5, 0x02, 0x0a, 0x18,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x02, 0x0a, 0x18, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c,
	0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x3a,
	0x01, 0x2a, 0x22, 0x38, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d,
	0x2a, 0x7d, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x45, 0xca, 0x41,
	0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x42, 0x89, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1d, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x35, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescData = file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDesc
)

func file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDescData
}

var file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v15_services_extension_feed_item_service_proto_goTypes = []interface{}{
	(*MutateExtensionFeedItemsRequest)(nil),                // 0: google.ads.googleads.v15.services.MutateExtensionFeedItemsRequest
	(*ExtensionFeedItemOperation)(nil),                     // 1: google.ads.googleads.v15.services.ExtensionFeedItemOperation
	(*MutateExtensionFeedItemsResponse)(nil),               // 2: google.ads.googleads.v15.services.MutateExtensionFeedItemsResponse
	(*MutateExtensionFeedItemResult)(nil),                  // 3: google.ads.googleads.v15.services.MutateExtensionFeedItemResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v15.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 5: google.protobuf.FieldMask
	(*resources.ExtensionFeedItem)(nil),                    // 6: google.ads.googleads.v15.resources.ExtensionFeedItem
	(*status.Status)(nil),                                  // 7: google.rpc.Status
}
var file_google_ads_googleads_v15_services_extension_feed_item_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v15.services.MutateExtensionFeedItemsRequest.operations:type_name -> google.ads.googleads.v15.services.ExtensionFeedItemOperation
	4, // 1: google.ads.googleads.v15.services.MutateExtensionFeedItemsRequest.response_content_type:type_name -> google.ads.googleads.v15.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v15.services.ExtensionFeedItemOperation.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.ads.googleads.v15.services.ExtensionFeedItemOperation.create:type_name -> google.ads.googleads.v15.resources.ExtensionFeedItem
	6, // 4: google.ads.googleads.v15.services.ExtensionFeedItemOperation.update:type_name -> google.ads.googleads.v15.resources.ExtensionFeedItem
	7, // 5: google.ads.googleads.v15.services.MutateExtensionFeedItemsResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 6: google.ads.googleads.v15.services.MutateExtensionFeedItemsResponse.results:type_name -> google.ads.googleads.v15.services.MutateExtensionFeedItemResult
	6, // 7: google.ads.googleads.v15.services.MutateExtensionFeedItemResult.extension_feed_item:type_name -> google.ads.googleads.v15.resources.ExtensionFeedItem
	0, // 8: google.ads.googleads.v15.services.ExtensionFeedItemService.MutateExtensionFeedItems:input_type -> google.ads.googleads.v15.services.MutateExtensionFeedItemsRequest
	2, // 9: google.ads.googleads.v15.services.ExtensionFeedItemService.MutateExtensionFeedItems:output_type -> google.ads.googleads.v15.services.MutateExtensionFeedItemsResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v15_services_extension_feed_item_service_proto_init() }
func file_google_ads_googleads_v15_services_extension_feed_item_service_proto_init() {
	if File_google_ads_googleads_v15_services_extension_feed_item_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateExtensionFeedItemsRequest); i {
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
		file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionFeedItemOperation); i {
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
		file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateExtensionFeedItemsResponse); i {
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
		file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateExtensionFeedItemResult); i {
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
	file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ExtensionFeedItemOperation_Create)(nil),
		(*ExtensionFeedItemOperation_Update)(nil),
		(*ExtensionFeedItemOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v15_services_extension_feed_item_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_services_extension_feed_item_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v15_services_extension_feed_item_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_services_extension_feed_item_service_proto = out.File
	file_google_ads_googleads_v15_services_extension_feed_item_service_proto_rawDesc = nil
	file_google_ads_googleads_v15_services_extension_feed_item_service_proto_goTypes = nil
	file_google_ads_googleads_v15_services_extension_feed_item_service_proto_depIdxs = nil
}
