package configs

var configLocator *ConfigLocator

type ConfigLocator struct {
	dbConfig  *DBConfig
}

func GetConfigLocator() *ConfigLocator {
	if configLocator == nil {
		configLocator = &ConfigLocator{}
	}
	
	return configLocator
}

func (locator *ConfigLocator) DBConfigInstance() *DBConfig {
	if configLocator.dbConfig == nil {
		configLocator.dbConfig = NewDBConfig()
	}
	
	return configLocator.dbConfig
}
