package router

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary		User Get Info
// @Description	Get user information
// @Tags			User
// @Accept			json
// @Produce		json
// @success		200	{object}	db.UserGetInfoRow
// @Failure		500	{object}	echo.HTTPError
// @Router			/user/info [get]
func userGetInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"
		user, err := pg.Queries.UserGetInfo(c.Request().Context(), username)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, user)
	}
}

// @Summary		User Edit Info
// @Description	Edit user information
// @Tags			User
// @Param			name		body	string	true	"name of coupon"
// @Param			address		body	string	true	"user address"
// @Param			image_id	body	string	true	"image id"
// @Accept			json
// @Produce		json
// @success		200	{object}	db.UserUpdateInfoRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/user/info [patch]
func userEditInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.UserUpdateInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		param.Username = username
		info, err := pg.Queries.UserUpdateInfo(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, info)
	}
}

// @Summary		User Upload Avatar
// @Description	Upload user avatar
// @Tags			User
// @Accept			png,jpeg,gif
// @Produce		json
// @Param			img	formData	file	true	"Image to upload"
// @Success		200
// @Failure		401
// @Router			/user/avatar [post]
func userUploadAvatar(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary		User Edit Password
// @Description	Change user password
// @Tags			User
// @Param			current_password	body	string	true	"current password"
// @Param			new_password		body	string	true	"new password"
// @Accept			json
// @Produce		json
// @success		200	{object}	db.UserUpdatePasswordRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/user/security/password [post]
func userEditPassword(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"
		var param db.UserUpdatePasswordParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		param.Username = username
		orders, err := pg.Queries.UserUpdatePassword(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, orders)
	}
}

// @Summary		User Get Credit Card
// @Description	Get all credit cards of the user
// @Tags			CreditCard
// @Accept			json
// @Produce		json
// @Success		200	{object}	json.RawMessage
// @Failure		500	{object}	echo.HTTPError
// @Router			/user/security/credit_card [get]
func userGetCreditCard(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		credit_card, err := pg.Queries.UserGetCreditCard(c.Request().Context(), username)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, credit_card)
	}
}

// @Summary		User Delete Credit Card
// @Description	Delete a credit card by its ID
// @Tags			CreditCard
// @Accept			json
// @Produce		json
// @Param			credit_card	body		json.RawMessage	true	"Credit Card"
// @Success		200			{object}	json.RawMessage
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/user/security/credit_card [patch]
func userUpdateCreditCard(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"
		var param db.UserUpdateCreditCardParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		param.Username = username
		credit_cards, err := pg.Queries.UserUpdateCreditCard(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, credit_cards)
	}
}
