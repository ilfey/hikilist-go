package builderInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"net/http"
)

type User interface {
	BuildListRequestDTOFromRequest(r *http.Request) (*dto.UserListRequestDTO, error)
	BuildDetailRequestDTOFromRequest(r *http.Request) (*dto.UserDetailRequestDTO, error)
	BuildMeRequestDTOFromRequest(r *http.Request) (*dto.UserMeRequestDTO, error)
	BuildCollectionListRequestDTOFromRequest(r *http.Request) (*dto.UserCollectionListRequestDTO, error)
}
