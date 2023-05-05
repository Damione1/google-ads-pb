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
// source: google/ads/googleads/v13/services/bidding_data_exclusion_service.proto

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
	BiddingDataExclusionService_MutateBiddingDataExclusions_FullMethodName = "/google.ads.googleads.v13.services.BiddingDataExclusionService/MutateBiddingDataExclusions"
)

// BiddingDataExclusionServiceClient is the client API for BiddingDataExclusionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BiddingDataExclusionServiceClient interface {
	// Creates, updates, or removes data exclusions.
	// Operation statuses are returned.
	MutateBiddingDataExclusions(ctx context.Context, in *MutateBiddingDataExclusionsRequest, opts ...grpc.CallOption) (*MutateBiddingDataExclusionsResponse, error)
}

type biddingDataExclusionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBiddingDataExclusionServiceClient(cc grpc.ClientConnInterface) BiddingDataExclusionServiceClient {
	return &biddingDataExclusionServiceClient{cc}
}

func (c *biddingDataExclusionServiceClient) MutateBiddingDataExclusions(ctx context.Context, in *MutateBiddingDataExclusionsRequest, opts ...grpc.CallOption) (*MutateBiddingDataExclusionsResponse, error) {
	out := new(MutateBiddingDataExclusionsResponse)
	err := c.cc.Invoke(ctx, BiddingDataExclusionService_MutateBiddingDataExclusions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiddingDataExclusionServiceServer is the server API for BiddingDataExclusionService service.
// All implementations must embed UnimplementedBiddingDataExclusionServiceServer
// for forward compatibility
type BiddingDataExclusionServiceServer interface {
	// Creates, updates, or removes data exclusions.
	// Operation statuses are returned.
	MutateBiddingDataExclusions(context.Context, *MutateBiddingDataExclusionsRequest) (*MutateBiddingDataExclusionsResponse, error)
	mustEmbedUnimplementedBiddingDataExclusionServiceServer()
}

// UnimplementedBiddingDataExclusionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBiddingDataExclusionServiceServer struct {
}

func (UnimplementedBiddingDataExclusionServiceServer) MutateBiddingDataExclusions(context.Context, *MutateBiddingDataExclusionsRequest) (*MutateBiddingDataExclusionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBiddingDataExclusions not implemented")
}
func (UnimplementedBiddingDataExclusionServiceServer) mustEmbedUnimplementedBiddingDataExclusionServiceServer() {
}

// UnsafeBiddingDataExclusionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BiddingDataExclusionServiceServer will
// result in compilation errors.
type UnsafeBiddingDataExclusionServiceServer interface {
	mustEmbedUnimplementedBiddingDataExclusionServiceServer()
}

func RegisterBiddingDataExclusionServiceServer(s grpc.ServiceRegistrar, srv BiddingDataExclusionServiceServer) {
	s.RegisterService(&BiddingDataExclusionService_ServiceDesc, srv)
}

func _BiddingDataExclusionService_MutateBiddingDataExclusions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBiddingDataExclusionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingDataExclusionServiceServer).MutateBiddingDataExclusions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BiddingDataExclusionService_MutateBiddingDataExclusions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingDataExclusionServiceServer).MutateBiddingDataExclusions(ctx, req.(*MutateBiddingDataExclusionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BiddingDataExclusionService_ServiceDesc is the grpc.ServiceDesc for BiddingDataExclusionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BiddingDataExclusionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.BiddingDataExclusionService",
	HandlerType: (*BiddingDataExclusionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateBiddingDataExclusions",
			Handler:    _BiddingDataExclusionService_MutateBiddingDataExclusions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/bidding_data_exclusion_service.proto",
}
