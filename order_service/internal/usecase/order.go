package usecase

import (
	"context"
	"fmt"
	"order_service/internal/model"
)

type Order struct {
	orderRepo        OrderRepository
	inventoryService InventoryService
}

func NewOrder(orderRepo OrderRepository, inventoryService InventoryService) *Order {
	return &Order{
		orderRepo:        orderRepo,
		inventoryService: inventoryService,
	}
}

func (u *Order) Create(ctx context.Context, request model.Order) (model.OrderResponce, error) {

	// Metadata of items
	var orderItemResponces []model.OrderItemResponce
	var successOrderItems []model.OrderItem

	var totalPrice int64
	for _, item := range request.OrderItems {
		var orderItemResp model.OrderItemResponce
		orderItemResp.ProductID = item.ProductID

		// Getting inventory from inventory service
		inventoryItem, err := u.inventoryService.GetById(item.ProductID)
		if err != nil {
			orderItemResp.Status = "rejected"
			orderItemResp.Reason = err.Error()
			orderItemResponces = append(orderItemResponces, orderItemResp)
			continue
		}

		orderItemResp.Name = inventoryItem.Name

		price := inventoryItem.Price * item.Quantity
		newAvailability := inventoryItem.Available - item.Quantity

		if newAvailability < 0 {
			orderItemResp.Status = "rejected"
			orderItemResp.Reason = "insufficient_inventory"
			orderItemResponces = append(orderItemResponces, orderItemResp)
			continue
		}

		// Trying to set new availability
		err = u.inventoryService.SetAvailability(item.ProductID, newAvailability, inventoryItem.Version)
		if err != nil {
			orderItemResp.Status = "rejected"
			orderItemResp.Reason = err.Error()
			orderItemResponces = append(orderItemResponces, orderItemResp)
			continue
		}

		orderItemResp.Price = price
		orderItemResp.Status = "accepted"

		totalPrice += price

		successOrderItems = append(successOrderItems, item)
		orderItemResponces = append(orderItemResponces, orderItemResp)
	}

	// Inserting only accepted orders to database
	request.OrderItems = successOrderItems
	orderID, err := u.orderRepo.Create(ctx, request)
	if err != nil {
		return model.OrderResponce{}, err
	}

	fmt.Println(orderItemResponces)

	responce := model.OrderResponce{
		OrderID:      orderID,
		CustomerName: request.CustomerName,
		Items:        orderItemResponces,
		Total:        totalPrice,
	}

	return responce, nil
}

func (u *Order) GetList(ctx context.Context) ([]model.Order, error) {
	orders, err := u.orderRepo.GetListWithFilter(ctx, model.OrderFilter{})
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (u *Order) Get(ctx context.Context, id int64) (model.Order, error) {
	order, err := u.orderRepo.GetWithFilter(ctx, model.OrderFilter{ID: id})
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func (u *Order) SetStatus(ctx context.Context, req model.UpdateStatus) (model.Order, error) {
	order, err := u.Get(ctx, req.OrderID)
	if err != nil {
		return model.Order{}, err
	}

	var updatedOrder model.OrderUpdateData
	updatedOrder.ID = &order.ID
	updatedOrder.Status = &req.Status

	err = u.orderRepo.Update(ctx, updatedOrder)

	if err != nil {
		return model.Order{}, err
	}

	order.Status = req.Status
	return order, nil
}
