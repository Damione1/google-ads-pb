// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: google/ads/googleads/v12/services/feed_item_target_service.proto

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

// FeedItemTargetServiceClient is the client API for FeedItemTargetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedItemTargetServiceClient interface {
	// Creates or removes feed item targets. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CriterionError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FeedItemTargetError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NotEmptyError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateFeedItemTargets(ctx context.Context, in *MutateFeedItemTargetsRequest, opts ...grpc.CallOption) (*MutateFeedItemTargetsResponse, error)
}

type feedItemTargetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedItemTargetServiceClient(cc grpc.ClientConnInterface) FeedItemTargetServiceClient {
	return &feedItemTargetServiceClient{cc}
}

func (c *feedItemTargetServiceClient) MutateFeedItemTargets(ctx context.Context, in *MutateFeedItemTargetsRequest, opts ...grpc.CallOption) (*MutateFeedItemTargetsResponse, error) {
	out := new(MutateFeedItemTargetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v12.services.FeedItemTargetService/MutateFeedItemTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedItemTargetServiceServer is the server API for FeedItemTargetService service.
// All implementations must embed UnimplementedFeedItemTargetServiceServer
// for forward compatibility
type FeedItemTargetServiceServer interface {
	// Creates or removes feed item targets. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CriterionError]()
	//   [DatabaseError]()
	//   [DistinctError]()
	//   [FeedItemTargetError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NotEmptyError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateFeedItemTargets(context.Context, *MutateFeedItemTargetsRequest) (*MutateFeedItemTargetsResponse, error)
	mustEmbedUnimplementedFeedItemTargetServiceServer()
}

// UnimplementedFeedItemTargetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedItemTargetServiceServer struct {
}

func (UnimplementedFeedItemTargetServiceServer) MutateFeedItemTargets(context.Context, *MutateFeedItemTargetsRequest) (*MutateFeedItemTargetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateFeedItemTargets not implemented")
}
func (UnimplementedFeedItemTargetServiceServer) mustEmbedUnimplementedFeedItemTargetServiceServer() {}

// UnsafeFeedItemTargetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedItemTargetServiceServer will
// result in compilation errors.
type UnsafeFeedItemTargetServiceServer interface {
	mustEmbedUnimplementedFeedItemTargetServiceServer()
}

func RegisterFeedItemTargetServiceServer(s grpc.ServiceRegistrar, srv FeedItemTargetServiceServer) {
	s.RegisterService(&FeedItemTargetService_ServiceDesc, srv)
}

func _FeedItemTargetService_MutateFeedItemTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedItemTargetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemTargetServiceServer).MutateFeedItemTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v12.services.FeedItemTargetService/MutateFeedItemTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemTargetServiceServer).MutateFeedItemTargets(ctx, req.(*MutateFeedItemTargetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedItemTargetService_ServiceDesc is the grpc.ServiceDesc for FeedItemTargetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedItemTargetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v12.services.FeedItemTargetService",
	HandlerType: (*FeedItemTargetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateFeedItemTargets",
			Handler:    _FeedItemTargetService_MutateFeedItemTargets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v12/services/feed_item_target_service.proto",
}
