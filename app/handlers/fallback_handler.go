package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func hasError(c *fiber.Ctx, err error, errCode int) bool {
	if err != nil {
		c.Status(errCode).JSON(fiber.Map{
			"message": err.Error(),
		})
		return true
	}
	return false
}
