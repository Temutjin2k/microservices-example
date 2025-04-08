package model

import "time"

type Inventory struct {
	ID          int64
	CreatedAt   time.Time
	Name        string
	Description string
	Price       float64
	Available   int64
	IsDeleted   bool

	Version int32
}

type InventoryUpdateData struct {
	ID          *int64
	CreatedAt   *time.Time
	Name        *string
	Description *string
	Price       *float64
	Available   *int64
	IsDeleted   *bool

	Version *int32
}
