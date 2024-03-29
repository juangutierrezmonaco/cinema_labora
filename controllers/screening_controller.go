package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/services"
)

func CreateScreening(w http.ResponseWriter, r *http.Request) {
	var newScreening models.Screening
	CreateControllerItem(w, r, &newScreening, func(data interface{}) (int, error) {
		screening := data.(*models.Screening)
		return services.CreateScreening(*screening)
	}, "Screening")
}

func GetScreenings(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	showtime, _ := strconv.ParseInt(r.URL.Query().Get("showtime"), 10, 64)
	showtimeGt, _ := strconv.ParseInt(r.URL.Query().Get("showtime_gt"), 10, 64)
	showtimeLt, _ := strconv.ParseInt(r.URL.Query().Get("showtime_lt"), 10, 64)
	price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
	priceGt, _ := strconv.ParseFloat(r.URL.Query().Get("price_gt"), 64)
	priceLt, _ := strconv.ParseFloat(r.URL.Query().Get("price_lt"), 64)
	language := r.URL.Query().Get("language")
	viewsCount, _ := strconv.Atoi(r.URL.Query().Get("views_count"))
	viewsCountGt, _ := strconv.Atoi(r.URL.Query().Get("views_count_gt"))
	viewsCountLt, _ := strconv.Atoi(r.URL.Query().Get("views_count_lt"))

	GetControllerItems(w, r, func() (interface{}, error) {
		return services.GetScreenings(name, showtime, showtimeGt, showtimeLt, price, priceGt, priceLt, language, viewsCount, viewsCountGt, viewsCountLt)
	}, "Screenings")
}

func GetScreeningByID(w http.ResponseWriter, r *http.Request) {
	GetControllerItemByID(w, r, func(id int) (interface{}, error) {
		return services.GetScreeningByID(id)
	}, "Screening")
}

func GetScreeningByMovieIdOrTheaterId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	isSearchingByMovie := strings.Contains(r.URL.String(), "movie")
	screenings, err := services.GetScreeningByMovieIdOrTheaterId(id, isSearchingByMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(screenings)
}

func UpdateScreening(w http.ResponseWriter, r *http.Request) {
	var updatedScreening models.Screening
	UpdateControllerItem(w, r, &updatedScreening, func(id int, data interface{}) error {
		screening := data.(*models.Screening)
		return services.UpdateScreening(id, *screening)
	}, "Screening")
}

func DeleteScreening(w http.ResponseWriter, r *http.Request) {
	DeleteControllerItem(w, r, func(id int) error {
		return services.DeleteScreening(id)
	}, "Screening")
}
