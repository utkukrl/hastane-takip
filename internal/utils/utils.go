package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	log.Printf("Error: %v", err)
	statusCode := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		statusCode = e.Code
	}
	return c.Status(statusCode).JSON(fiber.Map{
		"error": err.Error(),
	})
}
