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

-- name: UpdateOderItem :one
UPDATE order_items SET order_id = $2, product_id = $3, quantity = $4
WHERE id = $1
RETURNING *;
