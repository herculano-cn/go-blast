package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error, errCode int) error {
	return c.Status(errCode).JSON(fiber.Map{
		"message": err.Error(),
	})
}
