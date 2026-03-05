ALTER TABLE "series"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");