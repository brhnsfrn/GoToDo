package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func Authorize() fiber.Handler {
	err := basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("USER"): os.Getenv("PASSWD"),
		},
	})
	return err
}
