package env

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Env struct {
	APP_PORT int
}

func GetEnv() {
	root, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	if filepath.Base(root) == "tests" {
		root = filepath.Dir(root)
	}

	envPath := filepath.Join(root, ".env")

	if err := godotenv.Load(envPath); err != nil {
		panic("No .env file")
	}
}
