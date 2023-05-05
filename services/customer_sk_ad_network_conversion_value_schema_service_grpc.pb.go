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
// source: google/ads/googleads/v13/services/customer_sk_ad_network_conversion_value_schema_service.proto

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
	CustomerSkAdNetworkConversionValueSchemaService_MutateCustomerSkAdNetworkConversionValueSchema_FullMethodName = "/google.ads.googleads.v13.services.CustomerSkAdNetworkConversionValueSchemaService/MutateCustomerSkAdNetworkConversionValueSchema"
)

// CustomerSkAdNetworkConversionValueSchemaServiceClient is the client API for CustomerSkAdNetworkConversionValueSchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerSkAdNetworkConversionValueSchemaServiceClient interface {
	// Creates or updates the CustomerSkAdNetworkConversionValueSchema.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[InternalError]()
	//	[MutateError]()
	MutateCustomerSkAdNetworkConversionValueSchema(ctx context.Context, in *MutateCustomerSkAdNetworkConversionValueSchemaRequest, opts ...grpc.CallOption) (*MutateCustomerSkAdNetworkConversionValueSchemaResponse, error)
}

type customerSkAdNetworkConversionValueSchemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerSkAdNetworkConversionValueSchemaServiceClient(cc grpc.ClientConnInterface) CustomerSkAdNetworkConversionValueSchemaServiceClient {
	return &customerSkAdNetworkConversionValueSchemaServiceClient{cc}
}

func (c *customerSkAdNetworkConversionValueSchemaServiceClient) MutateCustomerSkAdNetworkConversionValueSchema(ctx context.Context, in *MutateCustomerSkAdNetworkConversionValueSchemaRequest, opts ...grpc.CallOption) (*MutateCustomerSkAdNetworkConversionValueSchemaResponse, error) {
	out := new(MutateCustomerSkAdNetworkConversionValueSchemaResponse)
	err := c.cc.Invoke(ctx, CustomerSkAdNetworkConversionValueSchemaService_MutateCustomerSkAdNetworkConversionValueSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerSkAdNetworkConversionValueSchemaServiceServer is the server API for CustomerSkAdNetworkConversionValueSchemaService service.
// All implementations must embed UnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer
// for forward compatibility
type CustomerSkAdNetworkConversionValueSchemaServiceServer interface {
	// Creates or updates the CustomerSkAdNetworkConversionValueSchema.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[InternalError]()
	//	[MutateError]()
	MutateCustomerSkAdNetworkConversionValueSchema(context.Context, *MutateCustomerSkAdNetworkConversionValueSchemaRequest) (*MutateCustomerSkAdNetworkConversionValueSchemaResponse, error)
	mustEmbedUnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer()
}

// UnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer struct {
}

func (UnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer) MutateCustomerSkAdNetworkConversionValueSchema(context.Context, *MutateCustomerSkAdNetworkConversionValueSchemaRequest) (*MutateCustomerSkAdNetworkConversionValueSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerSkAdNetworkConversionValueSchema not implemented")
}
func (UnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer) mustEmbedUnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer() {
}

// UnsafeCustomerSkAdNetworkConversionValueSchemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerSkAdNetworkConversionValueSchemaServiceServer will
// result in compilation errors.
type UnsafeCustomerSkAdNetworkConversionValueSchemaServiceServer interface {
	mustEmbedUnimplementedCustomerSkAdNetworkConversionValueSchemaServiceServer()
}

func RegisterCustomerSkAdNetworkConversionValueSchemaServiceServer(s grpc.ServiceRegistrar, srv CustomerSkAdNetworkConversionValueSchemaServiceServer) {
	s.RegisterService(&CustomerSkAdNetworkConversionValueSchemaService_ServiceDesc, srv)
}

func _CustomerSkAdNetworkConversionValueSchemaService_MutateCustomerSkAdNetworkConversionValueSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerSkAdNetworkConversionValueSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerSkAdNetworkConversionValueSchemaServiceServer).MutateCustomerSkAdNetworkConversionValueSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerSkAdNetworkConversionValueSchemaService_MutateCustomerSkAdNetworkConversionValueSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerSkAdNetworkConversionValueSchemaServiceServer).MutateCustomerSkAdNetworkConversionValueSchema(ctx, req.(*MutateCustomerSkAdNetworkConversionValueSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerSkAdNetworkConversionValueSchemaService_ServiceDesc is the grpc.ServiceDesc for CustomerSkAdNetworkConversionValueSchemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerSkAdNetworkConversionValueSchemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.CustomerSkAdNetworkConversionValueSchemaService",
	HandlerType: (*CustomerSkAdNetworkConversionValueSchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerSkAdNetworkConversionValueSchema",
			Handler:    _CustomerSkAdNetworkConversionValueSchemaService_MutateCustomerSkAdNetworkConversionValueSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/customer_sk_ad_network_conversion_value_schema_service.proto",
}
