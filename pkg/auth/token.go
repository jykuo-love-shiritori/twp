package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
)

func Token(c echo.Context) error {
	params := make(map[string]any)

	c.Bind(&params)

	b, _ := json.MarshalIndent(params, "", "  ")
	fmt.Println(string(b))

	code := params["code"].(string)
	verifier := params["code_verifier"].(string)

	sha := sha256.Sum256([]byte(verifier))
	challenge := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sha[:])

	mu.Lock()
	match := codeChallengePairs[code] == challenge
	mu.Unlock()

	if !match {
		return c.NoContent(http.StatusUnauthorized)
	}

	claims := &jwtCustomClaims{
		"someone", // TODO: remove hard-coded value
		constants.ADMIN,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := unsignedToken.SignedString([]byte(os.Getenv("TWP_JWT_SECRET")))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	fmt.Println(token)

	return c.JSON(http.StatusOK, echo.Map{
		"access_token": token,
	})
}
