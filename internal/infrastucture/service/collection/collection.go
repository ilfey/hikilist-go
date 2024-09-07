package collection

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/collection/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"golang.org/x/sync/errgroup"
)

type Collection struct {
	log loggerInterface.Logger

	animeCollection repositoryInterface.AnimeCollection
	collection      repositoryInterface.Collection
}

func NewCollection(
	log loggerInterface.Logger,

	animeCollection repositoryInterface.AnimeCollection,
	collection repositoryInterface.Collection,
) collectionInterface.Collection {
	return &Collection{
		log: log,

		animeCollection: animeCollection,
		collection:      collection,
	}
}

func (s *Collection) Create(ctx context.Context, cm *dto.CollectionCreateRequestDTO) error {
	if cm.IsPublic == nil {
		t := true

		cm.IsPublic = &t
	}

	err := s.collection.Create(ctx, cm)
	if err != nil {
		return s.log.Propagate(err)
	}

	return nil
}

func (s *Collection) Get(ctx context.Context, conds any) (*agg.CollectionDetail, error) {
	dm, err := s.collection.Get(ctx, conds)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	return dm, nil
}

func (s *Collection) GetListDTO(ctx context.Context, p *dto.CollectionListRequestDTO, conds any) (*agg.CollectionList, error) {
	var lm agg.CollectionList

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
		return nil, s.log.Propagate(err)
	}

	return &lm, nil
}

func (s *Collection) Find(ctx context.Context, dto *dto.CollectionListRequestDTO, conds any) ([]*agg.CollectionListItem, error) {
	items, err := s.collection.Find(ctx, dto, conds)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	return items, nil
}

func (s *Collection) Count(ctx context.Context, conds any) (uint64, error) {
	count, err := s.collection.Count(ctx, conds)
	if err != nil {
		return 0, s.log.Propagate(err)
	}

	return count, nil
}

func (s *Collection) Update(ctx context.Context, um *dto.CollectionUpdateRequestDTO) error {
	err := s.collection.Update(ctx, um)
	if err != nil {
		return s.log.Propagate(err)
	}

	return nil
}

/* ===== AnimeCollection Implementation ===== */

func (s *Collection) AddAnimes(ctx context.Context, aam *dto.CollectionAddAnimeRequestDTO) error {
	err := s.animeCollection.AddAnimes(ctx, aam)
	if err != nil {
		return s.log.Propagate(err)
	}

	return nil
}

func (s *Collection) RemoveAnimes(ctx context.Context, ram *dto.CollectionRemoveAnimeRequestDTO) error {
	err := s.animeCollection.RemoveAnimes(ctx, ram)
	if err != nil {
		return s.log.Propagate(err)
	}

	return nil
}
