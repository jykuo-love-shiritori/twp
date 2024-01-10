// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: test.sql

package db

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
)

const testDeleteCoupon = `-- name: TestDeleteCoupon :exec
DELETE FROM "coupon"
`

func (q *Queries) TestDeleteCoupon(ctx context.Context) error {
	_, err := q.db.Exec(ctx, testDeleteCoupon)
	return err
}

const testDeleteCouponById = `-- name: TestDeleteCouponById :one
DELETE FROM "coupon"
WHERE "id" = $1
RETURNING id, type, scope, shop_id, name, description, discount, start_date, expire_date
`

func (q *Queries) TestDeleteCouponById(ctx context.Context, id int32) (Coupon, error) {
	row := q.db.QueryRow(ctx, testDeleteCouponById, id)
	var i Coupon
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Scope,
		&i.ShopID,
		&i.Name,
		&i.Description,
		&i.Discount,
		&i.StartDate,
		&i.ExpireDate,
	)
	return i, err
}

const testDeleteOrderById = `-- name: TestDeleteOrderById :one
DELETE FROM "order_history"
WHERE "id" = $1
RETURNING id, user_id, shop_id, shipment, total_price, status, created_at
`

func (q *Queries) TestDeleteOrderById(ctx context.Context, id int32) (OrderHistory, error) {
	row := q.db.QueryRow(ctx, testDeleteOrderById, id)
	var i OrderHistory
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ShopID,
		&i.Shipment,
		&i.TotalPrice,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const testDeleteOrderDetail = `-- name: TestDeleteOrderDetail :exec
DELETE FROM "order_detail"
`

func (q *Queries) TestDeleteOrderDetail(ctx context.Context) error {
	_, err := q.db.Exec(ctx, testDeleteOrderDetail)
	return err
}

const testDeleteOrderDetailByOrderId = `-- name: TestDeleteOrderDetailByOrderId :one
DELETE FROM "order_detail"
WHERE "order_id" = $1
    AND "product_id" = $2
    AND "product_version" = $3
RETURNING order_id, product_id, product_version, quantity
`

type TestDeleteOrderDetailByOrderIdParams struct {
	OrderID        int32 `json:"order_id" param:"id"`
	ProductID      int32 `json:"product_id"`
	ProductVersion int32 `json:"product_version"`
}

func (q *Queries) TestDeleteOrderDetailByOrderId(ctx context.Context, arg TestDeleteOrderDetailByOrderIdParams) (OrderDetail, error) {
	row := q.db.QueryRow(ctx, testDeleteOrderDetailByOrderId, arg.OrderID, arg.ProductID, arg.ProductVersion)
	var i OrderDetail
	err := row.Scan(
		&i.OrderID,
		&i.ProductID,
		&i.ProductVersion,
		&i.Quantity,
	)
	return i, err
}

const testDeleteOrderHistory = `-- name: TestDeleteOrderHistory :exec
DELETE FROM "order_history"
`

func (q *Queries) TestDeleteOrderHistory(ctx context.Context) error {
	_, err := q.db.Exec(ctx, testDeleteOrderHistory)
	return err
}

const testDeleteProduct = `-- name: TestDeleteProduct :exec
DELETE FROM "product"
`

func (q *Queries) TestDeleteProduct(ctx context.Context) error {
	_, err := q.db.Exec(ctx, testDeleteProduct)
	return err
}

const testDeleteProductArchive = `-- name: TestDeleteProductArchive :exec
DELETE FROM "product_archive"
`

func (q *Queries) TestDeleteProductArchive(ctx context.Context) error {
	_, err := q.db.Exec(ctx, testDeleteProductArchive)
	return err
}

const testDeleteProductArchiveByIdVersion = `-- name: TestDeleteProductArchiveByIdVersion :one
DELETE FROM "product_archive"
WHERE "id" = $1
    AND "version" = $2
RETURNING id, version, name, description, price, image_id
`

type TestDeleteProductArchiveByIdVersionParams struct {
	ID      int32 `json:"id"`
	Version int32 `json:"version"`
}

func (q *Queries) TestDeleteProductArchiveByIdVersion(ctx context.Context, arg TestDeleteProductArchiveByIdVersionParams) (ProductArchive, error) {
	row := q.db.QueryRow(ctx, testDeleteProductArchiveByIdVersion, arg.ID, arg.Version)
	var i ProductArchive
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.ImageID,
	)
	return i, err
}

const testDeleteProductById = `-- name: TestDeleteProductById :one
DELETE FROM "product"
WHERE "id" = $1
RETURNING id, version, shop_id, name, description, price, image_id, expire_date, edit_date, stock, sales, enabled
`

func (q *Queries) TestDeleteProductById(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRow(ctx, testDeleteProductById, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.ShopID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.ImageID,
		&i.ExpireDate,
		&i.EditDate,
		&i.Stock,
		&i.Sales,
		&i.Enabled,
	)
	return i, err
}

const testDeleteShop = `-- name: TestDeleteShop :exec
DELETE FROM "shop"
`

func (q *Queries) TestDeleteShop(ctx context.Context) error {
	_, err := q.db.Exec(ctx, testDeleteShop)
	return err
}

const testDeleteShopById = `-- name: TestDeleteShopById :one
DELETE FROM "shop"
WHERE "id" = $1
RETURNING id, seller_name, image_id, name, description, enabled
`

func (q *Queries) TestDeleteShopById(ctx context.Context, id int32) (Shop, error) {
	row := q.db.QueryRow(ctx, testDeleteShopById, id)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.SellerName,
		&i.ImageID,
		&i.Name,
		&i.Description,
		&i.Enabled,
	)
	return i, err
}

const testDeleteTag = `-- name: TestDeleteTag :exec
DELETE FROM "tag"
`

func (q *Queries) TestDeleteTag(ctx context.Context) error {
	_, err := q.db.Exec(ctx, testDeleteTag)
	return err
}

const testDeleteTagById = `-- name: TestDeleteTagById :one
DELETE FROM "tag"
WHERE "id" = $1
RETURNING id, shop_id, name
`

func (q *Queries) TestDeleteTagById(ctx context.Context, id int32) (Tag, error) {
	row := q.db.QueryRow(ctx, testDeleteTagById, id)
	var i Tag
	err := row.Scan(&i.ID, &i.ShopID, &i.Name)
	return i, err
}

const testDeleteUser = `-- name: TestDeleteUser :exec
DELETE FROM "user"
`

func (q *Queries) TestDeleteUser(ctx context.Context) error {
	_, err := q.db.Exec(ctx, testDeleteUser)
	return err
}

const testDeleteUserById = `-- name: TestDeleteUserById :one
DELETE FROM "user"
WHERE "id" = $1
RETURNING "id",
    "username",
    "password",
    "name",
    "email",
    "address",
    "image_id",
    "role",
    "credit_card",
    "enabled"
`

type TestDeleteUserByIdRow struct {
	ID         int32           `json:"id" param:"id"`
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	Name       string          `form:"name" json:"name"`
	Email      string          `form:"email" json:"email"`
	Address    string          `form:"address" json:"address"`
	ImageID    string          `json:"image_id" swaggertype:"string"`
	Role       RoleType        `json:"role"`
	CreditCard json.RawMessage `json:"credit_card"`
	Enabled    bool            `json:"enabled"`
}

func (q *Queries) TestDeleteUserById(ctx context.Context, id int32) (TestDeleteUserByIdRow, error) {
	row := q.db.QueryRow(ctx, testDeleteUserById, id)
	var i TestDeleteUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.Email,
		&i.Address,
		&i.ImageID,
		&i.Role,
		&i.CreditCard,
		&i.Enabled,
	)
	return i, err
}

const testInsertCart = `-- name: TestInsertCart :one
INSERT INTO "cart" ("id", "user_id", "shop_id")
VALUES ($1, $2, $3)
RETURNING id, user_id, shop_id
`

type TestInsertCartParams struct {
	ID     int32 `json:"id" param:"cart_id"`
	UserID int32 `json:"user_id"`
	ShopID int32 `json:"shop_id"`
}

func (q *Queries) TestInsertCart(ctx context.Context, arg TestInsertCartParams) (Cart, error) {
	row := q.db.QueryRow(ctx, testInsertCart, arg.ID, arg.UserID, arg.ShopID)
	var i Cart
	err := row.Scan(&i.ID, &i.UserID, &i.ShopID)
	return i, err
}

const testInsertCartCoupon = `-- name: TestInsertCartCoupon :one
INSERT INTO "cart_coupon" ("cart_id", "coupon_id")
VALUES ($1, $2)
RETURNING cart_id, coupon_id
`

type TestInsertCartCouponParams struct {
	CartID   int32 `json:"cart_id"`
	CouponID int32 `json:"coupon_id"`
}

func (q *Queries) TestInsertCartCoupon(ctx context.Context, arg TestInsertCartCouponParams) (CartCoupon, error) {
	row := q.db.QueryRow(ctx, testInsertCartCoupon, arg.CartID, arg.CouponID)
	var i CartCoupon
	err := row.Scan(&i.CartID, &i.CouponID)
	return i, err
}

const testInsertCartProduct = `-- name: TestInsertCartProduct :one
INSERT INTO "cart_product" (
        "cart_id",
        "product_id",
        "quantity"
    )
VALUES ($1, $2, $3)
RETURNING cart_id, product_id, quantity
`

type TestInsertCartProductParams struct {
	CartID    int32 `json:"cart_id"`
	ProductID int32 `json:"product_id" param:"id"`
	Quantity  int32 `json:"quantity"`
}

func (q *Queries) TestInsertCartProduct(ctx context.Context, arg TestInsertCartProductParams) (CartProduct, error) {
	row := q.db.QueryRow(ctx, testInsertCartProduct, arg.CartID, arg.ProductID, arg.Quantity)
	var i CartProduct
	err := row.Scan(&i.CartID, &i.ProductID, &i.Quantity)
	return i, err
}

const testInsertCoupon = `-- name: TestInsertCoupon :one
INSERT INTO "coupon" (
        "id",
        "type",
        "scope",
        "shop_id",
        "name",
        "description",
        "discount",
        "start_date",
        "expire_date"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, type, scope, shop_id, name, description, discount, start_date, expire_date
`

type TestInsertCouponParams struct {
	ID          int32              `json:"id" param:"id"`
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	ShopID      pgtype.Int4        `json:"shop_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) TestInsertCoupon(ctx context.Context, arg TestInsertCouponParams) (Coupon, error) {
	row := q.db.QueryRow(ctx, testInsertCoupon,
		arg.ID,
		arg.Type,
		arg.Scope,
		arg.ShopID,
		arg.Name,
		arg.Description,
		arg.Discount,
		arg.StartDate,
		arg.ExpireDate,
	)
	var i Coupon
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Scope,
		&i.ShopID,
		&i.Name,
		&i.Description,
		&i.Discount,
		&i.StartDate,
		&i.ExpireDate,
	)
	return i, err
}

const testInsertCouponTag = `-- name: TestInsertCouponTag :one
INSERT INTO "coupon_tag" ("tag_id", "coupon_id")
VALUES ($1, $2)
RETURNING coupon_id, tag_id
`

type TestInsertCouponTagParams struct {
	TagID    int32 `json:"tag_id"`
	CouponID int32 `json:"coupon_id" param:"id"`
}

func (q *Queries) TestInsertCouponTag(ctx context.Context, arg TestInsertCouponTagParams) (CouponTag, error) {
	row := q.db.QueryRow(ctx, testInsertCouponTag, arg.TagID, arg.CouponID)
	var i CouponTag
	err := row.Scan(&i.CouponID, &i.TagID)
	return i, err
}

const testInsertOrderDetail = `-- name: TestInsertOrderDetail :one
INSERT INTO "order_detail" (
        "order_id",
        "product_id",
        "product_version",
        "quantity"
    )
VALUES ($1, $2, $3, $4)
RETURNING order_id, product_id, product_version, quantity
`

type TestInsertOrderDetailParams struct {
	OrderID        int32 `json:"order_id" param:"id"`
	ProductID      int32 `json:"product_id"`
	ProductVersion int32 `json:"product_version"`
	Quantity       int32 `json:"quantity"`
}

func (q *Queries) TestInsertOrderDetail(ctx context.Context, arg TestInsertOrderDetailParams) (OrderDetail, error) {
	row := q.db.QueryRow(ctx, testInsertOrderDetail,
		arg.OrderID,
		arg.ProductID,
		arg.ProductVersion,
		arg.Quantity,
	)
	var i OrderDetail
	err := row.Scan(
		&i.OrderID,
		&i.ProductID,
		&i.ProductVersion,
		&i.Quantity,
	)
	return i, err
}

const testInsertOrderHistory = `-- name: TestInsertOrderHistory :one
INSERT INTO "order_history" (
        "id",
        "user_id",
        "shop_id",
        "shipment",
        "total_price",
        "status"
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, user_id, shop_id, shipment, total_price, status, created_at
`

type TestInsertOrderHistoryParams struct {
	ID         int32       `json:"id" param:"id"`
	UserID     int32       `json:"user_id"`
	ShopID     int32       `json:"shop_id"`
	Shipment   int32       `json:"shipment"`
	TotalPrice int32       `json:"total_price"`
	Status     OrderStatus `json:"status"`
}

func (q *Queries) TestInsertOrderHistory(ctx context.Context, arg TestInsertOrderHistoryParams) (OrderHistory, error) {
	row := q.db.QueryRow(ctx, testInsertOrderHistory,
		arg.ID,
		arg.UserID,
		arg.ShopID,
		arg.Shipment,
		arg.TotalPrice,
		arg.Status,
	)
	var i OrderHistory
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ShopID,
		&i.Shipment,
		&i.TotalPrice,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const testInsertProduct = `-- name: TestInsertProduct :one
INSERT INTO "product" (
        "id",
        "version",
        "shop_id",
        "name",
        "description",
        "price",
        "image_id",
        "expire_date",
        "edit_date",
        "stock",
        "sales",
        "enabled"
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        NOW(),
        $9,
        $10,
        $11
    )
RETURNING id, version, shop_id, name, description, price, image_id, expire_date, edit_date, stock, sales, enabled
`

type TestInsertProductParams struct {
	ID          int32              `json:"id" param:"id"`
	Version     int32              `json:"version"`
	ShopID      int32              `json:"shop_id"`
	Name        string             `form:"name" json:"name"`
	Description string             `form:"description" json:"description"`
	Price       pgtype.Numeric     `json:"price" swaggertype:"number"`
	ImageID     string             `json:"image_id"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
	Stock       int32              `form:"stock" json:"stock"`
	Sales       int32              `json:"sales"`
	Enabled     bool               `form:"enabled" json:"enabled"`
}

func (q *Queries) TestInsertProduct(ctx context.Context, arg TestInsertProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, testInsertProduct,
		arg.ID,
		arg.Version,
		arg.ShopID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.ImageID,
		arg.ExpireDate,
		arg.Stock,
		arg.Sales,
		arg.Enabled,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.ShopID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.ImageID,
		&i.ExpireDate,
		&i.EditDate,
		&i.Stock,
		&i.Sales,
		&i.Enabled,
	)
	return i, err
}

const testInsertProductArchive = `-- name: TestInsertProductArchive :one
INSERT INTO "product_archive" (
        "id",
        "version",
        "name",
        "description",
        "price",
        "image_id"
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, version, name, description, price, image_id
`

type TestInsertProductArchiveParams struct {
	ID          int32          `json:"id"`
	Version     int32          `json:"version"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       pgtype.Numeric `json:"price" swaggertype:"number"`
	ImageID     string         `json:"image_id"`
}

func (q *Queries) TestInsertProductArchive(ctx context.Context, arg TestInsertProductArchiveParams) (ProductArchive, error) {
	row := q.db.QueryRow(ctx, testInsertProductArchive,
		arg.ID,
		arg.Version,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.ImageID,
	)
	var i ProductArchive
	err := row.Scan(
		&i.ID,
		&i.Version,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.ImageID,
	)
	return i, err
}

const testInsertProductTag = `-- name: TestInsertProductTag :one
INSERT INTO "product_tag" ("tag_id", "product_id")
VALUES ($1, $2)
RETURNING tag_id, product_id
`

type TestInsertProductTagParams struct {
	TagID     int32 `json:"tag_id"`
	ProductID int32 `json:"product_id" param:"id"`
}

func (q *Queries) TestInsertProductTag(ctx context.Context, arg TestInsertProductTagParams) (ProductTag, error) {
	row := q.db.QueryRow(ctx, testInsertProductTag, arg.TagID, arg.ProductID)
	var i ProductTag
	err := row.Scan(&i.TagID, &i.ProductID)
	return i, err
}

const testInsertShop = `-- name: TestInsertShop :one
INSERT INTO "shop" (
        "id",
        "seller_name",
        "name",
        "image_id",
        "description",
        "enabled"
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, seller_name, image_id, name, description, enabled
`

type TestInsertShopParams struct {
	ID          int32  `json:"id"`
	SellerName  string `json:"seller_name" param:"seller_name"`
	Name        string `form:"name" json:"name"`
	ImageID     string `json:"image_id" swaggertype:"string"`
	Description string `form:"description" json:"description"`
	Enabled     bool   `form:"enabled" json:"enabled"`
}

func (q *Queries) TestInsertShop(ctx context.Context, arg TestInsertShopParams) (Shop, error) {
	row := q.db.QueryRow(ctx, testInsertShop,
		arg.ID,
		arg.SellerName,
		arg.Name,
		arg.ImageID,
		arg.Description,
		arg.Enabled,
	)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.SellerName,
		&i.ImageID,
		&i.Name,
		&i.Description,
		&i.Enabled,
	)
	return i, err
}

const testInsertTag = `-- name: TestInsertTag :one
INSERT INTO "tag" ("id", "shop_id", "name")
VALUES ($1, $2, $3)
RETURNING id, shop_id, name
`

type TestInsertTagParams struct {
	ID     int32  `json:"id"`
	ShopID int32  `json:"shop_id"`
	Name   string `json:"name"`
}

func (q *Queries) TestInsertTag(ctx context.Context, arg TestInsertTagParams) (Tag, error) {
	row := q.db.QueryRow(ctx, testInsertTag, arg.ID, arg.ShopID, arg.Name)
	var i Tag
	err := row.Scan(&i.ID, &i.ShopID, &i.Name)
	return i, err
}

const testInsertUser = `-- name: TestInsertUser :one
INSERT INTO "user" (
        "id",
        "username",
        "password",
        "name",
        "email",
        "address",
        "image_id",
        "role",
        "credit_card",
        "enabled"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING "id",
    "username",
    "password",
    "name",
    "email",
    "address",
    "image_id",
    "role",
    "credit_card",
    "enabled"
`

type TestInsertUserParams struct {
	ID         int32           `json:"id" param:"id"`
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	Name       string          `form:"name" json:"name"`
	Email      string          `form:"email" json:"email"`
	Address    string          `form:"address" json:"address"`
	ImageID    string          `json:"image_id" swaggertype:"string"`
	Role       RoleType        `json:"role"`
	CreditCard json.RawMessage `json:"credit_card"`
	Enabled    bool            `json:"enabled"`
}

type TestInsertUserRow struct {
	ID         int32           `json:"id" param:"id"`
	Username   string          `json:"username"`
	Password   string          `json:"password"`
	Name       string          `form:"name" json:"name"`
	Email      string          `form:"email" json:"email"`
	Address    string          `form:"address" json:"address"`
	ImageID    string          `json:"image_id" swaggertype:"string"`
	Role       RoleType        `json:"role"`
	CreditCard json.RawMessage `json:"credit_card"`
	Enabled    bool            `json:"enabled"`
}

func (q *Queries) TestInsertUser(ctx context.Context, arg TestInsertUserParams) (TestInsertUserRow, error) {
	row := q.db.QueryRow(ctx, testInsertUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Name,
		arg.Email,
		arg.Address,
		arg.ImageID,
		arg.Role,
		arg.CreditCard,
		arg.Enabled,
	)
	var i TestInsertUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.Email,
		&i.Address,
		&i.ImageID,
		&i.Role,
		&i.CreditCard,
		&i.Enabled,
	)
	return i, err
}
