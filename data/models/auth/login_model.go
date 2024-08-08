package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
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
