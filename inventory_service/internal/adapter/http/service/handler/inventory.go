package handler

import "github.com/gin-gonic/gin"

type Inventory struct {
	invUseCase InventoryUseCase
}

func NewInventory(invUseCase InventoryUseCase) *Inventory {
	return &Inventory{invUseCase: invUseCase}
}

func (h *Inventory) Create(ctx *gin.Context) {

}

func (h *Inventory) GetList(ctx *gin.Context) {

}

func (h *Inventory) GetByID(ctx *gin.Context) {

}

func (h *Inventory) Update(ctx *gin.Context) {

}

func (h *Inventory) Delete(ctx *gin.Context) {

}
