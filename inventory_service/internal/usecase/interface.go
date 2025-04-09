package usecase

import (
	"context"
	"inventory_service/internal/model"
)

type InventoryRepostiry interface {
	Create(ctx context.Context, item model.Inventory) (int64, error)
	Get(ctx context.Context, id int64) (model.Inventory, error)
	GetList(ctx context.Context, filters model.Filters) ([]model.Inventory, int, error)
	Update(ctx context.Context, item model.Inventory) error
	Delete(ctx context.Context, id int64) error
}
