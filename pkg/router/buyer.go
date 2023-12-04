package router

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
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
		q := common.NewQueryParams(0, 10)
		if err := c.Bind(q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := q.Validate(); err != nil {
			logger.Errorw("invalid query parameter", "offset", q.Offset, "limit", q.Limit)
			return echo.NewHTTPError(http.StatusBadRequest, "invalid query parameter")
		}
		orders, err := pg.Queries.GetOrderHistory(c.Request().Context(), db.GetOrderHistoryParams{Username: username, Offset: q.Offset, Limit: q.Limit})
		if err != nil {
			logger.Errorw("failed to get order history", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, orders)
	}
}

type OrderDetail struct {
	Info    db.GetOrderInfoRow     `json:"info"`
	Details []db.GetOrderDetailRow `json:"details"`
}

// @Summary		Buyer Get Order Detail
// @Description	Get specific order detail
// @Tags			Buyer, Order
// @Produce		json
// @Param			id	path		int	true	"Order ID"
// @Success		200	{object}	OrderDetail
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/buyer/order/{id} [get]
func buyerGetOrderDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "🫠"
		var orderID int32
		var orderDetail OrderDetail
		var err error
		if echo.PathParamsBinder(c).Int32("id", &orderID).BindError() != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if orderDetail.Info, err = pg.Queries.GetOrderInfo(c.Request().Context(), db.GetOrderInfoParams{Username: username, OrderID: orderID}); err != nil {
			if err == pgx.ErrNoRows {
				return echo.NewHTTPError(http.StatusBadRequest)
			}
			logger.Errorw("failed to get order info", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if orderDetail.Details, err = pg.Queries.GetOrderDetail(c.Request().Context(), orderID); err != nil {
			logger.Errorw("failed to get order detail", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, orderDetail)
	}
}

type Cart struct {
	Seller_name string
	Products    []db.GetProductInCartRow
}

// @Summary		Buyer Get Cart
// @Description	Get all Carts of the user
// @Tags			Buyer, Cart
// @Produce		json
// @Success		200	{array}		Cart
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/buyer/cart [get]
func buyerGetCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "🤡"
		carts, err := pg.Queries.GetCart(c.Request().Context(), username)
		if err != nil {
			logger.Errorw("failed to get cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		var result []Cart
		for _, cartInfo := range carts {
			var cart Cart
			products, err := pg.Queries.GetProductInCart(c.Request().Context(), cartInfo.ID)
			if err != nil {
				logger.Errorw("failed to get product in cart", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			cart.Seller_name = cartInfo.SellerName
			cart.Products = products
			result = append(result, cart)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Buyer Edit Product In Cart
// @Description	Edit product quantity in cart (The product must be in the cart)
// @Tags			Buyer, Cart
// @Accept			json
// @Produce		json
// @Param			cart_id		path	int	true	"Cart ID"
// @Param			product_id	path	int	true	"Product ID"
// @Success		200 {string} string constants.SUCCESS
// @Failure		400 {object} echo.HTTPError
// @Failure		500 {object} echo.HTTPError
// @Router			/buyer/cart/{cart_id}/product/{product_id} [patch]
func buyerEditProductInCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "👷🏿"
		var param db.UpdateProductInCartParams
		param.Username = username
		if err := c.Bind(&param); err != nil {
			logger.Errorw("failed to bind product in cart", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if param.Quantity < 0 {
			logger.Errorw("invalid quantity", "quantity", param.Quantity)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if param.Quantity == 0 {
			if rows, err := pg.Queries.DeleteProductInCart(c.Request().Context(),
				db.DeleteProductInCartParams{
					Username:  username,
					ID:        param.CartID,
					ProductID: param.ProductID,
				}); err != nil {
				logger.Errorw("failed to delete product in cart", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			} else if rows == 0 {
				return echo.NewHTTPError(http.StatusBadRequest)
			}
			return c.JSON(http.StatusOK, constants.SUCCESS)
		}
		if _, err := pg.Queries.UpdateProductInCart(c.Request().Context(), param); err != nil {
			if err == pgx.ErrNoRows {
				return echo.NewHTTPError(http.StatusBadRequest)
			}
			logger.Errorw("failed to update product in cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

// @Summary		Buyer Add Product To Cart
// @Description	Add product to cart
// @Tags			Buyer, Cart
// @Accept			json
// @Produce		json
// @Param			cart_id		path	int	true	"Cart ID"
// @Param			product_id	path	int	true	"Product ID"
// @Success		200 {integer} int "product quantity in cart"
// @Failure		400 {object} echo.HTTPError
// @Failure		500 {object} echo.HTTPError
// @Router			/buyer/cart/product/{id} [post]
func buyerAddProductToCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "🧑🏿‍⚕️"
		var param db.AddProductToCartParams
		param.Username = username
		if err := c.Bind(&param); err != nil {
			logger.Errorw("failed to bind product in cart", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if param.Quantity <= 0 {
			logger.Errorw("invalid quantity", "quantity", param.Quantity)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		cnt, err := pg.Queries.AddProductToCart(c.Request().Context(), param)
		if err != nil {
			logger.Errorw("failed to add product to cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, cnt)
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
