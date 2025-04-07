package dto

import (
	"order_service/internal/model"

	"github.com/gin-gonic/gin"
)

type OrderCreateRequest struct {
	CustomerName string              `json:"customer_name"`
	OrderItems   []OrderItemsRequest `json:"items"`
}

type OrderItemsRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type OrderResponceRequest struct {
	OrderID      int64  `json:"order_id"`
	CustomerName string `json:"customer_name"`
}

func FromOrderCreateRequest(ctx *gin.Context) (model.Order, error) {
	var req OrderCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return model.Order{}, err
	}

	var order model.Order
	order.CustomerName = req.CustomerName

	for _, v := range req.OrderItems {
		orderItems := model.OrderItem{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		}
		order.OrderItems = append(order.OrderItems, orderItems)
	}

	return order, nil
}

func ToClientCreateResponse(order model.Order) OrderResponceRequest {
	return OrderResponceRequest{
		OrderID:      order.ID,
		CustomerName: order.CustomerName,
	}
}
