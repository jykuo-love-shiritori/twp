package router

import (
	"math"
	"net/http"
	"sort"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
		username := "Buyer"
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
		username := "Buyer"
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
	CartInfo db.GetCartRow
	Products []db.GetProductFromCartRow
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
func buyerGetCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
			cart.Products, err = pg.Queries.GetProductFromCart(c.Request().Context(), cartInfo.ID)
			if err != nil {
				logger.Errorw("failed to get product in cart", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			cart.Coupons, err = pg.Queries.GetCouponsFromCart(c.Request().Context(), db.GetCouponsFromCartParams{Username: username, CartID: cartInfo.ID})
			if err != nil {
				logger.Errorw("failed to get coupon in cart", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			cart.CartInfo = cartInfo
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
// @Param			cart_id		path		int		true	"Cart ID"
// @Param			product_id	path		int		true	"Product ID"
// @Param			quantity	body		ProductQuantity	true	"Quantity"
// @Success		200			{string}	string	constants.SUCCESS
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/buyer/cart/{cart_id}/product/{product_id} [patch]
func buyerEditProductInCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
			rows, err := pg.Queries.DeleteProductFromCart(c.Request().Context(),
				db.DeleteProductFromCartParams{
					Username:  username,
					CartID:    param.CartID,
					ProductID: param.ProductID,
				})
			if err != nil {
				logger.Errorw("failed to delete product in cart", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			if rows == 0 {
				return echo.NewHTTPError(http.StatusBadRequest)
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
// @Param			id	path		int	true	"Product ID"
// @Param			quantity	body		ProductQuantity	true	"Quantity"
// @Success		200			{integer}	int	"product quantity in cart"
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/buyer/cart/product/{id} [post]
func buyerAddProductToCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
// @Param			cart_id		path		int	true	"Cart ID"
// @Success		200			{array}		db.GetUsableCouponsRow
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/buyer/cart/{cart_id}/coupon [get]
func buyerGetCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
func buyerAddCouponToCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
func buyerDeleteProductFromCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var param db.DeleteProductFromCartParams
		param.Username = username
		if err := echo.PathParamsBinder(c).Int32("cart_id", &param.CartID).Int32("product_id", &param.ProductID).BindError(); err != nil {
			logger.Errorw("failed to parse cart_id or product_id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if rows, err := pg.Queries.DeleteProductFromCart(c.Request().Context(), param); err != nil {
			logger.Errorw("failed to delete product in cart", "error", err)
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
func buyerDeleteCouponFromCart(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
}

func getShipmentFee(total int32) int32 {
	return int32(math.Log(0.05 * float64(total)))
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
func buyerGetCheckout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var result checkout
		var cartID int32
		if err := echo.PathParamsBinder(c).Int32("cart_id", &cartID).BindError(); err != nil {
			logger.Errorw("failed to parse cart_id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		// this will validate cart and product legitimacy
		subtotal, err := pg.Queries.GetCartSubtotal(c.Request().Context(), cartID)
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
				cp.DiscountValue = result.Shipment
			}
			totalDiscount += cp.DiscountValue
			result.Coupons = append(result.Coupons, cp)
		}
		result.TotalDiscount = totalDiscount
		result.Total = max(0, result.Subtotal+result.Shipment-result.TotalDiscount)
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Buyer Checkout
// @Description	Checkout
// @Tags			Buyer, Checkout
// @Accept			json
// @Produce		json
// @param			cart_id	path		int		true	"Cart ID"
// @Success		200		{string}	string	constants.SUCCESS
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/buyer/cart/{cart_id}/checkout [post]
func buyerCheckout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var cartID int32
		if err := echo.PathParamsBinder(c).Int32("cart_id", &cartID).BindError(); err != nil {
			logger.Errorw("failed to parse cart_id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		products, err := pg.Queries.GetProductFromCart(c.Request().Context(), cartID)
		if err != nil {
			logger.Errorw("failed to get product from cart")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for _, product := range products {
			if product.Quantity > product.Stock {
				return echo.NewHTTPError(http.StatusBadRequest, "some product out of stock")
			}
			err := pg.Queries.UpdateProductVersion(c.Request().Context(), product.ProductID)
			if err != nil {
				logger.Errorw("failed to update product version")
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}
		// this will validate cart and product legitimacy
		subtotal, err := pg.Queries.GetCartSubtotal(c.Request().Context(), cartID)
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
				couponShippingFlag = true
			}
		}
		total := max(0, int32(subtotal)+shipment-totalDiscount) // if total < 0 => get achievement "🤑"

		if err := pg.Queries.Checkout(c.Request().Context(), db.CheckoutParams{Username: username, Shipment: shipment, CartID: cartID, TotalPrice: total}); err != nil {
			logger.Errorw("failed to checkout", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}
