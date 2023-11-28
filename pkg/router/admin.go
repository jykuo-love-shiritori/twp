package router

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary		Admin Get User
// @Description	Get all user information. Include user's icon, name, email, created time and role.
// @Tags			Admin, User
// @Produce		json
// @Success		200
// @Failure		401
// @Router			/admin/user [get]
func adminGetUser(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Admin Delete User
// @Description	Delete user.
// @Tags			Admin, User
// @Produce		json
// @param			id	path	int	true	"User ID"
// @Success		200
// @Failure		401
// @Router			/admin/user/{id} [delete]
func adminDeleteUser(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Admin Get Coupon
// @Description	Get all coupons (include shops).
// @Tags			Admin, Coupon
// @Produce		json
// @Success		200
// @Failure		401
// @Router			/admin/coupon [get]
func adminGetCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Admin Get Coupon Detail
// @Description	Get coupon details.
// @Tags			Admin, Coupon, Shop
// @Produce		json
// @Param			id	path	int	true	"Coupon ID"
// @Success		200
// @Failure		401
// @Router			/admin/coupon/{id} [get]
func adminGetCouponDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Admin Add Coupon
// @Description	Add global coupon.
// @Tags			Admin, Coupon
// @Accept			json
// @Produce		json
// @Success		200
// @Failure		401
// @Router			/admin/coupon [post]
func adminAddCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Admin Edit Coupon
// @Description	Edit global coupon.
// @Tags			Admin, Coupon
// @Accept			json
// @Produce		json
// @Param			id	path	int	true	"Coupon ID"
// @Success		200
// @Failure		401
// @Router			/admin/coupon/{id} [patch]
func adminEditCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Admin Delete Coupon
// @Description	Delete coupon (include shops').
// @Tags			Admin, Coupon
// @Produce		json
// @Param			id	path	int	true	"Coupon ID"
// @Success		200
// @Failure		401
// @Router			/admin/coupon/{id} [delete]
func adminDeleteCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}

// @Summary		Admin Get Site Report
// @Description	Get site report.
// @Tags			Admin, Report
// @Produce		json
// @Success		200
// @Failure		401
// @Router			/admin/report [get]
func adminGetReport(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}
