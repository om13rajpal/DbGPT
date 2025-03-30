package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT           string
	MONGO_URI      string
	POSTGRES_URI   string
	ISS            string
	JWT_SECRET     string
	EMAIL          string
	EMAIL_PASSWORD string
)

func InitConfig() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	PORT = getEnv("PORT", "3000")
	MONGO_URI = getEnv("MONGO_URI", "mongodb://localhost:27017/DBGPT")
	POSTGRES_URI = getEnv("POSTGRES_URI", "postgres://postgres:postgres@localhost:5432/postgres")
	ISS = getEnv("ISS", "dbgpt")
	JWT_SECRET = getEnv("JWT_SECRET", "dbgpt")
	EMAIL = getEnv("EMAIL", "")
	EMAIL_PASSWORD = getEnv("EMAIL_PASSWORD", "")

}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return fallback
}
