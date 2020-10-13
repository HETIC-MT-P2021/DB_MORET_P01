package structs

import (
	"github.com/valmrt77/DB_MORET_P01/models"
)

// OfficeView : Architecture rendered to view
type OfficeView struct {
	Offices []models.OfficeWithEmployees `json:"offices"`
}
