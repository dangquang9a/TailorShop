-- name: GetProducts :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM Products
ORDER BY name;

-- name: CreateProduct :one
INSERT INTO products (
  name, price, type_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM Products
WHERE id = $1;