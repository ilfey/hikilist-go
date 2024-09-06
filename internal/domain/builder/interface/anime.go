package builderInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"net/http"
)

type Anime interface {
	BuildDetailRequestDTOFromRequest(r *http.Request) (*dto.AnimeDetailRequestDTO, error)
	BuildListRequestDTOFromRequest(r *http.Request) (*dto.AnimeListRequestDTO, error)
	BuildCreateRequestDTOFromRequest(r *http.Request) (*dto.AnimeCreateRequestDTO, error)
}
