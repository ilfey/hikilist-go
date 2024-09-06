package actionInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type Action interface {
	// CreateRegisterAction(ctx context.Context, userId uint64) error
	// CreateCollectionCreateAction(ctx context.Context, userId uint64, collectionTitle string) error
	// CreateUpdateUsernameAction(ctx context.Context, userId uint64, oldUsername, newUsername string) error

	// Create(ctx context.Context, cm *action.ActionCreateRequestDTO) error
	GetListDTO(ctx context.Context, dto *dto.ActionListRequestDTO, conds any) (*agg.ActionList, error)
	// ActionListRequestDTO(ctx context.Context, p *paginate.Paginator, conds any) ([]*action.ActionListItem, error)
	// Count(ctx context.Context, conds any) (uint64, error)
}
