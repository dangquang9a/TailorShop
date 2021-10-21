-- name: GetOrderItem :one
SELECT * FROM order_items
WHERE order_id = $1 LIMIT 1;

-- name: ListOrderItems :many
SELECT * FROM order_items;

-- name: CreatOrderItem :one
INSERT INTO order_items (
  product_id, quantity
) VALUES (
  $1, $2
  
)
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE order_id = $1;

-- name: DeletOrderItemByProductID :exec
DELETE FROM order_items
WHERE product_id = $1;

