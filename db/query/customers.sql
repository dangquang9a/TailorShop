-- name: GetCustomer :one
SELECT * FROM customers
WHERE id = $1 LIMIT 1;

-- name: GetCustomerByPhone :one
SELECT * FROM customers
WHERE phone_number = $1 LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM Customers
ORDER BY id
LIMIT $1
OFFSET $2;

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

-- name: UpdateCustomer :one
UPDATE Customers SET address = $2, full_name = $3, phone_number = $4
WHERE id = $1
RETURNING *;