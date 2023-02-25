package postgresql

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	Querier
	ExecTx(ctx context.Context, fn func(Querier) error) error
}

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Repository {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func (store *Store) ExecTx(ctx context.Context, fn func(Querier) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(store.Queries)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
