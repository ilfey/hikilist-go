package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

type Anime struct {
	log loggerInterface.Logger
}

func NewAnime(container diInterface.ServiceContainer) (*Anime, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	return &Anime{
		log: log,
	}, nil
}

func (v *Anime) ValidateCreateRequestDTO(dto *dto.AnimeCreateRequestDTO) error {
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
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil
}

func (v *Anime) ValidateDetailRequestDTO(dto *dto.AnimeDetailRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"CollectionID": {
				options.GreaterThan[uint64](0),
			},
		},
	)
	if !ok {
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil
}

func (v *Anime) ValidateListRequestDTO(dto *dto.AnimeListRequestDTO) error {
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
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil
}
