package anime

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/anime/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"golang.org/x/sync/errgroup"
)

type Anime struct {
	log loggerInterface.Logger

	anime repositoryInterface.Anime
}

func NewAnime(log loggerInterface.Logger, anime repositoryInterface.Anime) animeInterface.Anime {
	return &Anime{
		log: log,

		anime: anime,
	}
}

func (s *Anime) Create(ctx context.Context, cm *dto.AnimeCreateRequestDTO) error {
	err := s.anime.Create(ctx, cm)
	if err != nil {
		return s.log.LogPropagate(err)
	}

	return nil
}

func (s *Anime) Get(ctx context.Context, conds any) (*agg.AnimeDetail, error) {
	dm, err := s.anime.Get(ctx, conds)
	if err != nil {
		return nil, s.log.LogPropagate(err)
	}

	return dm, nil
}

func (s *Anime) GetListModel(ctx context.Context, dto *dto.AnimeListRequestDTO, conds any) (*agg.AnimeList, error) {
	var lm agg.AnimeList

	g := errgroup.Group{}

	g.Go(func() error {
		items, err := s.FindWithPaginator(ctx, dto, conds)
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
		return nil, s.log.LogPropagate(err)
	}

	return &lm, nil
}

func (s *Anime) GetFromCollectionListDTO(ctx context.Context, dto *dto.AnimeListFromCollectionRequestDTO) (*agg.AnimeList, error) {
	var lm agg.AnimeList

	g := errgroup.Group{}

	g.Go(func() error {
		items, err := s.FindFromCollectionWithPaginator(ctx, dto)
		if err != nil {
			return err
		}

		lm.Results = items

		return nil
	})

	g.Go(func() error {
		count, err := s.CountInCollection(ctx, dto)
		if err != nil {
			return err
		}

		lm.Count = &count

		return nil
	})

	err := g.Wait()
	if err != nil {
		return nil, s.log.LogPropagate(err)
	}

	return &lm, nil
}

func (s *Anime) FindWithPaginator(ctx context.Context, dto *dto.AnimeListRequestDTO, conds any) ([]*agg.AnimeListItem, error) {
	items, err := s.anime.FindWithPaginator(ctx, dto, conds)
	if err != nil {
		return nil, s.log.LogPropagate(err)
	}

	return items, nil
}

func (s *Anime) FindFromCollectionWithPaginator(
	ctx context.Context,
	dto *dto.AnimeListFromCollectionRequestDTO,
) ([]*agg.AnimeListItem, error) {
	items, err := s.anime.FindFromCollectionWithPaginator(ctx, dto)
	if err != nil {
		return nil, s.log.LogPropagate(err)
	}

	return items, nil
}

func (s *Anime) CountInCollection(ctx context.Context, dto *dto.AnimeListFromCollectionRequestDTO) (uint64, error) {
	count, err := s.anime.CountInCollection(ctx, dto)
	if err != nil {
		return 0, s.log.LogPropagate(err)
	}

	return count, nil
}

func (s *Anime) Count(ctx context.Context, conds any) (uint64, error) {
	count, err := s.anime.Count(ctx, conds)
	if err != nil {
		return 0, s.log.LogPropagate(err)
	}

	return count, nil
}
