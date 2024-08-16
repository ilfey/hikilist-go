package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Read interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Write interface {
	Exec(ctx context.Context, sql string, args ...any) (commandTag pgconn.CommandTag, err error)
	Begin(ctx context.Context) (Tx, error)
	RunTx(ctx context.Context, fn func(tx Tx) error) error
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
}

// DB interface for work with DB
type DB interface {
	Read
	Write

	Statistics() *pgxpool.Stat
	Close()
}

type Tx interface {
	Read
	Write

	Commit() error
	CommitCtx(ctx context.Context) error
	Rollback() error
	RollbackCtx(ctx context.Context) error
}
