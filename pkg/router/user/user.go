package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type updatePasswordParams struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

// @Summary		User Get Info
// @Description	Get user information
// @Tags			User
// @Accept			json
// @Produce		json
// @success		200	{object}	db.UserGetInfoRow
// @Failure		500	{object}	echo.HTTPError
// @Router			/user/info [get]
func GetInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"
		user, err := pg.Queries.UserGetInfo(c.Request().Context(), username)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		user.ImageUrl = mc.GetFileURL(c.Request().Context(), user.ImageUrl)
		return c.JSON(http.StatusOK, user)
	}
}

// @Summary		User Edit Info
// @Description	Edit user information
// @Tags			User
// @Param			name	formData	string	true	"name of coupon"
// @Param			address	formData	string	true	"user address"
// @Param			email	formData	string	true	"email"
// @Param			image	formData	file	true	"image file"
// @Accept			json
// @Produce		json
// @success		200	{object}	db.UserUpdateInfoRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/user/info [patch]
func EditInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.UserUpdateInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		//check email format
		if _, err := mail.ParseAddress(param.Email); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		fileHeader, err := c.FormFile("image")
		//if have file then store in miniopkg/router/seller
		if err == nil {
			imageID, err := mc.PutFile(c.Request().Context(), fileHeader, common.GetEncodeName(fileHeader))
			if err != nil {
				logger.Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			param.ImageID = imageID
		} else if errors.Is(err, http.ErrMissingFile) {
			//use the origin image
			param.ImageID = ""
		} else {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		param.Username = username
		user, err := pg.Queries.UserUpdateInfo(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		user.ImageUrl = mc.GetFileURL(c.Request().Context(), user.ImageUrl)
		return c.JSON(http.StatusOK, user)
	}
}

// @Summary		User Edit Password
// @Description	Change user password
// @Tags			User
// @Param			password	body	updatePasswordParams	true	"password"
// @Accept			json
// @Produce		json
// @success		200	{object}	db.UserUpdatePasswordRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/user/security/password [post]
func EditPassword(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"
		var param updatePasswordParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		//todo check password strong

		currentPassword, err := pg.Queries.UserGetPassword(c.Request().Context(), username)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(param.CurrentPassword)) != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
		}

		hashNewPassword, err := bcrypt.GenerateFromPassword([]byte(param.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		userInfo, err := pg.Queries.UserUpdatePassword(c.Request().Context(), db.UserUpdatePasswordParams{Username: username, NewPassword: string(hashNewPassword)})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, userInfo)
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
func GetCreditCard(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
// @Param			credit_card	body		interface{}	true	"Credit Card"
// @Success		200			{object}	json.RawMessage
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/user/security/credit_card [patch]
func UpdateCreditCard(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"
		var param json.RawMessage
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		credit_cards, err := pg.Queries.UserUpdateCreditCard(c.Request().Context(), db.UserUpdateCreditCardParams{Username: username, CreditCard: param})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, credit_cards)
	}
}
