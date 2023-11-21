CREATE TYPE
    "order_status" AS ENUM (
        'pending',
        'paid',
        'shipped',
        'delivered',
        'cancelled'
    );

CREATE TYPE
    "coupon_type" AS ENUM (
        'percentage',
        'fixed',
        'shipping'
    );

CREATE TYPE "role_type" AS ENUM ( 'admin', 'customer' );

CREATE TABLE
    "cart" (
        "id" SERIAL PRIMARY KEY,
        "user_id" INT NOT NULL,
        "shop_id" INT NOT NULL
    );

CREATE TABLE
    "cart_product" (
        "cart_id" INT NOT NULL,
        "product_id" INT NOT NULL,
        "quantity" INT NOT NULL
    );

CREATE TABLE
    "order_history" (
        "id" SERIAL PRIMARY KEY,
        "user_id" INT NOT NULL,
        "shop_id" INT NOT NULL,
        "shipment" INT NOT NULL,
        "total_price" INT NOT NULL,
        "status" order_status NOT NULL,
        "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );

CREATE TABLE
    "order_detail" (
        "order_id" INT NOT NULL,
        "product_id" INT NOT NULL,
        "product_version" INT NOT NULL,
        "quantity" INT NOT NULL,
        PRIMARY KEY ("order_id", "product_id", "product_version")
    );

CREATE TABLE
    "cart_coupon" (
        "cart_id" INT NOT NULL,
        "coupon_id" INT NOT NULL
    );

CREATE TABLE
    "product" (
        "id" SERIAL PRIMARY KEY,
        "version" INT NOT NULL,
        "shop_id" INT NOT NULL,
        "name" VARCHAR(255) NOT NULL,
        "description" TEXT NOT NULL,
        "price" DECIMAL(10, 2) NOT NULL,
        "image_id" UUID NOT NULL,
        "exp_date" TIMESTAMPTZ NOT NULL,
        "edit_date" TIMESTAMPTZ NOT NULL, -- to limit the edit frequency
        "stock" INT NOT NULL,
        "sales" INT NOT NULL,
        "enabled" BOOLEAN NOT NULL DEFAULT TRUE
    );
CREATE TABLE
    "product_archive" (
        "id" INT NOT NULL,
        "version" INT NOT NULL DEFAULT 1,
        "name" VARCHAR(255) NOT NULL,
        "description" TEXT NOT NULL,
        "price" DECIMAL(10, 2) NOT NULL,
        "image_id" UUID NOT NULL,
        PRIMARY KEY ("id", "version")
    );
CREATE TABLE
    "coupon" (
        "id" SERIAL PRIMARY KEY,
        "type" coupon_type NOT NULL,
        "shop_id" INT NOT NULL,
        "description" TEXT NOT NULL,
        "discount" DECIMAL(5, 2) NOT NULL,
        "start_date" TIMESTAMPTZ NOT NULL,
        "expire_date" TIMESTAMPTZ NOT NULL CHECK ("expire_date" > "start_date")
    );

CREATE TABLE
    "user" (
        "id" SERIAL PRIMARY KEY,
        "username" VARCHAR(255) NOT NULL UNIQUE,
        "password" VARCHAR(255) NOT NULL,
        "name" VARCHAR(255) NOT NULL,
        "email" VARCHAR(255) NOT NULL UNIQUE,
        "address" VARCHAR(255) NOT NULL,
        "image_id" UUID NOT NULL,
        "role" role_type NOT NULL,
        "session_token" VARCHAR(255) NOT NULL,
        "credit_card" JSONB NOT NULL
    );

CREATE TABLE
    "shop" (
        "id" SERIAL PRIMARY KEY,
        "seller_name" VARCHAR(255) NOT NULL,
        "image_id" UUID NOT NULL,
        "name" VARCHAR(255) NOT NULL,
        "enabled" BOOLEAN NOT NULL
    );

CREATE TABLE
    "tag" (
        "id" SERIAL PRIMARY KEY,
        "shop_id" INT NOT NULL,
        "name" VARCHAR(255) NOT NULL
    );

CREATE TABLE
    "product_tag" (
        "tag_id" INT NOT NULL,
        "product_id" INT NOT NULL
    );

CREATE TABLE
    "coupon_tag" (
        "coupon_id" INT NOT NULL,
        "tag_id" INT NOT NULL
    );

ALTER TABLE "cart_product"
ADD
    FOREIGN KEY ("cart_id") REFERENCES "cart" ("id");

ALTER TABLE "cart_coupon"
ADD
    FOREIGN KEY ("cart_id") REFERENCES "cart" ("id");

ALTER TABLE "product"
ADD
    FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "shop"
ADD
    FOREIGN KEY ("seller_name") REFERENCES "user" ("username");

ALTER TABLE "coupon"
ADD
    FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "tag"
ADD
    FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "product_tag"
ADD
    FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "coupon_tag"
ADD
    FOREIGN KEY ("coupon_id") REFERENCES "coupon" ("id");

ALTER TABLE "cart"
ADD
    FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "cart"
ADD
    FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "order_detail"
ADD
    FOREIGN KEY ("order_id") REFERENCES "order_history" ("id");

ALTER TABLE "order_detail"
ADD
    FOREIGN KEY ("product_id", "product_version") REFERENCES "product_archive" ("id", "version");

ALTER TABLE "order_history"
ADD
    FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "order_history"
ADD
    FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");
