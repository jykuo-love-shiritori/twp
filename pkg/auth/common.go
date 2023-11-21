package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
)

type codeChallengeMethod string

const (
	s256  codeChallengeMethod = "S256"
	plain codeChallengeMethod = "plain"
)

type responseType string

const (
	code  responseType = "code"
	token responseType = "token"
)

type jwtCustomClaims struct {
	Username string         `json:"username"`
	Role     constants.Role `json:"role"`
	jwt.RegisteredClaims
}
