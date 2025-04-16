package dto

import (
	"inventory_service/internal/adapter/grpc/genproto/inventorypb"
	"inventory_service/internal/model"
	"time"
)

func FromCreateReqestToInventory(ctx *inventorypb.CreateInventoryRequest) model.Inventory {
	return model.Inventory{
		Name:        ctx.GetName(),
		Description: ctx.GetDescription(),
		Price:       ctx.GetPrice(),
		Available:   ctx.GetAvailable(),
	}
}

func ToInventoryProto(inv model.Inventory) *inventorypb.Inventory {
	return &inventorypb.Inventory{
		Id:          inv.ID,
		Name:        inv.Name,
		Description: inv.Description,
		Price:       inv.Price,
		Available:   inv.Available,
		CreatedAt:   inv.CreatedAt.Format(time.RFC3339),
		Version:     inv.Version,
	}
}

func ToInventoryUpdateModel(req *inventorypb.UpdateInventoryRequest) model.InventoryUpdateData {
	return model.InventoryUpdateData{
		ID:          &req.Id,
		Name:        nullString(req.Name),
		Description: nullString(req.Description),
		Price:       nullFloat64(req.Price),
		Available:   nullInt64(req.Available),
		Version:     &req.ExpectedVersion,
	}
}

func nullString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
func nullFloat64(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}
func nullInt64(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}
