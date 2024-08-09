package userroutes

import (
	"hastane-takip/internal/trait"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type PasswordResetRoute struct {
	dig.In
	PasswordResetTrait trait.IPasswordResetHandler
}

func (h *PasswordResetRoute) Handler(c *fiber.Ctx) error {
	type Request struct {
		Phone       string `json:"phone"`
		Code        string `json:"code"`
		NewPassword string `json:"new_password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	err := h.PasswordResetTrait.VerifyResetCode(req.Phone, req.Code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to verify reset code",
		})
	} else {
		if err := h.PasswordResetTrait.ResetPassword(req.Phone, req.NewPassword); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to reset password",
			})
		}
	}
	return c.JSON(fiber.Map{})
}
