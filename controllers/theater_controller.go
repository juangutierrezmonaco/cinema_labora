package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/services"
)

func CreateTheater(w http.ResponseWriter, r *http.Request) {
	var newTheater models.Theater
	CreateControllerItem(w, r, &newTheater, func(data interface{}) (int, error) {
		theater := data.(*models.Theater)
		return services.CreateTheater(*theater)
	}, "Theater")
}

func GetTheaters(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	capacity, _ := strconv.Atoi(r.URL.Query().Get("capacity"))
	capacityGt, _ := strconv.Atoi(r.URL.Query().Get("capacity_gt"))
	capacityLt, _ := strconv.Atoi(r.URL.Query().Get("capacity_lt"))

	GetControllerItems(w, r, func() (interface{}, error) {
		return services.GetTheaters(name, capacity, capacityGt, capacityLt)
	}, "Theaters")
}

func GetTheaterByID(w http.ResponseWriter, r *http.Request) {
	GetControllerItemByID(w, r, func(id int) (interface{}, error) {
		return services.GetTheaterByID(id)
	}, "Theater")
}

func UpdateTheater(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	theaterID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid theater ID", http.StatusBadRequest)
		return
	}

	var updatedTheater models.Theater
	err = json.NewDecoder(r.Body).Decode(&updatedTheater)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.UpdateTheater(theaterID, updatedTheater)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Theater with ID %d updated successfully.", theaterID)
	w.Write([]byte(response))
}

func DeleteTheater(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	theaterID, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid theater ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteTheater(theaterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Theater with ID %d deleted successfully.", theaterID)
	w.Write([]byte(response))
}
