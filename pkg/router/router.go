package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"

	"github.com/jykuo-love-shiritori/twp/db"
	_ "github.com/jykuo-love-shiritori/twp/docs"
	"github.com/jykuo-love-shiritori/twp/pkg/auth"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
)

//	@title			twp API
//	@version		0.o
//	@description	twp server api.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func RegisterDocs(e *echo.Echo) {
	docs := e.Group(constants.SWAGGER_PATH)
	docs.GET("/*", echoSwagger.WrapHandler)
}

func RegisterApi(e *echo.Echo, pg *db.DB, logger *zap.SugaredLogger) {
	api := e.Group("/api")

	api.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "pong"})
	}, auth.IsRole(pg, logger, db.RoleTypeAdmin))

	api.GET("/delay", func(c echo.Context) error {
		time.Sleep(1 * time.Second)
		return c.JSON(http.StatusOK, map[string]string{"message": "delay"})
	})

	api.POST("/signup", auth.Signup(pg, logger))

	api.POST("/oauth/authorize", auth.Authorize(pg, logger))
	api.POST("/oauth/token", auth.Token(pg, logger))

	// admin
	api.GET("/admin/user", adminGetUser(pg, logger))
	api.DELETE("/admin/user/:username", adminDisableUser(pg, logger))

	api.GET("/admin/coupon", adminGetCoupon(pg, logger))
	api.GET("/admin/coupon/:id", adminGetCouponDetail(pg, logger))
	api.POST("/admin/coupon", adminAddCoupon(pg, logger))
	api.PATCH("/admin/coupon/:id", adminEditCoupon(pg, logger))
	api.DELETE("/admin/coupon/:id", adminDeleteCoupon(pg, logger))

	api.GET("/admin/report", adminGetReport(pg, logger))

	// user
	api.GET("/user/info", userGetInfo(pg, logger))
	api.PATCH("/user/info", userEditInfo(pg, logger))
	api.POST("/user/info/upload", userUploadAvatar(pg, logger))
	api.POST("/user/security/password", userEditPassword(pg, logger))

	api.GET("/user/security/credit_card", userGetCreditCard(pg, logger))
	api.PATCH("/user/security/credit_card", userUpdateCreditCard(pg, logger))

	// general
	api.GET("/shop/:seller_name", getShopInfo(pg, logger)) // user
	api.GET("/shop/:seller_name/coupon", getShopCoupon(pg, logger))
	api.GET("/shop/:seller_name/search", searchShopProduct(pg, logger))

	api.GET("/tag/:id", getTagInfo(pg, logger))

	api.GET("/search", search(pg, logger)) // search both product and shop
	api.GET("/search/shop", searchShopByName(pg, logger))

	api.GET("/news", getNews(pg, logger))
	api.GET("/news/:id", getNewsDetail(pg, logger))
	api.GET("/discover", getDiscover(pg, logger))

	api.GET("/product/:id", getProductInfo(pg, logger))

	// buyer
	api.GET("/buyer/order", buyerGetOrderHistory(pg, logger))
	api.GET("/buyer/order/:id", buyerGetOrderDetail(pg, logger))

	api.GET("/buyer/cart", buyerGetCart(pg, logger)) // include product and coupon
	api.POST("/buyer/cart/:cart_id/product/:product_id", buyerAddProductToCart(pg, logger))
	api.POST("/buyer/cart/:cart_id/coupon/:coupon_id", buyerAddCouponToCart(pg, logger))
	api.PATCH("buyer/cart/:cart_id/product/:product_id", buyerEditProductInCart(pg, logger))
	api.DELETE("/buyer/:cart_id/product/:product_id", buyerDeleteProductFromCart(pg, logger))
	api.DELETE("/buyer/cart/:cart_id/coupon/:coupon_id", buyerDeleteCouponFromCart(pg, logger))

	api.GET("/buyer/cart/:cart_id/checkout", buyerGetCheckout(pg, logger))
	api.POST("/buyer/cart/:cart_id/checkout", buyerCheckout(pg, logger))

	// seller
	api.GET("/seller/info", sellerGetShopInfo(pg, logger))
	api.PATCH("/seller/info", sellerEditInfo(pg, logger))
	api.GET("/seller/tag", sellerGetTag(pg, logger))  // search available tag
	api.POST("/seller/tag", sellerAddTag(pg, logger)) // add tag for shop

	api.GET("/seller/coupon", sellerGetShopCoupon(pg, logger))
	api.GET("/seller/coupon/:id", sellerGetCouponDetail(pg, logger))
	api.POST("/seller/coupon", sellerAddCoupon(pg, logger))
	api.PATCH("/seller/coupon/:id", sellerEditCoupon(pg, logger))
	api.DELETE("/seller/coupon/:id", sellerDeleteCoupon(pg, logger))
	api.POST("/seller/coupon/:id/tag", sellerAddCouponTag(pg, logger))
	api.DELETE("/seller/coupon/:id/tag", sellerDeleteCouponTag(pg, logger))

	api.GET("/seller/order", sellerGetOrder(pg, logger))
	api.GET("/seller/order/:id", sellerGetOrderDetail(pg, logger))
	api.PATCH("/seller/order/:id", sellerUpdateOrderStatus(pg, logger))

	api.GET("/seller/report", sellerGetReport(pg, logger))
	api.GET("/seller/report/:year/:month", sellerGetReportDetail(pg, logger))

	api.GET("/seller/product", sellerListProduct(pg, logger))
	api.POST("/seller/product", sellerAddProduct(pg, logger))
	api.POST("/seller/product/:id/upload", sellerUploadProductImage(pg, logger))
	api.GET("/seller/product/:id", sellerGetProductDetail(pg, logger))
	api.PATCH("/seller/product/:id", sellerEditProduct(pg, logger))
	api.POST("/seller/product/:id/tag", sellerAddProductTag(pg, logger))
	api.DELETE("/seller/product/:id/tag", sellerDeleteProductTag(pg, logger))
	api.DELETE("/seller/product/:id", sellerDeleteProduct(pg, logger))
}
