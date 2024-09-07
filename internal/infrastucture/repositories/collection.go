package repositories

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
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

	//ErrCollectionsDeleteFailed = errtype.NewInternalRepositoryError("unable to delete collection")

	ErrCollectionGetFailed = errtype.NewInternalRepositoryError("unable to get collection")
)

type Collection struct {
	log    loggerInterface.Logger
	action repositoryInterface.Action
	db     postgres.RW
}

func NewCollection(log loggerInterface.Logger, db postgres.RW, actionRepo repositoryInterface.Action) *Collection {
	return &Collection{
		log:    log,
		action: actionRepo,
		db:     db,
	}
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
			return r.log.CriticalPropagate(err)
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&cm.ID)
		if err != nil {
			r.log.Log(err)

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
		return nil, r.log.CriticalPropagate(err)
	}

	var dm agg.CollectionDetail

	err = pgxscan.Get(ctx, r.db, &dm, sql, args...)
	if err != nil {
		r.log.Log(err)

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
		return nil, r.log.CriticalPropagate(err)
	}

	var items []*agg.CollectionListItem

	err = pgxscan.Select(ctx, r.db, &items, sql, args...)
	if err != nil {
		r.log.Log(err)

		return nil, ErrCollectionsFindFailed
	}

	return items, nil
}

func (r *Collection) Count(ctx context.Context, conds any) (uint64, error) {
	sql, args, err := r.CountSQL(conds)
	if err != nil {
		return 0, r.log.CriticalPropagate(err)
	}

	var count uint64

	err = r.db.QueryRow(ctx, sql, args...).Scan(&count)
	if err != nil {
		r.log.Log(err)

		return 0, ErrCollectionsCountFailed
	}

	return count, nil
}

func (r *Collection) Update(ctx context.Context, um *dto.CollectionUpdateRequestDTO) error {
	sql, args, err := r.UpdateSQL(um)
	if err != nil {
		return r.log.CriticalPropagate(err)
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Log(err)

		return ErrCollectionUpdateFailed
	}

	return err
}
