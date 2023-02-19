-- name: CreateLabel :one
INSERT INTO labels (
  name,
  color
) VALUES (
  $1, $2
) RETURNING *;

-- name: DeleteLabel :exec
DELETE FROM labels
WHERE id = $1;

-- name: UpdateLabel :one
UPDATE labels SET
name = COALESCE(sqlc.narg(name), name),
color = COALESCE(sqlc.narg(color), color)
WHERE id = $1
RETURNING *;

-- name: GetLabel :one
SELECT * FROM labels
WHERE id = $1;

-- name: CountLabel :one
SELECT count(*) FROM labels
WHERE (@name::text = '' OR name LIKE '%' || @name || '%');

-- name: ListLabel :many
SELECT * FROM labels
WHERE (@name::text = '' OR name LIKE '%' || @name || '%')
ORDER BY id
LIMIT $1
OFFSET $2;