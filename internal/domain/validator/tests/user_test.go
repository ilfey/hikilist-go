package tests

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/validator"
	"github.com/ilfey/hikilist-go/pkg/logger"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserValidatorSuite struct {
	suite.Suite

	Validator *validator.User
}

func (s *UserValidatorSuite) SetupTest() {
	s.Validator = validator.NewUser(logger.NewTest(s.T()))
}

func (s *UserValidatorSuite) TestValidateListRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.UserListRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.UserListRequestDTO{
				Page:  1,
				Limit: 10,
			},
			isValid: true,
		},
		{
			desc: "Invalid offset",
			req: &dto.UserListRequestDTO{
				Page:  0,
				Limit: 10,
			},
			isValid: false,
		},
		{
			desc: "Invalid limit",
			req: &dto.UserListRequestDTO{
				Page:  1,
				Limit: 0,
			},
			isValid: false,
		},
		{
			desc:    "Empty request",
			req:     &dto.UserListRequestDTO{},
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

func TestUserValidatorSuite(t *testing.T) {
	suite.Run(t, new(UserValidatorSuite))
}
