package buyer

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/jackc/pgx/v5"
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
	return int32(math.Log(float64(total)))
}

// @Summary		Buyer Get Checkout
// @Description	Get all checkout data
// @Tags			Buyer, Checkout
// @Produce		json
// @Param			cart_id	path		int	true	"Cart ID"
// @Success		200		{object}	checkout
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/buyer/cart/{cart_id}/checkout [get]
func GetCheckout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		result := checkout{Coupons: []couponDiscount{}}
		var cartID int32
		if err := echo.PathParamsBinder(c).Int32("cart_id", &cartID).BindError(); err != nil {
			logger.Errorw("failed to parse cart_id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		// this will validate cart and product legitimacy
		subtotal, err := pg.Queries.GetCartSubtotal(c.Request().Context(),
			db.GetCartSubtotalParams{
				Username: username,
				CartID:   cartID})
		if err != nil {
			if err == pgx.ErrNoRows {
				return echo.NewHTTPError(http.StatusBadRequest, "there might have some product are not available now")
			}
			logger.Errorw("failed to get cart subtotal", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.Subtotal = int32(subtotal)
		result.Shipment = getShipmentFee(result.Subtotal)

		var params db.GetCouponsFromCartParams
		params.CartID = cartID
		params.Username = username
		coupons, err := pg.Queries.GetCouponsFromCart(c.Request().Context(), params)
		if err != nil {
			logger.Errorw("failed to get coupons from cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		couponPercentageFlag := false
		couponShippingFlag := false
		totalDiscount := int32(0)
		for _, coupon := range coupons {
			var cp couponDiscount = couponDiscount{
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
			cp.Discount = discount.Float64
			switch coupon.Type {
			case db.CouponTypePercentage:
				if couponPercentageFlag {
					logger.Errorw("multiple percentage coupon", "error", err)
					return echo.NewHTTPError(http.StatusBadRequest, "multiple percentage coupon")
				}
				couponPercentageFlag = true
				cp.DiscountValue = int32(float64(result.Subtotal) * discount.Float64 / 100)
			case db.CouponTypeFixed:
				cp.DiscountValue = int32(discount.Float64)
			case db.CouponTypeShipping:
				if couponShippingFlag {
					logger.Errorw("multiple shipping coupon", "error", err)
					return echo.NewHTTPError(http.StatusBadRequest, "multiple shipping coupon")
				}
				couponShippingFlag = true
				cp.DiscountValue = int32(float64(result.Shipment) * (discount.Float64 / 100.0))

			}
			totalDiscount += cp.DiscountValue
			result.Coupons = append(result.Coupons, cp)
		}
		result.TotalDiscount = totalDiscount
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
	CreditCard json.RawMessage `swaggertype:"object"`
}

// @Summary		Buyer Checkout
// @Description	Checkout
// @Tags			Buyer, Checkout
// @Accept			json
// @Produce		json
// @param			cart_id			path		int				true	"Cart ID"
// @Param			payment_method	body		PaymentMethod	true	"Payment"
// @Success		200				{string}	string			constants.SUCCESS
// @Failure		400				{object}	echo.HTTPError
// @Failure		500				{object}	echo.HTTPError
// @Router			/buyer/cart/{cart_id}/checkout [post]
func Checkout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var cartID int32
		if err := echo.PathParamsBinder(c).Int32("cart_id", &cartID).BindError(); err != nil {
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
		products, err := pg.Queries.GetProductFromCart(c.Request().Context(), cartID)
		if err != nil {
			logger.Errorw("failed to get product from cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for _, product := range products {
			if product.Quantity > product.Stock {
				return echo.NewHTTPError(http.StatusBadRequest, "some product out of stock")
			}
			err := pg.Queries.UpdateProductVersion(c.Request().Context(), product.ProductID)
			if err != nil {
				logger.Errorw("failed to update product version", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}
		// this will validate cart and product legitimacy
		subtotal, err := pg.Queries.GetCartSubtotal(c.Request().Context(),
			db.GetCartSubtotalParams{
				CartID:   cartID,
				Username: username})
		if err != nil {
			if err == pgx.ErrNoRows {
				return echo.NewHTTPError(http.StatusBadRequest, "there might have some product are not available now")
			}
			logger.Errorw("failed to get cart subtotal", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		shipment := getShipmentFee(int32(subtotal))
		var params db.GetCouponsFromCartParams
		params.CartID = cartID
		params.Username = username

		// this will validate coupon legitimacy
		coupons, err := pg.Queries.GetCouponsFromCart(c.Request().Context(), params)
		if err != nil {
			logger.Errorw("failed to get coupons from cart", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		couponPercentageFlag := false
		couponShippingFlag := false
		totalDiscount := int32(0)
		for _, coupon := range coupons {
			discount, err := coupon.Discount.Float64Value()
			if err != nil {
				logger.Errorw("failed to get discount", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			switch coupon.Type {
			case db.CouponTypePercentage:
				if couponPercentageFlag {
					logger.Errorw("multiple percentage coupon", "error", err)
					return echo.NewHTTPError(http.StatusBadRequest, "multiple percentage coupon")
				}
				couponPercentageFlag = true
				totalDiscount += int32(float64(subtotal) * discount.Float64 / 100)
			case db.CouponTypeFixed:
				totalDiscount += int32(discount.Float64)
			case db.CouponTypeShipping:
				if couponShippingFlag {
					logger.Errorw("multiple shipping coupon", "error", err)
					return echo.NewHTTPError(http.StatusBadRequest, "multiple shipping coupon")
				}
				totalDiscount += int32(float64(shipment) * (discount.Float64 / 100.0))
				couponShippingFlag = true
			}
		}
		total := max(0, int32(subtotal)+shipment-totalDiscount) // if total < 0 => get achievement "🤑"

		if err := pg.Queries.Checkout(c.Request().Context(),
			db.CheckoutParams{
				Username:   username,
				Shipment:   shipment,
				CartID:     cartID,
				TotalPrice: total}); err != nil {
			logger.Errorw("failed to checkout", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}
