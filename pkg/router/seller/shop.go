package seller

import (
	"errors"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/auth"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/image"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

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
		username, valid := auth.GetUsername(c)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		shopInfo, err := pg.Queries.SellerGetInfo(c.Request().Context(), username)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		shopInfo.ImageUrl = image.GetUrl(shopInfo.ImageUrl)
		return c.JSON(http.StatusOK, shopInfo)
	}
}

// @Summary		Seller edit shop info
// @Description	Edit shop name, description, visibility.
// @Tags			Seller, Shop
// @Accept			mpfd
// @Param			name		formData	string	true	"update shop name"
// @Param			image		formData	file	false	"image file"
// @Param			description	formData	string	true	"update description"
// @Param			enabled		formData	bool	true	"update enabled status"
// @Produce		json
// @success		200	{object}	db.SellerUpdateInfoRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/info [patch]
func EditInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, valid := auth.GetUsername(c)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		var param db.SellerUpdateInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if enabled := c.FormValue("enabled"); enabled == "" {
			logger.Errorw("enabled cant be empty")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if param.Name == "" || param.Description == "" {
			logger.Errorw("columns cant be empty")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		fileHeader, err := c.FormFile("image")
		if err == nil {
			imageID, err := mc.PutFile(c.Request().Context(), fileHeader, common.CreateUniqueFileName(fileHeader.Filename))
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
		shopInfo.ImageUrl = image.GetUrl(shopInfo.ImageUrl)

		return c.JSON(http.StatusOK, shopInfo)
	}
}

// @Summary		Seller get report detail
// @Description	Get report detail by year and month for shop.
// @Tags			Seller, Shop, Report
// @Produce		json
// @Param			time	query		string	true	"time"
// @Success		200		{object}	ReportDetail
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/seller/report [get]
func GetReportDetail(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		username, valid := auth.GetUsername(c)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		inputTime := pgtype.Timestamptz{}
		if err := echo.QueryParamsBinder(c).Time("time", &inputTime.Time, time.RFC3339).BindError(); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		// if use 0 day will be the last day of last month
		inputTime.Time = time.Date(inputTime.Time.Year(), inputTime.Time.Month(), 1, 0, 0, 0, 0, inputTime.Time.Location())
		inputTime.Valid = true
		var result ReportDetail
		result.Report, err = pg.Queries.SellerReport(c.Request().Context(), db.SellerReportParams{SellerName: username, Time: inputTime})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.Products, err = pg.Queries.SellerBestSellProduct(c.Request().Context(), db.SellerBestSellProductParams{SellerName: username, Time: inputTime, Limit: 3})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range result.Products {
			result.Products[i].ImageUrl = image.GetUrl(result.Products[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, result)
	}
}
