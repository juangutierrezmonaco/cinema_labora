package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/labora/labora-golang/cinema_labora/config"
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/lib/pq"
)

func CreateScreening(newScreening models.Screening) (int, error) {
	insertQuery := "INSERT INTO screening(name, movie_id, theater_id, available_seats, taken_seats, showtime, price, language, views_count, created_at, updated_at)	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	currentTime := time.Now().Unix()
	fields := []interface{}{newScreening.Name, newScreening.MovieID, newScreening.TheaterID, newScreening.AvailableSeats, pq.Array(newScreening.TakenSeats), newScreening.Showtime, newScreening.Price, newScreening.Language, newScreening.ViewsCount, currentTime, currentTime}
	requiredFields := []interface{}{newScreening.Name, newScreening.MovieID, newScreening.TheaterID, newScreening.AvailableSeats, newScreening.Showtime, newScreening.Price}
	requiredFieldMsg := "Name, movie_id, theater_id, available_seats, showtime, and price are required fields"
	return CreateDatabaseItem(newScreening, insertQuery, fields, requiredFields, requiredFieldMsg)
}

func buildSearchScreeningQuery(name string, showtime int64, showtimeGt int64, showtimeLt int64, price float64, priceGt float64, priceLt float64, language string, viewsCount int, viewsCountGt int, viewsCountLt int) string {
	query := "SELECT * FROM screening"
	if name == "" && showtime == 0 && showtimeGt == 0 && showtimeLt == 0 && price == 0 && priceGt == 0 && priceLt == 0 && language == "" && viewsCount == 0 && viewsCountGt == 0 && viewsCountLt == 0 {
		return query
	}
	var subqueries []string
	if name != "" {
		subqueries = append(subqueries, fmt.Sprintf("name ILIKE '%%%s%%'", name))
	}
	if showtime != 0 {
		subqueries = append(subqueries, fmt.Sprintf("showtime = %d", showtime))
	}
	if showtimeGt != 0 {
		subqueries = append(subqueries, fmt.Sprintf("showtime > %d", showtimeGt))
	}
	if showtimeLt != 0 {
		subqueries = append(subqueries, fmt.Sprintf("showtime < %d", showtimeLt))
	}
	if price != 0 {
		subqueries = append(subqueries, fmt.Sprintf("price = %f", price))
	}
	if priceGt != 0 {
		subqueries = append(subqueries, fmt.Sprintf("price > %f", priceGt))
	}
	if priceLt != 0 {
		subqueries = append(subqueries, fmt.Sprintf("price < %f", priceLt))
	}
	if language != "" {
		subqueries = append(subqueries, fmt.Sprintf("language = '%s'", language))
	}
	if viewsCount != 0 {
		subqueries = append(subqueries, fmt.Sprintf("views_count = %d", viewsCount))
	}
	if viewsCountGt != 0 {
		subqueries = append(subqueries, fmt.Sprintf("views_count > %d", viewsCountGt))
	}
	if viewsCountLt != 0 {
		subqueries = append(subqueries, fmt.Sprintf("views_count < %d", viewsCountLt))
	}
	if len(subqueries) > 0 {
		query += " WHERE " + strings.Join(subqueries, " AND ")
	}
	return query
}

func convertUint8ToStringArray(bytes []uint8) []string {
	str := strings.Trim(string(bytes), "{}")
	return strings.Split(str, ",")
}

func GetScreenings(name string, showtime int64, showtimeGt int64, showtimeLt int64, price float64, priceGt float64, priceLt float64, language string, viewsCount int, viewsCountGt int, viewsCountLt int) ([]models.Screening, error) {
	query := buildSearchScreeningQuery(name, showtime, showtimeGt, showtimeLt, price, priceGt, priceLt, language, viewsCount, viewsCountGt, viewsCountLt)
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

	var screenings []models.Screening
	for rows.Next() {
		var screening models.Screening
		var auxTakenSeats []uint8
		err := rows.Scan(
			&screening.ID, &screening.Name, &screening.MovieID,
			&screening.TheaterID, &screening.AvailableSeats, &auxTakenSeats,
			&screening.Showtime, &screening.Price, &screening.Language,
			&screening.ViewsCount, &screening.CreatedAt, &screening.UpdatedAt,
		)
		screening.TakenSeats = convertUint8ToStringArray(auxTakenSeats)
		if err != nil {
			return nil, err
		}
		screenings = append(screenings, screening)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return screenings, nil
}

func GetScreeningByID(id int) (*models.Screening, error) {
	stmt, err := config.DbConnection.Prepare("SELECT * FROM screening WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var screening models.Screening
	var auxTakenSeats []uint8
	err = stmt.QueryRow(id).Scan(
		&screening.ID, &screening.Name, &screening.MovieID,
		&screening.TheaterID, &screening.AvailableSeats, &auxTakenSeats,
		&screening.Showtime, &screening.Price, &screening.Language,
		&screening.ViewsCount, &screening.CreatedAt, &screening.UpdatedAt,
	)
	screening.TakenSeats = convertUint8ToStringArray(auxTakenSeats)
	if err != nil {
		return nil, err
	}

	return &screening, nil
}

func GetScreeningByMovieIdOrTheaterId(id int, isSearchingByMovie bool) (*models.Screening, error) {
	var query string
	if isSearchingByMovie {
		query = "SELECT * FROM screening WHERE movie_id = $1"
	} else {
		query = "SELECT * FROM screening WHERE theater_id = $1"
	}
	stmt, err := config.DbConnection.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var screening models.Screening
	var auxTakenSeats []uint8
	err = stmt.QueryRow(id).Scan(
		&screening.ID, &screening.Name, &screening.MovieID,
		&screening.TheaterID, &screening.AvailableSeats, &auxTakenSeats,
		&screening.Showtime, &screening.Price, &screening.Language,
		&screening.ViewsCount, &screening.CreatedAt, &screening.UpdatedAt,
	)
	screening.TakenSeats = convertUint8ToStringArray(auxTakenSeats)
	if err != nil {
		return nil, err
	}

	return &screening, nil
}

func buildUpdateScreeningQuery(updatedScreening models.Screening) (string, error) {
	query := "UPDATE screening SET "
	paramsSize := 1
	if updatedScreening.Name != nil {
		query += fmt.Sprintf("name = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedScreening.MovieID != nil {
		query += fmt.Sprintf("movie_id = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedScreening.TheaterID != nil {
		query += fmt.Sprintf("theater_id = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedScreening.AvailableSeats != nil {
		query += fmt.Sprintf("available_seats = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedScreening.TakenSeats != nil {
		query += fmt.Sprintf("taken_seats = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedScreening.Showtime != nil {
		query += fmt.Sprintf("showtime = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedScreening.Price != nil {
		query += fmt.Sprintf("price = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedScreening.Language != nil {
		query += fmt.Sprintf("language = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedScreening.ViewsCount != nil {
		query += fmt.Sprintf("views_count = $%d, ", paramsSize)
		paramsSize++
	}
	if paramsSize == 1 {
		return "", fmt.Errorf("You must modify at least one field.")
	}

	query += fmt.Sprintf("updated_at = $%d, ", paramsSize)
	paramsSize++
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d", paramsSize)
	return query, nil
}

func getNonNullScreeningFields(screening models.Screening) []interface{} {
	var nonNullFields []interface{}

	if screening.Name != nil {
		nonNullFields = append(nonNullFields, screening.Name)
	}
	if screening.MovieID != nil {
		nonNullFields = append(nonNullFields, screening.MovieID)
	}
	if screening.TheaterID != nil {
		nonNullFields = append(nonNullFields, screening.TheaterID)
	}
	if screening.AvailableSeats != nil {
		nonNullFields = append(nonNullFields, screening.AvailableSeats)
	}
	if screening.TakenSeats != nil {
		nonNullFields = append(nonNullFields, pq.Array(screening.TakenSeats))
	}
	if screening.Showtime != nil {
		nonNullFields = append(nonNullFields, screening.Showtime)
	}
	if screening.Price != nil {
		nonNullFields = append(nonNullFields, screening.Price)
	}
	if screening.Language != nil {
		nonNullFields = append(nonNullFields, screening.Language)
	}
	if screening.ViewsCount != nil {
		nonNullFields = append(nonNullFields, screening.ViewsCount)
	}
	nonNullFields = append(nonNullFields, time.Now().Unix())

	return nonNullFields
}

func UpdateScreening(id int, updatedScreening models.Screening) error {
	query, err := buildUpdateScreeningQuery(updatedScreening)
	if err != nil {
		return err
	}

	stmt, err := config.DbConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	fields := getNonNullScreeningFields(updatedScreening)
	fields = append(fields, id)
	res, err := stmt.Exec(fields...)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("There's no screening with the ID %d", id)
	}

	return nil
}

func DeleteScreening(id int) error {
	stmt, err := config.DbConnection.Prepare("DELETE FROM screening WHERE id = $1")
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
		return fmt.Errorf("There's no screening with the ID %d", id)
	}

	return nil
}
