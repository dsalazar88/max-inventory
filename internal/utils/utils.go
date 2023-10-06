package utils

type Utils interface {
	ValidRolesToAddProduct(validRoles []int64, role int64) bool
}

type MyUtils struct{}
