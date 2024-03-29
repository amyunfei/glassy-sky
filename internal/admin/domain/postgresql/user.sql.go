// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package postgresql

import (
	"context"
	"database/sql"
)

const countUser = `-- name: CountUser :one
SELECT count(*) FROM users
WHERE ($1::text = '' OR username LIKE '%' || $1 || '%') AND
($2::text = '' OR nickname LIKE '%' || $2 || '%')
`

type CountUserParams struct {
	Username string
	Nickname string
}

func (q *Queries) CountUser(ctx context.Context, arg CountUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUser, arg.Username, arg.Nickname)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username, password, email, nickname
) VALUES (
  $1, $2, $3, $4
) RETURNING id, username, password, email, nickname, avatar, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	Username string
	Password string
	Email    string
	Nickname string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.Nickname,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Nickname,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password, email, nickname, avatar, created_at, updated_at, deleted_at FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Nickname,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, password, email, nickname, avatar, created_at, updated_at, deleted_at FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Nickname,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password, email, nickname, avatar, created_at, updated_at, deleted_at FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Nickname,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, username, password, email, nickname, avatar, created_at, updated_at, deleted_at FROM users
WHERE ($3::text = '' OR username LIKE '%' || $3 || '%') AND
($4::text = '' OR nickname LIKE '%' || $4 || '%')
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUserParams struct {
	Limit    int32
	Offset   int32
	Username string
	Nickname string
}

func (q *Queries) ListUser(ctx context.Context, arg ListUserParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUser,
		arg.Limit,
		arg.Offset,
		arg.Username,
		arg.Nickname,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Email,
			&i.Nickname,
			&i.Avatar,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users SET
nickname = COALESCE($2, nickname),
avatar = COALESCE($3, avatar),
password = COALESCE($4, password)
WHERE id = $1
RETURNING id, username, password, email, nickname, avatar, created_at, updated_at, deleted_at
`

type UpdateUserParams struct {
	ID       int64
	Nickname sql.NullString
	Avatar   sql.NullString
	Password sql.NullString
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Nickname,
		arg.Avatar,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Nickname,
		&i.Avatar,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
