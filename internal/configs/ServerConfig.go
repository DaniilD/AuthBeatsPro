package configs

type ServerConfig struct {
	Host string
	Port int
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{}
}
