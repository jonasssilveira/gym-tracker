CREATE TABLE "sets"
(
    "id"           BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "series_id"    INTEGER NOT NULL,
    "weight"       FLOAT,
    "reps"         INTEGER,
    "time"         INTEGER,
    "rest_time"    INTEGER,
    "date_created" TIMESTAMP DEFAULT NOW(),
    "date_updated" TIMESTAMP DEFAULT NOW()
);