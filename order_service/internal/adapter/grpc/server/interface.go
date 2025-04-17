package server

import "order_service/internal/adapter/grpc/server/frontend"

type OrderUsecase interface {
	frontend.OrderUsecase
}
