package db

import (
	"TailorShop/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func createRandomProduct(t *testing.T) Product {

	arg := CreateProductParams{
		Name:   util.RandomString(30),
		Price:  int32(util.RandomInt(1, 234)),
		TypeID: sql.NullInt32{Int32: createRandomProductType(t).ID, Valid: true},
	}
	product, err := testQueries.CreateProduct(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, product.Name, arg.Name)
	require.Equal(t, product.Price, arg.Price)
	require.Equal(t, product.TypeID, arg.TypeID)
	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)

	return product

}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	product2, err := testQueries.GetProducts(context.Background(), product1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Price, product2.Price)
	require.Equal(t, product1.TypeID, product2.TypeID)
	require.Equal(t, product1.CreatedAt, product2.CreatedAt)

	require.WithinDuration(t, product1.CreatedAt.Time, product2.CreatedAt.Time, time.Second)

}

func TestUpdateProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	arg := UpdateProductParams{
		ID:     product1.ID,
		Name:   util.RandomString(30),
		Price:  int32(util.RandomInt(1, 234)),
		TypeID: sql.NullInt32{Int32: 3, Valid: true},
	}
	product2, err := testQueries.UpdateProduct(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, product2)
	require.NotEmpty(t, product1)

	require.Equal(t, product1.ID, product2.ID)
	require.NotEqual(t, product1.Name, product2.Name)
	require.NotEqual(t, product1.Price, product2.Price)
	require.NotEqual(t, product1.TypeID, product2.TypeID)
	require.Equal(t, product1.CreatedAt, product2.CreatedAt)

	require.WithinDuration(t, product1.CreatedAt.Time, product2.CreatedAt.Time, time.Second)

}

func TestDeleteProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	err := testQueries.DeleteProduct(context.Background(), product1.ID)

	require.NoError(t, err)

	product2, err := testQueries.GetProducts(context.Background(), product1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}

func TestListProduct(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	arg := ListProductsParams{
		Limit:  5,
		Offset: 5,
	}

	products, err := testQueries.ListProducts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, value := range products {
		require.NotEmpty(t, value)
	}

}
