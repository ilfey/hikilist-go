package builderInterface

import "github.com/ilfey/hikilist-go/internal/domain/dto"

type Pagination interface {
	BuilderPaginationRequestFromPageAndLimit(page, limit *uint64) (*dto.PaginationRequestDTO, error)
}
