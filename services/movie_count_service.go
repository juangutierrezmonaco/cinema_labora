package services

import (
	"errors"
	"time"

	"github.com/labora/labora-golang/cinema_labora/config"
	"github.com/labora/labora-golang/cinema_labora/models"
)

func CreateMovieCount(newMovieCount models.MovieCount) (int, error) {
	if newMovieCount.MovieID == nil {
		return -1, errors.New("Movie ID is a required field")
	}

	stmt, err := config.DbConnection.Prepare("INSERT INTO movie_count(movie_id, views_count, updated_at) VALUES ($1, $2, $3) RETURNING views_count")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	var movieViewsCount int
	currentTime := time.Now().Unix()
	if newMovieCount.ViewsCount == nil {
		newMovieCount.ViewsCount = new(int)
		*newMovieCount.ViewsCount = 1
	}
	err = stmt.QueryRow(newMovieCount.MovieID, newMovieCount.ViewsCount, currentTime).Scan(&movieViewsCount)
	if err != nil {
		return -1, err
	}

	return movieViewsCount, nil
}

func GetMoviesCounts() ([]models.MovieCount, error) {
	stmt, err := config.DbConnection.Prepare("SELECT * FROM movie_count")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moviesCounts []models.MovieCount
	for rows.Next() {
		var movieCount models.MovieCount
		err := rows.Scan(
			&movieCount.MovieID, &movieCount.ViewsCount, &movieCount.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		err = UpdateMovieCount(movieCount)
		if err != nil {
			return nil, err
		}

		moviesCounts = append(moviesCounts, movieCount)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return moviesCounts, nil
}

func GetMovieCountByID(movieID int) (*models.MovieCount, error) {
	stmt, err := config.DbConnection.Prepare("SELECT * FROM movie_count WHERE movie_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var movieCount models.MovieCount
	err = stmt.QueryRow(movieID).Scan(&movieCount.MovieID, &movieCount.ViewsCount, &movieCount.UpdatedAt)
	if err != nil {
		return nil, err
	}

	err = UpdateMovieCount(movieCount)
	if err != nil {
		return nil, err
	}

	return &movieCount, nil
}

func UpdateMovieCount(movieCount models.MovieCount) error {
	stmt, err := config.DbConnection.Prepare("UPDATE movie_count SET views_count = $1, updated_at = $2 WHERE movie_id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()
	*movieCount.ViewsCount++
	_, err = stmt.Exec(movieCount.ViewsCount, time.Now().Unix(), movieCount.MovieID)
	if err != nil {
		return err
	}

	return nil
}
