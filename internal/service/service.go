package service

import (
	"context"
	"courses/golang/inventory-project/internal/models"
	"courses/golang/inventory-project/internal/repository"
	"courses/golang/inventory-project/internal/utils"
	"errors"
)

// Message for ERRORS
var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrRoleAlreadyAdded   = errors.New("role was already added for this user")
	ErrRoleNotFound       = errors.New("roles was not found")
)

var (
	MyUtils utils.MyUtils
)

// The above type defines a service interface with methods for user registration and login, managing
// user roles, and retrieving products.
// @property {error} RegisterUser - This method is used to register a new user. It takes the user's
// email, name, and password as parameters and returns an error if the registration fails.
// @property LoginUser - This method is used to authenticate a user by their email and password. It
// takes the email and password as input parameters and returns the authenticated user object along
// with an error if any.
// @property {error} SaveUserRole - This method is used to save a user's role in the system. It takes
// the user ID and role ID as parameters and returns an error if there is any issue saving the user's
// role.
// @property {error} RemoveUserRole - The RemoveUserRole method is used to remove a specific role from
// a user. It takes the user ID and role ID as parameters and returns an error if the operation fails.
// @property GetProducts - This method is used to retrieve a list of all products available in the
// system. It returns an array of models.Product and an error if any occurred during the retrieval
// process.
// @property GetProduct - This method is used to retrieve a specific product by its ID. It takes a
// context.Context object and an int64 value representing the product ID as parameters. It returns a
// pointer to a models.Product object and an error.
//
//go:generate mockery --name=Service --output=test
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)

	SaveUserRole(ctx context.Context, userId, roleId int64) error
	RemoveUserRole(ctx context.Context, userId, roleId int64) error

	AddProduct(ctx context.Context, product models.Product, email string) error
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProduct(ctx context.Context, id int64) (*models.Product, error)
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
