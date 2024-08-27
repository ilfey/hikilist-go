package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/sirupsen/logrus"
)

type LogoutModel struct {
	Refresh string `json:"refresh"`
}

func LogoutModelFromRequest(request *http.Request) *LogoutModel {
	model := new(LogoutModel)

	err := json.NewDecoder(request.Body).Decode(model)
	if err != nil {
		logrus.Infof("Error occurred while decoding LogoutModel %v", err)
	}

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
