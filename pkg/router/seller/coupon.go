package seller

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/pkg/auth"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CouponDetail struct {
	CouponInfo db.SellerGetCouponDetailRow `json:"coupon_info"`
	Tags       []db.SellerGetCouponTagRow  `json:"tags"`
}

type InsertCouponParams struct {
	Type        db.CouponType      `json:"type" example:"fixed"`
	Name        string             `json:"name" example:"product name"`
	Description string             `json:"description" example:"some description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number" example:"19.99"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string" example:"2024-10-12T07:20:50.52Z"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string" example:"2024-11-12T07:20:50.52Z"`
	Tags        []int32            `json:"tags" example:"10001,10002"`
}

// @Summary		Seller get shop coupon
// @Description	Get all coupons for shop.
// @Tags			Seller, Shop, Coupon
// @Param			offset	query	int	true	"offset"	default(0)	minimum(0)
// @Param			limit	query	int	true	"limit"		default(10)	maximum(20)
// @Produce		json
// @success		200	{array}		db.SellerGetCouponRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/coupon [get]
func GetShopCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, valid := auth.GetUsername(c)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		var requestParam common.QueryParams
		if err := c.Bind(&requestParam); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		if requestParam.Validate() != nil {
			logger.Errorw("invalid query parameter", "offset", requestParam.Offset, "limit", requestParam.Limit)
			return echo.NewHTTPError(http.StatusBadRequest, "offset or limit is invalid")
		}

		param := db.SellerGetCouponParams{SellerName: username, Limit: requestParam.Limit, Offset: requestParam.Offset}
		coupons, err := pg.Queries.SellerGetCoupon(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, coupons)
	}
}

// @Summary		Seller get coupon detail
// @Description	Get coupon detail by ID for shop.
// @Tags			Seller, Shop, Coupon
// @Produce		json
// @Param			id	path		int	true	"Coupon ID"
// @success		200	{object}	CouponDetail
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/coupon/{id} [get]
func GetCouponDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		username, valid := auth.GetUsername(c)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		var param db.SellerGetCouponDetailParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		var result CouponDetail
		var err error
		param.SellerName = username
		result.CouponInfo, err = pg.Queries.SellerGetCouponDetail(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.Tags, err = pg.Queries.SellerGetCouponTag(c.Request().Context(), db.SellerGetCouponTagParams{SellerName: param.SellerName, CouponID: param.ID})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Seller add coupon
// @Description	Add coupon for shop.
// @Tags			Seller, Shop, Coupon
// @Param			coupon	body	InsertCouponParams	true	"coupon"
// @Produce		json
// @success		200	{object}	db.SellerInsertCouponRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/coupon [post]
func AddCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, valid := auth.GetUsername(c)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		var param InsertCouponParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		valid, err := pg.Queries.SellerCheckTags(c.Request().Context(), db.SellerCheckTagsParams{SellerName: username, Tags: param.Tags})
		if err != nil {
			logger.Error(valid, err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if !valid {
			logger.Errorw("tags is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		//check start/expire time
		if param.StartDate.Time.Before(time.Now()) {
			param.StartDate.Time = time.Now()
		}
		if param.ExpireDate.Time.Before(param.StartDate.Time) {
			logger.Errorw("expire date is invalid", "start date", param.StartDate)
			return echo.NewHTTPError(http.StatusBadRequest, "expire date is invalid")
		}
		//check discount value
		if v, err := param.Discount.Float64Value(); err != nil || v.Float64 < 0 || (param.Type == db.CouponTypePercentage && v.Float64 > 100) {
			logger.Errorw("discount is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		coupon, err := pg.Queries.SellerInsertCoupon(c.Request().Context(), db.SellerInsertCouponParams{
			SellerName:  username,
			Type:        param.Type,
			Name:        param.Name,
			Description: param.Description,
			Discount:    param.Discount,
			StartDate:   param.StartDate,
			ExpireDate:  param.ExpireDate,
		})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		err = pg.Queries.SellerInsertCouponTags(c.Request().Context(), db.SellerInsertCouponTagsParams{CouponID: coupon.ID, Tags: param.Tags})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, coupon)
	}
}

// @Summary		Seller edit coupon
// @Description	Edit coupon for shop.
// @Tags			Seller, Shop, Coupon
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Coupon ID"
// @Param			coupon	body		InsertCouponParams	true	"coupon"
// @success		200		{object}	db.SellerUpdateCouponInfoRow
// @Failure		400		{object}	echo.HTTPError
// @Failure		500		{object}	echo.HTTPError
// @Router			/seller/coupon/{id} [patch]
func EditCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, valid := auth.GetUsername(c)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		var param db.SellerUpdateCouponInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		//check start/expire time
		if param.StartDate.Time.Before(time.Now()) {
			param.StartDate.Time = time.Now()
		}
		if param.ExpireDate.Time.Before(param.StartDate.Time) {
			logger.Errorw("expire date is invalid", "start date", param.StartDate)
			return echo.NewHTTPError(http.StatusBadRequest, "expire date is invalid")
		}
		//check discount value
		if v, err := param.Discount.Float64Value(); err != nil || v.Float64 < 0 || (param.Type == db.CouponTypePercentage && v.Float64 > 100) {
			logger.Errorw("discount is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.SellerName = username
		coupon, err := pg.Queries.SellerUpdateCouponInfo(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, coupon)
	}
}

// @Summary		Seller delete coupon
// @Description	Delete coupon for shop.
// @Tags			Seller, Shop, Coupon
// @Param			id	path	int	true	"Coupon ID"
// @Accept			json
// @Produce		json
// @Success		200	{string}	string	"success"
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/coupon/{id} [delete]
func DeleteCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, valid := auth.GetUsername(c)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		var param db.SellerDeleteCouponParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		param.SellerName = username
		effectRow, err := pg.Queries.SellerDeleteCoupon(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if effectRow == 0 {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

// @Summary		Seller add coupon tag
// @Description	Add tag on coupon
// @Tags			Seller, Shop, Coupon,Tag
// @Accept			json
// @Param			id		path	int				true	"coupon id"
// @Param			tag_id	body	GetTagParams	true	"add tag id"
// @Produce		json
// @success		200	{object}	db.CouponTag
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/coupon/{id}/tag [post]
func AddCouponTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, valid := auth.GetUsername(c)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		var param db.SellerInsertCouponTagParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.SellerName = username
		couponTag, err := pg.Queries.SellerInsertCouponTag(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, couponTag)
	}
}

// @Summary		Seller delete coupon tag
// @Description	Delete coupon for shop.
// @Tags			Seller, Shop, Coupon,Tag
// @Param			id		path	int				true	"coupon id"
// @Param			tag_id	body	GetTagParams	true	"add tag id"
// @Accept			json
// @Produce		json
// @Success		200	{string}	string	"success"
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/coupon/{id}/tag [delete]
func DeleteCouponTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, valid := auth.GetUsername(c)
		if !valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		var param db.SellerDeleteCouponTagParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		param.SellerName = username
		effectRow, err := pg.Queries.SellerDeleteCouponTag(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if effectRow == 0 {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}
