package structs

import "github.com/valmrt77/DB_MORET_P01/models"

// CustomerView : Architecture rendered to view
type CustomerView struct {
	models.Customer
	Orders []models.TotalAndPriceProduct `json:"orders"`
}
