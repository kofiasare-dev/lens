package contollers

import (
	"github.com/gofiber/fiber/v3"
)

func Root(c fiber.Ctx) (err error) {
	err = c.SendString("OK")

	return
}
