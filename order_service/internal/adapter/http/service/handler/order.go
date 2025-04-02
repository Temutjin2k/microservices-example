package handler

import (
	"github.com/gin-gonic/gin"
)

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
