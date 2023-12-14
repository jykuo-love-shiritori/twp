package router

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jykuo-love-shiritori/twp/db"
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
// @Param			limit		query		int		false	"limit"			default(10)
// @Success		200			{object}	shopInfo
// @Failure		400			{object}	echo.HTTPError
// @Failure		404			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/shop/{seller_name} [get]
func getShopInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
			if err == pgx.ErrNoRows {
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
		result.Products, err = pg.Queries.GetShopProducts(c.Request().Context(), db.GetShopProductsParams{
			Offset: qp.Offset, Limit: qp.Limit, SellerName: sellerName})

		if err != nil {
			logger.Errorw("failed to get shop info", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
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
// @Param			limit		query		int		false	"limit"			default(10)
// @Success		200			{array}		db.GetShopCouponsRow
// @Failure		400			{object}	echo.HTTPError
// @Failure		404			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/shop/{seller_name}/coupon [get]
func getShopCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
			if err == pgx.ErrNoRows {
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

// TODO
//
//	@Summary		Search Shop Products
//	@Description	Search products within a shop by seller username
//	@Tags			Shop,Product,Search
//	@Accept			json
//	@Produce		json
//	@Param			seller_name	path	int		true	"Seller username"
//	@Param			q			query	string	true	"search word"
//	@Success		200
//	@Failure		401
//	@Router			/shop/{seller_name}/search [get]
func searchShopProduct(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

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
func getTagInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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

// TODO
//
//	@Summary		Search for Products and Shops
//	@Description	Search for products and shops
//	@Tags			Search
//	@Accept			json
//	@Produce		json
//	@Param			q	query	string	true	"search word"
//	@Success		200
//	@Failure		401
//	@Router			/search [get]
func search(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// TODO
//
//	@Summary		Search for Shops by Name
//	@Description	Search for shops by name
//	@Tags			Search,Shop
//	@Accept			json
//	@Produce		json
//	@Param			q	query	string	true	"Search Name"
//	@Success		200
//	@Failure		401
//	@Router			/search/shop [get]
func searchShopByName(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// TODO
//
//	@Summary		Get News
//	@Description	Get news
//	@Tags			News
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		401
//	@Router			/news [get]
func getNews(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// TODO
//
//	@Summary		Get News Detail
//	@Description	Get details of a specific news item by ID
//	@Tags			News
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"News ID"
//	@Success		200
//	@Failure		401
//	@Router			/news/{id} [get]
func getNewsDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Get Discover
// @Description	Get discover content
// @Tags			Discover,Product
// @Accept			json
// @Produce		json
// @Param			offset	query	int	false	"Begin index"	default(0)
// @Param			limit	query	int	false	"limit"			default(10)
// @Success		200 {array} db.GetRandomProductsRow
// @Failure		500 {object} echo.HTTPError
// @Router			/discover [get]
func getDiscover(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
// @Success		200 {array} popular
// @Failure		500 {object} echo.HTTPError
// @Router			/popular [get]
func getPopular(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result popular
		var err error
		result.PopularProducts, err = pg.Queries.GetProductsFromPopularShop(c.Request().Context())
		if err != nil {
			logger.Errorw("failed to get popular products", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.LocalProducts, err = pg.Queries.GetProductsFromNearByShop(c.Request().Context())
		if err != nil {
			logger.Errorw("failed to get popular products", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
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
func getProductInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int32
		if err := echo.PathParamsBinder(c).Int32("id", &id).BindError(); err != nil {
			logger.Errorw("failed to parse id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		result, err := pg.Queries.GetProductInfo(c.Request().Context(), id)
		if err != nil {
			if err == pgx.ErrNoRows {
				return echo.NewHTTPError(http.StatusNotFound, "Product Not Found")
			}
			logger.Errorw("failed to get product info", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}
