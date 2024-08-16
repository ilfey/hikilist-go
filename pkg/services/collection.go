package services

import (
	"context"

	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/pkg/models/collection"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"golang.org/x/sync/errgroup"
)

type Collection interface {
	Create(ctx context.Context, cm *collection.CreateModel) error
	Get(ctx context.Context, conds any) (*collection.DetailModel, error)
	GetListModel(ctx context.Context, p *paginate.Paginator, conds any) (*collection.ListModel, error)
	// Find(ctx context.Context, p *paginate.Paginator, conds any) ([]*collection.ListItemModel, error)
	// Count(ctx context.Context, conds any) (uint, error)
	Update(ctx context.Context, um *collection.UpdateModel) error
}

type CollectionImpl struct {
	collection repositories.Collection
}

func NewCollection(collectionRepo repositories.Collection) Collection {
	return &CollectionImpl{
		collection: collectionRepo,
	}
}

func (s *CollectionImpl) Create(ctx context.Context, cm *collection.CreateModel) error {
	if cm.IsPublic == nil {
		t := true

		cm.IsPublic = &t
	}

	err := s.collection.Create(ctx, cm)
	if err != nil {
		logger.Debugf("Error occurred while creating collection %v", err)

		return err
	}

	return nil
}

func (s *CollectionImpl) Get(ctx context.Context, conds any) (*collection.DetailModel, error) {
	dm, err := s.collection.Get(ctx, conds)
	if err != nil {
		logger.Debugf("Error occurred while getting collection %v", err)

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
		logger.Debugf("Error occurred while finding collections %v", err)

		return nil, err
	}

	return items, nil
}

func (s *CollectionImpl) Count(ctx context.Context, conds any) (uint, error) {
	count, err := s.collection.Count(ctx, conds)
	if err != nil {
		logger.Debugf("Error occurred while counting collections")

		return 0, err
	}

	return count, nil
}

func (s *CollectionImpl) Update(ctx context.Context, um *collection.UpdateModel) error {
	err := s.collection.Update(ctx, um)
	if err != nil {
		logger.Debugf("Error occurred while updating collection %v", err)

		return err
	}

	return nil
}
