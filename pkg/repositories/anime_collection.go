package repositories

import (
	"context"

	"github.com/ilfey/hikilist-go/internal/postgres"
	animecollection "github.com/ilfey/hikilist-go/pkg/models/anime_collection"
)

type AnimeCollection interface {
	WithTx(tx DBRW) AnimeCollection

	AddAnimes(ctx context.Context, aam *animecollection.AddAnimesModel) error
	RemoveAnimes(ctx context.Context, ram *animecollection.RemoveAnimesModel) error
}

type AnimeCollectionImpl struct {
	db DBRW
}

func NewAnimeCollection(db DBRW) AnimeCollection {
	return &AnimeCollectionImpl{
		db: db,
	}
}

func (r *AnimeCollectionImpl) WithTx(tx DBRW) AnimeCollection {
	return &AnimeCollectionImpl{
		db: tx,
	}
}

func (r *AnimeCollectionImpl) AddAnimes(ctx context.Context, aam *animecollection.AddAnimesModel) error {
	return r.db.RunTx(ctx, func(tx postgres.Tx) error {
		// Get collection id.
		sql, args, err := r.GetCollectionIdSQL(aam.UserID, aam.CollectionID)
		if err != nil {
			return err
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&aam.CollectionID)
		if err != nil {
			return err
		}

		// Add animes.
		sql, args, err = r.AddAnimesSQL(aam)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, sql, args...)

		return err
	})
}

func (r *AnimeCollectionImpl) RemoveAnimes(ctx context.Context, ram *animecollection.RemoveAnimesModel) error {
	return r.db.RunTx(ctx, func(tx postgres.Tx) error {
		// Get collection id.
		sql, args, err := r.GetCollectionIdSQL(ram.UserID, ram.CollectionID)
		if err != nil {
			return err
		}

		err = tx.QueryRow(ctx, sql, args...).Scan(&ram.CollectionID)
		if err != nil {
			return err
		}

		// Remove animes.
		sql, args, err = r.RemoveAnimesSQL(ram)
		if err != nil {
			return err
		}

		_, err = r.db.Exec(ctx, sql, args...)

		return err
	})
}
