package usecase

import (
	"context"
	"order_service/internal/model"
)

type Order struct {
	orderRepo OrderRepository
}

func NewOrder(orderRepo OrderRepository) *Order {
	return &Order{
		orderRepo: orderRepo,
	}
}

func (u *Order) Create(ctx context.Context, request model.Order) (model.Order, error) {
	orderID, err := u.orderRepo.Create(ctx, request)
	if err != nil {
		return model.Order{}, err
	}

	request.ID = orderID
	return request, nil
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
