package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens-api/contollers"
	"github.com/kofiasare/lens-api/db"
	"github.com/kofiasare/lens-api/models"
	"github.com/kofiasare/lens-api/utils"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	api := fiber.New()

	pg := db.GetPgClient()
	defer pg.Close()

	// models
	models := []interface{}{&models.VerificationRequest{}}
	pg.Migrator().DropTable(models...)
	pg.AutoMigrate(models...)

	api.Get("/", contollers.Root)
	api.Post("/api/v1/verifications", contollers.VerificationsCreate)
	api.Get("/api/v1/verifications/:reference", contollers.VerificationsStatus)

	api.Use(utils.NotFoundMiddleware)
	log.Fatal(api.Listen(":8080"))
}
