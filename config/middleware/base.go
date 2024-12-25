package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/kofiasare/lens-api/models"
	"github.com/kofiasare/lens-api/utils"
)

var NotFoundMiddleware = func(c fiber.Ctx) (err error) {
	err = c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")

	return
}

var AdminAuthMiddleware = func(c fiber.Ctx) (err error) {
	authKey := c.Get("Authorization")
	expectedKey := os.Getenv("ADMIN_KEY")

	if utils.Blank(authKey) || authKey != expectedKey {
		err = c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authentication Required",
		})

		return
	}

	return c.Next()
}

var ApiAuthMiddleware = func(c fiber.Ctx) (err error) {
	authKey := c.Get("Authorization")

	if utils.Blank(authKey) || !isValidApiKey(authKey) {
		err = c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authentication Required",
		})

		return
	}

	return c.Next()
}

func isValidApiKey(apiKey string) bool {
	key, err := models.ApiKeyWithKey(apiKey)

	if err != nil {
		log.Info("Error:", err)
	}

	log.Info("Key found:", key)

	return false
}
