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

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var newTicket models.Ticket
	CreateControllerItem(w, r, &newTicket, func(data interface{}) (int, error) {
		ticket := data.(*models.Ticket)
		return services.CreateTicket(*ticket)
	}, "Ticket")
}

func GetTickets(w http.ResponseWriter, r *http.Request) {
	pickupID := r.URL.Query().Get("pickup_id")
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	screeningID, _ := strconv.Atoi(r.URL.Query().Get("screening_id"))

	tickets, err := services.GetTickets(pickupID, userID, screeningID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Error fetching the tickets."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tickets)
}

func GetTicketByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	ticket, err := services.GetTicketByID(ticketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ticket)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	var updatedTicket models.Ticket
	err = json.NewDecoder(r.Body).Decode(&updatedTicket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.UpdateTicket(ticketID, updatedTicket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Ticket with ID %d updated successfully.", ticketID)
	w.Write([]byte(response))
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ticketID, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteTicket(ticketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Ticket with ID %d deleted successfully.", ticketID)
	w.Write([]byte(response))
}
