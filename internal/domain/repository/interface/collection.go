package repositoryInterface

import (
	"context"
	"github.com/ilfey/hikilist-database/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/pkg/postgres"
)

type Collection interface {
	WithTx(tx postgres.RW) Collection

	// Create creates new collection and action.
	Create(ctx context.Context, req *dto.CollectionCreateRequestDTO) error

	// Get returns collection.
	Get(ctx context.Context, conds any) (*agg.CollectionDetail, error)

	// Find returns collection list.
	Find(ctx context.Context, req *dto.CollectionListRequestDTO, conds any) ([]*agg.CollectionListItem, error)

	// FindUserPublicCollectionList returns user public collection list.
	FindUserPublicCollectionList(ctx context.Context, req *dto.UserCollectionListRequestDTO) ([]*agg.CollectionListItem, error)

	// FindUserCollectionList returns user collection list with private and public collections.
	FindUserCollectionList(ctx context.Context, req *dto.UserCollectionListRequestDTO) ([]*agg.CollectionListItem, error)

	// Count returns count of collections.
	Count(ctx context.Context, conds any) (uint64, error)

	// CountUserPublicCollection returns count of user public collections.
	CountUserPublicCollection(ctx context.Context, req *dto.UserCollectionListRequestDTO) (uint64, error)

	// CountUserCollection returns count of user collections with private and public collections.
	CountUserCollection(ctx context.Context, req *dto.UserCollectionListRequestDTO) (uint64, error)

	// Update updates collection.
	Update(ctx context.Context, req *agg.CollectionDetail) error

	// Delete deletes collection.
	Delete(ctx context.Context, req *dto.CollectionDeleteRequestDTO) error
}
