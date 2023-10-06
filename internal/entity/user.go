package entity

type User struct {
	Id       int64  `db:"id"`
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
