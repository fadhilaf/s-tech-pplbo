DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
        RAISE NOTICE 'order_status type already exists';
    ELSE
        DROP TYPE order_status;
    END IF;
END $$;

DROP TABLE IF EXISTS orders;
