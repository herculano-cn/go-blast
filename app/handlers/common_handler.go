package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Returning a view
func SiteIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "Just a web framework, to have a blast!",
		"Description": "High performance, less latency, minimalist Go web framework",
	}, "layouts/main")
}

// Returning a JSON
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "OK",
	})
}
