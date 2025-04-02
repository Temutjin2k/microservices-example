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

func (r *Order) Create(ctx context.Context, client model.Order) error {
	panic("impliment me")
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
