package usecase

import (
	"context"
	"inventory_service/internal/adapter/http/service/handler/dto"
	"inventory_service/internal/model"
	"inventory_service/pkg/validator"
)

type Inventory struct {
	invRepo InventoryRepostiry
}

func NewInventory(invRepo InventoryRepostiry) *Inventory {
	return &Inventory{invRepo: invRepo}
}

func (u *Inventory) Create(ctx context.Context, request model.Inventory) (model.Inventory, error) {
	id, err := u.invRepo.Create(ctx, request)
	if err != nil {
		return model.Inventory{}, err
	}

	request.ID = id
	return request, nil
}

func (u *Inventory) GetList(ctx context.Context, filters model.Filters) ([]model.Inventory, dto.Metadata, error) {
	v := validator.New()

	model.ValidateFilters(v, filters)
	if !v.Valid() {
		return nil, dto.Metadata{}, ErrInvalidFilters
	}

	items, totalRecords, err := u.invRepo.GetList(ctx, filters)
	if err != nil {
		return nil, dto.Metadata{}, err
	}

	metadata := dto.CalculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return items, metadata, nil
}

func (u *Inventory) Get(ctx context.Context, id int64) (model.Inventory, error) {
	inv, err := u.invRepo.Get(ctx, id)
	if err != nil {
		return model.Inventory{}, err
	}
	return inv, nil
}

func (u *Inventory) Update(ctx context.Context, request model.Inventory) (model.Inventory, error) {
	panic("implement me")
}

func (u *Inventory) Delete(ctx context.Context, id int64) error {
	err := u.invRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
