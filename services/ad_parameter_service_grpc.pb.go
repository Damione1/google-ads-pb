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
// source: google/ads/googleads/v13/services/ad_parameter_service.proto

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
	AdParameterService_MutateAdParameters_FullMethodName = "/google.ads.googleads.v13.services.AdParameterService/MutateAdParameters"
)

// AdParameterServiceClient is the client API for AdParameterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdParameterServiceClient interface {
	// Creates, updates, or removes ad parameters. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AdParameterError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ContextError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateAdParameters(ctx context.Context, in *MutateAdParametersRequest, opts ...grpc.CallOption) (*MutateAdParametersResponse, error)
}

type adParameterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdParameterServiceClient(cc grpc.ClientConnInterface) AdParameterServiceClient {
	return &adParameterServiceClient{cc}
}

func (c *adParameterServiceClient) MutateAdParameters(ctx context.Context, in *MutateAdParametersRequest, opts ...grpc.CallOption) (*MutateAdParametersResponse, error) {
	out := new(MutateAdParametersResponse)
	err := c.cc.Invoke(ctx, AdParameterService_MutateAdParameters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdParameterServiceServer is the server API for AdParameterService service.
// All implementations must embed UnimplementedAdParameterServiceServer
// for forward compatibility
type AdParameterServiceServer interface {
	// Creates, updates, or removes ad parameters. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AdParameterError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ContextError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateAdParameters(context.Context, *MutateAdParametersRequest) (*MutateAdParametersResponse, error)
	mustEmbedUnimplementedAdParameterServiceServer()
}

// UnimplementedAdParameterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdParameterServiceServer struct {
}

func (UnimplementedAdParameterServiceServer) MutateAdParameters(context.Context, *MutateAdParametersRequest) (*MutateAdParametersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdParameters not implemented")
}
func (UnimplementedAdParameterServiceServer) mustEmbedUnimplementedAdParameterServiceServer() {}

// UnsafeAdParameterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdParameterServiceServer will
// result in compilation errors.
type UnsafeAdParameterServiceServer interface {
	mustEmbedUnimplementedAdParameterServiceServer()
}

func RegisterAdParameterServiceServer(s grpc.ServiceRegistrar, srv AdParameterServiceServer) {
	s.RegisterService(&AdParameterService_ServiceDesc, srv)
}

func _AdParameterService_MutateAdParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdParametersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdParameterServiceServer).MutateAdParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdParameterService_MutateAdParameters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdParameterServiceServer).MutateAdParameters(ctx, req.(*MutateAdParametersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdParameterService_ServiceDesc is the grpc.ServiceDesc for AdParameterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdParameterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.AdParameterService",
	HandlerType: (*AdParameterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdParameters",
			Handler:    _AdParameterService_MutateAdParameters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/ad_parameter_service.proto",
}
