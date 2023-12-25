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

	// admin
	api.GET("/admin/user", admin.GetUser(pg, mc, logger))
	api.DELETE("/admin/user/:username", admin.DisableUser(pg, logger))

	api.GET("/admin/coupon", admin.GetCoupon(pg, logger))
	api.GET("/admin/coupon/:id", admin.GetCouponDetail(pg, logger))
	api.POST("/admin/coupon", admin.AddCoupon(pg, logger))
	api.PATCH("/admin/coupon/:id", admin.EditCoupon(pg, logger))
	api.DELETE("/admin/coupon/:id", admin.DeleteCoupon(pg, logger))

	api.GET("/admin/report", admin.GetReport(pg, mc, logger))

	// user
	api.GET("/user/info", user.GetInfo(pg, mc, logger))
	api.PATCH("/user/info", user.EditInfo(pg, mc, logger))
	api.POST("/user/security/password", user.EditPassword(pg, logger))

	api.GET("/user/security/credit_card", user.GetCreditCard(pg, logger))
	api.PATCH("/user/security/credit_card", user.UpdateCreditCard(pg, logger))

	// general
	api.GET("/shop/:seller_name", general.GetShopInfo(pg, mc, logger)) // user
	api.GET("/shop/:seller_name/coupon", general.GetShopCoupon(pg, logger))
	api.GET("/shop/:seller_name/search", general.SearchShopProduct(pg, mc, logger))

	api.GET("/tag/:id", general.GetTagInfo(pg, logger))

	api.GET("/search", general.Search(pg, mc, logger)) // search both product and shop
	api.GET("/search/shop", general.SearchShopByName(pg, mc, logger))

	api.GET("/news", general.GetNews(pg, logger))
	api.GET("/news/:id", general.GetNewsDetail(pg, logger))
	api.GET("/discover", general.GetDiscover(pg, mc, logger))
	api.GET("/popular", general.GetPopular(pg, mc, logger))

	api.GET("/product/:id", general.GetProductInfo(pg, mc, logger))

	// buyer
	api.GET("/buyer/order", buyer.GetOrderHistory(pg, mc, logger))
	api.GET("/buyer/order/:id", buyer.GetOrderDetail(pg, mc, logger))
	api.PATCH("/buyer/order/:id", buyer.UpdateOrderStatus(pg, logger))

	api.GET("/buyer/cart", buyer.GetCart(pg, mc, logger)) // include product and coupon
	api.GET("/buyer/cart/:id/coupon", buyer.GetCoupon(pg, logger))
	api.POST("/buyer/cart/product/:id", buyer.AddProductToCart(pg, logger))
	api.POST("/buyer/cart/:cart_id/coupon/:coupon_id", buyer.AddCouponToCart(pg, logger))
	api.PATCH("/buyer/cart/:cart_id/product/:product_id", buyer.EditProductInCart(pg, logger))
	api.DELETE("/buyer/cart/:cart_id/product/:product_id", buyer.DeleteProductFromCart(pg, logger))
	api.DELETE("/buyer/cart/:cart_id/coupon/:coupon_id", buyer.DeleteCouponFromCart(pg, logger))

	api.GET("/buyer/cart/:id/checkout", buyer.GetCheckout(pg, logger))
	api.POST("/buyer/cart/:id/checkout", buyer.Checkout(pg, logger))

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

	api.GET("/seller/report", seller.GetReportDetail(pg, mc, logger))

	api.GET("/seller/product", seller.ListProduct(pg, mc, logger))
	api.POST("/seller/product", seller.AddProduct(pg, mc, logger))
	api.GET("/seller/product/:id", seller.GetProductDetail(pg, mc, logger))
	api.PATCH("/seller/product/:id", seller.EditProduct(pg, mc, logger))
	api.POST("/seller/product/:id/tag", seller.AddProductTag(pg, logger))
	api.DELETE("/seller/product/:id/tag", seller.DeleteProductTag(pg, logger))
	api.DELETE("/seller/product/:id", seller.DeleteProduct(pg, logger))

}
