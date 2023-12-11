package router

import (
	"net/http"
	"time"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"

	"go.uber.org/zap"
)

// @Summary		Admin Get User
// @Description	Get all user information. Include user's icon, name, email, created time and role.
// @Tags			Admin, User
// @Produce		json
// @Param			offset	query		int	false	"Begin index"	default(0)
// @Param			limit	query		int	false	"limit"			default(10)	maximum(20)
// @Success		200		{array}		db.GetUsersRow
// @Failure		400		{object}	echo.HTTPError
// @Router			/admin/user [get]
func adminGetUser(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
		users, err := pg.Queries.GetUsers(c.Request().Context(), db.GetUsersParams{Offset: q.Offset, Limit: q.Limit})
		if err != nil {
			logger.Errorw("failed to get users", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, users)
	}
}

// @Summary		Admin Disable User
// @Description	Disable user.
// @Tags			Admin, User
// @Produce		json
// @param			username	path		string	true	"Username"
// @Success		200			{string}	string	constants.SUCCESS
// @Failure		400			{object}	echo.HTTPError
// @Failure		404			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/admin/user/{username} [delete]
func adminDisableUser(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		if username == "" { // would not happen in future
			logger.Errorw("username is empty")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if execRows, err := pg.Queries.DisableUser(c.Request().Context(), username); err != nil {
			logger.Errorw("failed to disable user", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to disable user")
		} else if execRows == 0 {
			logger.Infow("user not found", "username", username)
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

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
func adminGetCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
func adminGetCouponDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
			logger.Infow("coupon not found", "id", id)
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
func adminAddCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var coupon db.AddCouponParams
		if err := c.Bind(&coupon); err != nil {
			logger.Errorw("failed to bind coupon", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		discount, err := coupon.Discount.Float64Value()
		if err != nil {
			logger.Errorw("failed to parse discount", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid discount")
		}
		if discount.Float64 < 0 || discount.Float64 > 100 {
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
// @Description	Edit global coupon.
// @Tags			Admin, Coupon
// @Accept			json
// @Produce		json
// @Param			id		path		int				true	"Coupon ID"
// @Param			coupon	body		PrettierCoupon	true	"Coupon"
// @Success		200		{object}	db.EditCouponRow
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/admin/coupon/{id} [patch]
func adminEditCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
		if discount.Float64 < 0 || discount.Float64 > 100 {
			logger.Errorw("discount is invalid", "discount", discount)
			return echo.NewHTTPError(http.StatusBadRequest, "Discount is invalid")
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
func adminDeleteCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
			logger.Infow("coupon not found", "id", id)
			return echo.NewHTTPError(http.StatusNotFound, "Coupon not found")
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

type adminReport struct {
	Year        int32                `json:"year"`
	Month       int32                `json:"month"`
	TotalAmount int32                `json:"total"`
	Sellers     []db.GetTopSellerRow `json:"sellers"`
}

// @Summary		Admin Get Site Report
// @Description	Get site report (top 3 sellers and total amount).
// @Tags			Admin, Report
// @Produce		json
// @Param			date	query		string	true	"Start year/month"
// @Success		200		{object}	adminReport
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/admin/report [get]
func adminGetReport(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		date := c.QueryParam("date")
		// verify the date is any start day of the month
		t, err := time.Parse(time.RFC3339, date)
		if err != nil {
			logger.Errorw("failed to parse date", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if t.Day() != 1 {
			logger.Errorw("date is not the first day of the month", "date", date)
			return echo.NewHTTPError(http.StatusBadRequest, "Date is not the first day of the month")
		}
		year, month, day := t.Date()
		result := adminReport{Year: int32(year), Month: int32(month)}
		// set date's time to 00:00:00
		formattedDate := time.Date(year, month, day, 0, 0, 0, 0, t.Location()).Format(time.RFC3339)
		result.Sellers, err = pg.Queries.GetTopSeller(c.Request().Context(), formattedDate)
		if err != nil {
			logger.Errorw("failed to get top seller", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		// get total amount by sum the total_sales of sellers
		result.TotalAmount, err = pg.Queries.GetTotalSales(c.Request().Context(), date)
		if err != nil {
			logger.Errorw("failed to get total sales", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}
