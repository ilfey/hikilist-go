package authModels

import (
	"encoding/json"

	"github.com/ilfey/hikilist-go/internal/utils/errorsx"
)

// Модель токенов
type TokensModel struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// Преобразовать в json
func (m *TokensModel) JSON() []byte {
	return errorsx.Must(json.Marshal(m))
}
