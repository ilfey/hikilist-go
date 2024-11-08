package tests

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/validator"
	"github.com/ilfey/hikilist-go/pkg/logger"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ActionValidatorSuite struct {
	suite.Suite

	Validator *validator.Action
}

func (s *ActionValidatorSuite) SetupTest() {
	s.Validator = validator.NewAction(logger.NewTest(s.T()))
}

func (s *ActionValidatorSuite) TestValidatePaginate() {
	testCases := []struct {
		desc    string
		req     *dto.ActionListRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.ActionListRequestDTO{
				UserID: 0,
				Page:   1,
				Limit:  10,
				Order:  "-id",
			},
			isValid: true,
		},
		{
			desc: "Invalid page",
			req: &dto.ActionListRequestDTO{
				UserID: 0,
				Page:   0,
				Limit:  10,
				Order:  "-id",
			},
			isValid: false,
		},
		{
			desc: "Invalid limit",
			req: &dto.ActionListRequestDTO{
				UserID: 0,
				Page:   1,
				Limit:  0,
				Order:  "-id",
			},
			isValid: false,
		},
		{
			desc:    "Empty request",
			req:     &dto.ActionListRequestDTO{},
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

func TestActionValidatorSuite(t *testing.T) {
	suite.Run(t, new(ActionValidatorSuite))
}
