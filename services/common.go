package services

import (
	"errors"

	"github.com/labora/labora-golang/cinema_labora/config"
)

func CreateDatabaseItem(newItem interface{}, insertQuery string, fields []interface{}, requiredFields []interface{}, requiredFieldsMsg string) (int, error) {
	for _, field := range requiredFields {
		if field == nil {
			return -1, errors.New(requiredFieldsMsg)
		}
	}

	stmt, err := config.DbConnection.Prepare(insertQuery)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	var newItemID int
	err = stmt.QueryRow(fields...).Scan(&newItemID)
	if err != nil {
		return -1, err
	}

	return newItemID, nil
}
