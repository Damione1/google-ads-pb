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
// source: google/ads/googleads/v13/services/custom_interest_service.proto

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
	CustomInterestService_MutateCustomInterests_FullMethodName = "/google.ads.googleads.v13.services.CustomInterestService/MutateCustomInterests"
)

// CustomInterestServiceClient is the client API for CustomInterestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomInterestServiceClient interface {
	// Creates or updates custom interests. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CriterionError]()
	//	[CustomInterestError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[PolicyViolationError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[StringLengthError]()
	MutateCustomInterests(ctx context.Context, in *MutateCustomInterestsRequest, opts ...grpc.CallOption) (*MutateCustomInterestsResponse, error)
}

type customInterestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomInterestServiceClient(cc grpc.ClientConnInterface) CustomInterestServiceClient {
	return &customInterestServiceClient{cc}
}

func (c *customInterestServiceClient) MutateCustomInterests(ctx context.Context, in *MutateCustomInterestsRequest, opts ...grpc.CallOption) (*MutateCustomInterestsResponse, error) {
	out := new(MutateCustomInterestsResponse)
	err := c.cc.Invoke(ctx, CustomInterestService_MutateCustomInterests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomInterestServiceServer is the server API for CustomInterestService service.
// All implementations must embed UnimplementedCustomInterestServiceServer
// for forward compatibility
type CustomInterestServiceServer interface {
	// Creates or updates custom interests. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CriterionError]()
	//	[CustomInterestError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[PolicyViolationError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[StringLengthError]()
	MutateCustomInterests(context.Context, *MutateCustomInterestsRequest) (*MutateCustomInterestsResponse, error)
	mustEmbedUnimplementedCustomInterestServiceServer()
}

// UnimplementedCustomInterestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomInterestServiceServer struct {
}

func (UnimplementedCustomInterestServiceServer) MutateCustomInterests(context.Context, *MutateCustomInterestsRequest) (*MutateCustomInterestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomInterests not implemented")
}
func (UnimplementedCustomInterestServiceServer) mustEmbedUnimplementedCustomInterestServiceServer() {}

// UnsafeCustomInterestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomInterestServiceServer will
// result in compilation errors.
type UnsafeCustomInterestServiceServer interface {
	mustEmbedUnimplementedCustomInterestServiceServer()
}

func RegisterCustomInterestServiceServer(s grpc.ServiceRegistrar, srv CustomInterestServiceServer) {
	s.RegisterService(&CustomInterestService_ServiceDesc, srv)
}

func _CustomInterestService_MutateCustomInterests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomInterestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomInterestServiceServer).MutateCustomInterests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomInterestService_MutateCustomInterests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomInterestServiceServer).MutateCustomInterests(ctx, req.(*MutateCustomInterestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomInterestService_ServiceDesc is the grpc.ServiceDesc for CustomInterestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomInterestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.CustomInterestService",
	HandlerType: (*CustomInterestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomInterests",
			Handler:    _CustomInterestService_MutateCustomInterests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/custom_interest_service.proto",
}
