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
// source: google/ads/googleads/v13/services/campaign_conversion_goal_service.proto

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
	CampaignConversionGoalService_MutateCampaignConversionGoals_FullMethodName = "/google.ads.googleads.v13.services.CampaignConversionGoalService/MutateCampaignConversionGoals"
)

// CampaignConversionGoalServiceClient is the client API for CampaignConversionGoalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignConversionGoalServiceClient interface {
	// Creates, updates or removes campaign conversion goals. Operation statuses
	// are returned.
	MutateCampaignConversionGoals(ctx context.Context, in *MutateCampaignConversionGoalsRequest, opts ...grpc.CallOption) (*MutateCampaignConversionGoalsResponse, error)
}

type campaignConversionGoalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignConversionGoalServiceClient(cc grpc.ClientConnInterface) CampaignConversionGoalServiceClient {
	return &campaignConversionGoalServiceClient{cc}
}

func (c *campaignConversionGoalServiceClient) MutateCampaignConversionGoals(ctx context.Context, in *MutateCampaignConversionGoalsRequest, opts ...grpc.CallOption) (*MutateCampaignConversionGoalsResponse, error) {
	out := new(MutateCampaignConversionGoalsResponse)
	err := c.cc.Invoke(ctx, CampaignConversionGoalService_MutateCampaignConversionGoals_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignConversionGoalServiceServer is the server API for CampaignConversionGoalService service.
// All implementations must embed UnimplementedCampaignConversionGoalServiceServer
// for forward compatibility
type CampaignConversionGoalServiceServer interface {
	// Creates, updates or removes campaign conversion goals. Operation statuses
	// are returned.
	MutateCampaignConversionGoals(context.Context, *MutateCampaignConversionGoalsRequest) (*MutateCampaignConversionGoalsResponse, error)
	mustEmbedUnimplementedCampaignConversionGoalServiceServer()
}

// UnimplementedCampaignConversionGoalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignConversionGoalServiceServer struct {
}

func (UnimplementedCampaignConversionGoalServiceServer) MutateCampaignConversionGoals(context.Context, *MutateCampaignConversionGoalsRequest) (*MutateCampaignConversionGoalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignConversionGoals not implemented")
}
func (UnimplementedCampaignConversionGoalServiceServer) mustEmbedUnimplementedCampaignConversionGoalServiceServer() {
}

// UnsafeCampaignConversionGoalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignConversionGoalServiceServer will
// result in compilation errors.
type UnsafeCampaignConversionGoalServiceServer interface {
	mustEmbedUnimplementedCampaignConversionGoalServiceServer()
}

func RegisterCampaignConversionGoalServiceServer(s grpc.ServiceRegistrar, srv CampaignConversionGoalServiceServer) {
	s.RegisterService(&CampaignConversionGoalService_ServiceDesc, srv)
}

func _CampaignConversionGoalService_MutateCampaignConversionGoals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignConversionGoalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignConversionGoalServiceServer).MutateCampaignConversionGoals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignConversionGoalService_MutateCampaignConversionGoals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignConversionGoalServiceServer).MutateCampaignConversionGoals(ctx, req.(*MutateCampaignConversionGoalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignConversionGoalService_ServiceDesc is the grpc.ServiceDesc for CampaignConversionGoalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignConversionGoalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.CampaignConversionGoalService",
	HandlerType: (*CampaignConversionGoalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignConversionGoals",
			Handler:    _CampaignConversionGoalService_MutateCampaignConversionGoals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/campaign_conversion_goal_service.proto",
}
