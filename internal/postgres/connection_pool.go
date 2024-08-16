package postgres

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rotisserie/eris"
)

// ConnectionPool is struct with connection pool
type ConnectionPool struct {
	*pgxpool.Pool
}

// Query sql
func (p *ConnectionPool) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return p.Pool.Query(ctx, sql, args...)
}

// QueryRow sql
func (p *ConnectionPool) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.Pool.QueryRow(ctx, sql, args...)
}

// Exec sql with context
func (p *ConnectionPool) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	return p.Pool.Exec(ctx, sql, arguments...)
}

// Begin return new transaction with context
func (p *ConnectionPool) Begin(ctx context.Context) (Tx, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "create transaction")
	}

	return &Transaction{
		Tx: tx,
		mu: &sync.Mutex{},
	}, nil
}

// RunTx exec sql with transaction
func (p *ConnectionPool) RunTx(ctx context.Context, fn func(tx Tx) error) error {
	var err error
	tx, err := p.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		p := recover()
		switch {
		case p != nil:
			// a panic occurred, rollback and repanic
			_ = tx.Rollback()
			panic(p)
		case err != nil:
			// something went wrong, rollback
			_ = tx.Rollback()
		default:
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)

	return err
}

func (p *ConnectionPool) Statistics() *pgxpool.Stat {
	return p.Stat()
}

// Close ...
func (p *ConnectionPool) Close() {
	p.Pool.Close()
}
