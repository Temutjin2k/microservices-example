package usecase

import (
	"context"
	"order_service/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, client model.Order) error
	GetWithFilter(ctx context.Context, filter model.OrderFilter) (model.Order, error)
	GetListWithFilter(ctx context.Context, filter model.OrderFilter) ([]model.Order, error)
	Update(ctx context.Context, filter model.OrderFilter, update model.OrderUpdateData) error
	Delete(ctx context.Context, filter model.OrderFilter) error
}
