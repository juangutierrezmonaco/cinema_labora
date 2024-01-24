package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func BuildRoutes(router *mux.Router) {
	// Routes
	router.HandleFunc("/api/movie", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")
	router.HandleFunc("/api/movie/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("PUT")
	router.HandleFunc("/api/movie", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	router.HandleFunc("/api/movie/{id}", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	router.HandleFunc("/api/movie", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
}
