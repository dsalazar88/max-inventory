package entity

type Product struct {
	Id          int64   `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float32 `db:"price"`
	CreatedBy   int64   `db:"created_by"`
}
