package config

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/kofiasare/lens/app/constants"
	"github.com/kofiasare/lens/app/contollers"
	"github.com/kofiasare/lens/app/models"
	"github.com/kofiasare/lens/app/services/bg"
	"github.com/kofiasare/lens/config/middleware"
)

func HookRoutesTo(app *fiber.App) {
	app.Get("/up", contollers.Root)

	registerAdminRoutes(app)
	registerApiV1Routes(app)
	setupTaskHandlers(app)

	// middlewares
	app.Use(middleware.NotFoundMiddleware)
}

func registerAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin", middleware.AdminAuthMiddleware)
	admin.Get("/apikeys/supported_permissions", contollers.ApiKeysSupportedPermissions)
	admin.Post("/apikeys", contollers.ApiKeysCreate)
}

func registerApiV1Routes(app *fiber.App) {
	v1 := app.Group("/api/v1", middleware.ApiAuthMiddleware)
	v1.Post("/face-match", contollers.FaceMatch)
	v1.Get("/verifications/:reference", contollers.VerificationsShow)
}

func setupTaskHandlers(app *fiber.App) {
	bg := bg.GetInstance()

	if bg.Web != nil {
		app.Get("monitoring/*", adaptor.HTTPHandler(bg.Web))
	}

	bg.Mux.HandleFunc(constants.VERIFICATION_REQUEST_PROCESSING_TASK,
		func(c context.Context, t *asynq.Task) (err error) {
			var ge models.GuaranteedExecution
			if err = json.Unmarshal(t.Payload(), &ge); err != nil {
				return
			}

			vr, err := models.FindVerificationRequest(uuid.MustParse(ge.GuaranteeableID))
			if err != nil {
				return
			}

			vr.ProcessVerificationRequest()

			return
		},
	)

}
