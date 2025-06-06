package main

import (
	"log"

	"tourlink/migrations"
	"tourlink/pkg/config"
	"tourlink/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := database.ConnectPostgres(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	migrations.RunAutoMigrations()

	app := fiber.New()

	// Setup routes from router.go
	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
