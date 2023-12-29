package seller

import (
	"errors"
	"net/http"
	"time"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/common"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ProductDetail struct {
	ProductInfo db.SellerGetProductDetailRow `json:"product_info"`
	Tags        []db.SellerGetProductTagRow  `json:"tags"`
}

// @Summary		Seller get product
// @Description	Delete product for shop.
// @Tags			Seller, Shop, Product
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Product ID"
// @Success		200	{object}	ProductDetail
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/product/{id} [get]
func GetProductDetail(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error
		var username string = "user1"
		var param db.SellerGetProductDetailParams

		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.SellerName = username

		var result ProductDetail
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
		result.ProductInfo.ImageUrl = mc.GetFileURL(c.Request().Context(), result.ProductInfo.ImageUrl)
		return c.JSON(http.StatusOK, result)
	}
}

// @Summary		Seller get product
// @Description	seller get product
// @Tags			Seller, Shop, Product
// @Param			offset	query	int	true	"offset"	default(0)	minimum(0)
// @Param			limit	query	int	true	"limit"		default(10)	maximum(20)
// @Accept			json
// @Produce		json
// @Success		200	{array}		db.SellerProductListRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/product [get]
func ListProduct(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
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
		param := db.SellerProductListParams{SellerName: username, Limit: requestParam.Limit, Offset: requestParam.Offset}
		productsRow, err := pg.Queries.SellerProductList(c.Request().Context(), param)
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		for i := range productsRow {
			productsRow[i].ImageUrl = mc.GetFileURL(c.Request().Context(), productsRow[i].ImageUrl)
		}
		return c.JSON(http.StatusOK, productsRow)
	}
}

// @Summary		Seller add product
// @Description	Add product for shop.
// @Tags			Seller, Shop, Product
// @Param			name		formData	string	true	"name of product"			default(A)
// @Param			description	formData	string	true	"description of product"	default(description)
// @Param			price		formData	number	true	"price"						default(19.99)
// @Param			image		formData	file	true	"image file"
// @Param			expire_date	formData	string	true	"expire date"	default(2024-10-12T07:20:50.52Z)
// @Param			stock		formData	int		true	"stock"			default(10)
// @Param			enabled		formData	bool	true	"enabled"		default(true)
// @Param			tags		formData	[]int32	false	"init tags"
// @Accept			mpfd
// @Produce		json
// @Success		200	{object}	db.SellerInsertProductRow
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/product [post]
func AddProduct(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerInsertProductParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		// fail to bind the pgtype.decimal have to manually load with scan function
		if err := param.Price.Scan(c.FormValue("price")); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := echo.FormFieldBinder(c).Time("expire_date", &param.ExpireDate.Time, time.RFC3339).BindError(); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.ExpireDate.Valid = true
		//check expire time
		if param.ExpireDate.Time.Before(time.Now()) {
			logger.Errorw("expire date is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		//check Price range
		if v, err := param.Price.Float64Value(); err != nil || v.Float64 < 0 {
			logger.Errorw("price is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		//check stock
		if param.Stock < 0 {
			logger.Errorw("stock is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		// formData is string so have to manually convert sting to int array
		tags := []int32{}
		if err := echo.FormFieldBinder(c).BindWithDelimiter("tags", &tags, ",").BindError(); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		valid, err := pg.Queries.SellerCheckTags(c.Request().Context(), db.SellerCheckTagsParams{SellerName: username, Tags: tags})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		if !valid {
			logger.Errorw("tags is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		//add product must have image
		fileHeader, err := c.FormFile("image")
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		//put file to minio
		ImageID, err := mc.PutFile(c.Request().Context(), fileHeader, common.CreateUniqueFileName(fileHeader))
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
		err = pg.Queries.SellerInsertProductTags(c.Request().Context(), db.SellerInsertProductTagsParams{ProductID: product.ID, Tags: tags})
		if err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		product.ImageUrl = mc.GetFileURL(c.Request().Context(), product.ImageUrl)
		return c.JSON(http.StatusOK, product)
	}
}

// @Summary		Seller edit product
// @Description	Edit product for shop.
// @Tags			Seller, Shop, Product
// @Accept			mpfd
// @Produce		json
// @Param			id			path		int		true	"Product ID"				default(10001)
// @Param			name		formData	string	true	"name of product"			default(product new 10001)
// @Param			description	formData	string	true	"description of product"	default(description)
// @Param			price		formData	number	true	"price"						default(19.99)
// @Param			image		formData	file	false	"image file"
// @Param			expire_date	formData	string	true	"expire date"	default(2024-10-12T07:20:50.52Z)
// @Param			stock		formData	int		true	"stock"			default(10)
// @Param			enabled		formData	bool	false	"enabled"		default(true)
// @Success		200			{object}	db.SellerUpdateProductInfoRow
// @Failure		400			{object}	echo.HTTPError
// @Failure		500			{object}	echo.HTTPError
// @Router			/seller/product/{id} [patch]
func EditProduct(pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user1"

		var param db.SellerUpdateProductInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := param.Price.Scan(c.FormValue("price")); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := echo.FormFieldBinder(c).Time("expire_date", &param.ExpireDate.Time, time.RFC3339).BindError(); err != nil {
			logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		param.ExpireDate.Valid = true
		//check Price range
		if v, err := param.Price.Float64Value(); err != nil || v.Float64 < 0 {
			logger.Errorw("price is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		//check expire time
		if param.ExpireDate.Time.Before(time.Now()) {
			logger.Errorw("expire date is invalid")
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		fileHeader, err := c.FormFile("image")
		if err == nil {
			imageID, err := mc.PutFile(c.Request().Context(), fileHeader, common.CreateUniqueFileName(fileHeader))
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
		product.ImageUrl = mc.GetFileURL(c.Request().Context(), product.ImageUrl)
		return c.JSON(http.StatusOK, product)
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
func DeleteProduct(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)

	}
}

// @Summary		Seller add product tag
// @Description	Add tag on product
// @Tags			Seller, Shop, Product,Tag
// @Accept			json
// @Param			id		path	string		true	"product id"
// @Param			tag_id	body	TagParams	true	"add tag id" // FIXME
// @Produce		json
// @Success		200	{object}	db.ProductTag
// @Failure		400	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/product/{id}/tag [post]
func AddProductTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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

// @Summary		Seller delete product tag
// @Description	Delete product for shop.
// @Tags			Seller, Shop, Coupon
// @Param			id		path	int				true	"product id"
// @Param			tag_id	body	GetTagParams	true	"add tag id"
// @Accept			json
// @Produce		json
// @Success		200	{string}	string	"success"
// @Failure		400	{object}	echo.HTTPError
// @Failure		404	{object}	echo.HTTPError
// @Failure		500	{object}	echo.HTTPError
// @Router			/seller/product/{id}/tag [delete]
func DeleteProductTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, constants.SUCCESS)
	}
}
