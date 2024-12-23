package utils

import "github.com/gofiber/fiber/v3"

var NotFoundMiddleware = func(c fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
}
