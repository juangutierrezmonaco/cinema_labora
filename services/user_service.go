package services

import (
	"fmt"
	"time"

	"github.com/labora/labora-golang/cinema_labora/config"
	"github.com/labora/labora-golang/cinema_labora/models"
)

func CreateUser(newUser models.User) (int, error) {
	insertQuery := "INSERT INTO \"user\"(first_name, last_name, email, password, gender, picture_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	currentTime := time.Now().Unix()
	fields := []interface{}{newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password, newUser.Gender, newUser.PictureURL, currentTime, currentTime}
	requiredFields := []interface{}{newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password, newUser.PictureURL}
	requiredFieldMsg := "FirstName, lastname, email, password, and pictureurl are required fields"
	return CreateDatabaseItem(newUser, insertQuery, fields, requiredFields, requiredFieldMsg)
}

func buildSearchUserQuery(firstName, lastName, email, gender string) string {
	qb := NewQueryBuilder("\"user\"")
	if firstName != "" {
		qb.AddCondition(fmt.Sprintf("first_name ILIKE '%%%s%%'", firstName))
	}
	if lastName != "" {
		qb.AddCondition(fmt.Sprintf("last_name ILIKE '%%%s%%'", lastName))
	}
	if email != "" {
		qb.AddCondition(fmt.Sprintf("email ILIKE '%%%s%%'", email))
	}
	if gender != "" {
		qb.AddCondition(fmt.Sprintf("gender ILIKE '%%%s%%'", gender))
	}
	return qb.BuildQuery()
}

func GetUsers(firstName, lastName, email, gender string) ([]models.User, error) {
	query := buildSearchUserQuery(firstName, lastName, email, gender)
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

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Gender, &user.PictureURL, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id int) (*models.User, error) {
	stmt, err := config.DbConnection.Prepare("SELECT * FROM \"user\" WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Gender, &user.PictureURL, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func buildUpdateUserQuery(updatedUser models.User) (string, error) {
	query := "UPDATE \"user\" SET "
	paramsSize := 1
	if updatedUser.FirstName != nil {
		query += fmt.Sprintf("first_name = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedUser.LastName != nil {
		query += fmt.Sprintf("last_name = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedUser.Email != nil {
		query += fmt.Sprintf("email = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedUser.Password != nil {
		query += fmt.Sprintf("password = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedUser.Gender != nil {
		query += fmt.Sprintf("gender = $%d, ", paramsSize)
		paramsSize++
	}
	if updatedUser.PictureURL != nil {
		query += fmt.Sprintf("picture_url = $%d, ", paramsSize)
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

func getNonNullUserFields(user models.User) []interface{} {
	var nonNullFields []interface{}

	if user.FirstName != nil {
		nonNullFields = append(nonNullFields, user.FirstName)
	}
	if user.LastName != nil {
		nonNullFields = append(nonNullFields, user.LastName)
	}
	if user.Email != nil {
		nonNullFields = append(nonNullFields, user.Email)
	}
	if user.Password != nil {
		nonNullFields = append(nonNullFields, user.Password)
	}
	if user.Gender != nil {
		nonNullFields = append(nonNullFields, user.Gender)
	}
	if user.PictureURL != nil {
		nonNullFields = append(nonNullFields, user.PictureURL)
	}
	nonNullFields = append(nonNullFields, time.Now().Unix())

	return nonNullFields
}

func UpdateUser(id int, updatedUser models.User) error {
	query, err := buildUpdateUserQuery(updatedUser)
	if err != nil {
		return err
	}

	stmt, err := config.DbConnection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	fields := getNonNullUserFields(updatedUser)
	fields = append(fields, id)
	res, err := stmt.Exec(fields...)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("There's no user with the ID %d", id)
	}

	return nil
}

func DeleteUser(id int) error {
	stmt, err := config.DbConnection.Prepare("DELETE FROM \"user\" WHERE id = $1")
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
		return fmt.Errorf("There's no user with the ID %d", id)
	}

	return nil
}
