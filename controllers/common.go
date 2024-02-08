package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
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
