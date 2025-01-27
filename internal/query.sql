-- name: GetOneProduct :one
SELECT * FROM products
WHERE id = ? LIMIT 1;

-- name: ListAllProducts :many
SELECT * FROM products
ORDER BY name;

-- name: CreateProduct :execresult
INSERT INTO products (
    name, content
) VALUES (
    ?, ?
);

-- name: UpdateProduct :execresult
UPDATE products 
SET name = ?, content = ?
WHERE id = ?;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = ?;