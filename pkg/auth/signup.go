package auth

import (
	"net/http"
	"net/mail"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type signupParams struct {
	Username string `json:"username" example:"john"`
	Password string `json:"password" example:"secretp@ssword123"`
	Name     string `json:"name" example:"John Doe"`
	Email    string `json:"email" example:"test@gmail.com"`
}

// Signup
//
//	@Summary		Customer signup
//	@Description	signup
//	@Tags			User
//	@Produce		json
//	@Param			request	body	signupParams	true	"something"
//	@Success		200
//	@Failure		400	{object}	echo.HTTPError
//	@Failure		500	{object}	echo.HTTPError
//	@Router			/signup [post]
func Signup(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		params := signupParams{}

		err = c.Bind(&params)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		userExists, err := pg.Queries.UserExists(c.Request().Context(), db.UserExistsParams{
			Username: params.Username,
			Email:    params.Email,
		})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
		}
		if userExists {
			return echo.NewHTTPError(http.StatusBadRequest, "Username or email already exists")
		}

		_, err = mail.ParseAddress(params.Email)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid email")
		}

		hash, err := hashPassword(params.Password)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
		}

		err = pg.Queries.AddUser(c.Request().Context(), db.AddUserParams{
			Username: params.Username,
			Password: hash,
			Name:     params.Name,
			Email:    params.Email,
			ImageID:  common.DefaultImageUuid,
		})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
		}

		return c.NoContent(http.StatusOK)
	}
}