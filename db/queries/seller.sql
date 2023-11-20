-- name: GetSellerInfo :one

SELECT s.*
FROM "user" u
    JOIN "shop" s ON u.username = s.seller_name
WHERE u.id = $1;

-- name: UpdateSellerInfo :exec

UPDATE "shop"
SET
    "image_id" = COALESCE($2, "image_id"),
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "enabled" = COALESCE($5, "enabled")
WHERE "seller_name" IN (
        SELECT "username"
        FROM "user" u
        WHERE u.id = $1
    );

-- name: SearchTag :many

SELECT t."id", t."name"
FROM "tag" t
    JOIN "shop" s ON "shop_id" = s.id
    JOIN "user" u ON s.seller_name = u.username
WHERE u.id = $1 AND t."name" ~* $2
ORDER BY LENGTH(t."name")
LIMIT 10;

-- name: InsertTag :one

WITH user_shop_info AS (
        SELECT
            u."id" AS "user_id",
            s."id" AS "shop_id"
        FROM "user" u
            JOIN "shop" s ON u."username" = s."seller_name"
        WHERE u."id" = $1
    )
INSERT INTO
    "tag" ("shop_id", "name")
SELECT u.shop_id, $2
FROM user_shop_info
RETURNING ("id", "name");

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

SELECT * FROM "coupon" WHERE "id"= $1 and"shop_id"= $2;

-- name: SellerInsertCoupon :one

INSERT INTO
    "coupon"(
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
    "type" = COALESCE($3, "type"),
    "description" = COALESCE($4, "description"),
    "discount" = COALESCE($4, "discount"),
    "start_date" = COALESCE($4, "start_date"),
    "expire_date" = COALESCE($4, "expire_date")
WHERE "id" = $1 AND "shop_id" = $2;

-- name: DeleteCoupon :exec

DELETE FROM"coupon"WHERE"id"= $1 AND"shop_id"=$2;

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
    "product"(
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
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "price" = COALESCE($5, "price"),
    "image_id" = COALESCE($6, "image_id"),
    "exp_date" = COALESCE($7, "exp_date"),
    "description" = COALESCE($8, "description"),
    "edit_date" = NOW(),
    "version" = "version" + 1
WHERE "id" = $1 AND "shop_id" = $2;

-- name: DeleteProduct :exec

UPDATE "product"
SET "enabled" = false
WHERE "id" = $1 AND "shop_id" = $2;