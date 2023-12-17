package auth

import (
	"net/http"
	"time"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary		Logout
// @Description	Logout the current user
// @Tags			Auth
// @Produce		json
// @Success		200
// @Failure		500
// @Router			/oauth/logout [post]
func Logout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenCookie, err := c.Cookie("refresh_token")
		if err != nil {
			logger.Errorln(err)
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		err = pg.Queries.DeleteRefreshToken(c.Request().Context(), tokenCookie.Value)
		if err != nil {
			logger.Errorln(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		// delete httpOnly cookie by resetting expire time to 0
		tokenCookie.Expires = time.Unix(0, 0)
		tokenCookie.SameSite = http.SameSiteStrictMode
		c.SetCookie(tokenCookie)

		return c.NoContent(http.StatusOK)
	}
}
