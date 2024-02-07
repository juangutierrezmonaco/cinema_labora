package routes

import (
	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/cinema_labora/controllers"
	"github.com/labora/labora-golang/cinema_labora/middlewares"
)

func BuildRoutes(router *mux.Router) {
	// Logging every request info
	router.Use(middlewares.LoggingMiddleware)

	// Theater Routes
	router.HandleFunc("/api/theater", controllers.CreateTheater).Methods("POST")
	router.HandleFunc("/api/theater", controllers.GetTheaters).Methods("GET")
	router.HandleFunc("/api/theater/{id}", controllers.GetTheaterByID).Methods("GET")
	router.HandleFunc("/api/theater/{id}", controllers.UpdateTheater).Methods("PUT")
	router.HandleFunc("/api/theater/{id}", controllers.DeleteTheater).Methods("DELETE")

	// Screening Routes
	router.HandleFunc("/api/screening", controllers.CreateScreening).Methods("POST")
	router.HandleFunc("/api/screening", controllers.GetScreenings).Methods("GET")
	router.HandleFunc("/api/screening/{id}", controllers.GetScreeningByID).Methods("GET")
	router.HandleFunc("/api/screening/movie/{id}", controllers.GetScreeningByMovieIdOrTheaterId).Methods("GET")
	router.HandleFunc("/api/screening/theater/{id}", controllers.GetScreeningByMovieIdOrTheaterId).Methods("GET")
	router.HandleFunc("/api/screening/{id}", controllers.UpdateScreening).Methods("PUT")
	router.HandleFunc("/api/screening/{id}", controllers.DeleteScreening).Methods("DELETE")

	// User Routes
	router.HandleFunc("/api/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE")

	// Comment Routes
	router.HandleFunc("/api/comment", controllers.CreateComment).Methods("POST")
	router.HandleFunc("/api/comment", controllers.GetComments).Methods("GET")
	router.HandleFunc("/api/comment/{id}", controllers.GetCommentByID).Methods("GET")
	router.HandleFunc("/api/comment/{id}", controllers.UpdateComment).Methods("PUT")
	router.HandleFunc("/api/comment/{id}", controllers.DeleteComment).Methods("DELETE")

	// Ticket Routes
	router.HandleFunc("/api/ticket", controllers.CreateTicket).Methods("POST")
	router.HandleFunc("/api/ticket", controllers.GetTickets).Methods("GET")
	router.HandleFunc("/api/ticket/{id}", controllers.GetTicketByID).Methods("GET")
	router.HandleFunc("/api/ticket/{id}", controllers.UpdateTicket).Methods("PUT")
	router.HandleFunc("/api/ticket/{id}", controllers.DeleteTicket).Methods("DELETE")

	// Movie Routes
	router.HandleFunc("/api/movie", controllers.GetMoviesByTitleAndOverview).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controllers.GetTMDBMovieDetails).Methods("GET")
}
