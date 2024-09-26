package actionInterface

import (
	"context"
	"github.com/ilfey/hikilist-database/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type Action interface {
	GetListDTO(ctx context.Context, dto *dto.ActionListRequestDTO, conds any) (*agg.ActionList, error)
}
