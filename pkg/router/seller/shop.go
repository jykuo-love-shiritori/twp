package seller

import (
	"errors"
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ReportDetailParam struct {
	Year  int32 `json:"year" param:"year"`
	Month int32 `json:"month" param:"month"`
}
type ReportDetail struct {
	Products []db.SellerBestSellProductRow `json:"products"`
	Report   db.SellerReportRow            `json:"report"`
}

// @Summary		Seller get shop info
// @Description	Get shop info, includes user picture, name, description.
// @Tags			Seller, Shop
// @Produce		json
// @success		200	{object}	db.SellerGetInfoRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/info [get]
func GetShopInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"
		shopInfo, err := pg.Queries.SellerGetInfo(c.Request().Context(), username)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		shopInfo.ImageUrl = mc.GetFileURL(c.Request().Context(), shopInfo.ImageUrl)
		return c.JSON(http.StatusOK, shopInfo)
	}
}

// @Summary		Seller edit shop info
// @Description	Edit shop name, description, visibility.
// @Tags			Seller, Shop
// @Accept			mpfd
// @Param			name		formData	string	true	"update shop name"	minlength(6)
// @Param			image		formData	file	true	"image file"
// @Param			description	formData	string	true	"update description"
// @Param			enabled		formData	bool	true	"update enabled status"
// @Produce		json
// @success		200	{object}	db.SellerUpdateInfoRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/info [patch]
func EditInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerUpdateInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		fileHeader, err := c.FormFile("image")
		if err == nil {
			imageID, err := mc.PutFile(c.Request().Context(), fileHeader, common.GetFileName(fileHeader))
			if err != nil {
				logger.Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			param.ImageID = imageID
		} else if errors.Is(err, http.ErrMissingFile) {
			//use the origin image
			param.ImageID = ""
		} else {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		param.SellerName = username
		shopInfo, err := pg.Queries.SellerUpdateInfo(c.Request().Context(), param)
		if err != nil {
			if param.ImageID != "" {
				err := mc.RemoveFile(c.Request().Context(), param.ImageID)
				if err != nil {
					logger.Error(err)
					return echo.NewHTTPError(http.StatusInternalServerError)
				}
			}
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		shopInfo.ImageUrl = mc.GetFileURL(c.Request().Context(), shopInfo.ImageUrl)

		return c.JSON(http.StatusOK, shopInfo)
	}
}

// @Summary		Seller get report detail
// @Description	Get report detail by year and month for shop.
// @Tags			Seller, Shop, Report
// @Produce		json
// @Param			year	path		int	true	"Year"
// @Param			month	path		int	true	"Month"
// @Success		200		{object}	db.SellerInsertCouponRow
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/seller/report/{year}/{month} [get]
func GetReportDetail(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var username string = "user1"
		var param ReportDetailParam
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		var result ReportDetail
		result.Report, err = pg.Queries.SellerReport(c.Request().Context(), db.SellerReportParams{SellerName: username, Month: param.Month, Year: param.Year})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.Products, err = pg.Queries.SellerBestSellProduct(c.Request().Context(), db.SellerBestSellProductParams{SellerName: username, Month: param.Month, Year: param.Year, Limit: 3})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range result.Products {
			result.Products[i].ImageUrl = mc.GetFileURL(c.Request().Context(), result.Products[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, result)
	}
}