// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: blogs.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBlog = `-- name: CreateBlog :one
INSERT INTO blogs 
    (author_id, 
    title, 
    content, 
    tsv_content, 
    thumbnail_s3_path, 
    "description",
    read_time,
    category)
VALUES ($1, $2, $3, to_tsvector($3), $4, $5, $6, $7::TEXT[])
RETURNING id
`

type CreateBlogParams struct {
	AuthorID        int32    `json:"author_id"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	ThumbnailS3Path string   `json:"thumbnail_s3_path"`
	Description     string   `json:"description"`
	ReadTime        int32    `json:"read_time"`
	Category        []string `json:"category"`
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (int32, error) {
	row := q.db.QueryRow(ctx, createBlog,
		arg.AuthorID,
		arg.Title,
		arg.Content,
		arg.ThumbnailS3Path,
		arg.Description,
		arg.ReadTime,
		arg.Category,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const featureBlog = `-- name: FeatureBlog :exec
UPDATE blogs
SET is_featured = true
WHERE id =$1
`

func (q *Queries) FeatureBlog(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, featureBlog, id)
	return err
}

const getBlogById = `-- name: GetBlogById :one
SELECT 
    b.id AS blog_id,
    b.title,
    b.content,
    b.thumbnail_s3_path AS blog_thumbnail_url,
    b.category,
    b.description,
    b.read_time,
    b.created_at AS blog_created_at,
    a.name AS author_name,
    a.thumbnail_s3_path AS author_profile_url
FROM 
    blogs b
JOIN 
    authors a ON b.author_id = a.id
WHERE b.id =$1
`

type GetBlogByIdRow struct {
	BlogID           int32            `json:"blog_id"`
	Title            string           `json:"title"`
	Content          string           `json:"content"`
	BlogThumbnailUrl string           `json:"blog_thumbnail_url"`
	Category         []string         `json:"category"`
	Description      string           `json:"description"`
	ReadTime         int32            `json:"read_time"`
	BlogCreatedAt    pgtype.Timestamp `json:"blog_created_at"`
	AuthorName       string           `json:"author_name"`
	AuthorProfileUrl string           `json:"author_profile_url"`
}

func (q *Queries) GetBlogById(ctx context.Context, id int32) (GetBlogByIdRow, error) {
	row := q.db.QueryRow(ctx, getBlogById, id)
	var i GetBlogByIdRow
	err := row.Scan(
		&i.BlogID,
		&i.Title,
		&i.Content,
		&i.BlogThumbnailUrl,
		&i.Category,
		&i.Description,
		&i.ReadTime,
		&i.BlogCreatedAt,
		&i.AuthorName,
		&i.AuthorProfileUrl,
	)
	return i, err
}

const getBlogInCategory = `-- name: GetBlogInCategory :many
SELECT id, author_id, description, title, content, read_time, tsv_content, thumbnail_s3_path, category, is_featured, created_at 
FROM blogs
WHERE $1 = ANY(category)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetBlogInCategoryParams struct {
	Category []string `json:"category"`
	Limit    int32    `json:"limit"`
	Offset   int32    `json:"offset"`
}

func (q *Queries) GetBlogInCategory(ctx context.Context, arg GetBlogInCategoryParams) ([]Blog, error) {
	rows, err := q.db.Query(ctx, getBlogInCategory, arg.Category, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Blog{}
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Description,
			&i.Title,
			&i.Content,
			&i.ReadTime,
			&i.TsvContent,
			&i.ThumbnailS3Path,
			&i.Category,
			&i.IsFeatured,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBlogs = `-- name: GetBlogs :many
SELECT 
    b.id AS blog_id,
    b.title,
    b.thumbnail_s3_path AS blog_thumbnail_url,
    b.category,
    b.description,
    b.read_time,
    b.created_at AS blog_created_at,
    a.name AS author_name,
    a.thumbnail_s3_path AS author_profile_url
FROM 
    blogs b
JOIN 
    authors a ON b.author_id = a.id
ORDER BY 
    b.created_at DESC
LIMIT 
    $1 OFFSET $2
`

type GetBlogsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetBlogsRow struct {
	BlogID           int32            `json:"blog_id"`
	Title            string           `json:"title"`
	BlogThumbnailUrl string           `json:"blog_thumbnail_url"`
	Category         []string         `json:"category"`
	Description      string           `json:"description"`
	ReadTime         int32            `json:"read_time"`
	BlogCreatedAt    pgtype.Timestamp `json:"blog_created_at"`
	AuthorName       string           `json:"author_name"`
	AuthorProfileUrl string           `json:"author_profile_url"`
}

func (q *Queries) GetBlogs(ctx context.Context, arg GetBlogsParams) ([]GetBlogsRow, error) {
	rows, err := q.db.Query(ctx, getBlogs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetBlogsRow{}
	for rows.Next() {
		var i GetBlogsRow
		if err := rows.Scan(
			&i.BlogID,
			&i.Title,
			&i.BlogThumbnailUrl,
			&i.Category,
			&i.Description,
			&i.ReadTime,
			&i.BlogCreatedAt,
			&i.AuthorName,
			&i.AuthorProfileUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFeaturedBlog = `-- name: GetFeaturedBlog :one
SELECT 
    b.id AS blog_id,
    b.title,
    b.content,
    b.thumbnail_s3_path AS blog_thumbnail_url,
    b.category,
    b.description,
    b.read_time,
    b.created_at AS blog_created_at,
    a.name AS author_name,
    a.thumbnail_s3_path AS author_profile_url
FROM 
    blogs b
JOIN 
    authors a ON b.author_id = a.id
WHERE b.is_featured = TRUE
`

type GetFeaturedBlogRow struct {
	BlogID           int32            `json:"blog_id"`
	Title            string           `json:"title"`
	Content          string           `json:"content"`
	BlogThumbnailUrl string           `json:"blog_thumbnail_url"`
	Category         []string         `json:"category"`
	Description      string           `json:"description"`
	ReadTime         int32            `json:"read_time"`
	BlogCreatedAt    pgtype.Timestamp `json:"blog_created_at"`
	AuthorName       string           `json:"author_name"`
	AuthorProfileUrl string           `json:"author_profile_url"`
}

func (q *Queries) GetFeaturedBlog(ctx context.Context) (GetFeaturedBlogRow, error) {
	row := q.db.QueryRow(ctx, getFeaturedBlog)
	var i GetFeaturedBlogRow
	err := row.Scan(
		&i.BlogID,
		&i.Title,
		&i.Content,
		&i.BlogThumbnailUrl,
		&i.Category,
		&i.Description,
		&i.ReadTime,
		&i.BlogCreatedAt,
		&i.AuthorName,
		&i.AuthorProfileUrl,
	)
	return i, err
}
