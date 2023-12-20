package buyer

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Cart struct {
	CartInfo db.GetCartRow
	Products []db.GetProductFromCartOrderByPriceDescRow
	Coupons  []db.GetCouponsFromCartRow
}

// @Summary		Buyer Get Cart
// @Description	Get all Carts of the user
// @Tags			Buyer, Cart
// @Produce		json
// @Success		200	{array}		Cart
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/buyer/cart [get]
func GetCart(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		carts, err := pg.Queries.GetCart(c.Request().Context(), username)
		if err != nil {
			logger.Errorw("failed to get cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result := []Cart{}
		for _, cartInfo := range carts {
			var cart Cart
			var err error
			cart.Products, err = pg.Queries.GetProductFromCartOrderByPriceDesc(c.Request().Context(), cartInfo.ID)
			if err != nil {
				logger.Errorw("failed to get product in cart", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			for i := range cart.Products {
				cart.Products[i].ImageUrl = mc.GetFileURL(c.Request().Context(), cart.Products[i].ImageUrl)
			}
			cart.Coupons, err = pg.Queries.GetCouponsFromCart(c.Request().Context(), db.GetCouponsFromCartParams{Username: username, CartID: cartInfo.ID})
			if err != nil {
				logger.Errorw("failed to get coupon in cart", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			cartInfo.ShopImageUrl = mc.GetFileURL(c.Request().Context(), cartInfo.ShopImageUrl)
			cart.CartInfo = cartInfo
			result = append(result, cart)
		}
		return c.JSON(http.StatusOK, result)
	}
}
