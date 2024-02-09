package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/labora/labora-golang/cinema_labora/config"
	"github.com/labora/labora-golang/cinema_labora/models"
)

func CreateTheater(newTheater models.Theater) (int, error) {
	insertQuery := "INSERT INTO theater(name, capacity, last_row, last_column, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	currentTime := time.Now().Unix()
	fields := []interface{}{newTheater.Name, newTheater.Capacity, newTheater.LastRow, newTheater.LastColumn, currentTime}
	requiredFields := []interface{}{newTheater.Name, newTheater.Capacity, newTheater.LastRow, newTheater.LastColumn}
	requiredFieldMsg := "Name, capacity, last row and last column are required fields"
	return CreateDatabaseItem(newTheater, insertQuery, fields, requiredFields, requiredFieldMsg)
}

func buildSearchTheaterQuery(name string, capacity int, capacityGt int, capacityLt int) string {
	qb := NewQueryBuilder("theater")
	if name != "" {
		qb.AddCondition(fmt.Sprintf("name ILIKE '%%%s%%'", name))
	}
	if capacity > 0 {
		qb.AddCondition(fmt.Sprintf("capacity = %d", capacity))
	}
	if capacityGt > 0 {
		qb.AddCondition(fmt.Sprintf("capacity > %d", capacityGt))
	}
	if capacityLt > 0 {
		qb.AddCondition(fmt.Sprintf("capacity < %d", capacityLt))
	}
	return qb.BuildQuery()
}

func GetTheaters(name string, capacity int, capacityGt int, capacityLt int) ([]models.Theater, error) {
	query := buildSearchTheaterQuery(name, capacity, capacityGt, capacityLt)
	scanRowFunc := func(rows *sql.Rows) (interface{}, error) {
		var theater models.Theater
		err := rows.Scan(&theater.ID, &theater.Name, &theater.Capacity, &theater.LastRow, &theater.LastColumn, &theater.CreatedAt)
		if err != nil {
			return nil, err
		}
		return theater, nil
	}

	items, err := GetDatabaseItems(query, scanRowFunc)
	if err != nil {
		return nil, err
	}

	var theaters []models.Theater
	for _, item := range items {
		theaters = append(theaters, item.(models.Theater))
	}

	return theaters, nil
}

func GetTheaterByID(id int) (*models.Theater, error) {
	scanRowFunc := func(row *sql.Row) (interface{}, error) {
		var theater models.Theater
		err := row.Scan(&theater.ID, &theater.Name, &theater.Capacity, &theater.LastRow, &theater.LastColumn, &theater.CreatedAt)
		if err != nil {
			return nil, err
		}

		return &theater, nil
	}

	item, err := GetDatabaseItemByID(id, "theater", scanRowFunc)
	if err != nil {
		return nil, err
	}

	return item.(*models.Theater), nil
}

func buildUpdateTheaterQuery(updatedTheater models.Theater) (string, error) {
	query := "UPDATE theater SET "
	paramsSize := 1
	if updatedTheater.Name != nil {
		query += fmt.Sprintf("name = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedTheater.Capacity != nil {
		query += fmt.Sprintf("capacity = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedTheater.LastRow != nil {
		query += fmt.Sprintf("last_row = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedTheater.LastColumn != nil {
		query += fmt.Sprintf("last_column = $%d, ", paramsSize)
		paramsSize++
	}
	if paramsSize == 1 {
		return "", fmt.Errorf("You must modify at least one field.")
	}

	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d", paramsSize)
	return query, nil
}

func getNonNullTheaterFields(theater models.Theater) []interface{} {
	var nonNullFields []interface{}

	if theater.Name != nil {
		nonNullFields = append(nonNullFields, theater.Name)
	}
	if theater.Capacity != nil {
		nonNullFields = append(nonNullFields, theater.Capacity)
	}
	if theater.LastRow != nil {
		nonNullFields = append(nonNullFields, theater.LastRow)
	}
	if theater.LastColumn != nil {
		nonNullFields = append(nonNullFields, theater.LastColumn)
	}

	return nonNullFields
}

func UpdateTheater(id int, updatedTheater models.Theater) error {
	query, err := buildUpdateTheaterQuery(updatedTheater)
	if err != nil {
		return err
	}

	stmt, err := config.DbConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	fields := getNonNullTheaterFields(updatedTheater)
	fields = append(fields, id)
	res, err := stmt.Exec(fields...)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("There's no theater with the ID %d", id)
	}

	return nil
}

func DeleteTheater(id int) error {
	stmt, err := config.DbConnection.Prepare("DELETE FROM theater WHERE id = $1")
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
		return fmt.Errorf("There's no theater with the ID %d", id)
	}

	return nil
}
