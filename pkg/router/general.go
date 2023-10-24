package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getShopInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func getShopCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func searchShopProduct(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func getTagInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func search(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func searchShopByName(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func getNews(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func getNewsDetail(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func getDiscover(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func getProductInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// api.GET("/shop/:id", getShopInfo) // user
// api.GET("/shop/:id/coupon", getShopCoupon)
// api.GET("/shop/:id/search", searchShop)

// api.GET("/tag/:id", getTagInfo)

// api.GET("/search", search)
// api.GET("/search/shop", searchShop)

// api.GET("/news", getNews)
// api.GET("/news/:id", getNewsDetail)
// api.GET("/discover", getDiscover)

// api.GET("/product/:id", getProductInfo)
