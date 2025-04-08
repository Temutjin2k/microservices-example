package handler

import (
	"inventory_service/internal/adapter/http/service/handler/dto"
	"inventory_service/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Inventory struct {
	invUseCase InventoryUseCase
}

func NewInventory(invUseCase InventoryUseCase) *Inventory {
	return &Inventory{invUseCase: invUseCase}
}

func (h *Inventory) Create(ctx *gin.Context) {
	inventory, err := dto.ToInventoryCreateRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v := validator.New()
	if dto.ValidateInventory(v, inventory); !v.Valid() {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": v.Errors})
		return
	}

	inventoryNew, err := h.invUseCase.Create(ctx.Request.Context(), inventory)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"inventory": dto.ToInventoryCreateResponce(inventoryNew)})
}

func (h *Inventory) GetList(ctx *gin.Context) {

}

func (h *Inventory) GetByID(ctx *gin.Context) {

}

func (h *Inventory) Update(ctx *gin.Context) {

}

func (h *Inventory) Delete(ctx *gin.Context) {

}
