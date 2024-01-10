DELETE FROM "coupon" WHERE "shop_id" IS NULL AND "scope" = 'shop';
ALTER TABLE "coupon"
    ADD CONSTRAINT "shop_coupon_constraint" CHECK (
        NOT ("shop_id" IS NULL AND "scope" = 'shop')
    );
