-- +goose Up
-- SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS books
(
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    published_date DATE NOT NULL,
    image_url TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NOT NULL
);
-- +goose Down
-- SELECT 'down SQL query';
DROP TABLE IF EXISTS books;