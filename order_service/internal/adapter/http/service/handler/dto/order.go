package dto

import (
	"order_service/internal/model"

	"github.com/gin-gonic/gin"
)

type OrderCreateRequest struct {
	CustomerName string `json:"customerName"`
}

type OrderResponceRequest struct {
}

func FromClientCreateRequest(ctx *gin.Context) (model.Order, error) {
	panic("implement me")
}
