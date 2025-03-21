-- name: CreateFeedFollow :many
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows 
        (
            id,     
            created_at, 
            updated_at, 
            user_id,
            feed_id
        )
        VALUES (
            $1,     
            $2, 
            $3, 
            $4,
            $5
        )
        RETURNING *
)
SELECT 
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
 FROM inserted_feed_follow
    JOIN feeds 
        ON feeds.id = inserted_feed_follow.feed_id
    JOIN users
        ON users.id = inserted_feed_follow.user_id;

-- name: GetFeedFollowsForUser :many
WITH feed_follow_item AS (
    SELECT * FROM feed_follows
)
SELECT
    feed_follow_item.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follow_item
    JOIN feeds 
        ON feeds.id = feed_follow_item.feed_id
    JOIN users
        ON users.id = feed_follow_item.user_id
WHERE users.id = $1;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows
WHERE feed_id = $1 AND user_id = $2;