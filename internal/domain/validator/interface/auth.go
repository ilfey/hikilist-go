package validatorInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type Auth interface {
	//ValidateChangePasswordRequestDTO(req *dto.AuthChangePasswordRequestDTO) error
	//ValidateChangeUsernameRequestDTO(req *dto.AuthChangeUsernameRequestDTO) error

	ValidateDeleteRequestDTO(req *dto.UserDeleteRequestDTO) error
	ValidateLoginRequestDTO(req *dto.AuthLoginRequestDTO) error
	ValidateLogoutRequestDTO(req *dto.AuthLogoutRequestDTO) error
	ValidateRefreshRequestDTO(req *dto.AuthRefreshRequestDTO) error
	ValidateRegisterRequestDTO(req *dto.AuthRegisterRequestDTO) error
}
