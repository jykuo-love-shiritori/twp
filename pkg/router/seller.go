package router

import (
	"context"
	"net/http"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary Seller get shop info
// @Description Get shop info, includes user picture, name, description.
// @Tags Seller, Shop
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller [get]
func sellerGetShopInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID int32 = 0

		shopInfo, err := pg.Queries.GetSellerInfo(context.Background(), userID)
		logger.Info(shopInfo)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		return c.JSON(http.StatusOK, shopInfo)

	}
}

// @Summary Seller edit shop info
// @Description Edit shop name, description, visibility.
// @Tags Seller, Shop
// @Param  image_id      body     string  true  "update image UUID"
// @Param  name          body     string  true  "update shop name"       minlength(6)
// @Param  description   body     string  true  "update description"
// @Param  enabled       body     bool    true  "update enabled status"
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller [patch]
func sellerEditInfo(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID int32 = 0

		var param db.UpdateSellerInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.ID = userID
		err := pg.Queries.UpdateSellerInfo(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		return c.JSON(http.StatusOK, param)
	}
}

// @Summary Seller get available tag
// @Description Get all available tags for shop.
// @Tags Seller, Shop, Tag
// @Param  name   body    string  true  "search tagname start with"     minlength(1)
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/tag [get]
func sellerGetTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {

	return func(c echo.Context) error {
		var userID int32 = 0

		var param db.SearchTagParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		if param.Name == "" || hasSpecialChars(param.Name) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "tag name can be ''"})
		}
		param.ID = userID
		param.Name = "^" + param.Name
		tags, err := pg.Queries.SearchTag(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		return c.JSON(http.StatusOK, tags)
	}
}

// @Summary Seller add tag
// @Description Add tag for shop.
// @Tags Seller, Shop, Tag
// @Accept json
// @Param  name   body    string  true  "insert tag"     minlength(1)
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/tag [post]
func sellerAddTag(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user0"

		var param db.HaveTagNameParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = username
		have, err := pg.Queries.HaveTagName(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		if have {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request (tagname have to be unique)"})
		}
		tag, err := pg.Queries.InsertTag(context.Background(), db.InsertTagParams(param))
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		return c.JSON(http.StatusOK, tag)
	}
}

// @Summary Seller get shop coupon
// @Description Get all coupons for shop.
// @Tags Seller, Shop, Coupon
// @Param  offset   body  int   true  "offset page"   minimum(0)
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/coupon [get]
func sellerGetShopCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user0"
		var couponPerPage int32 = 20

		var param db.SellerGetCouponParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = username
		param.Limit = couponPerPage
		param.Offset = param.Offset * couponPerPage
		coupons, err := pg.Queries.SellerGetCoupon(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		return c.JSON(http.StatusOK, coupons)
	}
}

// @Summary Seller get coupon detail
// @Description Get coupon detail by ID for shop.
// @Tags Seller, Shop, Coupon
// @Produce json
// @Param id path int true "Coupon ID"
// @Success 200
// @Failure 401
// @Router /seller/coupon/{id} [get]
func sellerGetCouponDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		var username string = "user0"

		var param db.SellerGetCouponDetailParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = username
		coupons, err := pg.Queries.SellerGetCouponDetail(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		if len(coupons) != 1 {
			logger.Errorw("Error ", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not Found Coupon"})
		}

		return c.JSON(http.StatusOK, coupons[0])
	}
}

// @Summary Seller add coupon
// @Description Add coupon for shop.
// @Tags Seller, Shop, Coupon
// @Param  type          body     string  true  "Coupon type" Enums('percentage', 'fixed', 'shipping')
// @Param  name          body     string  true  "name of coupon"
// @Param  description   body     string  true  "description of coupon"
// @Param  discount      body     number  false "discount perscent"
// @Param  start_date    body     time    true  "start date"
// @Param  expire_date   body     time    true  "expire date"
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/coupon [post]
func sellerAddCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user0"

		var param db.SellerInsertCouponParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = username
		coupon, err := pg.Queries.SellerInsertCoupon(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		return c.JSON(http.StatusOK, coupon)
	}
}

// @Summary Seller edit coupon
// @Description Edit coupon for shop.
// @Tags Seller, Shop, Coupon
// @Accept json
// @Produce json
// @Param  id            path     int     true  "Coupon ID"
// @Param  type          body     string  true  "Coupon type" Enums('percentage', 'fixed', 'shipping')
// @Param  name          body     string  true  "name of coupon"
// @Param  description   body     string  true  "description of coupon"
// @Param  discount      body     number  false "discount perscent"
// @Param  start_date    body     time    true  "start date"
// @Param  expire_date   body     time    true  "expire date"
// @Success 200
// @Failure 401
// @Router /seller/coupon/{id} [patch]
func sellerEditCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user0"

		var param db.UpdateCouponInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = username
		effectRow, err := pg.Queries.UpdateCouponInfo(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		if effectRow != 1 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not Found Coupon"})
		}
		return c.JSON(http.StatusOK, param)
	}
}

// @Summary Seller delete coupon
// @Description Delete coupon for shop.
// @Tags Seller, Shop, Coupon
// @Param  id       path     int     true  "Coupon ID"
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/coupon/{id} [delete]
func sellerDeleteCoupon(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user0"

		var param db.DeleteCouponParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = username
		effectRow, err := pg.Queries.DeleteCoupon(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		if effectRow != 1 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not Found Coupon"})
		}
		return c.JSON(http.StatusOK, map[string]string{"msg": "success"})
	}
}

// @Summary Seller get order
// @Description Get all orders for shop.
// @Tags Seller, Shop, Order
// @Param  offset   body   int   true  "offset page"   minimum(0)
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/order [get]
func sellerGetOrder(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		var username string = "user0"
		var orderPerPage int32 = 20

		var param db.SellerGetOrderParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		logger.Info(param)
		param.SellerName = username
		param.Limit = orderPerPage
		param.Offset = param.Offset * orderPerPage
		order, err := pg.Queries.SellerGetOrder(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		return c.JSON(http.StatusOK, order)
	}
}

// @Summary Seller get order detail
// @Description Get order detail by ID for shop.
// @Tags Seller, Shop, Order
// @Produce json
// @Param id path int true "Order ID"
// @Success 200
// @Failure 401
// @Router /seller/order/{id} [get]
func sellerGetOrderDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user0"

		var param db.SellerGetOrderDetailParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = username
		order, err := pg.Queries.SellerGetOrderDetail(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		return c.JSON(http.StatusOK, order)
	}
}

// @Summary Seller get order
// @Description Get all orders for shop.
// @Tags Seller, Shop, Order
// @Param  offset   body   int   true  "offset page"   minimum(0)
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/order [patch]
func sellerUpdateOrderStatus(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller get report
// @Description Get all available reports for shop.
// @Tags Seller, Shop, Report
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/report [get]
func sellerGetReport(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller get report detail
// @Description Get report detail by year and month for shop.
// @Tags Seller, Shop, Report
// @Produce json
// @Param year path int true "Year"
// @Param month path int true "Month"
// @Success 200
// @Failure 401
// @Router /seller/report/{year}/{month} [get]
func sellerGetReportDetail(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller add product
// @Description Add product for shop.
// @Tags Seller, Shop, Product
// @Param  id            path     int     true  "Product ID"
// @Param  name          body     string  true  "name of product"
// @Param  description   body     string  true  "description of product"
// @Param  price         body     number  false "price"
// @Param  image_id      body     string  true  "image id"
// @Param  exp_date      body     time    true  "expire date"
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/product [post]
func sellerAddProduct(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user0"

		var param db.SellerInsertProductParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = username
		product, err := pg.Queries.SellerInsertProduct(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		return c.JSON(http.StatusOK, product)
	}
}

// @Summary Seller upload product image
// @Description Upload product image for shop.
// @Tags Seller, Shop, Product
// @Accept png,jpeg,gif
// @Produce json
// @Param id path int true "Product ID"
// @Param img formData file true "image to upload"
// @Success 200
// @Failure 401
// @Router /seller/product/{id}/upload [post]
func sellerUploadProductImage(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller edit product
// @Description Edit product for shop.
// @Tags Seller, Shop, Product
// @Accept json
// @Produce json
// @Param  id            path     int     true  "Product ID"
// @Param  name          body     string  true  "name of product"
// @Param  description   body     string  true  "description of product"
// @Param  price         body     number  false "price"
// @Param  image_id      body     string  true  "image id"
// @Param  exp_date      body     time    true  "expire date"
// @Success 200
// @Failure 401
// @Router /seller/product/{id} [patch]
func sellerEditProduct(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user0"

		var param db.UpdateProductInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = username
		effectRow, err := pg.Queries.UpdateProductInfo(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		if effectRow != 1 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not Found Product"})
		}
		return c.JSON(http.StatusOK, param)
	}
}

// @Summary Seller delete product
// @Description Delete product for shop.
// @Tags Seller, Shop, Product
// @Accept json
// @Produce json
// @Param   id path int true "Product ID"
// @Success 200
// @Failure 401
// @Router /seller/product/{id} [delete]
func sellerDeleteProduct(pg *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var username string = "user0"

		var param db.DeleteProductParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = username
		effectRow, err := pg.Queries.DeleteProduct(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return mapDBErrorToHTTPError(err, c)
		}
		if effectRow != 1 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not Found Product"})
		}
		return c.JSON(http.StatusOK, param)
	}
}
