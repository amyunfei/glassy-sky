// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgresql

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int64
	Name      string
	ParentID  int64
	Color     int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type Label struct {
	ID        int64
	Name      string
	Color     int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
