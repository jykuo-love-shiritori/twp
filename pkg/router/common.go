package router

import (
	"net/http"
	"regexp"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type failure struct {
	Fail string `json:"fail"`
}

func hasSpecialChars(input string) bool {
	regexPattern := `[.*+?()|{}\\^$]`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(input)
}

func mapDBErrorToHTTPError(err error, c echo.Context) error {
	// Customize this function based on the specific errors you want to handle
	switch err {
	case pgx.ErrNoRows:
		return c.JSON(http.StatusOK, failure{"Not Found"})
	case pgx.ErrTooManyRows:
		return c.JSON(http.StatusInternalServerError, failure{"Too Many result"})
	default:
		return c.JSON(http.StatusInternalServerError, failure{"Internal Server Error"})
	}
}
