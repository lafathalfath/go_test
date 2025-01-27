-- +goose Up
CREATE TABLE products (
    id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name text    NOT NULL,
    content  text NOT NULL,
    createdAt datetime NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE products;