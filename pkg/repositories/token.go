package repositories

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/pkg/models/token"
)

type Token interface {
	WithTx(tx DBRW) Token

	Create(ctx context.Context, cm *token.CreateModel) error
	Get(ctx context.Context, conds any) (*token.DetailModel, error)
	Delete(ctx context.Context, conds any) error
}

type TokenImpl struct {
	db DBRW
}

func NewToken(db DBRW) Token {
	return &TokenImpl{
		db: db,
	}
}

func (TokenImpl) WithTx(tx DBRW) Token {
	return &TokenImpl{
		db: tx,
	}
}

func (r *TokenImpl) Create(ctx context.Context, cm *token.CreateModel) error {
	sql, args, err := r.CreateSQL(cm)
	if err != nil {
		return err
	}

	return r.db.QueryRow(ctx, sql, args...).Scan(&cm.ID)
}

func (r *TokenImpl) Get(ctx context.Context, conds any) (*token.DetailModel, error) {
	sql, args, err := r.GetSQL(conds)
	if err != nil {
		return nil, err
	}

	var dm token.DetailModel

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		return nil, err
	}

	return &dm, nil
}

func (r *TokenImpl) Delete(ctx context.Context, conds any) error {
	sql, args, err := r.DeleteSQL(conds)
	if err != nil {
		return err
	}

	var id uint

	return r.db.QueryRow(ctx, sql, args...).Scan(&id)
}
