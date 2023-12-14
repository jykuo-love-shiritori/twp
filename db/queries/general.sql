-- name: GetShopInfo :one
SELECT "seller_name",
    "image_id",
    "name",
    "description"
FROM "shop"
WHERE "seller_name" = $1
    AND "enabled" = TRUE;

-- name: ShopExists :one
SELECT "id"
FROM "shop" AS s
WHERE s."seller_name" = $1
    AND s."enabled" = TRUE;

-- name: GetShopCoupons :many
SELECT "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date"
FROM "coupon"
WHERE "shop_id" = $1
    OR "scope" = 'global'
ORDER BY "id" ASC
LIMIT $2 OFFSET $3;

-- name: GetTagInfo :one
SELECT "id",
    "name"
FROM "tag"
WHERE "id" = $1;

-- name: GetProductInfo :one
SELECT "id",
    "name",
    "description",
    "price",
    "image_id",
    "expire_date",
    "stock",
    "sales"
FROM "product"
WHERE "id" = $1
    AND "enabled" = TRUE;

-- name: GetShopProducts :many
SELECT P."id",
    P."name",
    P."description",
    P."price",
    P."image_id",
    P."expire_date",
    P."stock",
    P."sales"
FROM "product" P,
    "shop" S
WHERE S."seller_name" = $1
    AND P."shop_id" = S."id"
    AND P."enabled" = TRUE
ORDER BY P."sales" DESC
LIMIT $2 OFFSET $3;

-- name: GetSellerNameByShopID :one
SELECT "seller_name"
FROM "shop"
WHERE "id" = $1;

-- name: GetProductsFromPopularShop :many
WITH popular_shop AS (
    SELECT S."id" AS "id"
    FROM "shop" S,
        "order_history" O
    WHERE S."id" = O."shop_id"
        AND S."enabled" = TRUE
        AND O."created_at" >= (NOW() - (INTERVAL '1 month'))
        AND (
            SELECt COUNT("product"."id")
            FROM "product"
            WHERE "product"."shop_id" = S."id"
                AND "product"."enabled" = TRUE
        ) >= 1
    GROUP BY S."id"
    ORDER BY COUNT(O."id") DESC
    LIMIT 1
)
SELECT "id",
    "name",
    "description",
    "price",
    "image_id",
    "sales"
FROM "product"
WHERE "shop_id" = (
        SELECT "id"
        FROM popular_shop
    )
    AND "enabled" = TRUE
ORDER BY "sales" DESC
LIMIT 4;

-- name: GetProductsFromNearByShop :many
WITH nearby_shop AS (
    SELECT S."id" AS "id"
    FROM "shop" S
    WHERE S."enabled" = TRUE
        AND (
            SELECt COUNT("product"."id")
            FROM "product"
            WHERE "product"."shop_id" = S."id"
                AND "product"."enabled" = TRUE
        ) >= 1
    ORDER BY RANDOM() -- implement distance in future
    LIMIT 1
)
SELECT "id",
    "name",
    "description",
    "price",
    "image_id",
    "sales"
FROM "product"
WHERE "shop_id" = (
        SELECT "id"
        FROM nearby_shop
    )
    AND "enabled" = TRUE
ORDER BY "sales" DESC
LIMIT 4;

-- name: GetRandomProducts :many
SELECT "id",
    "name",
    "description",
    "price",
    "image_id",
    "sales"
FROM "product"
WHERE "enabled" = TRUE
ORDER BY "image_id" -- random but stable
LIMIT $1 OFFSET $2;
