package validatorInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type Collection interface {
	// ValidateAddAnimeRequestDTO validates create collection request DTO.
	ValidateAddAnimeRequestDTO(req *dto.CollectionAddAnimeRequestDTO) error

	// ValidateRemoveAnimeRequestDTO validates remove anime request DTO.
	ValidateRemoveAnimeRequestDTO(req *dto.CollectionRemoveAnimeRequestDTO) error

	// ValidateUpdateRequestDTO validates update collection request DTO.
	ValidateUpdateRequestDTO(req *dto.CollectionUpdateRequestDTO) error

	// ValidateDeleteRequestDTO validates delete collection request DTO.
	ValidateDeleteRequestDTO(req *dto.CollectionDeleteRequestDTO) error

	// ValidateCreateRequestDTO validates create collection request DTO.
	ValidateCreateRequestDTO(req *dto.CollectionCreateRequestDTO) error

	// ValidateListRequestDTO validates list collection request DTO.
	ValidateListRequestDTO(req *dto.CollectionListRequestDTO) error

	// ValidateAnimeListFromCollectionRequestDTO validates anime list from collection request DTO.
	ValidateAnimeListFromCollectionRequestDTO(req *dto.AnimeListFromCollectionRequestDTO) error

	// ValidateDetailRequestDTO validates detail collection request DTO.
	ValidateDetailRequestDTO(req *dto.CollectionDetailRequestDTO) error
}
