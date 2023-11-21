package auth

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary Logout
// @Description Logout the current user
// @Tags Auth
// @Produce json
// @Success 200
// @Failure 401
// @Router /logout [post]
func Logout(db *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}
