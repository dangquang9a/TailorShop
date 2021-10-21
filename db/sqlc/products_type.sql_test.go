package db

import (
	"TailorShop/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
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
