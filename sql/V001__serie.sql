CREATE TABLE "series"
(
    "id"           BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name"         VARCHAR(50),
    "total_time"   INTEGER,
    "finished"     BOOL DEFAULT false,
    "date_created" TIMESTAMP DEFAULT NOW(),
    "date_updated" TIMESTAMP DEFAULT NOW()
);
