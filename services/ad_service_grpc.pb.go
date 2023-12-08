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
// source: google/ads/googleads/v15/services/ad_service.proto

package services

import (
	context "context"
	resources "github.com/Damione1/google-ads-pb/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AdService_GetAd_FullMethodName     = "/google.ads.googleads.v15.services.AdService/GetAd"
	AdService_MutateAds_FullMethodName = "/google.ads.googleads.v15.services.AdService/MutateAds"
)

// AdServiceClient is the client API for AdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdServiceClient interface {
	// Returns the requested ad in full detail.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*resources.Ad, error)
	// Updates ads. Operation statuses are returned. Updating ads is not supported
	// for TextAd, ExpandedDynamicSearchAd, GmailAd and ImageAd.
	//
	// List of thrown errors:
	//
	//	[AdCustomizerError]()
	//	[AdError]()
	//	[AdSharingError]()
	//	[AdxError]()
	//	[AssetError]()
	//	[AssetLinkError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FeedAttributeReferenceError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[FunctionError]()
	//	[FunctionParsingError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[ImageError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MediaBundleError]()
	//	[MediaFileError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[PolicyFindingError]()
	//	[PolicyViolationError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateAds(ctx context.Context, in *MutateAdsRequest, opts ...grpc.CallOption) (*MutateAdsResponse, error)
}

type adServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdServiceClient(cc grpc.ClientConnInterface) AdServiceClient {
	return &adServiceClient{cc}
}

func (c *adServiceClient) GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*resources.Ad, error) {
	out := new(resources.Ad)
	err := c.cc.Invoke(ctx, AdService_GetAd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) MutateAds(ctx context.Context, in *MutateAdsRequest, opts ...grpc.CallOption) (*MutateAdsResponse, error) {
	out := new(MutateAdsResponse)
	err := c.cc.Invoke(ctx, AdService_MutateAds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdServiceServer is the server API for AdService service.
// All implementations must embed UnimplementedAdServiceServer
// for forward compatibility
type AdServiceServer interface {
	// Returns the requested ad in full detail.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GetAd(context.Context, *GetAdRequest) (*resources.Ad, error)
	// Updates ads. Operation statuses are returned. Updating ads is not supported
	// for TextAd, ExpandedDynamicSearchAd, GmailAd and ImageAd.
	//
	// List of thrown errors:
	//
	//	[AdCustomizerError]()
	//	[AdError]()
	//	[AdSharingError]()
	//	[AdxError]()
	//	[AssetError]()
	//	[AssetLinkError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FeedAttributeReferenceError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[FunctionError]()
	//	[FunctionParsingError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[ImageError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MediaBundleError]()
	//	[MediaFileError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[PolicyFindingError]()
	//	[PolicyViolationError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateAds(context.Context, *MutateAdsRequest) (*MutateAdsResponse, error)
	mustEmbedUnimplementedAdServiceServer()
}

// UnimplementedAdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdServiceServer struct {
}

func (UnimplementedAdServiceServer) GetAd(context.Context, *GetAdRequest) (*resources.Ad, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAd not implemented")
}
func (UnimplementedAdServiceServer) MutateAds(context.Context, *MutateAdsRequest) (*MutateAdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAds not implemented")
}
func (UnimplementedAdServiceServer) mustEmbedUnimplementedAdServiceServer() {}

// UnsafeAdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdServiceServer will
// result in compilation errors.
type UnsafeAdServiceServer interface {
	mustEmbedUnimplementedAdServiceServer()
}

func RegisterAdServiceServer(s grpc.ServiceRegistrar, srv AdServiceServer) {
	s.RegisterService(&AdService_ServiceDesc, srv)
}

func _AdService_GetAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).GetAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdService_GetAd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).GetAd(ctx, req.(*GetAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_MutateAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).MutateAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdService_MutateAds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).MutateAds(ctx, req.(*MutateAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdService_ServiceDesc is the grpc.ServiceDesc for AdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.AdService",
	HandlerType: (*AdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAd",
			Handler:    _AdService_GetAd_Handler,
		},
		{
			MethodName: "MutateAds",
			Handler:    _AdService_MutateAds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/ad_service.proto",
}
