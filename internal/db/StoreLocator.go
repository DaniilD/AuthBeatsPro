package db

import "AuthBeatsPro/internal/configs"

var storeLocator *StoreLocator

type StoreLocator struct {
	storeFactory *StoreFactory
	store        Store
}

func GetStoreLocator() *StoreLocator {
	if storeLocator == nil {
		storeLocator = &StoreLocator{
			storeFactory: NewStoreFactory(configs.GetConfigLocator().DBConfigInstance()),
		}
	}

	return storeLocator
}

func (storeLocator *StoreLocator) GetStore() Store {
	if storeLocator.store == nil {
		storeLocator.store = storeLocator.storeFactory.makeStore()
	}

	return storeLocator.store
}
