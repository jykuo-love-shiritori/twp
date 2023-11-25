package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/jykuo-love-shiritori/twp/db"
)

func TestDummy(t *testing.T) {
	if false {
		t.Fatal("dummy test failed")
	}

}

type testTable struct {
	User           []db.TestInsertUserParams           `json:"user"`
	Shop           []db.TestInsertShopParams           `json:"shop"`
	Coupon         []db.TestInsertCouponParams         `json:"coupon"`
	Product        []db.TestInsertProductParams        `json:"product"`
	ProductArchive []db.TestInsertProductArchiveParams `json:"product_archive"`
	Tag            []db.TestInsertTagParams            `json:"tag"`
	ProductTag     []db.TestInsertProductTagParams     `json:"product_tag"`
	CouponTag      []db.TestInsertCouponTagParams      `json:"coupon_tag"`
	Cart           []db.TestInsertCartParams           `json:"cart"`
	CartProduct    []db.TestInsertCartProductParams    `json:"cart_product"`
	CartCoupon     []db.TestInsertCartCouponParams     `json:"cart_coupon"`
	Order          []db.TestInsertOrderParams          `json:"order_history"`
	OrderDetail    []db.TestInsertOrderDetailParams    `json:"order_detail"`
}

func TestData(t *testing.T) {
	var err error
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	jsonFile, err := os.Open("data.json")
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	var data testTable
	json.Unmarshal(byteValue, &data)
	// fmt.Println(data.User[0].CreditCard)
	fmt.Println("user")
	// for _, userParam := range data.User {
	// 	_, err = db.Queries.TestInsertUser(context.Background(), userParam)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// }
	fmt.Println("shop")
	// for _, shopParam := range data.Shop {
	// 	_, err = db.Queries.TestInsertShop(context.Background(), shopParam)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	fmt.Println("coupon")
	// for _, couponParam := range data.Coupon {
	// 	_, err = db.Queries.TestInsertCoupon(context.Background(), couponParam)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	fmt.Println("product")
	// for _, productParam := range data.Product {
	// 	_, err = db.Queries.TestInsertProduct(context.Background(), productParam)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	fmt.Println("product Archive")
	// for _, productArchiveParam := range data.ProductArchive {
	// 	_, err = db.Queries.TestInsertProductArchive(context.Background(), productArchiveParam)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	fmt.Println("tag")
	// for _, tagParam := range data.Tag {
	// 	_, err = db.Queries.TestInsertTag(context.Background(), tagParam)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	fmt.Println("product tag")
	// for _, productTagParam := range data.ProductTag {
	// 	_, err = db.Queries.TestInsertProductTag(context.Background(), productTagParam)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	fmt.Println("coupon tag")
	// for _, couponTagParam := range data.CouponTag {
	// 	_, err = db.Queries.TestInsertCouponTag(context.Background(), couponTagParam)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	fmt.Println("cart")
	// for _, cartParam := range data.Cart {
	// 	_, err = db.Queries.TestInsertCart(context.Background(), cartParam)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	fmt.Println("order")
	// for _, orderParam := range data.Order {
	// 	_, err = db.Queries.TestInsertOrder(context.Background(), orderParam)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	fmt.Println("order detail")
	for _, orderDetailParam := range data.OrderDetail {
		_, err = db.Queries.TestInsertOrderDetail(context.Background(), orderDetailParam)
		if err != nil {
			log.Fatal(err)
		}
	}
}
