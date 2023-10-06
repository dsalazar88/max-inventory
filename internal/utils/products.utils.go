package utils

func (u *MyUtils) ValidRolesToAddProduct(validRoles []int64, role int64) bool {

	for _, vr := range validRoles {
		if vr == role {
			return true
		}
	}

	return false
}
