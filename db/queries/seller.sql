-- name: SellerGetInfo :one

SELECT
    "seller_name",
    "image_id",
    "description",
    "enabled"
FROM "shop"
WHERE "seller_name" = $1;

-- name: SellerUpdateInfo :one

UPDATE "shop"
SET
    "image_id" = COALESCE($2, "image_id"),
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "enabled" = COALESCE($5, "enabled")
WHERE "seller_name" = $1
RETURNING
    "seller_name",
    "image_id",
    "name",
    "enabled";

-- name: SellerSearchTag :many

SELECT t."id", t."name"
FROM "tag" t
    LEFT JOIN "shop" s ON "shop_id" = s.id
WHERE
    s."seller_name" = $1
    AND t."name" ~* $2
ORDER BY LENGTH(t."name") ASC
LIMIT $3;

-- name: HaveTagName :one

SELECT EXISTS (
        SELECT 1
        FROM "tag" t
            LEFT JOIN "shop" s ON "shop_id" = s.id
        WHERE
            s."seller_name" = $1
            AND t."name" = $2
    );

-- name: SellerInsertTag :one

INSERT INTO
    "tag" ("shop_id", "name")
VALUES ( (
            SELECT s."id"
            FROM "shop" s
            WHERE
                s."seller_name" = $1
                AND s."enabled" = true
        ),
        $2
    )
RETURNING "id", "name";

-- name: SellerGetCoupon :many

SELECT
    c."id",
    c."type",
    c."scope",
    c."name",
    c."discount",
    c."expire_date"
FROM "coupon" c
    JOIN "shop" s ON c."shop_id" = s.id
WHERE s.seller_name = $1
ORDER BY "start_date" DESC
LIMIT $2
OFFSET $3;

-- name: SellerGetCouponDetail :one

SELECT
    c."type",
    c."scope",
    c."name",
    c."discount",
    c."expire_date"
FROM "coupon" c
    JOIN "shop" s ON c."shop_id" = s.id
WHERE
    s."seller_name" = $1
    AND c."id" = $2;

-- name: SellerInsertCoupon :one

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
VALUES (
        $2, 'shop', (
            SELECT s."id"
            FROM "shop" s
            WHERE
                s."seller_name" = $1
                AND s."enabled" = true
        ),
        $3,
        $4,
        $5,
        $6,
        $7
    )
RETURNING
    "id",
    "type",
    "scope",
    "name",
    "discount",
    "expire_date";

-- name: SellerUpdateCouponInfo :one

UPDATE "coupon" c
SET
    "type" = COALESCE($3, "type"),
    "name" = COALESCE($4, "name"),
    "description" = COALESCE($5, "description"),
    "discount" = COALESCE($6, "discount"),
    "start_date" = COALESCE($7, "start_date"),
    "expire_date" = COALESCE($8, "expire_date")
WHERE c."id" = $2 AND "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
RETURNING
    c."id",
    c."type",
    c."scope",
    c."name",
    c."discount",
    c."expire_date";

-- name: SellerDeleteCoupon :execrows

DELETE FROM "coupon" c
WHERE c."id" = $2 AND "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    );

-- name: SellerGetOrder :many

SELECT
    "id",
    "shipment",
    "total_price",
    "status",
    "created_at"
FROM "order_history"
WHERE "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
    )
ORDER BY "created_at" DESC
LIMIT $2
OFFSET $3;

-- name: SellerGetOrderHistory :one

SELECT
    "order_history"."id",
    "order_history"."shipment",
    "order_history"."total_price",
    "order_history"."status",
    "order_history"."created_at"
FROM "order_history"
    JOIN shop ON order_history.shop_id = shop.id
WHERE
    shop.seller_name = $1
    AND order_history.id = $2;

-- name: SellerGetOrderDetail :many

SELECT
    product_archive.*,
    order_detail.quantity
FROM "order_detail"
    LEFT JOIN product_archive ON order_detail.product_id = product_archive.id AND order_detail.product_version = product_archive.version
    LEFT JOIN order_history ON order_history.id = order_detail.order_id
    LEFT JOIN shop ON order_history.shop_id = shop.id
WHERE
    shop.seller_name = $1
    AND order_detail.order_id = $2
ORDER BY quantity * price DESC;

-- name: SellerUpdateOrderStatus :one

UPDATE "order_history" oh
SET
    "status" = sqlc.arg(set_status)
WHERE "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
    AND oh."id" = $2
    AND oh."status" = sqlc.arg(current_status)
RETURNING
    oh."id",
    oh."shipment",
    oh."total_price",
    oh."status",
    oh."created_at";

-- TODO

-- SellerGetReport :many

-- SellerGetReportDetail :many

-- name: SellerGetProductDetail :one

SELECT
    p."name",
    p."image_id",
    p."price",
    p."sales",
    p."stock",
    p."enabled"
FROM "product" p
    JOIN "shop" s ON p."shop_id" = s.id
WHERE
    s.seller_name = $1
    AND p."id" = $2;

-- name: SellerProductList :many

SELECT
    p."id",
    p."name",
    p."image_id",
    p."price",
    p."sales",
    p."stock",
    p."enabled"
FROM "product" p
    JOIN "shop" s ON p."shop_id" = s.id
WHERE s.seller_name = $1
ORDER BY "sales" DESC
LIMIT $2
OFFSET $3;

-- name: SellerInsertProduct :one

INSERT INTO
    "product"(
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
        1, (
            SELECT s."id"
            FROM "shop" s
            WHERE
                s."seller_name" = $1
                AND s."enabled" = true
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
RETURNING
    "id",
    "name",
    "description",
    "price",
    "image_id",
    "expire_date",
    "edit_date",
    "stock",
    "sales";

-- name: SellerUpdateProductInfo :one

UPDATE "product" p
SET
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "price" = COALESCE($5, "price"),
    "image_id" = COALESCE($6, "image_id"),
    "expire_date" = COALESCE($7, "expire_date"),
    "enabled" = COALESCE($8, "enabled"),
    "stock" = COALESCE($9, "stock"),
    "edit_date" = NOW(),
    "version" = "version" + 1
WHERE "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
    AND p."id" = $2
RETURNING
    "id",
    "name",
    "description",
    "price",
    "image_id",
    "expire_date",
    "edit_date",
    "stock",
    "sales";

-- name: SellerDeleteProduct :execrows

DELETE FROM "product" p
WHERE "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
    AND p."id" = $2;

-- name: SellerGetProductTag :many

SELECT pt."tag_id", t."name"
FROM "product_tag" pt
    JOIN "product" p ON p."id" = pt."product_id"
    JOIN "shop" s ON s."id" = p."shop_id"
    JOIN "tag" t ON t."id" = pt."tag_id"
WHERE
    s."seller_name" = $1
    AND "product_id" = $2
    AND s."enabled" = true;

-- name: SellerGetCouponTag :many

SELECT ct."tag_id", t."name"
FROM "coupon_tag" ct
    JOIN "coupon" c ON c."id" = ct."coupon_id"
    JOIN "shop" s ON s."id" = c."shop_id"
    JOIN "tag" t ON t."id" = ct."tag_id"
WHERE
    s."seller_name" = $1
    AND "coupon_id" = $2;

-- name: SellerInsertProductTag :one

INSERT INTO
    "product_tag" ("tag_id", "product_id")
SELECT $2, $3
WHERE EXISTS (
        SELECT 1
        FROM "tag" t
            JOIN "shop" s ON s."id" = t."shop_id"
        WHERE
            s."seller_name" = $1
            AND t."id" = $2
            AND s."enabled" = true
    )
    AND EXISTS (
        SELECT 1
        FROM "product" p
            JOIN "shop" s ON s."id" = p."shop_id"
        WHERE
            s."seller_name" = $1
            AND p."id" = $3
    )
RETURNING *;

-- name: SellerDeleteProductTag :execrows

DELETE FROM "product_tag" tp
WHERE EXISTS (
        SELECT 1
        FROM "product" p
            JOIN "shop" s ON s."id" = p."shop_id"
        WHERE
            s."seller_name" = $1
            AND p."id" = $3
            AND s."enabled" = true
    )
    AND "product_id" = $3
    AND "tag_id" = $2;

-- name: SellerInsertCouponTag :one

INSERT INTO
    "coupon_tag" ("tag_id", "coupon_id")
SELECT $2, $3
WHERE EXISTS (
        SELECT 1
        FROM "tag" t
            JOIN "shop" s ON s."id" = t."shop_id"
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
            AND t."id" = $2
    )
    AND EXISTS (
        SELECT 1
        FROM "coupon" c
            JOIN "shop" s ON s."id" = c."shop_id"
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
            AND c."id" = $3
    )
RETURNING *;

-- name: SellerDeleteCouponTag :execrows

DELETE FROM "coupon_tag" tp
WHERE EXISTS (
        SELECT 1
        FROM "coupon" c
            JOIN "shop" s ON s."id" = c."shop_id"
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
            AND c."id" = $3
    )
    AND "coupon_id" = $3
    AND "tag_id" = $2;