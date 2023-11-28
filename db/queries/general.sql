-- name: GetShopInfo :one

SELECT
    "seller_name",
    "image_id",
    "name",
    "description"
FROM "shop"
WHERE "seller_name" = $1;

-- name: ShopExists :one

SELECT "id"
FROM "shop" AS s
WHERE EXISTS(
        SELECT 1
        FROM "shop"
        WHERE
            "enabled" = TRUE
    )
    AND s."seller_name" = $1;

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
FROM "coupon"
WHERE "shop_id" = $1;

-- name: GetTagInfo :one

SELECT "id", "name" FROM "tag" WHERE "id" = $1;

-- name: GetProductInfo :one

SELECT
    "id",
    "version",
    "name",
    "description",
    "price",
    "image_id",
    "exp_date",
    "stock",
    "sales",
    "enabled"
FROM "product"
WHERE "id" = $1;
