package dto

import (
	"inventory_service/internal/model"
	"inventory_service/pkg/validator"
)

func ValidateInventory(v *validator.Validator, inv model.Inventory) {
	v.Check(len(inv.Name) != 0, "name", "must be provided")
	v.Check(len(inv.Name) < 50, "name", "must not be more than 50 bytes long")

	v.Check(len(inv.Description) != 0, "description", "must be provided")
	v.Check(inv.Price > 0, "price", "must be positive")
	v.Check(inv.Available > 0, "availabilty", "must be positive")
}
