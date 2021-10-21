// Code generated by sqlc. DO NOT EDIT.
// source: order_items.sql

package db

import (
	"context"
)

const creatOrderItem = `-- name: CreatOrderItem :one
INSERT INTO order_items (
  product_id, quantity
) VALUES (
  $1, $2
  
)
RETURNING id, order_id, product_id, quantity
`

type CreatOrderItemParams struct {
	ProductID int32 `json:"productID"`
	Quantity  int32 `json:"quantity"`
}

func (q *Queries) CreatOrderItem(ctx context.Context, arg CreatOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, creatOrderItem, arg.ProductID, arg.Quantity)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.Quantity,
	)
	return i, err
}

const deletOrderItemByProductID = `-- name: DeletOrderItemByProductID :exec
DELETE FROM order_items
WHERE product_id = $1
`

func (q *Queries) DeletOrderItemByProductID(ctx context.Context, productID int32) error {
	_, err := q.db.ExecContext(ctx, deletOrderItemByProductID, productID)
	return err
}

const deleteOrderItem = `-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE order_id = $1
`

func (q *Queries) DeleteOrderItem(ctx context.Context, orderID int32) error {
	_, err := q.db.ExecContext(ctx, deleteOrderItem, orderID)
	return err
}

const getOrderItem = `-- name: GetOrderItem :one
SELECT id, order_id, product_id, quantity FROM order_items
WHERE order_id = $1 LIMIT 1
`

func (q *Queries) GetOrderItem(ctx context.Context, orderID int32) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, getOrderItem, orderID)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.Quantity,
	)
	return i, err
}

const listOrderItems = `-- name: ListOrderItems :many
SELECT id, order_id, product_id, quantity FROM order_items
`

func (q *Queries) ListOrderItems(ctx context.Context) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, listOrderItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrderItem
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.Quantity,
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
