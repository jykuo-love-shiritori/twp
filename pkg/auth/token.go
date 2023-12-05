package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type tokenParams struct {
	ClientId     string `form:"client_id" json:"client_id"`
	Code         string `form:"code" json:"code"`
	CodeVerifier string `form:"code_verifier" json:"code_verifier"`
	GrantType    string `form:"grant_type" json:"grant_type"`
}

func Token(db *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var params tokenParams
		err := c.Bind(&params)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse data")
		}

		fmt.Printf("%+v\n", params)
		code := strings.TrimSpace(params.Code)
		verifier := strings.TrimSpace(params.CodeVerifier)

		mu.Lock()
		user, found := codeChallengePairs[code]
		// delete(codeChallengePairs, params.Code)
		mu.Unlock()

		fmt.Println(found)

		if !verifyCodeChallenge(verifier, user) {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		claims := &jwtCustomClaims{
			user.Username,
			user.Role,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("TWP_JWT_SECRET")))
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"access_token": tokenString,
		})
	}
}

func verifyCodeChallenge(verifier string, challenge challengeUser) bool {
	switch challenge.ChallengeMethod {
	case s256:
		sha := sha256.Sum256([]byte(verifier))
		hash := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sha[:])
		fmt.Println(hash)
		fmt.Println(challenge.CodeChallenge)
		return challenge.CodeChallenge == hash

	case plain:
		return challenge.CodeChallenge == verifier

	default:
		return false
	}
}
