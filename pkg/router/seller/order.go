package seller

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type OrderDetail struct {
	OrderInfo db.SellerGetOrderHistoryRow  `json:"order_info"`
	Products  []db.SellerGetOrderDetailRow `json:"products"`
}

type OrderUpdateStatusParams struct {
	ID            int32          `json:"id" param:"id"`
	CurrentStatus db.OrderStatus `json:"current_status"`
}

// @Summary		Seller get order
// @Description	Get all orders for shop.
// @Tags			Seller, Shop, Order
// @Param			offset	query	int	true	"offset"	default(0)	minimum(0)
// @Param			limit	query	int	true	"limit"		default(10)	maximum(20)
// @Produce		json
// @Success		200	{array}		db.SellerGetOrderRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/order [get]
func GetOrder(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		var username string = "user1"

		var requestParam common.QueryParams
		if err := c.Bind(&requestParam); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		if requestParam.Validate() != nil {
			logger.Errorw("invalid query parameter", "offset", requestParam.Offset, "limit", requestParam.Limit)
			return echo.NewHTTPError(http.StatusBadRequest, "offset or limit is invalid")
		}
		param := db.SellerGetOrderParams{SellerName: username, Limit: requestParam.Limit, Offset: requestParam.Offset}

		orders, err := pg.Queries.SellerGetOrder(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range orders {
			orders[i].ImageUrl = mc.GetFileURL(c.Request().Context(), orders[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, orders)
	}
}

// @Summary		Seller get order detail
// @Description	Get order detail by ID for shop.
// @Tags			Seller, Shop, Order
// @Produce		json
// @Param			id	path		int	true	"Order ID"
// @Success		200	{object}	OrderDetail
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/order/{id} [get]
func GetOrderDetail(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerGetOrderDetailParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		param.SellerName = username
		var result OrderDetail
		var err error
		result.OrderInfo, err = pg.Queries.SellerGetOrderHistory(c.Request().Context(), db.SellerGetOrderHistoryParams{SellerName: param.SellerName, ID: param.OrderID})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.Products, err = pg.Queries.SellerGetOrderDetail(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.OrderInfo.ThumbnailUrl = mc.GetFileURL(c.Request().Context(), result.OrderInfo.ThumbnailUrl)
		result.OrderInfo.UserImageUrl = mc.GetFileURL(c.Request().Context(), result.OrderInfo.UserImageUrl)
		for i := range result.Products {
			result.Products[i].ImageUrl = mc.GetFileURL(c.Request().Context(), result.Products[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Seller update order status
// @Description	seller update orders status.
// @Tags			Seller, Shop, Order
// @Param			id		path		int						true	"Order ID"
// @Param			param	body		OrderUpdateStatusParams	true	"order current status"
// @Success		200		{object}	db.SellerUpdateOrderStatusRow
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/seller/order/{id} [patch]
func UpdateOrderStatus(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param OrderUpdateStatusParams
		if err := c.Bind(&param); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)

		}

		// shop can only a prove the status traction {paid > shipped ,shipped > delivered}
		// paid > shipped > delivered > (cancelled || finished)
		var NextStatus db.OrderStatus
		switch param.CurrentStatus {
		case db.OrderStatusPaid:
			NextStatus = db.OrderStatusShipped
		case db.OrderStatusShipped:
			NextStatus = db.OrderStatusDelivered
		default:
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		order, err := pg.Queries.SellerUpdateOrderStatus(c.Request().Context(), db.SellerUpdateOrderStatusParams{
			SellerName:    username,
			ID:            param.ID,
			CurrentStatus: param.CurrentStatus,
			SetStatus:     NextStatus,
		})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, order)
	}
}
