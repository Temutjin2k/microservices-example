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
	panic("implement me")
}
func (u *Order) Get(ctx context.Context, id uint64) (model.Order, error) {
	panic("implement me")

}
func (u *Order) GetList(ctx context.Context) ([]model.Order, error) {
	panic("implement me")

}
func (u *Order) Update(ctx context.Context, request model.Order) (model.Order, error) {
	panic("implement me")

}
func (u *Order) Delete(ctx context.Context, id uint64) error {
	panic("implement me")
}
