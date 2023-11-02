-- Drop foreign keys

ALTER TABLE "order_history"
DROP CONSTRAINT "order_history_user_id_fkey";

ALTER TABLE "order_history"
DROP CONSTRAINT "order_history_shop_id_fkey";

ALTER TABLE "order_detail"
DROP CONSTRAINT "order_detail_order_id_fkey";

ALTER TABLE "cart"
DROP CONSTRAINT "cart_user_id_fkey";

ALTER TABLE "cart"
DROP CONSTRAINT "cart_shop_id_fkey";

ALTER TABLE "product_tag"
DROP CONSTRAINT "product_tag_product_id_fkey";

ALTER TABLE "coupon_tag"
DROP CONSTRAINT "coupon_tag_coupon_id_fkey";

ALTER TABLE "shop"
DROP CONSTRAINT "shop_seller_name_fkey";

ALTER TABLE "coupon"
DROP CONSTRAINT "coupon_shop_id_fkey";

ALTER TABLE "product"
DROP CONSTRAINT "product_shop_id_fkey";

ALTER TABLE "cart_coupon"
DROP CONSTRAINT "cart_coupon_cart_id_fkey";

ALTER TABLE "cart_product"
DROP CONSTRAINT "cart_product_cart_id_fkey";

-- Drop tables

DROP TABLE "cart_coupon";

DROP TABLE "product_tag";

DROP TABLE "coupon_tag";

DROP TABLE "tag";

DROP TABLE "shop";

DROP TABLE "user";

DROP TABLE "coupon";

DROP TABLE "product";

DROP TABLE "order_detail";

DROP TABLE "order_history";

DROP TABLE "cart_product";

DROP TABLE "cart";

-- Drop ENUM types

DROP TYPE "role_type";

DROP TYPE "coupon_type";

DROP TYPE "order_status";
