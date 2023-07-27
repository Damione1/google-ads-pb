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
	"google.golang.org/grpc/metadata"
)

var newCustomerCustomizerClientHook clientHook

// CustomerCustomizerCallOptions contains the retry settings for each method of CustomerCustomizerClient.
type CustomerCustomizerCallOptions struct {
	MutateCustomerCustomizers []gax.CallOption
}

func defaultCustomerCustomizerGRPCClientOptions() []option.ClientOption {
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

func defaultCustomerCustomizerCallOptions() *CustomerCustomizerCallOptions {
	return &CustomerCustomizerCallOptions{
		MutateCustomerCustomizers: []gax.CallOption{
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

// internalCustomerCustomizerClient is an interface that defines the methods available from Google Ads API.
type internalCustomerCustomizerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCustomerCustomizers(context.Context, *servicespb.MutateCustomerCustomizersRequest, ...gax.CallOption) (*servicespb.MutateCustomerCustomizersResponse, error)
}

// CustomerCustomizerClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage customer customizer
type CustomerCustomizerClient struct {
	// The internal transport-dependent client.
	internalClient internalCustomerCustomizerClient

	// The call options for this service.
	CallOptions *CustomerCustomizerCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CustomerCustomizerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CustomerCustomizerClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CustomerCustomizerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCustomerCustomizers creates, updates or removes customer customizers. Operation statuses are
// returned.
func (c *CustomerCustomizerClient) MutateCustomerCustomizers(ctx context.Context, req *servicespb.MutateCustomerCustomizersRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerCustomizersResponse, error) {
	return c.internalClient.MutateCustomerCustomizers(ctx, req, opts...)
}

// customerCustomizerGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type customerCustomizerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CustomerCustomizerClient
	CallOptions **CustomerCustomizerCallOptions

	// The gRPC API client.
	customerCustomizerClient servicespb.CustomerCustomizerServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCustomerCustomizerClient creates a new customer customizer service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage customer customizer
func NewCustomerCustomizerClient(ctx context.Context, opts ...option.ClientOption) (*CustomerCustomizerClient, error) {
	clientOpts := defaultCustomerCustomizerGRPCClientOptions()
	if newCustomerCustomizerClientHook != nil {
		hookOpts, err := newCustomerCustomizerClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CustomerCustomizerClient{CallOptions: defaultCustomerCustomizerCallOptions()}

	c := &customerCustomizerGRPCClient{
		connPool:                 connPool,
		customerCustomizerClient: servicespb.NewCustomerCustomizerServiceClient(connPool),
		CallOptions:              &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *customerCustomizerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *customerCustomizerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *customerCustomizerGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *customerCustomizerGRPCClient) MutateCustomerCustomizers(ctx context.Context, req *servicespb.MutateCustomerCustomizersRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerCustomizersResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCustomerCustomizers[0:len((*c.CallOptions).MutateCustomerCustomizers):len((*c.CallOptions).MutateCustomerCustomizers)], opts...)
	var resp *servicespb.MutateCustomerCustomizersResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerCustomizerClient.MutateCustomerCustomizers(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
