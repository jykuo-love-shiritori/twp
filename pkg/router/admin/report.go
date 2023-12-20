package admin

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type adminReport struct {
	Year        int32                `json:"year"`
	Month       int32                `json:"month"`
	TotalAmount int32                `json:"total"`
	Sellers     []db.GetTopSellerRow `json:"sellers"`
}

// @Summary		Admin Get Site Report
// @Description	Get site report (top 3 sellers and total amount).
// @Tags			Admin, Report
// @Produce		json
// @Param			date	query		string	true	"Start year/month"
// @Success		200		{object}	adminReport
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/admin/report [get]
func GetReport(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		date := c.QueryParam("date")
		// verify the date is any start day of the month
		t, err := time.Parse(time.RFC3339, date)
		if err != nil {
			logger.Errorw("failed to parse date", "error", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if t.Day() != 1 {
			logger.Errorw("date is not the first day of the month", "date", date)
			return echo.NewHTTPError(http.StatusBadRequest, "Date is not the first day of the month")
		}
		year, month, day := t.Date()
		result := adminReport{Year: int32(year), Month: int32(month)}
		// set date's time to 00:00:00
		formattedDate := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
		result.Sellers, err = pg.Queries.GetTopSeller(c.Request().Context(), pgtype.Timestamptz{
			Time:  formattedDate,
			Valid: true,
		})
		for i := range result.Sellers {
			if result.Sellers[i].ImageUrl != "" {
				result.Sellers[i].ImageUrl = mc.GetFileURL(c.Request().Context(), result.Sellers[i].ImageUrl)
			}
		}
		if err != nil {
			logger.Errorw("failed to get top seller", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		// get total amount by sum the total_sales of sellers
		result.TotalAmount, err = pg.Queries.GetTotalSales(c.Request().Context(), pgtype.Timestamptz{
			Time:  formattedDate,
			Valid: true,
		})
		if err != nil {
			logger.Errorw("failed to get total sales", "error", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, result)
	}
}
