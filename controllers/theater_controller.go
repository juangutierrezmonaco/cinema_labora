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

	err := json.NewDecoder(r.Body).Decode(&newTheater)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	theaterID, err := services.CreateTheater(newTheater)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error while creating the theater."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := fmt.Sprintf("Theater created correcly with ID: %d\n", theaterID)
	w.Write([]byte(response))
}

func GetTheaters(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	capacity, err := strconv.Atoi(r.URL.Query().Get("capacity"))
	capacityGt, err := strconv.Atoi(r.URL.Query().Get("capacity_gt"))
	capacityLt, err := strconv.Atoi(r.URL.Query().Get("capacity_lt"))

	theaters, err := services.GetTheaters(name, capacity, capacityGt, capacityLt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error fetching the theaters."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(theaters)
}

func GetTheaterByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	theaterID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid theater ID", http.StatusBadRequest)
		return
	}

	theater, err := services.GetTheaterByID(theaterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(theater)
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
