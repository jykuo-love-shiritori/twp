package buyer

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type couponDiscount struct {
	ID            int32          `json:"id"`
	Name          string         `json:"name"`
	Type          db.CouponType  `json:"type"`
	Scope         db.CouponScope `json:"scope"`
	Description   string         `json:"description"`
	Discount      float64        `json:"discount"`
	DiscountValue int32          `json:"discount_value"`
}

type checkout struct {
	Subtotal      int32            `json:"subtotal"`
	Shipment      int32            `json:"shipment"`
	TotalDiscount int32            `json:"total_discount"`
	Coupons       []couponDiscount `json:"coupons"`
	Total         int32            `json:"total"`
	Payments      json.RawMessage  `json:"payments"`
}

func getShipmentFee(total int32) int32 {
	return int32(10 * math.Log(100*float64(total)))
}
func getDiscountValue(price float64, discount float64, couponType db.CouponType) int32 {
	switch couponType {
	case db.CouponTypePercentage:
		return int32(price * discount / 100)
	case db.CouponTypeFixed:
		return min(int32(discount), int32(price))
	}
	return 0
}

// @Summary		Buyer Get Checkout
// @Description	Get all checkout data
// @Tags			Buyer, Checkout
// @Produce		json
// @Param			id	path		int	true	"Cart ID"
// @Success		200		{object}	checkout
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/buyer/cart/{id}/checkout [get]
func GetCheckout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		result := checkout{Coupons: []couponDiscount{}}
		var cartID int32
		if err := echo.PathParamsBinder(c).Int32("id", &cartID).BindError(); err != nil {
			logger.Errorw("failed to parse cart_id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		// this will validate product stock and cart legitimacy
		valid, err := pg.Queries.ValidateProductsInCart(c.Request().Context(), db.ValidateProductsInCartParams{
			Username: username,
			CartID:   cartID,
		})
		if err != nil {
			logger.Errorw("failed to validate product in cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if !valid {
			return echo.NewHTTPError(http.StatusBadRequest, "Some product is not available now")
		}
		productCount := make(map[int32]int32)
		productTag := make(map[int32]*tagSet)
		products, err := pg.Queries.GetProductFromCartOrderByPriceDesc(c.Request().Context(), cartID)
		if err != nil {
			logger.Errorw("failed to get product from cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		// sort to make customer get most discount
		// calculate subtotal and get tags and counts
		for _, product := range products {
			// each product can have its coupon
			productCount[product.ProductID] = product.Quantity
			price, err := product.Price.Float64Value()
			if err != nil {
				logger.Errorw("failed to get price", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			// calculate subtotal
			result.Subtotal += int32(price.Float64 * float64(product.Quantity))
			tags, err := pg.Queries.GetProductTag(c.Request().Context(), product.ProductID)
			if err != nil {
				logger.Errorw("failed to get product tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			productTag[product.ProductID] = NewTagSet(tags)
		}
		result.Shipment = getShipmentFee(result.Subtotal)
		coupons, err := pg.Queries.GetSortedCouponsFromCart(c.Request().Context(), db.GetSortedCouponsFromCartParams{
			Username: username,
			CartID:   cartID,
		})
		if err != nil {
			logger.Errorw("failed to get coupons from cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		totalDiscount := int32(0)
		for _, coupon := range coupons {
			// init couponDiscount for returning
			var cd couponDiscount = couponDiscount{
				ID:          coupon.ID,
				Name:        coupon.Name,
				Type:        coupon.Type,
				Scope:       coupon.Scope,
				Description: coupon.Description,
			}
			discount, err := coupon.Discount.Float64Value()
			if err != nil {
				logger.Errorw("failed to get discount", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			cd.Discount = discount.Float64
			// if coupon is shipping coupon, calculate discount and continue
			if cd.Type == db.CouponTypeShipping {
				cd.DiscountValue = result.Shipment * (int32(cd.Discount / 100))
				totalDiscount += cd.DiscountValue
				result.Coupons = append(result.Coupons, cd)
				continue
			}
			tags, err := pg.Queries.GetCouponTag(c.Request().Context(), coupon.ID)
			if err != nil {
				logger.Errorw("failed to get coupon tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			couponTags := NewTagSet(tags)
			// match coupon with product
			for _, product := range products {
				if productCount[product.ProductID] == 0 {
					continue
				}
				if productTag[product.ProductID].Intersect(couponTags) {
					productCount[product.ProductID] -= 1 // decrease product count
					price, err := product.Price.Float64Value()
					if err != nil {
						logger.Errorw("failed to get price", "error", err)
						return echo.NewHTTPError(http.StatusInternalServerError)
					}
					cd.DiscountValue = getDiscountValue(price.Float64, cd.Discount, cd.Type)
					break
				}
			}
			totalDiscount += cd.DiscountValue
			result.Coupons = append(result.Coupons, cd)
		}
		result.TotalDiscount = totalDiscount
		// make it geq 0
		result.Total = max(0, result.Subtotal+result.Shipment-result.TotalDiscount)
		result.Payments, err = pg.Queries.GetCreditCard(c.Request().Context(), username)
		if err != nil {
			logger.Errorw("failed to get credit card", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}

type PaymentMethod struct {
	CreditCard json.RawMessage `json:"credit_card" swaggertype:"object"`
}

// @Summary		Buyer Checkout
// @Description	Checkout
// @Tags			Buyer, Checkout
// @Accept			json
// @Produce		json
// @param			id			path		int				true	"Cart ID"
// @Param			payment_method	body		PaymentMethod	true	"Payment" Example
// @Success		200				{string}	string			constants.SUCCESS
// @Failure		400				{object}	echo.HTTPError
// @Failure		500				{object}	echo.HTTPError
// @Router			/buyer/cart/{id}/checkout [post]
func Checkout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var cartID int32
		if err := echo.PathParamsBinder(c).Int32("id", &cartID).BindError(); err != nil {
			logger.Errorw("failed to parse cart_id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		var param db.ValidatePaymentParams
		param.Username = username
		if err := c.Bind(&param); err != nil {
			logger.Errorw("failed to bind payment", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if valid, err := pg.Queries.ValidatePayment(c.Request().Context(), param); err != nil {
			logger.Errorw("failed to validate payment", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if !valid {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid payment")
		}
		valid, err := pg.Queries.ValidateProductsInCart(c.Request().Context(), db.ValidateProductsInCartParams{
			Username: username,
			CartID:   cartID,
		})
		if err != nil {
			logger.Errorw("failed to validate product in cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if !valid {
			return echo.NewHTTPError(http.StatusBadRequest, "Some product is not available now")
		}
		// all the following logic is basically same as the GetCheckout
		// the only diff is that we need to update product version and no return
		products, err := pg.Queries.GetProductFromCartOrderByPriceDesc(c.Request().Context(), cartID)
		if err != nil {
			logger.Errorw("failed to get product from cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		tx, err := pg.NewTx(c.Request().Context())
		if err != nil {
			logger.Errorw("failed to create transaction", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		subtotal := int32(0)
		productCount := make(map[int32]int32)
		productTag := make(map[int32]*tagSet)
		// calculate subtotal and get tags and counts
		for _, product := range products {
			err := pg.Queries.WithTx(tx).UpdateProductVersion(c.Request().Context(), product.ProductID)
			if err != nil {
				logger.Errorw("failed to update product version", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			price, err := product.Price.Float64Value()
			if err != nil {
				logger.Errorw("failed to get price", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			subtotal += int32(price.Float64 * float64(product.Quantity))
			productCount[product.ProductID] = product.Quantity
			tags, err := pg.Queries.GetProductTag(c.Request().Context(), product.ProductID)
			if err != nil {
				logger.Errorw("failed to get product tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			productTag[product.ProductID] = NewTagSet(tags)
		}
		// this will validate cart and product legitimacy
		if valid, err := pg.Queries.ValidateProductsInCart(c.Request().Context(), db.ValidateProductsInCartParams{
			Username: username,
			CartID:   cartID,
		}); err != nil {
			logger.Errorw("failed to validate product in cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if !valid {
			return echo.NewHTTPError(http.StatusBadRequest, "Some product is not available now")
		}
		shipment := getShipmentFee(int32(subtotal))
		var params db.GetSortedCouponsFromCartParams
		params.CartID = cartID
		params.Username = username
		coupons, err := pg.Queries.GetSortedCouponsFromCart(c.Request().Context(), params)
		if err != nil {
			logger.Errorw("failed to get coupons from cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		totalDiscount := int32(0)
		// match coupon with product
		for _, coupon := range coupons {
			discount, err := coupon.Discount.Float64Value()
			if err != nil {
				logger.Errorw("failed to get discount", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			dc := discount.Float64
			if coupon.Type == db.CouponTypeShipping {
				totalDiscount += shipment * (int32(dc / 100))
				continue
			}
			tags, err := pg.Queries.GetCouponTag(c.Request().Context(), coupon.ID)
			if err != nil {
				logger.Errorw("failed to get coupon tag", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			couponTags := NewTagSet(tags)
			for _, product := range products {
				if productCount[product.ProductID] == 0 {
					continue
				}
				if productTag[product.ProductID].Intersect(couponTags) {
					productCount[product.ProductID] -= 1
					price, err := product.Price.Float64Value()
					if err != nil {
						logger.Errorw("failed to get price", "error", err)
						return echo.NewHTTPError(http.StatusInternalServerError)
					}
					totalDiscount += getDiscountValue(price.Float64, dc, coupon.Type)
					break
				}
			}
		}
		total := max(0, int32(subtotal)+shipment-totalDiscount)
		//  if total < 0, consider get achievement “There is nothing more expensive than something free”
		if err := pg.Queries.WithTx(tx).Checkout(c.Request().Context(),
			db.CheckoutParams{
				Username:   username,
				Shipment:   shipment,
				CartID:     cartID,
				TotalPrice: total}); err != nil {
			logger.Errorw("failed to checkout", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if err := tx.Commit(c.Request().Context()); err != nil {
			logger.Errorw("failed to commit transaction", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}
