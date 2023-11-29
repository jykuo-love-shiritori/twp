package router

import (
	"net/http"

	"github.com/jackc/pgx/v5"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary		Buyer Get Order History
// @Description	Get all order history of the user
// @Tags			Buyer, Order
// @Produce		json
// @Param			offset	query		int	false	"Begin index"	default(0)
// @Param			limit	query		int	false	"limit"			default(10)
// @Success		200		{array}		db.GetOrderHistoryRow
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/buyer/order [get]
func buyerGetOrderHistory(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "🤡"
		userId, err := pg.Queries.GetUserIDByUsername(c.Request().Context(), username)
		if err != nil {
			if err == pgx.ErrNoRows {
				return echo.NewHTTPError(http.StatusNotFound, "User Not Found/Disabled")
			}
			logger.Errorw("failed to get user id", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		q := common.NewQueryParams(0, 10)
		if err := c.Bind(q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		orders, err := pg.Queries.GetOrderHistory(c.Request().Context(), userId)
		if err != nil {
			logger.Errorw("failed to get order history", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if int(q.Offset) > len(orders) {
			return echo.NewHTTPError(http.StatusBadRequest, "Offset out of range")
		}
		q.Limit = min(q.Limit, int32(len(orders))-q.Offset)
		return c.JSON(http.StatusOK, orders[q.Offset:q.Offset+q.Limit])
	}
}

// @Summary		Buyer Get Order Detail
// @Description	Get specific order detail
// @Tags			Buyer, Order
// @Produce		json
// @Param			id	path	int	true	"Order ID"
// @Success		200
// @Failure		401
// @Router			/buyer/order/{id} [get]
func buyerGetOrderDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Buyer Get Cart
// @Description	Get all Carts of the user
// @Tags			Buyer, Cart
// @Produce		json
// @Success		200	{array}		common.Cart
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/buyer/cart [get]
func buyerGetCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "🤡"
		userId, err := pg.Queries.GetUserIDByUsername(c.Request().Context(), username)
		if err != nil {
			logger.Errorw("failed to get user id", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		carts, err := pg.Queries.GetCart(c.Request().Context(), userId)
		if err != nil {
			logger.Errorw("failed to get cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		var result []common.Cart
		for _, cartInfo := range carts {
			var cart common.Cart
			products, err := pg.Queries.GetProductInCart(c.Request().Context(), cartInfo.ID)
			if err != nil {
				logger.Errorw("failed to get product in cart", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			seller_name, err := pg.Queries.GetSellerNameByShopID(c.Request().Context(), cartInfo.ShopID)
			if err != nil {
				logger.Errorw("failed to get seller name", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			cart.Seller_name = seller_name
			cart.Products = products
			result = append(result, cart)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Buyer Edit Product In Cart
// @Description	Edit product quantity in cart
// @Tags			Buyer, Cart
// @Accept			json
// @Produce		json
// @Param			cart_id		path	int	true	"Cart ID"
// @Param			product_id	path	int	true	"Product ID"
// @Success		200
// @Failure		401
// @Router			/buyer/cart/{cart_id}/product/{product_id} [patch]
func buyerEditProductInCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Buyer Add Product To Cart
// @Description	Add product to cart
// @Tags			Buyer, Cart
// @Accept			json
// @Produce		json
// @Param			cart_id		path	int	true	"Cart ID"
// @Param			product_id	path	int	true	"Product ID"
// @Success		200
// @Failure		401
// @Router			/buyer/cart/{cart_id}/product/{product_id} [post]
func buyerAddProductToCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Buyer Add Coupon To Cart
// @Description	Add coupon to cart
// @Tags			Buyer, Cart, Coupon
// @Accept			json
// @Produce		json
// @Param			cart_id		path	int	true	"Cart ID"
// @Param			coupon_id	path	int	true	"Coupon ID"
// @Success		200
// @Failure		401
// @Router			/buyer/cart/{cart_id}/coupon/{coupon_id} [post]
func buyerAddCouponToCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Buyer Delete Product From Cart
// @Description	Delete product from cart
// @Tags			Buyer, Cart
// @Accept			json
// @Produce		json
// @Param			cart_id		path	int	true	"Cart ID"
// @Param			product_id	path	int	true	"Product ID"
// @Success		200
// @Failure		401
// @Router			/buyer/cart/{cart_id}/product/{product_id} [delete]
func buyerDeleteProductFromCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Buyer Delete Coupon From Cart
// @Description	Delete coupon from cart
// @Tags			Buyer, Cart, Coupon
// @Accept			json
// @Produce		json
// @Param			cart_id		path	int	true	"Cart ID"
// @Param			coupon_id	path	int	true	"Coupon ID"
// @Success		200
// @Failure		401
// @Router			/buyer/cart/{cart_id}/coupon/{coupon_id} [delete]
func buyerDeleteCouponFromCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Buyer Get Checkout
// @Description	Get all checkout data
// @Tags			Buyer, Checkout
// @Produce		json
// @Param			cart_id	path	int	true	"Cart ID"
// @Success		200
// @Failure		401
// @Router			/buyer/cart/{cart_id}/checkout [get]
func buyerGetCheckout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Buyer Checkout
// @Description	Checkout
// @Tags			Buyer, Checkout
// @Accept			json
// @Produce		json
// @param			cart_id	path	int	true	"Cart ID"
// @Success		200
// @Failure		401
// @Router			/buyer/cart/{cart_id}/checkout [post]
func buyerCheckout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}
