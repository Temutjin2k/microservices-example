package handler

import (
	"context"
	"inventory_service/internal/adapter/http/service/handler/dto"
	"inventory_service/internal/model"
)

type InventoryUseCase interface {
	Create(ctx context.Context, request model.Inventory) (model.Inventory, error)
	GetList(ctx context.Context, filters model.Filters) ([]model.Inventory, dto.Metadata, error)
	Get(ctx context.Context, id int64) (model.Inventory, error)
	Update(ctx context.Context, request model.InventoryUpdateData) (model.Inventory, error)
	Delete(ctx context.Context, id int64) error
}
