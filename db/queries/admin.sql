-- name: GetUsers :many

SELECT
    "username",
    "name",
    "email",
    "address",
    "role",
    "credit_card",
    "enabled"
FROM "user";

-- name: AddUser :exec

INSERT INTO
    "user" (
        "username",
        "password",
        "name",
        "email",
        "address",
        "image_id",
        "role",
        "credit_card"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: DisableUser :exec

UPDATE "user" SET "enabled" = FALSE WHERE "id" = $1;

-- name: DisableShop :exec

UPDATE "shop" AS s
SET s."enabled" = FALSE
WHERE s."seller_name" = $1;

-- name: DisableProductsFromShop :exec

UPDATE "product" AS p SET p."enabled" = FALSE WHERE p."shop_id" = $1;

-- name: GetAnyCoupons :many

SELECT * FROM "coupon";

-- name: GetCouponDetail :one

SELECT * FROM "coupon" WHERE "id" = $1;

-- name: AddCoupon :exec

INSERT INTO
    "coupon" (
        "id",
        "type",
        "shop_id",
        "name",
        "description",
        "discount",
        "start_date",
        "expire_date"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: EditCoupon :exec

UPDATE "coupon"
SET
    "type" = COALESCE($2, "type"),
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "discount" = COALESCE($5, "discount"),
    "start_date" = COALESCE($6, "start_date"),
    "expire_date" = COALESCE($7, "expire_date")
WHERE "id" = $1;

-- name: DeleteCoupon :exec

DELETE FROM "coupon" WHERE "id" = $1;

-- name: GetReport :many

SELECT *
FROM "order_history"
WHERE
    "created_at" >= $1
    AND "created_at" <= $2;

-- name: GetUserIDByUsername :one

SELECT "id" FROM "user" WHERE "username" = $1;
