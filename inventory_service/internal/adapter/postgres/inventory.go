package postgres

import (
	"context"
	"fmt"
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
		INSERT INTO inventory (name, description, price, available)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRow(ctx, query,
		item.Name,
		item.Description,
		item.Price,
		item.Available,
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

func (r *InventoryRepository) GetList(ctx context.Context, filters model.Filters) ([]model.Inventory, int, error) {
	query := fmt.Sprintf(`
        SELECT count(*) OVER(), id, created_at, name, description, price, available, isdeleted, version
        FROM inventory
        WHERE isdeleted = false
        ORDER BY %s %s, id ASC
        LIMIT $1 OFFSET $2`, filters.SortColumn(), filters.SortDirection())

	args := []any{filters.Limit(), filters.Offset()}

	var totalRecords int
	var items []model.Inventory

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Inventory
		err := rows.Scan(
			&totalRecords,
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
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows error: %w", err)
	}

	return items, totalRecords, nil
}

func (r *InventoryRepository) Update(ctx context.Context, item *model.Inventory) error {
	query := `
		UPDATE inventory
		SET name = $1, description = $2, price = $3, available = $4, version = version + 1
		WHERE id = $5 AND version = $6
		RETURNING version
	`

	args := []any{
		item.Name,
		item.Description,
		item.Price,
		item.Available,
		item.ID,
		item.Version,
	}

	err := r.db.QueryRow(ctx, query, args...).Scan(&item.Version)
	if err != nil {
		return err
	}

	return nil
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
