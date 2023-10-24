package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/jykuo-love-shiritori/twp/docs"
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
	e.GET("/swagger/*", echoSwagger.WrapHandler)

}
func RegisterApi(e *echo.Echo) {

	api := e.Group("/api")

	api.POST("/logout", logout)

	// admin
	api.GET("/admin/user", adminGetUser)
	api.DELETE("/admin/user/:id", adminDeleteUser)

	api.GET("/admin/coupon", adminGetCoupon)
	api.GET("/admin/coupon/:id", adminGetCouponDetail)
	api.POST("/admin/coupon", adminAddCoupon)
	api.PATCH("/admin/coupon/:id", adminEditCoupon)
	api.DELETE("/admin/coupon/:id", adminDeleteCoupon)

	api.GET("/admin/report", adminGetReport)

	// user
	api.GET("/user/info", userGetInfo)
	api.PATCH("/user/info", userEditInfo)
	api.POST("/user/info/upload", userUploadAvatar)
	api.POST("/user/security/password", userEditPassword)

	api.GET("/user/security/credit_card", userGetCreditCard)
	api.DELETE("/user/security/credit_card", userDeleteCreditCard)
	api.POST("/user/security/credit_card", userAddCreditCard)

	// general
	api.GET("/shop/:id", getShopInfo) // user
	api.GET("/shop/:id/coupon", getShopCoupon)
	api.GET("/shop/:id/search", searchShopProduct)

	api.GET("/tag/:id", getTagInfo)

	api.GET("/search", search) // search both product and shop
	api.GET("/search/shop", searchShopByName)

	api.GET("/news", getNews)
	api.GET("/news/:id", getNewsDetail)
	api.GET("/discover", getDiscover)

	api.GET("/product/:id", getProductInfo)

	// buyer
	api.GET("/buyer/order", buyerGetOrderHistrory)
	api.GET("/buyer/order/:id", buyerGetOrderDetail)

	api.GET("/buyer/cart", buyerGetCart) // include procuct and coupon
	api.POST("/buyer/cart/product:id", buyerAddProductToCart)
	api.POST("/buyer/cart/coupon:id", buyerAddCouponToCart)
	api.PATCH("buyer/cart/product:id", buyerEditProductInCart)
	api.DELETE("/buyer/cart/product:id", buyerDeleteProductFromCart)
	api.DELETE("/buyer/cart/coupon:id", buyerDeleteCouponFromCart)

	api.GET("/buyer/checkout", buyerGetCheckout)
	api.POST("/buyer/checkout", buyerCheckout)

	// seller
	api.GET("/seller", sellerGetShopInfo)
	api.PATCH("/seller", sellerEditInfo)
	api.GET("/seller/tag", sellerGetTag)  // search avaliable tag
	api.POST("/seller/tag", sellerAddTag) // add tag for shop

	api.GET("/seller/coupon", sellerGetShopCoupon)
	api.GET("/seller/coupon/:id", sellerGetCouponDetail)
	api.POST("/seller/coupon", sellerAddCoupon)
	api.PATCH("/seller/coupon/:id", sellerEditCoupon)
	api.DELETE("/seller/coupon/:id", sellerDeleteCoupon)

	api.GET("/seller/order", sellerGetOrder)
	api.GET("/seller/order/:id", sellerGetOrderDetail)

	api.GET("/seller/report", sellerGetReport)
	api.GET("/seller/report/:year/:month", sellerGetReportDetail)

	api.POST("/seller/product", sellerAddProduct)
	api.POST("/seller/product/:id/upload", sellerUploadProductImage)
	api.PATCH("/seller/product/:id", sellerEditProduct)
	api.DELETE("/seller/product/:id", sellerDeleteProduct)

}

/*
## User
- User Login
- User Sign up

- User Get its personal data
- User Get order history
- User Search Product
- User Get specific seller's shop
- User Get specific seller's coupon
- User Get specific Product data
- User Add Product into cart
- User Get shopping cart inventory
- User Get Checkout data
- User Get all avaliable coupon in cart
- User Add Coupon into cart
- User Get specific order data
- User Get Main Page News /pending
- User Get Popular Products
## Seller
- Seller Add Product
- Seller Add tag for Product
- Seller Edit tag for Product
- Seller Edit exist Product
- Seller Get all Shipments
- Seller Get all its coupon
- Seller Add coupon
- Seller Edit coupon
- Seller Get Sell Report
- Sell Get Specific Sell Report
## Admin
- Admin Get all user Data
- Admin ban/delete user
- Admin peek site report

*/
