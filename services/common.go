package services

import (
	"errors"
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