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
// source: google/ads/googleads/v13/services/google_ads_field_service.proto

package services

import (
	context "context"
	resources "github.com/shenzhencenter/google-ads-pb/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GoogleAdsFieldService_GetGoogleAdsField_FullMethodName     = "/google.ads.googleads.v13.services.GoogleAdsFieldService/GetGoogleAdsField"
	GoogleAdsFieldService_SearchGoogleAdsFields_FullMethodName = "/google.ads.googleads.v13.services.GoogleAdsFieldService/SearchGoogleAdsFields"
)

// GoogleAdsFieldServiceClient is the client API for GoogleAdsFieldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoogleAdsFieldServiceClient interface {
	// Returns just the requested field.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GetGoogleAdsField(ctx context.Context, in *GetGoogleAdsFieldRequest, opts ...grpc.CallOption) (*resources.GoogleAdsField, error)
	// Returns all fields that match the search query.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QueryError]()
	//	[QuotaError]()
	//	[RequestError]()
	SearchGoogleAdsFields(ctx context.Context, in *SearchGoogleAdsFieldsRequest, opts ...grpc.CallOption) (*SearchGoogleAdsFieldsResponse, error)
}

type googleAdsFieldServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoogleAdsFieldServiceClient(cc grpc.ClientConnInterface) GoogleAdsFieldServiceClient {
	return &googleAdsFieldServiceClient{cc}
}

func (c *googleAdsFieldServiceClient) GetGoogleAdsField(ctx context.Context, in *GetGoogleAdsFieldRequest, opts ...grpc.CallOption) (*resources.GoogleAdsField, error) {
	out := new(resources.GoogleAdsField)
	err := c.cc.Invoke(ctx, GoogleAdsFieldService_GetGoogleAdsField_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleAdsFieldServiceClient) SearchGoogleAdsFields(ctx context.Context, in *SearchGoogleAdsFieldsRequest, opts ...grpc.CallOption) (*SearchGoogleAdsFieldsResponse, error) {
	out := new(SearchGoogleAdsFieldsResponse)
	err := c.cc.Invoke(ctx, GoogleAdsFieldService_SearchGoogleAdsFields_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoogleAdsFieldServiceServer is the server API for GoogleAdsFieldService service.
// All implementations must embed UnimplementedGoogleAdsFieldServiceServer
// for forward compatibility
type GoogleAdsFieldServiceServer interface {
	// Returns just the requested field.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GetGoogleAdsField(context.Context, *GetGoogleAdsFieldRequest) (*resources.GoogleAdsField, error)
	// Returns all fields that match the search query.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QueryError]()
	//	[QuotaError]()
	//	[RequestError]()
	SearchGoogleAdsFields(context.Context, *SearchGoogleAdsFieldsRequest) (*SearchGoogleAdsFieldsResponse, error)
	mustEmbedUnimplementedGoogleAdsFieldServiceServer()
}

// UnimplementedGoogleAdsFieldServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoogleAdsFieldServiceServer struct {
}

func (UnimplementedGoogleAdsFieldServiceServer) GetGoogleAdsField(context.Context, *GetGoogleAdsFieldRequest) (*resources.GoogleAdsField, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoogleAdsField not implemented")
}
func (UnimplementedGoogleAdsFieldServiceServer) SearchGoogleAdsFields(context.Context, *SearchGoogleAdsFieldsRequest) (*SearchGoogleAdsFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGoogleAdsFields not implemented")
}
func (UnimplementedGoogleAdsFieldServiceServer) mustEmbedUnimplementedGoogleAdsFieldServiceServer() {}

// UnsafeGoogleAdsFieldServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoogleAdsFieldServiceServer will
// result in compilation errors.
type UnsafeGoogleAdsFieldServiceServer interface {
	mustEmbedUnimplementedGoogleAdsFieldServiceServer()
}

func RegisterGoogleAdsFieldServiceServer(s grpc.ServiceRegistrar, srv GoogleAdsFieldServiceServer) {
	s.RegisterService(&GoogleAdsFieldService_ServiceDesc, srv)
}

func _GoogleAdsFieldService_GetGoogleAdsField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoogleAdsFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleAdsFieldServiceServer).GetGoogleAdsField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoogleAdsFieldService_GetGoogleAdsField_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleAdsFieldServiceServer).GetGoogleAdsField(ctx, req.(*GetGoogleAdsFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleAdsFieldService_SearchGoogleAdsFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGoogleAdsFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleAdsFieldServiceServer).SearchGoogleAdsFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoogleAdsFieldService_SearchGoogleAdsFields_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleAdsFieldServiceServer).SearchGoogleAdsFields(ctx, req.(*SearchGoogleAdsFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoogleAdsFieldService_ServiceDesc is the grpc.ServiceDesc for GoogleAdsFieldService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoogleAdsFieldService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v13.services.GoogleAdsFieldService",
	HandlerType: (*GoogleAdsFieldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoogleAdsField",
			Handler:    _GoogleAdsFieldService_GetGoogleAdsField_Handler,
		},
		{
			MethodName: "SearchGoogleAdsFields",
			Handler:    _GoogleAdsFieldService_SearchGoogleAdsFields_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v13/services/google_ads_field_service.proto",
}
