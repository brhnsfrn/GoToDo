package configs

import (
	"log"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

func SetEnv() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")

	err := godotenv.Load(path.Join(dir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
