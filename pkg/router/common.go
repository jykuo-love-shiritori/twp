package router

import (
	"net/http"
	"regexp"

	"github.com/jackc/pgx/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
)

type failure struct {
	Msg string `json:"msg"`
}

type OrderDetail struct {
	OrderInfo db.OrderHistory              `json:"order_info"`
	Products  []db.SellerGetOrderDetailRow `json:"products"`
}

func hasSpecialChars(input string) bool {
	regexPattern := `[.*+?()|{}\\^$]`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(input)
}

func mapDBErrorToHTTPError(err error, c echo.Context) error {
	// Customize this function based on the specific errors you want to handle
	switch err {
	case nil:
		return c.JSON(http.StatusOK, failure{"success"})
	case pgx.ErrNoRows:
		return c.JSON(http.StatusInternalServerError, failure{"Not Found"})
	case pgx.ErrTooManyRows:
		return c.JSON(http.StatusInternalServerError, failure{"Too Many Result"})
	default:
		return c.JSON(http.StatusInternalServerError, failure{"Internal Server Error"})
	}
}
