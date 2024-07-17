package authModels

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/hikilist-go/internal/validator"
)

// Модель обновления токена
type RefreshModel struct {
	Refresh string `json:"refresh"`
}

// Собрать модель `RefreshModel` из `http.Request`
func RefreshModelFromRequest(request *http.Request) *RefreshModel {
	model := new(RefreshModel)

	json.NewDecoder(request.Body).Decode(model)

	return model
}

// Валидация модели
func (m RefreshModel) Validate() validator.ValidateError {
	return validator.Validate(
		m,
		map[string][]validator.Option{
			"Refresh": {
				validator.Required(),
			},
		},
	)
}
