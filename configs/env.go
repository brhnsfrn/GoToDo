package configs

import (
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

func SetEnv() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")

	godotenv.Load(dir + "\\.env")
}
