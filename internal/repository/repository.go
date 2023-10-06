package repository

import (
	"context"
	"courses/golang/inventory-project/internal/entity"

	"github.com/jmoiron/sqlx"
)

// The Repository interface defines methods for saving and retrieving user data.
// @property {error} SaveUser - This method is used to save a user in the repository. It takes the
// user's email, name, and password as parameters and returns an error if there is any issue while
// saving the user.
//
// @property GetUserByEmail - This method is used to retrieve a user from the repository based on their
// email address. It takes a context object and the email as parameters and returns a User entity and
// an error.
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)

	SaveUserRole(ctx context.Context, userId, roleId int64) error
	RemoveUserRole(ctx context.Context, userId, roleId int64) error
	GetUserRoles(ctx context.Context, userId int64) ([]entity.UserRole, error)

	SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error
	GetProducts(ctx context.Context) ([]entity.Product, error)
	GetProduct(ctx context.Context, id int64) (*entity.Product, error)
}

type repo struct {
	db *sqlx.DB
}

// The New function returns a new instance of a repository with the given database connection.
func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
