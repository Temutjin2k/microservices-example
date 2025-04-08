package dto

import (
	"inventory_service/internal/model"

	"github.com/gin-gonic/gin"
)

type InventoryCreateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Available   int64   `json:"available"`
}

type InventoryCreateResponce struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func ToInventoryCreateRequest(ctx *gin.Context) (model.Inventory, error) {
	var req InventoryCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return model.Inventory{}, err
	}

	inventory := model.Inventory{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Available:   req.Available,
	}

	return inventory, nil
}

func ToInventoryCreateResponce(inv model.Inventory) InventoryCreateResponce {
	return InventoryCreateResponce{
		ID:   inv.ID,
		Name: inv.Name,
	}
}
