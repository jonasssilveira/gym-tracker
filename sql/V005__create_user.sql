CREATE TABLE "users"
(
    "id"           BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "full_name"    VARCHAR(50),
    "chat_id"      BIGINT,
    "finished"     BOOL DEFAULT FALSE,
    "date_created" TIMESTAMP DEFAULT NOW(),
    "date_updated" TIMESTAMP DEFAULT NOW()
);
