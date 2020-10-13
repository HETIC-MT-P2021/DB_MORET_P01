package router

import (
	"github.com/gorilla/mux"
	"github.com/valmrt77/DB_MORET_P01/controllers"
)

// Load controller (handler)
func InitRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/customers/{id}", controllers.GetCustomer).Methods("GET")
	r.HandleFunc("/orders/{id}", controllers.GetOrder).Methods("GET")
	r.HandleFunc("/employees", controllers.GetEmployees).Methods("GET")
	r.HandleFunc("/offices", controllers.GetOffices).Methods("GET")
	return r
}