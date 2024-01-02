package general

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/image"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type GetNewsInfo struct {
	ID      int32  `json:"id"`
	Title   string `json:"news"`
	ImageID string `json:"image_id"`
}

// @Summary		Get News
// @Description	Get news
// @Tags			News
// @Accept			json
// @Produce		json
// @Success		200	{array}		common.NewsInfo
// @Failure		400	{object}	echo.HTTPError
// @Router			/news [get]
func GetNews(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, common.GetNewsInfo())
	}
}

// @Summary		Get News Detail
// @Description	Get details of a specific news item by ID
// @Tags			News
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"News ID"
// @Success		200	{object}	common.News
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Router			/news/{id} [get]
func GetNewsDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int32
		if err := echo.PathParamsBinder(c).Int32("id", &id).BindError(); err != nil {
			logger.Errorw("failed to parse id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		result, err := common.GetNews(id)
		if err != nil {
			logger.Errorw("failed to get news detail", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "News Not Found")
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Get Discover
// @Description	Get discover content
// @Tags			Discover,Product
// @Accept			json
// @Produce		json
// @Param			offset	query		int	false	"Begin index"	default(0)
// @Param			limit	query		int	false	"limit"			default(10)	Maximum(20)
// @Success		200		{array}		db.GetRandomProductsRow
// @Failure		500		{object}	echo.HTTPError
// @Router			/discover [get]
func GetDiscover(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := common.NewQueryParams(0, 10)
		if err := c.Bind(&q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := q.Validate(); err != nil {
			logger.Errorw("failed to validate query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "query parameter is invalid")
		}
		result, err := pg.Queries.GetRandomProducts(c.Request().Context(), db.GetRandomProductsParams{Offset: q.Offset, Limit: q.Limit})
		if err != nil {
			logger.Errorw("failed to get discover", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range result {
			result[i].ImageUrl = image.GetUrl(result[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, result)
	}
}

type popular struct {
	PopularProducts []db.GetProductsFromPopularShopRow `json:"popular_products"`
	LocalProducts   []db.GetProductsFromNearByShopRow  `json:"local_products"`
}

// @Summary		Get Popular products and Local products
// @Description	Get discover content
// @Tags			Discover, Product
// @Accept			json
// @Produce		json
// @Success		200	{array}		popular
// @Failure		500	{object}	echo.HTTPError
// @Router			/popular [get]
func GetPopular(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result popular
		var err error
		result.PopularProducts, err = pg.Queries.GetProductsFromPopularShop(c.Request().Context())
		if err != nil {
			logger.Errorw("failed to get popular products", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range result.PopularProducts {
			result.PopularProducts[i].ImageUrl = image.GetUrl(result.PopularProducts[i].ImageUrl)
		}
		result.LocalProducts, err = pg.Queries.GetProductsFromNearByShop(c.Request().Context())
		if err != nil {
			logger.Errorw("failed to get popular products", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range result.LocalProducts {
			result.LocalProducts[i].ImageUrl = mc.GetFileURL(c.Request().Context(), result.LocalProducts[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Get Product Info
// @Description	Get product information with product ID
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Product ID"
// @Success		200	{object}	db.GetProductInfoRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/product/{id} [get]
func GetProductInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int32
		if err := echo.PathParamsBinder(c).Int32("id", &id).BindError(); err != nil {
			logger.Errorw("failed to parse id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		result, err := pg.Queries.GetProductInfo(c.Request().Context(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return echo.NewHTTPError(http.StatusNotFound, "Product Not Found")
			}
			logger.Errorw("failed to get product info", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.ImageUrl = image.GetUrl(result.ImageUrl)
		return c.JSON(http.StatusOK, result)
	}
}
