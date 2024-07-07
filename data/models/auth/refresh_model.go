package authModels

import (
	"encoding/json"
	"net/http"
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
