package buyer

import (
	"net/http"
	"sort"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// numericToFloat64 converts a pgtype.Numeric to a float64.

func numericToFloat64(n pgtype.Numeric) (float64, error) {
	f, err := n.Float64Value()
	return f.Float64, err
}

type couponV2 struct {
	db.GetCouponsFromCartRow
	DiscountFloat float64
}

// @Summary		Buyer Get usable coupon of cart/shop
// @Description	Buyer get usable coupon of cart/shop
// @Tags			Buyer, Cart, Coupon
// @Accept			json
// @Produce		json
// @Param			cart_id	path		int	true	"Cart ID"
// @Success		200		{array}		db.GetUsableCouponsRow
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/buyer/cart/{cart_id}/coupon [get]
func GetCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var param db.GetUsableCouponsParams
		param.Username = username
		result := []db.GetUsableCouponsRow{}
		if err := c.Bind(&param); err != nil {
			logger.Errorw("failed to bind coupon in cart", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		usableCoupon, err := pg.Queries.GetUsableCoupons(c.Request().Context(), param)
		if err != nil {
			logger.Errorw("failed to get coupon", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		cartCoupon, err := pg.Queries.GetCouponsFromCart(c.Request().Context(), db.GetCouponsFromCartParams{Username: username, CartID: param.CartID})
		if err != nil {
			logger.Errorw("failed to get coupon", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		cartProduct, err := pg.Queries.GetProductFromCart(c.Request().Context(), param.CartID)
		if err != nil {
			logger.Errorw("failed to get product in cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		productCount := make(map[int32]int32)
		productTag := make(map[int32][]db.GetProductTagRow)
		for _, product := range cartProduct {
			productCount[product.ProductID] = product.Quantity
			productTag[product.ProductID], err = pg.Queries.GetProductTag(c.Request().Context(), product.ProductID)
			if err != nil {
				logger.Errorw("failed to get product tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}
		var cartCouponV2 []couponV2
		for _, coupon := range cartCoupon {
			discount, err := numericToFloat64(coupon.Discount)
			if err != nil {
				logger.Errorw("failed to get discount", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			cartCouponV2 = append(cartCouponV2, couponV2{GetCouponsFromCartRow: coupon, DiscountFloat: discount})
		}
		shippingFlag := false
		// sort CartCoupon by type (percentage => fixed), discount (high => low)
		sort.Slice(cartCouponV2, func(i, j int) bool {
			if cartCouponV2[i].Type == cartCouponV2[j].Type {
				return cartCouponV2[i].DiscountFloat > cartCouponV2[j].DiscountFloat
			}
			return cartCouponV2[i].Type < cartCouponV2[j].Type
		})
		// all the coupon in cart should be eligible
		for _, coupon := range cartCouponV2 {
			couponTag, err := pg.Queries.GetCouponTag(c.Request().Context(), coupon.ID)
			if err != nil {
				logger.Errorw("failed to get coupon tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if coupon.Type == db.CouponTypeShipping {
				shippingFlag = true
				continue
			}
		outerLoop:
			for _, product := range cartProduct {
				if productCount[product.ProductID] == 0 {
					continue
				}
				for _, ct := range couponTag {
					for _, pt := range productTag[product.ProductID] {
						if ct.TagID == pt.TagID {
							productCount[product.ProductID] -= 1
							break outerLoop
						}
					}
				}
			}
		}

		for _, usable := range usableCoupon {
			couponTag, err := pg.Queries.GetCouponTag(c.Request().Context(), usable.ID)
			if err != nil {
				logger.Errorw("failed to get coupon tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if usable.Type == db.CouponTypeShipping && shippingFlag {
				continue
			}

			// Check if the coupon is valid for any of the products in the cart
			if usable.Scope == db.CouponScopeGlobal {
				result = append(result, usable)
				continue
			}
			matchFlag := false
		outerLoop2:
			for _, product := range cartProduct {
				if productCount[product.ProductID] == 0 {
					continue
				}
				for _, ct := range couponTag {
					for _, pt := range productTag[product.ProductID] {
						if ct.TagID == pt.TagID {
							matchFlag = true
							break outerLoop2
						}
					}
				}
			}
			if matchFlag {
				result = append(result, usable)
			}
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Buyer Add Coupon To Cart
// @Description	Add coupon to cart
// @Tags			Buyer, Cart, Coupon
// @Accept			json
// @Produce		json
// @Param			cart_id		path		int		true	"Cart ID"
// @Param			coupon_id	path		int		true	"Coupon ID"
// @Success		200			{string}	string	constants.SUCCESS
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/buyer/cart/{cart_id}/coupon/{coupon_id} [post]
func AddCouponToCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var param db.AddCouponToCartParams
		param.Username = username
		if err := echo.PathParamsBinder(c).Int32("cart_id", &param.CartID).Int32("coupon_id", &param.CouponID).BindError(); err != nil {
			logger.Errorw("failed to parse cart_id or coupon_id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		cartProduct, err := pg.Queries.GetProductFromCart(c.Request().Context(), param.CartID)
		if err != nil {
			logger.Errorw("failed to get product in cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		productCount := make(map[int32]int32)
		productTag := make(map[int32][]db.GetProductTagRow)
		for _, product := range cartProduct {
			productCount[product.ProductID] = product.Quantity
			productTag[product.ProductID], err = pg.Queries.GetProductTag(c.Request().Context(), product.ProductID)
			if err != nil {
				logger.Errorw("failed to get product tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}
		cartCoupon, err := pg.Queries.GetCouponsFromCart(c.Request().Context(), db.GetCouponsFromCartParams{Username: username, CartID: param.CartID})
		if err != nil {
			logger.Errorw("failed to get coupon", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		var cartCouponV2 []couponV2
		for _, coupon := range cartCoupon {
			discount, err := numericToFloat64(coupon.Discount)
			if err != nil {
				logger.Errorw("failed to get discount", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			cartCouponV2 = append(cartCouponV2, couponV2{GetCouponsFromCartRow: coupon, DiscountFloat: discount})
		}
		shippingFlag := false
		// sort CartCoupon by type (percentage => fixed), discount (high => low)
		sort.Slice(cartCouponV2, func(i, j int) bool {
			if cartCouponV2[i].Type == cartCouponV2[j].Type {
				return cartCouponV2[i].DiscountFloat > cartCouponV2[j].DiscountFloat
			}
			return cartCouponV2[i].Type < cartCouponV2[j].Type
		})
		// all the coupon in cart should be eligible
		// match coupons to its product
		for _, coupon := range cartCouponV2 {
			couponTag, err := pg.Queries.GetCouponTag(c.Request().Context(), coupon.ID)
			if err != nil {
				logger.Errorw("failed to get coupon tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if coupon.Type == db.CouponTypeShipping {
				shippingFlag = true
				continue
			}
		outerLoop:
			for _, product := range cartProduct {
				if productCount[product.ProductID] == 0 {
					continue
				}
				for _, ct := range couponTag {
					for _, pt := range productTag[product.ProductID] {
						if ct.TagID == pt.TagID {
							productCount[product.ProductID] -= 1
							break outerLoop
						}
					}
				}
			}
		}
		// check new coupon eligibility
		couponDetail, err := pg.Queries.GetCouponDetail(c.Request().Context(), param.CouponID)
		if err != nil {
			logger.Errorw("failed to get coupon detail", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if couponDetail.Type == db.CouponTypeShipping && shippingFlag {
			return echo.NewHTTPError(http.StatusBadRequest, "multiple shipping coupon")
		}
		couponTag, err := pg.Queries.GetCouponTag(c.Request().Context(), param.CouponID)
		if err != nil {
			logger.Errorw("failed to get coupon tag", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		couponValid := couponDetail.Scope == db.CouponScopeGlobal

	outerLoop2:
		for _, product := range cartProduct {
			if productCount[product.ProductID] == 0 {
				continue
			}
			for _, ct := range couponTag {
				for _, pt := range productTag[product.ProductID] {
					if ct.TagID == pt.TagID {
						couponValid = true
						break outerLoop2
					}
				}
			}
		}
		if !couponValid {
			return echo.NewHTTPError(http.StatusBadRequest, "coupon not valid for any product in cart")
		}
		rows, err := pg.Queries.AddCouponToCart(c.Request().Context(), param)
		if err != nil {
			logger.Errorw("failed to add coupon to cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if rows == 0 {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

// @Summary		Buyer Delete Coupon In Cart
// @Description	Delete coupon In cart
// @Tags			Buyer, Cart, Coupon
// @Accept			json
// @Produce		json
// @Param			cart_id		path		int		true	"Cart ID"
// @Param			coupon_id	path		int		true	"Coupon ID"
// @Success		200			{string}	string	constants.SUCCESS
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/buyer/cart/{cart_id}/coupon/{coupon_id} [delete]
func DeleteCouponFromCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var param db.DeleteCouponFromCartParams
		param.Username = username
		if err := echo.PathParamsBinder(c).Int32("cart_id", &param.CartID).Int32("coupon_id", &param.CouponID).BindError(); err != nil {
			logger.Errorw("failed to parse cart_id or coupon_id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if rows, err := pg.Queries.DeleteCouponFromCart(c.Request().Context(), param); err != nil {
			logger.Errorw("failed to delete coupon in cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if rows == 0 {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}
