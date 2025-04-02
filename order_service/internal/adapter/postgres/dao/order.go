package dao

type Order struct {
	ID           int64
	CustomerName string
	Items        []OrderItem
}

type OrderItem struct {
	ProductID int
	Quantity  int
}
