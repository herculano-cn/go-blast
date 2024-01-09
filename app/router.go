package app

import (
	"blast/app/handlers"
	"blast/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func BuildNewRoutes() {
	fiberConfig := fiber.Config{
		Views: config.SetEngine(),
	}

	r := fiber.New(fiberConfig)

	if viper.GetString("ENV") == "development" {
		r.Use(logger.New())
	}

	r.Get("/", handlers.SiteIndex)
	r.Get("/health", handlers.HealthCheck)

	log.Fatal(r.Listen(":" + viper.GetString("server.port")))
}
