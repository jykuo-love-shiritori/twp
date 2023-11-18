package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Token(c echo.Context) error {
	params := make(map[string]any)

	c.Bind(&params)

	b, _ := json.MarshalIndent(params, "", "  ")
	fmt.Println(string(b))

	_code := params["code"]
	verifier := params["code_verifier"].(string)

	sha := sha256.Sum256([]byte(verifier))
	hash := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sha[:])

	fmt.Println(_code, hash)
	fmt.Println(code, challenge)

	if _code != code || hash != challenge {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token": "test",
	})
}
