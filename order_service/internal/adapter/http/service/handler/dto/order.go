package dto

import (
	"order_service/internal/adapter/postgres/dao"
	"order_service/internal/model"
	"time"

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

type OrderCreateResponceRequest struct {
	OrderID      int64  `json:"order_id"`
	CustomerName string `json:"customer_name"`
}

type OrderResponce struct {
	OrderID      int64               `json:"order_id"`
	CustomerName string              `json:"customer_name"`
	Items        []OrderItemsRequest `json:"items"`
	Status       string              `json:"status"`
	CreatedAt    time.Time           `json:"created_at"`
}

type OrderSetStatusRequest struct {
	Status string `json:"status"`
}

func FromOrderCreateRequest(ctx *gin.Context) (model.Order, error) {
	var req OrderCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return model.Order{}, err
	}

	var order model.Order
	order.CustomerName = req.CustomerName
	order.Status = dao.OrderStatusPending

	for _, v := range req.OrderItems {
		orderItems := model.OrderItem{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		}
		order.OrderItems = append(order.OrderItems, orderItems)
	}

	return order, nil
}

func ToOrderCreateResponse(order model.Order) OrderCreateResponceRequest {
	return OrderCreateResponceRequest{
		OrderID:      order.ID,
		CustomerName: order.CustomerName,
	}
}

func ToOrderListResponce(orders []model.Order) []OrderResponce {
	resp := []OrderResponce{}

	for _, order := range orders {
		orderResponce := ToOrderResponce(order)
		resp = append(resp, orderResponce)
	}

	return resp
}

func ToOrderResponce(order model.Order) OrderResponce {
	var orderResponce OrderResponce

	orderResponce.OrderID = order.ID
	orderResponce.CustomerName = order.CustomerName
	orderResponce.Status = order.Status
	orderResponce.CreatedAt = order.Created_at

	for _, item := range order.OrderItems {
		var itemRequest OrderItemsRequest
		itemRequest.ProductID = item.ProductID
		itemRequest.Quantity = item.ProductID
		orderResponce.Items = append(orderResponce.Items, itemRequest)
	}

	return orderResponce
}
