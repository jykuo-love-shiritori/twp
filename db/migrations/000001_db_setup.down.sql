-- Drop foreign key constraints

ALTER TABLE
    "order_history" DROP CONSTRAINT "order_history_user_id_fkey";

ALTER TABLE
    "order_history" DROP CONSTRAINT "order_history_shop_id_fkey";

ALTER TABLE
    "order_detail" DROP CONSTRAINT "order_detail_order_id_fkey";

ALTER TABLE "cart" DROP CONSTRAINT "cart_user_id_fkey";

ALTER TABLE "cart" DROP CONSTRAINT "cart_shop_id_fkey";

ALTER TABLE "coupon_tag" DROP CONSTRAINT "coupon_tag_coupon_id_fkey";

ALTER TABLE "coupon_tag" DROP CONSTRAINT "coupon_tag_tag_id_fkey";

ALTER TABLE "product_tag" DROP CONSTRAINT "product_tag_tag_id_fkey";

ALTER TABLE
    "product_tag" DROP CONSTRAINT "product_tag_product_id_fkey";

ALTER TABLE "tag" DROP CONSTRAINT "tag_shop_id_fkey";

ALTER TABLE "coupon" DROP CONSTRAINT "coupon_shop_id_fkey";

ALTER TABLE "cart_coupon" DROP CONSTRAINT "cart_coupon_cart_id_fkey";

ALTER TABLE
    "cart_product" DROP CONSTRAINT "cart_product_cart_id_fkey";

ALTER TABLE
    "cart_product" DROP CONSTRAINT "cart_product_product_id_fkey";

ALTER TABLE "product" DROP CONSTRAINT "product_shop_id_fkey";

ALTER TABLE "user" DROP CONSTRAINT "user_id_fkey";

-- Drop tables

DROP TABLE IF EXISTS "order_detail";

DROP TABLE IF EXISTS "order_history";

DROP TABLE IF EXISTS "cart_coupon";

DROP TABLE IF EXISTS "cart_product";

DROP TABLE IF EXISTS "coupon_tag";

DROP TABLE IF EXISTS "product_tag";

DROP TABLE IF EXISTS "coupon";

DROP TABLE IF EXISTS "product";

DROP TABLE IF EXISTS "tag";

DROP TABLE IF EXISTS "cart";

DROP TABLE IF EXISTS "user";

DROP TABLE IF EXISTS "shop";

-- Drop types

DROP TYPE IF EXISTS "order_status";

DROP TYPE IF EXISTS "coupon_type";

DROP TYPE IF EXISTS "role_type";