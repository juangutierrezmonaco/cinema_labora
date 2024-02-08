package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateControllerItem(w http.ResponseWriter, r *http.Request, newItem interface{}, createFunc func(interface{}) (int, error), itemName string) {
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemID, err := createFunc(newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error while creating the %s.", itemName)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := fmt.Sprintf("%s created correctly with ID: %d\n", itemName, itemID)
	w.Write([]byte(response))
}

func GetControllerItems(w http.ResponseWriter, r *http.Request, getFunc func() (interface{}, error), itemName string) {
	items, err := getFunc()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error fetching the %s.", itemName)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func GetControllerItemByID(w http.ResponseWriter, r *http.Request, getFunc func(int) (interface{}, error), itemName string) {
	params := mux.Vars(r)
	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid %s ID", itemName), http.StatusBadRequest)
		return
	}

	item, err := getFunc(itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func UpdateControllerItem(w http.ResponseWriter, r *http.Request, updatedItem interface{}, updateFunc func(int, interface{}) error, itemName string) {
	params := mux.Vars(r)
	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid %s ID", itemName), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = updateFunc(itemID, updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("%s with ID %d updated successfully.", itemName, itemID)
	w.Write([]byte(response))
}

func DeleteControllerItem(w http.ResponseWriter, r *http.Request, deleteFunc func(int) error, itemName string) {
	params := mux.Vars(r)
	itemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid %s ID", itemName), http.StatusBadRequest)
		return
	}

	err = deleteFunc(itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("%s with ID %d deleted successfully.", itemName, itemID)
	w.Write([]byte(response))
}
