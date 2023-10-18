package pkg

import (
	"github.com/labstack/echo/v4"
)

func RegisterApi(e *echo.Echo) {
	api := e.Group("/api")
	api.POST("/login", login)
	api.POST("/signup", signup)

	api.GET("/admin/user", adminGetUser)
	api.DELETE("/admin/user/:id", adminDeleteUser)

	api.GET("/admin/coupon", adminGetCoupon)
	api.POST("/admin/coupon", adminAddCoupon)
	api.DELETE("/admin/coupon/:id", adminDeleteCoupon)

	api.GET("/admin/report", adminGetReport)

	api.GET("/user/info", userGetInfo)
	api.POST("/user/info", userEditInfo)
	api.POST("/user/password", userEditPassword)

	api.GET("/user/credit_card", userGetCreditCard)
	api.DELETE("/user/credit_card", userDeleteCreditCard)
	api.POST("/user/credit_card", userAddCreditCard)

	api.GET("/user/order", userGetOrderHistrory)
	api.GET("/user/order/:id", userGetOrderDetail)

	api.GET("/shop/:id", getShopInfo) // user
	api.GET("/shop/:id/coupon", getShopCoupon)
	api.GET("/shop/:id/search", searchShop)

	api.GET("/tag/:id", getTagInfo)

	api.GET("/search", searchProduct)

	api.GET("/product/:id", getProductInfo)

	api.GET("/user/cart", getCart) // include procuct and coupon
	api.POST("/user/cart/product:id", addProductToCart)
	api.POST("/user/cart/coupon:id", addCouponToCart)
	api.DELETE("/user/cart/product:id", deleteProductFromCart)
	api.DELETE("/user/cart/coupon:id", deleteCouponFromCart)

	api.GET("/user/checkout", getCheckout)
	api.POST("/user/checkout", checkout)

	// seller
	api.GET("/user/shop", getShopInfo)
	api.PATCH("/user/shop", editInfo)
	api.GET("/user/shop/tag", getTag)  // search avaliable tag
	api.POST("/user/shop/tag", addTag) // add tag for shop

	api.GET("/user/shop/coupon", getCoupon)
	api.POST("/user/shop/coupon", addCoupon)
	api.PATCH("/user/shop/coupon/:id", editCoupon)
	api.DELETE("/user/shop/coupon/:id", deleteCoupon)

	api.GET("/user/shop/shipment", getShipment)
	api.GET("/user/shop/shipment/:id", getShipmentDetail)

	api.GET("/user/shop/report", getReport)
	api.GET("/user/shop/report/:year/:month", getReportDetail)

	api.POST("/user/shop/product", addProduct)
	api.POST("/user/shop/product/:id/upload", uploadProductImage)
	api.PATCH("/user/shop/product/:id", editProduct)
	api.DELETE("/user/shop/product/:id", deleteProduct)

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
