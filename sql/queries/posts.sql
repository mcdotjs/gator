-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, feed_id, description, published_at )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;


-- name: GetPostsForUser :many
SELECT * FROM posts 
WHERE posts.feed_id IN (
  SELECT feed_id FROM feed_follows WHERE feed_follows.user_id = $1
)
ORDER BY published_at DESC
LIMIT $2;


-- name: GetPostsForUserTroughJoin :many
SELECT posts.* FROM posts 
INNER JOIN feed_follows ON feed_follows.user_id = $1 AND posts.feed_id = feed_follows.feed_id
ORDER BY published_at DESC
LIMIT $2;


