package builder

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
)

type Pagination struct {
	log loggerInterface.Logger
}

func NewPagination(container diInterface.ServiceContainer) (*Pagination, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	return &Pagination{
		log: log,
	}, nil
}

func (p Pagination) BuilderPaginationRequestFromPageAndLimit(page, limit *uint64) (*dto.PaginationRequestDTO, error) {
	var (
		pg uint64 = 1
		lt uint64 = 10
	)

	if page != nil {
		pg = *page
	}

	if limit != nil {
		lt = *limit
	}

	return &dto.PaginationRequestDTO{
		Page:  pg,
		Limit: lt,
	}, nil
}
