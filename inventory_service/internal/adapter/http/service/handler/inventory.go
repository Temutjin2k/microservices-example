package handler

import "github.com/gin-gonic/gin"

type Inventory struct {
	invUseCase InventoryUseCase
}

func NewInventory(invUseCase InventoryUseCase) *Inventory {
	return &Inventory{invUseCase: invUseCase}
}

func Create(ctx *gin.Context) {

}

func GetList(ctx *gin.Context) {

}

func GetByID(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}
