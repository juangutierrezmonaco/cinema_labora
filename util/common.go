package util

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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
	DbData      DbData
	ServerData  ServerData
	MovieDbData MovieDbData
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
		MovieDbData: MovieDbData{
			ApiKey: os.Getenv("TMDB_API_KEY"),
		},
	}

	return nil
}

func HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func PasswordMatch(password string, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

func ParseTime(unixTime int64) string {
	time := time.Unix(unixTime, 0)
	timeStr := time.Format("02-01-2006 15:04:05")
	return timeStr
}
