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

var newAccountLinkClientHook clientHook

// AccountLinkCallOptions contains the retry settings for each method of AccountLinkClient.
type AccountLinkCallOptions struct {
	CreateAccountLink []gax.CallOption
	MutateAccountLink []gax.CallOption
}

func defaultAccountLinkGRPCClientOptions() []option.ClientOption {
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

func defaultAccountLinkCallOptions() *AccountLinkCallOptions {
	return &AccountLinkCallOptions{
		CreateAccountLink: []gax.CallOption{
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
		MutateAccountLink: []gax.CallOption{
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

// internalAccountLinkClient is an interface that defines the methods available from Google Ads API.
type internalAccountLinkClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateAccountLink(context.Context, *servicespb.CreateAccountLinkRequest, ...gax.CallOption) (*servicespb.CreateAccountLinkResponse, error)
	MutateAccountLink(context.Context, *servicespb.MutateAccountLinkRequest, ...gax.CallOption) (*servicespb.MutateAccountLinkResponse, error)
}

// AccountLinkClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This service allows management of links between Google Ads accounts and other
// accounts.
type AccountLinkClient struct {
	// The internal transport-dependent client.
	internalClient internalAccountLinkClient

	// The call options for this service.
	CallOptions *AccountLinkCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AccountLinkClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AccountLinkClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AccountLinkClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateAccountLink creates an account link.
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
// ThirdPartyAppAnalyticsLinkError (at )
func (c *AccountLinkClient) CreateAccountLink(ctx context.Context, req *servicespb.CreateAccountLinkRequest, opts ...gax.CallOption) (*servicespb.CreateAccountLinkResponse, error) {
	return c.internalClient.CreateAccountLink(ctx, req, opts...)
}

// MutateAccountLink creates or removes an account link.
// From V5, create is not supported through
// AccountLinkService.MutateAccountLink. Use
// AccountLinkService.CreateAccountLink instead.
//
// List of thrown errors:
// AccountLinkError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *AccountLinkClient) MutateAccountLink(ctx context.Context, req *servicespb.MutateAccountLinkRequest, opts ...gax.CallOption) (*servicespb.MutateAccountLinkResponse, error) {
	return c.internalClient.MutateAccountLink(ctx, req, opts...)
}

// accountLinkGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type accountLinkGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AccountLinkClient
	CallOptions **AccountLinkCallOptions

	// The gRPC API client.
	accountLinkClient servicespb.AccountLinkServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAccountLinkClient creates a new account link service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This service allows management of links between Google Ads accounts and other
// accounts.
func NewAccountLinkClient(ctx context.Context, opts ...option.ClientOption) (*AccountLinkClient, error) {
	clientOpts := defaultAccountLinkGRPCClientOptions()
	if newAccountLinkClientHook != nil {
		hookOpts, err := newAccountLinkClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AccountLinkClient{CallOptions: defaultAccountLinkCallOptions()}

	c := &accountLinkGRPCClient{
		connPool:          connPool,
		accountLinkClient: servicespb.NewAccountLinkServiceClient(connPool),
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
func (c *accountLinkGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *accountLinkGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *accountLinkGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *accountLinkGRPCClient) CreateAccountLink(ctx context.Context, req *servicespb.CreateAccountLinkRequest, opts ...gax.CallOption) (*servicespb.CreateAccountLinkResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateAccountLink[0:len((*c.CallOptions).CreateAccountLink):len((*c.CallOptions).CreateAccountLink)], opts...)
	var resp *servicespb.CreateAccountLinkResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.accountLinkClient.CreateAccountLink(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *accountLinkGRPCClient) MutateAccountLink(ctx context.Context, req *servicespb.MutateAccountLinkRequest, opts ...gax.CallOption) (*servicespb.MutateAccountLinkResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateAccountLink[0:len((*c.CallOptions).MutateAccountLink):len((*c.CallOptions).MutateAccountLink)], opts...)
	var resp *servicespb.MutateAccountLinkResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.accountLinkClient.MutateAccountLink(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
