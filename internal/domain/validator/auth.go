package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	diInterface "github.com/ilfey/hikilist-go/internal/domain/service/di/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

type Auth struct {
	log loggerInterface.Logger
}

func NewAuth(container diInterface.ServiceContainer) (*Auth, error) {
	log, err := container.GetLogger()
	if err != nil {
		return nil, err
	}

	return &Auth{
		log: log,
	}, nil
}

func (v *Auth) ValidateDeleteRequestDTO(dto *dto.UserDeleteRequestDTO) error {
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
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil

}

func (v *Auth) ValidateLoginRequestDTO(dto *dto.AuthLoginRequestDTO) error {
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
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil

}

func (v *Auth) ValidateLogoutRequestDTO(dto *dto.AuthLogoutRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Refresh": {
				options.Required(),
			},
		},
	)
	if !ok {
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil

}

func (v *Auth) ValidateRefreshRequestDTO(dto *dto.AuthRefreshRequestDTO) error {
	expectations, ok := validator.Validate(
		dto,
		map[string][]options.Option{
			"Refresh": {
				options.Required(),
			},
		},
	)

	if !ok {
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil

}

func (v *Auth) ValidateRegisterRequestDTO(dto *dto.AuthRegisterRequestDTO) error {
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
		return v.log.Propagate(errtype.NewValidatorError(expectations))
	}

	return nil

}
