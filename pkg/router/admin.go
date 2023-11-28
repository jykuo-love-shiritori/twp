package router

import (
	"net/http"
	"strconv"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"

	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"go.uber.org/zap"
)

// @Summary Admin Get User
// @Description Get all user information. Include user's icon, name, email, created time and role.
// @Tags Admin, User
// @Produce json
// @Param offset query int false "Begin index" default(0)
// @Param limit query int false "limit" default(10)
// @Success 200 {array} db.GetUsersRow
// @Failure 400 {object} Failure
// @Router /admin/user [get]
func adminGetUser(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := common.NewQueryParams(0, 10)
		if err := c.Bind(&q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		users, err := pg.Queries.GetUsers(c.Request().Context())
		if err != nil {
			logger.Errorw("failed to get users", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if int(q.Offset) > len(users) {
			return echo.NewHTTPError(http.StatusBadRequest, "Offset out of range")
		}
		q.Limit = min(q.Limit, int32(len(users))-q.Offset)
		return c.JSON(http.StatusOK, users[q.Offset:q.Offset+q.Limit])
	}
}

// @Summary Admin Disable User
// @Description Disable user.
// @Tags Admin, User
// @Produce json
// @param username path string true "Username"
// @Success 200 {string} string constants.SUCCESS
// @Failure 400 {object} Failure
// @Failure 404 {object} Failure
// @Failure 500 {object} Failure
// @Router /admin/user/{username} [patch]
func adminDisableUser(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		if username == "" {
			logger.Errorw("username is empty")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		id, err := pg.Queries.GetUserIDByUsername(c.Request().Context(), username)
		if err != nil {
			logger.Errorw("failed to get user id", "error", err)
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		if err := pg.Queries.DisableUser(c.Request().Context(), id); err != nil {
			logger.Errorw("failed to disable user", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to disable user")
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

// @Summary Admin Get Coupon
// @Description Get all coupons (include shops' coupon).
// @Tags Admin, Coupon
// @Produce json
// @param offset query int false "Begin index" default(0)
// @param limit query int false "limit" default(10)
// @Success 200 {array} db.Coupon
// @Failure 400 {object} Failure
// @Failure 500 {object} Failure
// @Router /admin/coupon [get]
func adminGetCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		q := common.NewQueryParams(0, 10)
		if err := c.Bind(&q); err != nil {
			logger.Errorw("failed to bind query parameter", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		result, err := pg.Queries.GetAnyCoupons(c.Request().Context())
		if err != nil {
			logger.Errorw("failed to get coupons", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if int(q.Offset) > len(result) {
			return echo.NewHTTPError(http.StatusBadRequest, "Offset out of range")
		}
		q.Limit = min(q.Limit, int32(len(result))-q.Offset)
		return c.JSON(http.StatusOK, result[q.Offset:q.Offset+q.Limit])
	}
}

// @Summary Admin Get Coupon Detail
// @Description Get coupon details.
// @Tags Admin, Coupon, Shop
// @Produce json
// @Param id path int true "Coupon ID"
// @Success 200 {object} db.GetCouponDetailRow
// @Failure 400 {object} Failure
// @Failure 404 {object} Failure
// @Failure 500 {string} Failure
// @Router /admin/coupon/{id} [get]
func adminGetCouponDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			logger.Errorw("failed to parse id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		idInt32 := int32(idInt)
		if exist, err := pg.Queries.CouponExists(c.Request().Context(), idInt32); err != nil {
			logger.Errorw("failed to check coupon exist", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else if !exist {
			logger.Infow("coupon not found", "id", id)
			return echo.NewHTTPError(http.StatusNotFound, "Coupon not found")
		}
		if result, err := pg.Queries.GetCouponDetail(c.Request().Context(), idInt32); err != nil {
			logger.Errorw("failed to get coupon detail", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		} else {
			return c.JSON(http.StatusOK, result)
		}
	}
}

// @Summary Admin Add Coupon
// @Description Add global coupon.
// @Tags Admin, Coupon
// @Accept json
// @Produce json
// @Param coupon body db.AddCouponParams true "Coupon"
// @Success 200 {object} db.AddCouponRow
// @Failure 400 {object} Failure
// @Failure 500 {object} Failure
// @Router /admin/coupon [post]
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

// @Summary Admin Edit Coupon
// @Description Edit global coupon.
// @Tags Admin, Coupon
// @Accept json
// @Produce json
// @Param id path int true "Coupon ID"
// @Param coupon body db.EditCouponParams true "Coupon"
// @Success 200 {object} db.EditCouponRow
// @Failure 400 {object} Failure
// @Failure 500 {object} Failure
// @Router /admin/coupon/{id} [patch]
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

// @Summary Admin Delete Coupon
// @Description Delete coupon (include shops').
// @Tags Admin, Coupon
// @Produce json
// @Param id path int true "Coupon ID"
// @Success 200 {string} string constants.SUCCESS
// @Failure 400 {object} Failure
// @Failure 500 {object} Failure
// @Router /admin/coupon/{id} [delete]
func adminDeleteCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			logger.Errorw("failed to parse id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		idInt32 := int32(idInt)
		if err := c.Bind(&id); err != nil {
			logger.Errorw("failed to bind id", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := pg.Queries.DeleteCoupon(c.Request().Context(), idInt32); err != nil {
			logger.Errorw("failed to delete coupon", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

// @Summary Admin Get Site Report
// @Description Get site report.
// @Tags Admin, Report
// @Produce json
// @Param start_date query string true "Start date"
// @Param end_date query string true "End date"
// @Success 200 {array} db.OrderHistory
// @Failure 400 {object} Failure
// @Failure 500 {object} Failure
// @Router /admin/report [get]
func adminGetReport(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var report db.GetReportParams
		if err := c.Bind(&report); err != nil {
			logger.Errorw("failed to bind report", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		result, err := pg.Queries.GetReport(c.Request().Context(), report)
		if err != nil {
			logger.Errorw("failed to get report", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}
