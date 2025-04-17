package dto

import (
	"order_service/internal/adapter/grpc/genproto/orderpb"
	"order_service/internal/model"
)

func FromOrderCreateRequest(ctx *orderpb.OrderCreateRequest) model.Order {
	req := model.Order{
		CustomerName: ctx.GetCustomerName(),
	}

	for _, v := range ctx.GetItems() {
		reqItem := model.OrderItem{
			ProductID: v.GetProductId(),
			Quantity:  v.GetQuantity(),
		}

		req.OrderItems = append(req.OrderItems, reqItem)
	}

	return req
}

func FromOrderRequestToResponce(res model.OrderResponce) *orderpb.OrderCreateResponse {
	responce := &orderpb.OrderCreateResponse{
		OrderId:      res.OrderID,
		CustomerName: res.CustomerName,
		Total:        res.Total,
	}

	for _, v := range res.Items {
		item := &orderpb.OrderItemResponse{
			ProductId: v.ProductID,
			Name:      v.Name,
			Price:     v.Price,
			Status:    v.Status,
			Reason:    v.Reason,
		}
		responce.Items = append(responce.Items, item)
	}

	return responce
}

func ToOrderListResponse(orders []model.Order) *orderpb.OrderListResponse {
	responce := &orderpb.OrderListResponse{}
	for _, v := range orders {
		responce.Orders = append(responce.Orders, ToOrderResponceList(v))
	}

	return responce
}

func ToOrderResponce(order model.Order) *orderpb.OrderResponseWrapper {
	responce := &orderpb.OrderResponseWrapper{
		Order: &orderpb.Order{
			OrderId:      order.ID,
			CustomerName: order.CustomerName,
			Status:       order.Status,
		},
	}

	for _, v := range order.OrderItems {
		item := &orderpb.OrderItemRequest{
			ProductId: v.ProductID,
			Quantity:  v.Quantity,
		}
		responce.Order.Items = append(responce.Order.Items, item)
	}

	return responce
}

func ToOrderResponceList(order model.Order) *orderpb.Order {
	responce := &orderpb.Order{
		OrderId:      order.ID,
		CustomerName: order.CustomerName,
		Status:       order.Status,
	}

	for _, v := range order.OrderItems {
		item := &orderpb.OrderItemRequest{
			ProductId: v.ProductID,
			Quantity:  v.Quantity,
		}
		responce.Items = append(responce.Items, item)
	}

	return responce
}
