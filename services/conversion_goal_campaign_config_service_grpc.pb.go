// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: google/ads/googleads/v12/services/conversion_goal_campaign_config_service.proto

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

// ConversionGoalCampaignConfigServiceClient is the client API for ConversionGoalCampaignConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversionGoalCampaignConfigServiceClient interface {
	// Creates, updates or removes conversion goal campaign config. Operation
	// statuses are returned.
	MutateConversionGoalCampaignConfigs(ctx context.Context, in *MutateConversionGoalCampaignConfigsRequest, opts ...grpc.CallOption) (*MutateConversionGoalCampaignConfigsResponse, error)
}

type conversionGoalCampaignConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionGoalCampaignConfigServiceClient(cc grpc.ClientConnInterface) ConversionGoalCampaignConfigServiceClient {
	return &conversionGoalCampaignConfigServiceClient{cc}
}

func (c *conversionGoalCampaignConfigServiceClient) MutateConversionGoalCampaignConfigs(ctx context.Context, in *MutateConversionGoalCampaignConfigsRequest, opts ...grpc.CallOption) (*MutateConversionGoalCampaignConfigsResponse, error) {
	out := new(MutateConversionGoalCampaignConfigsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v12.services.ConversionGoalCampaignConfigService/MutateConversionGoalCampaignConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionGoalCampaignConfigServiceServer is the server API for ConversionGoalCampaignConfigService service.
// All implementations must embed UnimplementedConversionGoalCampaignConfigServiceServer
// for forward compatibility
type ConversionGoalCampaignConfigServiceServer interface {
	// Creates, updates or removes conversion goal campaign config. Operation
	// statuses are returned.
	MutateConversionGoalCampaignConfigs(context.Context, *MutateConversionGoalCampaignConfigsRequest) (*MutateConversionGoalCampaignConfigsResponse, error)
	mustEmbedUnimplementedConversionGoalCampaignConfigServiceServer()
}

// UnimplementedConversionGoalCampaignConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConversionGoalCampaignConfigServiceServer struct {
}

func (UnimplementedConversionGoalCampaignConfigServiceServer) MutateConversionGoalCampaignConfigs(context.Context, *MutateConversionGoalCampaignConfigsRequest) (*MutateConversionGoalCampaignConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateConversionGoalCampaignConfigs not implemented")
}
func (UnimplementedConversionGoalCampaignConfigServiceServer) mustEmbedUnimplementedConversionGoalCampaignConfigServiceServer() {
}

// UnsafeConversionGoalCampaignConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversionGoalCampaignConfigServiceServer will
// result in compilation errors.
type UnsafeConversionGoalCampaignConfigServiceServer interface {
	mustEmbedUnimplementedConversionGoalCampaignConfigServiceServer()
}

func RegisterConversionGoalCampaignConfigServiceServer(s grpc.ServiceRegistrar, srv ConversionGoalCampaignConfigServiceServer) {
	s.RegisterService(&ConversionGoalCampaignConfigService_ServiceDesc, srv)
}

func _ConversionGoalCampaignConfigService_MutateConversionGoalCampaignConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateConversionGoalCampaignConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionGoalCampaignConfigServiceServer).MutateConversionGoalCampaignConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v12.services.ConversionGoalCampaignConfigService/MutateConversionGoalCampaignConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionGoalCampaignConfigServiceServer).MutateConversionGoalCampaignConfigs(ctx, req.(*MutateConversionGoalCampaignConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversionGoalCampaignConfigService_ServiceDesc is the grpc.ServiceDesc for ConversionGoalCampaignConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversionGoalCampaignConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v12.services.ConversionGoalCampaignConfigService",
	HandlerType: (*ConversionGoalCampaignConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateConversionGoalCampaignConfigs",
			Handler:    _ConversionGoalCampaignConfigService_MutateConversionGoalCampaignConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v12/services/conversion_goal_campaign_config_service.proto",
}
