UPDATE "coupon"
SET "discount" = 999.99
    WHERE "discount" > 999.99;

ALTER TABLE "coupon"
    ALTER COLUMN "discount" TYPE DECIMAL(5, 2);
