package services

import (
	"epseed/internal/models"
)

type UserService struct{}

func (s *UserService) GetUsers() ([]*models.User, error) {
	return models.GetUsers()
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return models.GetUserByEmail(email)
}
