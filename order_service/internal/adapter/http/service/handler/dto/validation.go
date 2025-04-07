package dto

import (
	"fmt"
	"order_service/internal/adapter/postgres/dao"
	"order_service/internal/model"
	"order_service/pkg/validator"
	"strings"
)

func ValidateOrder(v *validator.Validator, order model.Order) {
	// v.Check(order.ID >= 0, "order_id", "must be equal or greater than zero")

	v.Check(order.CustomerName != "", "customer_name", "must be provided")
	v.Check(len(order.CustomerName) < 50, "customer_name", "must not be more than 50 bytes long")

	for _, item := range order.OrderItems {
		v.Check(item.ProductID > 0, "items_product_id", "must be greater than zero")
		v.Check(item.Quantity > 0, "items_quantity", "must be greater than zero")
		v.Check(item.Quantity <= 100, "items_quantity", "item quantity cannot be greater than 100")
	}
}

func ValidateSetOrderStatusRequest(v *validator.Validator, req OrderSetStatusRequest) {
	safeList := []string{dao.OrderStatusCanceled, dao.OrderStatusCompleted, dao.OrderStatusPending}
	v.Check(validator.PermittedValue(req.Status, safeList...), "status", fmt.Sprintf("invalid status. Available: %v", strings.Join(safeList, ", ")))
}
