CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX product_name_desc_trgm_idx ON "product" USING GIN(("name" || ' ' || "description") gin_trgm_ops);

CREATE INDEX shop_name_desc_trgm_idx ON "shop" USING GIN(("name" || ' ' || "description") gin_trgm_ops);

-- due to offset and limit are sql reserved words, we have to use offset_val and limit_val
-- for some reason, returns table does not work with sqlc although they say they already fix that
CREATE OR REPLACE FUNCTION search_products(IN search_query TEXT, IN shop_id INT, IN min_price NUMERIC, IN max_price NUMERIC, IN min_stock INT, IN max_stock INT, IN has_coupon BOOLEAN, IN sort_by TEXT, IN order_direction TEXT, IN offset_val INT, IN limit_val INT)
    RETURNS TABLE(
        id INT,
        name TEXT,
        price NUMERIC,
        image_url TEXT,
        sales INT
    )
    AS $$
DECLARE
    query TEXT;
BEGIN
    -- Start of query
    query := 'SELECT p.id, p.name, p.price, p.image_id, p.sales FROM product p WHERE (p.name || '' '' || p.description) %> $1 AND p.enabled = TRUE';
    -- Shop filter
    IF shop_id IS NOT NULL THEN
        query := query || ' AND p.shop_id = $2';
    END IF;
    -- Price range filter
    IF min_price IS NOT NULL AND max_price IS NOT NULL THEN
        query := query || ' AND p.price BETWEEN $3 AND $4';
    END IF;
    -- Stock range filter
    IF min_stock IS NOT NULL AND max_stock IS NOT NULL THEN
        query := query || ' AND p.stock BETWEEN $5 AND $6';
    END IF;
    -- Coupon availability filter
    IF has_coupon IS NOT NULL THEN
        query := query || ' AND EXISTS (SELECT 1 FROM coupon WHERE coupon.shop_id = p.shop_id AND NOW() BETWEEN coupon.start_date AND coupon.expire_date)';
    END IF;
    -- Add dynamic ORDER BY clause
    IF sort_by IS NOT NULL THEN
        -- Handle different order by conditions
        CASE sort_by
        WHEN 'relevancy' THEN
            query := query || ' ORDER BY similarity((p.name || '' '' || p.description), $1)';
        WHEN 'sales' THEN
            query := query || ' ORDER BY p.sales';
        WHEN 'price' THEN
            query := query || ' ORDER BY p.price';
        WHEN 'stock' THEN
            query := query || ' ORDER BY p.stock';
        ELSE
            RAISE EXCEPTION 'Invalid order_by parameter';
        END CASE;
        -- Sorting direction
        IF order_direction IS NOT NULL THEN
                CASE WHEN order_direction = 'asc' THEN
                    query := query || ' ASC';
                WHEN order_direction = 'desc' THEN
                    query := query || ' DESC';
                ELSE
                    RAISE EXCEPTION 'Invalid order_direction parameter';
                END CASE;
                END IF;
                -- Add offset and limit
                IF offset_val IS NOT NULL THEN
                        query := query || ' OFFSET $7';
                    END IF;
            IF limit_val IS NOT NULL THEN
                query := query || ' LIMIT $8';
            END IF;
        END IF;
        -- Execute the query
        RETURN QUERY EXECUTE query
        USING search_query, shop_id, min_price, max_price, min_stock, max_stock, offset_val, limit_val;
END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION search_shop(search_query TEXT, offset_val INT, limit_val INT)
    RETURNS TABLE(
        name TEXT,
        seller_name TEXT,
        image_url TEXT
    )
    AS $$
DECLARE
    query TEXT;
BEGIN
    query := 'SELECT s.name, s.seller_name, s.image_id  FROM shop s WHERE (name || '' '' || description) %> ' || quote_literal(search_query) || ' AND s.enabled = TRUE';
    IF offset_val IS NOT NULL AND limit_val IS NOT NULL THEN
        query := query || ' OFFSET ' || offset_val || ' LIMIT ' || limit_val;
    END IF;
    RETURN QUERY EXECUTE query;
END;
$$
LANGUAGE plpgsql;
