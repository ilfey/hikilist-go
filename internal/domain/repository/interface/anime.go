package repositoryInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/pkg/postgres"
)

type Anime interface {
	WithTx(tx postgres.RW) Anime

	Create(ctx context.Context, cm *dto.AnimeCreateRequestDTO) error
	Get(ctx context.Context, conds any) (*agg.Anime, error)
	GetByID(ctx context.Context, id uint64) (*agg.Anime, error)
	Find(ctx context.Context, req *dto.PaginationRequestDTO) ([]agg.Anime, error)
	FindWithPaginator(ctx context.Context, dto *dto.AnimeListRequestDTO, conds any) ([]*agg.AnimeListItem, error)
	FindFromCollectionWithPaginator(ctx context.Context, dto *dto.AnimeListFromCollectionRequestDTO) ([]*agg.AnimeListItem, error)
	CountInCollection(ctx context.Context, dto *dto.AnimeListFromCollectionRequestDTO) (uint64, error)
	Count(ctx context.Context, conds any) (uint64, error)
}
