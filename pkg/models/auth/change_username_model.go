package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/sirupsen/logrus"
)

type ChangeUsernameModel struct {
	Username string `json:"username"`
}

func (m ChangeUsernameModel) Validate() error {
	return validator.Validate(
		m,
		map[string][]options.Option{
			"Username": {
				options.LenLessThan(32),
				options.LenGreaterThan(3),
			},
		},
	)
}

func ChangeUsernameModelFromRequest(request *http.Request) *ChangeUsernameModel {
	model := new(ChangeUsernameModel)

	err := json.NewDecoder(request.Body).Decode(model)
	if err != nil {
		logrus.Infof("Error occurred while decoding ChangeUsernameModel %v", err)
	}

	return model
}
