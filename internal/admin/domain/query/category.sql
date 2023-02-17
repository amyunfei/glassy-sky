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
name = COALESCE(sqlc.narg(name), name),
parent_id = COALESCE(sqlc.narg(parent_id), parent_id),
color = COALESCE(sqlc.narg(color), color)
WHERE id = $1
RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1;

-- name: CountCategory :one
SELECT count(*) FROM categories
WHERE name LIKE sqlc.narg('name');

-- name: ListCategory :many
SELECT * FROM categories
WHERE name LIKE COALESCE(sqlc.narg(name), name)
ORDER BY id
LIMIT $1
OFFSET $2;