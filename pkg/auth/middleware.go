package auth

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const tokenPrefix string = "Bearer "

// Validate JWT only
func ValidateJwt(pg *db.DB, logger *zap.SugaredLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "No token found")
			}

			if !strings.HasPrefix(authorization, tokenPrefix) {
				return echo.NewHTTPError(http.StatusBadRequest, "Bad token")
			}

			tokenString := strings.TrimPrefix(authorization, tokenPrefix)

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("TWP_JWT_SECRET")), nil
			})
			if err != nil {
				logger.Error(err)
				return echo.NewHTTPError(http.StatusBadRequest, "Bad token")
			}

			if !token.Valid {
				return echo.NewHTTPError(http.StatusForbidden, "Token validation failed")
			}

			return next(c)
		}
	}
}

// Validate and parse JWT
func IsRole(pg *db.DB, logger *zap.SugaredLogger, role db.RoleType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "No token found")
			}

			if !strings.HasPrefix(authorization, tokenPrefix) {
				return echo.NewHTTPError(http.StatusBadRequest, "Bad token")
			}

			tokenString := strings.TrimPrefix(authorization, tokenPrefix)

			claims := jwtCustomClaims{}
			token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("TWP_JWT_SECRET")), nil
			})
			if err != nil {
				logger.Error(err)
				return echo.NewHTTPError(http.StatusBadRequest, "Bad token")
			}

			if !token.Valid {
				return echo.NewHTTPError(http.StatusForbidden, "Token validation failed")
			}

			if role != claims.Role {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}

			return next(c)
		}
	}
}
