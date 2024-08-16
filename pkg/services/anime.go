package services

import (
	"context"

	"github.com/ilfey/hikilist-go/internal/logger"
	"github.com/ilfey/hikilist-go/internal/paginate"
	"github.com/ilfey/hikilist-go/pkg/models/anime"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"golang.org/x/sync/errgroup"
)

type Anime interface {
	Create(ctx context.Context, cm *anime.CreateModel) error
	Get(ctx context.Context, conds any) (*anime.DetailModel, error)
	GetListModel(ctx context.Context, p *paginate.Paginator, conds any) (*anime.ListModel, error)
	GetFromCollectionListModel(ctx context.Context, p *paginate.Paginator, userId uint, collectionId uint) (*anime.ListModel, error)
	// FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*anime.ListItemModel, error)
	// FindFromCollectionWithPaginator(ctx context.Context, p *paginate.Paginator, userId, collectionId uint) ([]*anime.ListItemModel, error)
	// CountInCollection(ctx context.Context, userId, collectionId uint) (uint, error)
	// Count(ctx context.Context, conds any) (uint, error)
}

type AnimeImpl struct {
	anime repositories.Anime
}

func NewAnime(animeRepo repositories.Anime) Anime {
	return &AnimeImpl{
		anime: animeRepo,
	}
}

func (s *AnimeImpl) Create(ctx context.Context, cm *anime.CreateModel) error {
	err := s.anime.Create(ctx, cm)
	if err != nil {
		logger.Debugf("Error occurred while creating anime %v", err)

		return err
	}

	return nil
}

func (s *AnimeImpl) Get(ctx context.Context, conds any) (*anime.DetailModel, error) {
	dm, err := s.anime.Get(ctx, conds)
	if err != nil {
		logger.Debugf("Error occurred while getting anime %v", err)

		return nil, err
	}

	return dm, nil
}

func (s *AnimeImpl) GetListModel(ctx context.Context, p *paginate.Paginator, conds any) (*anime.ListModel, error) {
	var lm anime.ListModel

	g := errgroup.Group{}

	g.Go(func() error {
		items, err := s.FindWithPaginator(ctx, p, conds)
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

func (s *AnimeImpl) GetFromCollectionListModel(
	ctx context.Context,
	p *paginate.Paginator,
	userId uint,
	collectionId uint,
) (*anime.ListModel, error) {
	var lm anime.ListModel

	g := errgroup.Group{}

	g.Go(func() error {
		items, err := s.FindFromCollectionWithPaginator(ctx, p, userId, collectionId)
		if err != nil {
			return err
		}

		lm.Results = items

		return nil
	})

	g.Go(func() error {
		count, err := s.CountInCollection(ctx, userId, collectionId)
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

func (s *AnimeImpl) FindWithPaginator(ctx context.Context, p *paginate.Paginator, conds any) ([]*anime.ListItemModel, error) {
	items, err := s.anime.FindWithPaginator(ctx, p, conds)
	if err != nil {
		logger.Debugf("Error occurred while finding animes %v", err)

		return nil, err
	}

	return items, nil
}

func (s *AnimeImpl) FindFromCollectionWithPaginator(
	ctx context.Context,
	p *paginate.Paginator,
	userId uint,
	collectionId uint,
) ([]*anime.ListItemModel, error) {
	items, err := s.anime.FindFromCollectionWithPaginator(ctx, p, userId, collectionId)
	if err != nil {
		logger.Debugf("Error occurred while finding animes %v", err)

		return nil, err
	}

	return items, nil
}

func (s *AnimeImpl) CountInCollection(ctx context.Context, userId uint, collectionId uint) (uint, error) {
	count, err := s.anime.CountInCollection(ctx, userId, collectionId)
	if err != nil {
		logger.Debugf("Error occurred while counting animes %v", err)

		return 0, err
	}

	return count, nil
}

func (s *AnimeImpl) Count(ctx context.Context, conds any) (uint, error) {
	count, err := s.anime.Count(ctx, conds)
	if err != nil {
		logger.Debugf("Error occurred while counting animes %v", err)

		return 0, err
	}

	return count, nil
}
