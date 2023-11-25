package user

import (
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func Signup(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		params := db.AddUserParams{}

		err = c.Bind(&params)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), 14)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
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

		err = pg.Queries.AddUser(c.Request().Context(), db.AddUserParams{
			Username: params.Username,
			Password: string(hash),
			Name:     params.Name,
			Email:    params.Email,
			ImageID:  userPlaceholderImageUuid,
		})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
		}

		return c.NoContent(http.StatusOK)
	}
}
