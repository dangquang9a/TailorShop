package db

import (
	"context"
	"database/sql"
	"fmt"
	"math"
)

//Priovides all functions to execute db queries and transaction
type Store struct {
	*Queries
	db *sql.DB
}

//Create a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

//execute transactions
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type DeleteCustomerTxParams struct {
	ID int64 `json:"id"`
}

type DeleteCustomerTxResult struct {
	Success bool `json:"success"`
}

func (store *Store) DeleteCustomerTx(ctx context.Context, arg DeleteCustomerTxParams) (DeleteCustomerTxResult, error) {
	var result DeleteCustomerTxResult

	err := store.execTx(ctx, func(q *Queries) error {

		err := q.DeleteMeasureByCustomerID(ctx, (arg.ID))

		if err != nil {
			return err
		}

		orderParam := GetOrderByUserIdParams{
			UserID: int32(arg.ID),
			Offset: 0,
			Limit:  math.MaxInt32,
		}

		listOrders, err := q.GetOrderByUserId(ctx, orderParam)

		if err != nil {
			return err
		}

		for _, order := range listOrders {
			err := q.DeleteOrderItem(ctx, order.ID)

			if err != nil {
				return err
			}

			err = q.DeleteOrder(ctx, order.ID)

			if err != nil {
				return err
			}

		}

		if err != nil {
			return err
		}

		err = q.DeleteCustomer(ctx, int32(arg.ID))

		if err != nil {
			return err
		}

		return nil
	})

	result.Success = true
	return result, err
}
