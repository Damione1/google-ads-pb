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

var newReachPlanClientHook clientHook

// ReachPlanCallOptions contains the retry settings for each method of ReachPlanClient.
type ReachPlanCallOptions struct {
	ListPlannableLocations []gax.CallOption
	ListPlannableProducts  []gax.CallOption
	GenerateReachForecast  []gax.CallOption
}

func defaultReachPlanGRPCClientOptions() []option.ClientOption {
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

func defaultReachPlanCallOptions() *ReachPlanCallOptions {
	return &ReachPlanCallOptions{
		ListPlannableLocations: []gax.CallOption{
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
		ListPlannableProducts: []gax.CallOption{
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
		GenerateReachForecast: []gax.CallOption{
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

// internalReachPlanClient is an interface that defines the methods available from Google Ads API.
type internalReachPlanClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListPlannableLocations(context.Context, *servicespb.ListPlannableLocationsRequest, ...gax.CallOption) (*servicespb.ListPlannableLocationsResponse, error)
	ListPlannableProducts(context.Context, *servicespb.ListPlannableProductsRequest, ...gax.CallOption) (*servicespb.ListPlannableProductsResponse, error)
	GenerateReachForecast(context.Context, *servicespb.GenerateReachForecastRequest, ...gax.CallOption) (*servicespb.GenerateReachForecastResponse, error)
}

// ReachPlanClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Reach Plan Service gives users information about audience size that can
// be reached through advertisement on YouTube. In particular,
// GenerateReachForecast provides estimated number of people of specified
// demographics that can be reached by an ad in a given market by a campaign of
// certain duration with a defined budget.
type ReachPlanClient struct {
	// The internal transport-dependent client.
	internalClient internalReachPlanClient

	// The call options for this service.
	CallOptions *ReachPlanCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ReachPlanClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ReachPlanClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ReachPlanClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListPlannableLocations returns the list of plannable locations (for example, countries).
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *ReachPlanClient) ListPlannableLocations(ctx context.Context, req *servicespb.ListPlannableLocationsRequest, opts ...gax.CallOption) (*servicespb.ListPlannableLocationsResponse, error) {
	return c.internalClient.ListPlannableLocations(ctx, req, opts...)
}

// ListPlannableProducts returns the list of per-location plannable YouTube ad formats with allowed
// targeting.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *ReachPlanClient) ListPlannableProducts(ctx context.Context, req *servicespb.ListPlannableProductsRequest, opts ...gax.CallOption) (*servicespb.ListPlannableProductsResponse, error) {
	return c.internalClient.ListPlannableProducts(ctx, req, opts...)
}

// GenerateReachForecast generates a reach forecast for a given targeting / product mix.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RangeError (at )
// ReachPlanError (at )
// RequestError (at )
func (c *ReachPlanClient) GenerateReachForecast(ctx context.Context, req *servicespb.GenerateReachForecastRequest, opts ...gax.CallOption) (*servicespb.GenerateReachForecastResponse, error) {
	return c.internalClient.GenerateReachForecast(ctx, req, opts...)
}

// reachPlanGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type reachPlanGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ReachPlanClient
	CallOptions **ReachPlanCallOptions

	// The gRPC API client.
	reachPlanClient servicespb.ReachPlanServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewReachPlanClient creates a new reach plan service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Reach Plan Service gives users information about audience size that can
// be reached through advertisement on YouTube. In particular,
// GenerateReachForecast provides estimated number of people of specified
// demographics that can be reached by an ad in a given market by a campaign of
// certain duration with a defined budget.
func NewReachPlanClient(ctx context.Context, opts ...option.ClientOption) (*ReachPlanClient, error) {
	clientOpts := defaultReachPlanGRPCClientOptions()
	if newReachPlanClientHook != nil {
		hookOpts, err := newReachPlanClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ReachPlanClient{CallOptions: defaultReachPlanCallOptions()}

	c := &reachPlanGRPCClient{
		connPool:        connPool,
		reachPlanClient: servicespb.NewReachPlanServiceClient(connPool),
		CallOptions:     &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *reachPlanGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *reachPlanGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *reachPlanGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *reachPlanGRPCClient) ListPlannableLocations(ctx context.Context, req *servicespb.ListPlannableLocationsRequest, opts ...gax.CallOption) (*servicespb.ListPlannableLocationsResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListPlannableLocations[0:len((*c.CallOptions).ListPlannableLocations):len((*c.CallOptions).ListPlannableLocations)], opts...)
	var resp *servicespb.ListPlannableLocationsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.reachPlanClient.ListPlannableLocations(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *reachPlanGRPCClient) ListPlannableProducts(ctx context.Context, req *servicespb.ListPlannableProductsRequest, opts ...gax.CallOption) (*servicespb.ListPlannableProductsResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListPlannableProducts[0:len((*c.CallOptions).ListPlannableProducts):len((*c.CallOptions).ListPlannableProducts)], opts...)
	var resp *servicespb.ListPlannableProductsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.reachPlanClient.ListPlannableProducts(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *reachPlanGRPCClient) GenerateReachForecast(ctx context.Context, req *servicespb.GenerateReachForecastRequest, opts ...gax.CallOption) (*servicespb.GenerateReachForecastResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GenerateReachForecast[0:len((*c.CallOptions).GenerateReachForecast):len((*c.CallOptions).GenerateReachForecast)], opts...)
	var resp *servicespb.GenerateReachForecastResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.reachPlanClient.GenerateReachForecast(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
