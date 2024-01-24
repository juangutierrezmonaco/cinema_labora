package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbData struct {
	Host    string
	Port    string
	DbName  string
	RolName string
	RolPass string
}

type ServerData struct {
	Host string
	Port string
}

type MovieDbData struct {
	ApiKey string
}

type AllEnvData struct {
	DbData     DbData
	ServerData ServerData
}

var EnvData AllEnvData

func LoadEnv() error {
	godotenv.Load("./.env")
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error while loading env variables. Error: %v\n", err)
	}

	EnvData = AllEnvData{
		DbData: DbData{
			Host:    os.Getenv("DB_HOST"),
			Port:    os.Getenv("DB_PORT"),
			DbName:  os.Getenv("DB_NAME"),
			RolName: os.Getenv("ROL_NAME"),
			RolPass: os.Getenv("ROL_PASS"),
		},
		ServerData: ServerData{
			Host: os.Getenv("HOST"),
			Port: os.Getenv("PORT"),
		},
	}

	return nil
}
