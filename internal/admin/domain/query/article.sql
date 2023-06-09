-- name: CreateArticle :one
INSERT INTO articles (
  title,
  excerpt,
  content,
  user_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = $1;

-- name: GetArticle :one
SELECT * FROM articles
WHERE id = $1;

-- name: CountArticle :one
SELECT count(*) FROM articles
WHERE (@title::text = '' OR title LIKE '%' || @title || '%');

-- name: ListArticle :many
SELECT * FROM articles
WHERE (@title::text = '' OR title LIKE '%' || @title || '%')
ORDER BY id
LIMIT $1
OFFSET $2;