package db

import (
	"TailorShop/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateMeasure(t *testing.T) {
	CreateRandomMeasure(t)
}

func CreateRandomMeasure(t *testing.T) Measure {

	arg := CreateMeasureParams{
		CustomerID: int64(createRandomCustomer(t).ID),
		Name:       util.RandomString(10),
		Number:     util.RandomString(10),
	}

	measure, err := testQueries.CreateMeasure(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, measure)

	require.Equal(t, measure.Name, arg.Name)
	require.Equal(t, measure.Number, arg.Number)
	require.Equal(t, measure.CustomerID, arg.CustomerID)

	require.NotZero(t, measure.ID)
	return measure
}

func CreateRandomMeasureByCustomerID(t *testing.T, customerID int64) Measure {

	arg := CreateMeasureParams{
		CustomerID: customerID,
		Name:       util.RandomString(10),
		Number:     util.RandomString(10),
	}

	measure, err := testQueries.CreateMeasure(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, measure)

	require.Equal(t, measure.Name, arg.Name)
	require.Equal(t, measure.Number, arg.Number)
	require.Equal(t, measure.CustomerID, arg.CustomerID)

	require.NotZero(t, measure.ID)
	return measure
}

func TestGetMeasure(t *testing.T) {

	measure1 := CreateRandomMeasure(t)

	measure2, err := testQueries.GetMeasureByID(context.Background(), measure1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, measure2)

	require.Equal(t, measure1.ID, measure2.ID)
	require.Equal(t, measure1.CustomerID, measure2.CustomerID)
	require.Equal(t, measure1.Name, measure2.Name)
	require.Equal(t, measure1.Number, measure2.Number)

}

func TestGetMeasureByCustomer(t *testing.T) {

	customer := createRandomCustomer(t)

	for i := 0; i < 10; i++ {
		CreateRandomMeasureByCustomerID(t, int64(customer.ID))
	}

	argForQuery := GetMeasureByCustomerIDParams{
		Limit:      5,
		Offset:     5,
		CustomerID: int64(customer.ID),
	}

	measures, err := testQueries.GetMeasureByCustomerID(context.Background(), argForQuery)

	require.NoError(t, err)
	require.Len(t, measures, 5)

	for _, value := range measures {
		require.NotEmpty(t, value)
	}

}

func TestDeleteMeasureByID(t *testing.T) {

	measure1 := CreateRandomMeasure(t)

	err := testQueries.DeleteMeasureByID(context.Background(), measure1.ID)

	require.NoError(t, err)

	measure2, err := testQueries.GetMeasureByID(context.Background(), measure1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, measure2)

}

func TestDeleteMeasureByCustomerID(t *testing.T) {

	measure1 := CreateRandomMeasure(t)

	err := testQueries.DeleteMeasureByCustomerID(context.Background(), measure1.CustomerID)

	require.NoError(t, err)

	arg := GetMeasureByCustomerIDParams{
		Limit:      1,
		Offset:     0,
		CustomerID: measure1.CustomerID,
	}

	measure2, _ := testQueries.GetMeasureByCustomerID(context.Background(), arg)

	require.Empty(t, measure2)
	require.Len(t, measure2, 0)

}

func TestGetListMeasures(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomMeasure(t)
	}

	arg := ListmeasuresParams{
		Limit:  5,
		Offset: 5,
	}

	Measures, err := testQueries.Listmeasures(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, Measures, 5)

	for _, value := range Measures {
		require.NotEmpty(t, value)
	}
}

func TestUpdateMeasures(t *testing.T) {
	measure1 := CreateRandomMeasure(t)

	arg := UpdateMeasureNumberParams{
		ID:     measure1.ID,
		Number: util.RandomString(3),
	}
	measure2, err := testQueries.UpdateMeasureNumber(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, measure2)
	require.NotEmpty(t, measure1)

	require.Equal(t, measure1.ID, measure2.ID)
	require.Equal(t, measure1.Name, measure2.Name)
	require.Equal(t, measure1.CustomerID, measure2.CustomerID)
	require.NotEqual(t, measure1.Number, measure2.Number)

}
