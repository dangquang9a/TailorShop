// Code generated by sqlc. DO NOT EDIT.
// source: customers.sql

package db

import (
	"context"
	"database/sql"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO Customers (
  full_name, address, phone_number
) VALUES (
  $1, $2, $3
)
RETURNING id, full_name, created_at, address, phone_number
`

type CreateCustomerParams struct {
	FullName    string         `json:"fullName"`
	Address     sql.NullString `json:"address"`
	PhoneNumber string         `json:"phoneNumber"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer, arg.FullName, arg.Address, arg.PhoneNumber)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreatedAt,
		&i.Address,
		&i.PhoneNumber,
	)
	return i, err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM Customers
WHERE id = $1
`

func (q *Queries) DeleteCustomer(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, id)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT id, full_name, created_at, address, phone_number FROM customers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCustomer(ctx context.Context, id int32) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreatedAt,
		&i.Address,
		&i.PhoneNumber,
	)
	return i, err
}

const getCustomerByPhone = `-- name: GetCustomerByPhone :one
SELECT id, full_name, created_at, address, phone_number FROM customers
WHERE phone_number = $1 LIMIT 1
`

func (q *Queries) GetCustomerByPhone(ctx context.Context, phoneNumber string) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomerByPhone, phoneNumber)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreatedAt,
		&i.Address,
		&i.PhoneNumber,
	)
	return i, err
}

const listCustomers = `-- name: ListCustomers :many
SELECT id, full_name, created_at, address, phone_number FROM Customers
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCustomersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCustomers(ctx context.Context, arg ListCustomersParams) ([]Customer, error) {
	rows, err := q.db.QueryContext(ctx, listCustomers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Customer{}
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.CreatedAt,
			&i.Address,
			&i.PhoneNumber,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCustomer = `-- name: UpdateCustomer :one
UPDATE Customers SET address = $2, full_name = $3, phone_number = $4
WHERE id = $1
RETURNING id, full_name, created_at, address, phone_number
`

type UpdateCustomerParams struct {
	ID          int32          `json:"id"`
	Address     sql.NullString `json:"address"`
	FullName    string         `json:"fullName"`
	PhoneNumber string         `json:"phoneNumber"`
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, updateCustomer,
		arg.ID,
		arg.Address,
		arg.FullName,
		arg.PhoneNumber,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreatedAt,
		&i.Address,
		&i.PhoneNumber,
	)
	return i, err
}
