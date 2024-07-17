package authModels

// Модель токенов
type TokensModel struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
