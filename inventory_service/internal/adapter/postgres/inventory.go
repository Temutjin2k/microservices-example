package postgres

import (
	"context"
	"inventory_service/internal/adapter/postgres/dao"
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
	query := `
		INSERT INTO inventory (name, description, price, available, isdeleted, version)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRow(ctx, query,
		item.Name,
		item.Description,
		item.Price,
		item.Available,
		item.IsDeleted,
		item.Version,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *InventoryRepository) Get(ctx context.Context, id int64) (model.Inventory, error) {
	query := `
		SELECT id, created_at, name, description, price, available, isdeleted, version
		FROM inventory
		WHERE id = $1 AND isdeleted = false
	`

	var item model.Inventory
	err := r.db.QueryRow(ctx, query, id).Scan(
		&item.ID,
		&item.CreatedAt,
		&item.Name,
		&item.Description,
		&item.Price,
		&item.Available,
		&item.IsDeleted,
		&item.Version,
	)
	if err != nil {
		return model.Inventory{}, err
	}

	return item, nil
}

func (r *InventoryRepository) GetList(ctx context.Context, filters model.Filters) ([]model.Inventory, error) {
	panic("implement me")
}

func (r *InventoryRepository) Update(ctx context.Context, item model.Inventory) error {
	panic("implement me")
}

func (r *InventoryRepository) Delete(ctx context.Context, id int64) error {
	query := `
		UPDATE inventory
		SET isdeleted = true,
			version = version + 1
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return dao.ErrRecordNotFound
	}

	return nil
}
