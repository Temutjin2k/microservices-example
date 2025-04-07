package handler

import (
	"github.com/gin-gonic/gin"
)

// OrderHandler
type Order struct {
	uc OrderUsecase
}

func NewClient(uc OrderUsecase) *Order {
	return &Order{
		uc: uc,
	}
}

func (c *Order) Create(ctx *gin.Context) {
	
}

func (c *Order) GetList(ctx *gin.Context) {

}

func (c *Order) GetByID(ctx *gin.Context) {

}

func (c *Order) Delete(ctx *gin.Context) {

}
