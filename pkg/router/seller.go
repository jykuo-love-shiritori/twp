package router

import (
	"errors"
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type orderDetail struct {
	OrderInfo db.SellerGetOrderHistoryRow  `json:"order_info"`
	Products  []db.SellerGetOrderDetailRow `json:"products"`
}
type productDetail struct {
	ProductInfo db.SellerGetProductDetailRow `json:"product_info"`
	Tags        []db.SellerGetProductTagRow  `json:"tags"`
}
type couponDetail struct {
	CouponInfo db.SellerGetCouponDetailRow `json:"coupon_info"`
	Tags       []db.SellerGetCouponTagRow  `json:"tags"`
}

type reportDetailParam struct {
	Year  int32 `json:"year" param:"year"`
	Month int32 `json:"month" param:"month"`
}
type reportDetail struct {
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
func sellerGetShopInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"
		shopInfo, err := pg.Queries.SellerGetInfo(c.Request().Context(), username)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		shopInfo.ImageID = mc.GetFileURL(c.Request().Context(), shopInfo.ImageID)
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
func sellerEditInfo(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
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
		shopInfo.ImageID = mc.GetFileURL(c.Request().Context(), shopInfo.ImageID)

		return c.JSON(http.StatusOK, shopInfo)
	}
}

// @Summary		Seller get available tag
// @Description	Get all available tags for shop.
// @Tags			Seller, Shop, Tag
// @Param			name	body	string	true	"search tag name start with"	minlength(1)
// @Produce		json
// @success		200	{array}		db.SellerSearchTagRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/tag [get]
func sellerGetTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"
		var tagPerPage int32 = 20

		var param db.SellerSearchTagParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		if param.Name == "" || common.HasRegexSpecialChars(param.Name) {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.SellerName = username
		param.Name = "^" + param.Name
		param.Limit = tagPerPage
		tags, err := pg.Queries.SellerSearchTag(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, tags)
	}
}

// @Summary		Seller add tag
// @Description	Add tag for shop.
// @Tags			Seller, Shop, Tag
// @Accept			json
// @Param			name	body	string	true	"insert tag"	minlength(1)
// @Produce		json
// @success		200	{object}	db.SellerInsertTagRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		409	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/tag [post]
func sellerAddTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.HaveTagNameParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		if param.Name == "" || common.HasRegexSpecialChars(param.Name) {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.SellerName = username
		have, err := pg.Queries.HaveTagName(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if have {
			return echo.NewHTTPError(http.StatusConflict)
		}
		tag, err := pg.Queries.SellerInsertTag(c.Request().Context(), db.SellerInsertTagParams(param))
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, tag)
	}
}

// @Summary		Seller get shop coupon
// @Description	Get all coupons for shop.
// @Tags			Seller, Shop, Coupon
// @Param			offset	query	int	true	"offset page"	default(0)	minimum(0)
// @Param			limit	query	int	true	"limit"			default(10)	minimum(20)
// @Produce		json
// @success		200	{array}		db.SellerGetCouponRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/coupon [get]
func sellerGetShopCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

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
// @success		200	{object}	couponDetail
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/coupon/{id} [get]
func sellerGetCouponDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		var username string = "user1"

		var param db.SellerGetCouponDetailParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		var result couponDetail
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

// todo insert tag
//
//	@Summary		Seller add coupon
//	@Description	Add coupon for shop.
//	@Tags			Seller, Shop, Coupon
//	@Param			type		body	string	true	"Coupon type"	Enums('percentage', 'fixed', 'shipping')
//	@Param			name		body	string	true	"name of coupon"
//	@Param			description	body	string	true	"description of coupon"
//	@Param			discount	body	number	false	"discount"
//	@Param			start_date	body	time	true	"start date"
//	@Param			expire_date	body	time	true	"expire date"
//	@Param			tags		body	[]int32	true	"init tags"
//	@Produce		json
//	@success		200	{object}	db.SellerInsertCouponRow
//	@Failure		400	{object}	echo.HTTPError
//	@Failure		500	{object}	echo.HTTPError
//	@Router			/seller/coupon [post]
func sellerAddCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerInsertCouponParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.SellerName = username
		coupon, err := pg.Queries.SellerInsertCoupon(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		// coupon, err := pg.Queries.SellerInsertCouponTag(c.Request().Context(), queryParams)
		// if err != nil {
		// 	logger.Error(err)
		// 	return echo.NewHTTPError(http.StatusInternalServerError)
		// // }
		return c.JSON(http.StatusOK, coupon)
	}
}

// @Summary		Seller add coupon tag
// @Description	Add tag on coupon
// @Tags			Seller, Shop, Coupon,Tag
// @Accept			json
// @Param			id		path	string	true	"coupon id"
// @Param			tag_id	body	int		true	"add tag id"
// @Produce		json
// @success		200	{object}	db.CouponTag
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/coupon/{id}/tag [post]
func sellerAddCouponTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

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

// @Summary		Seller edit coupon
// @Description	Edit coupon for shop.
// @Tags			Seller, Shop, Coupon
// @Accept			json
// @Produce		json
// @Param			coupon_id	path		int		true	"Coupon ID"
// @Param			type		body		string	true	"Coupon type"	Enums('percentage', 'fixed', 'shipping')
// @Param			name		body		string	true	"name of coupon"
// @Param			description	body		string	true	"description of coupon"
// @Param			discount	body		number	false	"discount"
// @Param			start_date	body		time	true	"start date"
// @Param			expire_date	body		time	true	"expire date"
// @success		200			{object}	db.SellerUpdateCouponInfoRow
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/seller/coupon/{coupon_id} [patch]
func sellerEditCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerUpdateCouponInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
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
func sellerDeleteCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

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
			logger.Error(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

// @Summary		Seller delete coupon tag
// @Description	Delete coupon for shop.
// @Tags			Seller, Shop, Coupon,Tag
// @Param			id		path	string	true	"coupon id"
// @Param			tag_id	body	int		true	"add tag id"
// @Accept			json
// @Produce		json
// @Success		200	{string}	string	"success"
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/coupon/{id}/tag [delete]
func sellerDeleteCouponTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

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
			logger.Error(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}

// @Summary		Seller get order
// @Description	Get all orders for shop.
// @Tags			Seller, Shop, Order
// @Param			offset	query	int	true	"offset"	default(0)	minimum(0)
// @Param			limit	query	int	true	"limit"		default(10)	maximum(20)
// @Produce		json
// @Success		200	{array}		db.SellerGetOrderRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/order [get]
func sellerGetOrder(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		var username string = "user1"

		var requestParam common.QueryParams
		if err := c.Bind(&requestParam); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		if requestParam.Validate() != nil {
			logger.Errorw("invalid query parameter", "offset", requestParam.Offset, "limit", requestParam.Limit)
			return echo.NewHTTPError(http.StatusBadRequest, "offset or limit is invalid")
		}
		param := db.SellerGetOrderParams{SellerName: username, Limit: requestParam.Limit, Offset: requestParam.Offset}

		orders, err := pg.Queries.SellerGetOrder(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range orders {
			orders[i].ImageID = mc.GetFileURL(c.Request().Context(), orders[i].ImageID)
		}
		return c.JSON(http.StatusOK, orders)
	}
}

// @Summary		Seller get order detail
// @Description	Get order detail by ID for shop.
// @Tags			Seller, Shop, Order
// @Produce		json
// @Param			id	path		int	true	"Order ID"
// @Success		200	{object}	orderDetail
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/order/{id} [get]
func sellerGetOrderDetail(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerGetOrderDetailParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		param.SellerName = username
		var result orderDetail
		var err error
		result.OrderInfo, err = pg.Queries.SellerGetOrderHistory(c.Request().Context(), db.SellerGetOrderHistoryParams{SellerName: param.SellerName, ID: param.OrderID})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.Products, err = pg.Queries.SellerGetOrderDetail(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range result.Products {
			result.Products[i].ImageID = mc.GetFileURL(c.Request().Context(), result.OrderInfo.ImageID)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Seller update order status
// @Description	seller update orders status.
// @Tags			Seller, Shop, Order
// @Param			id				path		int		true	"Order ID"
// @Param			current_status	body		string	true	"order status"	Enums('pending','paid','shipped','delivered','cancelled')
// @Param			set_status		body		string	true	"order status"	Enums('pending','paid','shipped','delivered','cancelled')
// @Success		200				{object}	db.SellerUpdateOrderStatusRow
// @Failure		400				{object}	echo.HTTPError
// @Failure		500				{object}	echo.HTTPError
// @Router			/seller/order [patch]
func sellerUpdateOrderStatus(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerUpdateOrderStatusParams
		if err := c.Bind(&param); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		param.SellerName = username

		// shop can only a prove the status traction {paid > shipped ,shipped > delivered}
		// paid > shipped > delivered > (cancelled || finished)
		if !((param.CurrentStatus == "paid" && param.SetStatus == "shipped") || (param.CurrentStatus == "shipped" && param.SetStatus == "delivered")) {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		order, err := pg.Queries.SellerUpdateOrderStatus(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, order)
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
func sellerGetReportDetail(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var username string = "user1"
		var param reportDetailParam
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		var result reportDetail
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
			result.Products[i].ImageID = mc.GetFileURL(c.Request().Context(), result.Products[i].ImageID)
		}
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Seller get product
// @Description	Delete product for shop.
// @Tags			Seller, Shop, Product
// @Accept			json
// @Produce		json
// @Param			product_id	path		int	true	"Product ID"
// @Success		200			{object}	productDetail
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/seller/product/{product_id} [get]
func sellerGetProductDetail(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var username string = "user1"
		var param db.SellerGetProductDetailParams

		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.SellerName = username

		var result productDetail
		result.ProductInfo, err = pg.Queries.SellerGetProductDetail(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.Tags, err = pg.Queries.SellerGetProductTag(c.Request().Context(), db.SellerGetProductTagParams{SellerName: username, ProductID: param.ID})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		result.ProductInfo.ImageID = mc.GetFileURL(c.Request().Context(), result.ProductInfo.ImageID)
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Seller get product
// @Description	seller get product
// @Tags			Seller, Shop, Product
// @Param			offset	query	int	true	"offset page"	default(0)	minimum(0)
// @Param			limit	query	int	true	"limit"			default(10)	maximum(20)
// @Accept			json
// @Produce		json
// @Success		200	{array}		db.SellerProductListRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/product [get]
func sellerListProduct(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var requestParam common.QueryParams
		if err := c.Bind(&requestParam); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if requestParam.Validate() != nil {
			logger.Errorw("invalid query parameter", "offset", requestParam.Offset, "limit", requestParam.Limit)
			return echo.NewHTTPError(http.StatusBadRequest, "offset or limit is invalid")
		}
		logger.Info(requestParam)
		param := db.SellerProductListParams{SellerName: username, Limit: requestParam.Limit, Offset: requestParam.Offset}
		productsRow, err := pg.Queries.SellerProductList(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range productsRow {
			productsRow[i].ImageID = mc.GetFileURL(c.Request().Context(), productsRow[i].ImageID)
		}
		return c.JSON(http.StatusOK, productsRow)
	}
}

// todo insert tag
// @Summary		Seller add product
// @Description	Add product for shop.
// @Tags			Seller, Shop, Product
// @Param			name		formData	string	true	"name of product"
// @Param			description	formData	string	true	"description of product"
// @Param			price		formData	number	false	"price"
// @Param			image		formData	file	true	"image id"
// @Param			expire_date	formData	string	true	"expire date"
// @Param			stock		formData	int		true	"stock"
// @Param			enabled		formData	string	true	"enabled"
// @Param			tags		formData	[]int32	true	"init tags"
// @Accept			mpfd
// @Produce		json
// @Success		200	{object}	db.SellerInsertProductRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/product [post]
func sellerAddProduct(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerInsertProductParams

		logger.Info()

		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := param.Price.Scan(c.FormValue("price")); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := param.ExpireDate.Scan(c.FormValue("expire_date")); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		fileHeader, err := c.FormFile("image")
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		//put file to minio
		ImageID, err := mc.PutFile(c.Request().Context(), fileHeader, common.GetFileName(fileHeader))
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		param.ImageID = ImageID
		param.SellerName = username
		product, err := pg.Queries.SellerInsertProduct(c.Request().Context(), param)
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

		product.ImageID = mc.GetFileURL(c.Request().Context(), product.ImageID)
		return c.JSON(http.StatusOK, product)
	}
}

// @Summary		Seller edit product
// @Description	Edit product for shop.
// @Tags			Seller, Shop, Product
// @Accept			mpfd
// @Produce		json
// @Param			id			path		int		true	"Product ID"
// @Param			name		formData	string	true	"name of product"
// @Param			description	formData	string	true	"description of product"
// @Param			price		formData	number	false	"price"
// @Param			image_id	formData	string	true	"image id"
// @Param			expire_date	formData	time	true	"expire date"
// @Param			stock		formData	int		true	"stock"
// @Param			enabled		formData	time	true	"enabled"
// @Success		200			{object}	db.SellerUpdateProductInfoRow
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/seller/product/{id} [patch]
func sellerEditProduct(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerUpdateProductInfoParams
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
		if err := param.Price.Scan(c.FormValue("price")); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		// param.Price = pgtype.Numeric{}
		product, err := pg.Queries.SellerUpdateProductInfo(c.Request().Context(), param)
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
		product.ImageID = mc.GetFileURL(c.Request().Context(), product.ImageID)
		return c.JSON(http.StatusOK, product)
	}
}

// @Summary		Seller add product tag
// @Description	Add tag on product
// @Tags			Seller, Shop, Product,Tag
// @Accept			json
// @Param			id		path	string	true	"product id"
// @Param			tag_id	body	int		true	"add tag id"
// @Produce		json
// @Success		200	{object}	db.ProductTag
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/product/{id}/tag [post]
func sellerAddProductTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerInsertProductTagParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.SellerName = username
		productTag, err := pg.Queries.SellerInsertProductTag(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, productTag)
	}
}

// @Summary		Seller delete product
// @Description	Delete product for shop.
// @Tags			Seller, Shop, Product
// @Accept			json
// @Produce		json
// @Param			id	path		int		true	"Product ID"
// @Success		200	{string}	string	"success"
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/product/{id} [delete]
func sellerDeleteProduct(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerDeleteProductParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		param.SellerName = username
		effectRow, err := pg.Queries.SellerDeleteProduct(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if effectRow == 0 {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)

	}
}

// @Summary		Seller delete product tag
// @Description	Delete product for shop.
// @Tags			Seller, Shop, Coupon
// @Param			id		path	int	true	"product id"
// @Param			tag_id	body	int	true	"add tag id"
// @Accept			json
// @Produce		json
// @Success		200	{string}	string	"success"
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/product/{id}/tag [delete]
func sellerDeleteProductTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerDeleteProductTagParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)

		}
		param.SellerName = username
		effectRow, err := pg.Queries.SellerDeleteProductTag(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if effectRow == 0 {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}
