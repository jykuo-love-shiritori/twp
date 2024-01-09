package general

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/image"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type searchParams struct {
	Q          string  `query:"q"`
	MinStock   int32   `query:"minStock"`
	MaxStock   int32   `query:"maxStock"`
	MinPrice   float64 `query:"minPrice"`
	MaxPrice   float64 `query:"maxPrice"`
	HaveCoupon bool    `query:"haveCoupon"`
	SortBy     string  `query:"sortBy"`
	Order      string  `query:"order"`
	common.QueryParams
}

func (q *searchParams) Validate() error {
	if q.Q == "" {
		return errors.New("search query is empty")
	}
	if q.MinPrice < 0 || q.MaxPrice < 0 || q.MinPrice > q.MaxPrice {
		return errors.New("invalid price range")
	}
	if q.SortBy != "" && q.SortBy != "price" && q.SortBy != "stock" && q.SortBy != "sales" && q.SortBy != "relevancy" {
		return errors.New("invalid sort by")
	}
	if q.Order != "" && q.Order != "asc" && q.Order != "desc" {
		return errors.New("invalid order")
	}
	return q.QueryParams.Validate()
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
	Shops    []PrettierShopSearchResult    `json:"shops"`
}
type searchResult struct {
	Products []db.SearchProductsRow `json:"products"`
	Shops    []db.SearchShopsRow    `json:"shops"`
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
// @Param			sortBy		query		string	false	"sort by"		Enums(price, stock, sales, relevancy)
// @Param			order		query		string	false	"sorting order"	Enums(asc, desc)
// @Param			offset		query		int		false	"Begin index"	default(0)
// @Param			limit		query		int		false	"limit"			default(10)	Maximum(20)
// @Success		200			{object}	PrettierProductSearchResult
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/shop/{seller_name}/search [get]
func SearchShopProduct(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
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
		if err := q.Validate(); err != nil {
			logger.Errorw("failed to validate query parameter", "error", err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		dbq := db.SearchProductsByShopParams{
			SellerName: seller_name,
			Offset:     int32(q.Offset),
			Limit:      int32(q.Limit),
			Query:      q.Q,
			HasCoupon: pgtype.Bool{
				Bool:  q.HaveCoupon,
				Valid: true,
			},
			SortBy: q.SortBy,
			Order:  q.Order,
		}
		if err := dbq.MinPrice.Scan(fmt.Sprintf("%f", q.MinPrice)); err != nil {
			logger.Errorw("failed to scan minPrice", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := dbq.MaxPrice.Scan(fmt.Sprintf("%f", q.MaxPrice)); err != nil {
			logger.Errorw("failed to scan maxPrice", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := dbq.MinStock.Scan(fmt.Sprintf("%d", q.MinStock)); err != nil {
			logger.Errorw("failed to scan minStock", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := dbq.MaxStock.Scan(fmt.Sprintf("%d", q.MaxStock)); err != nil {
			logger.Errorw("failed to scan maxStock", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		paramMap := map[string]interface{}{
			"minPrice":   &dbq.MinPrice,
			"maxPrice":   &dbq.MaxPrice,
			"minStock":   &dbq.MinStock,
			"maxStock":   &dbq.MaxStock,
			"haveCoupon": &dbq.HasCoupon,
		}
		for k, v := range paramMap {
			if c.QueryParam(k) == "" {
				switch val := v.(type) {
				case *pgtype.Numeric:
					val.Valid = false
				case *pgtype.Bool:
					val.Valid = false
				case *pgtype.Int4:
					val.Valid = false
				}
			}
		}
		products, err := pg.Queries.SearchProductsByShop(c.Request().Context(), dbq)
		if err != nil {
			logger.Errorw("failed to search products", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range products {
			products[i].ImageUrl = image.GetUrl(products[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, products)
	}
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
// @Param			sortBy		query		string	false	"sort by"		Enums(price, stock, sales, relevancy)
// @Param			order		query		string	false	"sorting order"	Enums(asc, desc)
// @Param			offset		query		int		false	"Begin index"	default(0)
// @Param			limit		query		int		false	"limit"			default(10)	Maximum(20)
// @Success		200			{object}	PrettierSearchResult
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/search [get]
func Search(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := newSearchParams()
		if err := c.Bind(&q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := q.Validate(); err != nil {
			logger.Errorw("failed to validate query parameter", "error", err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		dbq := db.SearchProductsParams{
			Offset: int32(q.Offset),
			Limit:  int32(q.Limit),
			Query:  q.Q,
			HasCoupon: pgtype.Bool{
				Bool:  q.HaveCoupon,
				Valid: true,
			},
			SortBy: q.SortBy,
			Order:  q.Order,
		}
		if err := dbq.MinPrice.Scan(fmt.Sprintf("%f", q.MinPrice)); err != nil {
			logger.Errorw("failed to scan minPrice", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := dbq.MaxPrice.Scan(fmt.Sprintf("%f", q.MaxPrice)); err != nil {
			logger.Errorw("failed to scan maxPrice", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := dbq.MinStock.Scan(fmt.Sprintf("%d", q.MinStock)); err != nil {
			logger.Errorw("failed to scan minStock", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := dbq.MaxStock.Scan(fmt.Sprintf("%d", q.MaxStock)); err != nil {
			logger.Errorw("failed to scan maxStock", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		paramMap := map[string]interface{}{
			"minPrice":   &dbq.MinPrice,
			"maxPrice":   &dbq.MaxPrice,
			"minStock":   &dbq.MinStock,
			"maxStock":   &dbq.MaxStock,
			"haveCoupon": &dbq.HasCoupon,
		}
		for k, v := range paramMap {
			if c.QueryParam(k) == "" {
				switch val := v.(type) {
				case *pgtype.Numeric:
					val.Valid = false
				case *pgtype.Bool:
					val.Valid = false
				case *pgtype.Int4:
					val.Valid = false
				}
			}
		}
		var sr searchResult
		var err error
		sr.Products, err = pg.Queries.SearchProducts(c.Request().Context(), dbq)
		if err != nil {
			logger.Errorw("failed to search products", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range sr.Products {
			sr.Products[i].ImageUrl = image.GetUrl(sr.Products[i].ImageUrl)
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
			sr.Shops[i].ImageUrl = image.GetUrl(sr.Shops[i].ImageUrl)
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
func SearchShopByName(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
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
