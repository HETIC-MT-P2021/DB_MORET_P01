package structs

import (
	"github.com/valmrt77/DB_MORET_P01/models"
)

// OrderView : Architecture rendered to view
type OrderView struct {
	OrderID  int              `json:"orderID"`
	Products []models.Product `json:"products"`
}
