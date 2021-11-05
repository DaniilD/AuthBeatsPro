package services

import (
	"AuthBeatsPro/internal/JWT"
	"AuthBeatsPro/internal/models"
	"AuthBeatsPro/internal/repositories"
	"time"
)

type AuthService struct {
	userRepository    *repositories.UserRepository
	sessionRepository *repositories.SessionRepository
	jwtHelper         *JWT.JWTHelper
}

func NewAuthService(userRepository *repositories.UserRepository,
	sessionRepository *repositories.SessionRepository,
	jwtHelper *JWT.JWTHelper) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		jwtHelper:         jwtHelper,
	}
}

func (service *AuthService) Login(email string, password string) (*JWT.Tokens, error) {
	response := new(JWT.Tokens)

	user, err := service.userRepository.GetByCredentials(email, password)

	if err != nil {
		return nil, err
	}

	token, err := service.jwtHelper.CreateToken(user.Id)

	if err != nil {
		return nil, err
	}

	refreshToken, err := service.jwtHelper.NewRefreshToken()

	if err != nil {
		return nil, err
	}

	response.AccessToken = token
	response.RefreshToken = refreshToken

	session := new(models.Session)
	session.RefreshToken = refreshToken
	session.ExpiresAt = time.Now()

	_, err = service.sessionRepository.CreateSession(session)

	if err != nil {
		return nil, err
	}

	return response, nil
}
