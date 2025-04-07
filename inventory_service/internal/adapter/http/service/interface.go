package service

import "inventory_service/internal/adapter/http/service/handler"

type InventoryUsecase interface {
	handler.InventoryUseCase
}
