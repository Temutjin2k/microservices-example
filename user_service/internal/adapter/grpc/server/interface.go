package server

import (
	"user_service/internal/adapter/grpc/frontend"
)

type UserUseCase interface {
	frontend.UserUseCase
}
