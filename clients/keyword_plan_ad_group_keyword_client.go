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

var newKeywordPlanAdGroupKeywordClientHook clientHook

// KeywordPlanAdGroupKeywordCallOptions contains the retry settings for each method of KeywordPlanAdGroupKeywordClient.
type KeywordPlanAdGroupKeywordCallOptions struct {
	MutateKeywordPlanAdGroupKeywords []gax.CallOption
}

func defaultKeywordPlanAdGroupKeywordGRPCClientOptions() []option.ClientOption {
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

func defaultKeywordPlanAdGroupKeywordCallOptions() *KeywordPlanAdGroupKeywordCallOptions {
	return &KeywordPlanAdGroupKeywordCallOptions{
		MutateKeywordPlanAdGroupKeywords: []gax.CallOption{
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

// internalKeywordPlanAdGroupKeywordClient is an interface that defines the methods available from Google Ads API.
type internalKeywordPlanAdGroupKeywordClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateKeywordPlanAdGroupKeywords(context.Context, *servicespb.MutateKeywordPlanAdGroupKeywordsRequest, ...gax.CallOption) (*servicespb.MutateKeywordPlanAdGroupKeywordsResponse, error)
}

// KeywordPlanAdGroupKeywordClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage Keyword Plan ad group keywords. KeywordPlanAdGroup is
// required to add ad group keywords. Positive and negative keywords are
// supported. A maximum of 10,000 positive keywords are allowed per keyword
// plan. A maximum of 1,000 negative keywords are allower per keyword plan. This
// includes campaign negative keywords and ad group negative keywords.
type KeywordPlanAdGroupKeywordClient struct {
	// The internal transport-dependent client.
	internalClient internalKeywordPlanAdGroupKeywordClient

	// The call options for this service.
	CallOptions *KeywordPlanAdGroupKeywordCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *KeywordPlanAdGroupKeywordClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *KeywordPlanAdGroupKeywordClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *KeywordPlanAdGroupKeywordClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateKeywordPlanAdGroupKeywords creates, updates, or removes Keyword Plan ad group keywords. Operation
// statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// KeywordPlanAdGroupKeywordError (at )
// KeywordPlanError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
func (c *KeywordPlanAdGroupKeywordClient) MutateKeywordPlanAdGroupKeywords(ctx context.Context, req *servicespb.MutateKeywordPlanAdGroupKeywordsRequest, opts ...gax.CallOption) (*servicespb.MutateKeywordPlanAdGroupKeywordsResponse, error) {
	return c.internalClient.MutateKeywordPlanAdGroupKeywords(ctx, req, opts...)
}

// keywordPlanAdGroupKeywordGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type keywordPlanAdGroupKeywordGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing KeywordPlanAdGroupKeywordClient
	CallOptions **KeywordPlanAdGroupKeywordCallOptions

	// The gRPC API client.
	keywordPlanAdGroupKeywordClient servicespb.KeywordPlanAdGroupKeywordServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewKeywordPlanAdGroupKeywordClient creates a new keyword plan ad group keyword service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage Keyword Plan ad group keywords. KeywordPlanAdGroup is
// required to add ad group keywords. Positive and negative keywords are
// supported. A maximum of 10,000 positive keywords are allowed per keyword
// plan. A maximum of 1,000 negative keywords are allower per keyword plan. This
// includes campaign negative keywords and ad group negative keywords.
func NewKeywordPlanAdGroupKeywordClient(ctx context.Context, opts ...option.ClientOption) (*KeywordPlanAdGroupKeywordClient, error) {
	clientOpts := defaultKeywordPlanAdGroupKeywordGRPCClientOptions()
	if newKeywordPlanAdGroupKeywordClientHook != nil {
		hookOpts, err := newKeywordPlanAdGroupKeywordClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := KeywordPlanAdGroupKeywordClient{CallOptions: defaultKeywordPlanAdGroupKeywordCallOptions()}

	c := &keywordPlanAdGroupKeywordGRPCClient{
		connPool:                        connPool,
		keywordPlanAdGroupKeywordClient: servicespb.NewKeywordPlanAdGroupKeywordServiceClient(connPool),
		CallOptions:                     &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *keywordPlanAdGroupKeywordGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *keywordPlanAdGroupKeywordGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *keywordPlanAdGroupKeywordGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *keywordPlanAdGroupKeywordGRPCClient) MutateKeywordPlanAdGroupKeywords(ctx context.Context, req *servicespb.MutateKeywordPlanAdGroupKeywordsRequest, opts ...gax.CallOption) (*servicespb.MutateKeywordPlanAdGroupKeywordsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateKeywordPlanAdGroupKeywords[0:len((*c.CallOptions).MutateKeywordPlanAdGroupKeywords):len((*c.CallOptions).MutateKeywordPlanAdGroupKeywords)], opts...)
	var resp *servicespb.MutateKeywordPlanAdGroupKeywordsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanAdGroupKeywordClient.MutateKeywordPlanAdGroupKeywords(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
