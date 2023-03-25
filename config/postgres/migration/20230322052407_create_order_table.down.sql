-- Kalo yang migration down diketiknyo urutannyo cak ini (tebalik dari urutan up)

DROP TABLE IF EXISTS orders;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
        DROP TYPE order_status;
    END IF;
END $$;

