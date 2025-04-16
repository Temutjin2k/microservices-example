package frontend

import (
	"context"
	"errors"
	"inventory_service/internal/adapter/grpc/genproto/inventorypb"
	"inventory_service/internal/adapter/grpc/server/frontend/dto"
	"inventory_service/internal/model"
	"inventory_service/pkg/validator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Inventory struct {
	inventorypb.UnimplementedInventoryServiceServer
	uc InventoryUsecase
}

func New(uc InventoryUsecase) *Inventory {
	return &Inventory{uc: uc}
}

func (h *Inventory) CreateInventory(ctx context.Context, req *inventorypb.CreateInventoryRequest) (*inventorypb.CreateInventoryResponse, error) {
	// Step 1: Convert request to domain model
	inventory := dto.FromCreateReqestToInventory(req)

	// Step 2: Validate
	v := validator.New()
	if dto.ValidateInventory(v, inventory); !v.Valid() {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", v.Errors)
	}

	// Step 3: Call use case
	newInventory, err := h.uc.Create(ctx, inventory)
	if err != nil {
		errCtx := dto.GRPCFromError(err)
		return nil, status.Errorf(codes.Code(errCtx.Code), "%s", errCtx.Message)
	}

	// Step 4: Convert domain model to gRPC response
	res := &inventorypb.CreateInventoryResponse{
		Id:   newInventory.ID,
		Name: newInventory.Name,
	}

	return res, nil
}

// GetInventoryList implements gRPC handler for fetching a paginated list
func (h *Inventory) GetInventoryList(ctx context.Context, req *inventorypb.GetInventoryListRequest) (*inventorypb.GetInventoryListResponse, error) {
	v := validator.New()

	filters := model.Filters{
		Page:         int(req.GetPage()),
		PageSize:     int(req.GetPageSize()),
		Sort:         req.GetSort(),
		SortSafelist: []string{"id", "name", "price", "-id", "-name", "-price"},
	}

	model.ValidateFilters(v, filters)
	if !v.Valid() {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", v.Errors)
	}

	inventories, metadata, err := h.uc.GetList(ctx, filters)
	if err != nil {
		return nil, dto.StatusFromDomainError(err)
	}

	var response inventorypb.GetInventoryListResponse
	for _, item := range inventories {
		response.Inventory = append(response.Inventory, dto.ToInventoryProto(item))
	}
	response.Metadata = &inventorypb.Metadata{
		CurrentPage:  int32(metadata.CurrentPage),
		PageSize:     int32(metadata.PageSize),
		FirstPage:    int32(metadata.FirstPage),
		LastPage:     int32(metadata.LastPage),
		TotalRecords: int32(metadata.TotalRecords),
	}
	return &response, nil
}

// GetInventoryByID implements gRPC handler for fetching inventory by ID
func (h *Inventory) GetInventoryByID(ctx context.Context, req *inventorypb.GetInventoryByIDRequest) (*inventorypb.Inventory, error) {
	item, err := h.uc.Get(ctx, req.GetId())
	if err != nil {
		return nil, dto.StatusFromDomainError(err)
	}
	return dto.ToInventoryProto(item), nil
}

// UpdateInventory implements gRPC handler for updating an inventory
func (h *Inventory) UpdateInventory(ctx context.Context, req *inventorypb.UpdateInventoryRequest) (*inventorypb.Inventory, error) {
	item := dto.ToInventoryUpdateModel(req)

	updated, err := h.uc.Update(ctx, item)
	if err != nil {
		if errors.Is(err, dto.ErrUnprocessableEntity) {
			v := validator.New()
			dto.ValidateInventory(v, updated)
			if !v.Valid() {
				return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", v.Errors)
			}
		}
		return nil, dto.StatusFromDomainError(err)
	}
	return dto.ToInventoryProto(updated), nil
}

// DeleteInventory implements gRPC handler for deleting an inventory
func (h *Inventory) DeleteInventory(ctx context.Context, req *inventorypb.DeleteInventoryRequest) (*inventorypb.DeleteInventoryResponse, error) {
	err := h.uc.Delete(ctx, req.GetId())
	if err != nil {
		return nil, dto.StatusFromDomainError(err)
	}
	return &inventorypb.DeleteInventoryResponse{Success: true}, nil
}
