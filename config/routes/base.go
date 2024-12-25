package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens-api/config/middleware"
	"github.com/kofiasare/lens-api/contollers"
)

func RegisterRoutes(app *fiber.App) {

	// Health check
	app.Get("/up", contollers.Root)

	registerAdminRoutes(app)
	registerApiV1Routes(app)

	// middlewares
	app.Use(middleware.NotFoundMiddleware)
}
