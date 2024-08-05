package drivers

import (
	"context"
	"sync"

	"github.com/ilfey/hikilist-go/internal/orm/database"
	"github.com/jmoiron/sqlx"
)

type SQLXDatabase struct {
	*sqlx.DB
}

func NewSQLX(base *sqlx.DB) database.DB {
	return &SQLXDatabase{
		DB: base,
	}
}

func (d *SQLXDatabase) Exec(ctx context.Context, sql string, args ...any) (database.Result, error) {
	return d.ExecContext(ctx, sql, args...)
}

func (d *SQLXDatabase) Query(ctx context.Context, sql string, args ...any) (database.Rows, error) {
	return d.QueryContext(ctx, sql, args...)
}

func (d *SQLXDatabase) QueryRow(ctx context.Context, sql string, args ...any) database.Row {
	return d.QueryRowContext(ctx, sql, args...)
}

func (d *SQLXDatabase) Begin(ctx context.Context) (database.Transaction, error) {
	tx, err := d.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &SQLXTransaction{
		Tx: tx,
		mu: &sync.Mutex{},
	}, nil
}

func (d *SQLXDatabase) RunTx(ctx context.Context, fn func(tx database.Transaction) error) error {
	tx, err := d.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		p := recover()
		switch {
		case p != nil:
			_ = tx.Rollback()

			panic(p)
		case err != nil:
			_ = tx.Rollback()
		default:
			err = tx.Commit()
		}
	}()

	err = fn(tx)

	return err
}
