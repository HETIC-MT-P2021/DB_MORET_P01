package services

import (
	"github.com/valmrt77/DB_MORET_P01/models"
	"github.com/valmrt77/DB_MORET_P01/structs"
)

// get employees and infos of their offices
func GetEmployees() (structs.EmployeesView, error) {

	var employeesRender structs.EmployeesView

	employees, err := models.GetEmployeesWithOffice()
	if err != nil {
		return employeesRender, err
	}

	// build returned object
	employeesRender = structs.EmployeesView{
		employees,
	}

	return employeesRender, nil
}
