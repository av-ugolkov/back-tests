CREATE TABLE IF NOT EXISTS
    "hash_index" ("id" UUID primary key, "hash" text not null);

CREATE INDEX "hash_index__hash" ON "hash_index" USING HASH ("hash");

DROP INDEX "hash_index__hash";

INSERT INTO
    "hash_index" ("id", "hash")
SELECT
    gen_random_uuid() AS "id",
    md5((random() * 100)::text) AS "hash"
FROM
    generate_series(1, 1_000_000); -- if the count of the table is more than 50_000_000, buckets can be overloaded

EXPLAIN ANALYZE SELECT "id", "hash" FROM "hash_index"
WHERE "hash"='18c4d76a5e2e6c1d33bc70cf4e0b094c';

SELECT indexrelid::regclass, idx_scan, idx_tup_read, idx_tup_fetch 
FROM pg_stat_user_indexes 
WHERE indexrelid = 'hash_index__hash'::regclass;