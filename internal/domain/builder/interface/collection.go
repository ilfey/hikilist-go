package builderInterface

import (
	"context"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"net/http"
)

type Collection interface {
	// BuildAnimeListFromCollectionRequestDTOFromRequest builds anime list from collection request DTO from request.
	BuildAnimeListFromCollectionRequestDTOFromRequest(req *http.Request) (*dto.AnimeListFromCollectionRequestDTO, error)

	// BuildUpdateRequestDTOFromRequest builds update collection request DTO from request.
	BuildUpdateRequestDTOFromRequest(req *http.Request) (*dto.CollectionUpdateRequestDTO, error)

	// BuildRemoveAnimeRequestDTOFromRequest builds remove anime request DTO from request.
	BuildRemoveAnimeRequestDTOFromRequest(req *http.Request) (*dto.CollectionRemoveAnimeRequestDTO, error)

	// BuildCreateRequestDTOFromRequest builds create collection request DTO from request.
	BuildCreateRequestDTOFromRequest(req *http.Request) (*dto.CollectionCreateRequestDTO, error)

	// BuildAddAnimeRequestDTOFromRequest builds add anime request DTO from request.
	BuildAddAnimeRequestDTOFromRequest(req *http.Request) (*dto.CollectionAddAnimeRequestDTO, error)

	// BuildListRequestDTOFromRequest builds list collection request DTO from request.
	BuildListRequestDTOFromRequest(req *http.Request) (*dto.CollectionListRequestDTO, error)

	// BuildDetailRequestDTOFromRequest builds detail collection request DTO from request.
	BuildDetailRequestDTOFromRequest(req *http.Request) (*dto.CollectionDetailRequestDTO, error)

	// BuildDeleteRequestDTOFromRequest builds delete collection request DTO from request.
	BuildDeleteRequestDTOFromRequest(req *http.Request) (*dto.CollectionDeleteRequestDTO, error)

	// BuildAggFromUpdateRequestDTO builds agg from update collection request DTO.
	BuildAggFromUpdateRequestDTO(ctx context.Context, req *dto.CollectionUpdateRequestDTO) (*agg.CollectionDetail, error)
}
