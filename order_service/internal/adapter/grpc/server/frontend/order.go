package frontend

import (
	"context"
	"errors"
	"order_service/internal/adapter/grpc/genproto/orderpb"
	"order_service/internal/adapter/grpc/server/frontend/dto"
	"order_service/internal/model"
	"order_service/pkg/validator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Order implements orderpb.OrderServiceServer
type Order struct {
	orderpb.UnimplementedOrderServiceServer
	uc OrderUsecase
}

func New(uc OrderUsecase) *Order {
	return &Order{uc: uc}
}

func (h *Order) CreateOrder(ctx context.Context, req *orderpb.OrderCreateRequest) (*orderpb.OrderCreateResponse, error) {
	// Convert gRPC request to domain model
	order := dto.FromOrderCreateRequest(req)

	// Validate request
	v := validator.New()
	dto.ValidateOrder(v, order)
	if !v.Valid() {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", v.Errors)
	}

	newOrder, err := h.uc.Create(ctx, order)
	if err != nil {
		return nil, toGRPCError(err)
	}

	// Convert domain model to gRPC response
	return dto.FromOrderRequestToResponce(newOrder), nil
}

func (h *Order) GetOrderById(ctx context.Context, req *orderpb.OrderID) (*orderpb.OrderResponseWrapper, error) {
	// Validate request
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "order ID is required")
	}

	// Call use case
	order, err := h.uc.Get(ctx, req.GetId())
	if err != nil {
		return nil, toGRPCError(err)
	}

	// Convert domain model to gRPC response
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
	return dto.ToOrderResponce(order), nil
}

func (h *Order) GetOrderList(ctx context.Context, req *orderpb.Empty) (*orderpb.OrderListResponse, error) {
	// Call use case
	orders, err := h.uc.GetList(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	// Convert domain models to gRPC response
	return dto.ToOrderListResponse(orders), nil
}

func (h *Order) SetOrderStatus(ctx context.Context, req *orderpb.SetOrderStatusRequest) (*orderpb.OrderResponseWrapper, error) {
	// Validate request
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "order ID is required")
	}
	if req.GetStatus() == "" {
		return nil, status.Error(codes.InvalidArgument, "status is required")
	}

	// Convert to domain model
	update := model.UpdateStatus{
		OrderID: req.GetId(),
		Status:  req.GetStatus(),
	}

	// Call use case
	order, err := h.uc.SetStatus(ctx, update)
	if err != nil {
		return nil, toGRPCError(err)
	}

	// Convert domain model to gRPC response
	return dto.ToOrderResponce(order), nil
}

// toGRPCError converts domain errors to gRPC status errors
func toGRPCError(err error) error {
	if err == nil {
		return nil
	}

	// Handle specific domain errors
	switch {
	case errors.Is(err, dto.ErrOrderNotFound):
		return status.Error(codes.NotFound, "order not found")
	case errors.Is(err, dto.ErrInvalidStatusTransition):
		return status.Error(codes.FailedPrecondition, "invalid status transition")
	// Add more specific error cases as needed
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
