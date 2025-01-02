package contollers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens/app/constants"
	"github.com/kofiasare/lens/app/contollers/inputs"
	"github.com/kofiasare/lens/app/models"
	"github.com/kofiasare/lens/app/validators"
)

func VerificationsCreate(c fiber.Ctx) (err error) {

	// base inputs
	var bi inputs.BaseVerificationInput
	if err = c.Bind().JSON(&bi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": validators.Format(err),
		})
	}

	// specific input
	si := bi.LoadSpecificInput()
	if err = c.Bind().JSON(si); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": validators.Format(err),
		})
	}

	// key authorization
	key := c.Locals("APIKEY").(*models.ApiKey)
	if !key.Supports(bi.Type) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": constants.ERROR_INSUFFICIENT_PRIVELEGES,
		})
	}

	// create
	vr, err := models.CreateVerificationRequest(key.Client, si)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"verificationRequest": vr})
}

func VerificationsShow(c fiber.Ctx) (err error) {
	v, err := models.FindVerificationByReference(c.Params("reference"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"verificationRequest": v})
}
