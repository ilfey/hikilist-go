package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

type Collection struct {
	logger loggerInterface.Logger
}

func NewCollection(logger loggerInterface.Logger) *Collection {
	return &Collection{
		logger: logger,
	}
}

func (b *Collection) ValidateAddAnimeRequestDTO(dto *dto.CollectionAddAnimeRequestDTO) error {
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
		return b.logger.LogPropagate(errtype.NewValidatorError("AddAnimeToCollection", expectations))
	}

	return nil
}

func (b *Collection) ValidateRemoveAnimeRequestDTO(dto *dto.CollectionRemoveAnimeRequestDTO) error {
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
		return b.logger.LogPropagate(errtype.NewValidatorError("AddAnimeToCollection", expectations))
	}

	return nil
}

func (b *Collection) ValidateUpdateRequestDTO(dto *dto.CollectionUpdateRequestDTO) error {
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
		return b.logger.LogPropagate(errtype.NewValidatorError("UpdateCollection", expectations))
	}

	return nil
}

func (b *Collection) ValidateRemoveRequestDTO(dto *dto.CollectionRemoveAnimeRequestDTO) error {
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
		return b.logger.LogPropagate(errtype.NewValidatorError("UpdateCollection", expectations))
	}

	return nil
}

func (b *Collection) ValidateCreateRequestDTO(dto *dto.CollectionCreateRequestDTO) error {
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
		return b.logger.LogPropagate(errtype.NewValidatorError("UpdateCollection", expectations))
	}

	return nil
}

func (b *Collection) ValidateListRequestDTO(dto *dto.CollectionListRequestDTO) error {
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

func (b *Collection) ValidateAnimeListFromCollectionRequestDTO(dto *dto.AnimeListFromCollectionRequestDTO) error {
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

func (b *Collection) ValidateDetailRequestDTO(dto *dto.CollectionDetailRequestDTO) error {
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
