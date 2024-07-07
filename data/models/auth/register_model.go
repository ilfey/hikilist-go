package authModels

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/validator"
)

// Модель регистрации
type RegisterModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Собрать модель `RegisterModel` из `http.Request`
func RegisterModelFromRequest(request *http.Request) *RegisterModel {
	model := new(RegisterModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}

// Валидация модели
func (m RegisterModel) Validate() ([]byte, bool) {
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
		return nil, true
	}

	return e.JSON(), false
}
