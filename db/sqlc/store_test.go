package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteCustomerTx(t *testing.T) {
	store := NewStore(testDB)
	// customer1 := createRandomCustomer(t)

	result, err := store.DeleteCustomerTx(context.Background(), DeleteCustomerTxParams{ID: int64(32)})

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, result.Success, true)

	customer2, err := store.GetCustomer(context.Background(), 32)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, customer2)

}
