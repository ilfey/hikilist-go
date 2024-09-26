package repositories

import (
	"context"
	"github.com/ilfey/hikilist-database/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/postgres"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type Action struct {
	log loggerInterface.Logger
	db  postgres.RW
}

var (
	ErrActionCreateFailed = errtype.NewInternalRepositoryError("unable to store action")
	ErrActionsFindFailed  = errtype.NewInternalRepositoryError("unable to find actions")
	ErrActionsCountFailed = errtype.NewInternalRepositoryError("unable to count actions")
)

func NewAction(container diInterface.AppContainer) (*Action, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	db, err := container.GetPostgresDatabase()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &Action{
		log: log,
		db:  db,
	}, nil
}

func (r *Action) WithTx(tx postgres.RW) repositoryInterface.Action {
	return &Action{
		db: tx,
	}
}

func (r *Action) Create(ctx context.Context, cm *dto.ActionCreateRequestDTO) error {
	sql, args, err := r.CreateSQL(cm)
	if err != nil {
		return r.log.Propagate(err)
	}

	err = r.db.QueryRow(ctx, sql, args...).Scan(&cm.ID)
	if err != nil {
		r.log.Error(err)

		return ErrActionCreateFailed
	}

	return nil
}

func (r *Action) Find(ctx context.Context, p *dto.ActionListRequestDTO, conds any) ([]*agg.ActionListItem, error) {
	sql, args, err := r.FindWithPaginatorSQL(p, conds)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var items []*agg.ActionListItem

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		r.log.Error(err)

		return nil, ErrActionsFindFailed
	}

	return items, nil
}

func (r *Action) Count(ctx context.Context, conds any) (uint64, error) {
	sql, args, err := r.CountSQL(conds)
	if err != nil {
		return 0, r.log.Propagate(err)
	}

	var count uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		r.log.Error(err)

		return 0, ErrActionsCountFailed
	}

	return count, nil
}
