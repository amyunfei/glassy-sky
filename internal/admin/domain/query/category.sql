-- name: CreateCategory :one
INSERT INTO categories (
  name,
  parent_id,
  color
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- name: UpdateCategory :one
UPDATE categories SET
name = $2,
parent_id = $3,
color = $4
WHERE id = $1
RETURNING *;

-- name: CountCategory :one
SELECT count(*) FROM categories;

-- name: ListCategory :many
SELECT * FROM categories
ORDER BY id
LIMIT $1
OFFSET $2;