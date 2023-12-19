package auth

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Refresh(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenCookie, err := c.Cookie("refresh_token")
		if err != nil {
			logger.Errorln(err)
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		user, err := pg.Queries.FindUserByRefreshToken(c.Request().Context(), tokenCookie.Value)
		if errors.Is(err, pgx.ErrNoRows) {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		if err != nil {
			logger.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		accessToken, err := generateAccessToken(user.Username, user.Role)
		if err != nil {
			logger.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"access_token": accessToken,
		})
	}
}
