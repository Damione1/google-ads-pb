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
// - protoc             v3.21.9
// source: google/ads/googleads/v13/services/campaign_label_service.proto

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
	CampaignLabelService_MutateCampaignLabels_FullMethodName = "/google.ads.googleads.v13.services.CampaignLabelService/MutateCampaignLabels"
)

// CampaignLabelServiceClient is the client API for CampaignLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignLabelServiceClient interface {
	// Creates and removes campaign-label relationships.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[LabelError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCampaignLabels(ctx context.Context, in *MutateCampaignLabelsRequest, opts ...grpc.CallOption) (*MutateCampaignLabelsResponse, error)
}

type campaignLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignLabelServiceClient(cc grpc.ClientConnInterface) CampaignLabelServiceClient {
	return &campaignLabelServiceClient{cc}
}

func (c *campaignLabelServiceClient) MutateCampaignLabels(ctx context.Context, in *MutateCampaignLabelsRequest, opts ...grpc.CallOption) (*MutateCampaignLabelsResponse, error) {
	out := new(MutateCampaignLabelsResponse)
	err := c.cc.Invoke(ctx, CampaignLabelService_MutateCampaignLabels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignLabelServiceServer is the server API for CampaignLabelService service.
// All implementations must embed UnimplementedCampaignLabelServiceServer
// for forward compatibility
type CampaignLabelServiceServer interface {
	// Creates and removes campaign-label relationships.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[LabelError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCampaignLabels(context.Context, *MutateCampaignLabelsRequest) (*MutateCampaignLabelsResponse, error)
	mustEmbedUnimplementedCampaignLabelServiceServer()
}

// UnimplementedCampaignLabelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignLabelServiceServer struct {
}

func (UnimplementedCampaignLabelServiceServer) MutateCampaignLabels(context.Context, *MutateCampaignLabelsRequest) (*MutateCampaignLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignLabels not implemented")
}
func (UnimplementedCampaignLabelServiceServer) mustEmbedUnimplementedCampaignLabelServiceServer() {}

// UnsafeCampaignLabelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignLabelServiceServer will
// result in compilation errors.
type UnsafeCampaignLabelServiceServer interface {
	mustEmbedUnimplementedCampaignLabelServiceServer()
}

func RegisterCampaignLabelServiceServer(s grpc.ServiceRegistrar, srv CampaignLabelServiceServer) {
	s.RegisterService(&CampaignLabelService_ServiceDesc, srv)
}

func _CampaignLabelService_MutateCampaignLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignLabelServiceServer).MutateCampaignLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignLabelService_MutateCampaignLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignLabelServiceServer).MutateCampaignLabels(ctx, req.(*MutateCampaignLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignLabelService_ServiceDesc is the grpc.ServiceDesc for CampaignLabelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignLabelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.CampaignLabelService",
	HandlerType: (*CampaignLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignLabels",
			Handler:    _CampaignLabelService_MutateCampaignLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/campaign_label_service.proto",
}
