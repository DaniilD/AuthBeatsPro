package operations

import (
	"AuthBeatsPro/internal/requests"
	"AuthBeatsPro/internal/response"
	"AuthBeatsPro/internal/services"
	"encoding/json"
	"net/http"
)

type LoginOperation struct {
	authService *services.AuthService
}

func NewLoginOperation(authService *services.AuthService) *LoginOperation {
	return &LoginOperation{
		authService: authService,
	}
}

func (operation *LoginOperation) Handle(r *http.Request) (interface{}, error) {
	var loginRequest requests.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		return nil, err
	}

	tokens, err := operation.authService.Login(loginRequest.Email, loginRequest.Password)

	if err != nil {
		return nil, err
	}

	loginResponse := response.NewLoginResponse(tokens.AccessToken, tokens.RefreshToken)

	return loginResponse, nil
}
