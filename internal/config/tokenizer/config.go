package tokenizer

// Config is the tokenizer configuration.
type Config struct {
	Salt []byte // Jwt token salt.

	Issuer string // Issuer of the token.

	AccessLifeTime  int // Access token lifetime.
	RefreshLifeTime int // Refresh token lifetime.
}
