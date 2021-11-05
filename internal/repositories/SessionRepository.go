package repositories

import (
	"AuthBeatsPro/internal/db"
	"AuthBeatsPro/internal/models"
	"fmt"
)

type SessionRepository struct {
	store db.Store
}

func NewSessionRepository(store db.Store) *SessionRepository {
	return &SessionRepository{
		store: store,
	}
}

func (repo *SessionRepository) CreateSession(session *models.Session) (int, error) {
	var id int64

	query := fmt.Sprintf("INSERT INTO %s (`refreshToken`, `expiresAt`, `userId`) VALUES (?, ?, ?);", SESSION_TABLE)
	result, err := repo.store.Exec(query, session.RefreshToken, session.ExpiresAt, session.UserId)

	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil

}
