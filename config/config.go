package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
	MONGO_URI string
	POSTGRES_URI string
)

func InitConfig() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	PORT = getEnv("PORT", "3000")
	MONGO_URI = getEnv("MONGO_URI", "mongodb://localhost:27017/DBGPT")
	POSTGRES_URI = getEnv("POSTGRES_URI", "postgres://postgres:postgres@localhost:5432/postgres")
}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return fallback
}