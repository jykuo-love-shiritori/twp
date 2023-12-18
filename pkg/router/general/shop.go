package general

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type shopInfo struct {
	Info     db.GetShopInfoRow       `json:"info"`
	Products []db.GetShopProductsRow `json:"products"`
}

// @Summary		Get Shop Info
// @Description	Get shop information with seller username
// @Tags			Shop
// @Accept			json
// @Produce		json
// @Param			seller_name	path		string	true	"seller username"
// @Param			offset		query		int		false	"Begin index"	default(0)
// @Param			limit		query		int		false	"limit"			default(10)	Maximum(20)
// @Success		200			{object}	shopInfo
// @Failure		400			{object}	echo.HTTPError
// @Failure		404			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/shop/{seller_name} [get]
func GetShopInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		qp := common.NewQueryParams(0, 10)
		sellerName := c.Param("seller_name")
		if sellerName == "" {
			logger.Errorw("seller_name is empty")
			return echo.NewHTTPError(http.StatusBadRequest, "seller_name is empty")
		}
		if err := c.Bind(&qp); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := qp.Validate(); err != nil {
			logger.Errorw("failed to validate query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "query parameter is invalid")
		}
		if _, err := pg.Queries.ShopExists(c.Request().Context(), sellerName); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return echo.NewHTTPError(http.StatusNotFound, "Shop Not Found")
			}
			logger.Errorw("failed to check shop exists", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		var result shopInfo
		var err error
		result.Info, err = pg.Queries.GetShopInfo(c.Request().Context(), sellerName)
		if err != nil {
			logger.Errorw("failed to get shop info", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.Info.ImageUrl = mc.GetFileURL(c.Request().Context(), result.Info.ImageUrl)
		result.Products, err = pg.Queries.GetShopProducts(c.Request().Context(), db.GetShopProductsParams{
			Offset: qp.Offset, Limit: qp.Limit, SellerName: sellerName})
		if err != nil {
			logger.Errorw("failed to get shop info", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range result.Products {
			result.Products[i].ImageUrl = mc.GetFileURL(c.Request().Context(), result.Products[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Get Shop Coupons
// @Description	Get coupons for a shop with seller username
// @Tags			Shop,Coupon
// @Accept			json
// @Produce		json
// @Param			seller_name	path		string	true	"seller username"
// @Param			offset		query		int		false	"Begin index"	default(0)
// @Param			limit		query		int		false	"limit"			default(10)	Maximum(20)
// @Success		200			{array}		db.GetShopCouponsRow
// @Failure		400			{object}	echo.HTTPError
// @Failure		404			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/shop/{seller_name}/coupon [get]
func GetShopCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		seller_name := c.Param("seller_name")
		if seller_name == "" {
			logger.Errorw("seller_name is empty")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		q := common.NewQueryParams(0, 10)
		if err := c.Bind(&q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := q.Validate(); err != nil {
			logger.Errorw("failed to validate query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "query parameter is invalid")
		}
		shop_id, err := pg.Queries.ShopExists(c.Request().Context(), seller_name)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return echo.NewHTTPError(http.StatusNotFound, "Shop Not Found")
			}
			logger.Errorw("failed to check shop exists", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		coupons, err := pg.Queries.GetShopCoupons(c.Request().Context(),
			db.GetShopCouponsParams{Offset: q.Offset, Limit: q.Limit, ShopID: pgtype.Int4{Int32: shop_id, Valid: true}})
		if err != nil {
			logger.Errorw("failed to get shop coupons", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, coupons)
	}
}
