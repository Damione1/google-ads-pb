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
// source: google/ads/googleads/v15/services/campaign_bid_modifier_service.proto

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
	CampaignBidModifierService_MutateCampaignBidModifiers_FullMethodName = "/google.ads.googleads.v15.services.CampaignBidModifierService/MutateCampaignBidModifiers"
)

// CampaignBidModifierServiceClient is the client API for CampaignBidModifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignBidModifierServiceClient interface {
	// Creates, updates, or removes campaign bid modifiers.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ContextError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateCampaignBidModifiers(ctx context.Context, in *MutateCampaignBidModifiersRequest, opts ...grpc.CallOption) (*MutateCampaignBidModifiersResponse, error)
}

type campaignBidModifierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignBidModifierServiceClient(cc grpc.ClientConnInterface) CampaignBidModifierServiceClient {
	return &campaignBidModifierServiceClient{cc}
}

func (c *campaignBidModifierServiceClient) MutateCampaignBidModifiers(ctx context.Context, in *MutateCampaignBidModifiersRequest, opts ...grpc.CallOption) (*MutateCampaignBidModifiersResponse, error) {
	out := new(MutateCampaignBidModifiersResponse)
	err := c.cc.Invoke(ctx, CampaignBidModifierService_MutateCampaignBidModifiers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignBidModifierServiceServer is the server API for CampaignBidModifierService service.
// All implementations must embed UnimplementedCampaignBidModifierServiceServer
// for forward compatibility
type CampaignBidModifierServiceServer interface {
	// Creates, updates, or removes campaign bid modifiers.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ContextError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateCampaignBidModifiers(context.Context, *MutateCampaignBidModifiersRequest) (*MutateCampaignBidModifiersResponse, error)
	mustEmbedUnimplementedCampaignBidModifierServiceServer()
}

// UnimplementedCampaignBidModifierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignBidModifierServiceServer struct {
}

func (UnimplementedCampaignBidModifierServiceServer) MutateCampaignBidModifiers(context.Context, *MutateCampaignBidModifiersRequest) (*MutateCampaignBidModifiersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignBidModifiers not implemented")
}
func (UnimplementedCampaignBidModifierServiceServer) mustEmbedUnimplementedCampaignBidModifierServiceServer() {
}

// UnsafeCampaignBidModifierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignBidModifierServiceServer will
// result in compilation errors.
type UnsafeCampaignBidModifierServiceServer interface {
	mustEmbedUnimplementedCampaignBidModifierServiceServer()
}

func RegisterCampaignBidModifierServiceServer(s grpc.ServiceRegistrar, srv CampaignBidModifierServiceServer) {
	s.RegisterService(&CampaignBidModifierService_ServiceDesc, srv)
}

func _CampaignBidModifierService_MutateCampaignBidModifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignBidModifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBidModifierServiceServer).MutateCampaignBidModifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignBidModifierService_MutateCampaignBidModifiers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBidModifierServiceServer).MutateCampaignBidModifiers(ctx, req.(*MutateCampaignBidModifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignBidModifierService_ServiceDesc is the grpc.ServiceDesc for CampaignBidModifierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignBidModifierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.CampaignBidModifierService",
	HandlerType: (*CampaignBidModifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignBidModifiers",
			Handler:    _CampaignBidModifierService_MutateCampaignBidModifiers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/campaign_bid_modifier_service.proto",
}
