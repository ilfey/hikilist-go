package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
	"github.com/sirupsen/logrus"
)

// Модель логина
type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Собрать модель `LoginModel` из `http.Request`
func LoginModelFromRequest(request *http.Request) *LoginModel {
	model := new(LoginModel)

	err := json.NewDecoder(request.Body).Decode(model)
	if err != nil {
		logrus.Infof("Error occurred while decoding LoginModel %v", err)
	}

	return model
}

// Валидация модели
func (m LoginModel) Validate() error {
	return validator.Validate(
		m,
		map[string][]options.Option{
			"Username": {
				options.LenLessThan(32),
				options.LenGreaterThan(3),
			},
			"Password": {
				options.LenLessThan(32),
				options.LenGreaterThan(5),
			},
		},
	)
}
