package repositories

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/pkg/models/anime"
)

type Anime interface {
	WithTx(tx DBRW) Anime

	Create(ctx context.Context, cm *anime.CreateModel) error
	Get(ctx context.Context, conds any) (*anime.DetailModel, error)
	FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*anime.ListItemModel, error)
	FindFromCollectionWithPaginator(ctx context.Context, p *paginate.Paginator, userId, collectionId uint) ([]*anime.ListItemModel, error)
	CountInCollection(ctx context.Context, userId, collectionId uint) (uint, error)
	Count(ctx context.Context, conds any) (uint, error)
}

type AnimeImpl struct {
	db DBRW
}

func NewAnime(db DBRW) Anime {
	return &AnimeImpl{
		db: db,
	}
}

func (r *AnimeImpl) WithTx(tx DBRW) Anime {
	return &AnimeImpl{
		db: tx,
	}
}

func (r *AnimeImpl) Create(ctx context.Context, cm *anime.CreateModel) error {
	sql, args, err := r.CreateSQL(cm)
	if err != nil {
		return err
	}

	return r.db.QueryRow(ctx, sql, args...).Scan(&cm.ID)
}

func (r *AnimeImpl) Get(ctx context.Context, conds any) (*anime.DetailModel, error) {
	sql, args, err := r.GetSQL(conds)
	if err != nil {
		return nil, err
	}

	var dm anime.DetailModel

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		return nil, err
	}

	return &dm, nil
}

func (r *AnimeImpl) FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*anime.ListItemModel, error) {
	sql, args, err := r.FindWithPaginatorSQL(p, conds)
	if err != nil {
		return nil, err
	}

	var items []*anime.ListItemModel

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *AnimeImpl) FindFromCollectionWithPaginator(ctx context.Context, p *paginate.Paginator, userId uint, collectionId uint) ([]*anime.ListItemModel, error) {
	sql, args, err := r.FindFromCollectionWithPaginatorSQL(p, userId, collectionId)
	if err != nil {
		return nil, err
	}

	var items []*anime.ListItemModel

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *AnimeImpl) CountInCollection(ctx context.Context, userId, collectionId uint) (uint, error) {
	sql, args, err := r.CountInCollectionSQL(userId, collectionId)
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

func (r *AnimeImpl) Count(ctx context.Context, conds any) (uint, error) {
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
