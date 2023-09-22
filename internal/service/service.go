package service

import (
	"context"
	"courses/golang/inventory-project/internal/models"
	"courses/golang/inventory-project/internal/repository"
)

// The above type defines a service interface for user registration and login operations.
// @property {error} RegisterUser - A method that registers a user with the provided email, name, and
// password. It returns an error if the registration fails.
//
// @property LoginUser - The LoginUser method is used to authenticate a user by their email and
// password. It takes the context, email, and password as input parameters and returns a pointer to the
// User model and an error.
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

// The New function returns a new instance of the Service interface with the provided repository.
func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
