package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jykuo-love-shiritori/twp/db"
	"github.com/jykuo-love-shiritori/twp/minio"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	var err error
	pg, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	mc, err := minio.NewMINIO()
	if err != nil {
		log.Fatal(err)
	}

	TestInsertData(pg, mc)
	// TestDeleteData(pg, mc)
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
	OrderHistory   []db.TestInsertOrderHistoryParams   `json:"order_history"`
	OrderDetail    []db.TestInsertOrderDetailParams    `json:"order_detail"`
}

func TestInsertData(pg *db.DB, mc *minio.MC) {

	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var data testTable
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatal(err)
	}
	if err = jsonFile.Close(); err != nil {
		log.Fatal(err)
	}

	for _, userParam := range data.User {
		hashNewPassword, err := bcrypt.GenerateFromPassword([]byte(userParam.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		userParam.ImageID = userParam.ImageID + ".jpg"
		_, err = mc.PutFileByPath(context.Background(), "./test.jpg", userParam.ImageID)
		if err != nil {
			log.Fatal(err)
		}
		userParam.Password = string(hashNewPassword)
		_, err = pg.Queries.TestInsertUser(context.Background(), userParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestUser success")
	for _, shopParam := range data.Shop {
		shopParam.ImageID = shopParam.ImageID + ".jpg"
		_, err := mc.PutFileByPath(context.Background(), "./test.jpg", shopParam.ImageID)
		if err != nil {
			log.Fatal(err)
		}
		_, err = pg.Queries.TestInsertShop(context.Background(), shopParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestShop success")
	for _, couponParam := range data.Coupon {
		_, err = pg.Queries.TestInsertCoupon(context.Background(), couponParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestCoupon success")
	for _, productParam := range data.Product {
		productParam.ImageID = productParam.ImageID + ".jpg"
		_, err := mc.PutFileByPath(context.Background(), "./test.jpg", productParam.ImageID)
		if err != nil {
			log.Fatal(err)
		}
		_, err = pg.Queries.TestInsertProduct(context.Background(), productParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestProduct success")
	for _, productArchiveParam := range data.ProductArchive {
		productArchiveParam.ImageID = productArchiveParam.ImageID + ".jpg"
		_, err := mc.PutFileByPath(context.Background(), "./test.jpg", productArchiveParam.ImageID)
		if err != nil {
			log.Fatal(err)
		}
		_, err = pg.Queries.TestInsertProductArchive(context.Background(), productArchiveParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestProductArchive success")
	for _, tagParam := range data.Tag {
		_, err = pg.Queries.TestInsertTag(context.Background(), tagParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestTag success")
	for _, productTagParam := range data.ProductTag {
		_, err = pg.Queries.TestInsertProductTag(context.Background(), productTagParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestProductTag success")
	for _, couponTagParam := range data.CouponTag {
		_, err = pg.Queries.TestInsertCouponTag(context.Background(), couponTagParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestCouponTag success")
	for _, cartParam := range data.Cart {
		_, err = pg.Queries.TestInsertCart(context.Background(), cartParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestCart success")
	for _, orderParam := range data.OrderHistory {
		orderParam.ImageID = orderParam.ImageID + ".jpg"
		_, err := mc.PutFileByPath(context.Background(), "./test.jpg", orderParam.ImageID)
		if err != nil {
			log.Fatal(err)
		}
		_, err = pg.Queries.TestInsertOrderHistory(context.Background(), orderParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestOrderHistory success")
	for _, orderDetailParam := range data.OrderDetail {
		_, err = pg.Queries.TestInsertOrderDetail(context.Background(), orderDetailParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestOrderDetail success")
	for _, cartProductParam := range data.CartProduct {
		_, err = pg.Queries.TestInsertCartProduct(context.Background(), cartProductParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestCartProduct success")
	for _, cartCouponParam := range data.CartCoupon {
		_, err = pg.Queries.TestInsertCartCoupon(context.Background(), cartCouponParam)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("InsertTestCartProduct success")
}

func TestDeleteData(pg *db.DB, mc *minio.MC) {

	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var data testTable
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatal(err)
	}
	if err = jsonFile.Close(); err != nil {
		log.Fatal(err)
	}
	for _, couponParam := range data.Coupon {
		_, err = pg.Queries.TestDeleteCouponById(context.Background(), couponParam.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("DeleteTestCoupon success")

	for _, productParam := range data.Product {
		err := mc.RemoveFile(context.Background(), productParam.ImageID)
		if err != nil {
			log.Fatal(err)
		}
		_, err = pg.Queries.TestDeleteProductById(context.Background(), productParam.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("DeleteTestProduct success")
	for _, orderDetailParam := range data.OrderDetail {
		_, err = pg.Queries.TestDeleteOrderDetailByOrderId(context.Background(), orderDetailParam.OrderID)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("DeleteTestOrderDetail success")

	for _, tagParam := range data.Tag {
		_, err = pg.Queries.TestDeleteTagById(context.Background(), tagParam.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("DeleteTestOrderDetail success")

	for _, orderParam := range data.OrderHistory {
		err := mc.RemoveFile(context.Background(), orderParam.ImageID)
		if err != nil {
			log.Fatal(err)
		}
		_, err = pg.Queries.TestDeleteOrderById(context.Background(), orderParam.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("DeleteTestOrderHistory success")
	for _, ProductArchiveParam := range data.ProductArchive {

		err := mc.RemoveFile(context.Background(), ProductArchiveParam.ImageID)
		if err != nil {
			log.Fatal(err)
		}
		_, err = pg.Queries.TestDeleteProductArchiveByIdVersion(context.Background(), db.TestDeleteProductArchiveByIdVersionParams{ID: ProductArchiveParam.ID, Version: ProductArchiveParam.Version})
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("DeleteTestProductArchive success")
	for _, shopParam := range data.Shop {
		err := mc.RemoveFile(context.Background(), shopParam.ImageID)
		if err != nil {
			log.Fatal(err)
		}
		_, err = pg.Queries.TestDeleteShopById(context.Background(), shopParam.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("DeleteTestShop success")

	for _, userParam := range data.User {
		err := mc.RemoveFile(context.Background(), userParam.ImageID)
		if err != nil {
			log.Fatal(err)
		}
		_, err = pg.Queries.TestDeleteUserById(context.Background(), userParam.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("DeleteTestUser success")

}
