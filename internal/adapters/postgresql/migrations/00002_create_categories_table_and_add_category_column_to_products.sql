-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE products
ADD COLUMN IF NOT EXISTS category_id UUID REFERENCES categories(id);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE products DROP COLUMN IF EXISTS category_id;
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd