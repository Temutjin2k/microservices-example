package handler

import (
	"net/http"
	"order_service/internal/adapter/http/service/handler/dto"

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
	order, err := dto.FromOrderCreateRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder, err := c.uc.Create(ctx.Request.Context(), order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"order": dto.ToClientCreateResponse(newOrder)})
}

func (c *Order) GetList(ctx *gin.Context) {

}

func (c *Order) GetByID(ctx *gin.Context) {

}

func (c *Order) Delete(ctx *gin.Context) {

}
