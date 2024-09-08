package validatorInterface

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
)

type User interface {
	ValidateCreateRequestDTO(req *dto.UserCreateRequestDTO) error
	ValidateListRequestDTO(req *dto.UserListRequestDTO) error
	ValidateDetailRequestDTO(req *dto.UserDetailRequestDTO) error
	ValidateMeRequestDTO(req *dto.UserMeRequestDTO) error
	ValidateCollectionListRequestDTO(req *dto.UserCollectionListRequestDTO) error
	//TODO: Add ValidateUpdateRequestDTO(req *dto.UserUpdateRequestDTO) error
}
