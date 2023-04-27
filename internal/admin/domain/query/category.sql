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
WHERE (@name::text = '' OR name LIKE '%' || @name || '%');

-- name: ListCategory :many
SELECT * FROM categories
WHERE (@name::text = '' OR name LIKE '%' || @name || '%')
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: TreeCategory :many
WITH RECURSIVE categoryTree AS (
  SELECT id, name, parent_id, color, 1::SMALLINT AS level
  FROM categories
  WHERE parent_id = 0
  UNION ALL
  SELECT c.id, c.name, c.parent_id, c.color, t.level::SMALLINT + 1::SMALLINT
  FROM categories c
  JOIN categoryTree t ON c.parent_id = t.id
) SELECT * FROM categoryTree;