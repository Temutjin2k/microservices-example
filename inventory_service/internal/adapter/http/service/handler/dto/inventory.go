package dto

import (
	"inventory_service/internal/model"
	"inventory_service/pkg/validator"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type InventoryCreateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Available   int64   `json:"available"`
}

type InventoryCreateResponce struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type InventoryResponse struct {
	ID          int64     `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Available   int64     `json:"available,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitzero"`
	Version     int32     `json:"version,omitempty"`
}

func ToInventoryCreateRequest(ctx *gin.Context) (model.Inventory, error) {
	var req InventoryCreateRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return model.Inventory{}, err
	}

	inventory := model.Inventory{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Available:   req.Available,
	}

	return inventory, nil
}

func ToInventoryCreateResponce(inv model.Inventory) InventoryCreateResponce {
	return InventoryCreateResponce{
		ID:   inv.ID,
		Name: inv.Name,
	}
}

func ToInventoryListReponce(invs []model.Inventory) []InventoryResponse {
	var responce []InventoryResponse

	for _, v := range invs {
		responce = append(responce, ToInventoryResponse(v))
	}

	return responce
}

func ToInventoryResponse(inv model.Inventory) InventoryResponse {
	return InventoryResponse{
		ID:          inv.ID,
		Name:        inv.Name,
		Description: inv.Description,
		Price:       inv.Price,
		Available:   inv.Available,
		CreatedAt:   inv.CreatedAt,
		Version:     inv.Version,
	}
}

func ParseListRequest(ctx *gin.Context, v *validator.Validator) model.Filters {
	// Default values
	page := 1
	pageSize := 20
	sort := "id"
	sortSafelist := []string{"id", "name", "price", "-id", "-name", "-price"}

	// Parse page parameter
	if pageStr := ctx.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	// Parse page_size parameter
	if pageSizeStr := ctx.Query("page_size"); pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
			pageSize = ps
		}
	}

	// Parse sort parameter
	if sortParam := ctx.Query("sort"); sortParam != "" {
		sort = sortParam
	}

	filter := model.Filters{
		Page:         page,
		PageSize:     pageSize,
		Sort:         sort,
		SortSafelist: sortSafelist,
	}

	model.ValidateFilters(v, filter)

	return filter
}
