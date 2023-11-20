package router

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary Buyer Get Order History
// @Description Get all order history of the user
// @Tags Buyer, Order
// @Produce json
// @Success 200
// @Failure 401
// @Router /buyer/order [get]
func buyerGetOrderHistory(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Buyer Get Order Detail
// @Description Get specific order detail
// @Tags Buyer, Order
// @Produce json
// @Param id path int true "Order ID"
// @Success 200
// @Failure 401
// @Router /buyer/order/{id} [get]
func buyerGetOrderDetail(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Buyer Get Cart
// @Description Get all Carts of the user
// @Tags Buyer, Cart
// @Produce json
// @Success 200
// @Failure 401
// @Router /buyer/cart [get]
func buyerGetCart(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Buyer Edit Product In Cart
// @Description Edit product quantity in cart
// @Tags Buyer, Cart
// @Accept json
// @Produce json
// @Param cart_id path int true "Cart ID"
// @Param product_id path int true "Product ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/{cart_id}/product/{product_id} [patch]
func buyerEditProductInCart(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Buyer Add Product To Cart
// @Description Add product to cart
// @Tags Buyer, Cart
// @Accept json
// @Produce json
// @Param cart_id path int true "Cart ID"
// @Param product_id path int true "Product ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/{cart_id}/product/{product_id} [post]
func buyerAddProductToCart(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Buyer Add Coupon To Cart
// @Description Add coupon to cart
// @Tags Buyer, Cart, Coupon
// @Accept json
// @Produce json
// @Param cart_id path int true "Cart ID"
// @Param coupon_id path int true "Coupon ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/{cart_id}/coupon/{coupon_id} [post]
func buyerAddCouponToCart(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Buyer Delete Product From Cart
// @Description Delete product from cart
// @Tags Buyer, Cart
// @Accept json
// @Produce json
// @Param cart_id path int true "Cart ID"
// @Param product_id path int true "Product ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/{cart_id}/product/{product_id} [delete]
func buyerDeleteProductFromCart(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Buyer Delete Coupon From Cart
// @Description Delete coupon from cart
// @Tags Buyer, Cart, Coupon
// @Accept json
// @Produce json
// @Param cart_id path int true "Cart ID"
// @Param coupon_id path int true "Coupon ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/{cart_id}/coupon/{coupon_id} [delete]
func buyerDeleteCouponFromCart(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Buyer Get Checkout
// @Description Get all checkout data
// @Tags Buyer, Checkout
// @Produce json
// @Param cart_id path int true "Cart ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/{cart_id}/checkout [get]
func buyerGetCheckout(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Buyer Checkout
// @Description Checkout
// @Tags Buyer, Checkout
// @Accept json
// @Produce json
// @param cart_id path int true "Cart ID"
// @Success 200
// @Failure 401
// @Router /buyer/cart/{cart_id}/checkout [post]
func buyerCheckout(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}
