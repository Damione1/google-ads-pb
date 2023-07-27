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
// - protoc             v3.21.12
// source: google/ads/googleads/v14/services/campaign_group_service.proto

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
	CampaignGroupService_MutateCampaignGroups_FullMethodName = "/google.ads.googleads.v14.services.CampaignGroupService/MutateCampaignGroups"
)

// CampaignGroupServiceClient is the client API for CampaignGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignGroupServiceClient interface {
	// Creates, updates, or removes campaign groups. Operation statuses are
	// returned.
	MutateCampaignGroups(ctx context.Context, in *MutateCampaignGroupsRequest, opts ...grpc.CallOption) (*MutateCampaignGroupsResponse, error)
}

type campaignGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignGroupServiceClient(cc grpc.ClientConnInterface) CampaignGroupServiceClient {
	return &campaignGroupServiceClient{cc}
}

func (c *campaignGroupServiceClient) MutateCampaignGroups(ctx context.Context, in *MutateCampaignGroupsRequest, opts ...grpc.CallOption) (*MutateCampaignGroupsResponse, error) {
	out := new(MutateCampaignGroupsResponse)
	err := c.cc.Invoke(ctx, CampaignGroupService_MutateCampaignGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignGroupServiceServer is the server API for CampaignGroupService service.
// All implementations must embed UnimplementedCampaignGroupServiceServer
// for forward compatibility
type CampaignGroupServiceServer interface {
	// Creates, updates, or removes campaign groups. Operation statuses are
	// returned.
	MutateCampaignGroups(context.Context, *MutateCampaignGroupsRequest) (*MutateCampaignGroupsResponse, error)
	mustEmbedUnimplementedCampaignGroupServiceServer()
}

// UnimplementedCampaignGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignGroupServiceServer struct {
}

func (UnimplementedCampaignGroupServiceServer) MutateCampaignGroups(context.Context, *MutateCampaignGroupsRequest) (*MutateCampaignGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignGroups not implemented")
}
func (UnimplementedCampaignGroupServiceServer) mustEmbedUnimplementedCampaignGroupServiceServer() {}

// UnsafeCampaignGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignGroupServiceServer will
// result in compilation errors.
type UnsafeCampaignGroupServiceServer interface {
	mustEmbedUnimplementedCampaignGroupServiceServer()
}

func RegisterCampaignGroupServiceServer(s grpc.ServiceRegistrar, srv CampaignGroupServiceServer) {
	s.RegisterService(&CampaignGroupService_ServiceDesc, srv)
}

func _CampaignGroupService_MutateCampaignGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignGroupServiceServer).MutateCampaignGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignGroupService_MutateCampaignGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignGroupServiceServer).MutateCampaignGroups(ctx, req.(*MutateCampaignGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignGroupService_ServiceDesc is the grpc.ServiceDesc for CampaignGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v14.services.CampaignGroupService",
	HandlerType: (*CampaignGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignGroups",
			Handler:    _CampaignGroupService_MutateCampaignGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v14/services/campaign_group_service.proto",
}
