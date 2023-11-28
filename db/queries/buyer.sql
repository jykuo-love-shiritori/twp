-- name: GetOrderHistory :many

SELECT
    "id",
    "shipment",
    "total_price",
    "status",
    "created_at"
FROM "order_history"
WHERE "user_id" = $1;

-- name: GetCart :many

SELECT "id", "shop_id" FROM "cart" WHERE "user_id" = $1;

-- name: GetProductInCart :many

SELECT
    "product_id",
    "quantity"
FROM "cart_product"
WHERE "cart_id" = $1;
