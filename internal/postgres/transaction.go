package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rotisserie/eris"
)

// Transaction ...
type Transaction struct {
	pgx.Tx
	mu                *sync.Mutex
	savePointSequence uint8
}

// Exec sql

// Exec sql with context
func (t *Transaction) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	return t.Tx.Exec(ctx, sql, arguments...)
}

// Query sql with context
func (t *Transaction) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return t.Tx.Query(ctx, sql, args...)
}

// QueryRow query row with context
func (t *Transaction) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return t.Tx.QueryRow(ctx, sql, args...)
}

// Begin create savepoint with context
func (t *Transaction) Begin(ctx context.Context) (Tx, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.savePointSequence++

	sql := fmt.Sprintf("SAVEPOINT savepoint_%d", t.savePointSequence)

	_, err := t.Tx.Exec(ctx, sql)
	if err != nil {
		return nil, eris.Wrap(err, "create savepoint")
	}

	return t, nil
}

// Rollback transaction or rollback to savepoint
func (t *Transaction) Rollback() error {
	return t.RollbackCtx(context.Background())
}

// RollbackCtx transaction or rollback to savepoint with context
func (t *Transaction) RollbackCtx(ctx context.Context) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.savePointSequence == 0 {
		return t.Tx.Rollback(context.Background())
	}

	sql := fmt.Sprintf("ROLLBACK TO SAVEPOINT savepoint_%d", t.savePointSequence)

	_, err := t.Tx.Exec(ctx, sql)
	t.savePointSequence--
	if err != nil {
		return eris.Wrap(err, "rollback to savepoint")
	}

	return nil
}

// Commit transaction
func (t *Transaction) Commit() error {
	return t.CommitCtx(context.Background())
}

// CommitCtx transaction with context
func (t *Transaction) CommitCtx(ctx context.Context) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.savePointSequence == 0 {
		return t.Tx.Commit(ctx)
	}

	t.savePointSequence--

	return nil
}

// RunTx exec sql with transaction
func (t *Transaction) RunTx(ctx context.Context, fn func(tx Tx) error) error {
	tx, err := t.Begin(ctx)
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

// Close ...
func (t *Transaction) Close() error {
	return t.CloseCtx(context.Background())
}

// CloseCtx with context
func (t *Transaction) CloseCtx(ctx context.Context) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	_ = t.Tx.Rollback(ctx)

	return nil
}

func (t *Transaction) Statistics() *pgxpool.Stat {
	return nil
}
