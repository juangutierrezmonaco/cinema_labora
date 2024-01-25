package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/labora/labora-golang/cinema_labora/config"
	"github.com/labora/labora-golang/cinema_labora/models"
)

func CreateTheater(newTheater models.Theater) (int, error) {
	if newTheater.Name == nil || newTheater.Capacity == nil || newTheater.LastRow == nil || newTheater.LastColumn == nil {
		return -1, errors.New("Name, capacity, last row and last column are required fields")
	}

	stmt, err := config.DbConnection.Prepare("INSERT INTO theater(name, capacity, last_row, last_column, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	var newTheaterID int
	currentTime := time.Now().Unix()
	err = stmt.QueryRow(newTheater.Name, newTheater.Capacity, newTheater.LastRow, newTheater.LastColumn, currentTime).Scan(&newTheaterID)
	if err != nil {
		return -1, err
	}

	return newTheaterID, nil
}

func buildSearchTheaterQuery(name string, capacity int, capacityGt int, capacityLt int) string {
	query := "SELECT * FROM theater"
	if name == "" && capacity == 0 && capacityGt == 0 && capacityLt == 0 {
		return query
	}
	var subqueries []string
	if name != "" {
		subqueries = append(subqueries, fmt.Sprintf("name ILIKE '%%%s%%'", name))
	}
	if capacity > 0 {
		subqueries = append(subqueries, fmt.Sprintf("capacity = %d", capacity))
	}
	if capacityGt > 0 {
		subqueries = append(subqueries, fmt.Sprintf("capacity > %d", capacityGt))
	}
	if capacityLt > 0 {
		subqueries = append(subqueries, fmt.Sprintf("capacity < %d", capacityLt))
	}
	if len(subqueries) > 0 {
		query += " WHERE " + strings.Join(subqueries, " AND ")
	}
	return query
}

func GetTheaters(name string, capacity int, capacityGt int, capacityLt int) ([]models.Theater, error) {
	query := buildSearchTheaterQuery(name, capacity, capacityGt, capacityLt)
	stmt, err := config.DbConnection.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var theaters []models.Theater

	for rows.Next() {
		var theater models.Theater
		err := rows.Scan(&theater.ID, &theater.Name, &theater.Capacity, &theater.LastRow, &theater.LastColumn, &theater.CreatedAt)
		if err != nil {
			return nil, err
		}
		theaters = append(theaters, theater)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return theaters, nil
}

func GetTheaterByID(id int) (*models.Theater, error) {
	stmt, err := config.DbConnection.Prepare("SELECT * FROM theater WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var theater models.Theater
	err = stmt.QueryRow(id).Scan(&theater.ID, &theater.Name, &theater.Capacity, &theater.LastRow, &theater.LastColumn, &theater.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &theater, nil
}

func buildUpdateTheaterQuery(updatedTheater models.Theater) (string, error) {
	query := "UPDATE theater SET "
	paramsSize := 1
	if updatedTheater.ID != nil {
		query += fmt.Sprintf("id = $%d, ", paramsSize)
		paramsSize++
	}
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
	if updatedTheater.CreatedAt != nil {
		query += fmt.Sprintf("created_at = $%d, ", paramsSize)
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

	if theater.ID != nil {
		nonNullFields = append(nonNullFields, theater.ID)
	}
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

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
