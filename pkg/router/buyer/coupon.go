package buyer

import (
	"net/http"
	"sort"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type tagSet struct {
	Set map[int32]bool
}

func NewTagSet(tags []int32) *tagSet {
	t := &tagSet{}
	t.Set = make(map[int32]bool)
	for _, tag := range tags {
		t.Set[tag] = true
	}
	return t
}
func (t *tagSet) Intersect(other *tagSet) bool {
	for k := range t.Set {
		if other.Set[k] {
			return true
		}
	}
	return false
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
		productTag := make(map[int32]*tagSet)
		for _, product := range cartProduct {
			productCount[product.ProductID] = product.Quantity
			tags, err := pg.Queries.GetProductTag(c.Request().Context(), product.ProductID)
			if err != nil {
				logger.Errorw("failed to get product tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			productTag[product.ProductID] = NewTagSet(tags)
		}
		shippingFlag := false
		// sort CartCoupon by type (percentage => fixed), discount (high => low)
		sort.Slice(cartCoupon, func(i, j int) bool {
			if cartCoupon[i].Type == cartCoupon[j].Type {
				return cartCoupon[i].Discount.Int.Cmp(cartCoupon[j].Discount.Int) == 1
			}
			return cartCoupon[i].Type < cartCoupon[j].Type
		})
		// match coupon to cart's product, all the coupon in cart should be eligible
		for _, coupon := range cartCoupon {
			tags, err := pg.Queries.GetCouponTag(c.Request().Context(), coupon.ID)
			if err != nil {
				logger.Errorw("failed to get coupon tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if coupon.Type == db.CouponTypeShipping {
				shippingFlag = true
				continue
			}
			couponTags := NewTagSet(tags)
			for _, product := range cartProduct {
				if productCount[product.ProductID] == 0 {
					continue
				}
				if productTag[product.ProductID].Intersect(couponTags) {
					productCount[product.ProductID] -= 1
					break
				}
			}
		}
		// check usable coupon eligibility
		for _, usable := range usableCoupon {
			tags, err := pg.Queries.GetCouponTag(c.Request().Context(), usable.ID)
			if err != nil {
				logger.Errorw("failed to get coupon tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			// one shipping coupon per cart
			if usable.Type == db.CouponTypeShipping && shippingFlag {
				continue
			}
			couponTags := NewTagSet(tags)
			// global coupon is always usable beside duplicate shipping coupon
			if usable.Scope == db.CouponScopeGlobal {
				result = append(result, usable)
				continue
			}
			// match coupon to cart's product
			for _, product := range cartProduct {
				if productCount[product.ProductID] == 0 {
					continue
				}
				if productTag[product.ProductID].Intersect(couponTags) {
					result = append(result, usable)
					break
				}
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
		productTag := make(map[int32]*tagSet)
		// init productCount and productTag
		for _, product := range cartProduct {
			productCount[product.ProductID] = product.Quantity
			tags, err := pg.Queries.GetProductTag(c.Request().Context(), product.ProductID)
			productTag[product.ProductID] = NewTagSet(tags)
			if err != nil {
				logger.Errorw("failed to get product tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}
		// get coupon in cart
		cartCoupon, err := pg.Queries.GetCouponsFromCart(c.Request().Context(), db.GetCouponsFromCartParams{Username: username, CartID: param.CartID})
		if err != nil {
			logger.Errorw("failed to get coupon", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		// one shipping coupon per cart
		shippingFlag := false
		// sort CartCoupon by type (percentage => fixed), discount (high => low)
		sort.Slice(cartCoupon, func(i, j int) bool {
			if cartCoupon[i].Type == cartCoupon[j].Type {
				return cartCoupon[i].Discount.Int.Cmp(cartCoupon[j].Discount.Int) == 1
			}
			return cartCoupon[i].Type < cartCoupon[j].Type
		})
		// match coupons to its product, the product count would be reduced if matched
		for _, coupon := range cartCoupon {
			tags, err := pg.Queries.GetCouponTag(c.Request().Context(), coupon.ID)
			if err != nil {
				logger.Errorw("failed to get coupon tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if coupon.Type == db.CouponTypeShipping {
				shippingFlag = true
				continue
			}
			couponTags := NewTagSet(tags)
			for _, product := range cartProduct {
				if productCount[product.ProductID] == 0 {
					continue
				}
				if productTag[product.ProductID].Intersect(couponTags) {
					productCount[product.ProductID] -= 1
					break
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
		tags, err := pg.Queries.GetCouponTag(c.Request().Context(), param.CouponID)
		if err != nil {
			logger.Errorw("failed to get coupon tag", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		couponTag := NewTagSet(tags)
		// global coupon is always usable beside duplicate shipping coupon
		couponValid := couponDetail.Scope == db.CouponScopeGlobal
		for _, product := range cartProduct {
			if productCount[product.ProductID] == 0 {
				continue
			}
			if productTag[product.ProductID].Intersect(couponTag) {
				couponValid = true
				break
			}
		}
		if !couponValid {
			return echo.NewHTTPError(http.StatusBadRequest, "coupon not valid for any product in cart")
		}
		// if coupon is valid, add it to cart
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
