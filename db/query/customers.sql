-- name: GetCustomer :one
SELECT * FROM customers
WHERE id = $1 LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM Customers
ORDER BY full_name;

-- name: CreateCustomer :one
INSERT INTO Customers (
  full_name, address, phone_number
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM Customers
WHERE id = $1;