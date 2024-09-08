package validatorInterface

import "github.com/ilfey/hikilist-go/internal/domain/dto"

type Pagination interface {
	Validate(req *dto.PaginationRequestDTO) error
}
