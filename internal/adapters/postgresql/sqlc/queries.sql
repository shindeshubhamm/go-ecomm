-- name: ListProducts :many
SELECT *
FROM products
ORDER BY name ASC;
-- name: FindProductById :one
SELECT *
FROM products
WHERE id = $1
LIMIT 1;