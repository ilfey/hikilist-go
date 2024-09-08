package repositories

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
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
	ErrCollectionCreateFailed = errtype.NewInternalRepositoryError("unable to store collection")
	ErrCollectionsFindFailed  = errtype.NewInternalRepositoryError("unable to find collections")
	ErrCollectionNotFoundById = errtype.NewEntityNotFoundError("database", "collection", "id")
	ErrCollectionsCountFailed = errtype.NewInternalRepositoryError("unable to count collections")
	ErrCollectionUpdateFailed = errtype.NewInternalRepositoryError("unable to update collection")
	ErrCollectionDeleteFailed = errtype.NewInternalRepositoryError("unable to delete collection")
	ErrCollectionGetFailed    = errtype.NewInternalRepositoryError("unable to get collection")
)

type Collection struct {
	log    loggerInterface.Logger
	action repositoryInterface.Action
	db     postgres.RW
}

func NewCollection(container diInterface.AppContainer) (*Collection, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	db, err := container.GetPostgresDatabase()
	if err != nil {
		return nil, log.Propagate(err)
	}

	action, err := container.GetActionRepository()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &Collection{
		log:    log,
		action: action,
		db:     db,
	}, nil
}

func (r *Collection) WithTx(tx postgres.RW) repositoryInterface.Collection {
	return &Collection{
		log:    r.log,
		action: r.action,
		db:     tx,
	}
}

func (r *Collection) Create(ctx context.Context, cm *dto.CollectionCreateRequestDTO) error {
	return r.db.RunTx(ctx, func(tx postgres.Tx) error {
		// Create collection.
		sql, args, err := r.CreateSQL(cm)
		if err != nil {
			return r.log.Propagate(err)
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&cm.ID)
		if err != nil {
			r.log.Error(err)

			return ErrCollectionCreateFailed
		}

		// Create action
		actionCm := dto.NewCreateCollectionAction(cm.UserID, cm.Title)

		return r.action.WithTx(tx).Create(ctx, actionCm)
	})
}

func (r *Collection) Get(ctx context.Context, conds any) (*agg.CollectionDetail, error) {
	sql, args, err := r.GetSQL(conds)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var dm agg.CollectionDetail

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		r.log.Error(err)

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrCollectionNotFoundById
		}

		return nil, ErrCollectionGetFailed
	}

	return &dm, nil
}

func (r *Collection) Find(ctx context.Context, dto *dto.CollectionListRequestDTO, conds any) ([]*agg.CollectionListItem, error) {
	sql, args, err := r.FindSQL(dto, conds)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var items []*agg.CollectionListItem

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		r.log.Error(err)

		return nil, ErrCollectionsFindFailed
	}

	return items, nil
}

func (r *Collection) FindUserPublicCollectionList(ctx context.Context, req *dto.UserCollectionListRequestDTO) ([]*agg.CollectionListItem, error) {
	sql, args, err := r.FindUserPublicCollectionListSQL(req)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var items []*agg.CollectionListItem

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		r.log.Error(err)

		return nil, ErrCollectionsFindFailed
	}

	return items, nil
}

func (r *Collection) FindUserCollectionList(ctx context.Context, req *dto.UserCollectionListRequestDTO) ([]*agg.CollectionListItem, error) {
	sql, args, err := r.FindUserCollectionListSQL(req)
	if err != nil {
		return nil, r.log.Propagate(err)
	}

	var items []*agg.CollectionListItem

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		r.log.Error(err)

		return nil, ErrCollectionsFindFailed
	}

	return items, nil
}

func (r *Collection) Count(ctx context.Context, conds any) (uint64, error) {
	sql, args, err := r.CountSQL(conds)
	if err != nil {
		return 0, r.log.Propagate(err)
	}

	var count uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		r.log.Error(err)

		return 0, ErrCollectionsCountFailed
	}

	return count, nil
}

func (r *Collection) CountUserPublicCollection(ctx context.Context, req *dto.UserCollectionListRequestDTO) (uint64, error) {
	sql, args, err := r.CountUserPublicCollectionSQL(req)
	if err != nil {
		return 0, r.log.Propagate(err)
	}

	var count uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		r.log.Error(err)

		return 0, ErrCollectionsCountFailed
	}

	return count, nil
}

func (r *Collection) CountUserCollection(ctx context.Context, req *dto.UserCollectionListRequestDTO) (uint64, error) {
	sql, args, err := r.CountUserCollectionSQL(req)
	if err != nil {
		return 0, r.log.Propagate(err)
	}

	var count uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		r.log.Error(err)

		return 0, ErrCollectionsCountFailed
	}

	return count, nil
}

func (r *Collection) Update(ctx context.Context, um *agg.CollectionDetail) error {
	sql, args, err := r.UpdateSQL(um)
	if err != nil {
		return r.log.Propagate(err)
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Error(err)

		return ErrCollectionUpdateFailed
	}

	return err
}

func (r *Collection) Delete(ctx context.Context, req *dto.CollectionDeleteRequestDTO) error {
	sql, args, err := r.DeleteSQL(req)
	if err != nil {
		return r.log.Propagate(err)
	}

	var id uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrCollectionNotFoundById
		}

		return ErrCollectionDeleteFailed
	}

	return err
}
