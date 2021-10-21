-- name: GetOrder :one
SELECT * FROM Orders
WHERE id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM Orders
ORDER BY created_at;

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