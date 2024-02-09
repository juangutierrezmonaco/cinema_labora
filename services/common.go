package services

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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

type QueryBuilder struct {
	Table      string
	Conditions []string
}

func NewQueryBuilder(table string) *QueryBuilder {
	return &QueryBuilder{
		Table: table,
	}
}

func (qb *QueryBuilder) AddCondition(condition string) {
	qb.Conditions = append(qb.Conditions, condition)
}

func (qb *QueryBuilder) BuildQuery() string {
	query := "SELECT * FROM " + qb.Table
	if len(qb.Conditions) > 0 {
		query += " WHERE " + strings.Join(qb.Conditions, " AND ")
	}
	return query
}

func GetDatabaseItems(query string, scanRowFunc func(rows *sql.Rows) (interface{}, error)) ([]interface{}, error) {
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

	var items []interface{}

	for rows.Next() {
		item, err := scanRowFunc(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func GetDatabaseItemByID(id int, tableName string, scanRowFunc func(row *sql.Row) (interface{}, error)) (interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableName)
	stmt, err := config.DbConnection.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	item, err := scanRowFunc(row)
	if err != nil {
		return nil, err
	}

	return item, nil
}
