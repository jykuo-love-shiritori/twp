// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: buyer.sql

package db

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgtype"
)

const addCouponToCart = `-- name: AddCouponToCart :execrows
INSERT INTO "cart_coupon"("cart_id", "coupon_id")
SELECT C."id",
    CO."id"
FROM "cart" AS C,
    "user" AS U,
    "coupon" AS CO
WHERE U."username" = $1
    AND C."user_id" = U."id"
    AND C."id" = $3
    AND (
        CO."scope" = 'global'
        OR (
            CO."scope" = 'shop'
            AND CO."shop_id" = C."shop_id"
        )
    )
    AND NOW() BETWEEN CO."start_date" AND CO."expire_date"
    AND NOT EXISTS (
        SELECT 1
        FROM "cart_coupon" AS CC
        WHERE CC."cart_id" = C."id"
            AND CC."coupon_id" = $2
    )
    AND CO."id" = $2
`

type AddCouponToCartParams struct {
	Username string `json:"username"`
	CouponID int32  `json:"coupon_id"`
	CartID   int32  `json:"cart_id" param:"cart_id"`
}

func (q *Queries) AddCouponToCart(ctx context.Context, arg AddCouponToCartParams) (int64, error) {
	result, err := q.db.Exec(ctx, addCouponToCart, arg.Username, arg.CouponID, arg.CartID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const addProductToCart = `-- name: AddProductToCart :one
WITH valid_product AS (
    SELECT P."id" AS product_id,
        S."id" AS shop_id
    FROM "product" P,
        "shop" S
    WHERE P."shop_id" = S."id"
        AND P."id" = $2
        AND P."enabled" = TRUE
        AND P."stock" >= $3
),
new_cart AS (
    INSERT INTO "cart"("user_id", "shop_id")
    SELECT U."id",
        S."id"
    FROM "user" AS U,
        "shop" AS S,
        "product" AS P
    WHERE U."username" = $1
        AND S."id" = P."shop_id"
        AND P."id" = $2
        AND NOT EXISTS (
            SELECT 1
            FROM "cart" AS C
            WHERE C."user_id" = U."id"
                AND C."shop_id" = S."id"
        )
    RETURNING id, user_id, shop_id
),
existed_cart AS (
    SELECT C."id"
    FROM "cart" AS C,
        "user" AS U,
        "shop" AS S
    WHERE U."username" = $1
        AND C."user_id" = U."id"
        AND C."shop_id" = S."id"
        AND S."id" =(
            SELECT "shop_id"
            FROM valid_product
        )
),
cart_id AS (
    SELECT "id"
    FROM new_cart
    UNION ALL
    SELECT "id"
    FROM existed_cart
),
insert_product AS (
    INSERT INTO "cart_product"("cart_id", "product_id", "quantity")
    SELECT (
            SELECT "id"
            FROM cart_id
        ),
        (
            SELECT "product_id"
            FROM valid_product
        ),
        $3 ON CONFLICT ("cart_id", "product_id") DO
    UPDATE
    SET "quantity" = "cart_product"."quantity" + $3
    RETURNING cart_id, product_id, quantity
)
SELECT COALESCE(SUM(CP."quantity"), 0) + $3 AS total_quantity
FROM "cart_product" AS CP,
    "cart" AS C,
    "user" AS U
WHERE U."username" = $1
    AND C."user_id" = U."id"
    AND CP."cart_id" = C."id"
`

type AddProductToCartParams struct {
	Username string `json:"username"`
	ID       int32  `json:"id" param:"id"`
	Quantity int32  `json:"quantity"`
}

// check product enabled ⬆️
func (q *Queries) AddProductToCart(ctx context.Context, arg AddProductToCartParams) (int64, error) {
	row := q.db.QueryRow(ctx, addProductToCart, arg.Username, arg.ID, arg.Quantity)
	var total_quantity int64
	err := row.Scan(&total_quantity)
	return total_quantity, err
}

const checkout = `-- name: Checkout :exec
WITH insert_order AS (
    INSERT INTO "order_history"(
            "user_id",
            "shop_id",
            "shipment",
            "total_price",
            "status"
        )
    SELECT U."id",
        S."id",
        $2,
        $3,
        'paid'
    FROM "user" AS U,
        "shop" AS S,
        "cart" AS C
    WHERE U."username" = $1
        AND U."id" = C."user_id"
        AND C."id" = $4
        AND S."id" = C."shop_id"
    RETURNING "id"
),
delete_cart AS (
    DELETE FROM "cart" AS C
    WHERE C."id" = $4
),
add_sales AS (
    UPDATE "product" AS P
    SET "sales" = "sales" + CP."quantity",
        "stock" = "stock" - CP."quantity"
    FROM "cart_product" AS CP
    WHERE CP."cart_id" = $4
        AND CP."product_id" = P."id"
        AND CP."quantity" <= P."stock"
)
INSERT INTO "order_detail"(
        "order_id",
        "product_id",
        "product_version",
        "quantity"
    )
SELECT (
        SELECT "id"
        FROM insert_order
    ),
    CP."product_id",
    P."version",
    CP."quantity"
FROM "cart_product" AS CP,
    "product" AS P,
    "cart" AS C,
    "user" AS U
WHERE C."id" = CP."cart_id"
    AND CP."product_id" = P."id"
    AND C."id" = $4
    AND C."user_id" = U."id"
`

type CheckoutParams struct {
	Username   string `json:"username"`
	Shipment   int32  `json:"shipment"`
	TotalPrice int32  `json:"total_price"`
	CartID     int32  `json:"cart_id" param:"cart_id"`
}

func (q *Queries) Checkout(ctx context.Context, arg CheckoutParams) error {
	_, err := q.db.Exec(ctx, checkout,
		arg.Username,
		arg.Shipment,
		arg.TotalPrice,
		arg.CartID,
	)
	return err
}

const deleteCouponFromCart = `-- name: DeleteCouponFromCart :execrows
DELETE FROM "cart_coupon" AS CC USING "cart" AS C,
    "user" AS U
WHERE U."username" = $1
    AND C."user_id" = U."id"
    AND C."id" = CC."cart_id"
    AND C."id" = $3
    AND CC."coupon_id" = $2
`

type DeleteCouponFromCartParams struct {
	Username string `json:"username"`
	CouponID int32  `json:"coupon_id"`
	CartID   int32  `json:"cart_id" param:"cart_id"`
}

func (q *Queries) DeleteCouponFromCart(ctx context.Context, arg DeleteCouponFromCartParams) (int64, error) {
	result, err := q.db.Exec(ctx, deleteCouponFromCart, arg.Username, arg.CouponID, arg.CartID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const deleteEmptyCart = `-- name: DeleteEmptyCart :exec
WITH valid_cart AS (
    SELECT C."id"
    FROM "cart" C
        JOIN "user" u ON u."id" = C."user_id"
    WHERE u."username" = $1
        AND C."id" = $2
        AND NOT EXISTS (
            SELECT 1
            FROM "cart_product" CP
            WHERE CP."cart_id" = C."id"
        )
),
delete_coupon AS (
    DELETE FROM "cart_coupon" AS CC
    WHERE "cart_id" =(
            SELECT "id"
            FROM valid_cart
        )
)
DELETE FROM "cart" AS C
WHERE "id" =(
        SELECT "id"
        FROM valid_cart
    )
`

type DeleteEmptyCartParams struct {
	Username string `json:"username"`
	CartID   int32  `json:"cart_id" param:"cart_id"`
}

func (q *Queries) DeleteEmptyCart(ctx context.Context, arg DeleteEmptyCartParams) error {
	_, err := q.db.Exec(ctx, deleteEmptyCart, arg.Username, arg.CartID)
	return err
}

const deleteProductFromCart = `-- name: DeleteProductFromCart :one
WITH valid_cart AS (
    SELECT C."id"
    FROM "cart" C
        JOIN "user" u ON u."id" = C."user_id"
    WHERE u."username" = $1
        AND C."id" = $2
),
deleted_products AS (
    DELETE FROM "cart_product" CP
    WHERE "cart_id" =(
            SELECT "id"
            FROM valid_cart
        )
        AND CP."product_id" = $3
    RETURNING cart_id, product_id, quantity
)
SELECT EXISTS (
        SELECT 1
        FROM deleted_products
    )
`

type DeleteProductFromCartParams struct {
	Username  string `json:"username"`
	CartID    int32  `json:"cart_id" param:"cart_id"`
	ProductID int32  `json:"product_id" param:"id"`
}

func (q *Queries) DeleteProductFromCart(ctx context.Context, arg DeleteProductFromCartParams) (bool, error) {
	row := q.db.QueryRow(ctx, deleteProductFromCart, arg.Username, arg.CartID, arg.ProductID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getCart = `-- name: GetCart :many
SELECT C."id",
    S."seller_name",
    S."image_id" AS "shop_image_url",
    S."name" AS "shop_name"
FROM "cart" AS C,
    "user" AS U,
    "shop" AS S
WHERE U."username" = $1
    AND U."id" = C."user_id"
    AND C."shop_id" = S."id"
`

type GetCartRow struct {
	ID           int32  `json:"id" param:"cart_id"`
	SellerName   string `json:"seller_name" param:"seller_name"`
	ShopImageUrl string `json:"shop_image_url" swaggertype:"string"`
	ShopName     string `form:"name" json:"shop_name"`
}

func (q *Queries) GetCart(ctx context.Context, username string) ([]GetCartRow, error) {
	rows, err := q.db.Query(ctx, getCart, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCartRow{}
	for rows.Next() {
		var i GetCartRow
		if err := rows.Scan(
			&i.ID,
			&i.SellerName,
			&i.ShopImageUrl,
			&i.ShopName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCouponDetail = `-- name: GetCouponDetail :one
SELECT "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date"
FROM "coupon"
WHERE "id" = $1
`

type GetCouponDetailRow struct {
	ID          int32              `json:"id" param:"id"`
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) GetCouponDetail(ctx context.Context, id int32) (GetCouponDetailRow, error) {
	row := q.db.QueryRow(ctx, getCouponDetail, id)
	var i GetCouponDetailRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Scope,
		&i.Name,
		&i.Description,
		&i.Discount,
		&i.StartDate,
		&i.ExpireDate,
	)
	return i, err
}

const getCouponTag = `-- name: GetCouponTag :many
SELECT "tag_id"
FROM "coupon_tag" AS CT,
    "tag" AS T
WHERE CT."coupon_id" = $1
    AND CT."tag_id" = T."id"
`

func (q *Queries) GetCouponTag(ctx context.Context, couponID int32) ([]int32, error) {
	rows, err := q.db.Query(ctx, getCouponTag, couponID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var tag_id int32
		if err := rows.Scan(&tag_id); err != nil {
			return nil, err
		}
		items = append(items, tag_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCouponsFromCart = `-- name: GetCouponsFromCart :many
WITH delete_expire_coupons AS (
    DELETE FROM "cart_coupon" AS CC USING "coupon" AS CO,
        "cart" AS C,
        "user" AS U
    WHERE U."username" = $1
        AND C."user_id" = U."id"
        AND C."id" = CC."cart_id"
        AND C."id" = $2
        AND CC."coupon_id" = CO."id"
        AND NOW() > CO."expire_date"
)
SELECT CO."id",
    CO."name",
    CO."type",
    CO."scope",
    CO."description",
    CO."discount"
FROM "cart_coupon" AS CC,
    "coupon" AS CO,
    "cart" AS C,
    "user" AS U
WHERE U."username" = $1
    AND C."user_id" = U."id"
    AND C."id" = CC."cart_id"
    AND C."id" = $2
    AND CC."coupon_id" = CO."id"
`

type GetCouponsFromCartParams struct {
	Username string `json:"username"`
	CartID   int32  `json:"cart_id" param:"cart_id"`
}

type GetCouponsFromCartRow struct {
	ID          int32          `json:"id" param:"id"`
	Name        string         `json:"name"`
	Type        CouponType     `json:"type"`
	Scope       CouponScope    `json:"scope"`
	Description string         `json:"description"`
	Discount    pgtype.Numeric `json:"discount" swaggertype:"number"`
}

func (q *Queries) GetCouponsFromCart(ctx context.Context, arg GetCouponsFromCartParams) ([]GetCouponsFromCartRow, error) {
	rows, err := q.db.Query(ctx, getCouponsFromCart, arg.Username, arg.CartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCouponsFromCartRow{}
	for rows.Next() {
		var i GetCouponsFromCartRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Scope,
			&i.Description,
			&i.Discount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCreditCard = `-- name: GetCreditCard :one
SELECT "credit_card"
FROM "user"
WHERE "username" = $1
`

func (q *Queries) GetCreditCard(ctx context.Context, username string) (json.RawMessage, error) {
	row := q.db.QueryRow(ctx, getCreditCard, username)
	var credit_card json.RawMessage
	err := row.Scan(&credit_card)
	return credit_card, err
}

const getOrderDetail = `-- name: GetOrderDetail :many
SELECT O."product_id",
    P."name",
    P."description",
    P."price",
    P."image_id" AS "image_url",
    O."quantity"
FROM "order_detail" AS O,
    "product_archive" AS P
WHERE O."order_id" = $1
    AND O."product_id" = P."id"
    AND O."product_version" = P."version"
`

type GetOrderDetailRow struct {
	ProductID   int32          `json:"product_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       pgtype.Numeric `json:"price" swaggertype:"number"`
	ImageUrl    string         `json:"image_url"`
	Quantity    int32          `json:"quantity"`
}

func (q *Queries) GetOrderDetail(ctx context.Context, orderID int32) ([]GetOrderDetailRow, error) {
	rows, err := q.db.Query(ctx, getOrderDetail, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrderDetailRow{}
	for rows.Next() {
		var i GetOrderDetailRow
		if err := rows.Scan(
			&i.ProductID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.ImageUrl,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderHistory = `-- name: GetOrderHistory :many
SELECT O."id",
    s."name" AS "shop_name",
    s."image_id" AS "shop_image_url",
    OP."thumbnail_url",
    OP."product_name",
    O."shipment",
    O."total_price",
    O."status",
    O."created_at"
FROM "order_history" AS O
    INNER JOIN "user" AS U ON U."id" = O."user_id"
    INNER JOIN "shop" AS S ON O."shop_id" = S."id"
    LEFT JOIN (
        SELECT OD."order_id",
            PA."name" AS "product_name",
            PA."image_id" AS "thumbnail_url",
            ROW_NUMBER() OVER (
                PARTITION BY OD."order_id"
                ORDER BY PA."price" DESC
            ) AS rn
        FROM "order_detail" AS OD
            INNER JOIN "product_archive" AS PA ON OD."product_id" = PA."id"
            AND OD."product_version" = PA."version"
        ORDER BY PA."price" DESC
    ) AS OP ON O."id" = OP."order_id"
    AND OP.rn = 1
WHERE U."username" = $1
ORDER BY O."created_at" ASC OFFSET $2
LIMIT $3
`

type GetOrderHistoryParams struct {
	Username string `json:"username"`
	Offset   int64  `json:"offset"`
	Limit    int64  `json:"limit"`
}

type GetOrderHistoryRow struct {
	ID           int32              `json:"id" param:"id"`
	ShopName     string             `form:"name" json:"shop_name"`
	ShopImageUrl string             `json:"shop_image_url" swaggertype:"string"`
	ThumbnailUrl string             `json:"thumbnail_url"`
	ProductName  string             `json:"product_name"`
	Shipment     int32              `json:"shipment"`
	TotalPrice   int32              `json:"total_price"`
	Status       OrderStatus        `json:"status"`
	CreatedAt    pgtype.Timestamptz `json:"created_at" swaggertype:"string"`
}

func (q *Queries) GetOrderHistory(ctx context.Context, arg GetOrderHistoryParams) ([]GetOrderHistoryRow, error) {
	rows, err := q.db.Query(ctx, getOrderHistory, arg.Username, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrderHistoryRow{}
	for rows.Next() {
		var i GetOrderHistoryRow
		if err := rows.Scan(
			&i.ID,
			&i.ShopName,
			&i.ShopImageUrl,
			&i.ThumbnailUrl,
			&i.ProductName,
			&i.Shipment,
			&i.TotalPrice,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderInfo = `-- name: GetOrderInfo :one
SELECT O."id",
    S."name" AS "shop_name",
    S."image_id" AS "shop_image_url",
    "shipment",
    "total_price",
    "status",
    "created_at",
    (T."subtotal" + "shipment" - "total_price") AS "discount"
FROM "order_history" AS O,
    "order_detail" AS D,
    "user" AS U,
    "shop" AS S,
    (
        SELECT SUM(P."price" * D."quantity")::INT AS "subtotal"
        FROM "order_detail" AS D,
            "product_archive" AS P
        WHERE D."order_id" = $1
            AND D."product_id" = P."id"
            AND D."product_version" = P."version"
    ) AS T
WHERE U."username" = $2
    AND O."id" = $1
    AND O."user_id" = U."id"
    AND O."id" = D."order_id"
    AND O."shop_id" = S."id"
`

type GetOrderInfoParams struct {
	OrderID  int32  `json:"order_id" param:"id"`
	Username string `json:"username"`
}

type GetOrderInfoRow struct {
	ID           int32              `json:"id" param:"id"`
	ShopName     string             `form:"name" json:"shop_name"`
	ShopImageUrl string             `json:"shop_image_url" swaggertype:"string"`
	Shipment     int32              `json:"shipment"`
	TotalPrice   int32              `json:"total_price"`
	Status       OrderStatus        `json:"status"`
	CreatedAt    pgtype.Timestamptz `json:"created_at" swaggertype:"string"`
	Discount     int32              `json:"discount"`
}

func (q *Queries) GetOrderInfo(ctx context.Context, arg GetOrderInfoParams) (GetOrderInfoRow, error) {
	row := q.db.QueryRow(ctx, getOrderInfo, arg.OrderID, arg.Username)
	var i GetOrderInfoRow
	err := row.Scan(
		&i.ID,
		&i.ShopName,
		&i.ShopImageUrl,
		&i.Shipment,
		&i.TotalPrice,
		&i.Status,
		&i.CreatedAt,
		&i.Discount,
	)
	return i, err
}

const getProductFromCartOrderByPriceDesc = `-- name: GetProductFromCartOrderByPriceDesc :many
SELECT "product_id",
    "name",
    "image_id" AS "image_url",
    "price",
    "quantity",
    "stock",
    "enabled"
FROM "cart_product" AS C,
    "product" AS P
WHERE "cart_id" = $1
    AND C."product_id" = P."id"
ORDER BY "price" DESC
`

type GetProductFromCartOrderByPriceDescRow struct {
	ProductID int32          `json:"product_id" param:"id"`
	Name      string         `form:"name" json:"name"`
	ImageUrl  string         `json:"image_url"`
	Price     pgtype.Numeric `json:"price" swaggertype:"number"`
	Quantity  int32          `json:"quantity"`
	Stock     int32          `form:"stock" json:"stock"`
	Enabled   bool           `form:"enabled" json:"enabled"`
}

func (q *Queries) GetProductFromCartOrderByPriceDesc(ctx context.Context, cartID int32) ([]GetProductFromCartOrderByPriceDescRow, error) {
	rows, err := q.db.Query(ctx, getProductFromCartOrderByPriceDesc, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductFromCartOrderByPriceDescRow{}
	for rows.Next() {
		var i GetProductFromCartOrderByPriceDescRow
		if err := rows.Scan(
			&i.ProductID,
			&i.Name,
			&i.ImageUrl,
			&i.Price,
			&i.Quantity,
			&i.Stock,
			&i.Enabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductTag = `-- name: GetProductTag :many
SELECT "tag_id"
FROM "product_tag" AS PT,
    "tag" AS T
WHERE PT."product_id" = $1
    AND PT."tag_id" = T."id"
`

// returning the number of products in any cart for US-SC-2 in SRS ⬆️
func (q *Queries) GetProductTag(ctx context.Context, productID int32) ([]int32, error) {
	rows, err := q.db.Query(ctx, getProductTag, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var tag_id int32
		if err := rows.Scan(&tag_id); err != nil {
			return nil, err
		}
		items = append(items, tag_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSortedCouponsFromCart = `-- name: GetSortedCouponsFromCart :many
SELECT C."id",
    C."name",
    "type",
    "scope",
    "description",
    "discount",
    "expire_date"
FROM "cart_coupon" AS CC,
    "coupon" AS C,
    "cart" AS 🛒,
    "user" AS U
WHERE U."username" = $1
    AND U."id" = 🛒."user_id"
    AND 🛒."id" = $2
    AND CC."cart_id" = 🛒."id"
    AND CC."coupon_id" = C."id"
    AND NOW() BETWEEN C."start_date" AND C."expire_date"
ORDER BY "type" ASC,
    "discount" DESC
`

type GetSortedCouponsFromCartParams struct {
	Username string `json:"username"`
	CartID   int32  `json:"cart_id" param:"cart_id"`
}

type GetSortedCouponsFromCartRow struct {
	ID          int32              `json:"id" param:"id"`
	Name        string             `json:"name"`
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) GetSortedCouponsFromCart(ctx context.Context, arg GetSortedCouponsFromCartParams) ([]GetSortedCouponsFromCartRow, error) {
	rows, err := q.db.Query(ctx, getSortedCouponsFromCart, arg.Username, arg.CartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSortedCouponsFromCartRow{}
	for rows.Next() {
		var i GetSortedCouponsFromCartRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Scope,
			&i.Description,
			&i.Discount,
			&i.ExpireDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSortedUsableCoupons = `-- name: GetSortedUsableCoupons :many
SELECT C."id",
    C."name",
    "type",
    "scope",
    "description",
    "discount",
    "expire_date"
FROM "coupon" AS C,
    "cart" AS 🛒,
    "user" AS U
WHERE U."username" = $1
    AND U."id" = 🛒."user_id"
    AND 🛒."id" = $2
    AND (
        C."scope" = 'global'
        OR (
            C."scope" = 'shop'
            AND C."shop_id" = 🛒."shop_id"
        )
    )
    AND NOW() BETWEEN C."start_date" AND C."expire_date"
    AND NOT EXISTS (
        SELECT 1
        FROM "cart_coupon" AS CC
        WHERE CC."cart_id" = 🛒."id"
            AND CC."coupon_id" = C."id"
    )
ORDER BY "type" ASC,
    "discount" DESC
`

type GetSortedUsableCouponsParams struct {
	Username string `json:"username"`
	CartID   int32  `json:"cart_id" param:"cart_id"`
}

type GetSortedUsableCouponsRow struct {
	ID          int32              `json:"id" param:"id"`
	Name        string             `json:"name"`
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) GetSortedUsableCoupons(ctx context.Context, arg GetSortedUsableCouponsParams) ([]GetSortedUsableCouponsRow, error) {
	rows, err := q.db.Query(ctx, getSortedUsableCoupons, arg.Username, arg.CartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSortedUsableCouponsRow{}
	for rows.Next() {
		var i GetSortedUsableCouponsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Scope,
			&i.Description,
			&i.Discount,
			&i.ExpireDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrderStatus = `-- name: UpdateOrderStatus :execrows
UPDATE "order_history"
SET "status" = $1
WHERE "id" = $2
    AND "user_id" =(
        SELECT "id"
        FROM "user"
        WHERE "username" = $3
    )
    AND "status" = 'delivered'
`

type UpdateOrderStatusParams struct {
	Status   NullOrderStatus `json:"status"`
	ID       pgtype.Int4     `json:"id"`
	Username string          `json:"username"`
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) (int64, error) {
	result, err := q.db.Exec(ctx, updateOrderStatus, arg.Status, arg.ID, arg.Username)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const updateProductFromCart = `-- name: UpdateProductFromCart :execrows
UPDATE "cart_product" AS CP
SET "quantity" = $3
FROM "user" AS U,
    "cart" AS C,
    "product" AS P
WHERE U."username" = $4
    AND U."id" = C."user_id"
    AND CP."cart_id" = $1
    AND CP."product_id" = $2
    AND $3 <= P."stock"
    AND P."id" = $2
    AND P."enabled" = TRUE
    AND C."id" = CP."cart_id"
`

type UpdateProductFromCartParams struct {
	CartID    int32  `json:"cart_id"`
	ProductID int32  `json:"product_id" param:"id"`
	Quantity  int32  `json:"quantity"`
	Username  string `json:"username"`
}

func (q *Queries) UpdateProductFromCart(ctx context.Context, arg UpdateProductFromCartParams) (int64, error) {
	result, err := q.db.Exec(ctx, updateProductFromCart,
		arg.CartID,
		arg.ProductID,
		arg.Quantity,
		arg.Username,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const updateProductVersion = `-- name: UpdateProductVersion :exec
WITH check_version AS (
    SELECT EXISTS (
            SELECT 1
            FROM "product" P,
                "product_archive" PA
            WHERE P."id" = $1
                AND P."id" = PA."id"
                AND P."version" = PA."version"
                AND P."name" = PA."name"
                AND P."description" = PA."description"
                AND P."price" = PA."price"
                AND P."image_id" = PA."image_id"
        ) AS "version_existed",
        EXISTS (
            SELECT 1
            FROM "product_archive" PA
            WHERE PA."id" = $1
        ) AS "product_existed"
),
insert_archive AS (
    INSERT INTO "product_archive"(
            "id",
            "version",
            "name",
            "description",
            "price",
            "image_id"
        )
    SELECT P."id",
        P."version" + COALESCE(
            (
                SELECT 1
                FROM Check_version
                WHERE "product_existed" = TRUE
                    AND "version_existed" = FALSE
            ),
            0
        ),
        P."name",
        P."description",
        P."price",
        P."image_id"
    FROM "product" AS P,
        check_version AS CV
    WHERE P."id" = $1
        AND CV."version_existed" = FALSE
)
UPDATE "product"
SET "version" = "version" + 1
FROM check_version CV
WHERE "product"."id" = $1
    AND CV."version_existed" = FALSE
    AND CV."product_existed" = TRUE
`

func (q *Queries) UpdateProductVersion(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, updateProductVersion, id)
	return err
}

const validatePayment = `-- name: ValidatePayment :one
SELECT EXISTS (
        SELECT 1
        FROM "user"
        WHERE "username" = $1
            AND (
                "credit_card" = $2
                OR random() < 0.999 -- random validate payment
            )
    )
`

type ValidatePaymentParams struct {
	Username   string          `json:"username"`
	CreditCard json.RawMessage `json:"credit_card"`
}

func (q *Queries) ValidatePayment(ctx context.Context, arg ValidatePaymentParams) (bool, error) {
	row := q.db.QueryRow(ctx, validatePayment, arg.Username, arg.CreditCard)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const validateProductsInCart = `-- name: ValidateProductsInCart :one
SELECT NOT EXISTS (
        SELECT 1
        FROM "cart_product" AS CP
            JOIN "product" AS P ON CP."product_id" = P."id"
            JOIN "cart" AS C ON C."id" = CP."cart_id"
            JOIN "user" AS U ON C."user_id" = U."id"
        WHERE C."id" = $2
            AND U."username" = $1
            AND (
                P."enabled" = FALSE
                OR CP."quantity" > P."stock"
            )
    )
    AND EXISTS (
        SELECT 1
        FROM "cart" AS C,
            "user" AS U
        WHERE C."id" = $2
            AND U."username" = $1
            AND C."user_id" = U."id"
    )
`

type ValidateProductsInCartParams struct {
	Username string `json:"username"`
	CartID   int32  `json:"cart_id" param:"cart_id"`
}

func (q *Queries) ValidateProductsInCart(ctx context.Context, arg ValidateProductsInCartParams) (pgtype.Bool, error) {
	row := q.db.QueryRow(ctx, validateProductsInCart, arg.Username, arg.CartID)
	var column_1 pgtype.Bool
	err := row.Scan(&column_1)
	return column_1, err
}
