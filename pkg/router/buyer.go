package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func buyerGetOrderHistrory(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func buyerGetOrderDetail(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func buyerGetCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func buyerAddProductToCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func buyerAddCouponToCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func buyerDeleteProductFromCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func buyerDeleteCouponFromCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func buyerGetCheckout(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func buyerCheckout(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// api.GET("/buyer/order", buyerGetOrderHistrory)
// api.GET("/buyer/order/:id", buyerGetOrderDetail)

// api.GET("/buyer/cart", buyerGetCart) // include procuct and coupon
// api.POST("/buyer/cart/product:id", buyerAddProductToCart)
// api.POST("/buyer/cart/coupon:id", buyerAddCouponToCart)
// api.DELETE("/buyer/cart/product:id", buyerDeleteProductFromCart)
// api.DELETE("/buyer/cart/coupon:id", buyerDeleteCouponFromCart)

// api.GET("/buyer/checkout", buyerGetCheckout)
// api.POST("/buyer/checkout", buyerCheckout)
