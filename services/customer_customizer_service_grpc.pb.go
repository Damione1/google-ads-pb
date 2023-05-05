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
// source: google/ads/googleads/v13/services/customer_customizer_service.proto

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
	CustomerCustomizerService_MutateCustomerCustomizers_FullMethodName = "/google.ads.googleads.v13.services.CustomerCustomizerService/MutateCustomerCustomizers"
)

// CustomerCustomizerServiceClient is the client API for CustomerCustomizerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerCustomizerServiceClient interface {
	// Creates, updates or removes customer customizers. Operation statuses are
	// returned.
	MutateCustomerCustomizers(ctx context.Context, in *MutateCustomerCustomizersRequest, opts ...grpc.CallOption) (*MutateCustomerCustomizersResponse, error)
}

type customerCustomizerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerCustomizerServiceClient(cc grpc.ClientConnInterface) CustomerCustomizerServiceClient {
	return &customerCustomizerServiceClient{cc}
}

func (c *customerCustomizerServiceClient) MutateCustomerCustomizers(ctx context.Context, in *MutateCustomerCustomizersRequest, opts ...grpc.CallOption) (*MutateCustomerCustomizersResponse, error) {
	out := new(MutateCustomerCustomizersResponse)
	err := c.cc.Invoke(ctx, CustomerCustomizerService_MutateCustomerCustomizers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerCustomizerServiceServer is the server API for CustomerCustomizerService service.
// All implementations must embed UnimplementedCustomerCustomizerServiceServer
// for forward compatibility
type CustomerCustomizerServiceServer interface {
	// Creates, updates or removes customer customizers. Operation statuses are
	// returned.
	MutateCustomerCustomizers(context.Context, *MutateCustomerCustomizersRequest) (*MutateCustomerCustomizersResponse, error)
	mustEmbedUnimplementedCustomerCustomizerServiceServer()
}

// UnimplementedCustomerCustomizerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerCustomizerServiceServer struct {
}

func (UnimplementedCustomerCustomizerServiceServer) MutateCustomerCustomizers(context.Context, *MutateCustomerCustomizersRequest) (*MutateCustomerCustomizersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerCustomizers not implemented")
}
func (UnimplementedCustomerCustomizerServiceServer) mustEmbedUnimplementedCustomerCustomizerServiceServer() {
}

// UnsafeCustomerCustomizerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerCustomizerServiceServer will
// result in compilation errors.
type UnsafeCustomerCustomizerServiceServer interface {
	mustEmbedUnimplementedCustomerCustomizerServiceServer()
}

func RegisterCustomerCustomizerServiceServer(s grpc.ServiceRegistrar, srv CustomerCustomizerServiceServer) {
	s.RegisterService(&CustomerCustomizerService_ServiceDesc, srv)
}

func _CustomerCustomizerService_MutateCustomerCustomizers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerCustomizersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerCustomizerServiceServer).MutateCustomerCustomizers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerCustomizerService_MutateCustomerCustomizers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerCustomizerServiceServer).MutateCustomerCustomizers(ctx, req.(*MutateCustomerCustomizersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerCustomizerService_ServiceDesc is the grpc.ServiceDesc for CustomerCustomizerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerCustomizerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.CustomerCustomizerService",
	HandlerType: (*CustomerCustomizerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerCustomizers",
			Handler:    _CustomerCustomizerService_MutateCustomerCustomizers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/customer_customizer_service.proto",
}
