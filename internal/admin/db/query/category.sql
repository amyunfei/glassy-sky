-- name: CreateCategory :one
INSERT INTO categories (
  name,
  parent_id,
  color
) VALUES (
  $1, $2, $3
) RETURNING *;