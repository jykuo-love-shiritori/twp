package buyer

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary		Buyer Edit Product In Cart
// @Description	Edit product quantity in cart (The product must be in the cart)
// @Tags			Buyer, Cart
// @Accept			json
// @Produce		json
// @Param			cart_id		path		int				true	"Cart ID"
// @Param			product_id	path		int				true	"Product ID"
// @Param			quantity	body		ProductQuantity	true	"Quantity"
// @Success		200			{string}	string			constants.SUCCESS
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/buyer/cart/{cart_id}/product/{product_id} [patch]
func EditProductInCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var param db.UpdateProductFromCartParams
		param.Username = username
		if err := echo.PathParamsBinder(c).Int32("cart_id", &param.CartID).Int32("product_id", &param.ProductID).BindError(); err != nil {
			logger.Errorw("failed to parse cart_id or product_id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := c.Bind(&param); err != nil {
			logger.Errorw("failed to bind product in cart", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if param.Quantity < 0 {
			logger.Errorw("invalid quantity", "quantity", param.Quantity)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if param.Quantity == 0 {
			tx, err := pg.NewTx(c.Request().Context())
			defer tx.Rollback(c.Request().Context()) //nolint:errcheck
			if err != nil {
				logger.Errorw("failed to create transaction", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			exec, err := pg.Queries.WithTx(tx).DeleteProductFromCart(c.Request().Context(),
				db.DeleteProductFromCartParams{
					Username:  username,
					CartID:    param.CartID,
					ProductID: param.ProductID,
				})
			if err != nil {
				logger.Errorw("failed to delete product in cart", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if !exec {
				return echo.NewHTTPError(http.StatusBadRequest)
			}
			// delete empty cart
			if err := pg.Queries.WithTx(tx).DeleteEmptyCart(c.Request().Context(), db.DeleteEmptyCartParams{
				Username: username,
				CartID:   param.CartID,
			}); err != nil {
				logger.Errorw("failed to delete empty cart", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if err := tx.Commit(c.Request().Context()); err != nil {
				logger.Errorw("failed to commit transaction", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			return c.JSON(http.StatusOK, constants.SUCCESS)
		}
		if rows, err := pg.Queries.UpdateProductFromCart(c.Request().Context(), param); err != nil {
			logger.Errorw("failed to update product in cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if rows == 0 {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

type ProductQuantity struct {
	Quantity int32 `json:"quantity"`
}

// @Summary		Buyer Add Product To Cart
// @Description	Add product to cart
// @Tags			Buyer, Cart
// @Accept			json
// @Produce		json
// @Param			id			path		int				true	"Product ID"
// @Param			quantity	body		ProductQuantity	true	"Quantity"
// @Success		200			{integer}	int				"product quantity in cart"
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/buyer/cart/product/{id} [post]
func AddProductToCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var param db.AddProductToCartParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("failed to bind product in cart", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.Username = username
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

// @Summary		Buyer Delete Product From Cart
// @Description	Delete product from cart
// @Tags			Buyer, Cart
// @Accept			json
// @Produce		json
// @Param			cart_id		path		int		true	"Cart ID"
// @Param			product_id	path		int		true	"Product ID"
// @Success		200			{string}	string	constants.SUCCESS
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/buyer/cart/{cart_id}/product/{product_id} [delete]
func DeleteProductFromCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var param db.DeleteProductFromCartParams
		param.Username = username
		if err := echo.PathParamsBinder(c).Int32("cart_id", &param.CartID).Int32("product_id", &param.ProductID).BindError(); err != nil {
			logger.Errorw("failed to parse cart_id or product_id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if exec, err := pg.Queries.DeleteProductFromCart(c.Request().Context(), param); err != nil {
			logger.Errorw("failed to delete product in cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if !exec {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		// delete empty cart
		if err := pg.Queries.DeleteEmptyCart(c.Request().Context(), db.DeleteEmptyCartParams{
			Username: username,
			CartID:   param.CartID,
		}); err != nil {
			logger.Errorw("failed to delete empty cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}
