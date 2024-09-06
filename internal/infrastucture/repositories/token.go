package repositories

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/georgysavva/scany/v2/pgxscan"
)

var (
	ErrTokenCreateFailed    = errtype.NewInternalRepositoryError("unable to store token")
	ErrTokenGetFailed       = errtype.NewInternalRepositoryError("unable to get token")
	ErrTokenNotFoundById    = errtype.NewEntityNotFoundError("database", "token", "id")
	ErrTokenNotFoundByToken = errtype.NewEntityNotFoundError("database", "token", "token")
	ErrTokenDeleteFailed    = errtype.NewInternalRepositoryError("unable to delete tokens")
)

type Token struct {
	logger loggerInterface.Logger
	db     postgres.RW
}

func NewToken(logger loggerInterface.Logger, db postgres.RW) *Token {
	return &Token{
		logger: logger,
		db:     db,
	}
}

func (r *Token) WithTx(tx postgres.RW) repositoryInterface.Token {
	return &Token{
		logger: r.logger,
		db:     tx,
	}
}

func (r *Token) Create(ctx context.Context, cm *agg.TokenCreate) error {
	sql, args, err := r.CreateSQL(cm)
	if err != nil {
		return r.logger.CriticalPropagate(err)
	}

	err = r.db.QueryRow(ctx, sql, args...).Scan(&cm.ID)
	if err != nil {
		r.logger.Log(err)

		return ErrTokenCreateFailed
	}

	return nil
}

func (r *Token) Get(ctx context.Context, conds any) (*agg.TokenDetail, error) {
	sql, args, err := r.GetSQL(conds)
	if err != nil {
		return nil, r.logger.CriticalPropagate(err)
	}

	var dm agg.TokenDetail

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		r.logger.Log(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTokenNotFoundById
		}

		return nil, ErrTokenGetFailed
	}

	return &dm, nil
}

func (r *Token) Has(ctx context.Context, token string) (bool, error) {
	sql, args, err := r.HasSQL(token)
	if err != nil {
		return false, r.logger.CriticalPropagate(err)
	}

	var id uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		r.logger.Log(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, ErrTokenNotFoundByToken
	}

	return true, nil
}

func (r *Token) Delete(ctx context.Context, conds any) error {
	sql, args, err := r.DeleteSQL(conds)
	if err != nil {
		return r.logger.CriticalPropagate(err)
	}

	var id uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		r.logger.Log(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return ErrTokenNotFoundById
		}

		return ErrTokenDeleteFailed
	}

	return nil
}
