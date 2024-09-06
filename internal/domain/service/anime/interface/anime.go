package animeInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type Anime interface {
	Create(ctx context.Context, cm *dto.AnimeCreateRequestDTO) error
	Get(ctx context.Context, conds any) (*agg.AnimeDetail, error)
	GetListModel(ctx context.Context, dto *dto.AnimeListRequestDTO, conds any) (*agg.AnimeList, error)
	GetFromCollectionListDTO(ctx context.Context, dto *dto.AnimeListFromCollectionRequestDTO) (*agg.AnimeList, error)
}
