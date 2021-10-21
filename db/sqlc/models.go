// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Customer struct {
	ID          int32          `json:"id"`
	FullName    string         `json:"fullName"`
	CreatedAt   sql.NullTime   `json:"createdAt"`
	Address     sql.NullString `json:"address"`
	PhoneNumber string         `json:"phoneNumber"`
}

type Measure struct {
	ID         int32  `json:"id"`
	CustomerID int64  `json:"customerID"`
	Name       string `json:"name"`
	Number     string `json:"number"`
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
	ID        int32 `json:"id"`
	OrderID   int32 `json:"orderID"`
	ProductID int32 `json:"productID"`
	Quantity  int32 `json:"quantity"`
}

type Product struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	// must be positive
	Price     int32         `json:"price"`
	TypeID    sql.NullInt32 `json:"typeID"`
	CreatedAt sql.NullTime  `json:"createdAt"`
}

type ProductsType struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
