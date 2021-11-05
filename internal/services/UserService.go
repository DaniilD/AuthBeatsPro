package services

import (
	"AuthBeatsPro/internal/models"
	"AuthBeatsPro/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) GetByCredentials(email string, password string) (*models.User, error) {
	return service.userRepository.GetByCredentials(email, password)
}

func (service *UserService) CreateUser(user *models.User) (int, error) {
	return service.userRepository.CreateUser(user)
}

func (service *UserService) UpdateUser(user *models.User) error {
	return service.userRepository.UpdateUser(user)
}

func (service *UserService) GetById(id int) (*models.User, error) {
	return service.userRepository.GetById(id)
}
