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
	"github.com/jykuo-love-shiritori/twp/pkg/router/admin"
	"github.com/jykuo-love-shiritori/twp/pkg/router/buyer"
	"github.com/jykuo-love-shiritori/twp/pkg/router/general"
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

	// general
	api.GET("/shop/:seller_name", general.GetShopInfo(pg, mc, logger)) // user
	api.GET("/shop/:seller_name/coupon", general.GetShopCoupon(pg, logger))
	api.GET("/shop/:seller_name/coupon/:id", general.GetShopCouponDetail(pg, logger))
	api.GET("/shop/:seller_name/search", general.SearchShopProduct(pg, mc, logger))

	api.GET("/tag/:id", general.GetTagInfo(pg, logger))

	api.GET("/search", general.Search(pg, mc, logger)) // search both product and shop
	api.GET("/search/shop", general.SearchShopByName(pg, mc, logger))

	api.GET("/news", general.GetNews(pg, logger))
	api.GET("/news/:id", general.GetNewsDetail(pg, logger))
	api.GET("/discover", general.GetDiscover(pg, mc, logger))
	api.GET("/popular", general.GetPopular(pg, mc, logger))

	api.GET("/product/:id", general.GetProductInfo(pg, mc, logger))

	// admin
	adminEndpoint := api.Group("", auth.IsRole(pg, logger, db.RoleTypeAdmin))
	adminEndpoint.GET("/admin/user", admin.GetUser(pg, mc, logger))
	adminEndpoint.DELETE("/admin/user/:username", admin.DisableUser(pg, logger))

	adminEndpoint.GET("/admin/coupon", admin.GetCoupon(pg, logger))
	adminEndpoint.GET("/admin/coupon/:id", admin.GetCouponDetail(pg, logger))
	adminEndpoint.POST("/admin/coupon", admin.AddCoupon(pg, logger))
	adminEndpoint.PATCH("/admin/coupon/:id", admin.EditCoupon(pg, logger))
	adminEndpoint.DELETE("/admin/coupon/:id", admin.DeleteCoupon(pg, logger))

	adminEndpoint.GET("/admin/report", admin.GetReport(pg, mc, logger))

	// user
	userEndpoint := api.Group("", auth.ValidateJwt(pg, logger))
	userEndpoint.GET("/user/info", user.GetInfo(pg, mc, logger))
	userEndpoint.PATCH("/user/info", user.EditInfo(pg, mc, logger))
	userEndpoint.POST("/user/security/password", user.EditPassword(pg, logger))

	userEndpoint.GET("/user/security/credit_card", user.GetCreditCard(pg, logger))
	userEndpoint.PATCH("/user/security/credit_card", user.UpdateCreditCard(pg, logger))

	// buyer
	customer := api.Group("", auth.IsRole(pg, logger, db.RoleTypeCustomer))
	customer.GET("/buyer/order", buyer.GetOrderHistory(pg, mc, logger))
	customer.GET("/buyer/order/:id", buyer.GetOrderDetail(pg, mc, logger))
	customer.PATCH("/buyer/order/:id", buyer.UpdateOrderStatus(pg, logger))

	customer.GET("/buyer/cart", buyer.GetCart(pg, mc, logger)) // include product and coupon
	customer.GET("/buyer/cart/:id/coupon", buyer.GetCoupon(pg, logger))
	customer.POST("/buyer/cart/product/:id", buyer.AddProductToCart(pg, logger))
	customer.POST("/buyer/cart/:cart_id/coupon/:coupon_id", buyer.AddCouponToCart(pg, logger))
	customer.PATCH("/buyer/cart/:cart_id/product/:product_id", buyer.EditProductInCart(pg, logger))
	customer.DELETE("/buyer/cart/:cart_id/product/:product_id", buyer.DeleteProductFromCart(pg, logger))
	customer.DELETE("/buyer/cart/:cart_id/coupon/:coupon_id", buyer.DeleteCouponFromCart(pg, logger))

	customer.GET("/buyer/cart/:id/checkout", buyer.GetCheckout(pg, logger))
	customer.POST("/buyer/cart/:id/checkout", buyer.Checkout(pg, logger))

	// seller
	customer.GET("/seller/info", seller.GetShopInfo(pg, mc, logger))
	customer.PATCH("/seller/info", seller.EditInfo(pg, mc, logger))
	customer.GET("/seller/tag", seller.GetTag(pg, logger))  // search available tag
	customer.POST("/seller/tag", seller.AddTag(pg, logger)) // add tag for shop

	customer.GET("/seller/coupon", seller.GetShopCoupon(pg, logger))
	//sqlc only take one path param tag overwrite
	customer.GET("/seller/coupon/:id", seller.GetCouponDetail(pg, logger))
	customer.POST("/seller/coupon", seller.AddCoupon(pg, logger))
	customer.PATCH("/seller/coupon/:id", seller.EditCoupon(pg, logger))
	customer.DELETE("/seller/coupon/:id", seller.DeleteCoupon(pg, logger))
	customer.POST("/seller/coupon/:id/tag", seller.AddCouponTag(pg, logger))
	customer.DELETE("/seller/coupon/:id/tag", seller.DeleteCouponTag(pg, logger))

	customer.GET("/seller/order", seller.GetOrder(pg, mc, logger))
	customer.GET("/seller/order/:id", seller.GetOrderDetail(pg, mc, logger))
	customer.PATCH("/seller/order/:id", seller.UpdateOrderStatus(pg, logger))

	customer.GET("/seller/report", seller.GetReportDetail(pg, mc, logger))

	customer.GET("/seller/product", seller.ListProduct(pg, mc, logger))
	customer.POST("/seller/product", seller.AddProduct(pg, mc, logger))
	customer.GET("/seller/product/:id", seller.GetProductDetail(pg, mc, logger))
	customer.PATCH("/seller/product/:id", seller.EditProduct(pg, mc, logger))
	customer.POST("/seller/product/:id/tag", seller.AddProductTag(pg, logger))
	customer.DELETE("/seller/product/:id/tag", seller.DeleteProductTag(pg, logger))
	customer.DELETE("/seller/product/:id", seller.DeleteProduct(pg, logger))
}
