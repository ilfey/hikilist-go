package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

type Action struct {
	logger loggerInterface.Logger
}

func NewAction(logger loggerInterface.Logger) *Action {
	return &Action{
		logger: logger,
	}
}

func (b *Action) ValidateListRequestDTO(dto *dto.ActionListRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Limit": {
				options.LessThan[uint64](101),
			},
			"Order": {
				options.InList([]string{
					"",
					"id",
					"-id",
				}),
			},
		},
	)
	if !ok {
		return errtype.NewValidatorError("GetActionList", expectations)
	}

	return nil
}
