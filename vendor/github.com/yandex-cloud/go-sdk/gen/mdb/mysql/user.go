// Code generated by sdkgen. DO NOT EDIT.

//nolint
package mysql

import (
	"context"

	"google.golang.org/grpc"

	mysql "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// UserServiceClient is a mysql.UserServiceClient with
// lazy GRPC connection initialization.
type UserServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ mysql.UserServiceClient = &UserServiceClient{}

// Create implements mysql.UserServiceClient
func (c *UserServiceClient) Create(ctx context.Context, in *mysql.CreateUserRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewUserServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements mysql.UserServiceClient
func (c *UserServiceClient) Delete(ctx context.Context, in *mysql.DeleteUserRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewUserServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements mysql.UserServiceClient
func (c *UserServiceClient) Get(ctx context.Context, in *mysql.GetUserRequest, opts ...grpc.CallOption) (*mysql.User, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewUserServiceClient(conn).Get(ctx, in, opts...)
}

// GrantPermission implements mysql.UserServiceClient
func (c *UserServiceClient) GrantPermission(ctx context.Context, in *mysql.GrantUserPermissionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewUserServiceClient(conn).GrantPermission(ctx, in, opts...)
}

// List implements mysql.UserServiceClient
func (c *UserServiceClient) List(ctx context.Context, in *mysql.ListUsersRequest, opts ...grpc.CallOption) (*mysql.ListUsersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewUserServiceClient(conn).List(ctx, in, opts...)
}

// RevokePermission implements mysql.UserServiceClient
func (c *UserServiceClient) RevokePermission(ctx context.Context, in *mysql.RevokeUserPermissionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewUserServiceClient(conn).RevokePermission(ctx, in, opts...)
}

// Update implements mysql.UserServiceClient
func (c *UserServiceClient) Update(ctx context.Context, in *mysql.UpdateUserRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewUserServiceClient(conn).Update(ctx, in, opts...)
}
