package controllers

import (
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/services"
	"net/http"
	"strconv"
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

	GetControllerItems(w, r, func() (interface{}, error) {
		return services.GetTickets(pickupID, userID, screeningID)
	}, "Tickets")
}

func GetTicketByID(w http.ResponseWriter, r *http.Request) {
	GetControllerItemByID(w, r, func(id int) (interface{}, error) {
		return services.GetTicketByID(id)
	}, "Ticket")
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	var updatedTicket models.Ticket
	UpdateControllerItem(w, r, &updatedTicket, func(id int, data interface{}) error {
		ticket := data.(*models.Ticket)
		return services.UpdateTicket(id, *ticket)
	}, "Ticket")
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	DeleteControllerItem(w, r, func(id int) error {
		return services.DeleteTicket(id)
	}, "Ticket")
}
