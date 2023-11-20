-- name: BuyerGetOrder :many

SELECT * FROM "order_history" WHERE "id" = $1 and "user_id" = $2;

-- name: BuyerCart :many

SELECT * FROM "cart" WHERE "user_id" = $1;