-- name: SellerGetInfo :one
SELECT "name",
    "image_id" AS "image_url",
    "description",
    "enabled"
FROM "shop"
WHERE "seller_name" = $1;
-- name: SellerUpdateInfo :one
UPDATE "shop"
SET "image_id" = CASE
        WHEN sqlc.arg('image_id')::TEXT = '' THEN "image_id"
        ELSE sqlc.arg('image_id')::TEXT
    END,
    "name" = COALESCE($2, "name"),
    "description" = COALESCE($3, "description"),
    "enabled" = COALESCE($4, "enabled")
WHERE "seller_name" = $1
RETURNING "name",
    "image_id" AS "image_url",
    "description",
    "enabled";
-- name: SellerSearchTag :many
SELECT t."id",
    t."name"
FROM "tag" t
    LEFT JOIN "shop" s ON "shop_id" = s.id
WHERE s."seller_name" = $1
    AND t."name" ~* $2
ORDER BY LENGTH(t."name") ASC
LIMIT $3;
-- name: HaveTagName :one
SELECT EXISTS (
        SELECT 1
        FROM "tag" t
            LEFT JOIN "shop" s ON "shop_id" = s.id
        WHERE s."seller_name" = $1
            AND t."name" = $2
    );
-- name: SellerInsertTag :one
INSERT INTO "tag"("shop_id", "name")
VALUES (
        (
            SELECT s."id"
            FROM "shop" s
            WHERE s."seller_name" = $1
                AND s."enabled" = TRUE
        ),
        $2
    )
RETURNING "id",
    "name";
-- name: SellerGetCoupon :many
SELECT c."id",
    c."type",
    c."scope",
    c."name",
    c."description",
    c."discount",
    c."start_date",
    c."expire_date"
FROM "coupon" c
    JOIN "shop" s ON c."shop_id" = s.id
WHERE s.seller_name = $1
ORDER BY "start_date" DESC
LIMIT $2 OFFSET $3;
-- name: SellerGetCouponDetail :one
SELECT c."type",
    c."scope",
    c."name",
    c."discount",
    c."description",
    c."start_date",
    c."expire_date"
FROM "coupon" c
    JOIN "shop" s ON c."shop_id" = s.id
WHERE s."seller_name" = $1
    AND c."id" = $2;
-- name: SellerInsertCoupon :one
INSERT INTO "coupon"(
        "type",
        "scope",
        "shop_id",
        "name",
        "description",
        "discount",
        "start_date",
        "expire_date"
    )
VALUES (
        $2,
        'shop',
        (
            SELECT s."id"
            FROM "shop" s
            WHERE s."seller_name" = $1
                AND s."enabled" = TRUE
        ),
        $3,
        $4,
        $5,
        $6,
        $7
    )
RETURNING "id",
    "type",
    "scope",
    "name",
    "discount",
    "description",
    "start_date",
    "expire_date";
-- name: SellerUpdateCouponInfo :one
UPDATE "coupon" c
SET "type" = COALESCE($3, "type"),
    "name" = COALESCE($4, "name"),
    "description" = COALESCE($5, "description"),
    "discount" = COALESCE($6, "discount"),
    "start_date" = COALESCE($7, "start_date"),
    "expire_date" = COALESCE($8, "expire_date")
WHERE c."id" = $2
    AND "shop_id" =(
        SELECT s."id"
        FROM "shop" s
        WHERE s."seller_name" = $1
            AND s."enabled" = TRUE
    )
RETURNING c."id",
    c."type",
    c."scope",
    c."name",
    c."description",
    c."discount",
    c."start_date",
    c."expire_date";
-- name: SellerDeleteCoupon :execrows
DELETE FROM "coupon" c
WHERE c."id" = $2
    AND "shop_id" =(
        SELECT s."id"
        FROM "shop" s
        WHERE s."seller_name" = $1
            AND s."enabled" = TRUE
    );
-- name: SellerGetOrder :many
SELECT oh."id",
    op."product_name",
    op."thumbnail_url",
    u."name" AS "user_name",
    u."image_id" AS "user_image_url",
    oh."shipment",
    oh."total_price",
    oh."status",
    oh."created_at"
FROM "order_history" AS oh
    INNER JOIN "shop" AS s ON oh."shop_id" = s."id"
    INNER JOIN "user" AS u ON oh."user_id" = u."id"
    LEFT JOIN (
        SELECT od."order_id",
            pa."name" AS "product_name",
            pa."image_id" AS "thumbnail_url",
            ROW_NUMBER() OVER (
                PARTITION BY od."order_id"
                ORDER BY pa."price" DESC
            ) AS rn
        FROM "order_detail" AS od
            INNER JOIN "product_archive" AS pa ON od."product_id" = pa."id"
            AND od."product_version" = pa."version"
        ORDER BY pa."price" DESC
    ) AS op ON oh."id" = op."order_id"
    AND op.rn = 1
WHERE s."seller_name" = $1
ORDER BY "created_at" DESC
LIMIT $2 OFFSET $3;
-- name: SellerGetOrderHistory :one
SELECT "order_history"."id",
    "order_history"."shipment",
    "order_history"."total_price",
    "order_history"."status",
    "order_history"."created_at",
    "user"."id" AS "user_id",
    "user"."name" AS "user_name",
    "user"."image_id" AS "user_image_url"
FROM "order_history"
    JOIN shop ON "order_history".shop_id = shop.id
    JOIN "user" ON "order_history".user_id = "user"."id"
WHERE shop.seller_name = $1
    AND "order_history".id = $2;
-- name: SellerGetOrderDetail :many
SELECT product_archive."id",
    product_archive."name",
    product_archive."description",
    product_archive."price",
    product_archive."image_id" AS "image_url",
    order_detail.quantity
FROM "order_detail"
    LEFT JOIN product_archive ON order_detail.product_id = product_archive.id
    AND order_detail.product_version = product_archive.version
    LEFT JOIN order_history ON order_history.id = order_detail.order_id
    LEFT JOIN shop ON order_history.shop_id = shop.id
WHERE shop.seller_name = $1
    AND order_detail.order_id = $2
ORDER BY quantity * price DESC;
-- name: SellerUpdateOrderStatus :one
UPDATE "order_history" oh
SET "status" = sqlc.arg(set_status)
WHERE "shop_id" =(
        SELECT s."id"
        FROM "shop" s
        WHERE s."seller_name" = $1
            AND s."enabled" = TRUE
    )
    AND oh."id" = $2
    AND oh."status" = sqlc.arg(current_status)
RETURNING oh."id",
    oh."shipment",
    oh."total_price",
    oh."status",
    oh."created_at";
-- name: SellerBestSellProduct :many
SELECT order_detail.product_id,
    product_archive.name,
    product_archive.price,
    product_archive.image_id AS "image_url",
    SUM(order_detail.quantity) AS total_quantity,
    SUM(order_detail.quantity * product_archive.price)::decimal(10, 2) AS total_sell,
    COUNT(order_history.id) AS order_count
FROM "order_detail"
    LEFT JOIN product_archive ON order_detail.product_id = product_archive.id
    AND order_detail.product_version = product_archive.version
    LEFT JOIN order_history ON order_history.id = order_detail.order_id
    LEFT JOIN shop ON order_history.shop_id = shop.id
WHERE shop.seller_name = $1
    AND order_history."created_at" > sqlc.arg('time')
    AND order_history."created_at" < sqlc.arg('time') + INTERVAL '1 month'
GROUP BY product_archive.id,
    product_archive.description,
    product_archive.name,
    product_archive.price,
    product_archive.image_id,
    order_detail.product_id
ORDER BY total_quantity DESC
LIMIT $2;
-- name: SellerReport :one
SELECT SUM(order_history.total_price)::decimal(10, 2) AS total_income,
    COUNT(order_history.id) AS order_count
FROM order_history
    LEFT JOIN shop ON order_history.shop_id = shop.id
WHERE shop.seller_name = $1
    AND order_history."created_at" > sqlc.arg('time')
    AND order_history."created_at" < sqlc.arg('time') + INTERVAL '1 month';
-- name: SellerGetProductDetail :one
SELECT p."name",
    p."description",
    p."image_id" as "image_url",
    p."price",
    p."sales",
    p."stock",
    p."enabled"
FROM "product" p
    JOIN "shop" s ON p."shop_id" = s.id
WHERE s.seller_name = $1
    AND p."id" = $2;
-- name: SellerProductList :many
SELECT p."id",
    p."name",
    p."image_id" AS "image_url",
    p."price",
    p."sales",
    p."stock",
    p."enabled"
FROM "product" p
    JOIN "shop" s ON p."shop_id" = s.id
WHERE s.seller_name = $1
ORDER BY "sales" DESC
LIMIT $2 OFFSET $3;
-- name: SellerCheckTags :one
SELECT NOT EXISTS (
        SELECT 1
        FROM unnest(@tags::INT []) t
            LEFT JOIN "tag" ON t = "tag"."id"
            LEFT JOIN "shop" s ON "tag"."shop_id" = s."id"
        WHERE "tag"."id" IS NULL
            OR s."seller_name" != $1
    );
-- name: SellerInsertProductTags :exec
INSERT INTO "product_tag"("product_id", "tag_id")
VALUES ($1, unnest(@tags::INT []));
-- name: SellerInsertCouponTags :exec
INSERT INTO "coupon_tag"("coupon_id", "tag_id")
VALUES ($1, unnest(@tags::INT []));
-- name: SellerInsertProduct :one
INSERT INTO "product"(
        "version",
        "shop_id",
        "name",
        "description",
        "price",
        "image_id",
        "expire_date",
        "edit_date",
        "stock",
        "enabled"
    )
VALUES (
        1,
        (
            SELECT s."id"
            FROM "shop" s
            WHERE s."seller_name" = $1
                AND s."enabled" = TRUE
        ),
        $2,
        $3,
        $4,
        $5,
        $6,
        NOW(),
        $7,
        $8
    )
RETURNING "id",
    "name",
    "description",
    "price",
    "image_id" AS "image_url",
    "expire_date",
    "edit_date",
    "stock",
    "sales",
    "enabled";
-- name: SellerUpdateProductInfo :one
UPDATE "product" p
SET "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "price" = COALESCE($5, "price"),
    "image_id" = CASE
        WHEN sqlc.arg('image_id')::TEXT = '' THEN "image_id"
        ELSE sqlc.arg('image_id')::TEXT
    END,
    "expire_date" = COALESCE($6, "expire_date"),
    "enabled" = COALESCE($7, "enabled"),
    "stock" = COALESCE($8, "stock"),
    "edit_date" = NOW(),
    "version" = "version" + 1
WHERE "shop_id" =(
        SELECT s."id"
        FROM "shop" s
        WHERE s."seller_name" = $1
            AND s."enabled" = TRUE
    )
    AND p."id" = $2
RETURNING "id",
    "name",
    "description",
    "price",
    "image_id" AS "image_url",
    "expire_date",
    "edit_date",
    "stock",
    "sales",
    "enabled";
-- name: SellerDeleteProduct :execrows
DELETE FROM "product" p
WHERE "shop_id" =(
        SELECT s."id"
        FROM "shop" s
        WHERE s."seller_name" = $1
            AND s."enabled" = TRUE
    )
    AND p."id" = $2;
-- name: SellerGetProductTag :many
SELECT pt."tag_id",
    t."name"
FROM "product_tag" pt
    JOIN "product" p ON p."id" = pt."product_id"
    JOIN "shop" s ON s."id" = p."shop_id"
    JOIN "tag" t ON t."id" = pt."tag_id"
WHERE s."seller_name" = $1
    AND "product_id" = $2
    AND s."enabled" = TRUE;
-- name: SellerGetCouponTag :many
SELECT ct."tag_id",
    t."name"
FROM "coupon_tag" ct
    JOIN "coupon" c ON c."id" = ct."coupon_id"
    JOIN "shop" s ON s."id" = c."shop_id"
    JOIN "tag" t ON t."id" = ct."tag_id"
WHERE s."seller_name" = $1
    AND "coupon_id" = $2;
-- name: SellerInsertProductTag :one
INSERT INTO "product_tag"("tag_id", "product_id")
SELECT $2,
    $3
WHERE EXISTS (
        SELECT 1
        FROM "tag" t
            JOIN "shop" s ON s."id" = t."shop_id"
        WHERE s."seller_name" = $1
            AND t."id" = $2
            AND s."enabled" = TRUE
    )
    AND EXISTS (
        SELECT 1
        FROM "product" p
            JOIN "shop" s ON s."id" = p."shop_id"
        WHERE s."seller_name" = $1
            AND p."id" = $3
    )
RETURNING *;
-- name: SellerDeleteProductTag :execrows
DELETE FROM "product_tag" tp
WHERE EXISTS (
        SELECT 1
        FROM "product" p
            JOIN "shop" s ON s."id" = p."shop_id"
        WHERE s."seller_name" = $1
            AND p."id" = $3
            AND s."enabled" = TRUE
    )
    AND "product_id" = $3
    AND "tag_id" = $2;
-- name: SellerInsertCouponTag :one
INSERT INTO "coupon_tag"("tag_id", "coupon_id")
SELECT $2,
    $3
WHERE EXISTS (
        SELECT 1
        FROM "tag" t
            JOIN "shop" s ON s."id" = t."shop_id"
        WHERE s."seller_name" = $1
            AND s."enabled" = TRUE
            AND t."id" = $2
    )
    AND EXISTS (
        SELECT 1
        FROM "coupon" c
            JOIN "shop" s ON s."id" = c."shop_id"
        WHERE s."seller_name" = $1
            AND s."enabled" = TRUE
            AND c."id" = $3
    )
RETURNING *;
-- name: SellerDeleteCouponTag :execrows
DELETE FROM "coupon_tag" tp
WHERE EXISTS (
        SELECT 1
        FROM "coupon" c
            JOIN "shop" s ON s."id" = c."shop_id"
        WHERE s."seller_name" = $1
            AND s."enabled" = TRUE
            AND c."id" = $3
    )
    AND "coupon_id" = $3
    AND "tag_id" = $2;
