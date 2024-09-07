package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

type Collection struct {
	log loggerInterface.Logger
}

func NewCollection(log loggerInterface.Logger) *Collection {
	return &Collection{
		log: log,
	}
}

func (v *Collection) ValidateAddAnimeRequestDTO(dto *dto.CollectionAddAnimeRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Animes": {
				options.Required(),
				options.LenGreaterThan(0),
			},
		},
	)
	if !ok {
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil
}

func (v *Collection) ValidateRemoveAnimeRequestDTO(dto *dto.CollectionRemoveAnimeRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Animes": {
				options.Required(),
				options.LenGreaterThan(0),
			},
		},
	)
	if !ok {
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil
}

func (v *Collection) ValidateUpdateRequestDTO(dto *dto.CollectionUpdateRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Title": {
				options.IfNotNil(
					options.LenGreaterThan(3),
					options.LenLessThan(256),
				),
			},
			"Description": {
				options.IfNotNil(
					options.LenGreaterThan(64),
					options.LenLessThan(4096),
				),
			},
		},
	)
	if !ok {
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil
}

func (v *Collection) ValidateRemoveRequestDTO(dto *dto.CollectionRemoveAnimeRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Animes": {
				options.Required(),
				options.LenGreaterThan(0),
			},
		},
	)
	if !ok {
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil
}

func (v *Collection) ValidateCreateRequestDTO(dto *dto.CollectionCreateRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Title": {
				options.LenGreaterThan(3),
				options.LenLessThan(256),
			},
			"Description": {
				options.IfNotNil(
					options.LenGreaterThan(64),
					options.LenLessThan(4096),
				),
			},
		},
	)
	if !ok {
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil
}

func (v *Collection) ValidateListRequestDTO(dto *dto.CollectionListRequestDTO) error {
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

func (v *Collection) ValidateAnimeListFromCollectionRequestDTO(dto *dto.AnimeListFromCollectionRequestDTO) error {
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

func (v *Collection) ValidateDetailRequestDTO(dto *dto.CollectionDetailRequestDTO) error {
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
