-- name: GetPostsForUser :many
SELECT * FROM posts
LEFT JOIN feeds 
    ON posts.feed_id = feeds.id
LEFT JOIN users
    ON feeds.user_id = users.id
ORDER BY posts.published_at DESC
LIMIT $1;
