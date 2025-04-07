package handler

import (
	"net/http"
	"order_service/internal/adapter/http/service/handler/dto"
	"order_service/internal/model"
	"order_service/pkg/validator"

	"github.com/gin-gonic/gin"
)

// OrderHandler
type Order struct {
	uc OrderUsecase
}

func NewClient(uc OrderUsecase) *Order {
	return &Order{
		uc: uc,
	}
}

func (c *Order) Create(ctx *gin.Context) {
	order, err := dto.FromOrderCreateRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v := validator.New()

	if dto.ValidateOrder(v, order); !v.Valid() {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": v.Errors})
		return
	}

	newOrder, err := c.uc.Create(ctx.Request.Context(), order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"order": dto.ToOrderCreateResponse(newOrder)})
}

func (c *Order) GetList(ctx *gin.Context) {
	orders, err := c.uc.GetList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"orders": dto.ToOrderListResponce(orders)})
}

func (c *Order) GetByID(ctx *gin.Context) {
	id, err := dto.ReadIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	// Get order from service
	order, err := c.uc.Get(ctx.Request.Context(), id)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"order": dto.ToOrderResponce(order)})
}

func (c *Order) SetStatus(ctx *gin.Context) {
	id, err := dto.ReadIDParam(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	var request dto.OrderSetStatusRequest

	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	v := validator.New()
	if dto.ValidateSetOrderStatusRequest(v, request); !v.Valid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": v.Errors})
		return
	}

	order, err := c.uc.SetStatus(ctx.Request.Context(), model.UpdateStatus{
		OrderID: id,
		Status:  request.Status,
	})
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"order": dto.ToOrderResponce(order)})
}

// func (c *Order) SetStatus(ctx *gin.Context) {
// 	id, err := dto.ReadIDParam(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
// 		return
// 	}

// 	// Get order from service
// 	order, err := c.uc.Get(ctx.Request.Context(), id)
// 	if err != nil {
// 		errCtx := dto.FromError(err)
// 		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})
// 		return
// 	}

// 	var request dto.OrderSetStatusRequest

// 	err = ctx.ShouldBindJSON(&request)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
// 		return
// 	}

// 	v := validator.New()
// 	if dto.ValidateSetOrderStatusRequest(v, request); !v.Valid() {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": v.Errors})
// 		return
// 	}
// 	order.Status = request.Status

// 	order, err = c.uc.Update(ctx.Request.Context(), order)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"order": dto.ToOrderResponce(order)})
// }
