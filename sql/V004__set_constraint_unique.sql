ALTER TABLE series
    ADD CONSTRAINT unique_date_created_name UNIQUE (date_created, name);