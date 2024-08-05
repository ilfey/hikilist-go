package drivers

import (
	"context"
	"fmt"
	"sync"

	"github.com/ilfey/hikilist-go/internal/orm/database"
	"github.com/jmoiron/sqlx"
	"github.com/rotisserie/eris"
)

type SQLXTransaction struct {
	*sqlx.Tx
	mu        *sync.Mutex
	savePoint uint8
}

func (t *SQLXTransaction) Exec(ctx context.Context, sql string, args ...any) (database.Result, error) {
	return t.ExecContext(ctx, sql, args...)
}

func (t *SQLXTransaction) Query(ctx context.Context, sql string, args ...any) (database.Rows, error) {
	return t.QueryContext(ctx, sql, args...)
}

func (t *SQLXTransaction) QueryRow(ctx context.Context, sql string, args ...any) database.Row {
	return t.QueryRowContext(ctx, sql, args...)
}

func (t *SQLXTransaction) Begin(ctx context.Context) (database.Transaction, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.savePoint++

	sql := fmt.Sprintf("SAVEPOINT savepoint_%d", t.savePoint)
	_, err := t.Exec(ctx, sql)
	if err != nil {
		return nil, eris.Wrap(err, "create savepoint")
	}

	return t, nil
}

func (t *SQLXTransaction) Rollback() error {
	return t.RollbackCtx(context.Background())
}

func (t *SQLXTransaction) RollbackCtx(ctx context.Context) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.savePoint == 0 {
		return t.Tx.Rollback()
	}

	sql := fmt.Sprintf("ROLLBACK TO SAVEPOINT savepoint_%d", t.savePoint)
	_, err := t.Exec(ctx, sql)
	t.savePoint--
	if err != nil {
		return eris.Wrap(err, "rollback to savepoint")
	}

	return nil
}

func (t *SQLXTransaction) Commit() error {
	return t.CommitCtx(context.Background())
}

func (t *SQLXTransaction) CommitCtx(ctx context.Context) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.savePoint == 0 {
		return t.Tx.Commit()
	}

	t.savePoint--

	return nil
}

func (d *SQLXTransaction) RunTx(ctx context.Context, fn func(tx database.Transaction) error) error {
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

func (t *SQLXTransaction) Close() error {
	return t.CloseCtx(context.Background())
}

func (t *SQLXTransaction) CloseCtx(ctx context.Context) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	_ = t.Tx.Rollback()

	return nil
}
