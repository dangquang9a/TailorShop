-- name: GetOrderItem :one
SELECT * FROM order_items
WHERE id = $1 LIMIT 1;

-- name: ListOrderItems :many
SELECT * FROM order_items
LIMIT $1
OFFSET $2;

-- name: ListOrderItemsByOrderId :many
SELECT * FROM order_items
WHERE order_id = $3
LIMIT $1
OFFSET $2;

-- name: CreateOrderItem :one
INSERT INTO order_items (
  product_id, quantity, order_id
) VALUES (
  $1, $2, $3
  
)
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE id = $1;

-- name: UpdateOderItem :one
UPDATE order_items SET order_id = $2, product_id = $3, quantity = $4
WHERE id = $1
RETURNING *;
