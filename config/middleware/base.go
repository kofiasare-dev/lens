package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare-dev/utils"
	"github.com/kofiasare/lens/app/constants"
	"github.com/kofiasare/lens/app/models"
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

	if utils.Blank(authKey) || invalidKey(authKey, c) {
		err = c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authentication Required",
		})

		return
	}

	return c.Next()
}

func invalidKey(apiKey string, c fiber.Ctx) bool {
	key, err := models.ApiKeyWithKey(apiKey)

	if err != nil ||
		key.State == constants.APIKEY_REVOKED ||
		key.Client.State == constants.CLIENT_INACTIVE {
		return true
	}

	c.Locals("APIKEY", key)

	return false
}
