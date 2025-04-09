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
		return nil, dto.Metadata{}, dto.ErrInvalidFilters
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

func (u *Inventory) Update(ctx context.Context, request model.InventoryUpdateData) (model.Inventory, error) {
	item, err := u.invRepo.Get(ctx, *request.ID)
	if err != nil {
		return model.Inventory{}, err
	}

	if request.Version != nil && item.Version != *request.Version {
		return model.Inventory{}, dto.ErrEditConflict
	}

	// If the input.Name value is nil then we know that no corresponding "name" key/
	// value pair was provided in the request body. So we move on and leave the
	// movie record unchanged. Otherwise, we update the movie record with the new name
	// value. Importantly, because input.Name is a now a pointer to a string, we need
	// to dereference the pointer using the * operator to get the underlying value
	// before assigning it to our movie record.
	if request.Name != nil {
		item.Name = *request.Name
	}
	if request.Description != nil {
		item.Description = *request.Description
	}
	if request.Price != nil {
		item.Price = *request.Price
	}
	if request.Available != nil {
		item.Available = *request.Available
	}

	v := validator.New()
	if dto.ValidateInventory(v, item); !v.Valid() {
		return item, dto.ErrUnprocessableEntity
	}

	err = u.invRepo.Update(ctx, &item)
	if err != nil {
		return model.Inventory{}, err
	}

	return item, nil

}

func (u *Inventory) Delete(ctx context.Context, id int64) error {
	err := u.invRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
