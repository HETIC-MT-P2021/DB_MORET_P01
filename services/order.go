package services

import (
	"github.com/valmrt77/DB_MORET_P01/models"
	"github.com/valmrt77/DB_MORET_P01/structs"
)

// Get order and products in
func GetOrder(orderID int) (structs.OrderView, error) {

	var orderRender structs.OrderView
	var products []models.Product

	orderID, orderProductsCode, err := models.GetProductsCode(10123)
	if err != nil {
		return orderRender, err
	}

	for _, productCode := range orderProductsCode {

		product, err := models.GetProduct(productCode)
		if err != nil {
			return orderRender, err
		}

		products = append(products, product)
	}

	// build returned object
	orderRender = structs.OrderView{
		orderID,
		products,
	}

	return orderRender, err
}
