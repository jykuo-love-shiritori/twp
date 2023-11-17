package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Get Shop Info
// @Description Get shop information with seller username
// @Tags Shop
// @Accept json
// @Produce json
// @Param seller_name path int true "seller username"
// @Success 200
// @Failure 401
// @Router /shop/{seller_name} [get]
func getShopInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Get Shop Coupons
// @Description Get coupons for a shop with seller username
// @Tags Shop,Coupon
// @Accept json
// @Produce json
// @Param seller_name path int true "seller username"
// @Success 200
// @Failure 401
// @Router /shop/{seller_name}/coupon [get]
func getShopCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Search Shop Products
// @Description Search products within a shop by seller username
// @Tags Shop,Product,Search
// @Accept json
// @Produce json
// @Param seller_name path int true "Seller username"
// @Param q query string true "search word"
// @Success 200
// @Failure 401
// @Router /shop/{seller_name}/search [get]
func searchShopProduct(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Get Tag Info
// @Description Get information about a tag by tag ID
// @Tags Tag
// @Accept json
// @Produce json
// @Param id path int true "Tag ID"
// @Success 200
// @Failure 401
// @Router /tag/{id} [get]
func getTagInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Search for Products and Shops
// @Description Search for products and shops
// @Tags Search
// @Accept json
// @Produce json
// @Param q query string true "search word"
// @Success 200
// @Failure 401
// @Router /search [get]
func search(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Search for Shops by Name
// @Description Search for shops by name
// @Tags Search,Shop
// @Accept json
// @Produce json
// @Param q query string true "Search Name"
// @Success 200
// @Failure 401
// @Router /search/shop [get]
func searchShopByName(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Get News
// @Description Get news
// @Tags News
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /news [get]
func getNews(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Get News Detail
// @Description Get details of a specific news item by ID
// @Tags News
// @Accept json
// @Produce json
// @Param id path int true "News ID"
// @Success 200
// @Failure 401
// @Router /news/{id} [get]
func getNewsDetail(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Get Discover
// @Description Get discover content
// @Tags Discover,Product
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /discover [get]
func getDiscover(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Get Product Info
// @Description Get product information with product ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200
// @Failure 401
// @Router /product/{id} [get]
func getProductInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
