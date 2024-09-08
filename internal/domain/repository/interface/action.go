package repositoryInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/pkg/postgres"
)

type Action interface {
	WithTx(tx postgres.RW) Action

	Create(ctx context.Context, dto *dto.ActionCreateRequestDTO) error
	//CreateCollectionCreate(ctx context.Context, *dto.ActionCreateCollectionRequestDTO) error
	//CreateUserRegister(ctx context.Context, *dto.ActionCreateUserRegisterRequestDTO) error
	//CreateUserUpdate(ctx context.Context, *dto.ActionCreateUserUpdateRequestDTO) error
	Find(ctx context.Context, dto *dto.ActionListRequestDTO, conds any) ([]*agg.ActionListItem, error)
	Count(ctx context.Context, conds any) (uint64, error)
}
