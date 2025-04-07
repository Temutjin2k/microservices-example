package handler

import (
	"context"
	"order_service/internal/model"
)

type OrderUsecase interface {
	Create(ctx context.Context, request model.Order) (model.Order, error)
	Get(ctx context.Context, id int64) (model.Order, error)
	GetList(ctx context.Context) ([]model.Order, error)
	SetStatus(ctx context.Context, request model.UpdateStatus) (model.Order, error)
}
