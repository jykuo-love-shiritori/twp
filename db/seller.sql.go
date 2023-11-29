// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: seller.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const haveTagName = `-- name: HaveTagName :one

SELECT EXISTS (
        SELECT t.id, shop_id, t.name, s.id, seller_name, image_id, s.name, description, enabled
        FROM "tag" t
            LEFT JOIN "shop" s ON "shop_id" = s.id
        WHERE
            s."seller_name" = $1
            AND t."name" = $2
    )
`

type HaveTagNameParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	Name       string `json:"name"`
}

func (q *Queries) HaveTagName(ctx context.Context, arg HaveTagNameParams) (bool, error) {
	row := q.db.QueryRow(ctx, haveTagName, arg.SellerName, arg.Name)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const sellerDeleteCoupon = `-- name: SellerDeleteCoupon :execrows

DELETE FROM "coupon" c
WHERE c."id" = $2 AND "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
`

type SellerDeleteCouponParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	ID         int32  `json:"id" param:"id"`
}

func (q *Queries) SellerDeleteCoupon(ctx context.Context, arg SellerDeleteCouponParams) (int64, error) {
	result, err := q.db.Exec(ctx, sellerDeleteCoupon, arg.SellerName, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const sellerDeleteCouponTag = `-- name: SellerDeleteCouponTag :execrows

DELETE FROM "coupon_tag" tp
WHERE EXISTS (
        SELECT 1
        FROM "coupon" c
            JOIN "shop" s ON s."id" = c."shop_id"
        WHERE
            s."seller_name" = $1
            AND c."id" = $3
    )
    AND "coupon_id" = $3
    AND "tag_id" = $2
`

type SellerDeleteCouponTagParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	TagID      int32  `json:"tag_id"`
	ID         int32  `json:"id" param:"id"`
}

func (q *Queries) SellerDeleteCouponTag(ctx context.Context, arg SellerDeleteCouponTagParams) (int64, error) {
	result, err := q.db.Exec(ctx, sellerDeleteCouponTag, arg.SellerName, arg.TagID, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const sellerDeleteProduct = `-- name: SellerDeleteProduct :execrows

DELETE FROM "product" p
WHERE "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
    AND p."id" = $2
`

type SellerDeleteProductParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	ID         int32  `json:"id" param:"id"`
}

func (q *Queries) SellerDeleteProduct(ctx context.Context, arg SellerDeleteProductParams) (int64, error) {
	result, err := q.db.Exec(ctx, sellerDeleteProduct, arg.SellerName, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const sellerDeleteProductTag = `-- name: SellerDeleteProductTag :execrows

DELETE FROM "product_tag" tp
WHERE EXISTS (
        SELECT 1
        FROM "product" p
            JOIN "shop" s ON s."id" = p."shop_id"
        WHERE
            s."seller_name" = $1
            AND p."id" = $3
    )
    AND "product_id" = $3
    AND "tag_id" = $2
`

type SellerDeleteProductTagParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	TagID      int32  `json:"tag_id"`
	ID         int32  `json:"id" param:"id"`
}

func (q *Queries) SellerDeleteProductTag(ctx context.Context, arg SellerDeleteProductTagParams) (int64, error) {
	result, err := q.db.Exec(ctx, sellerDeleteProductTag, arg.SellerName, arg.TagID, arg.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const sellerGetCoupon = `-- name: SellerGetCoupon :many

SELECT
    c."id",
    c."type",
    c."name",
    c."discount",
    c."expire_date"
FROM "coupon" c
    JOIN "shop" s ON c."shop_id" = s.id
WHERE s.seller_name = $1
ORDER BY "start_date" DESC
LIMIT $2
OFFSET $3
`

type SellerGetCouponParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

type SellerGetCouponRow struct {
	ID         int32              `json:"id" param:"id"`
	Type       CouponType         `json:"type"`
	Name       string             `json:"name"`
	Discount   pgtype.Numeric     `json:"discount" swaggertype:"string"`
	ExpireDate pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) SellerGetCoupon(ctx context.Context, arg SellerGetCouponParams) ([]SellerGetCouponRow, error) {
	rows, err := q.db.Query(ctx, sellerGetCoupon, arg.SellerName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SellerGetCouponRow{}
	for rows.Next() {
		var i SellerGetCouponRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Name,
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

const sellerGetCouponDetail = `-- name: SellerGetCouponDetail :one

SELECT
    c."id",
    c."type",
    c."name",
    c."discount",
    c."expire_date"
FROM "coupon" c
    JOIN "shop" s ON c."shop_id" = s.id
WHERE
    s."seller_name" = $1
    AND c."id" = $2
`

type SellerGetCouponDetailParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	ID         int32  `json:"id" param:"id"`
}

type SellerGetCouponDetailRow struct {
	ID         int32              `json:"id" param:"id"`
	Type       CouponType         `json:"type"`
	Name       string             `json:"name"`
	Discount   pgtype.Numeric     `json:"discount" swaggertype:"string"`
	ExpireDate pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) SellerGetCouponDetail(ctx context.Context, arg SellerGetCouponDetailParams) (SellerGetCouponDetailRow, error) {
	row := q.db.QueryRow(ctx, sellerGetCouponDetail, arg.SellerName, arg.ID)
	var i SellerGetCouponDetailRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Name,
		&i.Discount,
		&i.ExpireDate,
	)
	return i, err
}

const sellerGetCouponTag = `-- name: SellerGetCouponTag :many

SELECT ct.coupon_id, ct.tag_id
FROM "coupon_tag" ct
    JOIN "coupon" c ON c."id" = ct."coupon_id"
    JOIN "shop" s ON s."id" = c."shop_id"
WHERE
    s."seller_name" = $1
    AND "coupon_id" = $2
`

type SellerGetCouponTagParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	CouponID   int32  `json:"coupon_id" param:"id"`
}

func (q *Queries) SellerGetCouponTag(ctx context.Context, arg SellerGetCouponTagParams) ([]CouponTag, error) {
	rows, err := q.db.Query(ctx, sellerGetCouponTag, arg.SellerName, arg.CouponID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CouponTag{}
	for rows.Next() {
		var i CouponTag
		if err := rows.Scan(&i.CouponID, &i.TagID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const sellerGetInfo = `-- name: SellerGetInfo :one

SELECT
    "seller_name",
    "image_id",
    "description",
    "enabled"
FROM "shop"
WHERE "seller_name" = $1
`

type SellerGetInfoRow struct {
	SellerName  string      `json:"seller_name" param:"seller_name"`
	ImageID     pgtype.UUID `json:"image_id" swaggertype:"string"`
	Description string      `json:"description"`
	Enabled     bool        `json:"enabled"`
}

func (q *Queries) SellerGetInfo(ctx context.Context, sellerName string) (SellerGetInfoRow, error) {
	row := q.db.QueryRow(ctx, sellerGetInfo, sellerName)
	var i SellerGetInfoRow
	err := row.Scan(
		&i.SellerName,
		&i.ImageID,
		&i.Description,
		&i.Enabled,
	)
	return i, err
}

const sellerGetOrder = `-- name: SellerGetOrder :many

SELECT
    "id",
    "shipment",
    "total_price",
    "status",
    "created_at"
FROM "order_history"
WHERE "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
ORDER BY "created_at" DESC
LIMIT $2
OFFSET $3
`

type SellerGetOrderParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

type SellerGetOrderRow struct {
	ID         int32              `json:"id" param:"id"`
	Shipment   int32              `json:"shipment"`
	TotalPrice int32              `json:"total_price"`
	Status     OrderStatus        `json:"status"`
	CreatedAt  pgtype.Timestamptz `json:"created_at" swaggertype:"string"`
}

func (q *Queries) SellerGetOrder(ctx context.Context, arg SellerGetOrderParams) ([]SellerGetOrderRow, error) {
	rows, err := q.db.Query(ctx, sellerGetOrder, arg.SellerName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SellerGetOrderRow{}
	for rows.Next() {
		var i SellerGetOrderRow
		if err := rows.Scan(
			&i.ID,
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

const sellerGetOrderDetail = `-- name: SellerGetOrderDetail :many

SELECT
    product_archive.id, product_archive.version, product_archive.name, product_archive.description, product_archive.price, product_archive.image_id,
    order_detail.quantity
FROM "order_detail"
    LEFT JOIN product_archive ON order_detail.product_id = product_archive.id AND order_detail.product_version = product_archive.version
    LEFT JOIN order_history ON order_history.id = order_detail.order_id
    LEFT JOIN shop ON order_history.shop_id = shop.id
WHERE
    shop.seller_name = $1
    AND order_detail.order_id = $2
ORDER BY quantity * price DESC
`

type SellerGetOrderDetailParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	OrderID    int32  `json:"order_id" param:"id"`
}

type SellerGetOrderDetailRow struct {
	ID          pgtype.Int4    `json:"id"`
	Version     pgtype.Int4    `json:"version"`
	Name        pgtype.Text    `json:"name"`
	Description pgtype.Text    `json:"description"`
	Price       pgtype.Numeric `json:"price" swaggertype:"number"`
	ImageID     pgtype.UUID    `json:"image_id" swaggertype:"string"`
	Quantity    int32          `json:"quantity"`
}

func (q *Queries) SellerGetOrderDetail(ctx context.Context, arg SellerGetOrderDetailParams) ([]SellerGetOrderDetailRow, error) {
	rows, err := q.db.Query(ctx, sellerGetOrderDetail, arg.SellerName, arg.OrderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SellerGetOrderDetailRow{}
	for rows.Next() {
		var i SellerGetOrderDetailRow
		if err := rows.Scan(
			&i.ID,
			&i.Version,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.ImageID,
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

const sellerGetOrderHistory = `-- name: SellerGetOrderHistory :one

SELECT
    "order_history"."id",
    "order_history"."shipment",
    "order_history"."total_price",
    "order_history"."status",
    "order_history"."created_at"
FROM "order_history"
    JOIN shop ON order_history.shop_id = shop.id
WHERE
    shop.seller_name = $1
    AND order_history.id = $2
`

type SellerGetOrderHistoryParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	ID         int32  `json:"id" param:"id"`
}

type SellerGetOrderHistoryRow struct {
	ID         int32              `json:"id" param:"id"`
	Shipment   int32              `json:"shipment"`
	TotalPrice int32              `json:"total_price"`
	Status     OrderStatus        `json:"status"`
	CreatedAt  pgtype.Timestamptz `json:"created_at" swaggertype:"string"`
}

func (q *Queries) SellerGetOrderHistory(ctx context.Context, arg SellerGetOrderHistoryParams) (SellerGetOrderHistoryRow, error) {
	row := q.db.QueryRow(ctx, sellerGetOrderHistory, arg.SellerName, arg.ID)
	var i SellerGetOrderHistoryRow
	err := row.Scan(
		&i.ID,
		&i.Shipment,
		&i.TotalPrice,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const sellerGetProductDetail = `-- name: SellerGetProductDetail :one



SELECT
    p."id",
    p."name",
    p."image_id",
    p."price",
    p."sales",
    p."stock",
    p."enabled"
FROM "product" p
    JOIN "shop" s ON p."shop_id" = s.id
WHERE
    s.seller_name = $1
    AND p."id" = $2
`

type SellerGetProductDetailParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	ID         int32  `json:"id" param:"id"`
}

type SellerGetProductDetailRow struct {
	ID      int32          `json:"id" param:"id"`
	Name    string         `json:"name"`
	ImageID pgtype.UUID    `json:"image_id" swaggertype:"string"`
	Price   pgtype.Numeric `json:"price" swaggertype:"number"`
	Sales   int32          `json:"sales"`
	Stock   int32          `json:"stock"`
	Enabled bool           `json:"enabled"`
}

// SellerGetReport :many
// SellerGetReportDetail :many
func (q *Queries) SellerGetProductDetail(ctx context.Context, arg SellerGetProductDetailParams) (SellerGetProductDetailRow, error) {
	row := q.db.QueryRow(ctx, sellerGetProductDetail, arg.SellerName, arg.ID)
	var i SellerGetProductDetailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImageID,
		&i.Price,
		&i.Sales,
		&i.Stock,
		&i.Enabled,
	)
	return i, err
}

const sellerGetProductTag = `-- name: SellerGetProductTag :many

SELECT pt.tag_id, pt.product_id
FROM "product_tag" pt
    JOIN "product" p ON p."id" = pt."product_id"
    JOIN "shop" s ON s."id" = p."shop_id"
WHERE
    s."seller_name" = $1
    AND "product_id" = $2
`

type SellerGetProductTagParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	ProductID  int32  `json:"product_id" param:"id"`
}

func (q *Queries) SellerGetProductTag(ctx context.Context, arg SellerGetProductTagParams) ([]ProductTag, error) {
	rows, err := q.db.Query(ctx, sellerGetProductTag, arg.SellerName, arg.ProductID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductTag{}
	for rows.Next() {
		var i ProductTag
		if err := rows.Scan(&i.TagID, &i.ProductID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const sellerInsertCoupon = `-- name: SellerInsertCoupon :one

INSERT INTO
    "coupon" (
        "type",
        "shop_id",
        "name",
        "description",
        "discount",
        "start_date",
        "expire_date"
    )
VALUES (
        $2, (
            SELECT s."id"
            FROM "shop" s
            WHERE
                s."seller_name" = $1
                AND s."enabled" = true
        ),
        $3,
        $4,
        $5,
        $6,
        $7
    ) RETURNING "id",
    "type",
    "name",
    "discount",
    "expire_date"
`

type SellerInsertCouponParams struct {
	SellerName  string             `json:"seller_name" param:"seller_name"`
	Type        CouponType         `json:"type"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"string"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

type SellerInsertCouponRow struct {
	ID         int32              `json:"id" param:"id"`
	Type       CouponType         `json:"type"`
	Name       string             `json:"name"`
	Discount   pgtype.Numeric     `json:"discount" swaggertype:"string"`
	ExpireDate pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) SellerInsertCoupon(ctx context.Context, arg SellerInsertCouponParams) (SellerInsertCouponRow, error) {
	row := q.db.QueryRow(ctx, sellerInsertCoupon,
		arg.SellerName,
		arg.Type,
		arg.Name,
		arg.Description,
		arg.Discount,
		arg.StartDate,
		arg.ExpireDate,
	)
	var i SellerInsertCouponRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Name,
		&i.Discount,
		&i.ExpireDate,
	)
	return i, err
}

const sellerInsertCouponTag = `-- name: SellerInsertCouponTag :one

INSERT INTO
    "coupon_tag" ("tag_id", "coupon_id")
SELECT $2, $3
WHERE EXISTS (
        SELECT 1
        FROM "tag" t
            JOIN "shop" s ON s."id" = t."shop_id"
        WHERE
            s."seller_name" = $1
            AND t."id" = $2
    )
    AND EXISTS (
        SELECT 1
        FROM "coupon" c
            JOIN "shop" s ON s."id" = c."shop_id"
        WHERE
            s."seller_name" = $1
            AND c."id" = $3
    ) RETURNING coupon_id, tag_id
`

type SellerInsertCouponTagParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	TagID      int32  `json:"tag_id"`
	CouponID   int32  `json:"coupon_id" param:"id"`
}

func (q *Queries) SellerInsertCouponTag(ctx context.Context, arg SellerInsertCouponTagParams) (CouponTag, error) {
	row := q.db.QueryRow(ctx, sellerInsertCouponTag, arg.SellerName, arg.TagID, arg.CouponID)
	var i CouponTag
	err := row.Scan(&i.CouponID, &i.TagID)
	return i, err
}

const sellerInsertProduct = `-- name: SellerInsertProduct :one

INSERT INTO
    "product"(
        "version",
        "shop_id",
        "name",
        "description",
        "price",
        "image_id",
        "expire_date",
        "edit_date",
        "stock",
        "enabled"
    )
VALUES (
        1, (
            SELECT s."id"
            FROM "shop" s
            WHERE
                s."seller_name" = $1
                AND s."enabled" = true
        ),
        $2,
        $3,
        $4,
        $5,
        $6,
        NOW(),
        $7,
        $8
    ) RETURNING "id",
    "name",
    "description",
    "price",
    "image_id",
    "expire_date",
    "edit_date",
    "stock",
    "sales"
`

type SellerInsertProductParams struct {
	SellerName  string             `json:"seller_name" param:"seller_name"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       pgtype.Numeric     `json:"price" swaggertype:"number"`
	ImageID     pgtype.UUID        `json:"image_id" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
	Stock       int32              `json:"stock"`
	Enabled     bool               `json:"enabled"`
}

type SellerInsertProductRow struct {
	ID          int32              `json:"id" param:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       pgtype.Numeric     `json:"price" swaggertype:"number"`
	ImageID     pgtype.UUID        `json:"image_id" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
	EditDate    pgtype.Timestamptz `json:"edit_date" swaggertype:"string"`
	Stock       int32              `json:"stock"`
	Sales       int32              `json:"sales"`
}

func (q *Queries) SellerInsertProduct(ctx context.Context, arg SellerInsertProductParams) (SellerInsertProductRow, error) {
	row := q.db.QueryRow(ctx, sellerInsertProduct,
		arg.SellerName,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.ImageID,
		arg.ExpireDate,
		arg.Stock,
		arg.Enabled,
	)
	var i SellerInsertProductRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.ImageID,
		&i.ExpireDate,
		&i.EditDate,
		&i.Stock,
		&i.Sales,
	)
	return i, err
}

const sellerInsertProductTag = `-- name: SellerInsertProductTag :one

INSERT INTO
    "product_tag" ("tag_id", "product_id")
SELECT $2, $3
WHERE EXISTS (
        SELECT 1
        FROM "tag" t
            JOIN "shop" s ON s."id" = t."shop_id"
        WHERE
            s."seller_name" = $1
            AND t."id" = $2
    )
    AND EXISTS (
        SELECT 1
        FROM "product" p
            JOIN "shop" s ON s."id" = p."shop_id"
        WHERE
            s."seller_name" = $1
            AND p."id" = $3
    ) RETURNING tag_id, product_id
`

type SellerInsertProductTagParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	TagID      int32  `json:"tag_id"`
	ProductID  int32  `json:"product_id" param:"id"`
}

func (q *Queries) SellerInsertProductTag(ctx context.Context, arg SellerInsertProductTagParams) (ProductTag, error) {
	row := q.db.QueryRow(ctx, sellerInsertProductTag, arg.SellerName, arg.TagID, arg.ProductID)
	var i ProductTag
	err := row.Scan(&i.TagID, &i.ProductID)
	return i, err
}

const sellerInsertTag = `-- name: SellerInsertTag :one

INSERT INTO
    "tag" ("shop_id", "name")
VALUES ( (
            SELECT s."id"
            FROM "shop" s
            WHERE
                s."seller_name" = $1
                AND s."enabled" = true
        ),
        $2
    ) RETURNING "id",
    "name"
`

type SellerInsertTagParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	Name       string `json:"name"`
}

type SellerInsertTagRow struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) SellerInsertTag(ctx context.Context, arg SellerInsertTagParams) (SellerInsertTagRow, error) {
	row := q.db.QueryRow(ctx, sellerInsertTag, arg.SellerName, arg.Name)
	var i SellerInsertTagRow
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const sellerProductList = `-- name: SellerProductList :many

SELECT
    p."id",
    p."name",
    p."image_id",
    p."price",
    p."sales",
    p."stock",
    p."enabled"
FROM "product" p
    JOIN "shop" s ON p."shop_id" = s.id
WHERE s.seller_name = $1
ORDER BY "sales" DESC
LIMIT $2
OFFSET $3
`

type SellerProductListParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

type SellerProductListRow struct {
	ID      int32          `json:"id" param:"id"`
	Name    string         `json:"name"`
	ImageID pgtype.UUID    `json:"image_id" swaggertype:"string"`
	Price   pgtype.Numeric `json:"price" swaggertype:"number"`
	Sales   int32          `json:"sales"`
	Stock   int32          `json:"stock"`
	Enabled bool           `json:"enabled"`
}

func (q *Queries) SellerProductList(ctx context.Context, arg SellerProductListParams) ([]SellerProductListRow, error) {
	rows, err := q.db.Query(ctx, sellerProductList, arg.SellerName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SellerProductListRow{}
	for rows.Next() {
		var i SellerProductListRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ImageID,
			&i.Price,
			&i.Sales,
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

const sellerSearchTag = `-- name: SellerSearchTag :many

SELECT t."id", t."name"
FROM "tag" t
    LEFT JOIN "shop" s ON "shop_id" = s.id
WHERE
    s."seller_name" = $1
    AND t."name" ~* $2
ORDER BY LENGTH(t."name") ASC
LIMIT $3
`

type SellerSearchTagParams struct {
	SellerName string `json:"seller_name" param:"seller_name"`
	Name       string `json:"name"`
	Limit      int32  `json:"limit"`
}

type SellerSearchTagRow struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) SellerSearchTag(ctx context.Context, arg SellerSearchTagParams) ([]SellerSearchTagRow, error) {
	rows, err := q.db.Query(ctx, sellerSearchTag, arg.SellerName, arg.Name, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SellerSearchTagRow{}
	for rows.Next() {
		var i SellerSearchTagRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const sellerUpdateCouponInfo = `-- name: SellerUpdateCouponInfo :one

UPDATE "coupon" c
SET
    "type" = COALESCE($3, "type"),
    "name" = COALESCE($4, "name"),
    "description" = COALESCE($5, "description"),
    "discount" = COALESCE($6, "discount"),
    "start_date" = COALESCE($7, "start_date"),
    "expire_date" = COALESCE($8, "expire_date")
WHERE c."id" = $2 AND "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    ) RETURNING c."id",
    c."type",
    c."name",
    c."discount",
    c."expire_date"
`

type SellerUpdateCouponInfoParams struct {
	SellerName  string             `json:"seller_name" param:"seller_name"`
	ID          int32              `json:"id" param:"id"`
	Type        CouponType         `json:"type"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"string"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

type SellerUpdateCouponInfoRow struct {
	ID         int32              `json:"id" param:"id"`
	Type       CouponType         `json:"type"`
	Name       string             `json:"name"`
	Discount   pgtype.Numeric     `json:"discount" swaggertype:"string"`
	ExpireDate pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

func (q *Queries) SellerUpdateCouponInfo(ctx context.Context, arg SellerUpdateCouponInfoParams) (SellerUpdateCouponInfoRow, error) {
	row := q.db.QueryRow(ctx, sellerUpdateCouponInfo,
		arg.SellerName,
		arg.ID,
		arg.Type,
		arg.Name,
		arg.Description,
		arg.Discount,
		arg.StartDate,
		arg.ExpireDate,
	)
	var i SellerUpdateCouponInfoRow
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Name,
		&i.Discount,
		&i.ExpireDate,
	)
	return i, err
}

const sellerUpdateInfo = `-- name: SellerUpdateInfo :one

UPDATE "shop"
SET
    "image_id" = COALESCE($2, "image_id"),
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "enabled" = COALESCE($5, "enabled")
WHERE
    "seller_name" = $1 RETURNING "seller_name",
    "image_id",
    "name",
    "enabled"
`

type SellerUpdateInfoParams struct {
	SellerName  string      `json:"seller_name" param:"seller_name"`
	ImageID     pgtype.UUID `json:"image_id" swaggertype:"string"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Enabled     bool        `json:"enabled"`
}

type SellerUpdateInfoRow struct {
	SellerName string      `json:"seller_name" param:"seller_name"`
	ImageID    pgtype.UUID `json:"image_id" swaggertype:"string"`
	Name       string      `json:"name"`
	Enabled    bool        `json:"enabled"`
}

func (q *Queries) SellerUpdateInfo(ctx context.Context, arg SellerUpdateInfoParams) (SellerUpdateInfoRow, error) {
	row := q.db.QueryRow(ctx, sellerUpdateInfo,
		arg.SellerName,
		arg.ImageID,
		arg.Name,
		arg.Description,
		arg.Enabled,
	)
	var i SellerUpdateInfoRow
	err := row.Scan(
		&i.SellerName,
		&i.ImageID,
		&i.Name,
		&i.Enabled,
	)
	return i, err
}

const sellerUpdateOrderStatus = `-- name: SellerUpdateOrderStatus :one

UPDATE "order_history" oh
SET
    "status" = $3
WHERE "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
    AND oh."id" = $2
    AND oh."status" = $4 RETURNING oh."id",
    oh."shipment",
    oh."total_price",
    oh."status",
    oh."created_at"
`

type SellerUpdateOrderStatusParams struct {
	SellerName    string      `json:"seller_name" param:"seller_name"`
	ID            int32       `json:"id" param:"id"`
	SetStatus     OrderStatus `json:"set_status"`
	CurrentStatus OrderStatus `json:"current_status"`
}

type SellerUpdateOrderStatusRow struct {
	ID         int32              `json:"id" param:"id"`
	Shipment   int32              `json:"shipment"`
	TotalPrice int32              `json:"total_price"`
	Status     OrderStatus        `json:"status"`
	CreatedAt  pgtype.Timestamptz `json:"created_at" swaggertype:"string"`
}

func (q *Queries) SellerUpdateOrderStatus(ctx context.Context, arg SellerUpdateOrderStatusParams) (SellerUpdateOrderStatusRow, error) {
	row := q.db.QueryRow(ctx, sellerUpdateOrderStatus,
		arg.SellerName,
		arg.ID,
		arg.SetStatus,
		arg.CurrentStatus,
	)
	var i SellerUpdateOrderStatusRow
	err := row.Scan(
		&i.ID,
		&i.Shipment,
		&i.TotalPrice,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const sellerUpdateProductInfo = `-- name: SellerUpdateProductInfo :one

UPDATE "product" p
SET
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "price" = COALESCE($5, "price"),
    "image_id" = COALESCE($6, "image_id"),
    "expire_date" = COALESCE($7, "expire_date"),
    "enabled" = COALESCE($8, "enabled"),
    "stock" = COALESCE($9, "stock"),
    "edit_date" = NOW(),
    "version" = "version" + 1
WHERE "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
    AND p."id" = $2 RETURNING "id",
    "name",
    "description",
    "price",
    "image_id",
    "expire_date",
    "edit_date",
    "stock",
    "sales"
`

type SellerUpdateProductInfoParams struct {
	SellerName  string             `json:"seller_name" param:"seller_name"`
	ID          int32              `json:"id" param:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       pgtype.Numeric     `json:"price" swaggertype:"number"`
	ImageID     pgtype.UUID        `json:"image_id" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
	Enabled     bool               `json:"enabled"`
	Stock       int32              `json:"stock"`
}

type SellerUpdateProductInfoRow struct {
	ID          int32              `json:"id" param:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       pgtype.Numeric     `json:"price" swaggertype:"number"`
	ImageID     pgtype.UUID        `json:"image_id" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
	EditDate    pgtype.Timestamptz `json:"edit_date" swaggertype:"string"`
	Stock       int32              `json:"stock"`
	Sales       int32              `json:"sales"`
}

func (q *Queries) SellerUpdateProductInfo(ctx context.Context, arg SellerUpdateProductInfoParams) (SellerUpdateProductInfoRow, error) {
	row := q.db.QueryRow(ctx, sellerUpdateProductInfo,
		arg.SellerName,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.ImageID,
		arg.ExpireDate,
		arg.Enabled,
		arg.Stock,
	)
	var i SellerUpdateProductInfoRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.ImageID,
		&i.ExpireDate,
		&i.EditDate,
		&i.Stock,
		&i.Sales,
	)
	return i, err
}
