package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
)

type jwtCustomClaims struct {
	Username string
	Role     constants.Role
	jwt.RegisteredClaims
}
