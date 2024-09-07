package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

type Action struct {
	log loggerInterface.Logger
}

func NewAction(container diInterface.ServiceContainer) (*Action, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	return &Action{
		log: log,
	}, nil
}

func (v *Action) ValidateListRequestDTO(dto *dto.ActionListRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Page": {
				options.GreaterThan[uint64](0),
			},
			"Limit": {
				options.GreaterThan[uint64](0),
				options.LessThan[uint64](101),
			},
			"Order": {
				options.InList([]string{
					"id",
					"-id",
				}),
			},
		},
	)
	if !ok {
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil
}
