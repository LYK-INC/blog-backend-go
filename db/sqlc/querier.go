// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	AddRole(ctx context.Context, arg AddRoleParams) error
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (int32, error)
	CreateBlog(ctx context.Context, arg CreateBlogParams) (int32, error)
	GetBlogById(ctx context.Context, id int32) (Blog, error)
	GetBlogInCategory(ctx context.Context, arg GetBlogInCategoryParams) ([]Blog, error)
	GetBlogs(ctx context.Context, arg GetBlogsParams) ([]Blog, error)
	// Ensure role isn't already present
	RemoveRole(ctx context.Context, arg RemoveRoleParams) error
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error
}

var _ Querier = (*Queries)(nil)
