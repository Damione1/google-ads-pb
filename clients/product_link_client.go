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

var newProductLinkClientHook clientHook

// ProductLinkCallOptions contains the retry settings for each method of ProductLinkClient.
type ProductLinkCallOptions struct {
	CreateProductLink []gax.CallOption
	RemoveProductLink []gax.CallOption
}

func defaultProductLinkGRPCClientOptions() []option.ClientOption {
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

func defaultProductLinkCallOptions() *ProductLinkCallOptions {
	return &ProductLinkCallOptions{
		CreateProductLink: []gax.CallOption{
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
		RemoveProductLink: []gax.CallOption{
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

// internalProductLinkClient is an interface that defines the methods available from Google Ads API.
type internalProductLinkClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateProductLink(context.Context, *servicespb.CreateProductLinkRequest, ...gax.CallOption) (*servicespb.CreateProductLinkResponse, error)
	RemoveProductLink(context.Context, *servicespb.RemoveProductLinkRequest, ...gax.CallOption) (*servicespb.RemoveProductLinkResponse, error)
}

// ProductLinkClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This service allows management of links between  a Google
// Ads customer and another product.
type ProductLinkClient struct {
	// The internal transport-dependent client.
	internalClient internalProductLinkClient

	// The call options for this service.
	CallOptions *ProductLinkCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ProductLinkClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ProductLinkClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ProductLinkClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateProductLink creates a product link.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *ProductLinkClient) CreateProductLink(ctx context.Context, req *servicespb.CreateProductLinkRequest, opts ...gax.CallOption) (*servicespb.CreateProductLinkResponse, error) {
	return c.internalClient.CreateProductLink(ctx, req, opts...)
}

// RemoveProductLink removes a product link.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *ProductLinkClient) RemoveProductLink(ctx context.Context, req *servicespb.RemoveProductLinkRequest, opts ...gax.CallOption) (*servicespb.RemoveProductLinkResponse, error) {
	return c.internalClient.RemoveProductLink(ctx, req, opts...)
}

// productLinkGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type productLinkGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ProductLinkClient
	CallOptions **ProductLinkCallOptions

	// The gRPC API client.
	productLinkClient servicespb.ProductLinkServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewProductLinkClient creates a new product link service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This service allows management of links between  a Google
// Ads customer and another product.
func NewProductLinkClient(ctx context.Context, opts ...option.ClientOption) (*ProductLinkClient, error) {
	clientOpts := defaultProductLinkGRPCClientOptions()
	if newProductLinkClientHook != nil {
		hookOpts, err := newProductLinkClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ProductLinkClient{CallOptions: defaultProductLinkCallOptions()}

	c := &productLinkGRPCClient{
		connPool:          connPool,
		productLinkClient: servicespb.NewProductLinkServiceClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *productLinkGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *productLinkGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *productLinkGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *productLinkGRPCClient) CreateProductLink(ctx context.Context, req *servicespb.CreateProductLinkRequest, opts ...gax.CallOption) (*servicespb.CreateProductLinkResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateProductLink[0:len((*c.CallOptions).CreateProductLink):len((*c.CallOptions).CreateProductLink)], opts...)
	var resp *servicespb.CreateProductLinkResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productLinkClient.CreateProductLink(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *productLinkGRPCClient) RemoveProductLink(ctx context.Context, req *servicespb.RemoveProductLinkRequest, opts ...gax.CallOption) (*servicespb.RemoveProductLinkResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).RemoveProductLink[0:len((*c.CallOptions).RemoveProductLink):len((*c.CallOptions).RemoveProductLink)], opts...)
	var resp *servicespb.RemoveProductLinkResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productLinkClient.RemoveProductLink(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
