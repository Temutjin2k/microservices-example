package router

import (
	"order_service/internal/model"
)

type InventoryGRPCRouter struct {
	// inventory svc.InventoryServiceServer
}

func NewInventoryGrpcRouter() *InventoryGRPCRouter {
	return &InventoryGRPCRouter{}
}

func (r *InventoryGRPCRouter) GetById(id int64) (model.Inventory, error) {
	panic("implement me")
}

func (r *InventoryGRPCRouter) SetAvailability(id, newAvailability int64, version int32) error {
	panic("implement me")
}
