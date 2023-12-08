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
// source: google/ads/googleads/v15/services/feed_item_set_link_service.proto

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
	FeedItemSetLinkService_MutateFeedItemSetLinks_FullMethodName = "/google.ads.googleads.v15.services.FeedItemSetLinkService/MutateFeedItemSetLinks"
)

// FeedItemSetLinkServiceClient is the client API for FeedItemSetLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedItemSetLinkServiceClient interface {
	// Creates, updates, or removes feed item set links.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateFeedItemSetLinks(ctx context.Context, in *MutateFeedItemSetLinksRequest, opts ...grpc.CallOption) (*MutateFeedItemSetLinksResponse, error)
}

type feedItemSetLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedItemSetLinkServiceClient(cc grpc.ClientConnInterface) FeedItemSetLinkServiceClient {
	return &feedItemSetLinkServiceClient{cc}
}

func (c *feedItemSetLinkServiceClient) MutateFeedItemSetLinks(ctx context.Context, in *MutateFeedItemSetLinksRequest, opts ...grpc.CallOption) (*MutateFeedItemSetLinksResponse, error) {
	out := new(MutateFeedItemSetLinksResponse)
	err := c.cc.Invoke(ctx, FeedItemSetLinkService_MutateFeedItemSetLinks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedItemSetLinkServiceServer is the server API for FeedItemSetLinkService service.
// All implementations must embed UnimplementedFeedItemSetLinkServiceServer
// for forward compatibility
type FeedItemSetLinkServiceServer interface {
	// Creates, updates, or removes feed item set links.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateFeedItemSetLinks(context.Context, *MutateFeedItemSetLinksRequest) (*MutateFeedItemSetLinksResponse, error)
	mustEmbedUnimplementedFeedItemSetLinkServiceServer()
}

// UnimplementedFeedItemSetLinkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedItemSetLinkServiceServer struct {
}

func (UnimplementedFeedItemSetLinkServiceServer) MutateFeedItemSetLinks(context.Context, *MutateFeedItemSetLinksRequest) (*MutateFeedItemSetLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateFeedItemSetLinks not implemented")
}
func (UnimplementedFeedItemSetLinkServiceServer) mustEmbedUnimplementedFeedItemSetLinkServiceServer() {
}

// UnsafeFeedItemSetLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedItemSetLinkServiceServer will
// result in compilation errors.
type UnsafeFeedItemSetLinkServiceServer interface {
	mustEmbedUnimplementedFeedItemSetLinkServiceServer()
}

func RegisterFeedItemSetLinkServiceServer(s grpc.ServiceRegistrar, srv FeedItemSetLinkServiceServer) {
	s.RegisterService(&FeedItemSetLinkService_ServiceDesc, srv)
}

func _FeedItemSetLinkService_MutateFeedItemSetLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedItemSetLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemSetLinkServiceServer).MutateFeedItemSetLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeedItemSetLinkService_MutateFeedItemSetLinks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemSetLinkServiceServer).MutateFeedItemSetLinks(ctx, req.(*MutateFeedItemSetLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedItemSetLinkService_ServiceDesc is the grpc.ServiceDesc for FeedItemSetLinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedItemSetLinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.FeedItemSetLinkService",
	HandlerType: (*FeedItemSetLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateFeedItemSetLinks",
			Handler:    _FeedItemSetLinkService_MutateFeedItemSetLinks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/feed_item_set_link_service.proto",
}
