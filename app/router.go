package app

import (
	"blast/app/handlers"
	"blast/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func BuildNewRoutes() {
	fiberConfig := fiber.Config{
		Views: config.SetEngine(),
	}

	r := fiber.New(fiberConfig)

	if config.ConfigName == "development" {
		r.Use(logger.New())
	}

	r.Get("/", handlers.SiteIndex)
	r.Get("/health", handlers.HealthCheck)

	log.Fatal(r.Listen(":" + os.Getenv("SERVER_PORT")))
}
