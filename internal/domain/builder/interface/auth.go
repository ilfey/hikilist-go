package builderInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"net/http"
)

type Auth interface {
	BuildChangePasswordRequestDTOFromRequest(r *http.Request) (*dto.AuthChangePasswordRequestDTO, error)
	BuildChangeUsernameRequestDTOFromRequest(r *http.Request) (*dto.AuthChangeUsernameRequestDTO, error)
	BuildDeleteRequestDTOFromRequest(r *http.Request) (*dto.UserDeleteRequestDTO, error)
	BuildLoginRequestDTOFromRequest(r *http.Request) (*dto.AuthLoginRequestDTO, error)
	BuildLogoutRequestDTOFromRequest(r *http.Request) (*dto.AuthLogoutRequestDTO, error)
	BuildRefreshRequestDTOFromRequest(r *http.Request) (*dto.AuthRefreshRequestDTO, error)
	BuildRegisterRequestDTOFromRequest(r *http.Request) (*dto.AuthRegisterRequestDTO, error)
}
