package tests

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/validator"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuthValidatorSuite struct {
	suite.Suite

	Validator *validator.Auth
}

func (s *AuthValidatorSuite) SetupTest() {
	s.Validator = &validator.Auth{}
}

// TODO: remove this test case.
func (s *AuthValidatorSuite) TestValidateAuthChangePasswordRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.AuthChangePasswordRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.AuthChangePasswordRequestDTO{
				OldPassword: "oldPassword",
				NewPassword: "newPassword",
			},
			isValid: true,
		},
		{
			desc: "Invalid old password",
			req: &dto.AuthChangePasswordRequestDTO{
				OldPassword: "",
				NewPassword: "newPassword",
			},
			isValid: false,
		},
		{
			desc: "Invalid new password",
			req: &dto.AuthChangePasswordRequestDTO{
				OldPassword: "oldPassword",
				NewPassword: "",
			},
			isValid: false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateChangePasswordRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

// TODO: remove this test case.
func (s *AuthValidatorSuite) TestValidateChangeUsernameRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.AuthChangeUsernameRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.AuthChangeUsernameRequestDTO{
				NewUsername: "username",
				Password:    "password",
			},
			isValid: true,
		},
		{
			desc: "Invalid new username",
			req: &dto.AuthChangeUsernameRequestDTO{
				NewUsername: "",
				Password:    "password",
			},
			isValid: false,
		},
		{
			desc: "Invalid password",
			req: &dto.AuthChangeUsernameRequestDTO{
				NewUsername: "username",
				Password:    "",
			},
			isValid: false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateChangeUsernameRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *AuthValidatorSuite) TestValidateDeleteRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.UserDeleteRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.UserDeleteRequestDTO{
				Refresh:  "refresh",
				Password: "password",
			},
			isValid: true,
		},
		{
			desc: "Invalid password",
			req: &dto.UserDeleteRequestDTO{
				Refresh:  "refresh",
				Password: "",
			},
			isValid: false,
		},
		{
			desc: "Invalid refresh",
			req: &dto.UserDeleteRequestDTO{
				Refresh:  "",
				Password: "password",
			},
			isValid: false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateDeleteRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *AuthValidatorSuite) TestValidateLoginRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.AuthLoginRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.AuthLoginRequestDTO{
				Username: "username",
				Password: "password",
			},
			isValid: true,
		},
		{
			desc: "Invalid username",
			req: &dto.AuthLoginRequestDTO{
				Username: "",
				Password: "password",
			},
			isValid: false,
		},
		{
			desc: "Invalid password",
			req: &dto.AuthLoginRequestDTO{
				Username: "username",
				Password: "",
			},
			isValid: false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateLoginRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *AuthValidatorSuite) TestValidateLogoutRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.AuthLogoutRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.AuthLogoutRequestDTO{
				Refresh: "refresh",
			},
			isValid: true,
		},
		{
			desc: "Invalid refresh",
			req: &dto.AuthLogoutRequestDTO{
				Refresh: "",
			},
			isValid: false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateLogoutRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *AuthValidatorSuite) TestValidateRegisterRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.AuthRegisterRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.AuthRegisterRequestDTO{
				Username: "username",
				Password: "password",
			},
			isValid: true,
		},
		{
			desc: "Invalid username",
			req: &dto.AuthRegisterRequestDTO{
				Username: "",
				Password: "password",
			},
			isValid: false,
		},
		{
			desc: "Invalid password",
			req: &dto.AuthRegisterRequestDTO{
				Username: "username",
				Password: "",
			},
			isValid: false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateRegisterRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *AuthValidatorSuite) TestValidateRefreshRequestDTO() {
	testCases := []struct {
		desc    string
		req     *dto.AuthRefreshRequestDTO
		isValid bool
	}{
		{
			desc: "Valid request",
			req: &dto.AuthRefreshRequestDTO{
				Refresh: "refresh",
			},
			isValid: true,
		},
		{
			desc: "Invalid refresh",
			req: &dto.AuthRefreshRequestDTO{
				Refresh: "",
			},
			isValid: false,
		},
	}

	for _, tC := range testCases {
		s.Run(tC.desc, func() {
			err := s.Validator.ValidateRefreshRequestDTO(tC.req)
			if tC.isValid {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func TestAuthValidatorSuite(t *testing.T) {
	suite.Run(t, new(AuthValidatorSuite))
}
