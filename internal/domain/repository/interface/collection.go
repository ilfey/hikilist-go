package repositoryInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/pkg/postgres"
)

type Collection interface {
	WithTx(tx postgres.RW) Collection

	Create(ctx context.Context, cm *dto.CollectionCreateRequestDTO) error
	Get(ctx context.Context, conds any) (*agg.CollectionDetail, error)
	Find(ctx context.Context, dto *dto.CollectionListRequestDTO, conds any) ([]*agg.CollectionListItem, error)
	Count(ctx context.Context, conds any) (uint64, error)
	Update(ctx context.Context, um *dto.CollectionUpdateRequestDTO) error
}
