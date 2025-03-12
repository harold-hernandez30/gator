// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feeds.sql

package database

import (
	"context"
	"database/sql"
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
RETURNING id, created_at, updated_at, user_id, name, url, last_fetched_at
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
		&i.LastFetchedAt,
	)
	return i, err
}

const getFeed = `-- name: GetFeed :one
SELECT id, created_at, updated_at, user_id, name, url, last_fetched_at FROM feeds
WHERE url = $1
`

func (q *Queries) GetFeed(ctx context.Context, url string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeed, url)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Name,
		&i.Url,
		&i.LastFetchedAt,
	)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name FROM feeds
LEFT JOIN users
    ON feeds.user_id = users.id
`

type GetFeedsRow struct {
	Name   string
	Url    string
	Name_2 sql.NullString
}

func (q *Queries) GetFeeds(ctx context.Context) ([]GetFeedsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedsRow
	for rows.Next() {
		var i GetFeedsRow
		if err := rows.Scan(&i.Name, &i.Url, &i.Name_2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNextFeedToFetch = `-- name: GetNextFeedToFetch :one
SELECT id, created_at, updated_at, user_id, name, url, last_fetched_at FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1
`

func (q *Queries) GetNextFeedToFetch(ctx context.Context) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getNextFeedToFetch)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Name,
		&i.Url,
		&i.LastFetchedAt,
	)
	return i, err
}

const markFeedFetched = `-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = $1, updated_at = $1
WHERE id = $2
RETURNING id, created_at, updated_at, user_id, name, url, last_fetched_at
`

type MarkFeedFetchedParams struct {
	LastFetchedAt sql.NullTime
	ID            uuid.UUID
}

func (q *Queries) MarkFeedFetched(ctx context.Context, arg MarkFeedFetchedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, markFeedFetched, arg.LastFetchedAt, arg.ID)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Name,
		&i.Url,
		&i.LastFetchedAt,
	)
	return i, err
}
