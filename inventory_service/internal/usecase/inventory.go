package usecase

import (
	"context"
	"inventory_service/internal/model"
)

type Inventory struct {
	invRepo InventoryRepostiry
}

func NewInventory(invRepo InventoryRepostiry) *Inventory {
	return &Inventory{invRepo: invRepo}
}

func (u *Inventory) Create(ctx context.Context, request model.Inventory) (model.Inventory, error) {
	panic("implement me")
}

func (u *Inventory) GetList(ctx context.Context, filters model.Filters) ([]model.Inventory, error) {
	panic("implement me")
}

func (u *Inventory) Get(ctx context.Context, id int64) (model.Inventory, error) {
	panic("implement me")
}

func (u *Inventory) Update(ctx context.Context, request model.Inventory) (model.Inventory, error) {
	panic("implement me")
}

func (u *Inventory) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}
