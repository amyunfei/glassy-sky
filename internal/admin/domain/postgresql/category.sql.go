// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: category.sql

package postgresql

import (
	"context"
)

const countCategory = `-- name: CountCategory :one
SELECT count(*) FROM categories
`

func (q *Queries) CountCategory(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countCategory)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (
  name,
  parent_id,
  color
) VALUES (
  $1, $2, $3
) RETURNING id, name, parent_id, color, created_at, updated_at, deleted_at
`

type CreateCategoryParams struct {
	Name     string
	ParentID int64
	Color    int32
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, arg.Name, arg.ParentID, arg.Color)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ParentID,
		&i.Color,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const listCategory = `-- name: ListCategory :many
SELECT id, name, parent_id, color, created_at, updated_at, deleted_at FROM categories
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCategoryParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListCategory(ctx context.Context, arg ListCategoryParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategory, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ParentID,
			&i.Color,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories SET
name = $2,
parent_id = $3,
color = $4
WHERE id = $1
RETURNING id, name, parent_id, color, created_at, updated_at, deleted_at
`

type UpdateCategoryParams struct {
	ID       int64
	Name     string
	ParentID int64
	Color    int32
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, updateCategory,
		arg.ID,
		arg.Name,
		arg.ParentID,
		arg.Color,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ParentID,
		&i.Color,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
