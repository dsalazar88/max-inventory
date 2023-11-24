package service

import (
	"context"
	"courses/golang/inventory-project/internal/models"
	"errors"
)

var (
	validRolesToAddProduct []int64 = []int64{1, 2}

	ErrInvalidPermissions = errors.New("user does not have permission to add product")
)

// The `GetProducts` function is a method of a service struct (`serv`). It takes a `context.Context` as
// a parameter and returns a slice of `models.Product` and an error.
func (s *serv) GetProducts(ctx context.Context) ([]models.Product, error) {
	pp, err := s.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := []models.Product{}

	for _, p := range pp {
		products = append(products, models.Product{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}

	return products, nil
}

// The `GetProduct` function is a method of a service struct (`serv`). It takes a `context.Context` and
// an `int64` as parameters and returns a pointer to a `models.Product` and an error.
func (s *serv) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	p, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	product := &models.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}

	return product, nil
}

// The `AddProduct` function is a method of a service struct (`serv`). It takes a `context.Context`, a
// `models.Product`, and an email as parameters and returns an error.
func (s *serv) AddProduct(ctx context.Context, product models.Product, email string) error {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	roles, err := s.repo.GetUserRoles(ctx, u.Id)
	if err != nil {
		return err
	}

	userCanAdd := MyUtils.ValidRolesToAddProduct(roles, validRolesToAddProduct)

	if !userCanAdd {
		return ErrInvalidCredentials
	}

	return s.repo.SaveProduct(ctx, product.Name, product.Description, product.Price, u.Id)
}
