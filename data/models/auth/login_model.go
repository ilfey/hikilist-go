package authModels

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/validator"
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
func (m LoginModel) Validate() ([]byte, bool) {
	e := validator.ValidateStruct(
		m,
		map[string][]validator.Option{
			"Username": {
				validator.LenLessThan(32),
				validator.LenGreaterThan(3),
			},
			"Password": {
				validator.LenLessThan(32),
				validator.LenGreaterThan(5),
			},
		},
	)

	if e.Success() {
		return []byte{}, true
	}

	return e.JSON(), false
}
