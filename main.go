package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens-api/config/routes"
	"github.com/kofiasare/lens-api/db"
	"github.com/kofiasare/lens-api/models"
	"github.com/kofiasare/lens-api/utils"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	pg := db.GetPgClient()
	defer pg.Close()

	models := []interface{}{
		&models.Client{},
		&models.ApiKey{},
		&models.VerificationRequest{},
	}

	// migrate models
	// pg.Migrator().DropTable(models...)
	pg.AutoMigrate(models...)

	app := fiber.New(fiber.Config{
		StructValidator: utils.NewValidator(),
	})

	routes.RegisterRoutes(app)

	app.Listen(":8080")
}
