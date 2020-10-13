package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/valmrt77/DB_MORET_P01/services"
)

func GetOrder(w http.ResponseWriter, req *http.Request) {

	urlParams := mux.Vars(req)
	idStr := urlParams["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		services.WriteErrorJSON(w, http.StatusInternalServerError, "Impossible de récupérer la commande")
		return
	}

	order, err := services.GetOrder(id)
	if err != nil {
		log.Println(err)
		services.WriteErrorJSON(w, http.StatusInternalServerError, "Impossible de récupérer la commande")
		return
	}

	services.WriteJSON(w, http.StatusOK, order)
}
