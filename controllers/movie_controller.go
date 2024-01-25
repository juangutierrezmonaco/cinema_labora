package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/services"
)

func GetTMDBMovieDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	tmdbMovie, err := services.GetTMDBMovieDetails(movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If exists, automatically updates the views
	movieCount, err := services.GetMovieCountByID(movieID)
	if err != nil {
		_, err := services.CreateMovieCount(models.MovieCount{MovieID: &movieID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	tmdbMovie.ViewsCount = *movieCount.ViewsCount

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tmdbMovie)
}

func GetMoviesByTitleAndOverview(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	overview := r.URL.Query().Get("overview")

	if title == "" && overview == "" {
		http.Error(w, "Provide at least one search parameter (title or overview)", http.StatusBadRequest)
		return
	}

	movies, err := services.GetMoviesByTitleAndOverview(title, overview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}
