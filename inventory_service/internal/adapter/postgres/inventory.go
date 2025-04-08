package postgres

import (
	"context"
	"inventory_service/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type InventoryRepository struct {
	db *pgxpool.Pool
}

func NewInventoryRepository(db *pgxpool.Pool) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) Create(ctx context.Context, item model.Inventory) (int64, error) {
	panic("implement me")
}

func (r *InventoryRepository) Get(ctx context.Context, id int64) (model.Inventory, error) {
	panic("implement me")
}

func (r *InventoryRepository) GetList(ctx context.Context, filters model.Filters) ([]model.Inventory, error) {
	panic("implement me")
}

func (r *InventoryRepository) Update(ctx context.Context, item model.Inventory) error {
	panic("implement me")
}

func (r *InventoryRepository) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}
