-- name: GetOrderHistory :many
SELECT
    O."id",
    s."name" AS "shop_name",
    s."image_id" AS "shop_image_url",
    O."image_id" AS "thumbnail_url",
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
ORDER BY
    "created_at" ASC OFFSET $2
LIMIT $3;

-- name: GetOrderInfo :one
SELECT
    O."id",
    S."name" AS "shop_name",
    S."image_id" AS "shop_image_url",
    "shipment",
    "total_price",
    "status",
    "created_at",
(T."subtotal" + "shipment" - "total_price") AS "discount"
FROM
    "order_history" AS O,
    "order_detail" AS D,
    "user" AS U,
    "shop" AS S,
(
        SELECT
            SUM(P."price" * D."quantity")::INT AS "subtotal"
        FROM
            "order_detail" AS D,
            "product_archive" AS P
        WHERE
            D."order_id" = $1
            AND D."product_id" = P."id"
            AND D."product_version" = P."version") AS T
WHERE
    U."username" = $2
    AND O."id" = $1
    AND O."user_id" = U."id"
    AND O."id" = D."order_id"
    AND O."shop_id" = S."id";

-- name: GetOrderDetail :many
SELECT
    O."product_id",
    P."name",
    P."description",
    P."price",
    P."image_id" AS "image_url",
    O."quantity"
FROM
    "order_detail" AS O,
    "product_archive" AS P
WHERE
    O."order_id" = $1
    AND O."product_id" = P."id"
    AND O."product_version" = P."version";

-- name: GetCart :many
SELECT
    C."id",
    S."seller_name",
    S."image_id" AS "shop_image_url",
    S."name" AS "shop_name"
FROM
    "cart" AS C,
    "user" AS U,
    "shop" AS S
WHERE
    U."username" = $1
    AND U."id" = C."user_id"
    AND C."shop_id" = S."id";

-- name: GetCouponFromCart :many
SELECT
    C."id",
    C."name",
    "type",
    "scope",
    "description",
    "discount",
    "expire_date"
FROM
    "cart_coupon" AS CC,
    "coupon" AS C,
    "cart" AS 🛒,
    "user" AS U
WHERE
    U."username" = $1
    AND U."id" = 🛒."user_id"
    AND 🛒."id" = @cart_id
    AND CC."cart_id" = 🛒."id"
    AND CC."coupon_id" = C."id";

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
FROM
    "coupon"
WHERE
    "id" = $1;

-- name: GetProductFromCart :many
SELECT
    "product_id",
    "name",
    "image_id" AS "image_url",
    "price",
    "quantity",
    "stock",
    "enabled"
FROM
    "cart_product" AS C,
    "product" AS P
WHERE
    "cart_id" = $1
    AND C."product_id" = P."id";

-- name: UpdateProductFromCart :execrows
UPDATE
    "cart_product" AS CP
SET
    "quantity" = $3
FROM
    "user" AS U,
    "cart" AS C,
    "product" AS P
WHERE
    U."username" = $4
    AND U."id" = C."user_id"
    AND CP."cart_id" = $1
    AND CP."product_id" = $2
    AND $3 <= P."stock"
    AND P."id" = $2
    AND P."enabled" = TRUE
    AND C."id" = CP."cart_id";

-- name: DeleteProductFromCart :one
WITH valid_cart AS (
    SELECT
        C."id"
    FROM
        "cart" C
        JOIN "user" u ON u."id" = C."user_id"
    WHERE
        u."username" = $1
        AND C."id" = @cart_id
),
deleted_products AS (
    DELETE FROM "cart_product" CP
    WHERE "cart_id" =(
            SELECT
                "id"
            FROM
                valid_cart)
            AND CP."product_id" = @product_id
        RETURNING
            *
)
SELECT
    EXISTS (
        SELECT
            1
        FROM
            deleted_products);

-- name: DeleteEmptyCart :exec
WITH valid_cart AS (
    SELECT
        C."id"
    FROM
        "cart" C
        JOIN "user" u ON u."id" = C."user_id"
    WHERE
        u."username" = $1
        AND C."id" = @cart_id
        AND NOT EXISTS (
            SELECT
                1
            FROM
                "cart_product" CP
            WHERE
                CP."cart_id" = C."id")
),
delete_coupon AS (
    DELETE FROM "cart_coupon" AS CC
    WHERE "cart_id" =(
            SELECT
                "id"
            FROM
                valid_cart))
DELETE FROM "cart" AS C
WHERE "id" =(
        SELECT
            "id"
        FROM
            valid_cart);

-- name: GetUsableCoupons :many
SELECT
    C."id",
    C."name",
    "type",
    "scope",
    "description",
    "discount",
    "expire_date"
FROM
    "coupon" AS C,
    "cart" AS 🛒,
    "user" AS U
WHERE
    U."username" = $1
    AND U."id" = 🛒."user_id"
    AND 🛒."id" = @cart_id
    AND (C."scope" = 'global'
        OR (C."scope" = 'shop'
            AND C."shop_id" = 🛒."shop_id"))
    AND NOW() BETWEEN C."start_date" AND C."expire_date"
    AND NOT EXISTS (
        SELECT
            1
        FROM
            "cart_coupon" AS CC
        WHERE
            CC."cart_id" = 🛒."id"
            AND CC."coupon_id" = C."id");

-- name: AddProductToCart :one
WITH valid_product AS (
    SELECT
        P."id" AS product_id,
        S."id" AS shop_id
    FROM
        "product" P,
        "shop" S
    WHERE
        P."shop_id" = S."id"
        AND P."id" = $2
        AND P."enabled" = TRUE
        AND P."stock" >= @quantity
),
-- check product enabled ⬆️
new_cart AS (
INSERT INTO "cart"("user_id", "shop_id")
    SELECT
        U."id",
        S."id"
    FROM
        "user" AS U,
        "shop" AS S,
        "product" AS P
    WHERE
        U."username" = $1
        AND S."id" = P."shop_id"
        AND P."id" = $2
        AND NOT EXISTS (
            SELECT
                1
            FROM
                "cart" AS C
            WHERE
                C."user_id" = U."id"
                AND C."shop_id" = S."id")
        RETURNING
            *
),
existed_cart AS (
    SELECT
        C."id"
    FROM
        "cart" AS C,
        "user" AS U,
        "shop" AS S
    WHERE
        U."username" = $1
        AND C."user_id" = U."id"
        AND C."shop_id" = S."id"
        AND S."id" =(
            SELECT
                "shop_id"
            FROM
                valid_product)
),
cart_id AS (
    SELECT
        "id"
    FROM
        new_cart
    UNION ALL
    SELECT
        "id"
    FROM
        existed_cart
),
insert_product AS (
INSERT INTO "cart_product"("cart_id", "product_id", "quantity")
    SELECT
        (
            SELECT
                "id"
            FROM
                cart_id),
(
                SELECT
                    "product_id"
                FROM
                    valid_product),
                @quantity
            ON CONFLICT ("cart_id",
                "product_id")
                DO UPDATE SET
                    "quantity" = "cart_product"."quantity" + @quantity
                RETURNING
                    *
)
        SELECT
            COALESCE(SUM(CP."quantity"), 0) + @quantity AS total_quantity
        FROM
            "cart_product" AS CP,
            "cart" AS C,
            "user" AS U
        WHERE
            U."username" = $1
            AND C."user_id" = U."id"
            AND CP."cart_id" = C."id";

-- returning the number of products in any cart for US-SC-2 in SRS ⬆️
-- name: GetProductTag :many
SELECT
    "tag_id"
FROM
    "product_tag" AS PT,
    "tag" AS T
WHERE
    PT."product_id" = $1
    AND PT."tag_id" = T."id";

-- name: GetCouponTag :many
SELECT
    "tag_id"
FROM
    "coupon_tag" AS CT,
    "tag" AS T
WHERE
    CT."coupon_id" = $1
    AND CT."tag_id" = T."id";

-- name: AddCouponToCart :execrows
INSERT INTO "cart_coupon"("cart_id", "coupon_id")
SELECT
    C."id",
    CO."id"
FROM
    "cart" AS C,
    "user" AS U,
    "coupon" AS CO
WHERE
    U."username" = $1
    AND C."user_id" = U."id"
    AND C."id" = @Cart_id
    AND (CO."scope" = 'global'
        OR (CO."scope" = 'shop'
            AND CO."shop_id" = C."shop_id"))
    AND NOW() BETWEEN CO."start_date" AND CO."expire_date"
    AND NOT EXISTS (
        SELECT
            1
        FROM
            "cart_coupon" AS CC
        WHERE
            CC."cart_id" = C."id"
            AND CC."coupon_id" = $2)
    AND CO."id" = $2;

-- name: ValidateProductsInCart :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            "cart_product" AS CP,
            "product" AS P,
            "cart" AS C,
            "user" AS U
        WHERE
            C."id" = CP."cart_id"
            AND CP."product_id" = P."id"
            AND C."id" = @cart_id
            AND C."user_id" = U."id"
            AND U."username" = $1
            AND P."enabled" = TRUE
            AND CP."quantity" <= P."stock");

-- name: DeleteCouponFromCart :execrows
DELETE FROM "cart_coupon" AS CC USING "cart" AS C, "user" AS U
WHERE U."username" = $1
    AND C."user_id" = U."id"
    AND C."id" = CC."cart_id"
    AND C."id" = @cart_id
    AND CC."coupon_id" = $2;

-- name: GetCouponsFromCart :many
WITH delete_expire_coupons AS (
    DELETE FROM "cart_coupon" AS CC USING "coupon" AS CO, "cart" AS C, "user" AS U
WHERE U."username" = $1
    AND C."user_id" = U."id"
    AND C."id" = CC."cart_id"
    AND C."id" = @cart_id
    AND CC."coupon_id" = CO."id"
    AND NOW() > CO."expire_date"
)
SELECT
    CO."id",
    CO."name",
    CO."type",
    CO."scope",
    CO."description",
    CO."discount"
FROM
    "cart_coupon" AS CC,
    "coupon" AS CO,
    "cart" AS C,
    "user" AS U
WHERE
    U."username" = $1
    AND C."user_id" = U."id"
    AND C."id" = CC."cart_id"
    AND C."id" = @cart_id
    AND CC."coupon_id" = CO."id";

-- name: Checkout :exec
WITH insert_order AS (
INSERT INTO "order_history"("user_id", "shop_id", "image_id", "shipment", "total_price", "status")
    SELECT
        U."id",
        S."id",
        T."image_id",
        $2,
        $3,
        'paid'
    FROM
        "user" AS U,
        "shop" AS S,
        "cart" AS C,
(
            SELECT
                "image_id"
            FROM
                "product"
            WHERE
                "id" =(
                    SELECT
                        "product_id"
                    FROM
                        "cart_product"
                    WHERE
                        "cart_id" = @cart_id
                    ORDER BY
                        "price" DESC
                    LIMIT 1 -- the most expensive product's image_id will be used as the thumbnail ↙️
)) AS T
    WHERE
        U."username" = $1
        AND U."id" = C."user_id"
        AND C."id" = @cart_id
        AND S."id" = C."shop_id"
    RETURNING
        "id"
),
delete_cart AS (
    DELETE FROM "cart" AS C
    WHERE C."id" = @cart_id
),
add_sales AS (
    UPDATE
        "product" AS P
    SET
        "sales" = "sales" + CP."quantity",
        "stock" = "stock" - CP."quantity"
    FROM
        "cart_product" AS CP
    WHERE
        CP."cart_id" = @cart_id
        AND CP."product_id" = P."id"
        AND CP."quantity" <= P."stock")
INSERT INTO "order_detail"("order_id", "product_id", "product_version", "quantity")
SELECT
    (
        SELECT
            "id"
        FROM
            insert_order),
    CP."product_id",
    P."version",
    CP."quantity"
FROM
    "cart_product" AS CP,
    "product" AS P,
    "cart" AS C,
    "user" AS U
WHERE
    C."id" = CP."cart_id"
    AND CP."product_id" = P."id"
    AND C."id" = @cart_id
    AND C."user_id" = U."id";

-- name: UpdateProductVersion :exec
WITH check_version AS (
    SELECT
        EXISTS (
            SELECT
                1
            FROM
                "product" P,
                "product_archive" PA
            WHERE
                P."id" = $1
                AND P."id" = PA."id"
                AND P."version" = PA."version"
                AND P."name" = PA."name"
                AND P."description" = PA."description"
                AND P."price" = PA."price"
                AND P."image_id" = PA."image_id") AS "version_existed",
        EXISTS (
            SELECT
                1
            FROM
                "product_archive" PA
            WHERE
                PA."id" = $1) AS "product_existed"
),
insert_archive AS (
INSERT INTO "product_archive"("id", "version", "name", "description", "price", "image_id")
    SELECT
        P."id",
        P."version",
        P."name",
        P."description",
        P."price",
        P."image_id"
    FROM
        "product" AS P,
        check_version AS CV
    WHERE
        P."id" = $1
        AND CV."version_existed" = FALSE)
UPDATE
    "product"
SET
    "version" = "version" + 1
FROM
    check_version CV
WHERE
    "product"."id" = $1
    AND CV."version_existed" = FALSE
    AND CV."product_existed" = TRUE;

-- name: GetCreditCard :one
SELECT
    "credit_card"
FROM
    "user"
WHERE
    "username" = $1;

-- name: ValidatePayment :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            "user"
        WHERE
            "username" = $1
            AND ("credit_card" = $2
                OR random() < 0.999 -- random validate payment
));

-- name: UpdateOrderStatus :execrows
UPDATE
    "order_history"
SET
    "status" = @status
WHERE
    "id" = @id
    AND "user_id" =(
        SELECT
            "id"
        FROM
            "user"
        WHERE
            "username" = @username)
    AND "status" = 'delivered';
