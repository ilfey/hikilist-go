package builderInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"net/http"
)

type Collection interface {
	BuildAnimeListFromCollectionRequestDTOFromRequest(r *http.Request) (*dto.AnimeListFromCollectionRequestDTO, error)
	BuildUpdateRequestDTOFromRequest(r *http.Request) (*dto.CollectionUpdateRequestDTO, error)
	BuildRemoveAnimeRequestDTOFromRequest(r *http.Request) (*dto.CollectionRemoveAnimeRequestDTO, error)
	BuildCreateRequestDTOFromRequest(r *http.Request) (*dto.CollectionCreateRequestDTO, error)
	BuildAddAnimeRequestDTOFromRequest(r *http.Request) (*dto.CollectionAddAnimeRequestDTO, error)
	BuildListRequestDTOFromRequest(r *http.Request) (*dto.CollectionListRequestDTO, error)
	BuildDetailRequestDTOFromRequest(r *http.Request) (*dto.CollectionDetailRequestDTO, error)
}
