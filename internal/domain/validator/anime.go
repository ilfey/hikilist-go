package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

type Anime struct {
	logger loggerInterface.Logger
}

func NewAnime(logger loggerInterface.Logger) *Anime {
	return &Anime{
		logger: logger,
	}
}

func (b *Anime) ValidateCreateRequestDTO(dto *dto.AnimeCreateRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Title": {
				options.LenGreaterThan(3),
				options.LenLessThan(256),
			},
			"Description": {
				options.IfNotNil(
					options.LenLessThan(4096),
				),
			},
			"Poster": {
				options.IfNotNil(
					options.LenLessThan(256),
				),
			},
			"MalID": {
				options.IfNotNil(
					options.GreaterThan[uint64](0),
				),
			},
			"ShikiID": {
				options.IfNotNil(
					options.GreaterThan[uint64](0),
				),
			},
		},
	)
	if !ok {
		return errtype.NewValidatorError("CreateAnime", expectations)
	}

	return nil
}

func (b *Anime) ValidateDetailRequestDTO(dto *dto.AnimeDetailRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"CollectionID": {
				options.GreaterThan[uint64](0),
			},
		},
	)
	if !ok {
		return errtype.NewValidatorError("GetAnime", expectations)
	}

	return nil
}

func (b *Anime) ValidateListRequestDTO(dto *dto.AnimeListRequestDTO) error {
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
					"title",
					"-title",
					"episodes",
					"-episodes",
					"episodes_released",
					"-episodes_released",
				}),
			},
		},
	)
	if !ok {
		return errtype.NewValidatorError("GetAnimeList", expectations)
	}

	return nil
}
