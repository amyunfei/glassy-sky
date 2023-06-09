-- name: CreateArticleLabel :one
INSERT INTO articles_labels (
  article_id,
  label_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: DeleteArticleLabel :exec
DELETE FROM articles_labels
WHERE article_id = $1 AND label_id = $2;

-- name: GetArticleLabel :one
SELECT * FROM articles_labels
WHERE article_id = $1 AND label_id = $2;

-- name: GetArticlesByLabelID :many
SELECT * FROM articles_labels
WHERE label_id = $1;

-- name: GetLabelsByArticleID :many
SELECT * FROM articles_labels
WHERE article_id = $1;