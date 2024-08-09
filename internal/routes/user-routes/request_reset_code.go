package userroutes

import (
	"hastane-takip/internal/trait"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type RequestResetCodeRoute struct {
	dig.In
	PasswordResetTrait trait.IPasswordResetHandler
}

func (h *RequestResetCodeRoute) RequestResetCode(c *fiber.Ctx) error {
	type Request struct {
		Phone string `json:"phone"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	_, err := h.PasswordResetTrait.RequestResetCode(req.Phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create reset code",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Reset code generated",
	})
}
