package tests

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/validator"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AnimeValidatorSuite struct {
	suite.Suite

	Validator *validator.Anime
}

func (s *AnimeValidatorSuite) SetupTest() {
	s.Validator = &validator.Anime{
		// TODO: Provide logger.
	}
}

func (s *AnimeValidatorSuite) TestValidatePaginate() {
	testCases := []struct {
		desc    string
		req     *dto.AnimeListRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.AnimeListRequestDTO{
				Page:  1,
				Limit: 10,
				Order: "-id",
			},
			isValid: true,
		},
		{
			desc: "Invalid offset",
			req: &dto.AnimeListRequestDTO{
				Page:  -1,
				Limit: 10,
				Order: "-id",
			},
			isValid: false,
		},
		{
			desc: "Invalid limit",
			req: &dto.AnimeListRequestDTO{
				Page:  0,
				Limit: -1,
				Order: "-id",
			},
			isValid: false,
		},
		{
			desc:    "Empty request",
			req:     &dto.AnimeListRequestDTO{},
			isValid: true,
		},
	}
	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateListRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)

				return
			}

			s.Error(err)
		})
	}
}

func (s *AnimeValidatorSuite) TestValidateCreateModel() {
	testCases := []struct {
		desc    string
		model   *dto.AnimeCreateRequestDTO
		isValid bool
	}{
		{
			desc: "Model with only title",
			model: &dto.AnimeCreateRequestDTO{
				Title: "test",
			},
			isValid: true,
		},
		{
			desc: "Model with title and description",
			model: &dto.AnimeCreateRequestDTO{
				Title:       "test",
				Description: ptr("test"),
			},
			isValid: true,
		},
		{
			desc: "Model with empty title",
			model: &dto.AnimeCreateRequestDTO{
				Title: "",
			},
			isValid: false,
		},
		{
			desc: "Model with long title",
			model: &dto.AnimeCreateRequestDTO{
				Title: string(make([]byte, 256)),
			},
			isValid: false,
		},
		{
			desc:    "Empty model",
			model:   &dto.AnimeCreateRequestDTO{},
			isValid: false,
		},
		{
			desc: "Model with mal_id",
			model: &dto.AnimeCreateRequestDTO{
				Title: "test",
				MalID: ptr(uint64(1)),
			},
			isValid: true,
		},
		{
			desc: "Model with shiki_id",
			model: &dto.AnimeCreateRequestDTO{
				Title:   "test",
				ShikiID: ptr(uint64(1)),
			},
			isValid: true,
		},
		{
			desc: "Model with mal_id and shiki_id",
			model: &dto.AnimeCreateRequestDTO{
				Title:   "test",
				MalID:   ptr(uint64(1)),
				ShikiID: ptr(uint64(1)),
			},
			isValid: true,
		},
	}
	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateCreateRequestDTO(tC.model)
			if tC.isValid {
				s.NoError(err)

				return
			}

			s.Error(err)
		})
	}
}

func TestAnimeValidatorSuite(t *testing.T) {
	suite.Run(t, new(AnimeValidatorSuite))
}
