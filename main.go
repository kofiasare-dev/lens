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

	models := []interface{}{&models.VerificationRequest{}}

	// migrate models
	// pg.Migrator().DropTable(models...)
	pg.AutoMigrate(models...)

	api.Get("/", contollers.Root)

	v1 := api.Group("/api/v1")
	v1.Post("/verify/face", contollers.VerifyFace)

	api.Use(utils.NotFoundMiddleware)

	defer pg.Close()

	log.Fatal(api.Listen(":8080"))
}
