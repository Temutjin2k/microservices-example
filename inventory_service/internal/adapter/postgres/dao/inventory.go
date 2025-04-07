package dao

import "time"

type Product struct {
	ID          int64
	CreatedAt   time.Time
	Name        string
	Description string
	Price       float64
	Available   *int

	IsDeleted bool
	Version   int
}
