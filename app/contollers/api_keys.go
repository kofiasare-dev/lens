package contollers

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare-dev/utils"
	"github.com/kofiasare/lens/app/constants"
	"github.com/kofiasare/lens/app/contollers/inputs"
	"github.com/kofiasare/lens/app/models"
	"github.com/kofiasare/lens/app/validators"
)

func ApiKeysSupportedPermissions(c fiber.Ctx) error {
	return c.JSON(utils.Map(constants.SupportedApiKeyPermissions, func(p string) map[string]string {
		return map[string]string{
			"label": strings.ToUpper(strings.ReplaceAll(p, "_", " ")),
			"key":   p,
		}
	}))

}

func ApiKeysCreate(c fiber.Ctx) (err error) {
	var input inputs.CreateApiKeyInput

	if err = c.Bind().JSON(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": validators.Format(err),
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
