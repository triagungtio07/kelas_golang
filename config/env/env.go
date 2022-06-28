package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// app
	AppPort string

	// database
	DbHost        string
	DbPort        string
	DbUser        string
	DbPass        string
	DbName        string
	DbAutoMigrate bool
	Debug         bool
	SecretKey     string
)

func Load() {
	godotenv.Load()

	AppPort = os.Getenv("APP_PORT")

	DbName = os.Getenv("DB_NAME")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASSWORD")
	DbAutoMigrate = os.Getenv("DB_AUTO_CREATE") == "true"
	Debug = os.Getenv("DEBUG") == "true"
	SecretKey = os.Getenv("SECRET_KEY")

}
