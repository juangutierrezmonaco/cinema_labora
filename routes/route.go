package routes

import (
	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/cinema_labora/controllers"
)

func BuildRoutes(router *mux.Router) {
	// Theater Routes
	router.HandleFunc("/api/theater", controllers.CreateTheater).Methods("POST")
	router.HandleFunc("/api/theater", controllers.GetTheaters).Methods("GET")
	router.HandleFunc("/api/theater/{id}", controllers.GetTheaterByID).Methods("GET")
	router.HandleFunc("/api/theater/{id}", controllers.UpdateTheater).Methods("PUT")
	router.HandleFunc("/api/theater/{id}", controllers.DeleteTheater).Methods("DELETE")

	// Screening Routes
	router.HandleFunc("/api/screening", controllers.CreateScreening).Methods("POST")
	router.HandleFunc("/api/theater", controllers.GetScreenings).Methods("GET")
	router.HandleFunc("/api/theater/{id}", controllers.GetTheaters).Methods("GET")
	router.HandleFunc("/api/theater/{id}", controllers.UpdateTheater).Methods("PUT")
	router.HandleFunc("/api/theater/{id}", controllers.DeleteTheater).Methods("DELETE")
}
