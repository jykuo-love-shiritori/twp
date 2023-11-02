-- Drop foreign keys

ALTER TABLE "order_history"
DROP CONSTRAINT IF EXISTS "order_history_user_id_fkey";

ALTER TABLE "order_history"
DROP CONSTRAINT IF EXISTS "order_history_shop_id_fkey";

ALTER TABLE "order_detail"
DROP CONSTRAINT IF EXISTS "order_detail_order_id_fkey";

ALTER TABLE "cart"
DROP CONSTRAINT IF EXISTS "cart_user_id_fkey";

ALTER TABLE "cart"
DROP CONSTRAINT IF EXISTS "cart_shop_id_fkey";

ALTER TABLE "product_tag"
DROP CONSTRAINT IF EXISTS "product_tag_product_id_fkey";

ALTER TABLE "coupon_tag"
DROP CONSTRAINT IF EXISTS "coupon_tag_coupon_id_fkey";

ALTER TABLE "shop"
DROP CONSTRAINT IF EXISTS "shop_seller_name_fkey";

ALTER TABLE "coupon"
DROP CONSTRAINT IF EXISTS "coupon_shop_id_fkey";

ALTER TABLE "product"
DROP CONSTRAINT IF EXISTS "product_shop_id_fkey";

ALTER TABLE "cart_coupon"
DROP CONSTRAINT IF EXISTS "cart_coupon_cart_id_fkey";

ALTER TABLE "cart_product"
DROP CONSTRAINT IF EXISTS "cart_product_cart_id_fkey";

-- Drop tables

DROP TABLE IF EXISTS "cart_coupon";

DROP TABLE IF EXISTS "product_tag";

DROP TABLE IF EXISTS "coupon_tag";

DROP TABLE IF EXISTS "tag";

DROP TABLE IF EXISTS "shop";

DROP TABLE IF EXISTS "user";

DROP TABLE IF EXISTS "coupon";

DROP TABLE IF EXISTS "product";

DROP TABLE IF EXISTS "order_detail";

DROP TABLE IF EXISTS "order_history";

DROP TABLE IF EXISTS "cart_product";

DROP TABLE IF EXISTS "cart";

-- Drop ENUM types

DROP TYPE IF EXISTS "role_type";

DROP TYPE IF EXISTS "coupon_type";

DROP TYPE IF EXISTS "order_status";
