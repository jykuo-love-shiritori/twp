package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Login(db *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		params := make(map[string]any)

		c.Bind(&params)

		redirect_uri := params["redirect_uri"].(string)

		b, _ := json.MarshalIndent(params, "", "  ")
		fmt.Println(string(b))

		challenge := params["code_challenge"].(string)

		buf := make([]byte, 32)
		_, err := rand.Read(buf)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Unexpected Error",
			})
		}
		sha := sha256.Sum256(buf)
		code := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sha[:])

		mu.Lock()
		codeChallengePairs[code] = challenge
		mu.Unlock()

		return c.Redirect(http.StatusFound, redirect_uri+"?code="+code)
	}
}
