package controllers

import (
	"github.com/valmrt77/DB_MORET_P01/models"
	"github.com/valmrt77/DB_MORET_P01/services"
	"log"
	"net/http"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := 119
	customers, err := models.GetCustomer(customerId)

	if err!= nil{
		log.Println(err)
	}

	services.WriteJSON(w, http.StatusOK, customers)
}