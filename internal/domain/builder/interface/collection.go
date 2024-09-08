package builderInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"net/http"
)

type Collection interface {
	BuildAnimeListFromCollectionRequestDTOFromRequest(req *http.Request) (*dto.AnimeListFromCollectionRequestDTO, error)
	BuildUpdateRequestDTOFromRequest(req *http.Request) (*dto.CollectionUpdateRequestDTO, error)
	BuildRemoveAnimeRequestDTOFromRequest(req *http.Request) (*dto.CollectionRemoveAnimeRequestDTO, error)
	BuildCreateRequestDTOFromRequest(req *http.Request) (*dto.CollectionCreateRequestDTO, error)
	BuildAddAnimeRequestDTOFromRequest(req *http.Request) (*dto.CollectionAddAnimeRequestDTO, error)
	BuildListRequestDTOFromRequest(req *http.Request) (*dto.CollectionListRequestDTO, error)
	BuildDetailRequestDTOFromRequest(req *http.Request) (*dto.CollectionDetailRequestDTO, error)

	BuildAggFromUpdateRequestDTO(ctx context.Context, req *dto.CollectionUpdateRequestDTO) (*agg.CollectionDetail, error)
}
