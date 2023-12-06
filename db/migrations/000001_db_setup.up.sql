CREATE TYPE "order_status" AS ENUM (
    'paid',
    'shipped',
    'delivered',
    'cancelled',
    'finished'
);

CREATE TYPE "coupon_type" AS ENUM ('percentage', 'fixed', 'shipping');

CREATE TYPE "coupon_scope" AS ENUM ('global', 'shop');

CREATE TYPE "role_type" AS ENUM ('admin', 'customer');

CREATE TABLE "cart" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT NOT NULL,
    "shop_id" INT NOT NULL,
    CONSTRAINT unique_shop_user UNIQUE ("shop_id", "user_id")
);

CREATE TABLE "cart_product" (
    "cart_id" INT NOT NULL,
    "product_id" INT NOT NULL,
    "quantity" INT NOT NULL,
    CONSTRAINT unique_cart_product UNIQUE ("cart_id", "product_id")
);

CREATE TABLE "order_history" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT NOT NULL,
    "shop_id" INT NOT NULL,
    "image_id" TEXT,
    -- fot thumbnail
    "shipment" INT NOT NULL,
    "total_price" INT NOT NULL,
    "status" order_status NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE "order_detail" (
    "order_id" INT NOT NULL,
    "product_id" INT NOT NULL,
    "product_version" INT NOT NULL,
    "quantity" INT NOT NULL,
    PRIMARY KEY (
        "order_id",
        "product_id",
        "product_version"
    )
);

CREATE TABLE "cart_coupon" (
    "cart_id" INT NOT NULL,
    "coupon_id" INT NOT NULL,
    CONSTRAINT unique_cart_coupon UNIQUE ("cart_id", "coupon_id")
);

CREATE TABLE "product" (
    "id" SERIAL PRIMARY KEY,
    "version" INT NOT NULL,
    "shop_id" INT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "price" DECIMAL(10, 2) NOT NULL,
    "image_id" TEXT,
    "expire_date" TIMESTAMPTZ NOT NULL,
    "edit_date" TIMESTAMPTZ NOT NULL,
    -- to limit the edit frequency
    "stock" INT NOT NULL,
    "sales" INT NOT NULL DEFAULT 0,
    "enabled" BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE "product_archive" (
    "id" INT NOT NULL,
    "version" INT NOT NULL DEFAULT 1,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "price" DECIMAL(10, 2) NOT NULL,
    "image_id" TEXT,
    PRIMARY KEY ("id", "version")
);

CREATE TABLE "coupon" (
    "id" SERIAL PRIMARY KEY,
    "type" coupon_type NOT NULL,
    "scope" coupon_scope NOT NULL,
    "shop_id" INT,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "discount" DECIMAL(5, 2) NOT NULL,
    "start_date" TIMESTAMPTZ NOT NULL,
    "expire_date" TIMESTAMPTZ NOT NULL CHECK ("expire_date" > "start_date")
);

CREATE TABLE "user" (
    "id" SERIAL PRIMARY KEY,
    "username" TEXT NOT NULL UNIQUE,
    "password" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "email" TEXT NOT NULL UNIQUE,
    "address" TEXT NOT NULL,
    "image_id" TEXT,
    "role" role_type NOT NULL,
    "credit_card" JSONB NOT NULL,
    "refresh_token" TEXT NULL,
    "enabled" BOOLEAN NOT NULL DEFAULT TRUE -- if user deleted, set enabled to false
);

CREATE TABLE "shop" (
    "id" SERIAL PRIMARY KEY,
    "seller_name" TEXT NOT NULL,
    "image_id" TEXT,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "enabled" BOOLEAN NOT NULL
);

CREATE TABLE "tag" (
    "id" SERIAL PRIMARY KEY,
    "shop_id" INT NOT NULL,
    "name" TEXT NOT NULL,
    CONSTRAINT unique_name_shop UNIQUE ("shop_id", "name")
);

CREATE TABLE "product_tag" (
    "tag_id" INT NOT NULL,
    "product_id" INT NOT NULL,
    CONSTRAINT unique_tag_product UNIQUE ("tag_id", "product_id")
);

CREATE TABLE "coupon_tag" (
    "coupon_id" INT NOT NULL,
    "tag_id" INT NOT NULL,
    CONSTRAINT unique_tag_coupon UNIQUE ("tag_id", "coupon_id")
);

ALTER TABLE "shop"
ADD FOREIGN KEY ("seller_name") REFERENCES "user" ("username");

ALTER TABLE "product"
ADD FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "coupon"
ADD FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "tag"
ADD FOREIGN KEY ("shop_id") REFERENCES "shop" ("id") ON DELETE CASCADE;

ALTER TABLE "product_tag"
ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id") ON DELETE CASCADE;

ALTER TABLE "product_tag"
ADD FOREIGN KEY ("tag_id") REFERENCES "tag" ("id") ON DELETE CASCADE;

ALTER TABLE "coupon_tag"
ADD FOREIGN KEY ("coupon_id") REFERENCES "coupon" ("id") ON DELETE CASCADE;

ALTER TABLE "coupon_tag"
ADD FOREIGN KEY ("tag_id") REFERENCES "tag" ("id") ON DELETE CASCADE;

ALTER TABLE "cart"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "cart"
ADD FOREIGN KEY ("shop_id") REFERENCES "shop" ("id") ON DELETE CASCADE;

ALTER TABLE "cart_product"
ADD FOREIGN KEY ("cart_id") REFERENCES "cart" ("id") ON DELETE CASCADE;

ALTER TABLE "cart_product"
ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id") ON DELETE CASCADE;

ALTER TABLE "cart_coupon"
ADD FOREIGN KEY ("cart_id") REFERENCES "cart" ("id") ON DELETE CASCADE;

ALTER TABLE "cart_coupon"
ADD FOREIGN KEY ("coupon_id") REFERENCES "coupon" ("id") ON DELETE CASCADE;

ALTER TABLE "order_detail"
ADD FOREIGN KEY ("order_id") REFERENCES "order_history" ("id");

ALTER TABLE "order_detail"
ADD FOREIGN KEY (
        "product_id",
        "product_version"
    ) REFERENCES "product_archive" ("id", "version");

ALTER TABLE "order_history"
ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "order_history"
ADD FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");
