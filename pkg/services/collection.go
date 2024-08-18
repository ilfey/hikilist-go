package services

import (
	"context"

	"github.com/ilfey/hikilist-go/internal/paginate"
	animecollection "github.com/ilfey/hikilist-go/pkg/models/anime_collection"
	"github.com/ilfey/hikilist-go/pkg/models/collection"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type Collection interface {
	Create(ctx context.Context, cm *collection.CreateModel) error
	Get(ctx context.Context, conds any) (*collection.DetailModel, error)
	GetListModel(ctx context.Context, p *paginate.Paginator, conds any) (*collection.ListModel, error)
	// Find(ctx context.Context, p *paginate.Paginator, conds any) ([]*collection.ListItemModel, error)
	// Count(ctx context.Context, conds any) (uint, error)
	Update(ctx context.Context, um *collection.UpdateModel) error

	AddAnimes(ctx context.Context, aam *animecollection.AddAnimesModel) error
	RemoveAnimes(ctx context.Context, ram *animecollection.RemoveAnimesModel) error
}

type CollectionImpl struct {
	logger logrus.FieldLogger

	animeCollection repositories.AnimeCollection
	collection      repositories.Collection
}

func NewCollection(
	logger logrus.FieldLogger,

	animeCollectionRepo repositories.AnimeCollection,
	collectionRepo repositories.Collection,
) Collection {
	return &CollectionImpl{
		logger: logger,

		animeCollection: animeCollectionRepo,
		collection:      collectionRepo,
	}
}

func (s *CollectionImpl) Create(ctx context.Context, cm *collection.CreateModel) error {
	if cm.IsPublic == nil {
		t := true

		cm.IsPublic = &t
	}

	err := s.collection.Create(ctx, cm)
	if err != nil {
		s.logger.Debugf("Error occurred while creating collection %v", err)

		return err
	}

	return nil
}

func (s *CollectionImpl) Get(ctx context.Context, conds any) (*collection.DetailModel, error) {
	dm, err := s.collection.Get(ctx, conds)
	if err != nil {
		s.logger.Debugf("Error occurred while getting collection %v", err)

		return nil, err
	}

	return dm, nil
}

func (s *CollectionImpl) GetListModel(ctx context.Context, p *paginate.Paginator, conds any) (*collection.ListModel, error) {
	var lm collection.ListModel

	g := errgroup.Group{}

	g.Go(func() error {
		items, err := s.Find(ctx, p, conds)
		if err != nil {
			return err
		}

		lm.Results = items

		return nil
	})

	g.Go(func() error {
		count, err := s.Count(ctx, conds)
		if err != nil {
			return err
		}

		lm.Count = &count

		return nil
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	return &lm, nil
}

func (s *CollectionImpl) Find(ctx context.Context, p *paginate.Paginator, conds any) ([]*collection.ListItemModel, error) {
	items, err := s.collection.Find(ctx, p, conds)
	if err != nil {
		s.logger.Debugf("Error occurred while finding collections %v", err)

		return nil, err
	}

	return items, nil
}

func (s *CollectionImpl) Count(ctx context.Context, conds any) (uint, error) {
	count, err := s.collection.Count(ctx, conds)
	if err != nil {
		s.logger.Debugf("Error occurred while counting collections")

		return 0, err
	}

	return count, nil
}

func (s *CollectionImpl) Update(ctx context.Context, um *collection.UpdateModel) error {
	err := s.collection.Update(ctx, um)
	if err != nil {
		s.logger.Debugf("Error occurred while updating collection %v", err)

		return err
	}

	return nil
}

// Anime collection.

func (s *CollectionImpl) AddAnimes(ctx context.Context, aam *animecollection.AddAnimesModel) error {
	err := s.animeCollection.AddAnimes(ctx, aam)
	if err != nil {
		s.logger.Debugf("Error occurred while adding animes %v", err)

		return err
	}

	return nil
}

func (s *CollectionImpl) RemoveAnimes(ctx context.Context, ram *animecollection.RemoveAnimesModel) error {
	err := s.animeCollection.RemoveAnimes(ctx, ram)
	if err != nil {
		s.logger.Debugf("Error occurred while removing animes %v", err)

		return err
	}

	return nil
}
