package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"golang.org/x/crypto/bcrypt"
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

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
