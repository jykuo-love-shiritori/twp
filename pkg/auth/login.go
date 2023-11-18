package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	params := make(map[string]any)

	c.Bind(&params)

	redirect_uri := params["redirect_uri"].(string)

	b, _ := json.MarshalIndent(params, "", "  ")
	fmt.Println(string(b))

	challenge = params["code_challenge"].(string)

	_code := make([]byte, 32)
	_, err := rand.Read(_code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "",
		})
	}

	sha := sha256.Sum256(_code)
	code = base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sha[:])

	return c.Redirect(http.StatusFound, redirect_uri+"?code="+code)
}
