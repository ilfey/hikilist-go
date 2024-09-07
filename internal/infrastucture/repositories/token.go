package repositories

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
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
	log loggerInterface.Logger
	db  postgres.RW
}

func NewToken(container diInterface.ServiceContainer) (*Token, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	db, err := container.GetPostgresDatabase()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &Token{
		log: log,
		db:  db,
	}, nil
}

func (r *Token) WithTx(tx postgres.RW) repositoryInterface.Token {
	return &Token{
		log: r.log,
		db:  tx,
	}
}

func (r *Token) Create(ctx context.Context, cm *agg.TokenCreate) error {
	sql, args, err := r.CreateSQL(cm)
	if err != nil {
		return r.log.Propagate(err)
	}

	err = r.db.QueryRow(ctx, sql, args...).Scan(&cm.ID)
	if err != nil {
		r.log.Error(err)

		return ErrTokenCreateFailed
	}

	return nil
}

func (r *Token) Get(ctx context.Context, conds any) (*agg.TokenDetail, error) {
	sql, args, err := r.GetSQL(conds)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var dm agg.TokenDetail

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		r.log.Error(err)

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
		return false, r.log.Propagate(err)
	}

	var id uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		r.log.Error(err)

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
		return r.log.Propagate(err)
	}

	var id uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		r.log.Error(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return ErrTokenNotFoundById
		}

		return ErrTokenDeleteFailed
	}

	return nil
}
