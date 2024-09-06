package validatorInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type Action interface {
	ValidateListRequestDTO(req *dto.ActionListRequestDTO) error
}
