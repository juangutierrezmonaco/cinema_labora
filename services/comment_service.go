package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/labora/labora-golang/cinema_labora/config"
	"github.com/labora/labora-golang/cinema_labora/models"
)

func CreateComment(newComment models.Comment) (int, error) {
	insertQuery := "INSERT INTO comment(user_id, movie_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	currentTime := time.Now().Unix()
	fields := []interface{}{newComment.UserID, newComment.MovieID, newComment.Content, currentTime, currentTime}
	requiredFields := []interface{}{newComment.UserID, newComment.MovieID, newComment.Content}
	requiredFieldMsg := "User ID, movie ID, and content are required fields"
	return CreateDatabaseItem(newComment, insertQuery, fields, requiredFields, requiredFieldMsg)
}

func buildSearchCommentQuery(userID, movieID int, content string) string {
	query := "SELECT * FROM comment"
	if userID == 0 && movieID == 0 && content == "" {
		return query
	}

	var subqueries []string
	if userID != 0 {
		subqueries = append(subqueries, fmt.Sprintf("user_id = %d", userID))
	}
	if movieID != 0 {
		subqueries = append(subqueries, fmt.Sprintf("movie_id = %d", movieID))
	}
	if content != "" {
		subqueries = append(subqueries, fmt.Sprintf("content ILIKE '%%%s%%'", content))
	}

	if len(subqueries) > 0 {
		query += " WHERE " + strings.Join(subqueries, " AND ")
	}
	return query
}

func GetComments(userID, movieID int, content string) ([]models.Comment, error) {
	query := buildSearchCommentQuery(userID, movieID, content)
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

	var comments []models.Comment

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.MovieID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func GetCommentByID(id int) (*models.Comment, error) {
	stmt, err := config.DbConnection.Prepare("SELECT * FROM comment WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var comment models.Comment
	err = stmt.QueryRow(id).Scan(&comment.ID, &comment.UserID, &comment.MovieID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func buildUpdateCommentQuery(updatedComment models.Comment) (string, error) {
	query := "UPDATE comment SET "
	paramsSize := 1
	if updatedComment.UserID != nil {
		query += fmt.Sprintf("user_id = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedComment.MovieID != nil {
		query += fmt.Sprintf("movie_id = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedComment.Content != nil {
		query += fmt.Sprintf("content = $%d, ", paramsSize)
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

func getNonNullCommentFields(comment models.Comment) []interface{} {
	var nonNullFields []interface{}

	if comment.UserID != nil {
		nonNullFields = append(nonNullFields, comment.UserID)
	}
	if comment.MovieID != nil {
		nonNullFields = append(nonNullFields, comment.MovieID)
	}
	if comment.Content != nil {
		nonNullFields = append(nonNullFields, comment.Content)
	}
	nonNullFields = append(nonNullFields, time.Now().Unix())

	return nonNullFields
}

func UpdateComment(id int, updatedComment models.Comment) error {
	query, err := buildUpdateCommentQuery(updatedComment)
	if err != nil {
		return err
	}

	stmt, err := config.DbConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	fields := getNonNullCommentFields(updatedComment)
	fields = append(fields, id)
	res, err := stmt.Exec(fields...)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("There's no comment with the ID %d", id)
	}

	return nil
}

func DeleteComment(id int) error {
	stmt, err := config.DbConnection.Prepare("DELETE FROM comment WHERE id = $1")
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
		return fmt.Errorf("There's no comment with the ID %d", id)
	}

	return nil
}
