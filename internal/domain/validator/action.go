package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

type Action struct {
	log loggerInterface.Logger
}

func NewAction(log loggerInterface.Logger) *Action {
	return &Action{
		log: log,
	}
}

func (b *Action) ValidateListRequestDTO(dto *dto.ActionListRequestDTO) error {
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
		return errtype.NewValidatorError("GetActionList", expectations)
	}

	return nil
}
