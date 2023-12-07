package auth

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jykuo-love-shiritori/twp/db"
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
	Username string      `json:"username"`
	Role     db.RoleType `json:"role"`
	jwt.RegisteredClaims
}

func generateRandomString(len int) (string, error) {
	buf := make([]byte, len)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(buf), nil
}
