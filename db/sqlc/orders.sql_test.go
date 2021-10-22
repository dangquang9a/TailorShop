package db

import (
	"TailorShop/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func createRandomOrder(t *testing.T) Order {
	arg := CreateOrderParams{
		UserID:  createRandomCustomer(t).ID,
		Status:  sql.NullString{String: util.RandomString(5), Valid: true},
		Prepaid: sql.NullInt64{Int64: int64(util.RandomInt(100000, 600000)), Valid: true},
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, order.UserID, arg.UserID)
	require.Equal(t, order.Status, arg.Status)
	require.Equal(t, order.Prepaid, arg.Prepaid)

	require.NotZero(t, order.ID)

	return order
}

func createRandomOrderByUserID(t *testing.T, id int32) Order {
	arg := CreateOrderParams{
		UserID:  id,
		Status:  sql.NullString{String: util.RandomString(5), Valid: true},
		Prepaid: sql.NullInt64{Int64: int64(util.RandomInt(100000, 600000)), Valid: true},
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, order.UserID, arg.UserID)
	require.Equal(t, order.Status, arg.Status)
	require.Equal(t, order.Prepaid, arg.Prepaid)

	require.NotZero(t, order.ID)

	return order
}

func TestGetOrder(t *testing.T) {

	order1 := createRandomOrder(t)

	order2, err := testQueries.GetOrder(context.Background(), order1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.UserID, order2.UserID)
	require.Equal(t, order1.Status, order2.Status)
	require.Equal(t, order1.Prepaid, order2.Prepaid)
	require.Equal(t, order1.CreatedAt, order2.CreatedAt)

	require.WithinDuration(t, order1.CreatedAt.Time, order2.CreatedAt.Time, time.Second)

}

func TestDeleteOrder(t *testing.T) {

	order1 := createRandomOrder(t)

	err := testQueries.DeleteOrder(context.Background(), order1.ID)

	require.NoError(t, err)

	order2, err := testQueries.GetOrder(context.Background(), order1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, order2)

}

func TestUpdateOrder(t *testing.T) {
	order1 := createRandomOrder(t)
	arg := UpdateOderParams{
		ID:      order1.ID,
		UserID:  createRandomCustomer(t).ID,
		Status:  sql.NullString{String: util.RandomString(5), Valid: true},
		Prepaid: sql.NullInt64{Int64: int64(util.RandomInt(100000, 600000)), Valid: true},
	}

	order2, err := testQueries.UpdateOder(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order2)
	require.NotEmpty(t, order1)

	require.Equal(t, order1.ID, order2.ID)
	require.NotEqual(t, order1.UserID, order2.UserID)
	require.NotEqual(t, order1.Status, order2.Status)
	require.NotEqual(t, order1.Prepaid, order2.Prepaid)
	require.Equal(t, order1.CreatedAt, order2.CreatedAt)

	require.WithinDuration(t, order1.CreatedAt.Time, order2.CreatedAt.Time, time.Second)

}

func TestListOrders(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}

	arg := ListOrdersParams{
		Limit:  5,
		Offset: 5,
	}

	orders, err := testQueries.ListOrders(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, orders, 5)

	for _, value := range orders {
		require.NotEmpty(t, value)
	}
}

func TestGetOrderByUserId(t *testing.T) {

	uid := createRandomCustomer(t).ID

	for i := 0; i < 10; i++ {
		createRandomOrderByUserID(t, uid)
	}

	arg := GetOrderByUserIdParams{
		Limit:  5,
		Offset: 5,
		UserID: uid,
	}

	orders, err := testQueries.GetOrderByUserId(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, orders, 5)

	for _, value := range orders {
		require.NotEmpty(t, value)
	}

}
