-- name: GetProducts :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM Products
ORDER BY name
LIMIT $1
OFFSET $2;

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

-- name: UpdateProduct :one
UPDATE Products SET name = $2, price = $3, type_id = $4
WHERE id = $1
RETURNING *;