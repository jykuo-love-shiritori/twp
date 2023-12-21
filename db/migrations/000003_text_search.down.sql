DROP FUNCTION IF EXISTS search_products;

DROP FUNCTION IF EXISTS search_shop;

DROP INDEX IF EXISTS product_name_desc_trgm_idx;

DROP INDEX IF EXISTS shop_name_desc_trgm_idx;

DROP EXTENSION IF EXISTS pg_trgm;
