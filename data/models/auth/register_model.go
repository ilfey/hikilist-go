package authModels

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
	"github.com/ilfey/hikilist-go/internal/validator/options"
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
func (m RegisterModel) Validate() error {
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
