package model

import "time"

type Order struct {
	ID           int64
	CustomerName string
	OrderItems   []OrderItem
	Status       string
	Created_at   time.Time
}

type OrderItem struct {
	OrderID   int64
	ProductID int64
	Quantity  int64
}

type OrderUpdateData struct {
	CustomerName *string
	OrderItems   *[]OrderItem
	Status       *string
	Created_at   *time.Time
}

type OrderFilter struct {
}
