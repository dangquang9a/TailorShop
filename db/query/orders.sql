-- name: GetOrder :one
SELECT * FROM Orders
WHERE id = $1 LIMIT 1;

-- name: GetOrderByUserId :many
SELECT * FROM Orders
WHERE user_id = $3
LIMIT $1
OFFSET $2;

-- name: ListOrders :many
SELECT * FROM Orders
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: CreateOrder :one
INSERT INTO Orders (
  user_id, status, prepaid
) VALUES (
  $1, $2, $3
  
)
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM Orders
WHERE id = $1;

-- name: UpdateOder :one
UPDATE Orders SET user_id = $2, status = $3, prepaid = $4
WHERE id = $1
RETURNING *;