package anime

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	"github.com/ilfey/hikilist-go/internal/domain/service/anime/interface"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	validatorInterface "github.com/ilfey/hikilist-go/internal/domain/validator/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"golang.org/x/sync/errgroup"
)

type Anime struct {
	log loggerInterface.Logger

	validator           validatorInterface.Anime
	paginationValidator validatorInterface.Pagination

	anime repositoryInterface.Anime
}

func NewAnime(container diInterface.ServiceContainer) (animeInterface.Anime, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	validator, err := container.GetAnimeValidator()
	if err != nil {
		return nil, log.Propagate(err)
	}

	paginationValidator, err := container.GetPaginationValidator()
	if err != nil {
		return nil, log.Propagate(err)
	}

	repo, err := container.GetAnimeRepository()
	if err != nil {
		return nil, log.Propagate(err)
	}

	return &Anime{
		log:                 log,
		validator:           validator,
		paginationValidator: paginationValidator,
		anime:               repo,
	}, nil
}

func (s *Anime) Create(ctx context.Context, req *dto.AnimeCreateRequestDTO) error {
	// Validate.
	if err := s.validator.ValidateCreateRequestDTO(req); err != nil {
		return s.log.Propagate(err)
	}

	// Create.
	err := s.anime.Create(ctx, req)
	if err != nil {
		return s.log.Propagate(err)
	}

	return nil
}

func (s *Anime) Get(ctx context.Context, conds any) (*agg.Anime, error) {
	dm, err := s.anime.Get(ctx, conds)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	return dm, nil
}

func (s *Anime) GetByID(ctx context.Context, id uint64) (*agg.Anime, error) {
	dm, err := s.anime.GetByID(ctx, id)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	return dm, nil
}

func (s *Anime) Find(ctx context.Context, req *dto.PaginationRequestDTO) ([]agg.Anime, error) {
	// Validate.
	if err := s.paginationValidator.Validate(req); err != nil {
		return nil, s.log.Propagate(err)
	}

	// Find.
	list, err := s.anime.Find(ctx, req)
	if err != nil {
		return nil, s.log.Propagate(err)
	}

	return list, nil
}

func (s *Anime) GetListModel(ctx context.Context, dto *dto.AnimeListRequestDTO, conds any) (*agg.AnimeList, error) {
	// Validate.
	if err := s.validator.ValidateListRequestDTO(dto); err != nil {
		return nil, s.log.Propagate(err)
	}

	var lm agg.AnimeList

	g := errgroup.Group{}

	// Find.
	g.Go(func() error {
		items, err := s.anime.FindWithPaginator(ctx, dto, conds)
		if err != nil {
			return s.log.Propagate(err)
		}

		lm.Results = items

		return nil
	})

	// Count.
	g.Go(func() error {
		count, err := s.anime.Count(ctx, conds)
		if err != nil {
			return s.log.Propagate(err)
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

func (s *Anime) GetFromCollectionListDTO(ctx context.Context, dto *dto.AnimeListFromCollectionRequestDTO) (*agg.AnimeList, error) {
	// TODO: add validation.

	var lm agg.AnimeList

	g := errgroup.Group{}

	// Find.
	g.Go(func() error {
		items, err := s.anime.FindFromCollectionWithPaginator(ctx, dto)
		if err != nil {
			return s.log.Propagate(err)
		}

		lm.Results = items

		return nil
	})

	// Count.
	g.Go(func() error {
		count, err := s.anime.CountInCollection(ctx, dto)
		if err != nil {
			return s.log.Propagate(err)
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
