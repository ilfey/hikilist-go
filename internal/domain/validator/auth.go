package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

type Auth struct {
	logger loggerInterface.Logger
}

func NewAuth(logger loggerInterface.Logger) *Auth {
	return &Auth{
		logger: logger,
	}
}

//func (b *Auth) ValidateChangePasswordRequestDTO(dto *dto.AuthChangePasswordRequestDTO) error {
//	expectations, ok := validator.Validate(
//		dto,
//		map[string][]options.Option{
//			"OldPassword": {
//				options.Required(),
//				options.LenLessThan(32),
//				options.LenGreaterThan(5),
//			},
//			"NewPassword": {
//				options.Required(),
//				options.LenLessThan(32),
//				options.LenGreaterThan(5),
//			},
//		},
//	)
//	if !ok {
//		return errtype.NewValidatorError("ChangePassword", expectations)
//	}
//
//	return nil
//}
//
//func (b *Auth) ValidateChangeUsernameRequestDTO(dto *dto.AuthChangeUsernameRequestDTO) error {
//	expectations, ok := validator.Validate(
//		dto,
//		map[string][]options.Option{
//			"NewUsername": {
//				options.Required(),
//				options.LenLessThan(32),
//				options.LenGreaterThan(3),
//			},
//			"Password": {
//				options.Required(),
//				options.LenLessThan(32),
//				options.LenGreaterThan(5),
//			},
//		},
//	)
//
//	if !ok {
//		return errtype.NewValidatorError("ChangeUsername", expectations)
//	}
//
//	return nil
//
//}

func (b *Auth) ValidateDeleteRequestDTO(dto *dto.UserDeleteRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Refresh": {
				options.Required(),
			},
			"Password": {
				options.Required(),
				options.LenLessThan(32),
				options.LenGreaterThan(5),
			},
		},
	)

	if !ok {
		return errtype.NewValidatorError("Delete", expectations)
	}

	return nil

}

func (b *Auth) ValidateLoginRequestDTO(dto *dto.AuthLoginRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Username": {
				options.Required(),
				options.LenLessThan(32),
				options.LenGreaterThan(3),
			},
			"Password": {
				options.Required(),
				options.LenLessThan(32),
				options.LenGreaterThan(5),
			},
		},
	)

	if !ok {
		return errtype.NewValidatorError("Login", expectations)
	}

	return nil

}

func (b *Auth) ValidateLogoutRequestDTO(dto *dto.AuthLogoutRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Refresh": {
				options.Required(),
			},
		},
	)
	if !ok {
		return errtype.NewValidatorError("Logout", expectations)
	}

	return nil

}

func (b *Auth) ValidateRefreshRequestDTO(dto *dto.AuthRefreshRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Refresh": {
				options.Required(),
			},
		},
	)

	if !ok {
		return errtype.NewValidatorError("Refresh", expectations)
	}

	return nil

}

func (b *Auth) ValidateRegisterRequestDTO(dto *dto.AuthRegisterRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Username": {
				options.Required(),
				options.LenLessThan(32),
				options.LenGreaterThan(3),
			},
			"Password": {
				options.Required(),
				options.LenLessThan(32),
				options.LenGreaterThan(5),
			},
		},
	)
	if !ok {
		return errtype.NewValidatorError("Register", expectations)
	}

	return nil

}
