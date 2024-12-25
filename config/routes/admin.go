package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens-api/config/middleware"
	"github.com/kofiasare/lens-api/contollers"
)

func registerAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin", middleware.AdminAuthMiddleware)
	admin.Get("/apikeys/supported_permissions", contollers.ApiKeysSupportedPermissions)
	admin.Post("/apikeys", contollers.ApiKeysCreate)
}
