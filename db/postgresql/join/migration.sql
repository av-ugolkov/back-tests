CREATE TABLE IF NOT EXISTS
    "join_second" ("id" BIGINT UNIQUE, "data" TEXT NOT NULL);

INSERT INTO
    "join_second" ("id", "data")
SELECT
    g AS "id",
    md5((random() * 100)::text) AS "data"
FROM
    generate_series(1, 100_000) g; -- if the count of the table is more than 50_000_000, buckets can be overloaded


CREATE TABLE IF NOT EXISTS
    "join_first" ("id" UUID, "second_id" BIGINT references "join_second" ("id"));

CREATE INDEX "join_first__second_id" ON "join_first"("second_id");

DROP INDEX IF EXISTS "join_first__second_id";

INSERT INTO
    "join_first" ("id", "second_id")
SELECT
    gen_random_uuid() AS "id",
    s.id AS "second_id"
FROM "join_second" s; -- if the count of the table is more than 50_000_000, buckets can be overloaded



EXPLAIN ANALYZE
SELECT * FROM "join_first" f 
JOIN "join_second" c ON f.second_id = c.id;
