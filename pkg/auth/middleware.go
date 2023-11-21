package auth

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const tokenPrefix string = "Bearer "

func IsRole(db *db.DB, logger *zap.SugaredLogger, role constants.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "No token found",
				})
			}

			if !strings.HasPrefix(authorization, tokenPrefix) {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "Bad token",
				})
			}

			tokenString := strings.TrimPrefix(authorization, tokenPrefix)

			claims := &jwtCustomClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("TWP_SECRET_KEY")), nil
			})
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "Bad token",
				})
			}

			if !token.Valid {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": "Token validation failed",
				})
			}

			if role != claims.Role {
				return c.NoContent(http.StatusForbidden)
			}

			return next(c)
		}
	}
}
