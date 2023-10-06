package repository

import (
	"context"
	"courses/golang/inventory-project/internal/entity"
)

const (
	queryInsertUser = `
		insert into USERS (email, name, password)
		values (?, ?, ?);
	`

	queryGetUserByEmail = `
		select 
			id,
			email,
			name,
			password
		from USERS
		where email = ?;
	`

	queryInsertUserRole = `
		insert into USER_ROLES (user_id, role_id) 
			values (?, ?);
	`

	queryRemoveUserRole = `
		delete from USER_ROLES 
		where user_id = ? and role_id = ?;
	`

	queryGetUserRoles = `
		select user_id, role_id 
		from USER_ROLES
		where user_id = ?;
	`
)

// The `SaveUser` function is a method of the `repo` struct in the `repository` package. It takes a
// context, email, name, and password as input parameters and returns an error.
func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)
	return err
}

// The `GetUserByEmail` function is a method of the `repo` struct in the `repository` package. It takes
// a context and an email as input parameters and returns a pointer to an `entity.User` struct and an
// error.
func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}

	err := r.db.GetContext(ctx, u, queryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}

	return u, err
}

// The `SaveUserRole` function is a method of the `repo` struct in the `repository` package. It takes a
// context, userId, and roleId as input parameters and returns an error.
func (r *repo) SaveUserRole(ctx context.Context, userId, roleId int64) error {
	_, err := r.db.ExecContext(ctx, queryInsertUserRole, userId, roleId)

	return err
}

// The `RemoveUserRole` function is a method of the `repo` struct in the `repository` package. It takes
// a context, userId, and roleId as input parameters and returns an error.
func (r *repo) RemoveUserRole(ctx context.Context, userId, roleId int64) error {
	_, err := r.db.ExecContext(ctx, queryRemoveUserRole, userId, roleId)

	return err
}

// The `GetUserRoles` function is a method of the `repo` struct in the `repository` package. It takes a
// context and a userId as input parameters and returns a slice of `entity.UserRole` structs and an
// error.
func (r *repo) GetUserRoles(ctx context.Context, userId int64) ([]entity.UserRole, error) {
	roles := []entity.UserRole{}

	err := r.db.SelectContext(ctx, &roles, queryGetUserRoles, userId)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
