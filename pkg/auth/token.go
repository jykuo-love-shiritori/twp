package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type tokenParams struct {
	ClientId     string `form:"client_id" json:"client_id"`
	Code         string `form:"code" json:"code"`
	CodeVerifier string `form:"code_verifier" json:"code_verifier"`
	GrantType    string `form:"grant_type" json:"grant_type"`
	RedirectUri  string `form:"redirect_uri" json:"redirect_uri"`
}

func Token(db *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var params tokenParams
		err := c.Bind(&params)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse data")
		}

		mu.Lock()
		challenge := codeChallengePairs[params.Code]
		mu.Unlock()

		if !verifyCodeChallenge(params.CodeVerifier, challenge) {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		claims := &jwtCustomClaims{
			"someone", // TODO: remove hard-coded value
			constants.ADMIN,
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

func verifyCodeChallenge(verifier string, challenge challenge) bool {
	switch challenge.challengeMethod {
	case s256:
		sha := sha256.Sum256([]byte(verifier))
		hash := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sha[:])
		return challenge.challengeString == hash

	case plain:
		return challenge.challengeString == verifier

	default:
		return false
	}
}
