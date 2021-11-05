package operations

import "AuthBeatsPro/internal/services"

var operationLocator *OperationLocator

type OperationLocator struct {
	loginOperation *LoginOperation
}

func GetOperationLocator() *OperationLocator {
	if operationLocator == nil {
		operationLocator = &OperationLocator{}
	}

	return operationLocator
}

func (locator *OperationLocator) GetLoginOperation() *LoginOperation {
	if locator.loginOperation == nil {
		locator.loginOperation = &LoginOperation{
			authService: services.GetServiceLocator().AuthService(),
		}
	}

	return locator.loginOperation
}
