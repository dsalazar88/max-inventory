package models

type User struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
