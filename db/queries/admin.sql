-- name: GetUsers :many
SELECT
    "username",
    "name",
    "email",
    "address",
    "image_id" AS "icon_url",
    "role",
    "credit_card",
    "enabled"
FROM
    "user"
WHERE
    "enabled" = TRUE
ORDER BY
    "id" ASC
LIMIT $1 OFFSET $2;

-- name: EnabledShop :execrows
UPDATE
    "shop"
SET
    "enabled" = TRUE
WHERE
    "seller_name" = $1;

-- name: DisableUser :execrows
WITH disabled_user AS (
    UPDATE
        "user"
    SET
        "enabled" = FALSE
    WHERE
        "username" = $1
    RETURNING
        "username"
),
disabled_shop AS (
    UPDATE
        "shop"
    SET
        "enabled" = FALSE
    WHERE
        "seller_name" =(
            SELECT
                "username"
            FROM
                disabled_user)
        RETURNING
            "id")
UPDATE
    "product"
SET
    "enabled" = FALSE
WHERE
    "shop_id" =(
        SELECT
            "id"
        FROM
            disabled_shop);

-- name: DisableShop :execrows
WITH disable_shop AS (
    UPDATE
        "shop"
    SET
        "enabled" = FALSE
    WHERE
        "seller_name" = $1
    RETURNING
        "id")
UPDATE
    "product"
SET
    "enabled" = FALSE
WHERE
    "shop_id" =(
        SELECT
            "id"
        FROM
            disable_shop);

-- name: DisableProductsFromShop :execrows
UPDATE
    "product"
SET
    "enabled" = FALSE
WHERE
    "shop_id" = $1;

-- name: CouponExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            "coupon"
        WHERE
            "id" = $1);

-- name: GetGlobalCoupons :many
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
    "scope" = 'global'
ORDER BY
    "id" ASC
LIMIT $1 OFFSET $2;

-- name: GetGlobalCouponDetail :one
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
    "scope" = 'global'
    AND "id" = $1;

-- name: AddCoupon :one
INSERT INTO "coupon"("type", "scope", "name", "description", "discount", "start_date", "expire_date")
    VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    "id", "type", "scope", "name", "description", "discount", "start_date", "expire_date";

-- name: EditCoupon :one
UPDATE
    "coupon"
SET
    "type" = COALESCE($2, "type"),
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "discount" = COALESCE($5, "discount"),
    "start_date" = COALESCE($6, "start_date"),
    "expire_date" = COALESCE($7, "expire_date")
WHERE
    "id" = $1
    AND "scope" = 'global'
RETURNING
    "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date";

-- name: DeleteCoupon :execrows
DELETE FROM "coupon"
WHERE "id" = $1
    AND "scope" = 'global';

-- name: GetUserIDByUsername :one
SELECT
    "id"
FROM
    "user"
WHERE
    "username" = $1;

-- name: GetShopIDBySellerName :one
SELECT
    "id"
FROM
    "shop"
WHERE
    "seller_name" = $1;

-- name: GetTopSeller :many
SELECT
    S."seller_name",
    S."name",
    S."image_id" AS "image_url",
    SUM(O."total_price") AS "total_sales"
FROM
    "shop" AS S,
    "order_history" AS O
WHERE
    S."id" = O."shop_id"
    AND O."status" = 'paid'
    AND O."created_at" < CAST((CAST(@date::TEXT AS TIMESTAMPTZ) +(INTERVAL '1 month')) AS TIMESTAMPTZ)
    AND O."created_at" >= CAST(@date::TEXT AS TIMESTAMPTZ)
GROUP BY
    S."seller_name",
    S."name",
    S."image_id"
ORDER BY
    "total_sales" DESC
LIMIT 3;

-- name: GetTotalSales :one
SELECT
    COALESCE(SUM("total_price"), 0)::INTEGER AS "total_sales"
FROM
    "order_history"
WHERE
    "status" = 'paid'
    AND "created_at" < CAST((CAST(@date::TEXT AS TIMESTAMPTZ) +(INTERVAL '1 month')) AS TIMESTAMPTZ)
    AND "created_at" >= CAST(@date::TEXT AS TIMESTAMPTZ);
