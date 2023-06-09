-- name: CreateArticleCategory :one
INSERT INTO articles_categories (
  article_id,
  category_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: DeleteArticleCategory :exec
DELETE FROM articles_categories
WHERE article_id = $1 AND category_id = $2;

-- name: GetArticleCategory :one
SELECT * FROM articles_categories
WHERE article_id = $1 AND category_id = $2;

-- name: GetArticlesByCategoryID :many
SELECT * FROM articles_categories
WHERE category_id = $1;

-- name: GetCategoriesByArticleID :many
SELECT * FROM articles_categories
WHERE article_id = $1;