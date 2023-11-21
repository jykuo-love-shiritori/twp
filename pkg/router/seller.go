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
type failure struct {
	code int
	fail string
}

func sellerGetShopInfo(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {

	return func(c echo.Context) error {
		var userID int32 = 0

		shopInfo, err := d.Queries.GetSellerInfo(context.Background(), userID)
		logger.Info(shopInfo)
		if err != nil {
			logger.Fatal(err)
			return c.JSON(http.StatusInternalServerError, failure{"error"})
		}
		return c.JSON(http.StatusOK, shopInfo)
	}
}

// @Summary Seller edit shop info
// @Description Edit shop name, description, visibility.
// @Tags Seller, Shop
// @Param  image_id       query     string  false  "name search by q"  Format(email)
// @Param  name           query     string  false  "name search by q"  Format(email)
// @Param  Description    query     string  false  "name search by q"  Format(email)
// @Param  enable         query     string  false  "name search by q"  Format(email)
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller [patch]
func sellerEditInfo(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {

	return func(c echo.Context) error {
		var userID int32 = 0

		var param db.UpdateSellerInfoParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.ID = userID
		err := d.Queries.UpdateSellerInfo(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return c.JSON(http.StatusInternalServerError, failure{"error"})
		}
		return c.JSON(http.StatusOK, param)
	}
}

// @Summary Seller get available tag
// @Description Get all available tags for shop.
// @Tags Seller, Shop, Tag
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/tag [get]
func sellerGetTag(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {

	return func(c echo.Context) error {
		var userID int32 = 1

		var param db.SearchTagParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.ID = userID
		param.Name = "^" + param.Name

		tag, err := d.Queries.SearchTag(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return c.JSON(http.StatusInternalServerError, failure{"error"})
		}
		return c.JSON(http.StatusOK, tag)
	}
}

// @Summary Seller add tag
// @Description Add tag for shop.
// @Tags Seller, Shop, Tag
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/tag [post]
func sellerAddTag(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userName string = "user0"

		var param db.InsertTagParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = userName

		tag, err := d.Queries.InsertTag(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return c.JSON(http.StatusInternalServerError, failure{"error"})
		}
		return c.JSON(http.StatusOK, tag)
	}
}

// @Summary Seller get shop coupon
// @Description Get all coupons for shop.
// @Tags Seller, Shop, Coupon
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/coupon [get]
func sellerGetShopCoupon(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var userName string = "user0"

		var param db.SellerGetCouponParams
		if err := c.Bind(&param); err != nil {
			logger.Errorw("Error binding request parameters", "error", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
		}
		param.SellerName = userName

		tag, err := d.Queries.SellerGetCoupon(context.Background(), param)
		if err != nil {
			logger.Fatal(err)
			return c.JSON(http.StatusInternalServerError, failure{"error"})
		}
		return c.JSON(http.StatusOK, tag)
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
func sellerGetCouponDetail(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller add coupon
// @Description Add coupon for shop.
// @Tags Seller, Shop, Coupon
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/coupon [post]
func sellerAddCoupon(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller edit coupon
// @Description Edit coupon for shop.
// @Tags Seller, Shop, Coupon
// @Accept json
// @Produce json
// @Param id path int true "Coupon ID"
// @Success 200
// @Failure 401
// @Router /seller/coupon/{id} [patch]
func sellerEditCoupon(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller delete coupon
// @Description Delete coupon for shop.
// @Tags Seller, Shop, Coupon
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
func sellerDeleteCoupon(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller get order
// @Description Get all orders for shop.
// @Tags Seller, Shop, Order
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/order [get]
func sellerGetOrder(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
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
func sellerGetOrderDetail(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
func sellerGetReport(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
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
func sellerGetReportDetail(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller add product
// @Description Add product for shop.
// @Tags Seller, Shop, Product
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/product [post]
func sellerAddProduct(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
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
func sellerUploadProductImage(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller edit product
// @Description Edit product for shop.
// @Tags Seller, Shop, Product
// @Accept json
// @Produce json
// @Param id query int true "Product ID"
// @Success 200
// @Failure 401
// @Router /seller/product/{id} [patch]
func sellerEditProduct(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}

// @Summary Seller delete product
// @Description Delete product for shop.
// @Tags Seller, Shop, Product
// @Accept json
// @Produce json
// @Param id query int true "Product ID"
// @Success 200
// @Failure 401
// @Router /seller/product/{id} [delete]
func sellerDeleteProduct(d *db.DB, logger *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.NoContent(http.StatusOK)
	}
}
