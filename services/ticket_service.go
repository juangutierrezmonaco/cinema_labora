package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/labora/labora-golang/cinema_labora/config"
	"github.com/labora/labora-golang/cinema_labora/models"
)

func CreateTicket(newTicket models.Ticket) (int, error) {
	insertQuery := "INSERT INTO ticket(pickup_id, user_id, screening_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	currentTime := time.Now().Unix()
	fields := []interface{}{newTicket.PickupID, newTicket.UserID, newTicket.ScreeningID, currentTime}
	requiredFields := []interface{}{newTicket.PickupID, newTicket.UserID, newTicket.ScreeningID}
	requiredFieldMsg := "Pickup ID, user ID, and screening ID are required fields"
	return CreateDatabaseItem(newTicket, insertQuery, fields, requiredFields, requiredFieldMsg)
}

func buildSearchTicketQuery(pickupID string, userID, screeningID int) string {
	qb := NewQueryBuilder("ticket")
	if pickupID != "" {
		qb.AddCondition(fmt.Sprintf("pickup_id = '%s'", pickupID))
	}
	if userID != 0 {
		qb.AddCondition(fmt.Sprintf("user_id = %d", userID))
	}
	if screeningID != 0 {
		qb.AddCondition(fmt.Sprintf("screening_id = %d", screeningID))
	}
	return qb.BuildQuery()
}

func GetTickets(pickupID string, userID, screeningID int) ([]models.Ticket, error) {
	query := buildSearchTicketQuery(pickupID, userID, screeningID)
	scanRowFunc := func(rows *sql.Rows) (interface{}, error) {
		var ticket models.Ticket
		err := rows.Scan(&ticket.ID, &ticket.PickupID, &ticket.UserID, &ticket.ScreeningID, &ticket.CreatedAt)
		if err != nil {
			return nil, err
		}
		return ticket, nil
	}

	items, err := GetDatabaseItems(query, scanRowFunc)
	if err != nil {
		return nil, err
	}

	var tickets []models.Ticket
	for _, item := range items {
		tickets = append(tickets, item.(models.Ticket))
	}

	return tickets, nil
}

func GetTicketByID(id int) (*models.Ticket, error) {
	scanRowFunc := func(row *sql.Row) (interface{}, error) {
		var ticket models.Ticket
		err := row.Scan(&ticket.ID, &ticket.PickupID, &ticket.UserID, &ticket.ScreeningID, &ticket.CreatedAt)
		if err != nil {
			return nil, err
		}

		return &ticket, nil
	}

	item, err := GetDatabaseItemByID(id, "ticket", scanRowFunc)
	if err != nil {
		return nil, err
	}

	return item.(*models.Ticket), nil
}

func buildUpdateTicketQuery(updatedTicket models.Ticket) (string, error) {
	query := "UPDATE ticket SET "
	paramsSize := 1
	if updatedTicket.PickupID != nil {
		query += fmt.Sprintf("pickup_id = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedTicket.UserID != nil {
		query += fmt.Sprintf("user_id = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedTicket.ScreeningID != nil {
		query += fmt.Sprintf("screening_id = $%d, ", paramsSize)
		paramsSize++
	}
	if paramsSize == 1 {
		return "", fmt.Errorf("You must modify at least one field.")
	}

	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d", paramsSize)
	return query, nil
}

func getNonNullTicketFields(ticket models.Ticket) []interface{} {
	var nonNullFields []interface{}

	if ticket.PickupID != nil {
		nonNullFields = append(nonNullFields, ticket.PickupID)
	}
	if ticket.UserID != nil {
		nonNullFields = append(nonNullFields, ticket.UserID)
	}
	if ticket.ScreeningID != nil {
		nonNullFields = append(nonNullFields, ticket.ScreeningID)
	}

	return nonNullFields
}

func UpdateTicket(id int, updatedTicket models.Ticket) error {
	query, err := buildUpdateTicketQuery(updatedTicket)
	if err != nil {
		return err
	}

	stmt, err := config.DbConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	fields := getNonNullTicketFields(updatedTicket)
	fields = append(fields, id)
	res, err := stmt.Exec(fields...)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("There's no ticket with the ID %d", id)
	}

	return nil
}

func DeleteTicket(id int) error {
	stmt, err := config.DbConnection.Prepare("DELETE FROM ticket WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("There's no ticket with the ID %d", id)
	}

	return nil
}
