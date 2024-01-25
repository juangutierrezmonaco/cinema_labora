package routes

import (
	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/cinema_labora/controllers"
)

func BuildRoutes(router *mux.Router) {
	// Theater Routes
	router.HandleFunc("/api/theater", controllers.CreateTheater).Methods("POST")
	router.HandleFunc("/api/theater", controllers.GetAllTheaters).Methods("GET")
	router.HandleFunc("/api/theater/{id}", controllers.GetAllTheaters).Methods("GET")
	router.HandleFunc("/api/theater/{id}", controllers.UpdateTheater).Methods("PUT")
	router.HandleFunc("/api/theater/{id}", controllers.DeleteTheater).Methods("DELETE")
}
