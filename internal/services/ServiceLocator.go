package services

import (
	"AuthBeatsPro/internal/JWT"
	"AuthBeatsPro/internal/configs"
	"AuthBeatsPro/internal/db"
	"AuthBeatsPro/internal/repositories"
)

var serviceLocator *ServiceLocator

type ServiceLocator struct {
	sessionService *SessionService
	userService    *UserService
	authService    *AuthService
}

func GetServiceLocator() *ServiceLocator {
	if serviceLocator == nil {
		serviceLocator = &ServiceLocator{}
	}

	return serviceLocator
}

func (locator *ServiceLocator) SessionService() *SessionService {
	if locator.sessionService == nil {
		locator.sessionService = NewSessionService(
			repositories.NewSessionRepository(db.GetStoreLocator().GetStore()))
	}

	return locator.sessionService
}

func (locator *ServiceLocator) UserService() *UserService {
	if locator.userService == nil {
		locator.userService = NewUserService(
			repositories.NewUserRepository(db.GetStoreLocator().GetStore()))
	}

	return locator.userService
}

func (locator *ServiceLocator) AuthService() *AuthService {
	if locator.authService == nil {
		locator.authService = NewAuthService(
			repositories.NewUserRepository(db.GetStoreLocator().GetStore()),
			repositories.NewSessionRepository(db.GetStoreLocator().GetStore()),
			JWT.NewJWTHelper(configs.GetConfigLocator().JWTConfig()),
		)
	}

	return locator.authService
}
