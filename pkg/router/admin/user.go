package admin

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary		Admin Get User
// @Description	Get all user information. Include user's icon, name, email, created time and role.
// @Tags			Admin, User
// @Produce		json
// @Param			offset	query		int	false	"Begin index"	default(0)
// @Param			limit	query		int	false	"limit"			default(10)	maximum(20)
// @Success		200		{array}		db.GetUsersRow
// @Failure		400		{object}	echo.HTTPError
// @Router			/admin/user [get]
func GetUser(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := common.NewQueryParams(0, 10)
		if err := c.Bind(&q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if q.Validate() != nil {
			logger.Errorw("invalid query parameter", "offset", q.Offset, "limit", q.Limit)
			return echo.NewHTTPError(http.StatusBadRequest, "offset or limit is invalid")
		}
		users, err := pg.Queries.GetUsers(c.Request().Context(), db.GetUsersParams{Offset: q.Offset, Limit: q.Limit})
		if err != nil {
			logger.Errorw("failed to get users", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range users {
			users[i].IconUrl = mc.GetFileURL(c.Request().Context(), users[i].IconUrl)
		}
		return c.JSON(http.StatusOK, users)
	}
}

// @Summary		Admin Disable User
// @Description	Disable user.
// @Tags			Admin, User
// @Produce		json
// @param			username	path		string	true	"Username"
// @Success		200			{string}	string	constants.SUCCESS
// @Failure		400			{object}	echo.HTTPError
// @Failure		404			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/admin/user/{username} [delete]
func DisableUser(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		if username == "" { // would not happen in future
			logger.Errorw("username is empty")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if execRows, err := pg.Queries.DisableUser(c.Request().Context(), username); err != nil {
			logger.Errorw("failed to disable user", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to disable user")
		} else if execRows == 0 {
			logger.Errorw("user not found", "username", username)
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}
