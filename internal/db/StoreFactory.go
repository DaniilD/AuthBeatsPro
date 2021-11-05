package db

import "AuthBeatsPro/internal/configs"

type StoreFactory struct {
	config *configs.DBConfig
}

func NewStoreFactory(config *configs.DBConfig) *StoreFactory {
	return &StoreFactory{
		config: config,
	}
}

func (storeFactory *StoreFactory) makeStore() Store {
	switch driver := storeFactory.config.Driver; driver {
	case MYSQL:
		return NewMysqlStore(storeFactory.config)
	default:
		return NewMysqlStore(storeFactory.config)
	}
}
