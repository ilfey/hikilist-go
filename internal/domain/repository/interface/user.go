package repositoryInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/pkg/postgres"
)

type User interface {
	WithTx(tx postgres.RW) User

	Create(ctx context.Context, cm *dto.UserCreateRequestDTO) error
	Get(ctx context.Context, conds any) (*agg.UserDetail, error)
	Find(ctx context.Context, dto *dto.UserListRequestDTO, conds any) ([]*agg.UserListItem, error)
	Count(ctx context.Context, conds any) (uint64, error)
	UpdateUsername(ctx context.Context, userId uint64, oldUsername, newUsername string) error
	UpdateLastOnline(ctx context.Context, userId uint64) error
	UpdatePassword(ctx context.Context, userId uint64, hash string) error
	Delete(ctx context.Context, conds any) error
}
