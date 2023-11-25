package router

import (
	"context"
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary Logout
// @Description Logout the current user
// @Tags Auth
// @Produce json
// @Success 200
// @Failure 401
// @Router /logout [post]
func logout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary User Get Info
// @Description Get user information
// @Tags User
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /api/user/info [get]
func userGetInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID int32 = 1
		user, err := pg.Queries.UserGetInfo(context.Background(), userID)
		if err != nil {
			return DBResponse(c, err, "failed to get user info", logger)
		}
		return c.JSON(http.StatusOK, user)
	}
}

// @Summary User Edit Info
// @Description Edit user information
// @Tags User
// @Param  name          body     string  true  "name of coupon"
// @Param  address       body     string  true  "user address"
// @Param  image_id      body     string  true  "image id"
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /api/user/info [patch]
func userEditInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID int32 = 1

		var param db.UserUpdateInfoParams
		if err := c.Bind(&param); err != nil {
			return err
		}
		param.ID = userID
		order, err := pg.Queries.UserUpdateInfo(context.Background(), param)
		if err != nil {
			return DBResponse(c, err, "failed to user infomation", logger)
		}

		return c.JSON(http.StatusOK, order)
	}
}

// @Summary User Upload Avatar
// @Description Upload user avatar
// @Tags User
// @Accept png,jpeg,gif
// @Produce json
// @Param img formData file true "Image to upload"
// @Success 200
// @Failure 401
// @Router /api/user/avatar [post]
func userUploadAvatar(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary User Edit Password
// @Description Change user password
// @Tags User
// @Param  current_password  body     string  true  "current password"
// @Param  new_password      body     string  true  "new password"
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /api/user/security/password [post]
func userEditPassword(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID int32 = 1
		var param db.UserUpdatePasswordParams
		if err := c.Bind(&param); err != nil {
			return err
		}
		param.ID = userID
		orders, err := pg.Queries.UserUpdatePassword(context.Background(), param)
		if err != nil {
			return DBResponse(c, err, "failed to change user password", logger)
		}
		return c.JSON(http.StatusOK, orders)
	}
}

// @Summary User Get Credit Card
// @Description Get all credit cards of the user
// @Tags CreditCard
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /api/user/security/credit_card [get]
func userGetCreditCard(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID int32 = 1

		if err := c.Bind(&userID); err != nil {
			return err
		}
		userID = 1
		credit_card, err := pg.Queries.UserGetCreditCard(context.Background(), userID)
		if err != nil {
			return DBResponse(c, err, "failed to get credit card", logger)
		}
		return c.JSON(http.StatusOK, credit_card)
	}
}

// @Summary User Delete Credit Card
// @Description Delete a credit card by its ID
// @Tags CreditCard
// @Accept json
// @Produce json
// @Param  credit_card   body string true "Credit Card"
// @Success 200 {array} db.creditCard
// @Failure 401
// @Router /api/user/security/credit_card/delete [patch]
func userUpdateCreditCard(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID int32 = 1
		var param db.UserUpdateCreditCardParams
		if err := c.Bind(&param); err != nil {
			return err
		}
		param.ID = userID
		credit_card, err := pg.Queries.UserUpdateCreditCard(context.Background(), param)
		if err != nil {
			return DBResponse(c, err, "failed to update credit card", logger)
		}
		return c.JSON(http.StatusOK, credit_card)
	}
}
