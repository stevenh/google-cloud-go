// Copyright 2022 Google LLC
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

package resourcemanager

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newOrganizationsClientHook clientHook

// OrganizationsCallOptions contains the retry settings for each method of OrganizationsClient.
type OrganizationsCallOptions struct {
	GetOrganization     []gax.CallOption
	SearchOrganizations []gax.CallOption
	GetIamPolicy        []gax.CallOption
	SetIamPolicy        []gax.CallOption
	TestIamPermissions  []gax.CallOption
}

func defaultOrganizationsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudresourcemanager.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudresourcemanager.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudresourcemanager.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultOrganizationsCallOptions() *OrganizationsCallOptions {
	return &OrganizationsCallOptions{
		GetOrganization: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchOrganizations: []gax.CallOption{},
		GetIamPolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SetIamPolicy:       []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{},
	}
}

// internalOrganizationsClient is an interface that defines the methods available from Cloud Resource Manager API.
type internalOrganizationsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetOrganization(context.Context, *resourcemanagerpb.GetOrganizationRequest, ...gax.CallOption) (*resourcemanagerpb.Organization, error)
	SearchOrganizations(context.Context, *resourcemanagerpb.SearchOrganizationsRequest, ...gax.CallOption) *OrganizationIterator
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
}

// OrganizationsClient is a client for interacting with Cloud Resource Manager API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Allows users to manage their organization resources.
type OrganizationsClient struct {
	// The internal transport-dependent client.
	internalClient internalOrganizationsClient

	// The call options for this service.
	CallOptions *OrganizationsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *OrganizationsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *OrganizationsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *OrganizationsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetOrganization fetches an organization resource identified by the specified resource name.
func (c *OrganizationsClient) GetOrganization(ctx context.Context, req *resourcemanagerpb.GetOrganizationRequest, opts ...gax.CallOption) (*resourcemanagerpb.Organization, error) {
	return c.internalClient.GetOrganization(ctx, req, opts...)
}

// SearchOrganizations searches organization resources that are visible to the user and satisfy
// the specified filter. This method returns organizations in an unspecified
// order. New organizations do not necessarily appear at the end of the
// results, and may take a small amount of time to appear.
//
// Search will only return organizations on which the user has the permission
// resourcemanager.organizations.get
func (c *OrganizationsClient) SearchOrganizations(ctx context.Context, req *resourcemanagerpb.SearchOrganizationsRequest, opts ...gax.CallOption) *OrganizationIterator {
	return c.internalClient.SearchOrganizations(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for an organization resource. The policy may
// be empty if no such policy or resource exists. The resource field should
// be the organization’s resource name, for example: “organizations/123”.
//
// Authorization requires the IAM permission
// resourcemanager.organizations.getIamPolicy on the specified organization.
func (c *OrganizationsClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on an organization resource. Replaces any
// existing policy. The resource field should be the organization’s resource
// name, for example: “organizations/123”.
//
// Authorization requires the IAM permission
// resourcemanager.organizations.setIamPolicy on the specified organization.
func (c *OrganizationsClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns the permissions that a caller has on the specified organization.
// The resource field should be the organization’s resource name,
// for example: “organizations/123”.
//
// There are no permissions required for making this API call.
func (c *OrganizationsClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// organizationsGRPCClient is a client for interacting with Cloud Resource Manager API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type organizationsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing OrganizationsClient
	CallOptions **OrganizationsCallOptions

	// The gRPC API client.
	organizationsClient resourcemanagerpb.OrganizationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewOrganizationsClient creates a new organizations client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Allows users to manage their organization resources.
func NewOrganizationsClient(ctx context.Context, opts ...option.ClientOption) (*OrganizationsClient, error) {
	clientOpts := defaultOrganizationsGRPCClientOptions()
	if newOrganizationsClientHook != nil {
		hookOpts, err := newOrganizationsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := OrganizationsClient{CallOptions: defaultOrganizationsCallOptions()}

	c := &organizationsGRPCClient{
		connPool:            connPool,
		disableDeadlines:    disableDeadlines,
		organizationsClient: resourcemanagerpb.NewOrganizationsClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *organizationsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *organizationsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *organizationsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *organizationsGRPCClient) GetOrganization(ctx context.Context, req *resourcemanagerpb.GetOrganizationRequest, opts ...gax.CallOption) (*resourcemanagerpb.Organization, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetOrganization[0:len((*c.CallOptions).GetOrganization):len((*c.CallOptions).GetOrganization)], opts...)
	var resp *resourcemanagerpb.Organization
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.organizationsClient.GetOrganization(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *organizationsGRPCClient) SearchOrganizations(ctx context.Context, req *resourcemanagerpb.SearchOrganizationsRequest, opts ...gax.CallOption) *OrganizationIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).SearchOrganizations[0:len((*c.CallOptions).SearchOrganizations):len((*c.CallOptions).SearchOrganizations)], opts...)
	it := &OrganizationIterator{}
	req = proto.Clone(req).(*resourcemanagerpb.SearchOrganizationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcemanagerpb.Organization, string, error) {
		resp := &resourcemanagerpb.SearchOrganizationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.organizationsClient.SearchOrganizations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOrganizations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *organizationsGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.organizationsClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *organizationsGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.organizationsClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *organizationsGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.organizationsClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// OrganizationIterator manages a stream of *resourcemanagerpb.Organization.
type OrganizationIterator struct {
	items    []*resourcemanagerpb.Organization
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*resourcemanagerpb.Organization, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *OrganizationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OrganizationIterator) Next() (*resourcemanagerpb.Organization, error) {
	var item *resourcemanagerpb.Organization
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OrganizationIterator) bufLen() int {
	return len(it.items)
}

func (it *OrganizationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
