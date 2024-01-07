-- name: GetShopInfo :one
SELECT
    "seller_name",
    "image_id" AS "image_url",
    "name",
    "description"
FROM
    "shop"
WHERE
    "seller_name" = $1
    AND "enabled" = TRUE;

-- name: ShopExists :one
SELECT
    "id"
FROM
    "shop" AS s
WHERE
    s."seller_name" = $1
    AND s."enabled" = TRUE;

-- name: GetShopCoupons :many
SELECT
    "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date"
FROM
    "coupon"
WHERE
    "shop_id" = $1
    OR "scope" = 'global'
ORDER BY
    "id" ASC
LIMIT $2 OFFSET $3;

-- name: GetTagInfo :one
SELECT
    "id",
    "name"
FROM
    "tag"
WHERE
    "id" = $1;

-- name: GetProductInfo :one
SELECT
    P."id",
    P."name",
    P."description",
    P."price",
    P."image_id" AS "image_url",
    P."expire_date",
    P."stock",
    P."sales",
    S."seller_name" AS "seller_name"
FROM
    "product" AS P
    JOIN "shop" S ON S."id" = P."shop_id"
WHERE
    P."id" = $1
    AND P."enabled" = TRUE;

-- name: GetShopProducts :many
SELECT
    P."id",
    P."name",
    P."description",
    P."price",
    P."image_id" AS "image_url",
    P."expire_date",
    P."stock",
    P."sales"
FROM
    "product" P,
    "shop" S
WHERE
    S."seller_name" = $1
    AND P."shop_id" = S."id"
    AND P."enabled" = TRUE
ORDER BY
    P."sales" DESC
LIMIT $2 OFFSET $3;

-- name: GetSellerNameByShopID :one
SELECT
    "seller_name"
FROM
    "shop"
WHERE
    "id" = $1;

-- name: GetProductsFromPopularShop :many
WITH popular_shop AS (
    SELECT
        S."id" AS "id"
    FROM
        "shop" S,
        "order_history" O
    WHERE
        S."id" = O."shop_id"
        AND S."enabled" = TRUE
        AND O."created_at" >=(NOW() -(INTERVAL '1 month'))
        AND (
            SELECT
                COUNT("product"."id")
            FROM
                "product"
            WHERE
                "product"."shop_id" = S."id"
                AND "product"."enabled" = TRUE) >= 1
        GROUP BY
            S."id"
        ORDER BY
            COUNT(O."id") DESC
        LIMIT 1
)
SELECT
    "id",
    "name",
    "description",
    "price",
    "image_id" AS "image_url",
    "sales"
FROM
    "product"
WHERE
    "shop_id" =(
        SELECT
            "id"
        FROM
            popular_shop)
    AND "enabled" = TRUE
ORDER BY
    "sales" DESC
LIMIT 4;

-- name: GetProductsFromNearByShop :many
WITH nearby_shop AS (
    SELECT
        S."id" AS "id"
    FROM
        "shop" S
    WHERE
        S."enabled" = TRUE
        AND (
            SELECT
                COUNT("product"."id")
            FROM
                "product"
            WHERE
                "product"."shop_id" = S."id"
                AND "product"."enabled" = TRUE) >= 1
        ORDER BY
            RANDOM() -- implement distance in future
        LIMIT 1
)
SELECT
    "id",
    "name",
    "description",
    "price",
    "image_id" AS "image_url",
    "sales"
FROM
    "product"
WHERE
    "shop_id" =(
        SELECT
            "id"
        FROM
            nearby_shop)
    AND "enabled" = TRUE
ORDER BY
    "sales" DESC
LIMIT 4;

-- name: GetRandomProducts :many
SELECT
    "id",
    "name",
    "description",
    "price",
    "image_id" AS "image_url",
    "sales"
FROM
    "product"
WHERE
    "enabled" = TRUE
ORDER BY
    "image_id" -- random but stable
LIMIT $1 OFFSET $2;

-- name: SearchProducts :many
SELECT
    "id",
    "name",
    "price",
    "image_url",
    "sales"
FROM
    search_products(@query, NULL, sqlc.narg('min_price'), sqlc.narg('max_price'), sqlc.narg('min_stock'), sqlc.narg('max_stock'), sqlc.narg('has_coupon'), sqlc.narg('sort_by'), sqlc.narg('order'), sqlc.arg('offset'), sqlc.arg('limit'));

-- name: SearchProductsByShop :many
WITH shop_id AS (
    SELECT
        "id"
    FROM
        "shop"
    WHERE
        "seller_name" = $1
        AND "enabled" = TRUE
)
SELECT
    "id",
    "name",
    "price",
    "image_url",
    "sales"
FROM
    search_products(@query,(
            SELECT
                "id"
            FROM shop_id), sqlc.narg('min_price'), sqlc.narg('max_price'), sqlc.narg('min_stock'), sqlc.narg('max_stock'), sqlc.narg('has_coupon'), sqlc.narg('sort_by'), sqlc.narg('order'), sqlc.arg('offset'), sqlc.arg('limit'));

-- name: SearchShops :many
SELECT
    "name",
    "seller_name",
    "image_url"
FROM
    search_shop(@query, sqlc.arg('offset'), sqlc.arg('limit'));
