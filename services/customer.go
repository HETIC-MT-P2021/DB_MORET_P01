package services

import (
	"github.com/valmrt77/DB_MORET_P01/models"
	"github.com/valmrt77/DB_MORET_P01/structs"
)

// Get customer data ordered infos
func GetCustomer(customerID int) (structs.CustomerView, error) {

	var customerRender structs.CustomerView
	var totalAndPriceProducts []models.TotalAndPriceProduct

	customer, err := models.GetCustomer(customerID)
	if err != nil {
		return structs.CustomerView{}, err
	}

	// build returned object
	customerRender = structs.CustomerView{
		customer,
		totalAndPriceProducts,
	}

	return customerRender, nil
}
