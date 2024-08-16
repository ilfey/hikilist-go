package repositories

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/pkg/models/action"
)

type Action interface {
	WithTx(tx DBRW) Action

	Create(ctx context.Context, cm *action.CreateModel) error
	FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*action.ListItemModel, error)
	Count(ctx context.Context, conds any) (uint, error)
}

type ActionImpl struct {
	db DBRW
}

func NewAction(db DBRW) Action {
	return &ActionImpl{
		db: db,
	}
}

func (r *ActionImpl) WithTx(tx DBRW) Action {
	return &ActionImpl{
		db: tx,
	}
}

func (r *ActionImpl) Create(ctx context.Context, cm *action.CreateModel) error {
	sql, args, err := r.CreateSQL(cm)
	if err != nil {
		return err
	}

	return r.db.QueryRow(ctx, sql, args...).Scan(&cm.ID)
}

func (r *ActionImpl) FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*action.ListItemModel, error) {
	sql, args, err := r.FindWithPaginatorSQL(p, conds)
	if err != nil {
		return nil, err
	}

	var items []*action.ListItemModel

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ActionImpl) Count(ctx context.Context, conds any) (uint, error) {
	sql, args, err := r.CountSQL(conds)
	if err != nil {
		return 0, err
	}

	var count uint

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
