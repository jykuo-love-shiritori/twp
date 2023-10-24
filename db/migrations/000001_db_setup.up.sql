CREATE TABLE
    "cart" (
        "id" SERIAL PRIMARY KEY,
        "user_id" INT,
        "shop_id" INT
    );

CREATE TABLE
    "cart_product" (
        "cart_id" INT,
        "product_id" INT,
        "quantity" INT
    );

-- Define OrderStatus enum

CREATE TYPE
    order_status AS ENUM (
        'pending',
        'paid',
        'shipped',
        'delivered',
        'cancelled'
    );

CREATE TABLE
    "order_history" (
        "id" SERIAL PRIMARY KEY,
        "user_id" INT,
        "shop_id" INT,
        "shipment" INT,
        "coupon_id" INT [],
        "total_price" INT,
        "status" order_status,
        "created_at" TIMESTAMP
    );

CREATE TABLE
    "order_detail" (
        "order_id" INT,
        "product_price" INT,
        "quantity" INT
    );

CREATE TABLE "cart_coupon" ( "cart_id" INT, "coupon_id" INT );

CREATE TABLE
    "product" (
        "id" SERIAL PRIMARY KEY,
        "shop_id" INT,
        "name" VARCHAR(255),
        "description" TEXT,
        "price" DECIMAL(10, 2),
        "image_name" UUID,
        "due_date" TIMESTAMP,
        "stock" INT,
        "sales" INT,
        "enabled" BOOLEAN
    );

-- Define CouponType enum

CREATE TYPE
    coupon_type AS ENUM (
        'percentage',
        'fixed',
        'shipping'
    );

CREATE TABLE
    "coupon" (
        "id" SERIAL PRIMARY KEY,
        "type" coupon_type,
        "shop_id" INT,
        "information" TEXT,
        "discount" DECIMAL(5, 2),
        "start_date" TIMESTAMP,
        "expire_date" TIMESTAMP
    );

-- Define RoleType enum

CREATE TYPE role_type AS ENUM ( 'admin', 'customer');

CREATE TABLE
    "user" (
        "id" SERIAL PRIMARY KEY,
        "name" VARCHAR(255),
        "email" VARCHAR(255),
        "address" VARCHAR(255),
        "role" role_type,
        "session_token" VARCHAR(255),
        "credit_card" JSON,
        "password" VARCHAR(255)
    );

CREATE TABLE
    "shop" (
        "id" SERIAL PRIMARY KEY,
        "seller_id" INT,
        "name" VARCHAR(255),
        "enabled" BOOLEAN
    );

CREATE TABLE
    "tag" (
        "id" SERIAL PRIMARY KEY,
        "shop_id" INT,
        "name" VARCHAR(255)
    );

CREATE TABLE "product_tag" ( "tag_id" INT, "product_id" INT );

CREATE TABLE "coupon_tag" ( "coupon_id" INT, "tag_id" INT );

ALTER TABLE "cart_product"
ADD
    FOREIGN KEY ("cart_id") REFERENCES "cart" ("id");

ALTER TABLE "cart_coupon"
ADD
    FOREIGN KEY ("cart_id") REFERENCES "cart" ("id");

ALTER TABLE "product"
ADD
    FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "cart_product"
ADD
    FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "cart_coupon"
ADD
    FOREIGN KEY ("coupon_id") REFERENCES "coupon" ("id");

ALTER TABLE "coupon"
ADD
    FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "user"
ADD
    FOREIGN KEY ("id") REFERENCES "shop" ("id");

ALTER TABLE "tag"
ADD
    FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "product_tag"
ADD
    FOREIGN KEY ("tag_id") REFERENCES "tag" ("id");

ALTER TABLE "product_tag"
ADD
    FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "coupon_tag"
ADD
    FOREIGN KEY ("coupon_id") REFERENCES "coupon" ("id");

ALTER TABLE "coupon_tag"
ADD
    FOREIGN KEY ("tag_id") REFERENCES "tag" ("id");

ALTER TABLE "cart"
ADD
    FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "cart"
ADD
    FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "order_detail"
ADD
    FOREIGN KEY ("order_id") REFERENCES "order_history" ("id");

ALTER TABLE "order_history"
ADD
    FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "order_history"
ADD
    FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");