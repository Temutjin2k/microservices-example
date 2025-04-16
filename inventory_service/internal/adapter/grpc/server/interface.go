package server

import (
	"inventory_service/internal/adapter/grpc/server/frontend"
)

type InventoryUsecase interface {
	frontend.InventoryUsecase
}
