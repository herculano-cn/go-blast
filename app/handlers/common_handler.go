package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Returning a view
func SiteIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Go Blast! Do less, do Betterr!",
	})
}

// Returning a JSON
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "OK",
	})
}
