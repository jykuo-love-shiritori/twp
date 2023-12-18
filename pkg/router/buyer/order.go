package buyer

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
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
func GetOrderHistory(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		q := common.NewQueryParams(0, 10)
		if err := c.Bind(&q); err != nil {
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
		for i := range orders {
			orders[i].ShopImageUrl = mc.GetFileURL(c.Request().Context(), orders[i].ShopImageUrl)
			orders[i].ThumbnailUrl = mc.GetFileURL(c.Request().Context(), orders[i].ThumbnailUrl)
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
func GetOrderDetail(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := "Buyer"
		var orderID int32
		var orderDetail OrderDetail
		var err error
		if echo.PathParamsBinder(c).Int32("id", &orderID).BindError() != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if orderDetail.Info, err = pg.Queries.GetOrderInfo(c.Request().Context(),
			db.GetOrderInfoParams{
				Username: username,
				OrderID:  orderID}); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return echo.NewHTTPError(http.StatusBadRequest)
			}
			logger.Errorw("failed to get order info", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		orderDetail.Info.ShopImageUrl = mc.GetFileURL(c.Request().Context(), orderDetail.Info.ShopImageUrl)
		if orderDetail.Details, err = pg.Queries.GetOrderDetail(c.Request().Context(), orderID); err != nil {
			logger.Errorw("failed to get order detail", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range orderDetail.Details {
			orderDetail.Details[i].ImageUrl = mc.GetFileURL(c.Request().Context(), orderDetail.Details[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, orderDetail)
	}
}
