package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func sellerGetShopInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerEditInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerGetTag(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerAddTag(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerGetShopCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerAddCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerEditCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerDeleteCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerGetOrder(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerGetOrderDetail(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerGetReport(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerGetReportDetail(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerAddProduct(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerUploadProductImage(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerEditProduct(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func sellerDeleteProduct(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// api := e.Group("/api")
// api.Use(middleware.JWT([]byte("secret")))
// // seller
// api.GET("/seller", sellerGetShopInfo)
// api.PATCH("/seller", sellerEditInfo)
// api.GET("/seller/tag", sellerGetTag)  // search avaliable tag
// api.POST("/seller/tag", sellerAddTag) // add tag for shop

// api.GET("/seller/coupon", sellerGetShopCoupon)
// api.POST("/seller/coupon", sellerAddCoupon)
// api.PATCH("/seller/coupon/:id", sellerEditCoupon)
// api.DELETE("/seller/coupon/:id", sellerDeleteCoupon)

// api.GET("/seller/order", sellerGetOrder)
// api.GET("/seller/order/:id", sellerGetOrderDetail)

// api.GET("/seller/report", sellerGetReport)
// api.GET("/seller/report/:year/:month", sellerGetReportDetail)

// api.POST("/seller/product", sellerAddProduct)
// api.POST("/seller/product/:id/upload", sellerUploadProductImage)
// api.PATCH("/seller/product/:id", sellerEditProduct)
// api.DELETE("/seller/product/:id", sellerDeleteProduct)
