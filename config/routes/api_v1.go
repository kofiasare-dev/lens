package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens/app/contollers"
	"github.com/kofiasare/lens/config/middleware"
)

func registerApiV1Routes(app *fiber.App) {
	v1 := app.Group("/api/v1", middleware.ApiAuthMiddleware)
	v1.Post("/verifications", contollers.VerificationsCreate)
	v1.Get("/verifications/:reference", contollers.VerificationsShow)
}
