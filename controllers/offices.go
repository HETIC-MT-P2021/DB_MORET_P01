package controllers

import (
	"log"
	"net/http"

	"github.com/valmrt77/DB_MORET_P01/services"
)

func GetOffices(w http.ResponseWriter, req *http.Request) {

	order, err := services.GetOffices()
	if err != nil {
		log.Println(err)
		services.WriteErrorJSON(w, http.StatusInternalServerError, "Impossible de récupérer les magasins")
		return
	}

	services.WriteJSON(w, http.StatusOK, order)
}
