package env

import (
	"path/filepath"

	"github.com/joho/godotenv"
)

type Env struct {
	APP_PORT int
}

func GetEnv() {
	envPath := filepath.Join(".", ".env")

	if err := godotenv.Load(envPath); err != nil {
		panic("No .env file")
	}
}
