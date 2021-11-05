package services

import (
	"AuthBeatsPro/internal/models"
	"AuthBeatsPro/internal/repositories"
)

type SessionService struct {
	sessionRepository *repositories.SessionRepository
}

func NewSessionService(sessionRepository *repositories.SessionRepository) *SessionService {
	return &SessionService{
		sessionRepository: sessionRepository,
	}
}

func (service *SessionService) CreateSession(session *models.Session) (int, error) {
	return service.sessionRepository.CreateSession(session)
}
