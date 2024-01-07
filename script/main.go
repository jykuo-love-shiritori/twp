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
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
)

type loadDataTable struct {
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

func main() {
	pg, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	mc, err := minio.NewMINIO()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:                 "twp-CLI",
		Usage:                "A CLI application to load and delete data from a database",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:      "load",
				Usage:     "Load data into the database",
				UsageText: "load path/data.json",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return cli.ShowSubcommandHelp(c)
					}
					fileName := c.Args().Get(0)
					return LoadData(pg, mc, context.Background(), fileName)
				},
			},
			{
				Name:      "unload",
				Usage:     "Unload data from the database",
				UsageText: "unload or unload path/data.json",
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 {
						return ClearData(pg, context.Background())
					}
					fileName := c.Args().Get(0)
					return UnloadData(pg, mc, context.Background(), fileName)
				},
			},
		},
	}
	app.Suggest = true
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func LoadData(pg *db.DB, mc *minio.MC, ctx context.Context, filePath string) error {

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	var data loadDataTable
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return err
	}
	if err = jsonFile.Close(); err != nil {
		return err
	}

	for _, userParam := range data.User {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(userParam.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		userParam.Password = string(hashPassword)

		tx, err := pg.NewTx(ctx)
		if err != nil {
			return err
		}
		defer tx.Rollback(ctx) //nolint:errcheck

		userParam.ImageID, err = mc.PutFileByPath(ctx, userParam.ImageID)
		if err != nil {
			return err
		}
		_, err = pg.Queries.WithTx(tx).TestInsertUser(ctx, userParam)
		if err != nil {
			return err
		}
		err = tx.Commit(ctx)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert User Success")
	for _, shopParam := range data.Shop {
		tx, err := pg.NewTx(ctx)
		if err != nil {
			return err
		}
		defer tx.Rollback(ctx) //nolint:errcheck

		shopParam.ImageID, err = mc.PutFileByPath(ctx, shopParam.ImageID)
		if err != nil {
			return err
		}
		_, err = pg.Queries.WithTx(tx).TestInsertShop(ctx, shopParam)
		if err != nil {
			return err
		}
		err = tx.Commit(ctx)
		if err != nil {
			return err
		}

	}
	fmt.Println("Insert Shop Success")
	for _, couponParam := range data.Coupon {
		_, err = pg.Queries.TestInsertCoupon(ctx, couponParam)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert Coupon Success")
	for _, productParam := range data.Product {
		tx, err := pg.NewTx(ctx)
		if err != nil {
			return err
		}
		defer tx.Rollback(ctx) //nolint:errcheck

		productParam.ImageID, err = mc.PutFileByPath(ctx, productParam.ImageID)
		if err != nil {
			return err
		}
		_, err = pg.Queries.WithTx(tx).TestInsertProduct(ctx, productParam)
		if err != nil {
			return err
		}
		err = tx.Commit(ctx)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert Product Success")
	for _, productArchiveParam := range data.ProductArchive {
		tx, err := pg.NewTx(ctx)
		if err != nil {
			return err
		}
		defer tx.Rollback(ctx) //nolint:errcheck

		productArchiveParam.ImageID, err = mc.PutFileByPath(ctx, productArchiveParam.ImageID)
		if err != nil {
			return err
		}
		_, err = pg.Queries.WithTx(tx).TestInsertProductArchive(ctx, productArchiveParam)
		if err != nil {
			return err
		}
		err = tx.Commit(ctx)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert ProductArchive Success")
	for _, tagParam := range data.Tag {
		_, err = pg.Queries.TestInsertTag(ctx, tagParam)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert Tag Success")
	for _, productTagParam := range data.ProductTag {
		_, err = pg.Queries.TestInsertProductTag(ctx, productTagParam)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert ProductTag Success")
	for _, couponTagParam := range data.CouponTag {
		_, err = pg.Queries.TestInsertCouponTag(ctx, couponTagParam)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert CouponTag Success")
	for _, cartParam := range data.Cart {
		_, err = pg.Queries.TestInsertCart(ctx, cartParam)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert Cart Success")
	for _, orderParam := range data.OrderHistory {
		_, err = pg.Queries.TestInsertOrderHistory(ctx, orderParam)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert OrderHistory Success")
	for _, orderDetailParam := range data.OrderDetail {
		_, err = pg.Queries.TestInsertOrderDetail(ctx, orderDetailParam)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert OrderDetail Success")
	for _, cartProductParam := range data.CartProduct {
		_, err = pg.Queries.TestInsertCartProduct(ctx, cartProductParam)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert CartProduct Success")
	for _, cartCouponParam := range data.CartCoupon {
		_, err = pg.Queries.TestInsertCartCoupon(ctx, cartCouponParam)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert CartProduct Success")
	return nil
}

func UnloadData(pg *db.DB, mc *minio.MC, ctx context.Context, filePath string) error {

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	var data loadDataTable
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return err
	}
	if err = jsonFile.Close(); err != nil {
		return err
	}
	for _, couponParam := range data.Coupon {
		_, err = pg.Queries.TestDeleteCouponById(ctx, couponParam.ID)
		if err != nil {
			return err
		}
	}
	fmt.Println("Delete Coupon Success")

	for _, productParam := range data.Product {
		product, err := pg.Queries.TestDeleteProductById(ctx, productParam.ID)
		if err != nil {
			return err
		}
		err = mc.RemoveFile(ctx, product.ImageID)
		if err != nil {
			return err
		}

	}
	fmt.Println("Delete Product Success")
	for _, orderDetailParam := range data.OrderDetail {
		_, err = pg.Queries.TestDeleteOrderDetailByOrderId(ctx, db.TestDeleteOrderDetailByOrderIdParams{
			OrderID:        orderDetailParam.OrderID,
			ProductID:      orderDetailParam.ProductID,
			ProductVersion: orderDetailParam.ProductVersion,
		})
		if err != nil {
			return err
		}
	}
	fmt.Println("Delete OrderDetail Success")

	for _, tagParam := range data.Tag {
		_, err = pg.Queries.TestDeleteTagById(ctx, tagParam.ID)
		if err != nil {
			return err
		}
	}
	fmt.Println("Delete Tag Success")

	for _, orderParam := range data.OrderHistory {
		_, err = pg.Queries.TestDeleteOrderById(ctx, orderParam.ID)
		if err != nil {
			return err
		}
	}
	fmt.Println("Delete OrderHistory Success")
	for _, ProductArchiveParam := range data.ProductArchive {

		productArchive, err := pg.Queries.TestDeleteProductArchiveByIdVersion(ctx, db.TestDeleteProductArchiveByIdVersionParams{ID: ProductArchiveParam.ID, Version: ProductArchiveParam.Version})
		if err != nil {
			return err
		}
		err = mc.RemoveFile(ctx, productArchive.ImageID)
		if err != nil {
			return err
		}
	}
	fmt.Println("Delete ProductArchive Success")
	for _, shopParam := range data.Shop {
		shop, err := pg.Queries.TestDeleteShopById(ctx, shopParam.ID)
		if err != nil {
			return err
		}
		err = mc.RemoveFile(ctx, shop.ImageID)
		if err != nil {
			return err
		}
	}
	fmt.Println("Delete Shop Success")

	for _, userParam := range data.User {
		user, err := pg.Queries.TestDeleteUserById(ctx, userParam.ID)
		if err != nil {
			return err
		}
		err = mc.RemoveFile(ctx, user.ImageID)
		if err != nil {
			return err
		}
	}
	fmt.Println("Delete User Success")
	return nil
}
func ClearData(pg *db.DB, ctx context.Context) error {
	err := pg.Queries.TestDeleteCoupon(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Delete Coupon Success")
	err = pg.Queries.TestDeleteProduct(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Delete Product Success")
	err = pg.Queries.TestDeleteOrderDetail(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Delete OrderDetail Success")
	err = pg.Queries.TestDeleteTag(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Delete Tag Success")
	err = pg.Queries.TestDeleteOrderHistory(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Delete OrderHistory Success")
	err = pg.Queries.TestDeleteProductArchive(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Delete ProductArchive Success")
	err = pg.Queries.TestDeleteShop(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Delete Shop Success")
	err = pg.Queries.TestDeleteUser(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Delete User Success")
	return nil
}
