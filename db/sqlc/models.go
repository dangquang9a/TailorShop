// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Customer struct {
	ID          int32          `json:"id"`
	FullName    sql.NullString `json:"fullName"`
	CreatedAt   sql.NullTime   `json:"createdAt"`
	Address     sql.NullString `json:"address"`
	PhoneNumber sql.NullString `json:"phoneNumber"`
}

type Measure struct {
	Code       int32          `json:"code"`
	CustomerID sql.NullInt64  `json:"customerID"`
	Name       sql.NullString `json:"name"`
	Number     sql.NullString `json:"number"`
}

type Order struct {
	ID      int32          `json:"id"`
	UserID  int32          `json:"userID"`
	Status  sql.NullString `json:"status"`
	Prepaid sql.NullInt64  `json:"prepaid"`
	// When order created
	CreatedAt sql.NullTime `json:"createdAt"`
}

type OrderItem struct {
	OrderID   sql.NullInt32 `json:"orderID"`
	ProductID sql.NullInt32 `json:"productID"`
	Quantity  sql.NullInt32 `json:"quantity"`
}

type Product struct {
	ID   int32          `json:"id"`
	Name sql.NullString `json:"name"`
	// must be positive
	Price     int32         `json:"price"`
	TypeID    sql.NullInt32 `json:"typeID"`
	CreatedAt sql.NullTime  `json:"createdAt"`
}

type ProductsType struct {
	ID   int32          `json:"id"`
	Name sql.NullString `json:"name"`
}
