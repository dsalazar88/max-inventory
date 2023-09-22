package entity

type User struct {
	Id       int64  `db:"id"`
	Email    string `db:"name"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
