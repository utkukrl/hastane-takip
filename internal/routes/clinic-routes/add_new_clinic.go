package clinicroutes

import (
	"hastane-takip/internal/models"
	trait "hastane-takip/internal/trait"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type AddNewClinicRoute struct {
	dig.In
	ClinicHandler trait.IClinicHandler
}

func (r *AddNewClinicRoute) Handler(c *fiber.Ctx) error {
	var addedClinic models.Clinic

	if err := c.BodyParser(&addedClinic); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := r.ClinicHandler.CreateClinic(&addedClinic)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not add clinic",
		})
	}
	return c.JSON(fiber.Map{})

}
