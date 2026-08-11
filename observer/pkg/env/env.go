package env

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Env struct {
	APP_PORT int
}

func getEnv(base_folder string) error {
	root, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	if filepath.Base(root) == "tests" {
		root = filepath.Dir(root)
	}

	envPath := filepath.Join(root, base_folder, ".env")

	if err := godotenv.Load(envPath); err != nil {
		return err
	}
	return nil
}
func GetEnv() {
	var err error
	err = getEnv(".")
	if err == nil {
		return
	}

	err = getEnv("../..")
	if err == nil {
		return
	}
	panic(err)
}
