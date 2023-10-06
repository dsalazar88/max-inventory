package repository

import (
	"context"
	"courses/golang/inventory-project/internal/entity"
)

const (
	queryInsertProduct = `
		inser into PRODUCTS (name, description, price, createdBy)
			values (?, ?, ?, ?);
	`

	queryGetProducts = `
		select id, name, description, price, created_by
		from PRODUCTS;
	`

	queryGetProductById = `
	select id, name, description, price, created_by
	from PRODUCTS
	where id = ?;
`
)

// The `SaveProduct` function is responsible for saving a new product into the database. It takes the
// necessary parameters such as the product name, description, price, and the ID of the user who
// created the product.
func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, queryInsertProduct, name, description, price, createdBy)
	if err != nil {
		return err
	}

	return nil
}

// The `GetProducts` function is responsible for retrieving all products from the database. It takes
// the necessary parameters such as the context and returns a slice of `entity.Product` and an error.
func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {
	pp := []entity.Product{}

	err := r.db.SelectContext(ctx, &pp, queryGetProducts)
	if err != nil {
		return nil, err
	}

	return pp, nil
}

// The `GetProduct` function is responsible for retrieving a specific product from the database based
// on its ID. It takes the necessary parameters such as the context and the ID of the product.
func (r *repo) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	p := &entity.Product{}

	err := r.db.GetContext(ctx, &p, queryGetProductById, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}
