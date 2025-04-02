package handler

import (
	"context"
	"order_service/internal/model"
)

type OrderUsecase interface {
	Create(ctx context.Context, request model.Order) (model.Order, error)
	Get(ctx context.Context, id uint64) (model.Order, error)
	GetList(ctx context.Context) ([]model.Order, error)
	Update(ctx context.Context, request model.Order) (model.Order, error)
	Delete(ctx context.Context, id uint64) error
}
