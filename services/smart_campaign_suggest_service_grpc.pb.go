// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: google/ads/googleads/v11/services/smart_campaign_suggest_service.proto

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

// SmartCampaignSuggestServiceClient is the client API for SmartCampaignSuggestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmartCampaignSuggestServiceClient interface {
	// Returns BudgetOption suggestions.
	SuggestSmartCampaignBudgetOptions(ctx context.Context, in *SuggestSmartCampaignBudgetOptionsRequest, opts ...grpc.CallOption) (*SuggestSmartCampaignBudgetOptionsResponse, error)
	// Suggests a Smart campaign ad compatible with the Ad family of resources,
	// based on data points such as targeting and the business to advertise.
	SuggestSmartCampaignAd(ctx context.Context, in *SuggestSmartCampaignAdRequest, opts ...grpc.CallOption) (*SuggestSmartCampaignAdResponse, error)
	// Suggests keyword themes to advertise on.
	SuggestKeywordThemes(ctx context.Context, in *SuggestKeywordThemesRequest, opts ...grpc.CallOption) (*SuggestKeywordThemesResponse, error)
}

type smartCampaignSuggestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmartCampaignSuggestServiceClient(cc grpc.ClientConnInterface) SmartCampaignSuggestServiceClient {
	return &smartCampaignSuggestServiceClient{cc}
}

func (c *smartCampaignSuggestServiceClient) SuggestSmartCampaignBudgetOptions(ctx context.Context, in *SuggestSmartCampaignBudgetOptionsRequest, opts ...grpc.CallOption) (*SuggestSmartCampaignBudgetOptionsResponse, error) {
	out := new(SuggestSmartCampaignBudgetOptionsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.SmartCampaignSuggestService/SuggestSmartCampaignBudgetOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartCampaignSuggestServiceClient) SuggestSmartCampaignAd(ctx context.Context, in *SuggestSmartCampaignAdRequest, opts ...grpc.CallOption) (*SuggestSmartCampaignAdResponse, error) {
	out := new(SuggestSmartCampaignAdResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.SmartCampaignSuggestService/SuggestSmartCampaignAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartCampaignSuggestServiceClient) SuggestKeywordThemes(ctx context.Context, in *SuggestKeywordThemesRequest, opts ...grpc.CallOption) (*SuggestKeywordThemesResponse, error) {
	out := new(SuggestKeywordThemesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.SmartCampaignSuggestService/SuggestKeywordThemes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartCampaignSuggestServiceServer is the server API for SmartCampaignSuggestService service.
// All implementations must embed UnimplementedSmartCampaignSuggestServiceServer
// for forward compatibility
type SmartCampaignSuggestServiceServer interface {
	// Returns BudgetOption suggestions.
	SuggestSmartCampaignBudgetOptions(context.Context, *SuggestSmartCampaignBudgetOptionsRequest) (*SuggestSmartCampaignBudgetOptionsResponse, error)
	// Suggests a Smart campaign ad compatible with the Ad family of resources,
	// based on data points such as targeting and the business to advertise.
	SuggestSmartCampaignAd(context.Context, *SuggestSmartCampaignAdRequest) (*SuggestSmartCampaignAdResponse, error)
	// Suggests keyword themes to advertise on.
	SuggestKeywordThemes(context.Context, *SuggestKeywordThemesRequest) (*SuggestKeywordThemesResponse, error)
	mustEmbedUnimplementedSmartCampaignSuggestServiceServer()
}

// UnimplementedSmartCampaignSuggestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmartCampaignSuggestServiceServer struct {
}

func (UnimplementedSmartCampaignSuggestServiceServer) SuggestSmartCampaignBudgetOptions(context.Context, *SuggestSmartCampaignBudgetOptionsRequest) (*SuggestSmartCampaignBudgetOptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestSmartCampaignBudgetOptions not implemented")
}
func (UnimplementedSmartCampaignSuggestServiceServer) SuggestSmartCampaignAd(context.Context, *SuggestSmartCampaignAdRequest) (*SuggestSmartCampaignAdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestSmartCampaignAd not implemented")
}
func (UnimplementedSmartCampaignSuggestServiceServer) SuggestKeywordThemes(context.Context, *SuggestKeywordThemesRequest) (*SuggestKeywordThemesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestKeywordThemes not implemented")
}
func (UnimplementedSmartCampaignSuggestServiceServer) mustEmbedUnimplementedSmartCampaignSuggestServiceServer() {
}

// UnsafeSmartCampaignSuggestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmartCampaignSuggestServiceServer will
// result in compilation errors.
type UnsafeSmartCampaignSuggestServiceServer interface {
	mustEmbedUnimplementedSmartCampaignSuggestServiceServer()
}

func RegisterSmartCampaignSuggestServiceServer(s grpc.ServiceRegistrar, srv SmartCampaignSuggestServiceServer) {
	s.RegisterService(&SmartCampaignSuggestService_ServiceDesc, srv)
}

func _SmartCampaignSuggestService_SuggestSmartCampaignBudgetOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestSmartCampaignBudgetOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCampaignSuggestServiceServer).SuggestSmartCampaignBudgetOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.SmartCampaignSuggestService/SuggestSmartCampaignBudgetOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCampaignSuggestServiceServer).SuggestSmartCampaignBudgetOptions(ctx, req.(*SuggestSmartCampaignBudgetOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartCampaignSuggestService_SuggestSmartCampaignAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestSmartCampaignAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCampaignSuggestServiceServer).SuggestSmartCampaignAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.SmartCampaignSuggestService/SuggestSmartCampaignAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCampaignSuggestServiceServer).SuggestSmartCampaignAd(ctx, req.(*SuggestSmartCampaignAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartCampaignSuggestService_SuggestKeywordThemes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestKeywordThemesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCampaignSuggestServiceServer).SuggestKeywordThemes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.SmartCampaignSuggestService/SuggestKeywordThemes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCampaignSuggestServiceServer).SuggestKeywordThemes(ctx, req.(*SuggestKeywordThemesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmartCampaignSuggestService_ServiceDesc is the grpc.ServiceDesc for SmartCampaignSuggestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmartCampaignSuggestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.SmartCampaignSuggestService",
	HandlerType: (*SmartCampaignSuggestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuggestSmartCampaignBudgetOptions",
			Handler:    _SmartCampaignSuggestService_SuggestSmartCampaignBudgetOptions_Handler,
		},
		{
			MethodName: "SuggestSmartCampaignAd",
			Handler:    _SmartCampaignSuggestService_SuggestSmartCampaignAd_Handler,
		},
		{
			MethodName: "SuggestKeywordThemes",
			Handler:    _SmartCampaignSuggestService_SuggestKeywordThemes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/smart_campaign_suggest_service.proto",
}
