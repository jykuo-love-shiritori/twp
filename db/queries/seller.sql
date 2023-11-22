-- name: GetSellerInfo :one

SELECT s.*
FROM "user" u
    JOIN "shop" s ON u.username = s.seller_name
WHERE u.id = $1;

-- name: UpdateSellerInfo :exec

UPDATE "shop"
SET
    "image_id" = COALESCE($2, "image_id"),
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "enabled" = COALESCE($5, "enabled")
WHERE "seller_name" IN (
        SELECT "username"
        FROM "user" u
        WHERE u.id = $1
    );

-- name: SearchTag :many

SELECT t."id", t."name"
FROM "tag" t
    LEFT JOIN "shop" s ON "shop_id" = s.id
    LEFT JOIN "user" u ON s.seller_name = u.username
WHERE u.id = $1 AND t."name" ~* $2
ORDER BY LENGTH(t."name")
LIMIT 10;

-- name: HaveTagName :one

SELECT
    CASE
        WHEN EXISTS (
            SELECT *
            FROM "tag" t
                LEFT JOIN "shop" s ON "shop_id" = s.id
            WHERE
                s."seller_name" = $1
                AND t."name" = $2
        ) THEN true
        ELSE false
    END;

-- name: InsertTag :one

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
    ) ON CONFLICT ("shop_id", "name")
DO NOTHING
RETURNING "id", "name";

-- name: SellerGetCoupon :many

SELECT
    c."id",
    c."type",
    c."shop_id",
    c."name",
    c."discount",
    c."expire_date"
FROM "coupon" c
    JOIN "shop" s ON c."shop_id" = s.id
WHERE s.seller_name = $1
ORDER BY "start_date" DESC
LIMIT $2
OFFSET $3;

-- name: SellerGetCouponDetail :many

SELECT c.*
FROM "coupon" c
    JOIN "shop" s ON c."shop_id" = s.id
WHERE
    s."seller_name" = $1
    AND c."id" = $2;

-- name: SellerInsertCoupon :one

INSERT INTO
    "coupon" (
        "type",
        "shop_id",
        "name",
        "description",
        "discount",
        "start_date",
        "expire_date"
    )
VALUES (
        $2, (
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
RETURNING *;

-- name: UpdateCouponInfo :execrows

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
    );

-- name: DeleteCoupon :execrows

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
    "shop_id",
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
            AND s."enabled" = true
    )
ORDER BY "created_at" DESC
LIMIT $2
OFFSET $3;

-- name: SellerGetOrderDetail :many

SELECT
    product_archive.*,
    order_history.*,
    order_detail.*
FROM "order_detail"
    LEFT JOIN product_archive ON order_detail.product_id = product_archive.id AND order_detail.product_version = product_archive.version
    LEFT JOIN order_history ON order_history.id = order_detail.order_id
    LEFT JOIN shop ON order_history.shop_id = shop.id
WHERE
    shop.seller_name = $1
    OR order_detail.order_id = $2;

-- SellerGetReport :many

-- SellerGetReportDetail :many

-- name: SellerInsertProduct :one

INSERT INTO
    "product"(
        "version",
        "shop_id",
        "name",
        "description",
        "price",
        "image_id",
        "exp_date"
    )
VALUES (
        0, (
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
        $6
    )
RETURNING "id";

-- name: UpdateProductInfo :execrows

UPDATE "product" p
SET
    "name" = COALESCE($3, "name"),
    "description" = COALESCE($4, "description"),
    "price" = COALESCE($5, "price"),
    "image_id" = COALESCE($6, "image_id"),
    "exp_date" = COALESCE($7, "exp_date"),
    "edit_date" = NOW(),
    "version" = "version" + 1
WHERE "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
    AND p."id" = $2;

-- name: DeleteProduct :execrows

UPDATE "product" p
SET
    "enabled" = false,
    "edit_date" = NOW()
WHERE "shop_id" = (
        SELECT s."id"
        FROM "shop" s
        WHERE
            s."seller_name" = $1
            AND s."enabled" = true
    )
    AND p."id" = $2;
