package authModels

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
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
func (m RegisterModel) Validate() validator.ValidateError {
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
