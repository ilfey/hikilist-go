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
	GetUserPublicCollectionListDTO(ctx context.Context, dto *dto.UserCollectionListRequestDTO) (*agg.CollectionList, error)
	GetUserCollectionListDTO(ctx context.Context, dto *dto.UserCollectionListRequestDTO) (*agg.CollectionList, error)
	Update(ctx context.Context, um *agg.CollectionDetail) error

	// TODO: create Remove(ctx context.Context, dto *dto.CollectionRemoveRequestDTO) error
	//Remove(ctx context.Context, dto *dto.CollectionRemoveRequestDTO) error

	AddAnimes(ctx context.Context, aam *dto.CollectionAddAnimeRequestDTO) error
	RemoveAnimes(ctx context.Context, ram *dto.CollectionRemoveAnimeRequestDTO) error
}
