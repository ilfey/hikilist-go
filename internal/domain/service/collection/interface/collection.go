package collectionInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type Collection interface {
	Create(ctx context.Context, cm *dto.CollectionCreateRequestDTO) error
	Get(ctx context.Context, conds any) (*agg.CollectionDetail, error)
	GetListDTO(ctx context.Context, dto *dto.CollectionListRequestDTO, conds any) (*agg.CollectionList, error)
	// GetListDTO(ctx context.Context, p *paginate.Paginator, conds any) ([]*collection.CollectionListItem, error)
	// Count(ctx context.Context, conds any) (uint64, error)
	Update(ctx context.Context, um *dto.CollectionUpdateRequestDTO) error

	// TODO: create Remove(ctx context.Context, dto *dto.CollectionRemoveRequestDTO) error
	//Remove(ctx context.Context, dto *dto.CollectionRemoveRequestDTO) error

	AddAnimes(ctx context.Context, aam *dto.CollectionAddAnimeRequestDTO) error
	RemoveAnimes(ctx context.Context, ram *dto.CollectionRemoveAnimeRequestDTO) error
}
