package router

import (
	"net/http"
	"regexp"

	"github.com/jackc/pgx/v5"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type orderDetail struct {
	OrderInfo db.OrderHistory              `json:"order_info"`
	Products  []db.SellerGetOrderDetailRow `json:"products"`
}
type failure struct {
	Error string `json:"error"`
}

func hasSpecialChars(input string) bool {
	regexPattern := `[.*+?()|{}\\^$]`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(input)
}

func DBResponse(c echo.Context, err error, logger *zap.SugaredLogger) error {
	// Customize this function based on the specific errors you want to handle
	if httpErr, ok := err.(*echo.HTTPError); ok {
		switch httpErr.Code {
		case http.StatusBadRequest:
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, failure{"Bad Request (parameters)"})
		default:
			return err
		}
	}

	switch err {
	case nil:
		return c.JSON(http.StatusOK, failure{"success"})
	case pgx.ErrNoRows:
		logger.Error(err)
		return c.JSON(http.StatusNotFound, failure{"Not Found"})
	case pgx.ErrTooManyRows:
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, failure{"Too Many Result"})
	case pgx.ErrTooManyRows, pgx.ErrTxClosed, pgx.ErrTxCommitRollback:
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, failure{"Internal Server Error"})
	default:
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, failure{"Internal Server Error"})
	}
}
