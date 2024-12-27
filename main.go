package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens/app/models"
	"github.com/kofiasare/lens/app/services/face"
	"github.com/kofiasare/lens/app/validators"
	"github.com/kofiasare/lens/config/routes"
	"github.com/kofiasare/lens/db"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fm := face.GetFaceMatcher()
	defer fm.Close()

	pg := db.GetPgClient()
	defer pg.Close()

	// migrate models
	// pg.Migrator().DropTable(models.Migrations...)
	pg.AutoMigrate(models.Migrations...)

	app := fiber.New(fiber.Config{
		StructValidator: validators.NewValidator(),
	})

	routes.RegisterRoutes(app)
	app.Listen(":8080")
}
