package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"

	"github.com/jykuo-love-shiritori/twp/db"
	_ "github.com/jykuo-love-shiritori/twp/docs"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
)

// @title           twp API
// @version         0.o
// @description     twp server api.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func RegisterDocs(e *echo.Echo) {
	docs := e.Group(constants.SWAGGER_PATH)
	docs.GET("/*", echoSwagger.WrapHandler)
}

func RegisterApi(e *echo.Echo, db *db.DB, logger *zap.SugaredLogger) {
	api := e.Group("/api")

	api.GET("/ping", func(c echo.Context) error { return c.JSON(http.StatusOK, map[string]string{"message": "pong"}) })

	api.POST("/logout", logout(db, logger))

	// admin
	api.GET("/admin/user", adminGetUser(db, logger))
	api.PATCH("/admin/user/:username", adminDisableUser(db, logger))

	api.GET("/admin/coupon", adminGetCoupon(db, logger))
	api.GET("/admin/coupon/:id", adminGetCouponDetail(db, logger))
	api.POST("/admin/coupon", adminAddCoupon(db, logger))
	api.PATCH("/admin/coupon/:id", adminEditCoupon(db, logger))
	api.DELETE("/admin/coupon/:id", adminDeleteCoupon(db, logger))

	api.GET("/admin/report", adminGetReport(db, logger))

	// user
	api.GET("/user/info", userGetInfo(db, logger))
	api.PATCH("/user/info", userEditInfo(db, logger))
	api.POST("/user/info/upload", userUploadAvatar(db, logger))
	api.POST("/user/security/password", userEditPassword(db, logger))

	api.GET("/user/security/credit_card", userGetCreditCard(db, logger))
	api.DELETE("/user/security/credit_card", userDeleteCreditCard(db, logger))
	api.POST("/user/security/credit_card", userAddCreditCard(db, logger))

	// general
	api.GET("/shop/:seller_name", getShopInfo(db, logger)) // user
	api.GET("/shop/:seller_name/coupon", getShopCoupon(db, logger))
	api.GET("/shop/:seller_name/search", searchShopProduct(db, logger))

	api.GET("/tag/:id", getTagInfo(db, logger))

	api.GET("/search", search(db, logger)) // search both product and shop
	api.GET("/search/shop", searchShopByName(db, logger))

	api.GET("/news", getNews(db, logger))
	api.GET("/news/:id", getNewsDetail(db, logger))
	api.GET("/discover", getDiscover(db, logger))

	api.GET("/product/:id", getProductInfo(db, logger))

	// buyer
	api.GET("/buyer/order", buyerGetOrderHistory(db, logger))
	api.GET("/buyer/order/:id", buyerGetOrderDetail(db, logger))

	api.GET("/buyer/cart", buyerGetCart(db, logger)) // include product and coupon
	api.POST("/buyer/cart/:cart_id/product/:product_id", buyerAddProductToCart(db, logger))
	api.POST("/buyer/cart/:cart_id/coupon/:coupon_id", buyerAddCouponToCart(db, logger))
	api.PATCH("buyer/cart/:cart_id/product/:product_id", buyerEditProductInCart(db, logger))
	api.DELETE("/buyer/:cart_id/product/:product_id", buyerDeleteProductFromCart(db, logger))
	api.DELETE("/buyer/cart/:cart_id/coupon/:coupon_id", buyerDeleteCouponFromCart(db, logger))

	api.GET("/buyer/cart/:cart_id/checkout", buyerGetCheckout(db, logger))
	api.POST("/buyer/cart/:cart_id/checkout", buyerCheckout(db, logger))

	// seller
	api.GET("/seller", sellerGetShopInfo(db, logger))
	api.PATCH("/seller", sellerEditInfo(db, logger))
	api.GET("/seller/tag", sellerGetTag(db, logger))  // search available tag
	api.POST("/seller/tag", sellerAddTag(db, logger)) // add tag for shop

	api.GET("/seller/coupon", sellerGetShopCoupon(db, logger))
	api.GET("/seller/coupon/:id", sellerGetCouponDetail(db, logger))
	api.POST("/seller/coupon", sellerAddCoupon(db, logger))
	api.PATCH("/seller/coupon/:id", sellerEditCoupon(db, logger))
	api.DELETE("/seller/coupon/:id", sellerDeleteCoupon(db, logger))

	api.GET("/seller/order", sellerGetOrder(db, logger))
	api.GET("/seller/order/:id", sellerGetOrderDetail(db, logger))

	api.GET("/seller/report", sellerGetReport(db, logger))
	api.GET("/seller/report/:year/:month", sellerGetReportDetail(db, logger))

	api.POST("/seller/product", sellerAddProduct(db, logger))
	api.POST("/seller/product/:id/upload", sellerUploadProductImage(db, logger))
	api.PATCH("/seller/product/:id", sellerEditProduct(db, logger))
	api.DELETE("/seller/product/:id", sellerDeleteProduct(db, logger))

}
