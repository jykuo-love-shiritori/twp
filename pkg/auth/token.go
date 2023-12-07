package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type tokenParams struct {
	ClientId     string `form:"client_id" json:"client_id"`
	Code         string `form:"code" json:"code"`
	CodeVerifier string `form:"code_verifier" json:"code_verifier"`
	GrantType    string `form:"grant_type" json:"grant_type"`
}

func Token(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var params tokenParams
		err := c.Bind(&params)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse data")
		}

		code := strings.TrimSpace(params.Code)
		verifier := strings.TrimSpace(params.CodeVerifier)

		mu.Lock()
		user, found := codeChallengePairs[code]
		delete(codeChallengePairs, params.Code)
		mu.Unlock()
		if !found {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		if !verifyCodeChallenge(verifier, user) {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		accessToken, err := generateAccessToken(user.Username, user.Role)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		refreshToken, err := generateRandomString(32)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		expireDate := time.Now().Add(30 * 24 * time.Hour)
		err = pg.Queries.SetRefreshToken(c.Request().Context(), db.SetRefreshTokenParams{
			Username:     user.Username,
			RefreshToken: refreshToken,
			ExpireDate:   pgtype.Timestamptz{Time: expireDate, Valid: true},
		})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		cookie := new(http.Cookie)
		cookie.Name = "refresh_token"
		cookie.Value = refreshToken
		cookie.Secure = common.IsEnv(constants.PROD)
		cookie.Expires = expireDate
		cookie.HttpOnly = true
		cookie.SameSite = http.SameSiteStrictMode
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, echo.Map{
			"access_token": accessToken,
		})
	}
}

func verifyCodeChallenge(verifier string, challenge challengeUser) bool {
	switch challenge.ChallengeMethod {
	case s256:
		sha := sha256.Sum256([]byte(verifier))
		hash := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sha[:])
		return challenge.CodeChallenge == hash

	case plain:
		return challenge.CodeChallenge == verifier

	default:
		return false
	}
}

func generateAccessToken(username string, role db.RoleType) (string, error) {
	claims := &jwtCustomClaims{
		username,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("TWP_JWT_SECRET")))
}
