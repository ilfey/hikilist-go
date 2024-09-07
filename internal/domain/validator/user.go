package validator

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/ilfey/hikilist-go/pkg/validator"
	"github.com/ilfey/hikilist-go/pkg/validator/options"
)

type User struct {
	log loggerInterface.Logger
}

func NewUser(log loggerInterface.Logger) *User {
	return &User{
		log: log,
	}
}

func (v *User) ValidateCreateRequestDTO(req *dto.UserCreateRequestDTO) error {
	expectations, ok := validator.Validate(
		req,
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
		return errtype.NewValidatorError("List", expectations)
	}

	return nil
}

func (v *User) ValidateListRequestDTO(req *dto.UserListRequestDTO) error {
	expectations, ok := validator.Validate(
		req,
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
					"username",
					"-username",
				}),
			},
		},
	)
	if !ok {
		return errtype.NewValidatorError("List", expectations)
	}

	return nil
}

func (v *User) ValidateDetailRequestDTO(req *dto.UserDetailRequestDTO) error {
	expectations, ok := validator.Validate(
		req,
		map[string][]options.Option{
			"UserID": {
				options.GreaterThan[uint64](0),
			},
		},
	)
	if !ok {
		return errtype.NewValidatorError("Detail", expectations)
	}

	return nil
}

func (v *User) ValidateMeRequestDTO(req *dto.UserMeRequestDTO) error {
	expectations, ok := validator.Validate(
		req,
		map[string][]options.Option{
			"UserID": {
				options.GreaterThan[uint64](0),
			},
		},
	)
	if !ok {
		return errtype.NewValidatorError("Me", expectations)
	}

	return nil
}

func (v *User) ValidateCollectionRequestDTO(req *dto.UserCollectionsRequestDTO) error {
	expectations, ok := validator.Validate(
		req,
		map[string][]options.Option{
			"UserID": {
				options.GreaterThan[uint64](0),
			},
		},
	)
	if !ok {
		return errtype.NewValidatorError("UserCollections", expectations)
	}

	return nil
}
