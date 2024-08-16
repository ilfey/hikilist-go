package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
)

type ChangePasswordModel struct {
	OldPassword    string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (cpm ChangePasswordModel) Validate() error {
	return validator.Validate(
		cpm,
		map[string][]options.Option{
			"OldPassword": {
				options.LenLessThan(32),
				options.LenGreaterThan(5),
			},
			"NewPassword": {
				options.LenLessThan(32),
				options.LenGreaterThan(5),
			},
		},
	)
}

func ChangePasswordModelFromRequest(request *http.Request) *ChangePasswordModel {
	model := new(ChangePasswordModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}
