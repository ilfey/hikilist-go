package validatorInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type Anime interface {
	ValidateCreateRequestDTO(req *dto.AnimeCreateRequestDTO) error
	ValidateDetailRequestDTO(req *dto.AnimeDetailRequestDTO) error
	ValidateListRequestDTO(req *dto.AnimeListRequestDTO) error
}
