package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valmrt77/DB_MORET_P01/models"
	"github.com/valmrt77/DB_MORET_P01/router"
)

func main() {
	r := mux.NewRouter()
	router.InitRoutes(r)

	models.ConnectToBDD()

	log.Fatal(http.ListenAndServe(":9000", r))
}
