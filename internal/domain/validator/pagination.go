package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
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

func (v *Pagination) Validate(req *dto.PaginationRequestDTO) error {
	expectations, ok := validator.Validate(
		req,
		map[string][]options.Option{
			"Page": {
				options.GreaterThan[uint64](0),
			},
			"Limit": {
				options.GreaterThan[uint64](0),
				options.LessThan[uint64](101),
			},
		},
	)
	if !ok {
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil
}
