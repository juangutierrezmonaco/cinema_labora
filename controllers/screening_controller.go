package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/services"
)

func CreateScreening(w http.ResponseWriter, r *http.Request) {
	var newScreening models.Screening

	err := json.NewDecoder(r.Body).Decode(&newScreening)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	screeningID, err := services.CreateScreening(newScreening)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error while creating the screening."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := fmt.Sprintf("Screening created correctly with ID: %d\n", screeningID)
	w.Write([]byte(response))
}

func GetScreenings(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	showtime, err := strconv.ParseInt(r.URL.Query().Get("showtime"), 10, 64)
	showtimeGt, err := strconv.ParseInt(r.URL.Query().Get("showtime_gt"), 10, 64)
	showtimeLt, err := strconv.ParseInt(r.URL.Query().Get("showtime_lt"), 10, 64)
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
	priceGt, err := strconv.ParseFloat(r.URL.Query().Get("price_gt"), 64)
	priceLt, err := strconv.ParseFloat(r.URL.Query().Get("price_lt"), 64)
	language := r.URL.Query().Get("language")
	viewsCount, err := strconv.Atoi(r.URL.Query().Get("views_count"))
	viewsCountGt, err := strconv.Atoi(r.URL.Query().Get("views_count_gt"))
	viewsCountLt, err := strconv.Atoi(r.URL.Query().Get("views_count_lt"))

	screenings, err := services.GetScreenings(name, showtime, showtimeGt, showtimeLt, price, priceGt, priceLt, language, viewsCount, viewsCountGt, viewsCountLt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error fetching the screenings."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(screenings)
}

func GetScreeningByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	screeningID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid screening ID", http.StatusBadRequest)
		return
	}

	screening, err := services.GetScreeningByID(screeningID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(screening)
}

func GetScreeningByMovieIdOrTheaterId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid screening ID", http.StatusBadRequest)
		return
	}

	isSearchingByMovie := strings.Contains(r.URL.String(), "movie")
	screening, err := services.GetScreeningByMovieIdOrTheaterId(id, isSearchingByMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(screening)
}

func UpdateScreening(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	screeningID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid screening ID", http.StatusBadRequest)
		return
	}

	var updatedScreening models.Screening
	err = json.NewDecoder(r.Body).Decode(&updatedScreening)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.UpdateScreening(screeningID, updatedScreening)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Screening with ID %d updated successfully.", screeningID)
	w.Write([]byte(response))
}

func DeleteScreening(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	screeningID, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid screening ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteScreening(screeningID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Screening with ID %d deleted successfully.", screeningID)
	w.Write([]byte(response))
}
