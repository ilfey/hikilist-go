package repositories

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/postgres"
)

var (
	ErrAnimeAlreadyInCollection        = errtype.NewEntityAlreadyExistsError("anime", "id")
	ErrAnimeAddToCollectionFailed      = errtype.NewInternalRepositoryError("unable to add anime to collection")
	ErrAnimeRemoveFromCollectionFailed = errtype.NewInternalRepositoryError("unable to remove anime from collection")
)

type AnimeCollection struct {
	log loggerInterface.Logger
	db  postgres.RW
}

func NewAnimeCollection(log loggerInterface.Logger, db postgres.RW) *AnimeCollection {
	return &AnimeCollection{
		log: log,
		db:  db,
	}
}

func (r *AnimeCollection) WithTx(tx postgres.RW) repositoryInterface.AnimeCollection {
	return &AnimeCollection{
		log: r.log,
		db:  tx,
	}
}

func (r *AnimeCollection) AddAnimes(ctx context.Context, addAnimeModel *dto.CollectionAddAnimeRequestDTO) error {
	return r.db.RunTx(ctx, func(tx postgres.Tx) error {
		// Detail collection id.
		sql, args, err := r.GetCollectionIdSQL(addAnimeModel.UserID, addAnimeModel.CollectionID)
		if err != nil {
			return r.log.CriticalPropagate(err)
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&addAnimeModel.CollectionID)
		if err != nil {
			r.log.Log(err)

			if postgres.IsNotFound(err) {
				return ErrCollectionNotFoundById
			}

			return ErrCollectionGetFailed
		}

		// Add animes.
		sql, args, err = r.AddAnimesSQL(addAnimeModel)
		if err != nil {
			return r.log.CriticalPropagate(err)
		}

		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			r.log.Log(err)

			if postgres.PgErrCodeEquals(err, postgres.UniqueViolation) {
				return ErrAnimeAlreadyInCollection
			}

			if postgres.PgErrCodeEquals(err, postgres.ForeignKeyViolation) {
				return ErrAnimeNotFoundById
			}

			return ErrAnimeAddToCollectionFailed
		}

		return nil
	})
}

func (r *AnimeCollection) RemoveAnimes(ctx context.Context, removeAnimeModel *dto.CollectionRemoveAnimeRequestDTO) error {
	return r.db.RunTx(ctx, func(tx postgres.Tx) error {
		// Detail collection id.
		sql, args, err := r.GetCollectionIdSQL(removeAnimeModel.UserID, removeAnimeModel.CollectionID)
		if err != nil {
			return r.log.CriticalPropagate(err)
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&removeAnimeModel.CollectionID)
		if err != nil {
			r.log.Log(err)

			if postgres.IsNotFound(err) {
				return ErrCollectionNotFoundById
			}

			return ErrCollectionGetFailed
		}

		// Remove animes.
		sql, args, err = r.RemoveAnimesSQL(removeAnimeModel)
		if err != nil {
			return r.log.CriticalPropagate(err)
		}

		_, err = r.db.Exec(ctx, sql, args...)
		if err != nil {
			r.log.Log(err)

			return ErrAnimeRemoveFromCollectionFailed
		}

		return nil
	})
}
