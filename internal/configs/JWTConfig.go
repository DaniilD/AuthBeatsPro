package configs

type JWTConfig struct {
	SigningKey string
}

func NewJWTConfig() *JWTConfig {
	return &JWTConfig{}
}
