package postgres

import (
	"context"
	"order_service/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Order struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *Order {
	return &Order{db: db}
}

func (r *Order) Create(ctx context.Context, order model.Order) (int64, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO orders (customername, status) 
		VALUES ($1, $2)
		RETURNING ID;
	`

	var orderID int64
	err = tx.QueryRow(ctx, query, order.CustomerName, order.Status).Scan(&orderID)
	if err != nil {
		return 0, err
	}

	// Inserting order items. in case when same product id is given, it check on conflict, if so it's just adding quantity for previus row.
	queryOrderItems := `
		INSERT INTO order_items (OrderID, ProductID, Quantity) VALUES
		($1, $2, $3)
		ON CONFLICT (OrderID, ProductID)
		DO UPDATE SET Quantity = order_items.Quantity + EXCLUDED.Quantity;
	`

	for _, v := range order.OrderItems {
		_, err = tx.Exec(ctx, queryOrderItems, orderID, v.ProductID, v.Quantity)
		if err != nil {
			return 0, err
		}
	}

	return orderID, tx.Commit(ctx)
}

func (r *Order) GetWithFilter(ctx context.Context, filter model.OrderFilter) (model.Order, error) {
	panic("impliment me")
}

func (r *Order) GetListWithFilter(ctx context.Context, filter model.OrderFilter) ([]model.Order, error) {
	panic("impliment me")

}

func (r *Order) Update(ctx context.Context, filter model.OrderFilter, update model.OrderUpdateData) error {
	panic("impliment me")

}

func (r *Order) Delete(ctx context.Context, filter model.OrderFilter) error {
	panic("impliment me")
}
