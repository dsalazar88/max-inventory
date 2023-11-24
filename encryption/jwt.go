package encryption

import (
	"courses/golang/inventory-project/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

// The function SignedLoginToken generates a signed JWT token using the user's email and name.
func SignedLoginToken(u *models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.Name,
	})

	return token.SignedString([]byte(key))
}
