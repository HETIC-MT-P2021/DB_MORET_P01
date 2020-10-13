package controllers

import (
	"log"
	"net/http"

	"github.com/valmrt77/DB_MORET_P01/services"
)

func GetEmployees(w http.ResponseWriter, req *http.Request) {

	customer, err := services.GetEmployees()
	if err != nil {
		log.Println(err)
		services.WriteErrorJSON(w, http.StatusInternalServerError, "Impossible de récupérer les employés")
		return
	}

	services.WriteJSON(w, http.StatusOK, customer)
}
