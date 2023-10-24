package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Buyer Get Order History
// @Description Get all order history of the user
// @Tags Buyer, Order
// @Produce json
// @Success 200
// @Failure 401
// @Router /buyer/order [get]
func buyerGetOrderHistrory(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Buyer Get Order Detail
// @Description Get specific order detail
// @Tags Buyer, Order
// @Produce json
// @Param id path int true "Order ID"
// @Success 200
// @Failure 401
// @Router /buyer/order/{id} [get]
func buyerGetOrderDetail(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Buyer Get Cart
// @Description Get all products and coupons in cart
// @Tags Buyer, Cart
// @Produce json
// @Success 200
// @Failure 401
// @Router /buyer/cart [get]
func buyerGetCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Buyer Edit Product In Cart
// @Description Edit product quantity in cart
// @Tags Buyer, Cart
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/product/{id} [patch]
func buyerEditProductInCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Buyer Add Product To Cart
// @Description Add product to cart
// @Tags Buyer, Cart
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/product/{id} [post]
func buyerAddProductToCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Buyer Add Coupon To Cart
// @Description Add coupon to cart
// @Tags Buyer, Cart, Coupon
// @Accept json
// @Produce json
// @Param id path int true "Coupon ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/coupon/{id} [post]
func buyerAddCouponToCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Buyer Delete Product From Cart
// @Description Delete product from cart
// @Tags Buyer, Cart
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/product/{id} [delete]
func buyerDeleteProductFromCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Buyer Delete Coupon From Cart
// @Description Delete coupon from cart
// @Tags Buyer, Cart, Coupon
// @Accept json
// @Produce json
// @Param id path int true "Coupon ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/coupon/{id} [delete]
func buyerDeleteCouponFromCart(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Buyer Get Checkout
// @Description Get all checkout data
// @Tags Buyer, Checkout
// @Produce json
// @Success 200
// @Failure 401
// @Router /buyer/checkout [get]
func buyerGetCheckout(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Buyer Checkout
// @Description Checkout
// @Tags Buyer, Checkout
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /buyer/checkout [post]
func buyerCheckout(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
