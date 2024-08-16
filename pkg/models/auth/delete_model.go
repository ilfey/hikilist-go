package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
)

type DeleteModel struct {
	Refresh  string `json:"refresh"`
	Password string `json:"password"`
}

func (m DeleteModel) Validate() error {
	return validator.Validate(
		m,
		map[string][]options.Option{
			"Refresh": {
				options.Required(),
			},
			"Password": {
				options.LenLessThan(32),
				options.LenGreaterThan(5),
			},
		},
	)
}

func DeleteModelFromRequest(request *http.Request) *DeleteModel {
	model := new(DeleteModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}
