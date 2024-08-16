package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
)

type LogoutModel struct {
	Refresh string `json:"refresh"`
}

func LogoutModelFromRequest(request *http.Request) *LogoutModel {
	model := new(LogoutModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}

func (m LogoutModel) Validate() error {
	return validator.Validate(
		m,
		map[string][]options.Option{
			"Refresh": {
				options.Required(),
			},
		},
	)
}
