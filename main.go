package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/valmrt77/DB_MORET_P01/router"
)

func main() {
	r := mux.NewRouter()
	router.InitRoutes(r)
	log.Fatal(http.ListenAndServe(":9000", r))
}