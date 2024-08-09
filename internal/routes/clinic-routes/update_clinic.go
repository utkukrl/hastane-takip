package clinicroutes

import (
	"hastane-takip/internal/models"
	trait "hastane-takip/internal/trait"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type UpdateClinicRoute struct {
	dig.In
	ClinicHandler trait.IClinicHandler
}

func (r *UpdateClinicRoute) Handler(c *fiber.Ctx) error {
	var clinicUpdate models.Clinic
	if err := c.BodyParser(&clinicUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	err = r.ClinicHandler.UpdateClinic(uint(id), &clinicUpdate)
	if err != nil {
		if err.Error() == "clinic record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Clinic record not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update clinic record",
		})
	}

	return c.Status(fiber.StatusOK).JSON(clinicUpdate)
}
