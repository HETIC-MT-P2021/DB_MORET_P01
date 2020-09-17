package controllers

import (
	"github.com/valmrt77/DB_MORET_P01/models"
	"github.com/valmrt77/DB_MORET_P01/services"
	"net/http"
)

func GetCustomer(w http.ResponseWriter, r http.Request) {
	customerId := 119
	customer := models.GetCustomer(customerId)

	services.WriteJSON(w, http.StatusOK, customer)
}