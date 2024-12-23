package contollers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func Root(c fiber.Ctx) error {
	return c.SendString("OK")
}

func VerifyFace(c fiber.Ctx) error {
	c.Accepts("application/json")

	log.Info(c.IP())

	return c.SendStatus(400)
}
