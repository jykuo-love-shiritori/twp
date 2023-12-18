package seller

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary		Seller get available tag
// @Description	Get all available tags for shop.
// @Tags			Seller, Shop, Tag
// @Param			name	query	string	true	"search tag name start with"	minlength(1)
// @Produce		json
// @success		200	{array}		db.SellerSearchTagRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/tag [get]
func GetTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"
		var tagPerPage int32 = 20

		var param db.SellerSearchTagParams
		param.Name = c.QueryParam("name")
		if param.Name == "" || common.HasRegexSpecialChars(param.Name) {
			logger.Errorw("search tag name is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.SellerName = username
		param.Name = "^" + param.Name
		param.Limit = tagPerPage
		tags, err := pg.Queries.SellerSearchTag(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, tags)
	}
}

// @Summary		Seller add tag
// @Description	Add tag for shop.
// @Tags			Seller, Shop, Tag
// @Accept			json
// @Param			name	body	string	true	"insert tag"	minlength(1)
// @Produce		json
// @success		200	{object}	db.SellerInsertTagRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		409	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/tag [post]
func AddTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.HaveTagNameParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		if param.Name == "" || common.HasRegexSpecialChars(param.Name) {
			logger.Errorw("tag name is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.SellerName = username
		have, err := pg.Queries.HaveTagName(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if have {
			logger.Errorw("tag name not unique")
			return echo.NewHTTPError(http.StatusConflict)
		}
		tag, err := pg.Queries.SellerInsertTag(c.Request().Context(), db.SellerInsertTagParams(param))
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, tag)
	}
}
