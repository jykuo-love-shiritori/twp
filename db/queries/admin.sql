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

-- name: UserExists :one

SELECT EXISTS( SELECT 1 FROM "user" WHERE "username" = $1 );

-- name: AddUser :exec

WITH _ AS (
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
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    )
INSERT INTO
    "shop" (
        "seller_name",
        "image_id",
        "name",
        "description",
        "enabled"
    )
VALUES ($1, $6, '', '', FALSE);

-- name: EnabledShop :exec

UPDATE "shop" AS s SET s."enabled" = TRUE WHERE s."seller_name" = $1;

-- name: DisableUser :exec

WITH disabled_user AS (
        UPDATE "user" AS u
        SET "enabled" = FALSE
        WHERE u."id" = $1
        RETURNING
            "username"
    ),
    disabled_shop AS (
        UPDATE "shop"
        SET "enabled" = FALSE
        WHERE "seller_name" = (
                SELECT
                    "username"
                FROM
                    disabled_user
            )
        RETURNING "id"
    )
UPDATE "product"
SET "enabled" = FALSE
WHERE "shop_id" = (
        SELECT "id"
        FROM disabled_shop
    );

-- there are some sql 🪄 happening here

-- name: DisableShop :exec

WITH disable_shop AS (
        UPDATE "shop" AS s
        SET s."enabled" = FALSE
        WHERE
            s."seller_name" = $1
        RETURNING s."id"
    )
UPDATE "product" AS p
SET p."enabled" = FALSE
WHERE p."shop_id" = (
        SELECT "id"
        FROM disable_shop
    );

-- name: DisableProductsFromShop :exec

UPDATE "product" AS p SET p."enabled" = FALSE WHERE p."shop_id" = $1;

-- name: CouponExists :one

SELECT EXISTS( SELECT 1 FROM "coupon" WHERE "id" = $1 );

-- name: GetAnyCoupons :many

SELECT * FROM "coupon";

-- name: GetCouponDetail :one

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
WHERE "id" = $1;

-- name: AddCoupon :one

INSERT INTO
    "coupon" (
        "type",
        "scope",
        "shop_id",
        "name",
        "description",
        "discount",
        "start_date",
        "expire_date"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date";

-- i don't feel right about this

-- name: EditCoupon :one

UPDATE "coupon"
SET
    "type" = COALESCE($2, "type"),
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "discount" = COALESCE($5, "discount"),
    "start_date" = COALESCE($6, "start_date"),
    "expire_date" = COALESCE($7, "expire_date")
WHERE "id" = $1
RETURNING
    "id",
    "type",
    "scope",
    "name",
    "description",
    "discount",
    "start_date",
    "expire_date";

-- name: DeleteCoupon :exec

WITH _ AS (
        DELETE FROM
            "cart_coupon"
        WHERE "coupon_id" = $1
    )
DELETE FROM "coupon"
WHERE "id" = $1;

-- name: GetReport :many

SELECT *
FROM "order_history"
WHERE
    "created_at" >= $1
    AND "created_at" <= $2;

-- name: GetUserIDByUsername :one

SELECT "id" FROM "user" WHERE "username" = $1;

-- name: GetShopIDBySellerName :one

SELECT "id" FROM "shop" WHERE "seller_name" = $1;
