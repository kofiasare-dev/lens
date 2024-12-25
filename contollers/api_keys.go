package contollers

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens-api/contollers/inputs"
	"github.com/kofiasare/lens-api/models"
	"github.com/kofiasare/lens-api/utils"
)

func ApiKeysSupportedPermissions(c fiber.Ctx) error {
	return c.JSON(utils.Map(models.SupportedApiKeyPermissions, func(p string) map[string]string {
		return map[string]string{
			"key":   strings.ToLower(strings.ReplaceAll(p, " ", "_")),
			"label": p,
		}
	}))

}

func ApiKeysCreate(c fiber.Ctx) (err error) {
	var input inputs.CreateApiKeyInput

	if err = c.Bind().JSON(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": utils.FormatValidationError(err),
		})
	}

	apiKey, err := models.CreateApiKey(&input, utils.SecureRandomHex(16))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"key": apiKey})
}
