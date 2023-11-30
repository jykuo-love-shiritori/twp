package router

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
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
// @Description	Get all coupons (include shops' coupon).
// @Tags			Admin, Coupon
// @Produce		json
// @param			offset	query		int	false	"Begin index"	default(0)
// @param			limit	query		int	false	"limit"			default(10)
// @Success		200		{array}		db.Coupon
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
		result, err := pg.Queries.GetAnyCoupons(c.Request().Context(), db.GetAnyCouponsParams{Offset: q.Offset, Limit: q.Limit})
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
// @Success		200	{object}	db.GetCouponDetailRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/admin/coupon/{id} [get]
func adminGetCouponDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int32
		if err := echo.QueryParamsBinder(c).Int32("id", &id); err != nil {
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
		if result, err := pg.Queries.GetCouponDetail(c.Request().Context(), id); err != nil {
			logger.Errorw("failed to get coupon detail", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else {
			return c.JSON(http.StatusOK, result)
		}
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
		coupon.Scope = "global"
		result, err := pg.Queries.AddCoupon(c.Request().Context(), coupon)
		if err != nil {
			logger.Errorw("failed to add coupon", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Admin Edit Coupon
// @Description	Edit any coupon.
// @Tags			Admin, Coupon
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Coupon ID"
// @Param			coupon	body		db.EditCouponParams	true	"Coupon"
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
		if err := echo.QueryParamsBinder(c).Int32("id", &id); err != nil {
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

// For further implementation
type dateParams struct {
	StartDate pgtype.Timestamptz `query:"start_date"`
	EndDate   pgtype.Timestamptz `query:"end_date"`
}

// TODO
// @Summary		Admin Get Site Report
// @Description	Get site report.
// @Tags			Admin, Report
// @Produce		json
// @Param			start_date	query		string	true	"Start date"
// @Param			end_date	query		string	true	"End date"
// @Success		200			{string}	string "TODO"
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/admin/report [get]
func adminGetReport(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var _ dateParams // keep for future
		return echo.NewHTTPError(http.StatusNotImplemented)
	}
}
