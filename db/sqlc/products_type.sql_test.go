package db

import (
	"TailorShop/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateProductType(t *testing.T) {
	createRandomProductType(t)
}

func createRandomProductType(t *testing.T) ProductsType {
	randomName := util.RandomString(10)
	products_type, err := testQueries.CreateProductsType(context.Background(), randomName)

	require.NoError(t, err)
	require.NotEmpty(t, products_type)

	require.Equal(t, products_type.Name, randomName)
	require.NotZero(t, products_type.ID)

	return products_type
}
func TestGetProductTypes(t *testing.T) {
	productTypes1 := createRandomProductType(t)
	productTypes2, err := testQueries.GetProductsType(context.Background(), productTypes1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, productTypes2)

	require.Equal(t, productTypes1.ID, productTypes2.ID)
	require.Equal(t, productTypes1.Name, productTypes2.Name)

}
func TestUpdateProductTypes(t *testing.T) {
	productTypes1 := createRandomProductType(t)

	arg := UpdateProductTypeParams{
		ID:   productTypes1.ID,
		Name: util.RandomString(15),
	}

	productTypes2, err := testQueries.UpdateProductType(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, productTypes2)
	require.NotEmpty(t, productTypes1)

	require.Equal(t, productTypes1.ID, productTypes2.ID)
	require.NotEqual(t, productTypes1.Name, productTypes2.Name)

}

func TestDeleteProductTypes(t *testing.T) {
	productTypes1 := createRandomProductType(t)
	err := testQueries.DeleteProductsType(context.Background(), productTypes1.ID)

	require.NoError(t, err)

	productTypes2, err := testQueries.GetProductsType(context.Background(), productTypes1.ID)

	require.Empty(t, productTypes2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}

func TestListProductTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProductType(t)
	}

	arg := ListProductsTypesParams{
		Limit:  5,
		Offset: 5,
	}

	productsType, err := testQueries.ListProductsTypes(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, productsType, 5)

	for _, value := range productsType {
		require.NotEmpty(t, value)
	}
}
