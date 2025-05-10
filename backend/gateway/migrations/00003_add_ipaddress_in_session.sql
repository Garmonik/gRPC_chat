-- +goose Up
-- SQL in this section is executed when the migration is applied

ALTER TABLE session
    ADD COLUMN ip_address VARCHAR(255);

-- Обновляем существующие записи (опционально)
UPDATE session SET ip_address = '' WHERE ip_address IS NULL;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back

ALTER TABLE session
DROP COLUMN ip_address;