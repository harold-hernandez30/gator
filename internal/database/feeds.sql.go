// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, user_id, name, url)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id, created_at, updated_at, user_id, name, url
`

type CreateFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	Name      string
	Url       string
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.Name,
		arg.Url,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Name,
		&i.Url,
	)
	return i, err
}
