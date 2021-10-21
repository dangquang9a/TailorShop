package db

import (
	"TailorShop/util"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func createRandomCustomer(t *testing.T) Customer {
	arg := CreateCustomerParams{
		FullName:    util.RandomString(30),
		Address:     sql.NullString{util.RandomString(30), true},
		PhoneNumber: util.RandomString(30),
	}
	customer, err := testQueries.CreateCustomer(context.Background(), arg)

	// fmt.Println(customer)
	// fmt.Println(err)
	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.FullName, customer.FullName)
	require.Equal(t, arg.Address, customer.Address)
	require.Equal(t, arg.PhoneNumber, customer.PhoneNumber)

	require.NotZero(t, customer.ID)
	require.NotZero(t, customer.CreatedAt)

	return customer
}

func TestGetCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	customer2, err := testQueries.GetCustomer(context.Background(), customer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.ID, customer2.ID)
	require.Equal(t, customer1.Address, customer2.Address)
	require.Equal(t, customer1.FullName, customer2.FullName)
	require.Equal(t, customer1.PhoneNumber, customer2.PhoneNumber)
	require.Equal(t, customer1.CreatedAt, customer2.CreatedAt)

	require.WithinDuration(t, customer1.CreatedAt.Time, customer2.CreatedAt.Time, time.Second)
}

func TestUpdateCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	arg := UpdateCustomerParams{
		ID:          customer1.ID,
		FullName:    util.RandomString(30),
		Address:     sql.NullString{util.RandomString(30), true},
		PhoneNumber: util.RandomString(10),
	}

	customer2, err := testQueries.UpdateCustomer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, customer2)
	require.NotEmpty(t, customer1)

	require.Equal(t, customer1.ID, customer2.ID)
	require.NotEqual(t, customer1.Address, customer2.Address)
	require.NotEqual(t, customer1.FullName, customer2.FullName)
	require.NotEqual(t, customer1.PhoneNumber, customer2.PhoneNumber)
	require.Equal(t, customer1.CreatedAt, customer2.CreatedAt)

	require.WithinDuration(t, customer1.CreatedAt.Time, customer2.CreatedAt.Time, time.Second)

}

func TestDeleteCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	err := testQueries.DeleteCustomer(context.Background(), customer1.ID)

	require.NoError(t, err)

	customer2, err := testQueries.GetCustomer(context.Background(), customer1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, customer2)

	err2 := testQueries.DeleteCustomer(context.Background(), 4)
	fmt.Println(err2)
}

func TestListCustomer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCustomer(t)
	}

	arg := ListCustomersParams{
		Limit:  5,
		Offset: 5,
	}

	customers, err := testQueries.ListCustomers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, customers, 5)

	for _, customers := range customers {
		require.NotEmpty(t, customers)
	}
}
