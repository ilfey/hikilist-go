package userInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type CRUD interface {
	Create(ctx context.Context, dto *dto.UserCreateRequestDTO) error
	Detail(ctx context.Context, dto *dto.UserDetailRequestDTO) (*agg.UserDetail, error)
	List(ctx context.Context, dto *dto.UserListRequestDTO) (*agg.UserList, error)

	// TODO: Remove this methods
	//ChangeUsername(userId uint64, oldUsername, newUsername string) error
	//UpdatePassword(userId uint64, hash string) error

	UpdateLastOnline(ctx context.Context, userId uint64) error
	Delete(ctx context.Context, conds any) error
}
