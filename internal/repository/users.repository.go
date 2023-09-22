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
	return u, err
}
