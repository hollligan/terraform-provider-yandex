// Code generated by sdkgen. DO NOT EDIT.

//nolint
package postgresql

import (
	"context"

	"google.golang.org/grpc"

	postgresql "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1"
)

//revive:disable

// ResourcePresetServiceClient is a postgresql.ResourcePresetServiceClient with
// lazy GRPC connection initialization.
type ResourcePresetServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ postgresql.ResourcePresetServiceClient = &ResourcePresetServiceClient{}

// Get implements postgresql.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) Get(ctx context.Context, in *postgresql.GetResourcePresetRequest, opts ...grpc.CallOption) (*postgresql.ResourcePreset, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return postgresql.NewResourcePresetServiceClient(conn).Get(ctx, in, opts...)
}

// List implements postgresql.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) List(ctx context.Context, in *postgresql.ListResourcePresetsRequest, opts ...grpc.CallOption) (*postgresql.ListResourcePresetsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return postgresql.NewResourcePresetServiceClient(conn).List(ctx, in, opts...)
}
