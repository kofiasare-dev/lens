package contollers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens/app/constants"
	"github.com/kofiasare/lens/app/contollers/inputs"
	"github.com/kofiasare/lens/app/models"
	"github.com/kofiasare/lens/app/validators"
)

func FaceMatch(c fiber.Ctx) (err error) {
	key := c.Locals("APIKEY").(*models.ApiKey)
	if !key.Supports(constants.PERMISSION_FACIAL_MATCH) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": constants.ERROR_INSUFFICIENT_PRIVELEGES,
		})
	}

	var fmi inputs.FaceMatchInput

	if err := fmi.Validate(c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": validators.Format(err),
		})
	}

	if err := fmi.UploadImages(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	vr := &models.VerificationRequest{
		Type:        constants.PERMISSION_FACIAL_MATCH,
		Client:      key.Client,
		CallbackURL: fmi.CallbackURL,
		Reference:   fmi.Reference,
	}

	if err = models.CreateVerificationRequest(vr); err != nil {
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
