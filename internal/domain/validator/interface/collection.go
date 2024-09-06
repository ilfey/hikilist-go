package validatorInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type Collection interface {
	ValidateAddAnimeRequestDTO(req *dto.CollectionAddAnimeRequestDTO) error
	ValidateRemoveAnimeRequestDTO(req *dto.CollectionRemoveAnimeRequestDTO) error
	ValidateUpdateRequestDTO(req *dto.CollectionUpdateRequestDTO) error

	// TODO: Add ValidateDeleteRequestDTO(req *dto.CollectionRemoveRequestDTO) error

	ValidateCreateRequestDTO(req *dto.CollectionCreateRequestDTO) error
	ValidateListRequestDTO(req *dto.CollectionListRequestDTO) error
	ValidateAnimeListFromCollectionRequestDTO(req *dto.AnimeListFromCollectionRequestDTO) error
	ValidateDetailRequestDTO(req *dto.CollectionDetailRequestDTO) error
}
