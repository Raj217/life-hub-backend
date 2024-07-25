package config

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	PORT = "PORT"
)

// InitEnv initializes the .env file
func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}
	PORT = os.Getenv(PORT)
}
