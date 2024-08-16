package auth

// Конфиг аутентификации
type Config struct {
	Secret          []byte
	Issuer          string
	AccessLifeTime  int
	RefreshLifeTime int
}
