package router

import (
	"github.com/gorilla/mux"
	"github.com/valmrt77/DB_MORET_P01/controllers"
)

// Load controller (handler)
func InitRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/customers", controllers.GetCustomer).Methods("GET")
	return r
}