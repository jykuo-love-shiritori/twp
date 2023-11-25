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
type productDetail struct {
	ProductInfo db.Product      `json:"product_info"`
	Tags        []db.ProductTag `json:"tags"`
}
type couponDetail struct {
	CouponInfo db.Coupon      `json:"coupon_info"`
	Tags       []db.CouponTag `json:"tags"`
}
type failure struct {
	Error string `json:"error"`
}

func hasSpecialChars(input string) bool {
	regexPattern := `[.*+?()|{}\\^$]`
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(input)
}

func DBResponse(c echo.Context, err error, errorMsg string, logger *zap.SugaredLogger) error {
	switch err {
	case pgx.ErrNoRows:
		logger.Errorw(errorMsg, "error", err)
		return c.JSON(http.StatusNotFound, failure{"Not Found"})
	case pgx.ErrTooManyRows:
		logger.Errorw(errorMsg, "error", err)
		return c.JSON(http.StatusInternalServerError, failure{"Too Many Result"})
	case pgx.ErrTooManyRows, pgx.ErrTxClosed, pgx.ErrTxCommitRollback:
		logger.Errorw(errorMsg, "error", err)
		return c.JSON(http.StatusInternalServerError, failure{"Internal Server Error"})
	default:
		logger.Errorw(errorMsg, "error", err)
		return c.JSON(http.StatusInternalServerError, failure{errorMsg})
	}
}
