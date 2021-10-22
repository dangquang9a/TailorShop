package db

import (
	"TailorShop/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOrderItem(t *testing.T) {
	createRandomOrderItem(t)
}

func createRandomOrderItem(t *testing.T) OrderItem {

	arg := CreateOrderItemParams{
		ProductID: createRandomProduct(t).ID,
		Quantity:  int32(util.RandomInt(0, 10)),
		OrderID:   createRandomOrder(t).ID,
	}

	order_item, err := testQueries.CreateOrderItem(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order_item)

	require.Equal(t, order_item.ProductID, arg.ProductID)
	require.Equal(t, order_item.Quantity, arg.Quantity)
	require.Equal(t, order_item.OrderID, arg.OrderID)

	require.NotZero(t, order_item.ID)

	return order_item

}

func createRandomOrderItemByOrder(t *testing.T, o int32) OrderItem {

	arg := CreateOrderItemParams{
		ProductID: createRandomProduct(t).ID,
		Quantity:  int32(util.RandomInt(0, 10)),
		OrderID:   o,
	}

	order_item, err := testQueries.CreateOrderItem(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order_item)

	require.Equal(t, order_item.ProductID, arg.ProductID)
	require.Equal(t, order_item.Quantity, arg.Quantity)
	require.Equal(t, order_item.OrderID, arg.OrderID)

	require.NotZero(t, order_item.ID)

	return order_item

}

func TestGetOrderItem(t *testing.T) {

	order_item1 := createRandomOrderItem(t)

	order_item2, err := testQueries.GetOrderItem(context.Background(), order_item1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, order_item2)

	require.Equal(t, order_item1.ID, order_item2.ID)
	require.Equal(t, order_item1.ProductID, order_item2.ProductID)
	require.Equal(t, order_item1.Quantity, order_item2.Quantity)
	require.Equal(t, order_item1.OrderID, order_item2.OrderID)

}

func TestDeleteOrderItem(t *testing.T) {

	order_item1 := createRandomOrderItem(t)

	err := testQueries.DeleteOrderItem(context.Background(), order_item1.ID)

	require.NoError(t, err)

	order_item2, err := testQueries.GetOrderItem(context.Background(), order_item1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, order_item2)

}

func TestListOrderItem(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrderItem(t)
	}

	arg := ListOrderItemsParams{
		Limit:  5,
		Offset: 5,
	}

	order_item, err := testQueries.ListOrderItems(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, order_item, 5)

	for _, value := range order_item {
		require.NotEmpty(t, value)
	}
}

func TestListOrderItemByOrder(t *testing.T) {

	orderID := createRandomOrder(t).ID
	for i := 0; i < 10; i++ {
		createRandomOrderItemByOrder(t, orderID)
	}

	arg := ListOrderItemsByOrderIdParams{
		Limit:   5,
		Offset:  5,
		OrderID: orderID,
	}

	order_item, err := testQueries.ListOrderItemsByOrderId(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, order_item, 5)

	for _, value := range order_item {
		require.NotEmpty(t, value)
	}
}

func TestUpdateOrderItem(t *testing.T) {
	order_item1 := createRandomOrderItem(t)
	arg := UpdateOderItemParams{
		ID:        order_item1.ID,
		ProductID: createRandomProduct(t).ID,
		Quantity:  int32(util.RandomInt(0, 10)),
		OrderID:   createRandomOrder(t).ID,
	}

	order2, err := testQueries.UpdateOderItem(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order2)
	require.NotEmpty(t, order_item1)

	require.Equal(t, order_item1.ID, order2.ID)
	require.NotEqual(t, order_item1.ProductID, order2.ProductID)
	require.NotEqual(t, order_item1.Quantity, order2.Quantity)
	require.NotEqual(t, order_item1.OrderID, order2.OrderID)

}
