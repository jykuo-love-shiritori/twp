package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"

	"github.com/jykuo-love-shiritori/twp/db"
	_ "github.com/jykuo-love-shiritori/twp/docs"
	"github.com/jykuo-love-shiritori/twp/minio"
	"github.com/jykuo-love-shiritori/twp/pkg/auth"
	"github.com/jykuo-love-shiritori/twp/pkg/constants"
	"github.com/jykuo-love-shiritori/twp/pkg/router/seller"
	"github.com/jykuo-love-shiritori/twp/pkg/router/user"
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

func RegisterApi(e *echo.Echo, pg *db.DB, mc *minio.MC, logger *zap.SugaredLogger) {
	api := e.Group("/api")

	api.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "pong"})
	}, auth.IsRole(pg, logger, db.RoleTypeCustomer))

	api.GET("/delay", func(c echo.Context) error {
		time.Sleep(1 * time.Second)
		return c.JSON(http.StatusOK, map[string]string{"message": "delay"})
	})

	api.POST("/signup", auth.Signup(pg, logger))

	api.POST("/oauth/authorize", auth.Authorize(pg, logger))
	api.POST("/oauth/token", auth.Token(pg, logger))
	api.POST("/oauth/refresh", auth.Refresh(pg, logger))
	api.POST("/oauth/logout", auth.Logout(pg, logger), auth.ValidateJwt(pg, logger))

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
	api.GET("/user/info", user.GetInfo(pg, mc, logger))
	api.PATCH("/user/info", user.EditInfo(pg, mc, logger))
	api.POST("/user/security/password", user.EditPassword(pg, logger))

	api.GET("/user/security/credit_card", user.GetCreditCard(pg, logger))
	api.PATCH("/user/security/credit_card", user.UpdateCreditCard(pg, logger))

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

	api.GET("/product/:product_id", getProductInfo(pg, logger))

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
	api.GET("/seller/info", seller.GetShopInfo(pg, mc, logger))
	api.PATCH("/seller/info", seller.EditInfo(pg, mc, logger))
	api.GET("/seller/tag", seller.GetTag(pg, logger))  // search available tag
	api.POST("/seller/tag", seller.AddTag(pg, logger)) // add tag for shop

	api.GET("/seller/coupon", seller.GetShopCoupon(pg, logger))
	//sqlc only take one path param tag overwrite
	api.GET("/seller/coupon/:id", seller.GetCouponDetail(pg, logger))
	api.POST("/seller/coupon", seller.AddCoupon(pg, logger))
	api.PATCH("/seller/coupon/:id", seller.EditCoupon(pg, logger))
	api.DELETE("/seller/coupon/:id", seller.DeleteCoupon(pg, logger))
	api.POST("/seller/coupon/:id/tag", seller.AddCouponTag(pg, logger))
	api.DELETE("/seller/coupon/:id/tag", seller.DeleteCouponTag(pg, logger))

	api.GET("/seller/order", seller.GetOrder(pg, mc, logger))
	api.GET("/seller/order/:id", seller.GetOrderDetail(pg, mc, logger))
	api.PATCH("/seller/order/:id", seller.UpdateOrderStatus(pg, logger))

	api.GET("/seller/report/:year/:month", seller.GetReportDetail(pg, mc, logger))

	api.GET("/seller/product", seller.ListProduct(pg, mc, logger))
	api.POST("/seller/product", seller.AddProduct(pg, mc, logger))
	api.GET("/seller/product/:id", seller.GetProductDetail(pg, mc, logger))
	api.PATCH("/seller/product/:id", seller.EditProduct(pg, mc, logger))
	api.POST("/seller/product/:id/tag", seller.AddProductTag(pg, logger))
	api.DELETE("/seller/product/:id/tag", seller.DeleteProductTag(pg, logger))
	api.DELETE("/seller/product/:id", seller.DeleteProduct(pg, logger))

	api.GET("/seller/product", seller.ListProduct(pg, mc, logger))
	api.POST("/seller/product", seller.AddProduct(pg, mc, logger))
	api.GET("/seller/product/:id", seller.GetProductDetail(pg, mc, logger))
	api.PATCH("/seller/product/:id", seller.EditProduct(pg, mc, logger))
	api.POST("/seller/product/:id/tag", seller.AddProductTag(pg, logger))
	api.DELETE("/seller/product/:id/tag", seller.DeleteProductTag(pg, logger))
	api.DELETE("/seller/product/:id", seller.DeleteProduct(pg, logger))
}
