package auth

import (
	"net/http"
	"net/mail"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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
		ctx := c.Request().Context()
		params := signupParams{}

		err = c.Bind(&params)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		tx, err := pg.NewTx(ctx)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		defer tx.Rollback(ctx) //nolint:errcheck

		userExists, err := pg.Queries.WithTx(tx).UserExists(ctx, db.UserExistsParams{
			Username: params.Username,
			Email:    params.Email,
		})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if userExists {
			return echo.NewHTTPError(http.StatusBadRequest, "Username or email already exists")
		}

		_, err = mail.ParseAddress(params.Email)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid email")
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), 14)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		err = pg.Queries.WithTx(tx).AddUser(ctx, db.AddUserParams{
			Username: params.Username,
			Password: string(hash),
			Name:     params.Name,
			Email:    params.Email,
		})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		err = pg.Queries.WithTx(tx).AddShop(ctx, db.AddShopParams{
			SellerName: params.Username,
			Name:       params.Name,
		})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		err = tx.Commit(ctx)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}
