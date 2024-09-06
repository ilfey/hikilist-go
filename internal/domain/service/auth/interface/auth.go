package authInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"net/http"
)

type Auth interface {
	IsAuthed(request *http.Request) (uint64, error)
	ChangePassword(ctx context.Context, userId uint64, password string) error
	DeleteUser(ctx context.Context, deleteDTO *dto.UserDeleteRequestDTO) error
	Refresh(ctx context.Context, refreshDTO *dto.AuthRefreshRequestDTO) (*agg.TokenPair, error)
	Logout(ctx context.Context, lofoutDTO *dto.AuthLogoutRequestDTO) error
	Login(ctx context.Context, loginDTO *dto.AuthLoginRequestDTO) (*agg.TokenPair, error)
	Register(ctx context.Context, registerDTO *dto.AuthRegisterRequestDTO) (*dto.UserCreateRequestDTO, error)
}
