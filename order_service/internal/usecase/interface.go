package usecase

import (
	"context"
	"order_service/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order model.Order) (int64, error)
	GetWithFilter(ctx context.Context, filter model.OrderFilter) (model.Order, error)
	GetListWithFilter(ctx context.Context, filter model.OrderFilter) ([]model.Order, error)
	Update(ctx context.Context, update model.OrderUpdateData) error
}

type InventoryService interface {
	GetById(id int64) (model.Inventory, error)
	Substruct(id, newAvailability int64, version int32) error
}
