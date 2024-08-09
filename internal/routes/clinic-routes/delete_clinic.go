package clinicroutes

import (
	trait "hastane-takip/internal/trait"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type DeleteClinicRoute struct {
	dig.In
	ClinicHandler trait.IClinicHandler
}

func (r *DeleteClinicRoute) Handler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	err = r.ClinicHandler.DeleteClinic(uint(id))
	if err != nil {
		if err.Error() == "clinic record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Clinic record not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete clinic record",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Clinic record deleted successfully",
	})
}
