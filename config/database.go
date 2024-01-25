package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labora/labora-golang/cinema_labora/util"
	_ "github.com/lib/pq"
)

func InitDb() {
	fmt.Printf("Initiating connection to the database...\n")
	makeDbConnection()
}

var DbConnection *sql.DB

func makeDbConnection() error {
	dbData := util.EnvData.DbData
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbData.Host, dbData.Port, dbData.RolName, dbData.RolPass, dbData.DbName)

	dbConn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Error while opening the connection to the db. Error: %v\n", err)
	}

	DbConnection = dbConn
	fmt.Printf("Successfully connected to the database '%s'.\n", dbData.DbName)
	if err = DbConnection.Ping(); err != nil {
		DbConnection.Close()
		return err
	}

	return nil
}
