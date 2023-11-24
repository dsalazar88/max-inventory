package utils

import "courses/golang/inventory-project/internal/entity"

// The function `ValidRolesToAddProduct` is a method of the `MyUtils` struct. It takes in two
// parameters: `validRoles`, which is a slice of `int64` values, and `role`, which is an `int64` value.
func (u *MyUtils) ValidRolesToAddProduct(roles []entity.UserRole, validRoles []int64) bool {

	for _, r := range roles {
		for _, vr := range validRoles {
			if vr == r.RoleId {
				return true
			}
		}
	}

	return false
}
