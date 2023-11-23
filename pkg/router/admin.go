package router

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary Admin Get User
// @Description Get all user information. Include user's icon, name, email, created time and role.
// @Tags Admin, User
// @Produce json
// @Success 200 {array} db.GetUsersRow
// @Failure 401
// @Router /admin/user [get]
func adminGetUser(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := pg.Queries.GetUsers(c.Request().Context())
		if err != nil {
			logger.Errorw("failed to get users", "error", err)
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, users)
	}
}

// @Summary Admin Delete(disable) User
// @Description Delete(disable) user.
// @Tags Admin, User
// @Produce json
// @param id path int true "User ID"
// @Success 200 {string} string constants.SUCCESS
// @Failure 401
// @Router /admin/user/{username} [delete]
func adminDisableUser(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		id, err := pg.Queries.GetUserIDByUsername(c.Request().Context(), username)
		if err != nil {
			logger.Errorw("failed to get user id", "error", err)
			return c.JSON(http.StatusInternalServerError, nil)
		}
		if err := pg.Queries.DisableUser(c.Request().Context(), id); err != nil {
			logger.Errorw("failed to disable user", "error", err)
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

// @Summary Admin Get Coupon
// @Description Get all coupons (include shops).
// @Tags Admin, Coupon
// @Produce json
// @Success 200 {array} db.Coupon
// @Failure 401
// @Router /admin/coupon [get]
func adminGetCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := pg.Queries.GetAnyCoupons(c.Request().Context())
		if err != nil {
			logger.Errorw("failed to disable user", "error", err)
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary Admin Get Coupon Detail
// @Description Get coupon details.
// @Tags Admin, Coupon, Shop
// @Produce json
// @Param id path int true "Coupon ID"
// @Success 200 {object} db.Coupon
// @Failure 401
// @Router /admin/coupon/{id} [get]
func adminGetCouponDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int32
		if err := c.Bind(&id); err != nil {
			logger.Errorw("failed to bind id", "error", err)
			return c.JSON(http.StatusBadRequest, nil)
		}
		if result, err := pg.Queries.GetCouponDetail(c.Request().Context(), id); err != nil {
			logger.Errorw("failed to get coupon detail", "error", err)
			return c.JSON(http.StatusInternalServerError, nil)
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
// @Success 200
// @Failure 401
// @Router /admin/coupon [post]
func adminAddCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var coupon db.AddCouponParams
		if err := c.Bind(&coupon); err != nil {
			logger.Errorw("failed to bind coupon", "error", err)
			return c.JSON(http.StatusBadRequest, nil)
		}
		coupon.Scope = "global"
		result, err := pg.Queries.AddCoupon(c.Request().Context(), coupon)
		if err != nil {
			logger.Errorw("failed to add coupon", "error", err)
			return c.JSON(http.StatusInternalServerError, nil)
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
// @Success 200
// @Failure 401
// @Router /admin/coupon/{id} [patch]
func adminEditCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var coupon db.EditCouponParams
		if err := c.Bind(&coupon); err != nil {
			logger.Errorw("failed to bind coupon", "error", err)
			return c.JSON(http.StatusBadRequest, nil)
		}
		result, err := pg.Queries.EditCoupon(c.Request().Context(), coupon)
		if err != nil {
			logger.Errorw("failed to edit coupon", "error", err)
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary Admin Delete Coupon
// @Description Delete coupon (include shops').
// @Tags Admin, Coupon
// @Produce json
// @Param id path int true "Coupon ID"
// @Success 200
// @Failure 401
// @Router /admin/coupon/{id} [delete]
func adminDeleteCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int32
		if err := c.Bind(&id); err != nil {
			logger.Errorw("failed to bind id", "error", err)
			return c.JSON(http.StatusBadRequest, nil)
		}
		if err := pg.Queries.DeleteCoupon(c.Request().Context(), id); err != nil {
			logger.Errorw("failed to delete coupon", "error", err)
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

// @Summary Admin Get Site Report
// @Description Get site report.
// @Tags Admin, Report
// @Produce json
// @Success 200
// @Failure 401
// @Router /admin/report [get]
func adminGetReport(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var report db.GetReportParams
		if err := c.Bind(&report); err != nil {
			logger.Errorw("failed to bind report", "error", err)
			return c.JSON(http.StatusBadRequest, nil)
		}
		result, err := pg.Queries.GetReport(c.Request().Context(), report)
		if err != nil {
			logger.Errorw("failed to get report", "error", err)
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.JSON(http.StatusOK, result)
	}
}
