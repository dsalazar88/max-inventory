package service

import (
	"context"
	"courses/golang/inventory-project/internal/models"
	"errors"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// The `RegisterUser` function is responsible for registering a new user in the system. Here's a
// breakdown of what it does:
func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	//TODO has password
	return s.repo.SaveUser(ctx, email, name, password)
}

// The `LoginUser` function is responsible for authenticating a user by checking their email and
// password against the stored user data. Here's a breakdown of what it does:
func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//TODO decrypt password

	if u.Password != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		Id:    u.Id,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
