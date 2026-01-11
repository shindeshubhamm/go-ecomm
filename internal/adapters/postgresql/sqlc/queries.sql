-- name: ListProducts :many
SELECT *
FROM products
ORDER BY name ASC;
-- name: FindProductById :one
SELECT *
FROM products
WHERE id = $1
LIMIT 1;
-- name: ListOrders :many
SELECT *
FROM orders;
-- name: FindOrderById :one
SELECT *
FROM orders
WHERE id = $1
LIMIT 1;