-- name: GetShopInfo :one

SELECT
    "seller_name",
    "image_id",
    "name",
    "description"
FROM "shop"
WHERE
    "seller_name" = $1
    AND "enabled" = TRUE;

-- name: ShopExists :one

SELECT "id"
FROM "shop" AS s
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
FROM "coupon"
WHERE
    "shop_id" = $1
    OR "scope" = 'global'
ORDER BY "id" ASC
LIMIT $2
OFFSET $3;

-- name: GetTagInfo :one

SELECT "id", "name" FROM "tag" WHERE "id" = $1;

-- name: GetProductInfo :one

SELECT
    "id",
    "name",
    "description",
    "price",
    "image_id",
    "exp_date",
    "stock",
    "sales"
FROM "product"
WHERE
    "id" = $1
    AND "enabled" = TRUE;

-- name: GetSellerNameByShopID :one

SELECT "seller_name" FROM "shop" WHERE "id" = $1;
