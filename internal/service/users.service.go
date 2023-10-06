package service

import (
	"context"
	"courses/golang/inventory-project/encryption"
	"courses/golang/inventory-project/internal/models"
	"errors"
	"log"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrRoleAlreadyAdded   = errors.New("role was already added for this user")
	ErrRoleNotFound       = errors.New("roles was not found")
)

// The `RegisterUser` function is responsible for registering a new user in the system. Here's a
// breakdown of what it does:
func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, pass)
}

// The `LoginUser` function is responsible for authenticating a user by checking their email and
// password against the stored user data. Here's a breakdown of what it does:
func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		Id:    u.Id,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

// The `SaveUserRole` function is responsible for saving a user's role in the system. Here's a
// breakdown of what it does:
func (s *serv) SaveUserRole(ctx context.Context, userId, roleId int64) error {
	roles, err := s.repo.GetUserRoles(ctx, userId)
	if err != nil {
		return err
	}
	log.Println(roles)

	for _, r := range roles {
		if r.RoleId == roleId {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userId, roleId)
}

// The `RemoveUserRole` function is responsible for removing a user's role in the system. Here's a
// breakdown of what it does:
func (s *serv) RemoveUserRole(ctx context.Context, userId, roleId int64) error {
	roles, err := s.repo.GetUserRoles(ctx, userId)
	if err != nil {
		return err
	}
	roleFound := false
	log.Println(roles)
	for _, r := range roles {
		if r.RoleId == roleId {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}

	return s.repo.RemoveUserRole(ctx, userId, roleId)
}
