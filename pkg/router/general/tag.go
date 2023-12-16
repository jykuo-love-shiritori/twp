package general

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary		Get Tag Info
// @Description	Get information about a tag by tag ID
// @Tags			Tag
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Tag ID"
// @Success		200	{object}	db.GetTagInfoRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/tag/{id} [get]
func GetTagInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int32
		if err := echo.PathParamsBinder(c).Int32("id", &id).BindError(); err != nil {
			logger.Errorw("failed to parse id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		result, err := pg.Queries.GetTagInfo(c.Request().Context(), id)
		if err != nil {
			if err == pgx.ErrNoRows {
				return echo.NewHTTPError(http.StatusNotFound, "Tag Not Found")
			}
			logger.Errorw("failed to get tag info", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}
