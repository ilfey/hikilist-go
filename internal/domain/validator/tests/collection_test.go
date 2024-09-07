package tests

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/validator"
	"github.com/ilfey/hikilist-go/pkg/logger"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CollectionValidatorSuite struct {
	suite.Suite

	Validator *validator.Collection
}

func (s *CollectionValidatorSuite) SetupTest() {
	s.Validator = validator.NewCollection(logger.NewTest(s.T()))
}

func (s *CollectionValidatorSuite) TestValidateRemoveAnimeRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.CollectionRemoveAnimeRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.CollectionRemoveAnimeRequestDTO{
				Animes: []uint64{1},
			},
			isValid: true,
		},
		{
			desc:    "Empty request",
			req:     &dto.CollectionRemoveAnimeRequestDTO{},
			isValid: false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateRemoveAnimeRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *CollectionValidatorSuite) TestValidateAddAnimeRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.CollectionAddAnimeRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.CollectionAddAnimeRequestDTO{
				Animes: []uint64{1},
			},
			isValid: true,
		},
		{
			desc:    "Empty request",
			req:     &dto.CollectionAddAnimeRequestDTO{},
			isValid: false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateAddAnimeRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *CollectionValidatorSuite) TestValidateCreateRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.CollectionCreateRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.CollectionCreateRequestDTO{
				Title:       "title",
				Description: ptr(string(make([]byte, 128))),
			},
			isValid: true,
		},
		{
			desc: "Invalid title",
			req: &dto.CollectionCreateRequestDTO{
				Title:       "",
				Description: ptr(string(make([]byte, 128))),
			},
			isValid: false,
		},
		{
			desc: "Invalid description",
			req: &dto.CollectionCreateRequestDTO{
				Title:       "title",
				Description: ptr(""),
			},
			isValid: false,
		},
		{
			desc:    "Empty request",
			req:     &dto.CollectionCreateRequestDTO{},
			isValid: false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateCreateRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *CollectionValidatorSuite) TestValidateUpdateRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.CollectionUpdateRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.CollectionUpdateRequestDTO{
				Title:       ptr("title"),
				Description: ptr(string(make([]byte, 128))),
			},
			isValid: true,
		},
		{
			desc: "Invalid title",
			req: &dto.CollectionUpdateRequestDTO{
				Title:       ptr(""),
				Description: ptr(string(make([]byte, 128))),
			},
			isValid: false,
		},
		{
			desc: "Invalid description",
			req: &dto.CollectionUpdateRequestDTO{
				Title:       ptr("title"),
				Description: ptr(""),
			},
			isValid: false,
		},
		{
			desc:    "Empty request",
			req:     &dto.CollectionUpdateRequestDTO{},
			isValid: true,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateUpdateRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *CollectionValidatorSuite) TestValidateListRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.CollectionListRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.CollectionListRequestDTO{
				Page:  1,
				Limit: 10,
				Order: "-id",
			},
			isValid: true,
		},
		{
			desc: "Invalid offset",
			req: &dto.CollectionListRequestDTO{
				Page:  0,
				Limit: 10,
				Order: "-id",
			},
			isValid: false,
		},
		{
			desc: "Invalid limit",
			req: &dto.CollectionListRequestDTO{
				Page:  1,
				Limit: 0,
				Order: "-id",
			},
			isValid: false,
		},
		{
			desc:    "Empty request",
			req:     &dto.CollectionListRequestDTO{},
			isValid: false,
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

func TestCollectionValidatorSuite(t *testing.T) {
	suite.Run(t, new(CollectionValidatorSuite))
}
