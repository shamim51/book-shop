-- +goose Up
ALTER TABLE books
ALTER COLUMN deleted_at DROP NOT NULL;

-- +goose Down
ALTER TABLE books
ALTER COLUMN deleted_at SET NOT NULL;
