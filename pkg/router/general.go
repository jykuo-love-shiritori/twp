package router

import (
	"math/big"
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
func getShopInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
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

// @Summary		Search Shop Products
// @Description	Search products within a shop by seller username
// @Tags			Shop,Product,Search
// @Accept			json
// @Produce		json
// @Param			seller_name	path		string	true	"Seller username"
// @Param			q			query		string	true	"search query"
// @Param			minPrice	query		number	false	"price lower bound"
// @Param			maxPrice	query		number	false	"price upper bound"
// @Param			minStock	query		int		false	"stock lower bound"
// @Param			maxStock	query		int		false	"stock upper bound"
// @Param			haveCoupon	query		bool	false	"has coupon"
// @Param			sortBy		query		string	false	"sort by"		Enums("price", "stock", "sales", "relevancy")
// @Param			order		query		string	false	"sorting order"	Enums("asc", "desc")
// @Param			offset		query		int		false	"Begin index"	default(0)
// @Param			limit		query		int		false	"limit"			default(10)	Maximum(20)
// @Success		200			{object}	PrettierProductSearchResult
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/shop/{seller_name}/search [get]
func searchShopProduct(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		seller_name := c.Param("seller_name")
		if seller_name == "" {
			logger.Errorw("seller_name is empty")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		q := newSearchParams()
		if err := c.Bind(&q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := q.QueryParams.Validate(); err != nil {
			logger.Errorw("failed to validate query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "query parameter is invalid")
		}
		if q.Q == "" {
			logger.Errorw("search query is empty")
			return echo.NewHTTPError(http.StatusBadRequest, "search query is empty")
		}
		if q.MinPrice < 0 || q.MaxPrice < 0 || q.MinPrice > q.MaxPrice {
			logger.Errorw("invalid price range")
			return echo.NewHTTPError(http.StatusBadRequest, "invalid price range")
		}
		if q.SortBy != "" && q.SortBy != "price" && q.SortBy != "stock" && q.SortBy != "sales" && q.SortBy != "relevancy" {
			logger.Errorw("invalid sort by")
			return echo.NewHTTPError(http.StatusBadRequest, "invalid sort by")
		}
		if q.Order != "" && q.Order != "asc" && q.Order != "desc" {
			logger.Errorw("invalid order")
			return echo.NewHTTPError(http.StatusBadRequest, "invalid order")
		}
		dbq := db.SearchProductsByShopParams{
			SellerName: seller_name,
			Offset:     int32(q.Offset),
			Limit:      int32(q.Limit),
			Query:      q.Q,
			MinPrice: pgtype.Numeric{
				Int:   big.NewInt(int64(q.MinPrice)),
				Valid: true,
			},
			MaxPrice: pgtype.Numeric{
				Int:   big.NewInt(int64(q.MaxPrice)),
				Valid: true,
			},
			HasCoupon: pgtype.Bool{
				Bool:  q.HaveCoupon,
				Valid: true,
			},
			SortBy: q.SortBy,
			Order:  q.Order,
		}
		if c.QueryParam("minPrice") == "" {
			dbq.MinPrice.Valid = false
		}
		if c.QueryParam("maxPrice") == "" {
			dbq.MaxPrice.Valid = false
		}
		if c.QueryParam("minStock") == "" {
			dbq.MinStock.Valid = false
		}
		if c.QueryParam("maxStock") == "" {
			dbq.MaxStock.Valid = false
		}
		if c.QueryParam("haveCoupon") == "" {
			dbq.HasCoupon.Valid = false
		}
		products, err := pg.Queries.SearchProductsByShop(c.Request().Context(), dbq)
		if err != nil {
			logger.Errorw("failed to search products", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range products {
			products[i].ImageUrl = mc.GetFileURL(c.Request().Context(), products[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, products)
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

type searchParams struct {
	Q          string `query:"q"`
	MinPrice   int32  `query:"minPrice"`
	MaxPrice   int32  `query:"maxPrice"`
	HaveCoupon bool   `query:"haveCoupon"`
	SortBy     string `query:"sortBy"`
	Order      string `query:"order"`
	common.QueryParams
}

func newSearchParams() searchParams {
	return searchParams{HaveCoupon: false, SortBy: "relevancy", Order: "asc", QueryParams: common.NewQueryParams(0, 10)}
}

type PrettierProductSearchResult struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Price    int32  `json:"price"`
	Stock    int32  `json:"stock"`
	ImageUrl string `json:"image_url"`
}
type PrettierShopSearchResult struct {
	Id         int32  `json:"id"`
	Name       string `json:"name"`
	SellerName string `json:"seller_name"`
	ImageUrl   string `json:"image_url"`
}
type PrettierSearchResult struct {
	Products []PrettierProductSearchResult `json:"products"`
	Shops    []PrettierProductSearchResult `json:"shops"`
}
type searchResult struct {
	Products []db.SearchProductsRow `json:"products"`
	Shops    []db.SearchShopsRow    `json:"shops"`
}

// @Summary		Search for Products and Shops
// @Description	Search for products and shops
// @Tags			Search
// @Accept			json
// @Produce		json
// @Param			q			query		string	true	"search query"
// @Param			minPrice	query		number	false	"price lower bound"
// @Param			maxPrice	query		number	false	"price upper bound"
// @Param			minStock	query		int		false	"stock lower bound"
// @Param			maxStock	query		int		false	"stock upper bound"
// @Param			haveCoupon	query		bool	false	"has coupon"
// @Param			sortBy		query		string	false	"sort by"		Enums("price", "stock", "sales", "relevancy")
// @Param			order		query		string	false	"sorting order"	Enums("asc", "desc")
// @Param			offset		query		int		false	"Begin index"	default(0)
// @Param			limit		query		int		false	"limit"			default(10)	Maximum(20)
// @Success		200			{object}	PrettierSearchResult
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/search [get]
func search(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := newSearchParams()
		if err := c.Bind(&q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := q.QueryParams.Validate(); err != nil {
			logger.Errorw("failed to validate query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "query parameter is invalid")
		}
		if q.Q == "" {
			logger.Errorw("search query is empty")
			return echo.NewHTTPError(http.StatusBadRequest, "search query is empty")
		}
		if q.MinPrice < 0 || q.MaxPrice < 0 || q.MinPrice > q.MaxPrice {
			logger.Errorw("invalid price range")
			return echo.NewHTTPError(http.StatusBadRequest, "invalid price range")
		}
		if q.SortBy != "" && q.SortBy != "price" && q.SortBy != "stock" && q.SortBy != "sales" && q.SortBy != "relevancy" {
			logger.Errorw("invalid sort by")
			return echo.NewHTTPError(http.StatusBadRequest, "invalid sort by")
		}
		if q.Order != "" && q.Order != "asc" && q.Order != "desc" {
			logger.Errorw("invalid order")
			return echo.NewHTTPError(http.StatusBadRequest, "invalid order")
		}
		dbq := db.SearchProductsParams{
			Offset: int32(q.Offset),
			Limit:  int32(q.Limit),
			Query:  q.Q,
			MinPrice: pgtype.Numeric{
				Int:   big.NewInt(int64(q.MinPrice)),
				Valid: true,
			},
			MaxPrice: pgtype.Numeric{
				Int:   big.NewInt(int64(q.MaxPrice)),
				Valid: true,
			},
			HasCoupon: pgtype.Bool{
				Bool:  q.HaveCoupon,
				Valid: true,
			},
			SortBy: q.SortBy,
			Order:  q.Order,
		}
		if c.QueryParam("minPrice") == "" {
			dbq.MinPrice.Valid = false
		}
		if c.QueryParam("maxPrice") == "" {
			dbq.MaxPrice.Valid = false
		}
		if c.QueryParam("minStock") == "" {
			dbq.MinStock.Valid = false
		}
		if c.QueryParam("maxStock") == "" {
			dbq.MaxStock.Valid = false
		}
		if c.QueryParam("haveCoupon") == "" {
			dbq.HasCoupon.Valid = false
		}
		var sr searchResult
		var err error
		sr.Products, err = pg.Queries.SearchProducts(c.Request().Context(), dbq)
		if err != nil {
			logger.Errorw("failed to search products", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range sr.Products {
			sr.Products[i].ImageUrl = mc.GetFileURL(c.Request().Context(), sr.Products[i].ImageUrl)
		}
		sr.Shops, err = pg.Queries.SearchShops(c.Request().Context(), db.SearchShopsParams{
			Query:  q.Q,
			Offset: 0,
			Limit:  4,
		})
		if err != nil {
			logger.Errorw("failed to search shops", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range sr.Shops {
			sr.Shops[i].ImageUrl = mc.GetFileURL(c.Request().Context(), sr.Shops[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, sr)
	}
}

type searchShopByNameParams struct {
	Query string `query:"q"`
	common.QueryParams
}

// @Summary		Search for Shops by Name
// @Description	Search for shops by name
// @Tags			Search,Shop
// @Accept			json
// @Produce		json
// @Param			q		query		string	true	"Search Name"
// @Param			offset	query		int		false	"Begin index"	default(0)
// @Param			limit	query		int		false	"limit"			default(10)	Maximum(20)
// @Success		200		{array}		PrettierShopSearchResult
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/search/shop [get]
func searchShopByName(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := searchShopByNameParams{QueryParams: common.NewQueryParams(0, 10)}
		if err := c.Bind(&q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := q.Validate(); err != nil {
			logger.Errorw("failed to validate query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "query parameter is invalid")
		}
		if q.Query == "" {
			logger.Errorw("search query is empty")
			return echo.NewHTTPError(http.StatusBadRequest, "search query is empty")
		}
		result, err := pg.Queries.SearchShops(c.Request().Context(), db.SearchShopsParams{
			Query:  q.Query,
			Offset: int32(q.Offset),
			Limit:  int32(q.Limit),
		})
		if err != nil {
			logger.Errorw("failed to search shops", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}

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
func getNews(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
func getNewsDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
func getDiscover(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
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
			result[i].ImageUrl = mc.GetFileURL(c.Request().Context(), result[i].ImageUrl)
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
func getPopular(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result popular
		var err error
		result.PopularProducts, err = pg.Queries.GetProductsFromPopularShop(c.Request().Context())
		if err != nil {
			logger.Errorw("failed to get popular products", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range result.PopularProducts {
			result.PopularProducts[i].ImageUrl = mc.GetFileURL(c.Request().Context(), result.PopularProducts[i].ImageUrl)
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
func getProductInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
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
		result.ImageUrl = mc.GetFileURL(c.Request().Context(), result.ImageUrl)
		return c.JSON(http.StatusOK, result)
	}
}
