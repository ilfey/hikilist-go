package collectionInterface

import (
	"context"
	"github.com/ilfey/hikilist-database/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type Collection interface {
	// Create creates new collection and action.
	Create(ctx context.Context, req *dto.CollectionCreateRequestDTO) error

	// Get returns collection.
	Get(ctx context.Context, conds any) (*agg.CollectionDetail, error)

	// GetList returns list of collections.
	GetList(ctx context.Context, req *dto.CollectionListRequestDTO, conds any) (*agg.CollectionList, error)

	// GetUserPublicCollectionList returns list of public collections.
	GetUserPublicCollectionList(ctx context.Context, req *dto.UserCollectionListRequestDTO) (*agg.CollectionList, error)

	// GetUserCollectionList returns list of user collections with private collections.
	GetUserCollectionList(ctx context.Context, req *dto.UserCollectionListRequestDTO) (*agg.CollectionList, error)

	// Update updates collection.
	Update(ctx context.Context, req *agg.CollectionDetail) error

	// Delete deletes collection.
	Delete(ctx context.Context, req *dto.CollectionDeleteRequestDTO) error

	// AddAnimes adds animes to collection.
	AddAnimes(ctx context.Context, req *dto.CollectionAddAnimeRequestDTO) error

	// RemoveAnimes removes animes from collection.
	RemoveAnimes(ctx context.Context, req *dto.CollectionRemoveAnimeRequestDTO) error
}
