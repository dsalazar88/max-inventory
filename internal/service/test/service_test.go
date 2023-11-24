package test

import (
	"courses/golang/inventory-project/encryption"
	"courses/golang/inventory-project/internal/entity"
	"courses/golang/inventory-project/internal/repository"
	"courses/golang/inventory-project/internal/service"
	"errors"
	"os"
	"testing"

	mock "github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s service.Service

var (
	ErrInvalidPermissions = errors.New("no have permissions for add product")
)

// The TestMain function sets up mock repository methods for testing and runs the test suite.
func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)

	u := &entity.User{
		Id:       1,
		Email:    "test@exists.com",
		Password: encryptedPassword,
	}

	adminUser := &entity.User{
		Id:       1,
		Email:    "admin@email.com",
		Password: encryptedPassword,
	}

	customerUser := &entity.User{
		Id:       2,
		Email:    "customer@email.com",
		Password: encryptedPassword,
	}

	repo = &repository.MockRepository{}

	//Test for users
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(u, nil)
	repo.On("GetUserByEmail", mock.Anything, "admin@email.com").Return(adminUser, nil)
	repo.On("GetUserByEmail", mock.Anything, "customer@email.com").Return(customerUser, nil)

	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	//Test for roles
	repo.On("GetUserRoles", mock.Anything, int64(1)).Return([]entity.UserRole{
		{
			UserId: 1,
			RoleId: 1,
		},
	}, nil)
	repo.On("GetUserRoles", mock.Anything, int64(2)).Return([]entity.UserRole{
		{
			UserId: 2,
			RoleId: 3,
		},
	}, nil)

	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	//Test for products
	repo.On("SaveProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = service.New(repo)

	code := m.Run()
	os.Exit(code)
}
