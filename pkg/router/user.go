package router

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

//	@Summary		Logout
//	@Description	Logout the current user
//	@Tags			Auth
//	@Produce		json
//	@Success		200
//	@Failure		401
//	@Router			/logout [post]
func logout(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

//	@Summary		User Get Info
//	@Description	Get user information
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		401
//	@Router			/user [get]
func userGetInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

//	@Summary		User Edit Info
//	@Description	Edit user information
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		401
//	@Router			/user/edit [patch]
func userEditInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

//	@Summary		User Upload Avatar
//	@Description	Upload user avatar
//	@Tags			User
//	@Accept			png,jpeg,gif
//	@Produce		json
//	@Param			img	formData	file	true	"Image to upload"
//	@Success		200
//	@Failure		401
//	@Router			/user/avatar [post]
func userUploadAvatar(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

//	@Summary		User Edit Password
//	@Description	Change user password
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		401
//	@Router			/user/password [post]
func userEditPassword(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

//	@Summary		User Get Credit Card
//	@Description	Get all credit cards of the user
//	@Tags			CreditCard
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		401
//	@Router			/user/security/credit_card [get]
func userGetCreditCard(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

//	@Summary		User Delete Credit Card
//	@Description	Delete a credit card by its ID
//	@Tags			CreditCard
//	@Accept			json
//	@Produce		json
//	@Param			id	query	int	true	"Credit Card ID"
//	@Success		200
//	@Failure		401
//	@Router			/user/security/credit_card/delete [delete]
func userDeleteCreditCard(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

//	@Summary		User Add Credit Card
//	@Description	Add a new credit card for the user
//	@Tags			CreditCard
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		401
//	@Router			/user/security/credit_card/add [post]
func userAddCreditCard(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}
