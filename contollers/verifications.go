package contollers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func VerificationsCreate(c fiber.Ctx) (err error) {
	c.Accepts("application/json")

	err = c.SendStatus(400)

	return
}

func VerificationsStatus(c fiber.Ctx) (err error) {

	log.Info(c.Params("reference"))

	return
}
