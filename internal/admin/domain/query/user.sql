-- name: CreateUser :one
INSERT INTO users (
  username, password, email, nickname
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET
nickname = COALESCE(sqlc.narg(nickname), nickname),
avatar = COALESCE(sqlc.narg(avatar), avatar),
password = COALESCE(sqlc.narg(password), password)
WHERE id = $1
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: CountUser :one
SELECT count(*) FROM users
WHERE (@username::text = '' OR username LIKE '%' || @username || '%') AND
(@nickname::text = '' OR nickname LIKE '%' || @nickname || '%');

-- name: ListUser :many
SELECT * FROM users
WHERE (@username::text = '' OR username LIKE '%' || @username || '%') AND
(@nickname::text = '' OR nickname LIKE '%' || @nickname || '%')
ORDER BY id
LIMIT $1
OFFSET $2;