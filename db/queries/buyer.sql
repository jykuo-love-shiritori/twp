-- name: GetOrderHistory :many

SELECT
    O."id",
    s."name",
    s."image_id",
    "shipment",
    "total_price",
    "status",
    "created_at"
FROM
    "order_history" AS O,
    "user" AS U,
    "shop" AS S
WHERE
    U."username" = $1
    AND U."id" = O."user_id"
    AND O."shop_id" = S."id"

ORDER BY "created_at" ASC OFFSET $2 LIMIT $3;

-- name: GetOrderInfo :one

SELECT
    O."id",
    s."name",
    s."image_id",
    "shipment",
    "total_price",
    "status",
    "created_at"
FROM
    "order_history" AS O,
    "user" AS U,
    "shop" AS S
WHERE
    U."username" = $2
    AND O."id" = $1;

-- name: GetOrderDetail :many

SELECT
    O."product_id",
    P."name",
    P."description",
    P."price",
    P."image_id",
    O."quantity"
FROM
    "order_detail" AS O,
    "product_archive" AS P
WHERE
    O."order_id" = $1
    AND O."product_id" = P."id"
    AND O."product_version" = P."version";

-- name: GetCart :many

SELECT C."id", S."seller_name"
FROM
    "cart" AS C,
    "user" AS U,
    "shop" AS S
WHERE
    U."username" = $1
    AND U."id" = C."user_id"
    AND C."shop_id" = S."id";

-- name: GetProductInCart :many

SELECT
    "product_id",
    "quantity"
FROM "cart_product"
WHERE "cart_id" = $1;
