// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package clients

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	servicespb "github.com/Damione1/google-ads-pb/services"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newBrandSuggestionClientHook clientHook

// BrandSuggestionCallOptions contains the retry settings for each method of BrandSuggestionClient.
type BrandSuggestionCallOptions struct {
	SuggestBrands []gax.CallOption
}

func defaultBrandSuggestionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBrandSuggestionCallOptions() *BrandSuggestionCallOptions {
	return &BrandSuggestionCallOptions{
		SuggestBrands: []gax.CallOption{
			gax.WithTimeout(14400000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalBrandSuggestionClient is an interface that defines the methods available from Google Ads API.
type internalBrandSuggestionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SuggestBrands(context.Context, *servicespb.SuggestBrandsRequest, ...gax.CallOption) (*servicespb.SuggestBrandsResponse, error)
}

// BrandSuggestionClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This service will suggest brands based on a prefix.
type BrandSuggestionClient struct {
	// The internal transport-dependent client.
	internalClient internalBrandSuggestionClient

	// The call options for this service.
	CallOptions *BrandSuggestionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BrandSuggestionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BrandSuggestionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BrandSuggestionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SuggestBrands rpc to return a list of matching brands based on a prefix for this
// customer.
func (c *BrandSuggestionClient) SuggestBrands(ctx context.Context, req *servicespb.SuggestBrandsRequest, opts ...gax.CallOption) (*servicespb.SuggestBrandsResponse, error) {
	return c.internalClient.SuggestBrands(ctx, req, opts...)
}

// brandSuggestionGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type brandSuggestionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing BrandSuggestionClient
	CallOptions **BrandSuggestionCallOptions

	// The gRPC API client.
	brandSuggestionClient servicespb.BrandSuggestionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewBrandSuggestionClient creates a new brand suggestion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This service will suggest brands based on a prefix.
func NewBrandSuggestionClient(ctx context.Context, opts ...option.ClientOption) (*BrandSuggestionClient, error) {
	clientOpts := defaultBrandSuggestionGRPCClientOptions()
	if newBrandSuggestionClientHook != nil {
		hookOpts, err := newBrandSuggestionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BrandSuggestionClient{CallOptions: defaultBrandSuggestionCallOptions()}

	c := &brandSuggestionGRPCClient{
		connPool:              connPool,
		brandSuggestionClient: servicespb.NewBrandSuggestionServiceClient(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *brandSuggestionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *brandSuggestionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *brandSuggestionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *brandSuggestionGRPCClient) SuggestBrands(ctx context.Context, req *servicespb.SuggestBrandsRequest, opts ...gax.CallOption) (*servicespb.SuggestBrandsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SuggestBrands[0:len((*c.CallOptions).SuggestBrands):len((*c.CallOptions).SuggestBrands)], opts...)
	var resp *servicespb.SuggestBrandsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.brandSuggestionClient.SuggestBrands(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
