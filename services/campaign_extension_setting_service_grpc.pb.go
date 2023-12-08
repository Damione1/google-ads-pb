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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: google/ads/googleads/v15/services/campaign_extension_setting_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CampaignExtensionSettingService_MutateCampaignExtensionSettings_FullMethodName = "/google.ads.googleads.v15.services.CampaignExtensionSettingService/MutateCampaignExtensionSettings"
)

// CampaignExtensionSettingServiceClient is the client API for CampaignExtensionSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignExtensionSettingServiceClient interface {
	// Creates, updates, or removes campaign extension settings. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[ExtensionSettingError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateCampaignExtensionSettings(ctx context.Context, in *MutateCampaignExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCampaignExtensionSettingsResponse, error)
}

type campaignExtensionSettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignExtensionSettingServiceClient(cc grpc.ClientConnInterface) CampaignExtensionSettingServiceClient {
	return &campaignExtensionSettingServiceClient{cc}
}

func (c *campaignExtensionSettingServiceClient) MutateCampaignExtensionSettings(ctx context.Context, in *MutateCampaignExtensionSettingsRequest, opts ...grpc.CallOption) (*MutateCampaignExtensionSettingsResponse, error) {
	out := new(MutateCampaignExtensionSettingsResponse)
	err := c.cc.Invoke(ctx, CampaignExtensionSettingService_MutateCampaignExtensionSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignExtensionSettingServiceServer is the server API for CampaignExtensionSettingService service.
// All implementations must embed UnimplementedCampaignExtensionSettingServiceServer
// for forward compatibility
type CampaignExtensionSettingServiceServer interface {
	// Creates, updates, or removes campaign extension settings. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[ExtensionSettingError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateCampaignExtensionSettings(context.Context, *MutateCampaignExtensionSettingsRequest) (*MutateCampaignExtensionSettingsResponse, error)
	mustEmbedUnimplementedCampaignExtensionSettingServiceServer()
}

// UnimplementedCampaignExtensionSettingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignExtensionSettingServiceServer struct {
}

func (UnimplementedCampaignExtensionSettingServiceServer) MutateCampaignExtensionSettings(context.Context, *MutateCampaignExtensionSettingsRequest) (*MutateCampaignExtensionSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignExtensionSettings not implemented")
}
func (UnimplementedCampaignExtensionSettingServiceServer) mustEmbedUnimplementedCampaignExtensionSettingServiceServer() {
}

// UnsafeCampaignExtensionSettingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignExtensionSettingServiceServer will
// result in compilation errors.
type UnsafeCampaignExtensionSettingServiceServer interface {
	mustEmbedUnimplementedCampaignExtensionSettingServiceServer()
}

func RegisterCampaignExtensionSettingServiceServer(s grpc.ServiceRegistrar, srv CampaignExtensionSettingServiceServer) {
	s.RegisterService(&CampaignExtensionSettingService_ServiceDesc, srv)
}

func _CampaignExtensionSettingService_MutateCampaignExtensionSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignExtensionSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignExtensionSettingServiceServer).MutateCampaignExtensionSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignExtensionSettingService_MutateCampaignExtensionSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignExtensionSettingServiceServer).MutateCampaignExtensionSettings(ctx, req.(*MutateCampaignExtensionSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignExtensionSettingService_ServiceDesc is the grpc.ServiceDesc for CampaignExtensionSettingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignExtensionSettingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.CampaignExtensionSettingService",
	HandlerType: (*CampaignExtensionSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignExtensionSettings",
			Handler:    _CampaignExtensionSettingService_MutateCampaignExtensionSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/campaign_extension_setting_service.proto",
}
