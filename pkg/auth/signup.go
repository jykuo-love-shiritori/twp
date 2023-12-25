package auth

import (
	"net/http"
	"net/mail"
	"regexp"
	"unicode"

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

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidUsername(username string) bool {
	if len(username) > 32 {
		return false
	}

	r := regexp.MustCompile(`^[A-Za-z0-9]+$`)

	return r.MatchString(username)
}

func isValidPassword(password string) bool {
	// bcrypt only allows passwords with length up to 72 bytes
	if len(password) < 8 || len([]byte(password)) >= 72 {
		return false
	}

	// golang `regexp` doesn't support backtracking
	hasLower := false
	hasUpper := false
	hasNumber := false
	hasSpecial := false

	for _, v := range password {
		switch {
		case unicode.IsLower(v):
			hasLower = true
		case unicode.IsUpper(v):
			hasUpper = true
		case unicode.IsNumber(v):
			hasNumber = true
		case unicode.IsPunct(v) || unicode.IsSymbol(v):
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasNumber && hasSpecial
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

		userExists, err := pg.Queries.UserExists(ctx, db.UserExistsParams{
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

		if !isValidEmail(params.Email) {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid email")
		}
		if !isValidUsername(params.Username) {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid username")
		}
		if !isValidPassword(params.Password) {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid password")
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), 14)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		tx, err := pg.NewTx(ctx)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		defer tx.Rollback(ctx) //nolint:errcheck

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
