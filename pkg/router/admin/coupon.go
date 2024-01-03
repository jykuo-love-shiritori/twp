package admin

import (
	"net/http"
	"time"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary		Admin Get Coupon
// @Description	Get all global coupons .
// @Tags			Admin, Coupon
// @Produce		json
// @param			offset	query		int	false	"Begin index"	default(0)
// @param			limit	query		int	false	"limit"			default(10)
// @Success		200		{array}		db.GetGlobalCouponsRow
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/admin/coupon [get]
func GetCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := common.NewQueryParams(0, 10)
		if err := c.Bind(&q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if q.Validate() != nil {
			logger.Errorw("invalid query parameter", "offset", q.Offset, "limit", q.Limit)
			return echo.NewHTTPError(http.StatusBadRequest, "offset or limit is invalid")
		}
		result, err := pg.Queries.GetGlobalCoupons(c.Request().Context(), db.GetGlobalCouponsParams{Offset: q.Offset, Limit: q.Limit})
		if err != nil {
			logger.Errorw("failed to get coupons", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Admin Get Coupon Detail
// @Description	Get coupon details.
// @Tags			Admin, Coupon, Shop
// @Produce		json
// @Param			id	path		int	true	"Coupon ID"
// @Success		200	{object}	db.GetGlobalCouponDetailRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/admin/coupon/{id} [get]
func GetCouponDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int32
		if err := echo.PathParamsBinder(c).Int32("id", &id).BindError(); err != nil {
			logger.Errorw("failed to parse id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if exist, err := pg.Queries.CouponExists(c.Request().Context(), id); err != nil {
			logger.Errorw("failed to check coupon exist", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if !exist {
			logger.Errorw("coupon not found", "id", id)
			return echo.NewHTTPError(http.StatusNotFound, "Coupon not found")
		}
		result, err := pg.Queries.GetGlobalCouponDetail(c.Request().Context(), id)
		if err != nil {
			logger.Errorw("failed to get coupon detail", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Admin Add Coupon
// @Description	Add global coupon.
// @Tags			Admin, Coupon
// @Accept			json
// @Produce		json
// @Param			coupon	body		db.AddCouponParams	true	"Coupon"
// @Success		200		{object}	db.AddCouponRow
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/admin/coupon [post]
func AddCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var coupon db.AddCouponParams
		if err := c.Bind(&coupon); err != nil {
			logger.Errorw("failed to bind coupon", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		// if the the given time is invalid, make it become now 😇
		if coupon.StartDate.Time.Before(time.Now().Truncate(24 * time.Hour)) {
			coupon.StartDate.Time = time.Now()
		}
		if coupon.ExpireDate.Time.Before(coupon.StartDate.Time) {
			logger.Errorw("start date is invalid", "start date", coupon.StartDate)
			return echo.NewHTTPError(http.StatusBadRequest, "Start date is invalid")
		}
		discount, err := coupon.Discount.Float64Value()
		if err != nil {
			logger.Errorw("failed to parse discount", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid discount")
		}
		if discount.Float64 < 0 || (coupon.Type == db.CouponTypePercentage && discount.Float64 > 100) {
			logger.Errorw("discount is invalid", "discount", discount)
			return echo.NewHTTPError(http.StatusBadRequest, "Discount is invalid")
		}
		coupon.Scope = "global"
		result, err := pg.Queries.AddCoupon(c.Request().Context(), coupon)
		if err != nil {
			logger.Errorw("failed to add coupon", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}

type PrettierCoupon struct { // for swagger
	Type        db.CouponType `json:"type"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Discount    float32       `json:"discount"`
	StartDate   string        `json:"start_date"`
	ExpireDate  string        `json:"end_date"`
}

// @Summary		Admin Edit Coupon
// @Description	Edit global coupon. All the coupon properties are required.
// @Tags			Admin, Coupon
// @Accept			json
// @Produce		json
// @Param			id		path		int				true	"Coupon ID"
// @Param			coupon	body		PrettierCoupon	true	"Coupon"
// @Success		200		{object}	db.EditCouponRow
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/admin/coupon/{id} [patch]
func EditCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var coupon db.EditCouponParams
		if err := c.Bind(&coupon); err != nil {
			logger.Errorw("failed to bind coupon", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid coupon data")
		}
		discount, err := coupon.Discount.Float64Value()
		if err != nil {
			logger.Errorw("failed to parse discount", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid discount")
		}
		if discount.Float64 < 0 || (coupon.Type == db.CouponTypePercentage && discount.Float64 > 100) {
			logger.Errorw("discount is invalid", "discount", discount)
			return echo.NewHTTPError(http.StatusBadRequest, "Discount is invalid")
		}
		// if the the given time is invalid, make it become now 😇
		if coupon.StartDate.Time.Before(time.Now()) {
			coupon.StartDate.Time = time.Now()
		}
		if coupon.ExpireDate.Time.Before(coupon.StartDate.Time) {
			logger.Errorw("start date should later than ", "start date", coupon.StartDate)
			return echo.NewHTTPError(http.StatusBadRequest, "Expire date is invalid")
		}
		result, err := pg.Queries.EditCoupon(c.Request().Context(), coupon)
		if err != nil {
			logger.Errorw("failed to edit coupon", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Admin Delete Coupon
// @Description	Delete coupon (include shops').
// @Tags			Admin, Coupon
// @Produce		json
// @Param			id	path		int		true	"Coupon ID"
// @Success		200	{string}	string	constants.SUCCESS
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/admin/coupon/{id} [delete]
func DeleteCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int32
		if err := echo.PathParamsBinder(c).Int32("id", &id).BindError(); err != nil {
			logger.Errorw("failed to parse id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if execRows, err := pg.Queries.DeleteCoupon(c.Request().Context(), id); err != nil {
			logger.Errorw("failed to delete coupon", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if execRows == 0 {
			logger.Errorw("coupon not found", "id", id)
			return echo.NewHTTPError(http.StatusNotFound, "Coupon not found")
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}
