package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens/app/contollers"
	"github.com/kofiasare/lens/config/middleware"
)

func registerAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin", middleware.AdminAuthMiddleware)
	admin.Get("/apikeys/supported_permissions", contollers.ApiKeysSupportedPermissions)
	admin.Post("/apikeys", contollers.ApiKeysCreate)
}
