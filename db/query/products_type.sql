-- name: GetProductsType :one
SELECT * FROM products_type
WHERE id = $1 LIMIT 1;

-- name: ListProductsTypes :many
SELECT * FROM products_type
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: CreateProductsType :one
INSERT INTO products_type (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteProductsType :exec
DELETE FROM products_type
WHERE id = $1;

-- name: UpdateProductType :one
UPDATE products_type SET name = $2
WHERE id = $1
RETURNING *;