package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Seller get shop info
// @Description Get shop info, includes user picture, name, description.
// @Tags Seller, Shop
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller [get]
func sellerGetShopInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller edit shop info
// @Description Edit shop name, description, visibility.
// @Tags Seller, Shop
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller [patch]
func sellerEditInfo(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller get avaliable tag
// @Description Get all avaliable tags for shop.
// @Tags Seller, Shop, Tag
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/tag [get]
func sellerGetTag(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller add tag
// @Description Add tag for shop.
// @Tags Seller, Shop, Tag
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/tag [post]
func sellerAddTag(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller get shop coupon
// @Description Get all coupons for shop.
// @Tags Seller, Shop, Coupon
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/coupon [get]
func sellerGetShopCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller get coupon detail
// @Description Get coupon detail by ID for shop.
// @Tags Seller, Shop, Coupon
// @Produce json
// @Param id path int true "Coupon ID"
// @Success 200
// @Failure 401
// @Router /seller/coupon/{id} [get]
func sellerGetCouponDetail(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller add coupon
// @Description Add coupon for shop.
// @Tags Seller, Shop, Coupon
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/coupon [post]
func sellerAddCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
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
func sellerEditCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller delete coupon
// @Description Delete coupon for shop.
// @Tags Seller, Shop, Coupon
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
func sellerDeleteCoupon(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller get order
// @Description Get all orders for shop.
// @Tags Seller, Shop, Order
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/order [get]
func sellerGetOrder(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller get order detail
// @Description Get order detail by ID for shop.
// @Tags Seller, Shop, Order
// @Produce json
// @Param id path int true "Order ID"
// @Success 200
// @Failure 401
// @Router /seller/order/{id} [get]
func sellerGetOrderDetail(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller get report
// @Description Get all avaliable reports for shop.
// @Tags Seller, Shop, Report
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/report [get]
func sellerGetReport(c echo.Context) error {
	return c.NoContent(http.StatusOK)
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
func sellerGetReportDetail(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// @Summary Seller add product
// @Description Add product for shop.
// @Tags Seller, Shop, Product
// @Accept json
// @Produce json
// @Success 200
// @Failure 401
// @Router /seller/product [post]
func sellerAddProduct(c echo.Context) error {
	return c.NoContent(http.StatusOK)
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
func sellerUploadProductImage(c echo.Context) error {
	return c.NoContent(http.StatusOK)
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
func sellerEditProduct(c echo.Context) error {
	return c.NoContent(http.StatusOK)
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
func sellerDeleteProduct(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
