-- name: TestInsertUser :one
INSERT INTO "user" (
        "id",
        "username",
        "password",
        "name",
        "email",
        "address",
        "image_id",
        "role",
        "credit_card",
        "enabled"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING "id",
    "username",
    "password",
    "name",
    "email",
    "address",
    "image_id",
    "role",
    "credit_card",
    "enabled";
-- name: TestInsertShop :one
INSERT INTO "shop" (
        "id",
        "seller_name",
        "name",
        "image_id",
        "description",
        "enabled"
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: TestInsertCoupon :one
INSERT INTO "coupon" (
        "id",
        "type",
        "scope",
        "shop_id",
        "name",
        "description",
        "discount",
        "start_date",
        "expire_date"
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;
-- name: TestInsertProduct :one
INSERT INTO "product" (
        "id",
        "version",
        "shop_id",
        "name",
        "description",
        "price",
        "image_id",
        "expire_date",
        "edit_date",
        "stock",
        "sales",
        "enabled"
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        NOW(),
        $9,
        $10,
        $11
    )
RETURNING *;
-- name: TestInsertProductArchive :one
INSERT INTO "product_archive" (
        "id",
        "version",
        "name",
        "description",
        "price",
        "image_id"
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: TestInsertTag :one
INSERT INTO "tag" ("id", "shop_id", "name")
VALUES ($1, $2, $3)
RETURNING *;
-- name: TestInsertProductTag :one
INSERT INTO "product_tag" ("tag_id", "product_id")
VALUES ($1, $2)
RETURNING *;
-- name: TestInsertCouponTag :one
INSERT INTO "coupon_tag" ("tag_id", "coupon_id")
VALUES ($1, $2)
RETURNING *;
-- name: TestInsertCart :one
INSERT INTO "cart" ("id", "user_id", "shop_id")
VALUES ($1, $2, $3)
RETURNING *;
-- name: TestInsertCartProduct :one
INSERT INTO "cart_product" (
        "cart_id",
        "product_id",
        "quantity"
    )
VALUES ($1, $2, $3)
RETURNING *;
-- name: TestInsertCartCoupon :one
INSERT INTO "cart_coupon" ("cart_id", "coupon_id")
VALUES ($1, $2)
RETURNING *;
-- name: TestInsertOrderHistory :one
INSERT INTO "order_history" (
        "id",
        "user_id",
        "shop_id",
        "shipment",
        "total_price",
        "status"
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: TestInsertOrderDetail :one
INSERT INTO "order_detail" (
        "order_id",
        "product_id",
        "product_version",
        "quantity"
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: TestDeleteUserById :one
DELETE FROM "user"
WHERE "id" = $1
RETURNING "id",
    "username",
    "password",
    "name",
    "email",
    "address",
    "image_id",
    "role",
    "credit_card",
    "enabled";
-- name: TestDeleteShopById :one
DELETE FROM "shop"
WHERE "id" = $1
RETURNING *;
-- name: TestDeleteCouponById :one
DELETE FROM "coupon"
WHERE "id" = $1
RETURNING *;
-- name: TestDeleteProductById :one
DELETE FROM "product"
WHERE "id" = $1
RETURNING *;
-- name: TestDeleteOrderDetailByOrderId :one
DELETE FROM "order_detail"
WHERE "order_id" = $1
    AND "product_id" = $2
    AND "product_version" = $3
RETURNING *;
-- name: TestDeleteOrderById :one
DELETE FROM "order_history"
WHERE "id" = $1
RETURNING *;
-- name: TestDeleteProductArchiveByIdVersion :one
DELETE FROM "product_archive"
WHERE "id" = $1
    AND "version" = $2
RETURNING *;
-- name: TestDeleteTagById :one
DELETE FROM "tag"
WHERE "id" = $1
RETURNING *;
-- name: TestDeleteCoupon :exec
DELETE FROM "coupon";
-- name: TestDeleteProduct :exec
DELETE FROM "product";
-- name: TestDeleteOrderDetail :exec
DELETE FROM "order_detail";
-- name: TestDeleteTag :exec
DELETE FROM "tag";
-- name: TestDeleteOrderHistory :exec
DELETE FROM "order_history";
-- name: TestDeleteProductArchive :exec
DELETE FROM "product_archive";
-- name: TestDeleteShop :exec
DELETE FROM "shop";
-- name: TestDeleteUser :exec
DELETE FROM "user";
