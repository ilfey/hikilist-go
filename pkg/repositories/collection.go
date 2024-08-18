package repositories

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/internal/postgres"
	"github.com/ilfey/hikilist-go/pkg/models/action"
	"github.com/ilfey/hikilist-go/pkg/models/collection"
)

type Collection interface {
	WithTx(tx DBRW) Collection

	Create(ctx context.Context, cm *collection.CreateModel) error
	Get(ctx context.Context, conds any) (*collection.DetailModel, error)
	Find(ctx context.Context, p *paginate.Paginator, conds any) ([]*collection.ListItemModel, error)
	Count(ctx context.Context, conds any) (uint, error)
	Update(ctx context.Context, um *collection.UpdateModel) error
}

type CollectionImpl struct {
	action Action
	db     DBRW
}

func NewCollection(db DBRW, actionRepo Action) Collection {
	return &CollectionImpl{
		action: actionRepo,
		db:     db,
	}
}

func (r *CollectionImpl) WithTx(tx DBRW) Collection {
	return &CollectionImpl{
		action: r.action,
		db:     tx,
	}
}

func (r *CollectionImpl) Create(ctx context.Context, cm *collection.CreateModel) error {
	return r.db.RunTx(ctx, func(tx postgres.Tx) error {
		sql, args, err := r.CreateSQL(cm)
		if err != nil {
			return err
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&cm.ID)
		if err != nil {
			return err
		}

		// Create action
		actionCm := action.NewCreateCollectionAction(cm.UserID, cm.Title)

		return r.action.WithTx(tx).Create(ctx, actionCm)
	})
}

func (r *CollectionImpl) Get(ctx context.Context, conds any) (*collection.DetailModel, error) {
	sql, args, err := r.GetSQL(conds)
	if err != nil {
		return nil, err
	}

	var dm collection.DetailModel

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		return nil, err
	}

	return &dm, nil
}

func (r *CollectionImpl) Find(ctx context.Context, p *paginate.Paginator, conds any) ([]*collection.ListItemModel, error) {
	sql, args, err := r.FindSQL(p, conds)
	if err != nil {
		return nil, err
	}

	var items []*collection.ListItemModel

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *CollectionImpl) Count(ctx context.Context, conds any) (uint, error) {
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

func (r *CollectionImpl) Update(ctx context.Context, um *collection.UpdateModel) error {
	sql, args, err := r.UpdateSQL(um)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, sql, args...)

	return err
}
