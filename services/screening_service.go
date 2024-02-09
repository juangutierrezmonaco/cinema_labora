package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/labora/labora-golang/cinema_labora/config"
	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/util"
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
	qb := NewQueryBuilder("screening")
	if name != "" {
		qb.AddCondition(fmt.Sprintf("name ILIKE '%%%s%%'", name))
	}
	if showtime != 0 {
		qb.AddCondition(fmt.Sprintf("showtime = %d", showtime))
	}
	if showtimeGt != 0 {
		qb.AddCondition(fmt.Sprintf("showtime > %d", showtimeGt))
	}
	if showtimeLt != 0 {
		qb.AddCondition(fmt.Sprintf("showtime < %d", showtimeLt))
	}
	if price != 0 {
		qb.AddCondition(fmt.Sprintf("price = %f", price))
	}
	if priceGt != 0 {
		qb.AddCondition(fmt.Sprintf("price > %f", priceGt))
	}
	if priceLt != 0 {
		qb.AddCondition(fmt.Sprintf("price < %f", priceLt))
	}
	if language != "" {
		qb.AddCondition(fmt.Sprintf("language = '%s'", language))
	}
	if viewsCount != 0 {
		qb.AddCondition(fmt.Sprintf("views_count = %d", viewsCount))
	}
	if viewsCountGt != 0 {
		qb.AddCondition(fmt.Sprintf("views_count > %d", viewsCountGt))
	}
	if viewsCountLt != 0 {
		qb.AddCondition(fmt.Sprintf("views_count < %d", viewsCountLt))
	}
	return qb.BuildQuery()
}

func GetScreenings(name string, showtime int64, showtimeGt int64, showtimeLt int64, price float64, priceGt float64, priceLt float64, language string, viewsCount int, viewsCountGt int, viewsCountLt int) ([]models.Screening, error) {
	query := buildSearchScreeningQuery(name, showtime, showtimeGt, showtimeLt, price, priceGt, priceLt, language, viewsCount, viewsCountGt, viewsCountLt)
	scanRowFunc := func(rows *sql.Rows) (interface{}, error) {
		var screening models.Screening
		var auxTakenSeats []uint8
		err := rows.Scan(
			&screening.ID, &screening.Name, &screening.MovieID,
			&screening.TheaterID, &screening.AvailableSeats, &auxTakenSeats,
			&screening.Showtime, &screening.Price, &screening.Language,
			&screening.ViewsCount, &screening.CreatedAt, &screening.UpdatedAt,
		)
		screening.TakenSeats = util.ConvertSqlUint8ToStringArray(auxTakenSeats)
		if err != nil {
			return nil, err
		}
		return screening, nil
	}

	items, err := GetDatabaseItems(query, scanRowFunc)
	if err != nil {
		return nil, err
	}

	var screenings []models.Screening
	for _, item := range items {
		screenings = append(screenings, item.(models.Screening))
	}

	return screenings, nil
}

func GetScreeningByID(id int) (*models.Screening, error) {
	scanRowFunc := func(row *sql.Row) (interface{}, error) {
		var screening models.Screening
		var auxTakenSeats []uint8
		err := row.Scan(
			&screening.ID, &screening.Name, &screening.MovieID,
			&screening.TheaterID, &screening.AvailableSeats, &auxTakenSeats,
			&screening.Showtime, &screening.Price, &screening.Language,
			&screening.ViewsCount, &screening.CreatedAt, &screening.UpdatedAt,
		)
		screening.TakenSeats = util.ConvertSqlUint8ToStringArray(auxTakenSeats)
		if err != nil {
			return nil, err
		}

		return &screening, nil
	}

	item, err := GetDatabaseItemByID(id, "screening", scanRowFunc)
	if err != nil {
		return nil, err
	}

	return item.(*models.Screening), nil
}

func GetScreeningByMovieIdOrTheaterId(id int, isSearchingByMovie bool) ([]models.Screening, error) {
	var query string
	if isSearchingByMovie {
		query = fmt.Sprintf("SELECT * FROM screening WHERE movie_id = %d", id)
	} else {
		query = fmt.Sprintf("SELECT * FROM screening WHERE theater_id = %d", id)
	}
	scanRowFunc := func(rows *sql.Rows) (interface{}, error) {
		var screening models.Screening
		var auxTakenSeats []uint8
		err := rows.Scan(
			&screening.ID, &screening.Name, &screening.MovieID,
			&screening.TheaterID, &screening.AvailableSeats, &auxTakenSeats,
			&screening.Showtime, &screening.Price, &screening.Language,
			&screening.ViewsCount, &screening.CreatedAt, &screening.UpdatedAt,
		)
		screening.TakenSeats = util.ConvertSqlUint8ToStringArray(auxTakenSeats)
		if err != nil {
			return nil, err
		}
		return screening, nil
	}

	items, err := GetDatabaseItems(query, scanRowFunc)
	if err != nil {
		return nil, err
	}

	var screenings []models.Screening
	for _, item := range items {
		screenings = append(screenings, item.(models.Screening))
	}

	return screenings, nil
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
	return DeleteItemByID(id, "screening")
}
