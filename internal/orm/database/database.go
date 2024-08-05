package database

import (
	"context"
)

type DB interface {
	Read
	Write
}

type Read interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
}

type Write interface {
	Exec(ctx context.Context, sql string, args ...any) (Result, error)
	Begin(ctx context.Context) (Transaction, error)
	RunTx(ctx context.Context, fn func(tx Transaction) error) error
}

type Transaction interface {
	Read
	Write

	Rollback() error
	RollbackCtx(ctx context.Context) error
	Commit() error
	CommitCtx(ctx context.Context) error
	Close() error
	CloseCtx(ctx context.Context) error
}

type Rows interface {
	Row

	Next() bool
	Close() error
	Columns() ([]string, error)
	// ColumnTypes() ([]*sql.ColumnType, error)
}

type Row interface {
	Err() error
	Scan(dest ...any) error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
