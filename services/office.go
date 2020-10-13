package services

import (
	"github.com/valmrt77/DB_MORET_P01/models"
	"github.com/valmrt77/DB_MORET_P01/structs"
)

// get offices with employees working in
func GetOffices() (structs.OfficeView, error) {

	var officesRender structs.OfficeView
	var officesWithEmployees []models.OfficeWithEmployees

	offices, err := models.GetOffices()
	if err != nil {
		return officesRender, err
	}

	for _, office := range offices {

		employees, err := models.GetEmployeesOffice(office.OfficeCode)
		if err != nil {
			return officesRender, err
		}

		officeWithEmployees := models.OfficeWithEmployees{
			office,
			employees,
		}

		officesWithEmployees = append(officesWithEmployees, officeWithEmployees)
	}

	// build returned object
	officesRender = structs.OfficeView{
		officesWithEmployees,
	}

	return officesRender, nil
}
