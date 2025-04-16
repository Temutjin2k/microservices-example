package frontend

import (
	"context"
	"user_service/internal/adapter/grpc/genproto/userpb"
)

type User struct {
	userpb.UnimplementedUserServiceServer

	uc UserUseCase
}

func NewUser(uc UserUseCase) *User {
	return &User{uc: uc}
}

func (h *User) AuthenticateUser(context.Context, *userpb.AuthRequest) (*userpb.AuthResponse, error) {
	panic("implement me")
}
func (h *User) GetUserProfile(context.Context, *userpb.UserID) (*userpb.UserProfile, error) {
	panic("implement me")
}

func (h *User) RegisterUser(context.Context, *userpb.UserRequest) (*userpb.UserResponse, error) {
	panic("implement me")
}
