-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetUserFeeds :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: MarkFeedFetched :one
UPDATE feeds 
SET updated_at = NOW(), last_fetched_at= NOW() 
WHERE id = $1 AND user_id = $2 
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds 
WHERE user_id = $1
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
