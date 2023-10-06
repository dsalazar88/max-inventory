package entity

type UserRole struct {
	Id     int64 `db:"id"`
	UserId int64 `db:"user_id"`
	RoleId int64 `db:"role_id"`
}
