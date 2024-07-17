package authModels

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
)

// Модель логина
type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Собрать модель `LoginModel` из `http.Request`
func LoginModelFromRequest(request *http.Request) *LoginModel {
	model := new(LoginModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}

// Валидация модели
func (m LoginModel) Validate() validator.ValidateError {
	return validator.Validate(
		m,
		map[string][]validator.Option{
			"Username": {
				validator.LenLessThat(32),
				validator.LenGreaterThat(3),
			},
			"Password": {
				validator.LenLessThat(32),
				validator.LenGreaterThat(5),
			},
		},
	)
}
