package router

import (
	"net/http"
	"regexp"

	"github.com/jackc/pgx/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type failure struct {
	Error string `json:"error"`
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

func DBResponse(c echo.Context, err error, logger *zap.SugaredLogger) error {
	// Customize this function based on the specific errors you want to handle
	if httpErr, ok := err.(*echo.HTTPError); ok {
		logger.Errorw("Error binding request parameters", "error", err)
		return c.JSON(httpErr.Code, failure{httpErr.Message.(string)})
	}

	switch err {
	case nil:
		return c.JSON(http.StatusOK, failure{"success"})
	case pgx.ErrNoRows:
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, failure{"Not Found"})
	case pgx.ErrTooManyRows:
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, failure{"Too Many Result"})
	case pgx.ErrTxClosed:
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, failure{"Internal Server Error"})
	case pgx.ErrTxCommitRollback:
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, failure{"Internal Server Error"})
	default:
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, failure{"Internal Server Error"})
	}
}
