package configs

var configLocator *ConfigLocator

type ConfigLocator struct {
	dbConfig     *DBConfig
	serverConfig *ServerConfig
	jwtConfig    *JWTConfig
}

func GetConfigLocator() *ConfigLocator {
	if configLocator == nil {
		configLocator = &ConfigLocator{}
	}

	return configLocator
}

func (locator *ConfigLocator) DBConfigInstance() *DBConfig {
	if locator.dbConfig == nil {
		locator.dbConfig = NewDBConfig()
	}

	return locator.dbConfig
}

func (locator *ConfigLocator) ServerConfigInstance() *ServerConfig {
	if locator.serverConfig == nil {
		locator.serverConfig = NewServerConfig()
	}

	return locator.serverConfig
}

func (locator *ConfigLocator) JWTConfig() *JWTConfig {
	if locator.jwtConfig == nil {
		locator.jwtConfig = NewJWTConfig()
	}

	return locator.jwtConfig
}
