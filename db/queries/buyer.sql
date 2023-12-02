-- name: GetOrderHistory :many

SELECT
    O."id",
    "shipment",
    "total_price",
    "status",
    "created_at"
FROM
    "order_history" AS O,
    "user" AS U
WHERE
    U."username" = $1
    AND U."id" = O."user_id"
ORDER BY "created_at" ASC
OFFSET $2
LIMIT $3;

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
