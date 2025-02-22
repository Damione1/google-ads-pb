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
// source: google/ads/googleads/v15/services/feed_item_service.proto

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
	FeedItemService_MutateFeedItems_FullMethodName = "/google.ads.googleads.v15.services.FeedItemService/MutateFeedItems"
)

// FeedItemServiceClient is the client API for FeedItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedItemServiceClient interface {
	// Creates, updates, or removes feed items. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FeedItemError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateFeedItems(ctx context.Context, in *MutateFeedItemsRequest, opts ...grpc.CallOption) (*MutateFeedItemsResponse, error)
}

type feedItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedItemServiceClient(cc grpc.ClientConnInterface) FeedItemServiceClient {
	return &feedItemServiceClient{cc}
}

func (c *feedItemServiceClient) MutateFeedItems(ctx context.Context, in *MutateFeedItemsRequest, opts ...grpc.CallOption) (*MutateFeedItemsResponse, error) {
	out := new(MutateFeedItemsResponse)
	err := c.cc.Invoke(ctx, FeedItemService_MutateFeedItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedItemServiceServer is the server API for FeedItemService service.
// All implementations must embed UnimplementedFeedItemServiceServer
// for forward compatibility
type FeedItemServiceServer interface {
	// Creates, updates, or removes feed items. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FeedItemError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[ListOperationError]()
	//	[MutateError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UrlFieldError]()
	MutateFeedItems(context.Context, *MutateFeedItemsRequest) (*MutateFeedItemsResponse, error)
	mustEmbedUnimplementedFeedItemServiceServer()
}

// UnimplementedFeedItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedItemServiceServer struct {
}

func (UnimplementedFeedItemServiceServer) MutateFeedItems(context.Context, *MutateFeedItemsRequest) (*MutateFeedItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateFeedItems not implemented")
}
func (UnimplementedFeedItemServiceServer) mustEmbedUnimplementedFeedItemServiceServer() {}

// UnsafeFeedItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedItemServiceServer will
// result in compilation errors.
type UnsafeFeedItemServiceServer interface {
	mustEmbedUnimplementedFeedItemServiceServer()
}

func RegisterFeedItemServiceServer(s grpc.ServiceRegistrar, srv FeedItemServiceServer) {
	s.RegisterService(&FeedItemService_ServiceDesc, srv)
}

func _FeedItemService_MutateFeedItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemServiceServer).MutateFeedItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedItemService_MutateFeedItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemServiceServer).MutateFeedItems(ctx, req.(*MutateFeedItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedItemService_ServiceDesc is the grpc.ServiceDesc for FeedItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.FeedItemService",
	HandlerType: (*FeedItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateFeedItems",
			Handler:    _FeedItemService_MutateFeedItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/feed_item_service.proto",
}
