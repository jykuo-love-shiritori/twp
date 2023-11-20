-- name: GetSellerInfo :one

SELECT * FROM "shop" WHERE "id" = $1;

-- name: UpdateSellerInfo :exec

UPDATE "shop"
SET
    "seller_name" = CASE
        WHEN $2 IS NOT NULL THEN $2
        ELSE "seller_name"
    END,
    "image_id" = CASE
        WHEN $3 IS NOT NULL THEN $3
        ELSE "image_id"
    END,
    "name" = CASE
        WHEN $4 IS NOT NULL THEN $4
        ELSE "name"
    END,
    "description" = CASE
        WHEN $5 IS NOT NULL THEN $5
        ELSE "description"
    END,
    "enabled" = CASE
        WHEN $6 IS NOT NULL THEN $6
        ELSE "enabled"
    END
WHERE "id" = $1;

-- name: SearchTag :many

SELECT "id", "name"
FROM "tag"
WHERE
    "shop_id" = $1
    AND "name" LIKE $2 || '%'
ORDER BY LEN("name")
LIMIT 10;

-- name: InsertTag :one

INSERT INTO "tag" ( "shop_id", "name") VALUES ($1,$2) RETURNING "id";

-- name: SellerGetCoupon :many

SELECT
    "id",
    "type",
    "shop_id",
    "name",
    "discount",
    "expire_date"
FROM "coupon"
WHERE "shop_id" = $1
ORDER BY "start_date" DESC
LIMIT $2
OFFSET $3;

-- name: SellerGetCouponDetail :one

SELECT * FROM "coupon" WHERE "id" = $1 and "shop_id" = $2;

-- name: SellerInsertCoupon :one

INSERT INTO
    "coupon" (
        "type",
        "shop_id",
        "description",
        "discount",
        "start_date",
        "expire_date"
    )
VALUES ($1, $2, $3, $4, NOW(), $5)
RETURNING "id";

-- name: UpdateCouponInfo :exec

UPDATE "coupon"
SET
    "type" = CASE
        WHEN $3 IS NOT NULL THEN $3
        ELSE "type"
    END,
    "description" = CASE
        WHEN $4 IS NOT NULL THEN $4
        ELSE "description"
    END,
    "discount" = CASE
        WHEN $5 IS NOT NULL THEN $5
        ELSE "discount"
    END,
    "start_date" = CASE
        WHEN $6 IS NOT NULL THEN $6
        ELSE "start_date"
    END,
    "expire_date" = CASE
        WHEN $7 IS NOT NULL THEN $7
        ELSE "expire_date"
    END
WHERE "id" = $1 AND "shop_id" = $2;

-- name: DeleteCoupon :exec

DELETE FROM "coupon" WHERE "id" = $1 AND "shop_id"=$2;

-- name: SellerGetOrder :many

SELECT
    "id",
    "shop_id",
    "shipment",
    "total_price",
    "status",
    "created_at"
FROM "order_history"
WHERE "shop_id" = $1
ORDER BY "created_at" DESC
LIMIT $2
OFFSET $3;

-- name: OrderDetail :one

SELECT *
FROM "order_detail"
    LEFT JOIN "product_archive" ON "order_detail"."product_id" = "product"."id" AND "order_detail"."version" = "product"."id"
WHERE "order_id" = $1;

-- name: SellerGetReport :many

-- name: SellerGetReportDetail :many

-- name: SellerInsertProduct :one

INSERT INTO
    "product" (
        "version",
        "shop_id",
        "name",
        "description",
        "price",
        "image_id",
        "exp_date"
    )
VALUES (0, $1, $2, $3, $4, $5, $6)
RETURNING "id";

-- name: UpdateProductInfo :exec

UPDATE "product"
SET
    "name" = CASE
        WHEN $3 IS NOT NULL THEN $3
        ELSE "description"
    END,
    "description" = CASE
        WHEN $4 IS NOT NULL THEN $4
        ELSE "discount"
    END,
    "price" = CASE
        WHEN $5 IS NOT NULL THEN $5
        ELSE "start_date"
    END,
    "image_id" = CASE
        WHEN $6 IS NOT NULL THEN $6
        ELSE "image_id"
    END,
    "exp_date" = CASE
        WHEN $7 IS NOT NULL THEN $7
        ELSE "exp_date"
    END,
    "edit_date" = NOW(),
    "version" = "version" + 1
WHERE "id" = $1 AND "shop_id" = $2;

-- name: DeleteProduct :exec

UPDATE "product"
SET "enabled" = false
WHERE "id" = $1 AND "shop_id" = $2;