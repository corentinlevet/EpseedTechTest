package services

import (
	"epseed/internal/models"
	"errors"
)

type AuthService struct{}

func (s *AuthService) Login(username string, password string) (*models.User, error) {
	user, err := models.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Signup(username string, email string, password string) error {
	// Check if user already exists
	user, _ := models.GetUserByEmail(email)
	if user != nil {
		var e error = errors.New("User already exists")
		return e
	}

	err := models.CreateUser(username, email, password)
	if err != nil {
		return err
	}

	return nil
}
