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

	gax "github.com/googleapis/gax-go/v2"
	servicespb "github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newAdGroupCustomizerClientHook clientHook

// AdGroupCustomizerCallOptions contains the retry settings for each method of AdGroupCustomizerClient.
type AdGroupCustomizerCallOptions struct {
	MutateAdGroupCustomizers []gax.CallOption
}

func defaultAdGroupCustomizerGRPCClientOptions() []option.ClientOption {
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

func defaultAdGroupCustomizerCallOptions() *AdGroupCustomizerCallOptions {
	return &AdGroupCustomizerCallOptions{
		MutateAdGroupCustomizers: []gax.CallOption{
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

// internalAdGroupCustomizerClient is an interface that defines the methods available from Google Ads API.
type internalAdGroupCustomizerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAdGroupCustomizers(context.Context, *servicespb.MutateAdGroupCustomizersRequest, ...gax.CallOption) (*servicespb.MutateAdGroupCustomizersResponse, error)
}

// AdGroupCustomizerClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage ad group customizer
type AdGroupCustomizerClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupCustomizerClient

	// The call options for this service.
	CallOptions *AdGroupCustomizerCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupCustomizerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupCustomizerClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AdGroupCustomizerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAdGroupCustomizers creates, updates or removes ad group customizers. Operation statuses are
// returned.
func (c *AdGroupCustomizerClient) MutateAdGroupCustomizers(ctx context.Context, req *servicespb.MutateAdGroupCustomizersRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupCustomizersResponse, error) {
	return c.internalClient.MutateAdGroupCustomizers(ctx, req, opts...)
}

// adGroupCustomizerGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupCustomizerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AdGroupCustomizerClient
	CallOptions **AdGroupCustomizerCallOptions

	// The gRPC API client.
	adGroupCustomizerClient servicespb.AdGroupCustomizerServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewAdGroupCustomizerClient creates a new ad group customizer service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage ad group customizer
func NewAdGroupCustomizerClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupCustomizerClient, error) {
	clientOpts := defaultAdGroupCustomizerGRPCClientOptions()
	if newAdGroupCustomizerClientHook != nil {
		hookOpts, err := newAdGroupCustomizerClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AdGroupCustomizerClient{CallOptions: defaultAdGroupCustomizerCallOptions()}

	c := &adGroupCustomizerGRPCClient{
		connPool:                connPool,
		adGroupCustomizerClient: servicespb.NewAdGroupCustomizerServiceClient(connPool),
		CallOptions:             &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *adGroupCustomizerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupCustomizerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupCustomizerGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupCustomizerGRPCClient) MutateAdGroupCustomizers(ctx context.Context, req *servicespb.MutateAdGroupCustomizersRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupCustomizersResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateAdGroupCustomizers[0:len((*c.CallOptions).MutateAdGroupCustomizers):len((*c.CallOptions).MutateAdGroupCustomizers)], opts...)
	var resp *servicespb.MutateAdGroupCustomizersResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupCustomizerClient.MutateAdGroupCustomizers(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
