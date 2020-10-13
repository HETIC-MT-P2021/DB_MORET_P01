package structs

import (
	"github.com/valmrt77/DB_MORET_P01/models"
)

// EmployeesView : Architecture rendered to view
type EmployeesView struct {
	Employees []models.EmployeeWithOffice `json:"employees"`
}
